// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySignatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityToken(v string) *ModifySignatureRequest
	GetSecurityToken() *string
	SetSignatureId(v string) *ModifySignatureRequest
	GetSignatureId() *string
	SetSignatureKey(v string) *ModifySignatureRequest
	GetSignatureKey() *string
	SetSignatureName(v string) *ModifySignatureRequest
	GetSignatureName() *string
	SetSignatureSecret(v string) *ModifySignatureRequest
	GetSignatureSecret() *string
}

type ModifySignatureRequest struct {
	// The security token included in the WebSocket request header. The system uses this token to authenticate the request.
	//
	// example:
	//
	// 436fa39b-b3b9-40c5-ae5d-ce3e000e38c5
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The ID of the signature key that you want to manage.
	//
	// This parameter is required.
	//
	// example:
	//
	// dd05f1c54d6749eda95f9fa6d491449a
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	// The new Key value of the key. The value must be 6 to 20 characters in length and can contain letters, digits, and underscores (_). It must start with a letter.
	//
	// example:
	//
	// qwertyuiop
	SignatureKey *string `json:"SignatureKey,omitempty" xml:"SignatureKey,omitempty"`
	// The new name of the key. The name must be 4 to 50 characters in length and can contain letters, digits, and underscores (_). It must start with a letter.
	//
	// example:
	//
	// backendsignature
	SignatureName *string `json:"SignatureName,omitempty" xml:"SignatureName,omitempty"`
	// The new Secret value of the key. The value must be 6 to 30 characters in length and can contain letters, digits, and special characters. Special characters include underscores (_), at signs (@), number signs (#), exclamation points (!), and asterisks (\\*). The value must start with a letter.
	//
	// example:
	//
	// asdfghjkl
	SignatureSecret *string `json:"SignatureSecret,omitempty" xml:"SignatureSecret,omitempty"`
}

func (s ModifySignatureRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySignatureRequest) GoString() string {
	return s.String()
}

func (s *ModifySignatureRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifySignatureRequest) GetSignatureId() *string {
	return s.SignatureId
}

func (s *ModifySignatureRequest) GetSignatureKey() *string {
	return s.SignatureKey
}

func (s *ModifySignatureRequest) GetSignatureName() *string {
	return s.SignatureName
}

func (s *ModifySignatureRequest) GetSignatureSecret() *string {
	return s.SignatureSecret
}

func (s *ModifySignatureRequest) SetSecurityToken(v string) *ModifySignatureRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifySignatureRequest) SetSignatureId(v string) *ModifySignatureRequest {
	s.SignatureId = &v
	return s
}

func (s *ModifySignatureRequest) SetSignatureKey(v string) *ModifySignatureRequest {
	s.SignatureKey = &v
	return s
}

func (s *ModifySignatureRequest) SetSignatureName(v string) *ModifySignatureRequest {
	s.SignatureName = &v
	return s
}

func (s *ModifySignatureRequest) SetSignatureSecret(v string) *ModifySignatureRequest {
	s.SignatureSecret = &v
	return s
}

func (s *ModifySignatureRequest) Validate() error {
	return dara.Validate(s)
}
