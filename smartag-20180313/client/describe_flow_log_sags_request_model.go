// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowLogSagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowLogId(v string) *DescribeFlowLogSagsRequest
	GetFlowLogId() *string
	SetOwnerAccount(v string) *DescribeFlowLogSagsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeFlowLogSagsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeFlowLogSagsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeFlowLogSagsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeFlowLogSagsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeFlowLogSagsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeFlowLogSagsRequest
	GetResourceOwnerId() *int64
}

type DescribeFlowLogSagsRequest struct {
	// The instance ID of the flow log.
	//
	// example:
	//
	// fl-l934tsa5504yuc****
	FlowLogId    *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the SAG instance list. Minimum value: **1**. Default value: **1**
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in the case of a paged query result. Maximum value: **50**. Default value: **10**
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region to which the flow log belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghahi
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeFlowLogSagsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowLogSagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogSagsRequest) GetFlowLogId() *string {
	return s.FlowLogId
}

func (s *DescribeFlowLogSagsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeFlowLogSagsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeFlowLogSagsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeFlowLogSagsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFlowLogSagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFlowLogSagsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeFlowLogSagsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeFlowLogSagsRequest) SetFlowLogId(v string) *DescribeFlowLogSagsRequest {
	s.FlowLogId = &v
	return s
}

func (s *DescribeFlowLogSagsRequest) SetOwnerAccount(v string) *DescribeFlowLogSagsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeFlowLogSagsRequest) SetOwnerId(v int64) *DescribeFlowLogSagsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeFlowLogSagsRequest) SetPageNumber(v int32) *DescribeFlowLogSagsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeFlowLogSagsRequest) SetPageSize(v int32) *DescribeFlowLogSagsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFlowLogSagsRequest) SetRegionId(v string) *DescribeFlowLogSagsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowLogSagsRequest) SetResourceOwnerAccount(v string) *DescribeFlowLogSagsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeFlowLogSagsRequest) SetResourceOwnerId(v int64) *DescribeFlowLogSagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeFlowLogSagsRequest) Validate() error {
	return dara.Validate(s)
}
