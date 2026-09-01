// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentlessScanTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetSelectionType(v string) *CreateAgentlessScanTaskRequest
	GetAssetSelectionType() *string
	SetAutoDeleteDays(v int32) *CreateAgentlessScanTaskRequest
	GetAutoDeleteDays() *int32
	SetClientToken(v string) *CreateAgentlessScanTaskRequest
	GetClientToken() *string
	SetRegionId(v string) *CreateAgentlessScanTaskRequest
	GetRegionId() *string
	SetReleaseAfterScan(v bool) *CreateAgentlessScanTaskRequest
	GetReleaseAfterScan() *bool
	SetScanDataDisk(v bool) *CreateAgentlessScanTaskRequest
	GetScanDataDisk() *bool
	SetTargetType(v int32) *CreateAgentlessScanTaskRequest
	GetTargetType() *int32
	SetUuidList(v []*string) *CreateAgentlessScanTaskRequest
	GetUuidList() []*string
}

type CreateAgentlessScanTaskRequest struct {
	// The asset selection identifier.
	//
	// example:
	//
	// AGENTLESS_SCAN_ONCE_TASK_1720145******
	AssetSelectionType *string `json:"AssetSelectionType,omitempty" xml:"AssetSelectionType,omitempty"`
	// The image retention period, in days. This parameter takes effect only for host detection and does not take effect for user snapshot detection or user custom image detection.
	//
	// example:
	//
	// 1
	AutoDeleteDays *int32 `json:"AutoDeleteDays,omitempty" xml:"AutoDeleteDays,omitempty"`
	// The idempotency key.
	//
	// example:
	//
	// 66a9c708-d4a4-4fe
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the instance to query. Valid values:
	//
	// - **cn-hangzhou*	- (default): China.
	//
	// - **ap-southeast-1**: outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to enable the cost-saving mode. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	ReleaseAfterScan *bool `json:"ReleaseAfterScan,omitempty" xml:"ReleaseAfterScan,omitempty"`
	// Specifies whether to detect data cloud disks. Valid values:
	//
	// - **true**: Detected.
	//
	// - **false**: Not detected.
	//
	// example:
	//
	// true
	ScanDataDisk *bool `json:"ScanDataDisk,omitempty" xml:"ScanDataDisk,omitempty"`
	// The target type. Valid values:
	//
	// - **1**: Host detection - detection by snapshot.
	//
	// - **2**: Host detection - detection by image.
	//
	// - **3**: User snapshot detection.
	//
	// - **2**: User custom image detection.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	TargetType *int32 `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The UUIDs of the assets to be detected.
	//
	// > You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to obtain the UUIDs of servers.
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s CreateAgentlessScanTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentlessScanTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentlessScanTaskRequest) GetAssetSelectionType() *string {
	return s.AssetSelectionType
}

func (s *CreateAgentlessScanTaskRequest) GetAutoDeleteDays() *int32 {
	return s.AutoDeleteDays
}

func (s *CreateAgentlessScanTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAgentlessScanTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAgentlessScanTaskRequest) GetReleaseAfterScan() *bool {
	return s.ReleaseAfterScan
}

func (s *CreateAgentlessScanTaskRequest) GetScanDataDisk() *bool {
	return s.ScanDataDisk
}

func (s *CreateAgentlessScanTaskRequest) GetTargetType() *int32 {
	return s.TargetType
}

func (s *CreateAgentlessScanTaskRequest) GetUuidList() []*string {
	return s.UuidList
}

func (s *CreateAgentlessScanTaskRequest) SetAssetSelectionType(v string) *CreateAgentlessScanTaskRequest {
	s.AssetSelectionType = &v
	return s
}

func (s *CreateAgentlessScanTaskRequest) SetAutoDeleteDays(v int32) *CreateAgentlessScanTaskRequest {
	s.AutoDeleteDays = &v
	return s
}

func (s *CreateAgentlessScanTaskRequest) SetClientToken(v string) *CreateAgentlessScanTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAgentlessScanTaskRequest) SetRegionId(v string) *CreateAgentlessScanTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAgentlessScanTaskRequest) SetReleaseAfterScan(v bool) *CreateAgentlessScanTaskRequest {
	s.ReleaseAfterScan = &v
	return s
}

func (s *CreateAgentlessScanTaskRequest) SetScanDataDisk(v bool) *CreateAgentlessScanTaskRequest {
	s.ScanDataDisk = &v
	return s
}

func (s *CreateAgentlessScanTaskRequest) SetTargetType(v int32) *CreateAgentlessScanTaskRequest {
	s.TargetType = &v
	return s
}

func (s *CreateAgentlessScanTaskRequest) SetUuidList(v []*string) *CreateAgentlessScanTaskRequest {
	s.UuidList = v
	return s
}

func (s *CreateAgentlessScanTaskRequest) Validate() error {
	return dara.Validate(s)
}
