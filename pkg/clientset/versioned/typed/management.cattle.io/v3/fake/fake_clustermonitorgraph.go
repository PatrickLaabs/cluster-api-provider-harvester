/*
Copyright 2023 Rancher Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterMonitorGraphs implements ClusterMonitorGraphInterface
type FakeClusterMonitorGraphs struct {
	Fake *FakeManagementV3
	ns   string
}

var clustermonitorgraphsResource = schema.GroupVersionResource{Group: "management.cattle.io", Version: "v3", Resource: "clustermonitorgraphs"}

var clustermonitorgraphsKind = schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterMonitorGraph"}

// Get takes name of the clusterMonitorGraph, and returns the corresponding clusterMonitorGraph object, and an error if there is any.
func (c *FakeClusterMonitorGraphs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.ClusterMonitorGraph, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clustermonitorgraphsResource, c.ns, name), &v3.ClusterMonitorGraph{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterMonitorGraph), err
}

// List takes label and field selectors, and returns the list of ClusterMonitorGraphs that match those selectors.
func (c *FakeClusterMonitorGraphs) List(ctx context.Context, opts v1.ListOptions) (result *v3.ClusterMonitorGraphList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clustermonitorgraphsResource, clustermonitorgraphsKind, c.ns, opts), &v3.ClusterMonitorGraphList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.ClusterMonitorGraphList{ListMeta: obj.(*v3.ClusterMonitorGraphList).ListMeta}
	for _, item := range obj.(*v3.ClusterMonitorGraphList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterMonitorGraphs.
func (c *FakeClusterMonitorGraphs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clustermonitorgraphsResource, c.ns, opts))

}

// Create takes the representation of a clusterMonitorGraph and creates it.  Returns the server's representation of the clusterMonitorGraph, and an error, if there is any.
func (c *FakeClusterMonitorGraphs) Create(ctx context.Context, clusterMonitorGraph *v3.ClusterMonitorGraph, opts v1.CreateOptions) (result *v3.ClusterMonitorGraph, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clustermonitorgraphsResource, c.ns, clusterMonitorGraph), &v3.ClusterMonitorGraph{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterMonitorGraph), err
}

// Update takes the representation of a clusterMonitorGraph and updates it. Returns the server's representation of the clusterMonitorGraph, and an error, if there is any.
func (c *FakeClusterMonitorGraphs) Update(ctx context.Context, clusterMonitorGraph *v3.ClusterMonitorGraph, opts v1.UpdateOptions) (result *v3.ClusterMonitorGraph, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clustermonitorgraphsResource, c.ns, clusterMonitorGraph), &v3.ClusterMonitorGraph{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterMonitorGraph), err
}

// Delete takes name of the clusterMonitorGraph and deletes it. Returns an error if one occurs.
func (c *FakeClusterMonitorGraphs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(clustermonitorgraphsResource, c.ns, name, opts), &v3.ClusterMonitorGraph{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterMonitorGraphs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clustermonitorgraphsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v3.ClusterMonitorGraphList{})
	return err
}

// Patch applies the patch and returns the patched clusterMonitorGraph.
func (c *FakeClusterMonitorGraphs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.ClusterMonitorGraph, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clustermonitorgraphsResource, c.ns, name, pt, data, subresources...), &v3.ClusterMonitorGraph{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterMonitorGraph), err
}
