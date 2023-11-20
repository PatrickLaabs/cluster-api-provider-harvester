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

package v3

import (
	"context"
	"time"

	scheme "github.com/rancher-sandbox/cluster-api-provider-harvester/pkg/clientset/versioned/scheme"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterLoggingsGetter has a method to return a ClusterLoggingInterface.
// A group's client should implement this interface.
type ClusterLoggingsGetter interface {
	ClusterLoggings(namespace string) ClusterLoggingInterface
}

// ClusterLoggingInterface has methods to work with ClusterLogging resources.
type ClusterLoggingInterface interface {
	Create(ctx context.Context, clusterLogging *v3.ClusterLogging, opts v1.CreateOptions) (*v3.ClusterLogging, error)
	Update(ctx context.Context, clusterLogging *v3.ClusterLogging, opts v1.UpdateOptions) (*v3.ClusterLogging, error)
	UpdateStatus(ctx context.Context, clusterLogging *v3.ClusterLogging, opts v1.UpdateOptions) (*v3.ClusterLogging, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.ClusterLogging, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.ClusterLoggingList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.ClusterLogging, err error)
	ClusterLoggingExpansion
}

// clusterLoggings implements ClusterLoggingInterface
type clusterLoggings struct {
	client rest.Interface
	ns     string
}

// newClusterLoggings returns a ClusterLoggings
func newClusterLoggings(c *ManagementV3Client, namespace string) *clusterLoggings {
	return &clusterLoggings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterLogging, and returns the corresponding clusterLogging object, and an error if there is any.
func (c *clusterLoggings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.ClusterLogging, err error) {
	result = &v3.ClusterLogging{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterloggings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterLoggings that match those selectors.
func (c *clusterLoggings) List(ctx context.Context, opts v1.ListOptions) (result *v3.ClusterLoggingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.ClusterLoggingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterloggings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterLoggings.
func (c *clusterLoggings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clusterloggings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterLogging and creates it.  Returns the server's representation of the clusterLogging, and an error, if there is any.
func (c *clusterLoggings) Create(ctx context.Context, clusterLogging *v3.ClusterLogging, opts v1.CreateOptions) (result *v3.ClusterLogging, err error) {
	result = &v3.ClusterLogging{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clusterloggings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterLogging).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterLogging and updates it. Returns the server's representation of the clusterLogging, and an error, if there is any.
func (c *clusterLoggings) Update(ctx context.Context, clusterLogging *v3.ClusterLogging, opts v1.UpdateOptions) (result *v3.ClusterLogging, err error) {
	result = &v3.ClusterLogging{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterloggings").
		Name(clusterLogging.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterLogging).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterLoggings) UpdateStatus(ctx context.Context, clusterLogging *v3.ClusterLogging, opts v1.UpdateOptions) (result *v3.ClusterLogging, err error) {
	result = &v3.ClusterLogging{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterloggings").
		Name(clusterLogging.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterLogging).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterLogging and deletes it. Returns an error if one occurs.
func (c *clusterLoggings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterloggings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterLoggings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterloggings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterLogging.
func (c *clusterLoggings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.ClusterLogging, err error) {
	result = &v3.ClusterLogging{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clusterloggings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
