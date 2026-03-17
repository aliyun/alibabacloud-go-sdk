// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagVbrRelationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSagVbrRelationsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSagVbrRelationsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSagVbrRelationsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSagVbrRelationsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSagVbrRelationsRequest
	GetResourceOwnerId() *int64
	SetVbrInstanceIds(v []*string) *DescribeSagVbrRelationsRequest
	GetVbrInstanceIds() []*string
	SetVbrRegionId(v string) *DescribeSagVbrRelationsRequest
	GetVbrRegionId() *string
}

type DescribeSagVbrRelationsRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/69813.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IDs of the VBRs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-bp15ihkk93ezxppkd****
	VbrInstanceIds []*string `json:"VbrInstanceIds,omitempty" xml:"VbrInstanceIds,omitempty" type:"Repeated"`
	// The ID of the region where the VBR is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	VbrRegionId *string `json:"VbrRegionId,omitempty" xml:"VbrRegionId,omitempty"`
}

func (s DescribeSagVbrRelationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagVbrRelationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSagVbrRelationsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSagVbrRelationsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSagVbrRelationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSagVbrRelationsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSagVbrRelationsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSagVbrRelationsRequest) GetVbrInstanceIds() []*string {
	return s.VbrInstanceIds
}

func (s *DescribeSagVbrRelationsRequest) GetVbrRegionId() *string {
	return s.VbrRegionId
}

func (s *DescribeSagVbrRelationsRequest) SetOwnerAccount(v string) *DescribeSagVbrRelationsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSagVbrRelationsRequest) SetOwnerId(v int64) *DescribeSagVbrRelationsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSagVbrRelationsRequest) SetRegionId(v string) *DescribeSagVbrRelationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSagVbrRelationsRequest) SetResourceOwnerAccount(v string) *DescribeSagVbrRelationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSagVbrRelationsRequest) SetResourceOwnerId(v int64) *DescribeSagVbrRelationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSagVbrRelationsRequest) SetVbrInstanceIds(v []*string) *DescribeSagVbrRelationsRequest {
	s.VbrInstanceIds = v
	return s
}

func (s *DescribeSagVbrRelationsRequest) SetVbrRegionId(v string) *DescribeSagVbrRelationsRequest {
	s.VbrRegionId = &v
	return s
}

func (s *DescribeSagVbrRelationsRequest) Validate() error {
	return dara.Validate(s)
}
