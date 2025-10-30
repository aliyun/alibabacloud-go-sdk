// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeClusterAddonsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v []*UpgradeClusterAddonsRequestBody) *UpgradeClusterAddonsRequest
	GetBody() []*UpgradeClusterAddonsRequestBody
}

type UpgradeClusterAddonsRequest struct {
	// The request parameters.
	Body []*UpgradeClusterAddonsRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s UpgradeClusterAddonsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeClusterAddonsRequest) GoString() string {
	return s.String()
}

func (s *UpgradeClusterAddonsRequest) GetBody() []*UpgradeClusterAddonsRequestBody {
	return s.Body
}

func (s *UpgradeClusterAddonsRequest) SetBody(v []*UpgradeClusterAddonsRequestBody) *UpgradeClusterAddonsRequest {
	s.Body = v
	return s
}

func (s *UpgradeClusterAddonsRequest) Validate() error {
	if s.Body != nil {
		for _, item := range s.Body {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpgradeClusterAddonsRequestBody struct {
	// The name of the component.
	//
	// This parameter is required.
	//
	// example:
	//
	// coredns
	ComponentName *string `json:"component_name,omitempty" xml:"component_name,omitempty"`
	// The custom component settings that you want to use. The value is a JSON string.
	//
	// example:
	//
	// {\\"CpuRequest\\":\\"800m\\"}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// The version to which the component can be updated. You can call the `DescribeClusterAddonsVersion` operation to query the version to which the component can be updated.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.6.7
	NextVersion *string `json:"next_version,omitempty" xml:"next_version,omitempty"`
	// The update policy. Valid values:
	//
	// 	- overwrite
	//
	// 	- canary
	//
	// example:
	//
	// canary
	Policy *string `json:"policy,omitempty" xml:"policy,omitempty"`
	// The current version of the component.
	//
	// example:
	//
	// v1.6.2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpgradeClusterAddonsRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeClusterAddonsRequestBody) GoString() string {
	return s.String()
}

func (s *UpgradeClusterAddonsRequestBody) GetComponentName() *string {
	return s.ComponentName
}

func (s *UpgradeClusterAddonsRequestBody) GetConfig() *string {
	return s.Config
}

func (s *UpgradeClusterAddonsRequestBody) GetNextVersion() *string {
	return s.NextVersion
}

func (s *UpgradeClusterAddonsRequestBody) GetPolicy() *string {
	return s.Policy
}

func (s *UpgradeClusterAddonsRequestBody) GetVersion() *string {
	return s.Version
}

func (s *UpgradeClusterAddonsRequestBody) SetComponentName(v string) *UpgradeClusterAddonsRequestBody {
	s.ComponentName = &v
	return s
}

func (s *UpgradeClusterAddonsRequestBody) SetConfig(v string) *UpgradeClusterAddonsRequestBody {
	s.Config = &v
	return s
}

func (s *UpgradeClusterAddonsRequestBody) SetNextVersion(v string) *UpgradeClusterAddonsRequestBody {
	s.NextVersion = &v
	return s
}

func (s *UpgradeClusterAddonsRequestBody) SetPolicy(v string) *UpgradeClusterAddonsRequestBody {
	s.Policy = &v
	return s
}

func (s *UpgradeClusterAddonsRequestBody) SetVersion(v string) *UpgradeClusterAddonsRequestBody {
	s.Version = &v
	return s
}

func (s *UpgradeClusterAddonsRequestBody) Validate() error {
	return dara.Validate(s)
}
