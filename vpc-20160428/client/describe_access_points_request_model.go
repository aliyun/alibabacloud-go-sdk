// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessPointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeAccessPointsRequest
	GetAcceptLanguage() *string
	SetOwnerId(v int64) *DescribeAccessPointsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeAccessPointsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAccessPointsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeAccessPointsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAccessPointsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAccessPointsRequest
	GetResourceOwnerId() *int64
}

type DescribeAccessPointsRequest struct {
	// The language of the response. Valid values:
	//
	// 	- **zh-CN**: Chinese
	//
	// 	- **en-US*	- (default): English
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the access point.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAccessPointsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessPointsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeAccessPointsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAccessPointsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAccessPointsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAccessPointsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAccessPointsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAccessPointsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAccessPointsRequest) SetAcceptLanguage(v string) *DescribeAccessPointsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeAccessPointsRequest) SetOwnerId(v int64) *DescribeAccessPointsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAccessPointsRequest) SetPageNumber(v int32) *DescribeAccessPointsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessPointsRequest) SetPageSize(v int32) *DescribeAccessPointsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessPointsRequest) SetRegionId(v string) *DescribeAccessPointsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAccessPointsRequest) SetResourceOwnerAccount(v string) *DescribeAccessPointsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAccessPointsRequest) SetResourceOwnerId(v int64) *DescribeAccessPointsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAccessPointsRequest) Validate() error {
	return dara.Validate(s)
}
