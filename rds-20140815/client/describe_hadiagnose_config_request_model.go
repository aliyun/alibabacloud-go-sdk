// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHADiagnoseConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeHADiagnoseConfigRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DescribeHADiagnoseConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeHADiagnoseConfigRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeHADiagnoseConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeHADiagnoseConfigRequest
	GetResourceOwnerId() *int64
}

type DescribeHADiagnoseConfigRequest struct {
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeHADiagnoseConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHADiagnoseConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeHADiagnoseConfigRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeHADiagnoseConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeHADiagnoseConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHADiagnoseConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeHADiagnoseConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeHADiagnoseConfigRequest) SetDBInstanceId(v string) *DescribeHADiagnoseConfigRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeHADiagnoseConfigRequest) SetOwnerId(v int64) *DescribeHADiagnoseConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHADiagnoseConfigRequest) SetRegionId(v string) *DescribeHADiagnoseConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHADiagnoseConfigRequest) SetResourceOwnerAccount(v string) *DescribeHADiagnoseConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHADiagnoseConfigRequest) SetResourceOwnerId(v int64) *DescribeHADiagnoseConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeHADiagnoseConfigRequest) Validate() error {
	return dara.Validate(s)
}
