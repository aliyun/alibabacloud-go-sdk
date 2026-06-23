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
	// The component configuration.
	//
	// example:
	//
	// {"NetworkPolicy":"true"}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// The component name.
	//
	// example:
	//
	// terway-eniip
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The component status. Valid values:
	//
	// - initial: installing
	//
	// - active: installed
	//
	// - unhealthy: abnormal
	//
	// - upgrading: upgrading
	//
	// - updating: updating
	//
	// - deleting: uninstalling
	//
	// - deleted: deleted.
	//
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The component version.
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
