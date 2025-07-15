// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVSwitchCidrReservationUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *GetVSwitchCidrReservationUsageRequest
	GetMaxResults() *int64
	SetNextToken(v string) *GetVSwitchCidrReservationUsageRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *GetVSwitchCidrReservationUsageRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetVSwitchCidrReservationUsageRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetVSwitchCidrReservationUsageRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetVSwitchCidrReservationUsageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetVSwitchCidrReservationUsageRequest
	GetResourceOwnerId() *int64
	SetVSwitchCidrReservationId(v string) *GetVSwitchCidrReservationUsageRequest
	GetVSwitchCidrReservationId() *string
}

type GetVSwitchCidrReservationUsageRequest struct {
	// The number of entries to return on each page. Valid values: **1*	- to **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- If a value is returned for NextToken, specify the value in the next request to retrieve a new page of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the vSwitch.
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
	// The ID of the reserved CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// vcr-bp1m12saqteraw3rp****
	VSwitchCidrReservationId *string `json:"VSwitchCidrReservationId,omitempty" xml:"VSwitchCidrReservationId,omitempty"`
}

func (s GetVSwitchCidrReservationUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVSwitchCidrReservationUsageRequest) GoString() string {
	return s.String()
}

func (s *GetVSwitchCidrReservationUsageRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *GetVSwitchCidrReservationUsageRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetVSwitchCidrReservationUsageRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetVSwitchCidrReservationUsageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetVSwitchCidrReservationUsageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVSwitchCidrReservationUsageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetVSwitchCidrReservationUsageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetVSwitchCidrReservationUsageRequest) GetVSwitchCidrReservationId() *string {
	return s.VSwitchCidrReservationId
}

func (s *GetVSwitchCidrReservationUsageRequest) SetMaxResults(v int64) *GetVSwitchCidrReservationUsageRequest {
	s.MaxResults = &v
	return s
}

func (s *GetVSwitchCidrReservationUsageRequest) SetNextToken(v string) *GetVSwitchCidrReservationUsageRequest {
	s.NextToken = &v
	return s
}

func (s *GetVSwitchCidrReservationUsageRequest) SetOwnerAccount(v string) *GetVSwitchCidrReservationUsageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetVSwitchCidrReservationUsageRequest) SetOwnerId(v int64) *GetVSwitchCidrReservationUsageRequest {
	s.OwnerId = &v
	return s
}

func (s *GetVSwitchCidrReservationUsageRequest) SetRegionId(v string) *GetVSwitchCidrReservationUsageRequest {
	s.RegionId = &v
	return s
}

func (s *GetVSwitchCidrReservationUsageRequest) SetResourceOwnerAccount(v string) *GetVSwitchCidrReservationUsageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetVSwitchCidrReservationUsageRequest) SetResourceOwnerId(v int64) *GetVSwitchCidrReservationUsageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetVSwitchCidrReservationUsageRequest) SetVSwitchCidrReservationId(v string) *GetVSwitchCidrReservationUsageRequest {
	s.VSwitchCidrReservationId = &v
	return s
}

func (s *GetVSwitchCidrReservationUsageRequest) Validate() error {
	return dara.Validate(s)
}
