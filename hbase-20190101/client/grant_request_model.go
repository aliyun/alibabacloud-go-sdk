// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *GrantRequest
	GetAccountName() *string
	SetAclActions(v string) *GrantRequest
	GetAclActions() *string
	SetClusterId(v string) *GrantRequest
	GetClusterId() *string
	SetNamespace(v string) *GrantRequest
	GetNamespace() *string
	SetTableName(v string) *GrantRequest
	GetTableName() *string
}

type GrantRequest struct {
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

func (s GrantRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantRequest) GoString() string {
	return s.String()
}

func (s *GrantRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *GrantRequest) GetAclActions() *string {
	return s.AclActions
}

func (s *GrantRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GrantRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GrantRequest) GetTableName() *string {
	return s.TableName
}

func (s *GrantRequest) SetAccountName(v string) *GrantRequest {
	s.AccountName = &v
	return s
}

func (s *GrantRequest) SetAclActions(v string) *GrantRequest {
	s.AclActions = &v
	return s
}

func (s *GrantRequest) SetClusterId(v string) *GrantRequest {
	s.ClusterId = &v
	return s
}

func (s *GrantRequest) SetNamespace(v string) *GrantRequest {
	s.Namespace = &v
	return s
}

func (s *GrantRequest) SetTableName(v string) *GrantRequest {
	s.TableName = &v
	return s
}

func (s *GrantRequest) Validate() error {
	return dara.Validate(s)
}
