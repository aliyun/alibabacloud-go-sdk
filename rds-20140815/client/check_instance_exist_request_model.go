// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInstanceExistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CheckInstanceExistRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *CheckInstanceExistRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CheckInstanceExistRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckInstanceExistRequest
	GetResourceOwnerId() *int64
}

type CheckInstanceExistRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckInstanceExistRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceExistRequest) GoString() string {
	return s.String()
}

func (s *CheckInstanceExistRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CheckInstanceExistRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckInstanceExistRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckInstanceExistRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckInstanceExistRequest) SetDBInstanceId(v string) *CheckInstanceExistRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CheckInstanceExistRequest) SetOwnerId(v int64) *CheckInstanceExistRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckInstanceExistRequest) SetResourceOwnerAccount(v string) *CheckInstanceExistRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckInstanceExistRequest) SetResourceOwnerId(v int64) *CheckInstanceExistRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckInstanceExistRequest) Validate() error {
	return dara.Validate(s)
}
