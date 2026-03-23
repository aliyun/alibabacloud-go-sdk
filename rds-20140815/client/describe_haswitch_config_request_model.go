// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHASwitchConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeHASwitchConfigRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DescribeHASwitchConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeHASwitchConfigRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeHASwitchConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeHASwitchConfigRequest
	GetResourceOwnerId() *int64
}

type DescribeHASwitchConfigRequest struct {
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeHASwitchConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHASwitchConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeHASwitchConfigRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeHASwitchConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeHASwitchConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHASwitchConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeHASwitchConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeHASwitchConfigRequest) SetDBInstanceId(v string) *DescribeHASwitchConfigRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeHASwitchConfigRequest) SetOwnerId(v int64) *DescribeHASwitchConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHASwitchConfigRequest) SetRegionId(v string) *DescribeHASwitchConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHASwitchConfigRequest) SetResourceOwnerAccount(v string) *DescribeHASwitchConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHASwitchConfigRequest) SetResourceOwnerId(v int64) *DescribeHASwitchConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeHASwitchConfigRequest) Validate() error {
	return dara.Validate(s)
}
