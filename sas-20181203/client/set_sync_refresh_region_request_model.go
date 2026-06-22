// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSyncRefreshRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultRegion(v int32) *SetSyncRefreshRegionRequest
	GetDefaultRegion() *int32
	SetRegionIds(v []*string) *SetSyncRefreshRegionRequest
	GetRegionIds() []*string
	SetVendor(v string) *SetSyncRefreshRegionRequest
	GetVendor() *string
}

type SetSyncRefreshRegionRequest struct {
	// The access type of the multi-cloud site. Valid values:
	//
	// - **0**: no default site is configured, and you can select regions as needed.
	//
	// - **1**: the current site is already the default site.
	//
	// example:
	//
	// 0
	DefaultRegion *int32 `json:"DefaultRegion,omitempty" xml:"DefaultRegion,omitempty"`
	// The list of regions to be synchronized for the current site.
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	// The cloud asset vendor. Valid values:
	//
	// - **Tencent**: Tencent Cloud.
	//
	// - **HUAWEICLOUD**: Huawei Cloud.
	//
	// - **Azure**: Azure.
	//
	// - **AWS**: AWS.
	//
	// example:
	//
	// Tencent
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s SetSyncRefreshRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s SetSyncRefreshRegionRequest) GoString() string {
	return s.String()
}

func (s *SetSyncRefreshRegionRequest) GetDefaultRegion() *int32 {
	return s.DefaultRegion
}

func (s *SetSyncRefreshRegionRequest) GetRegionIds() []*string {
	return s.RegionIds
}

func (s *SetSyncRefreshRegionRequest) GetVendor() *string {
	return s.Vendor
}

func (s *SetSyncRefreshRegionRequest) SetDefaultRegion(v int32) *SetSyncRefreshRegionRequest {
	s.DefaultRegion = &v
	return s
}

func (s *SetSyncRefreshRegionRequest) SetRegionIds(v []*string) *SetSyncRefreshRegionRequest {
	s.RegionIds = v
	return s
}

func (s *SetSyncRefreshRegionRequest) SetVendor(v string) *SetSyncRefreshRegionRequest {
	s.Vendor = &v
	return s
}

func (s *SetSyncRefreshRegionRequest) Validate() error {
	return dara.Validate(s)
}
