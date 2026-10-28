// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeAssetRefreshTaskConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetRefreshConfigs(v []*ChangeAssetRefreshTaskConfigRequestAssetRefreshConfigs) *ChangeAssetRefreshTaskConfigRequest
	GetAssetRefreshConfigs() []*ChangeAssetRefreshTaskConfigRequestAssetRefreshConfigs
	SetRegionId(v string) *ChangeAssetRefreshTaskConfigRequest
	GetRegionId() *string
}

type ChangeAssetRefreshTaskConfigRequest struct {
	// The list of asset refresh configurations.
	AssetRefreshConfigs []*ChangeAssetRefreshTaskConfigRequestAssetRefreshConfigs `json:"AssetRefreshConfigs,omitempty" xml:"AssetRefreshConfigs,omitempty" type:"Repeated"`
	// The region of the Security Center instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ChangeAssetRefreshTaskConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeAssetRefreshTaskConfigRequest) GoString() string {
	return s.String()
}

func (s *ChangeAssetRefreshTaskConfigRequest) GetAssetRefreshConfigs() []*ChangeAssetRefreshTaskConfigRequestAssetRefreshConfigs {
	return s.AssetRefreshConfigs
}

func (s *ChangeAssetRefreshTaskConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ChangeAssetRefreshTaskConfigRequest) SetAssetRefreshConfigs(v []*ChangeAssetRefreshTaskConfigRequestAssetRefreshConfigs) *ChangeAssetRefreshTaskConfigRequest {
	s.AssetRefreshConfigs = v
	return s
}

func (s *ChangeAssetRefreshTaskConfigRequest) SetRegionId(v string) *ChangeAssetRefreshTaskConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ChangeAssetRefreshTaskConfigRequest) Validate() error {
	if s.AssetRefreshConfigs != nil {
		for _, item := range s.AssetRefreshConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChangeAssetRefreshTaskConfigRequestAssetRefreshConfigs struct {
	// The configuration type. Valid values:
	//
	// - **0**: Host refresh task
	//
	// - **1**: Cloud service refresh task
	//
	// - **2**: AccessKey scheduled check task
	//
	// example:
	//
	// 2
	RefreshConfigType *int32 `json:"RefreshConfigType,omitempty" xml:"RefreshConfigType,omitempty"`
	// The synchronization cycle. Valid values:
	//
	//  - **60**: 60 minutes
	//
	//  - **180**: 3 hours
	//
	//  - **360**: 6 hours
	//
	//  - **720**: 12 hours
	//
	//  - **1440**: 1 day
	//
	//  - **10080**: 7 days
	//
	// example:
	//
	// 360
	SchedulePeriod *int32 `json:"SchedulePeriod,omitempty" xml:"SchedulePeriod,omitempty"`
	// The enabling status of the configuration. Valid values:
	//
	//  - **1**: Enabled
	//
	// - **0**: Disabled
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the specified AccessKey record when you configure an AccessKey scheduled check task.
	//
	// example:
	//
	// 2308
	TargetId *int64 `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The cloud asset vendor. Valid values:
	//
	// - **3**: Tencent Cloud
	//
	// - **4**: Huawei Cloud
	//
	// - **7**: Amazon Web Services
	//
	// example:
	//
	// 3
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s ChangeAssetRefreshTaskConfigRequestAssetRefreshConfigs) String() string {
	return dara.Prettify(s)
}

func (s ChangeAssetRefreshTaskConfigRequestAssetRefreshConfigs) GoString() string {
	return s.String()
}

func (s *ChangeAssetRefreshTaskConfigRequestAssetRefreshConfigs) GetRefreshConfigType() *int32 {
	return s.RefreshConfigType
}

func (s *ChangeAssetRefreshTaskConfigRequestAssetRefreshConfigs) GetSchedulePeriod() *int32 {
	return s.SchedulePeriod
}

func (s *ChangeAssetRefreshTaskConfigRequestAssetRefreshConfigs) GetStatus() *int32 {
	return s.Status
}

func (s *ChangeAssetRefreshTaskConfigRequestAssetRefreshConfigs) GetTargetId() *int64 {
	return s.TargetId
}

func (s *ChangeAssetRefreshTaskConfigRequestAssetRefreshConfigs) GetVendor() *int32 {
	return s.Vendor
}

func (s *ChangeAssetRefreshTaskConfigRequestAssetRefreshConfigs) SetRefreshConfigType(v int32) *ChangeAssetRefreshTaskConfigRequestAssetRefreshConfigs {
	s.RefreshConfigType = &v
	return s
}

func (s *ChangeAssetRefreshTaskConfigRequestAssetRefreshConfigs) SetSchedulePeriod(v int32) *ChangeAssetRefreshTaskConfigRequestAssetRefreshConfigs {
	s.SchedulePeriod = &v
	return s
}

func (s *ChangeAssetRefreshTaskConfigRequestAssetRefreshConfigs) SetStatus(v int32) *ChangeAssetRefreshTaskConfigRequestAssetRefreshConfigs {
	s.Status = &v
	return s
}

func (s *ChangeAssetRefreshTaskConfigRequestAssetRefreshConfigs) SetTargetId(v int64) *ChangeAssetRefreshTaskConfigRequestAssetRefreshConfigs {
	s.TargetId = &v
	return s
}

func (s *ChangeAssetRefreshTaskConfigRequestAssetRefreshConfigs) SetVendor(v int32) *ChangeAssetRefreshTaskConfigRequestAssetRefreshConfigs {
	s.Vendor = &v
	return s
}

func (s *ChangeAssetRefreshTaskConfigRequestAssetRefreshConfigs) Validate() error {
	return dara.Validate(s)
}
