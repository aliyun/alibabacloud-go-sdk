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
	// Identification of asset selection.
	//
	// example:
	//
	// AGENTLESS_SCAN_ONCE_TASK_1720145******
	AssetSelectionType *string `json:"AssetSelectionType,omitempty" xml:"AssetSelectionType,omitempty"`
	// The retention period of images. Unit: days.
	//
	// example:
	//
	// 1
	AutoDeleteDays *int32 `json:"AutoDeleteDays,omitempty" xml:"AutoDeleteDays,omitempty"`
	// Specifies whether to enable the cost-saving mode. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ReleaseAfterScan *bool `json:"ReleaseAfterScan,omitempty" xml:"ReleaseAfterScan,omitempty"`
	// Specifies whether to check data disks. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ScanDataDisk *bool `json:"ScanDataDisk,omitempty" xml:"ScanDataDisk,omitempty"`
	// The type of the detection object. Valid values:
	//
	// 	- **2**: image
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	TargetType *int32 `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The UUIDs of the assets on which you want to run the detection task.
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
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
