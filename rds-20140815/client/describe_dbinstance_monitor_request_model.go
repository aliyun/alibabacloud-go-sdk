// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeDBInstanceMonitorRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DescribeDBInstanceMonitorRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeDBInstanceMonitorRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstanceMonitorRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBInstanceMonitorRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceMonitorRequest
	GetResourceOwnerId() *int64
}

type DescribeDBInstanceMonitorRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBInstanceMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceMonitorRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceMonitorRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeDBInstanceMonitorRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceMonitorRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstanceMonitorRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstanceMonitorRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstanceMonitorRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceMonitorRequest) SetClientToken(v string) *DescribeDBInstanceMonitorRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeDBInstanceMonitorRequest) SetDBInstanceId(v string) *DescribeDBInstanceMonitorRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceMonitorRequest) SetOwnerAccount(v string) *DescribeDBInstanceMonitorRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceMonitorRequest) SetOwnerId(v int64) *DescribeDBInstanceMonitorRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceMonitorRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceMonitorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceMonitorRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceMonitorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceMonitorRequest) Validate() error {
	return dara.Validate(s)
}
