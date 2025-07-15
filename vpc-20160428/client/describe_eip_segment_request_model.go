// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEipSegmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeEipSegmentRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DescribeEipSegmentRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeEipSegmentRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeEipSegmentRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEipSegmentRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeEipSegmentRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeEipSegmentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeEipSegmentRequest
	GetResourceOwnerId() *int64
	SetSegmentInstanceId(v string) *DescribeEipSegmentRequest
	GetSegmentInstanceId() *string
}

type DescribeEipSegmentRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001sdfg
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// The ID of the region to which the contiguous EIP group belongs. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the contiguous EIP group that you want to query.
	//
	// example:
	//
	// eipsg-2zett8ba055tbsxme****
	SegmentInstanceId *string `json:"SegmentInstanceId,omitempty" xml:"SegmentInstanceId,omitempty"`
}

func (s DescribeEipSegmentRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipSegmentRequest) GoString() string {
	return s.String()
}

func (s *DescribeEipSegmentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeEipSegmentRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeEipSegmentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeEipSegmentRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEipSegmentRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEipSegmentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEipSegmentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeEipSegmentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeEipSegmentRequest) GetSegmentInstanceId() *string {
	return s.SegmentInstanceId
}

func (s *DescribeEipSegmentRequest) SetClientToken(v string) *DescribeEipSegmentRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeEipSegmentRequest) SetOwnerAccount(v string) *DescribeEipSegmentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeEipSegmentRequest) SetOwnerId(v int64) *DescribeEipSegmentRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEipSegmentRequest) SetPageNumber(v int32) *DescribeEipSegmentRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEipSegmentRequest) SetPageSize(v int32) *DescribeEipSegmentRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEipSegmentRequest) SetRegionId(v string) *DescribeEipSegmentRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEipSegmentRequest) SetResourceOwnerAccount(v string) *DescribeEipSegmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeEipSegmentRequest) SetResourceOwnerId(v int64) *DescribeEipSegmentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeEipSegmentRequest) SetSegmentInstanceId(v string) *DescribeEipSegmentRequest {
	s.SegmentInstanceId = &v
	return s
}

func (s *DescribeEipSegmentRequest) Validate() error {
	return dara.Validate(s)
}
