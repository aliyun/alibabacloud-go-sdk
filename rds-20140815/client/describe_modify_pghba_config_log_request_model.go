// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModifyPGHbaConfigLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeModifyPGHbaConfigLogRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DescribeModifyPGHbaConfigLogRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeModifyPGHbaConfigLogRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *DescribeModifyPGHbaConfigLogRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeModifyPGHbaConfigLogRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DescribeModifyPGHbaConfigLogRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeModifyPGHbaConfigLogRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeModifyPGHbaConfigLogRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeModifyPGHbaConfigLogRequest
	GetStartTime() *string
}

type DescribeModifyPGHbaConfigLogRequest struct {
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// 1
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp1lymyn1v3i****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// example:
	//
	// 2021-11-25T06:00:40Z
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in Coordinated Universal Time (UTC).
	//
	// example:
	//
	// 2021-11-25T05:00:40Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeModifyPGHbaConfigLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeModifyPGHbaConfigLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeModifyPGHbaConfigLogRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeModifyPGHbaConfigLogRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeModifyPGHbaConfigLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeModifyPGHbaConfigLogRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeModifyPGHbaConfigLogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeModifyPGHbaConfigLogRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeModifyPGHbaConfigLogRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeModifyPGHbaConfigLogRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeModifyPGHbaConfigLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeModifyPGHbaConfigLogRequest) SetClientToken(v string) *DescribeModifyPGHbaConfigLogRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogRequest) SetDBInstanceId(v string) *DescribeModifyPGHbaConfigLogRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogRequest) SetEndTime(v string) *DescribeModifyPGHbaConfigLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogRequest) SetOwnerAccount(v string) *DescribeModifyPGHbaConfigLogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogRequest) SetOwnerId(v int64) *DescribeModifyPGHbaConfigLogRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogRequest) SetResourceGroupId(v string) *DescribeModifyPGHbaConfigLogRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogRequest) SetResourceOwnerAccount(v string) *DescribeModifyPGHbaConfigLogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogRequest) SetResourceOwnerId(v int64) *DescribeModifyPGHbaConfigLogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogRequest) SetStartTime(v string) *DescribeModifyPGHbaConfigLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeModifyPGHbaConfigLogRequest) Validate() error {
	return dara.Validate(s)
}
