// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountParameterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ModifyAccountParameterRequest
	GetAccountName() *string
	SetInstanceId(v string) *ModifyAccountParameterRequest
	GetInstanceId() *string
	SetParameters(v string) *ModifyAccountParameterRequest
	GetParameters() *string
	SetSecurityToken(v string) *ModifyAccountParameterRequest
	GetSecurityToken() *string
}

type ModifyAccountParameterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// demoaccount
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// r-bp1s4h3oosz5y8ilc7
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"access-db-id":"1","cu-limit":"10"}
	Parameters    *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyAccountParameterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountParameterRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountParameterRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyAccountParameterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyAccountParameterRequest) GetParameters() *string {
	return s.Parameters
}

func (s *ModifyAccountParameterRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyAccountParameterRequest) SetAccountName(v string) *ModifyAccountParameterRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountParameterRequest) SetInstanceId(v string) *ModifyAccountParameterRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyAccountParameterRequest) SetParameters(v string) *ModifyAccountParameterRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyAccountParameterRequest) SetSecurityToken(v string) *ModifyAccountParameterRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyAccountParameterRequest) Validate() error {
	return dara.Validate(s)
}
