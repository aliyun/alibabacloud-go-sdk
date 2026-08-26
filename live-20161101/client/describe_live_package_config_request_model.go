// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLivePackageConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLivePackageConfigRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLivePackageConfigRequest
	GetDomainName() *string
	SetOrder(v string) *DescribeLivePackageConfigRequest
	GetOrder() *string
	SetOwnerId(v int64) *DescribeLivePackageConfigRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *DescribeLivePackageConfigRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeLivePackageConfigRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeLivePackageConfigRequest
	GetRegionId() *string
	SetStreamName(v string) *DescribeLivePackageConfigRequest
	GetStreamName() *string
}

type DescribeLivePackageConfigRequest struct {
	// The app name. Leave empty to match all names.
	//
	// example:
	//
	// AppName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The live streaming domain (primary playback domain).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The sort order. Valid values:
	//
	// - **asc*	- (default): ascending order.
	//
	// - **desc**: descending order.
	//
	// example:
	//
	// asc
	Order   *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The current page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The page size. Default value: 10. Valid values: 5 to 30.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The live stream name. Leave empty to match all stream names.
	//
	// example:
	//
	// StreamName
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLivePackageConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePackageConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLivePackageConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLivePackageConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLivePackageConfigRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeLivePackageConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLivePackageConfigRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeLivePackageConfigRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLivePackageConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLivePackageConfigRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLivePackageConfigRequest) SetAppName(v string) *DescribeLivePackageConfigRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLivePackageConfigRequest) SetDomainName(v string) *DescribeLivePackageConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLivePackageConfigRequest) SetOrder(v string) *DescribeLivePackageConfigRequest {
	s.Order = &v
	return s
}

func (s *DescribeLivePackageConfigRequest) SetOwnerId(v int64) *DescribeLivePackageConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLivePackageConfigRequest) SetPageNum(v int32) *DescribeLivePackageConfigRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeLivePackageConfigRequest) SetPageSize(v int32) *DescribeLivePackageConfigRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLivePackageConfigRequest) SetRegionId(v string) *DescribeLivePackageConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLivePackageConfigRequest) SetStreamName(v string) *DescribeLivePackageConfigRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLivePackageConfigRequest) Validate() error {
	return dara.Validate(s)
}
