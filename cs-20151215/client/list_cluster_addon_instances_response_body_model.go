// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterAddonInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddons(v []*ListClusterAddonInstancesResponseBodyAddons) *ListClusterAddonInstancesResponseBody
	GetAddons() []*ListClusterAddonInstancesResponseBodyAddons
}

type ListClusterAddonInstancesResponseBody struct {
	// A list of components that are installed in the cluster.
	Addons []*ListClusterAddonInstancesResponseBodyAddons `json:"addons,omitempty" xml:"addons,omitempty" type:"Repeated"`
}

func (s ListClusterAddonInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterAddonInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterAddonInstancesResponseBody) GetAddons() []*ListClusterAddonInstancesResponseBodyAddons {
	return s.Addons
}

func (s *ListClusterAddonInstancesResponseBody) SetAddons(v []*ListClusterAddonInstancesResponseBodyAddons) *ListClusterAddonInstancesResponseBody {
	s.Addons = v
	return s
}

func (s *ListClusterAddonInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListClusterAddonInstancesResponseBodyAddons struct {
	// The component name.
	//
	// example:
	//
	// coredns
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the component. Valid values:
	//
	// 	- active: The component is installed.
	//
	// 	- updating: The component is being modified.
	//
	// 	- upgrading: The component is being updated.
	//
	// 	- deleting: The component is being uninstalled.
	//
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The version of the component.
	//
	// example:
	//
	// v1.9.3.10-7dfca203-aliyun
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListClusterAddonInstancesResponseBodyAddons) String() string {
	return dara.Prettify(s)
}

func (s ListClusterAddonInstancesResponseBodyAddons) GoString() string {
	return s.String()
}

func (s *ListClusterAddonInstancesResponseBodyAddons) GetName() *string {
	return s.Name
}

func (s *ListClusterAddonInstancesResponseBodyAddons) GetState() *string {
	return s.State
}

func (s *ListClusterAddonInstancesResponseBodyAddons) GetVersion() *string {
	return s.Version
}

func (s *ListClusterAddonInstancesResponseBodyAddons) SetName(v string) *ListClusterAddonInstancesResponseBodyAddons {
	s.Name = &v
	return s
}

func (s *ListClusterAddonInstancesResponseBodyAddons) SetState(v string) *ListClusterAddonInstancesResponseBodyAddons {
	s.State = &v
	return s
}

func (s *ListClusterAddonInstancesResponseBodyAddons) SetVersion(v string) *ListClusterAddonInstancesResponseBodyAddons {
	s.Version = &v
	return s
}

func (s *ListClusterAddonInstancesResponseBodyAddons) Validate() error {
	return dara.Validate(s)
}
