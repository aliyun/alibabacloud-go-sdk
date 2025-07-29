// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterAddonInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *DescribeClusterAddonInstanceResponseBody
	GetConfig() *string
	SetName(v string) *DescribeClusterAddonInstanceResponseBody
	GetName() *string
	SetState(v string) *DescribeClusterAddonInstanceResponseBody
	GetState() *string
	SetVersion(v string) *DescribeClusterAddonInstanceResponseBody
	GetVersion() *string
}

type DescribeClusterAddonInstanceResponseBody struct {
	// The configuration of the component.
	//
	// example:
	//
	// {"NetworkPolicy":"true"}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// The name of the component.
	//
	// example:
	//
	// terway-eniip
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the component. Valid values:
	//
	// 	- initial: The component is being installed.
	//
	// 	- active: The component has been installed.
	//
	// 	- unhealthy: The component is in an abnormal state.
	//
	// 	- upgrading: The component is undergoing an upgrade.
	//
	// 	- updating: Component configuration changes are being applied.
	//
	// 	- deleting: The component is being uninstalled.
	//
	// 	- deleted: The component has been deleted.
	//
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The version of the component.
	//
	// example:
	//
	// v1.4.3
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeClusterAddonInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAddonInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonInstanceResponseBody) GetConfig() *string {
	return s.Config
}

func (s *DescribeClusterAddonInstanceResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeClusterAddonInstanceResponseBody) GetState() *string {
	return s.State
}

func (s *DescribeClusterAddonInstanceResponseBody) GetVersion() *string {
	return s.Version
}

func (s *DescribeClusterAddonInstanceResponseBody) SetConfig(v string) *DescribeClusterAddonInstanceResponseBody {
	s.Config = &v
	return s
}

func (s *DescribeClusterAddonInstanceResponseBody) SetName(v string) *DescribeClusterAddonInstanceResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeClusterAddonInstanceResponseBody) SetState(v string) *DescribeClusterAddonInstanceResponseBody {
	s.State = &v
	return s
}

func (s *DescribeClusterAddonInstanceResponseBody) SetVersion(v string) *DescribeClusterAddonInstanceResponseBody {
	s.Version = &v
	return s
}

func (s *DescribeClusterAddonInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
