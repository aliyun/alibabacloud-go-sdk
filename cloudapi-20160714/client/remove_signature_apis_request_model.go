// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSignatureApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiIds(v string) *RemoveSignatureApisRequest
	GetApiIds() *string
	SetGroupId(v string) *RemoveSignatureApisRequest
	GetGroupId() *string
	SetSecurityToken(v string) *RemoveSignatureApisRequest
	GetSecurityToken() *string
	SetSignatureId(v string) *RemoveSignatureApisRequest
	GetSignatureId() *string
	SetStageName(v string) *RemoveSignatureApisRequest
	GetStageName() *string
}

type RemoveSignatureApisRequest struct {
	// The IDs of the APIs from which you want to unbind the signature key.
	//
	// 	- If this parameter is not specified, the signature key is unbound from all the APIs in the specified environment of the API group.
	//
	// 	- The IDs of the APIs that you want to manage. Separate multiple API IDs with commas (,). A maximum of 100 API IDs can be entered.
	//
	// example:
	//
	// 123
	ApiIds *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
	// The ID of the API group to which the API that you want to manage belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0009db9c828549768a200320714b8930
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The ID of the signature key.
	//
	// This parameter is required.
	//
	// example:
	//
	// dd05f1c54d6749eda95f9fa6d491449a
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// The name of the runtime environment. Valid values:
	//
	// 	- **RELEASE**
	//
	// 	- **TEST**
	//
	// This parameter is required.
	//
	// example:
	//
	// TEST
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s RemoveSignatureApisRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveSignatureApisRequest) GoString() string {
	return s.String()
}

func (s *RemoveSignatureApisRequest) GetApiIds() *string {
	return s.ApiIds
}

func (s *RemoveSignatureApisRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *RemoveSignatureApisRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RemoveSignatureApisRequest) GetSignatureId() *string {
	return s.SignatureId
}

func (s *RemoveSignatureApisRequest) GetStageName() *string {
	return s.StageName
}

func (s *RemoveSignatureApisRequest) SetApiIds(v string) *RemoveSignatureApisRequest {
	s.ApiIds = &v
	return s
}

func (s *RemoveSignatureApisRequest) SetGroupId(v string) *RemoveSignatureApisRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveSignatureApisRequest) SetSecurityToken(v string) *RemoveSignatureApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *RemoveSignatureApisRequest) SetSignatureId(v string) *RemoveSignatureApisRequest {
	s.SignatureId = &v
	return s
}

func (s *RemoveSignatureApisRequest) SetStageName(v string) *RemoveSignatureApisRequest {
	s.StageName = &v
	return s
}

func (s *RemoveSignatureApisRequest) Validate() error {
	return dara.Validate(s)
}
