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
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
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
