// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *RevokeRequest
	GetAccountName() *string
	SetAclActions(v string) *RevokeRequest
	GetAclActions() *string
	SetClusterId(v string) *RevokeRequest
	GetClusterId() *string
	SetNamespace(v string) *RevokeRequest
	GetNamespace() *string
	SetTableName(v string) *RevokeRequest
	GetTableName() *string
}

type RevokeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test01
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// READ,WRITE
	AclActions *string `json:"AclActions,omitempty" xml:"AclActions,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s RevokeRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeRequest) GoString() string {
	return s.String()
}

func (s *RevokeRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *RevokeRequest) GetAclActions() *string {
	return s.AclActions
}

func (s *RevokeRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *RevokeRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *RevokeRequest) GetTableName() *string {
	return s.TableName
}

func (s *RevokeRequest) SetAccountName(v string) *RevokeRequest {
	s.AccountName = &v
	return s
}

func (s *RevokeRequest) SetAclActions(v string) *RevokeRequest {
	s.AclActions = &v
	return s
}

func (s *RevokeRequest) SetClusterId(v string) *RevokeRequest {
	s.ClusterId = &v
	return s
}

func (s *RevokeRequest) SetNamespace(v string) *RevokeRequest {
	s.Namespace = &v
	return s
}

func (s *RevokeRequest) SetTableName(v string) *RevokeRequest {
	s.TableName = &v
	return s
}

func (s *RevokeRequest) Validate() error {
	return dara.Validate(s)
}
