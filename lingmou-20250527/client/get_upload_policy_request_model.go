// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJwtToken(v string) *GetUploadPolicyRequest
	GetJwtToken() *string
	SetType(v string) *GetUploadPolicyRequest
	GetType() *string
}

type GetUploadPolicyRequest struct {
	// example:
	//
	// Token
	JwtToken *string `json:"jwtToken,omitempty" xml:"jwtToken,omitempty"`
	// example:
	//
	// INPUT_CHAT_BACKGROUND_PIC
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetUploadPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUploadPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetUploadPolicyRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *GetUploadPolicyRequest) GetType() *string {
	return s.Type
}

func (s *GetUploadPolicyRequest) SetJwtToken(v string) *GetUploadPolicyRequest {
	s.JwtToken = &v
	return s
}

func (s *GetUploadPolicyRequest) SetType(v string) *GetUploadPolicyRequest {
	s.Type = &v
	return s
}

func (s *GetUploadPolicyRequest) Validate() error {
	return dara.Validate(s)
}
