// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableEngineVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeAvailableEngineVersionRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeAvailableEngineVersionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAvailableEngineVersionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeAvailableEngineVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAvailableEngineVersionRequest
	GetResourceOwnerId() *int64
}

type DescribeAvailableEngineVersionRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bpxxxxxxxx
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAvailableEngineVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableEngineVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableEngineVersionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeAvailableEngineVersionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAvailableEngineVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAvailableEngineVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAvailableEngineVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAvailableEngineVersionRequest) SetDBInstanceId(v string) *DescribeAvailableEngineVersionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAvailableEngineVersionRequest) SetOwnerAccount(v string) *DescribeAvailableEngineVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAvailableEngineVersionRequest) SetOwnerId(v int64) *DescribeAvailableEngineVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAvailableEngineVersionRequest) SetResourceOwnerAccount(v string) *DescribeAvailableEngineVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAvailableEngineVersionRequest) SetResourceOwnerId(v int64) *DescribeAvailableEngineVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAvailableEngineVersionRequest) Validate() error {
	return dara.Validate(s)
}
