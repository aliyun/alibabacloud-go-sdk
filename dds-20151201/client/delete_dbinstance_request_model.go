// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteDBInstanceRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DeleteDBInstanceRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DeleteDBInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteDBInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteDBInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteDBInstanceRequest
	GetResourceOwnerId() *int64
}

type DeleteDBInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
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

func (s DeleteDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDBInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteDBInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDBInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDBInstanceRequest) SetClientToken(v string) *DeleteDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetDBInstanceId(v string) *DeleteDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetOwnerAccount(v string) *DeleteDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetOwnerId(v int64) *DeleteDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetResourceOwnerAccount(v string) *DeleteDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetResourceOwnerId(v int64) *DeleteDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
