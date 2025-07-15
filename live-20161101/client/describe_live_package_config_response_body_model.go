// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLivePackageConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLivePackageConfigList(v *DescribeLivePackageConfigResponseBodyLivePackageConfigList) *DescribeLivePackageConfigResponseBody
	GetLivePackageConfigList() *DescribeLivePackageConfigResponseBodyLivePackageConfigList
	SetOrder(v string) *DescribeLivePackageConfigResponseBody
	GetOrder() *string
	SetPageNum(v int32) *DescribeLivePackageConfigResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeLivePackageConfigResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLivePackageConfigResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *DescribeLivePackageConfigResponseBody
	GetTotalNum() *int32
	SetTotalPage(v int32) *DescribeLivePackageConfigResponseBody
	GetTotalPage() *int32
}

type DescribeLivePackageConfigResponseBody struct {
	// The live stream encapsulation configurations.
	LivePackageConfigList *DescribeLivePackageConfigResponseBodyLivePackageConfigList `json:"LivePackageConfigList,omitempty" xml:"LivePackageConfigList,omitempty" type:"Struct"`
	// The sorting order. Valid values:
	//
	// 	- **asc*	- (default): ascending order
	//
	// 	- **desc**: descending order
	//
	// example:
	//
	// asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1FEDCFD8-4C5D-5CB3-A5A1-0B59E5AE57B0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of live stream encapsulation configurations.
	//
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 10
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeLivePackageConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePackageConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLivePackageConfigResponseBody) GetLivePackageConfigList() *DescribeLivePackageConfigResponseBodyLivePackageConfigList {
	return s.LivePackageConfigList
}

func (s *DescribeLivePackageConfigResponseBody) GetOrder() *string {
	return s.Order
}

func (s *DescribeLivePackageConfigResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeLivePackageConfigResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLivePackageConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLivePackageConfigResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *DescribeLivePackageConfigResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeLivePackageConfigResponseBody) SetLivePackageConfigList(v *DescribeLivePackageConfigResponseBodyLivePackageConfigList) *DescribeLivePackageConfigResponseBody {
	s.LivePackageConfigList = v
	return s
}

func (s *DescribeLivePackageConfigResponseBody) SetOrder(v string) *DescribeLivePackageConfigResponseBody {
	s.Order = &v
	return s
}

func (s *DescribeLivePackageConfigResponseBody) SetPageNum(v int32) *DescribeLivePackageConfigResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeLivePackageConfigResponseBody) SetPageSize(v int32) *DescribeLivePackageConfigResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLivePackageConfigResponseBody) SetRequestId(v string) *DescribeLivePackageConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLivePackageConfigResponseBody) SetTotalNum(v int32) *DescribeLivePackageConfigResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeLivePackageConfigResponseBody) SetTotalPage(v int32) *DescribeLivePackageConfigResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeLivePackageConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLivePackageConfigResponseBodyLivePackageConfigList struct {
	LivePackageConfig []*DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig `json:"LivePackageConfig,omitempty" xml:"LivePackageConfig,omitempty" type:"Repeated"`
}

func (s DescribeLivePackageConfigResponseBodyLivePackageConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePackageConfigResponseBodyLivePackageConfigList) GoString() string {
	return s.String()
}

func (s *DescribeLivePackageConfigResponseBodyLivePackageConfigList) GetLivePackageConfig() []*DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig {
	return s.LivePackageConfig
}

func (s *DescribeLivePackageConfigResponseBodyLivePackageConfigList) SetLivePackageConfig(v []*DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig) *DescribeLivePackageConfigResponseBodyLivePackageConfigList {
	s.LivePackageConfig = v
	return s
}

func (s *DescribeLivePackageConfigResponseBodyLivePackageConfigList) Validate() error {
	return dara.Validate(s)
}

type DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig struct {
	// The application name.
	//
	// example:
	//
	// AppName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The main streaming domain.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Indicates whether the transcoded stream is ignored. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IgnoreTranscode *bool `json:"IgnoreTranscode,omitempty" xml:"IgnoreTranscode,omitempty"`
	// The part length. Unit: milliseconds.
	//
	// example:
	//
	// 0
	PartDuration *int32 `json:"PartDuration,omitempty" xml:"PartDuration,omitempty"`
	// The streaming protocol and encapsulation format.
	//
	// example:
	//
	// HLS_CMAF
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The segment length. Unit: seconds.
	//
	// example:
	//
	// 5
	SegmentDuration *int32 `json:"SegmentDuration,omitempty" xml:"SegmentDuration,omitempty"`
	// The number of segments.
	//
	// example:
	//
	// 3
	SegmentNum *int32 `json:"SegmentNum,omitempty" xml:"SegmentNum,omitempty"`
	// The stream name.
	//
	// example:
	//
	// StreamName
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig) GoString() string {
	return s.String()
}

func (s *DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig) GetIgnoreTranscode() *bool {
	return s.IgnoreTranscode
}

func (s *DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig) GetPartDuration() *int32 {
	return s.PartDuration
}

func (s *DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig) GetSegmentDuration() *int32 {
	return s.SegmentDuration
}

func (s *DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig) GetSegmentNum() *int32 {
	return s.SegmentNum
}

func (s *DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig) SetAppName(v string) *DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig {
	s.AppName = &v
	return s
}

func (s *DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig) SetDomainName(v string) *DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig {
	s.DomainName = &v
	return s
}

func (s *DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig) SetIgnoreTranscode(v bool) *DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig {
	s.IgnoreTranscode = &v
	return s
}

func (s *DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig) SetPartDuration(v int32) *DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig {
	s.PartDuration = &v
	return s
}

func (s *DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig) SetProtocol(v string) *DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig {
	s.Protocol = &v
	return s
}

func (s *DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig) SetSegmentDuration(v int32) *DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig {
	s.SegmentDuration = &v
	return s
}

func (s *DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig) SetSegmentNum(v int32) *DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig {
	s.SegmentNum = &v
	return s
}

func (s *DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig) SetStreamName(v string) *DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig {
	s.StreamName = &v
	return s
}

func (s *DescribeLivePackageConfigResponseBodyLivePackageConfigListLivePackageConfig) Validate() error {
	return dara.Validate(s)
}
