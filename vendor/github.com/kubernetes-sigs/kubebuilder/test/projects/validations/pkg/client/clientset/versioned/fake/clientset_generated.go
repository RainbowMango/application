// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "github.com/kubernetes-sigs/kubebuilder/test/projects/validations/pkg/client/clientset/versioned"
	appsv1 "github.com/kubernetes-sigs/kubebuilder/test/projects/validations/pkg/client/clientset/versioned/typed/apps/v1"
	fakeappsv1 "github.com/kubernetes-sigs/kubebuilder/test/projects/validations/pkg/client/clientset/versioned/typed/apps/v1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	fakePtr := testing.Fake{}
	fakePtr.AddReactor("*", "*", testing.ObjectReaction(o))
	fakePtr.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return &Clientset{fakePtr, &fakediscovery.FakeDiscovery{Fake: &fakePtr}}
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

var _ clientset.Interface = &Clientset{}

// AppsV1 retrieves the AppsV1Client
func (c *Clientset) AppsV1() appsv1.AppsV1Interface {
	return &fakeappsv1.FakeAppsV1{Fake: &c.Fake}
}

// Apps retrieves the AppsV1Client
func (c *Clientset) Apps() appsv1.AppsV1Interface {
	return &fakeappsv1.FakeAppsV1{Fake: &c.Fake}
}
