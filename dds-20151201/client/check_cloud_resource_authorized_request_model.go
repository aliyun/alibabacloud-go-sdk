// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCloudResourceAuthorizedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CheckCloudResourceAuthorizedRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *CheckCloudResourceAuthorizedRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CheckCloudResourceAuthorizedRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CheckCloudResourceAuthorizedRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckCloudResourceAuthorizedRequest
	GetResourceOwnerId() *int64
	SetTargetRegionId(v string) *CheckCloudResourceAuthorizedRequest
	GetTargetRegionId() *string
}

type CheckCloudResourceAuthorizedRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp18f7d6b6a7****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/62010.html) operation to query the region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	TargetRegionId *string `json:"TargetRegionId,omitempty" xml:"TargetRegionId,omitempty"`
}

func (s CheckCloudResourceAuthorizedRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckCloudResourceAuthorizedRequest) GoString() string {
	return s.String()
}

func (s *CheckCloudResourceAuthorizedRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CheckCloudResourceAuthorizedRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckCloudResourceAuthorizedRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckCloudResourceAuthorizedRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckCloudResourceAuthorizedRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckCloudResourceAuthorizedRequest) GetTargetRegionId() *string {
	return s.TargetRegionId
}

func (s *CheckCloudResourceAuthorizedRequest) SetDBInstanceId(v string) *CheckCloudResourceAuthorizedRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetOwnerAccount(v string) *CheckCloudResourceAuthorizedRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetOwnerId(v int64) *CheckCloudResourceAuthorizedRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetResourceOwnerAccount(v string) *CheckCloudResourceAuthorizedRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetResourceOwnerId(v int64) *CheckCloudResourceAuthorizedRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) SetTargetRegionId(v string) *CheckCloudResourceAuthorizedRequest {
	s.TargetRegionId = &v
	return s
}

func (s *CheckCloudResourceAuthorizedRequest) Validate() error {
	return dara.Validate(s)
}
