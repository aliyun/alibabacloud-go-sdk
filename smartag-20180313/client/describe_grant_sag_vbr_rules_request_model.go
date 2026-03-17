// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGrantSagVbrRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeGrantSagVbrRulesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeGrantSagVbrRulesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeGrantSagVbrRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGrantSagVbrRulesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeGrantSagVbrRulesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeGrantSagVbrRulesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeGrantSagVbrRulesRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeGrantSagVbrRulesRequest
	GetSmartAGId() *string
	SetVbrInstanceId(v string) *DescribeGrantSagVbrRulesRequest
	GetVbrInstanceId() *string
}

type DescribeGrantSagVbrRulesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value:**10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// example:
	//
	// sag-0nnteglltw6z4b****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The ID of the VBR.
	//
	// example:
	//
	// vbr-bp13gtbhdp0pfqg6s****
	VbrInstanceId *string `json:"VbrInstanceId,omitempty" xml:"VbrInstanceId,omitempty"`
}

func (s DescribeGrantSagVbrRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantSagVbrRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGrantSagVbrRulesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeGrantSagVbrRulesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeGrantSagVbrRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGrantSagVbrRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGrantSagVbrRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGrantSagVbrRulesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeGrantSagVbrRulesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeGrantSagVbrRulesRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeGrantSagVbrRulesRequest) GetVbrInstanceId() *string {
	return s.VbrInstanceId
}

func (s *DescribeGrantSagVbrRulesRequest) SetOwnerAccount(v string) *DescribeGrantSagVbrRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeGrantSagVbrRulesRequest) SetOwnerId(v int64) *DescribeGrantSagVbrRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGrantSagVbrRulesRequest) SetPageNumber(v int32) *DescribeGrantSagVbrRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGrantSagVbrRulesRequest) SetPageSize(v int32) *DescribeGrantSagVbrRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGrantSagVbrRulesRequest) SetRegionId(v string) *DescribeGrantSagVbrRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGrantSagVbrRulesRequest) SetResourceOwnerAccount(v string) *DescribeGrantSagVbrRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeGrantSagVbrRulesRequest) SetResourceOwnerId(v int64) *DescribeGrantSagVbrRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeGrantSagVbrRulesRequest) SetSmartAGId(v string) *DescribeGrantSagVbrRulesRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeGrantSagVbrRulesRequest) SetVbrInstanceId(v string) *DescribeGrantSagVbrRulesRequest {
	s.VbrInstanceId = &v
	return s
}

func (s *DescribeGrantSagVbrRulesRequest) Validate() error {
	return dara.Validate(s)
}
