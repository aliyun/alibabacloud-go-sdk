// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSignatureApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiIds(v string) *SetSignatureApisRequest
	GetApiIds() *string
	SetGroupId(v string) *SetSignatureApisRequest
	GetGroupId() *string
	SetSecurityToken(v string) *SetSignatureApisRequest
	GetSecurityToken() *string
	SetSignatureId(v string) *SetSignatureApisRequest
	GetSignatureId() *string
	SetStageName(v string) *SetSignatureApisRequest
	GetStageName() *string
}

type SetSignatureApisRequest struct {
	// The API IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3b81fd160f5645e097cc8855d75a1cf6,46fbb52840d146f186e38e8e70fc8c12
	ApiIds *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
	// The API group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0009db9c828549768a200320714b8930
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The signature ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dd05f1c54d6749eda95f9fa6d491449a
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// The environment. Valid values:
	//
	// 	- **RELEASE**: the production environment
	//
	// 	- **PRE**: the staging environment
	//
	// 	- **TEST**: the testing environment
	//
	// This parameter is required.
	//
	// example:
	//
	// TEST
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s SetSignatureApisRequest) String() string {
	return dara.Prettify(s)
}

func (s SetSignatureApisRequest) GoString() string {
	return s.String()
}

func (s *SetSignatureApisRequest) GetApiIds() *string {
	return s.ApiIds
}

func (s *SetSignatureApisRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *SetSignatureApisRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SetSignatureApisRequest) GetSignatureId() *string {
	return s.SignatureId
}

func (s *SetSignatureApisRequest) GetStageName() *string {
	return s.StageName
}

func (s *SetSignatureApisRequest) SetApiIds(v string) *SetSignatureApisRequest {
	s.ApiIds = &v
	return s
}

func (s *SetSignatureApisRequest) SetGroupId(v string) *SetSignatureApisRequest {
	s.GroupId = &v
	return s
}

func (s *SetSignatureApisRequest) SetSecurityToken(v string) *SetSignatureApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetSignatureApisRequest) SetSignatureId(v string) *SetSignatureApisRequest {
	s.SignatureId = &v
	return s
}

func (s *SetSignatureApisRequest) SetStageName(v string) *SetSignatureApisRequest {
	s.StageName = &v
	return s
}

func (s *SetSignatureApisRequest) Validate() error {
	return dara.Validate(s)
}
