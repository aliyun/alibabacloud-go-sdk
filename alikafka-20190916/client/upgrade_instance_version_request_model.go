// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeInstanceVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpgradeInstanceVersionRequest
	GetInstanceId() *string
	SetRegionId(v string) *UpgradeInstanceVersionRequest
	GetRegionId() *string
	SetTargetVersion(v string) *UpgradeInstanceVersionRequest
	GetTargetVersion() *string
}

type UpgradeInstanceVersionRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The major version to be upgraded to. Valid values:
	//
	// 	- **0.10.2**
	//
	// 	- **2.2.0**
	//
	// If you set this parameter to the current major version, the system upgrades the instance to the latest minor version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.10.2
	TargetVersion *string `json:"TargetVersion,omitempty" xml:"TargetVersion,omitempty"`
}

func (s UpgradeInstanceVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeInstanceVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceVersionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpgradeInstanceVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradeInstanceVersionRequest) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *UpgradeInstanceVersionRequest) SetInstanceId(v string) *UpgradeInstanceVersionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeInstanceVersionRequest) SetRegionId(v string) *UpgradeInstanceVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeInstanceVersionRequest) SetTargetVersion(v string) *UpgradeInstanceVersionRequest {
	s.TargetVersion = &v
	return s
}

func (s *UpgradeInstanceVersionRequest) Validate() error {
	return dara.Validate(s)
}
