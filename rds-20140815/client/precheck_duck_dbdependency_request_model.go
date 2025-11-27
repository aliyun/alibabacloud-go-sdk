// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrecheckDuckDBDependencyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *PrecheckDuckDBDependencyRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *PrecheckDuckDBDependencyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *PrecheckDuckDBDependencyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *PrecheckDuckDBDependencyRequest
	GetResourceOwnerId() *int64
}

type PrecheckDuckDBDependencyRequest struct {
	// The primary instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-2zegy5pdkg58****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PrecheckDuckDBDependencyRequest) String() string {
	return dara.Prettify(s)
}

func (s PrecheckDuckDBDependencyRequest) GoString() string {
	return s.String()
}

func (s *PrecheckDuckDBDependencyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *PrecheckDuckDBDependencyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PrecheckDuckDBDependencyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *PrecheckDuckDBDependencyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *PrecheckDuckDBDependencyRequest) SetDBInstanceId(v string) *PrecheckDuckDBDependencyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *PrecheckDuckDBDependencyRequest) SetOwnerId(v int64) *PrecheckDuckDBDependencyRequest {
	s.OwnerId = &v
	return s
}

func (s *PrecheckDuckDBDependencyRequest) SetResourceOwnerAccount(v string) *PrecheckDuckDBDependencyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PrecheckDuckDBDependencyRequest) SetResourceOwnerId(v int64) *PrecheckDuckDBDependencyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PrecheckDuckDBDependencyRequest) Validate() error {
	return dara.Validate(s)
}
