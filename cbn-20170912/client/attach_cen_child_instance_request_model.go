// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachCenChildInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *AttachCenChildInstanceRequest
	GetCenId() *string
	SetChildInstanceId(v string) *AttachCenChildInstanceRequest
	GetChildInstanceId() *string
	SetChildInstanceOwnerId(v int64) *AttachCenChildInstanceRequest
	GetChildInstanceOwnerId() *int64
	SetChildInstanceRegionId(v string) *AttachCenChildInstanceRequest
	GetChildInstanceRegionId() *string
	SetChildInstanceType(v string) *AttachCenChildInstanceRequest
	GetChildInstanceType() *string
	SetOwnerAccount(v string) *AttachCenChildInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AttachCenChildInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AttachCenChildInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AttachCenChildInstanceRequest
	GetResourceOwnerId() *int64
}

type AttachCenChildInstanceRequest struct {
	// The ID of the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-7qthudw0ll6jmc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the network instance that you want to attach to the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp18sth14qii3pnvx****
	ChildInstanceId *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the network instance belongs.
	//
	// > If the network instance and the CEN instance belong to different Alibaba Cloud accounts, this parameter is required.
	//
	// example:
	//
	// 1688000000000000
	ChildInstanceOwnerId *int64 `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
	// The ID of the region where the network instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// 	- **VPC**: VPC
	//
	// 	- **VBR**: VBR
	//
	// 	- **CCN**: CCN instance
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	ChildInstanceType    *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AttachCenChildInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachCenChildInstanceRequest) GoString() string {
	return s.String()
}

func (s *AttachCenChildInstanceRequest) GetCenId() *string {
	return s.CenId
}

func (s *AttachCenChildInstanceRequest) GetChildInstanceId() *string {
	return s.ChildInstanceId
}

func (s *AttachCenChildInstanceRequest) GetChildInstanceOwnerId() *int64 {
	return s.ChildInstanceOwnerId
}

func (s *AttachCenChildInstanceRequest) GetChildInstanceRegionId() *string {
	return s.ChildInstanceRegionId
}

func (s *AttachCenChildInstanceRequest) GetChildInstanceType() *string {
	return s.ChildInstanceType
}

func (s *AttachCenChildInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AttachCenChildInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AttachCenChildInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AttachCenChildInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AttachCenChildInstanceRequest) SetCenId(v string) *AttachCenChildInstanceRequest {
	s.CenId = &v
	return s
}

func (s *AttachCenChildInstanceRequest) SetChildInstanceId(v string) *AttachCenChildInstanceRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *AttachCenChildInstanceRequest) SetChildInstanceOwnerId(v int64) *AttachCenChildInstanceRequest {
	s.ChildInstanceOwnerId = &v
	return s
}

func (s *AttachCenChildInstanceRequest) SetChildInstanceRegionId(v string) *AttachCenChildInstanceRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *AttachCenChildInstanceRequest) SetChildInstanceType(v string) *AttachCenChildInstanceRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *AttachCenChildInstanceRequest) SetOwnerAccount(v string) *AttachCenChildInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AttachCenChildInstanceRequest) SetOwnerId(v int64) *AttachCenChildInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachCenChildInstanceRequest) SetResourceOwnerAccount(v string) *AttachCenChildInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachCenChildInstanceRequest) SetResourceOwnerId(v int64) *AttachCenChildInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AttachCenChildInstanceRequest) Validate() error {
	return dara.Validate(s)
}
