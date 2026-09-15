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
	SetFrom(v string) *CreateAgentlessScanTaskRequest
	GetFrom() *string
	SetRegionId(v string) *CreateAgentlessScanTaskRequest
	GetRegionId() *string
	SetReleaseAfterScan(v bool) *CreateAgentlessScanTaskRequest
	GetReleaseAfterScan() *bool
	SetResourceRegionId(v string) *CreateAgentlessScanTaskRequest
	GetResourceRegionId() *string
	SetScanDataDisk(v bool) *CreateAgentlessScanTaskRequest
	GetScanDataDisk() *bool
	SetTargetType(v int32) *CreateAgentlessScanTaskRequest
	GetTargetType() *int32
	SetTargets(v []*CreateAgentlessScanTaskRequestTargets) *CreateAgentlessScanTaskRequest
	GetTargets() []*CreateAgentlessScanTaskRequestTargets
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
	// The image retention period, in days. This parameter takes effect only for host detection. It does not take effect for user snapshot detection or user custom image detection.
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
	// The source of the API call, which is used to collect statistics on scan task volume and scan data volume by source. If this parameter is not specified, the value is empty.
	//
	// example:
	//
	// image-console
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The region ID, which is usually automatically populated by the gateway.
	//
	// example:
	//
	// cn-shanghai
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
	// The region ID of the resource to be detected, such as cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
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
	// The list of targets for image security remediation. Each target specifies the source image, the region, the name of the remediated image, and the vulnerability identifiers to be fixed.
	Targets []*CreateAgentlessScanTaskRequestTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
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

func (s *CreateAgentlessScanTaskRequest) GetFrom() *string {
	return s.From
}

func (s *CreateAgentlessScanTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAgentlessScanTaskRequest) GetReleaseAfterScan() *bool {
	return s.ReleaseAfterScan
}

func (s *CreateAgentlessScanTaskRequest) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *CreateAgentlessScanTaskRequest) GetScanDataDisk() *bool {
	return s.ScanDataDisk
}

func (s *CreateAgentlessScanTaskRequest) GetTargetType() *int32 {
	return s.TargetType
}

func (s *CreateAgentlessScanTaskRequest) GetTargets() []*CreateAgentlessScanTaskRequestTargets {
	return s.Targets
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

func (s *CreateAgentlessScanTaskRequest) SetFrom(v string) *CreateAgentlessScanTaskRequest {
	s.From = &v
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

func (s *CreateAgentlessScanTaskRequest) SetResourceRegionId(v string) *CreateAgentlessScanTaskRequest {
	s.ResourceRegionId = &v
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

func (s *CreateAgentlessScanTaskRequest) SetTargets(v []*CreateAgentlessScanTaskRequestTargets) *CreateAgentlessScanTaskRequest {
	s.Targets = v
	return s
}

func (s *CreateAgentlessScanTaskRequest) SetUuidList(v []*string) *CreateAgentlessScanTaskRequest {
	s.UuidList = v
	return s
}

func (s *CreateAgentlessScanTaskRequest) Validate() error {
	if s.Targets != nil {
		for _, item := range s.Targets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAgentlessScanTaskRequestTargets struct {
	// The ID of the source ECS custom image to be remediated. The image must be located in the region specified by RegionId of this target.
	//
	// example:
	//
	// m-bp1example123456789
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the source ECS custom image to be remediated.
	//
	// example:
	//
	// source-image
	OriginImageName *string `json:"OriginImageName,omitempty" xml:"OriginImageName,omitempty"`
	// The name of the ECS image generated after remediation.
	//
	// example:
	//
	// patched-image-20260909
	OutputImageName *string `json:"OutputImageName,omitempty" xml:"OutputImageName,omitempty"`
	// The region ID of the source image to be remediated, such as cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of vulnerability identifiers to be fixed. At least one vulnerability identifier must be specified, and each identifier must be unique and non-empty.
	VulnerabilityIds []*string `json:"VulnerabilityIds,omitempty" xml:"VulnerabilityIds,omitempty" type:"Repeated"`
}

func (s CreateAgentlessScanTaskRequestTargets) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentlessScanTaskRequestTargets) GoString() string {
	return s.String()
}

func (s *CreateAgentlessScanTaskRequestTargets) GetImageId() *string {
	return s.ImageId
}

func (s *CreateAgentlessScanTaskRequestTargets) GetOriginImageName() *string {
	return s.OriginImageName
}

func (s *CreateAgentlessScanTaskRequestTargets) GetOutputImageName() *string {
	return s.OutputImageName
}

func (s *CreateAgentlessScanTaskRequestTargets) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAgentlessScanTaskRequestTargets) GetVulnerabilityIds() []*string {
	return s.VulnerabilityIds
}

func (s *CreateAgentlessScanTaskRequestTargets) SetImageId(v string) *CreateAgentlessScanTaskRequestTargets {
	s.ImageId = &v
	return s
}

func (s *CreateAgentlessScanTaskRequestTargets) SetOriginImageName(v string) *CreateAgentlessScanTaskRequestTargets {
	s.OriginImageName = &v
	return s
}

func (s *CreateAgentlessScanTaskRequestTargets) SetOutputImageName(v string) *CreateAgentlessScanTaskRequestTargets {
	s.OutputImageName = &v
	return s
}

func (s *CreateAgentlessScanTaskRequestTargets) SetRegionId(v string) *CreateAgentlessScanTaskRequestTargets {
	s.RegionId = &v
	return s
}

func (s *CreateAgentlessScanTaskRequestTargets) SetVulnerabilityIds(v []*string) *CreateAgentlessScanTaskRequestTargets {
	s.VulnerabilityIds = v
	return s
}

func (s *CreateAgentlessScanTaskRequestTargets) Validate() error {
	return dara.Validate(s)
}
