// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSignatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityToken(v string) *DeleteSignatureRequest
	GetSecurityToken() *string
	SetSignatureId(v string) *DeleteSignatureRequest
	GetSignatureId() *string
}

type DeleteSignatureRequest struct {
	// The security token included in the WebSocket request header. The system uses this token to authenticate the request.
	//
	// example:
	//
	// 9abe3317-3e22-4957-ab9f-dd893d0ac956
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The ID of the key to be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// dd05f1c54d6749eda95f9fa6d491449a
	SignatureId *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
}

func (s DeleteSignatureRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSignatureRequest) GoString() string {
	return s.String()
}

func (s *DeleteSignatureRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteSignatureRequest) GetSignatureId() *string {
	return s.SignatureId
}

func (s *DeleteSignatureRequest) SetSecurityToken(v string) *DeleteSignatureRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteSignatureRequest) SetSignatureId(v string) *DeleteSignatureRequest {
	s.SignatureId = &v
	return s
}

func (s *DeleteSignatureRequest) Validate() error {
	return dara.Validate(s)
}
