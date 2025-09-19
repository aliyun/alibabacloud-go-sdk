// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetRefreshTaskConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssetRefreshConfig(v []*ListAssetRefreshTaskConfigResponseBodyAssetRefreshConfig) *ListAssetRefreshTaskConfigResponseBody
	GetAssetRefreshConfig() []*ListAssetRefreshTaskConfigResponseBodyAssetRefreshConfig
	SetRequestId(v string) *ListAssetRefreshTaskConfigResponseBody
	GetRequestId() *string
}

type ListAssetRefreshTaskConfigResponseBody struct {
	// The asset synchronization configuration.
	AssetRefreshConfig []*ListAssetRefreshTaskConfigResponseBodyAssetRefreshConfig `json:"AssetRefreshConfig,omitempty" xml:"AssetRefreshConfig,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// D65AADFC-1D20-5A6A-8F6A-9FA53C0DC1F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAssetRefreshTaskConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAssetRefreshTaskConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListAssetRefreshTaskConfigResponseBody) GetAssetRefreshConfig() []*ListAssetRefreshTaskConfigResponseBodyAssetRefreshConfig {
	return s.AssetRefreshConfig
}

func (s *ListAssetRefreshTaskConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAssetRefreshTaskConfigResponseBody) SetAssetRefreshConfig(v []*ListAssetRefreshTaskConfigResponseBodyAssetRefreshConfig) *ListAssetRefreshTaskConfigResponseBody {
	s.AssetRefreshConfig = v
	return s
}

func (s *ListAssetRefreshTaskConfigResponseBody) SetRequestId(v string) *ListAssetRefreshTaskConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAssetRefreshTaskConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAssetRefreshTaskConfigResponseBodyAssetRefreshConfig struct {
	// The type of the configuration. Valid values:
	//
	// 	- **0**: server synchronization task
	//
	// 	- **1**: cloud service synchronization task
	//
	// 	- **2**: scheduled AccessKey pair verification task
	//
	// example:
	//
	// 1
	RefreshConfigType *int32 `json:"RefreshConfigType,omitempty" xml:"RefreshConfigType,omitempty"`
	// The synchronization cycle. Valid values:
	//
	// 	- **60**: 60 minutes
	//
	// 	- **180**: 3 hours
	//
	// 	- **360**: 6 hours
	//
	// 	- **720**: 12 hours
	//
	// 	- **1440**: 1 day
	//
	// 	- **10080**: 7 days
	//
	// example:
	//
	// 360
	SchedulePeriod *int32 `json:"SchedulePeriod,omitempty" xml:"SchedulePeriod,omitempty"`
	// The status of the configuration. Valid values:
	//
	// 	- **1**: enabled
	//
	// 	- **0**: disabled
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The service provider of the cloud asset. Valid values:
	//
	// 	- **3**: Tencent Cloud
	//
	// 	- **4**: Huawei Cloud
	//
	// 	- **7**: Amazon Web Services (AWS) Cloud
	//
	// example:
	//
	// 3
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s ListAssetRefreshTaskConfigResponseBodyAssetRefreshConfig) String() string {
	return dara.Prettify(s)
}

func (s ListAssetRefreshTaskConfigResponseBodyAssetRefreshConfig) GoString() string {
	return s.String()
}

func (s *ListAssetRefreshTaskConfigResponseBodyAssetRefreshConfig) GetRefreshConfigType() *int32 {
	return s.RefreshConfigType
}

func (s *ListAssetRefreshTaskConfigResponseBodyAssetRefreshConfig) GetSchedulePeriod() *int32 {
	return s.SchedulePeriod
}

func (s *ListAssetRefreshTaskConfigResponseBodyAssetRefreshConfig) GetStatus() *int32 {
	return s.Status
}

func (s *ListAssetRefreshTaskConfigResponseBodyAssetRefreshConfig) GetVendor() *int32 {
	return s.Vendor
}

func (s *ListAssetRefreshTaskConfigResponseBodyAssetRefreshConfig) SetRefreshConfigType(v int32) *ListAssetRefreshTaskConfigResponseBodyAssetRefreshConfig {
	s.RefreshConfigType = &v
	return s
}

func (s *ListAssetRefreshTaskConfigResponseBodyAssetRefreshConfig) SetSchedulePeriod(v int32) *ListAssetRefreshTaskConfigResponseBodyAssetRefreshConfig {
	s.SchedulePeriod = &v
	return s
}

func (s *ListAssetRefreshTaskConfigResponseBodyAssetRefreshConfig) SetStatus(v int32) *ListAssetRefreshTaskConfigResponseBodyAssetRefreshConfig {
	s.Status = &v
	return s
}

func (s *ListAssetRefreshTaskConfigResponseBodyAssetRefreshConfig) SetVendor(v int32) *ListAssetRefreshTaskConfigResponseBodyAssetRefreshConfig {
	s.Vendor = &v
	return s
}

func (s *ListAssetRefreshTaskConfigResponseBodyAssetRefreshConfig) Validate() error {
	return dara.Validate(s)
}
