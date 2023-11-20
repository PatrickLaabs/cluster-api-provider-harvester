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

// FakeClusterScans implements ClusterScanInterface
type FakeClusterScans struct {
	Fake *FakeManagementV3
	ns   string
}

var clusterscansResource = schema.GroupVersionResource{Group: "management.cattle.io", Version: "v3", Resource: "clusterscans"}

var clusterscansKind = schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterScan"}

// Get takes name of the clusterScan, and returns the corresponding clusterScan object, and an error if there is any.
func (c *FakeClusterScans) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.ClusterScan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clusterscansResource, c.ns, name), &v3.ClusterScan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterScan), err
}

// List takes label and field selectors, and returns the list of ClusterScans that match those selectors.
func (c *FakeClusterScans) List(ctx context.Context, opts v1.ListOptions) (result *v3.ClusterScanList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clusterscansResource, clusterscansKind, c.ns, opts), &v3.ClusterScanList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.ClusterScanList{ListMeta: obj.(*v3.ClusterScanList).ListMeta}
	for _, item := range obj.(*v3.ClusterScanList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterScans.
func (c *FakeClusterScans) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clusterscansResource, c.ns, opts))

}

// Create takes the representation of a clusterScan and creates it.  Returns the server's representation of the clusterScan, and an error, if there is any.
func (c *FakeClusterScans) Create(ctx context.Context, clusterScan *v3.ClusterScan, opts v1.CreateOptions) (result *v3.ClusterScan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clusterscansResource, c.ns, clusterScan), &v3.ClusterScan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterScan), err
}

// Update takes the representation of a clusterScan and updates it. Returns the server's representation of the clusterScan, and an error, if there is any.
func (c *FakeClusterScans) Update(ctx context.Context, clusterScan *v3.ClusterScan, opts v1.UpdateOptions) (result *v3.ClusterScan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clusterscansResource, c.ns, clusterScan), &v3.ClusterScan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterScan), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterScans) UpdateStatus(ctx context.Context, clusterScan *v3.ClusterScan, opts v1.UpdateOptions) (*v3.ClusterScan, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clusterscansResource, "status", c.ns, clusterScan), &v3.ClusterScan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterScan), err
}

// Delete takes name of the clusterScan and deletes it. Returns an error if one occurs.
func (c *FakeClusterScans) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(clusterscansResource, c.ns, name, opts), &v3.ClusterScan{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterScans) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clusterscansResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v3.ClusterScanList{})
	return err
}

// Patch applies the patch and returns the patched clusterScan.
func (c *FakeClusterScans) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.ClusterScan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clusterscansResource, c.ns, name, pt, data, subresources...), &v3.ClusterScan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterScan), err
}
