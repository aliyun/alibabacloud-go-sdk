// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurgeDBInstanceLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *PurgeDBInstanceLogRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *PurgeDBInstanceLogRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *PurgeDBInstanceLogRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *PurgeDBInstanceLogRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *PurgeDBInstanceLogRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *PurgeDBInstanceLogRequest
	GetResourceOwnerId() *int64
}

type PurgeDBInstanceLogRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
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

func (s PurgeDBInstanceLogRequest) String() string {
	return dara.Prettify(s)
}

func (s PurgeDBInstanceLogRequest) GoString() string {
	return s.String()
}

func (s *PurgeDBInstanceLogRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *PurgeDBInstanceLogRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *PurgeDBInstanceLogRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *PurgeDBInstanceLogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PurgeDBInstanceLogRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *PurgeDBInstanceLogRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *PurgeDBInstanceLogRequest) SetClientToken(v string) *PurgeDBInstanceLogRequest {
	s.ClientToken = &v
	return s
}

func (s *PurgeDBInstanceLogRequest) SetDBInstanceId(v string) *PurgeDBInstanceLogRequest {
	s.DBInstanceId = &v
	return s
}

func (s *PurgeDBInstanceLogRequest) SetOwnerAccount(v string) *PurgeDBInstanceLogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *PurgeDBInstanceLogRequest) SetOwnerId(v int64) *PurgeDBInstanceLogRequest {
	s.OwnerId = &v
	return s
}

func (s *PurgeDBInstanceLogRequest) SetResourceOwnerAccount(v string) *PurgeDBInstanceLogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PurgeDBInstanceLogRequest) SetResourceOwnerId(v int64) *PurgeDBInstanceLogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PurgeDBInstanceLogRequest) Validate() error {
	return dara.Validate(s)
}
