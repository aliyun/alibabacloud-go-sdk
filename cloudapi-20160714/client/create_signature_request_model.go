// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSignatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityToken(v string) *CreateSignatureRequest
	GetSecurityToken() *string
	SetSignatureKey(v string) *CreateSignatureRequest
	GetSignatureKey() *string
	SetSignatureName(v string) *CreateSignatureRequest
	GetSignatureName() *string
	SetSignatureSecret(v string) *CreateSignatureRequest
	GetSignatureSecret() *string
}

type CreateSignatureRequest struct {
	// The security token included in the WebSocket request header. The system uses this token to authenticate the request.
	//
	// example:
	//
	// fa876ffb-caab-4f0a-93b3-3409f2fa5199
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The Key value of the key. The value must be 6 to 20 characters in length and can contain letters, digits, and underscores (_). It must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// qwertyuiop
	SignatureKey *string `json:"SignatureKey,omitempty" xml:"SignatureKey,omitempty"`
	// The displayed name of the key. The name must be 4 to 50 characters in length and can contain letters, digits, and underscores (_). It must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// backendsignature
	SignatureName *string `json:"SignatureName,omitempty" xml:"SignatureName,omitempty"`
	// The Secret value of the key. The value must be 6 to 30 characters in length and can contain letters, digits, and special characters. Special characters include underscores (_), at signs (@), number signs (#), exclamation points (!), and asterisks (\\*). The value must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// asdfghjkl
	SignatureSecret *string `json:"SignatureSecret,omitempty" xml:"SignatureSecret,omitempty"`
}

func (s CreateSignatureRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSignatureRequest) GoString() string {
	return s.String()
}

func (s *CreateSignatureRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateSignatureRequest) GetSignatureKey() *string {
	return s.SignatureKey
}

func (s *CreateSignatureRequest) GetSignatureName() *string {
	return s.SignatureName
}

func (s *CreateSignatureRequest) GetSignatureSecret() *string {
	return s.SignatureSecret
}

func (s *CreateSignatureRequest) SetSecurityToken(v string) *CreateSignatureRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateSignatureRequest) SetSignatureKey(v string) *CreateSignatureRequest {
	s.SignatureKey = &v
	return s
}

func (s *CreateSignatureRequest) SetSignatureName(v string) *CreateSignatureRequest {
	s.SignatureName = &v
	return s
}

func (s *CreateSignatureRequest) SetSignatureSecret(v string) *CreateSignatureRequest {
	s.SignatureSecret = &v
	return s
}

func (s *CreateSignatureRequest) Validate() error {
	return dara.Validate(s)
}
