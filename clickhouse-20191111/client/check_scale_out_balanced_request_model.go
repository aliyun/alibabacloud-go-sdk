// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckScaleOutBalancedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CheckScaleOutBalancedRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *CheckScaleOutBalancedRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CheckScaleOutBalancedRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *CheckScaleOutBalancedRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *CheckScaleOutBalancedRequest
	GetPageSize() *int32
	SetRegionId(v string) *CheckScaleOutBalancedRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CheckScaleOutBalancedRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckScaleOutBalancedRequest
	GetResourceOwnerId() *int64
}

type CheckScaleOutBalancedRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of pages to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page. Default value: 30. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
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

func (s CheckScaleOutBalancedRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckScaleOutBalancedRequest) GoString() string {
	return s.String()
}

func (s *CheckScaleOutBalancedRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CheckScaleOutBalancedRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckScaleOutBalancedRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckScaleOutBalancedRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *CheckScaleOutBalancedRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *CheckScaleOutBalancedRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckScaleOutBalancedRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckScaleOutBalancedRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckScaleOutBalancedRequest) SetDBClusterId(v string) *CheckScaleOutBalancedRequest {
	s.DBClusterId = &v
	return s
}

func (s *CheckScaleOutBalancedRequest) SetOwnerAccount(v string) *CheckScaleOutBalancedRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckScaleOutBalancedRequest) SetOwnerId(v int64) *CheckScaleOutBalancedRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckScaleOutBalancedRequest) SetPageNumber(v int32) *CheckScaleOutBalancedRequest {
	s.PageNumber = &v
	return s
}

func (s *CheckScaleOutBalancedRequest) SetPageSize(v int32) *CheckScaleOutBalancedRequest {
	s.PageSize = &v
	return s
}

func (s *CheckScaleOutBalancedRequest) SetRegionId(v string) *CheckScaleOutBalancedRequest {
	s.RegionId = &v
	return s
}

func (s *CheckScaleOutBalancedRequest) SetResourceOwnerAccount(v string) *CheckScaleOutBalancedRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckScaleOutBalancedRequest) SetResourceOwnerId(v int64) *CheckScaleOutBalancedRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckScaleOutBalancedRequest) Validate() error {
	return dara.Validate(s)
}
