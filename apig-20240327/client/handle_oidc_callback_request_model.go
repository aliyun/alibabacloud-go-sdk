// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleOidcCallbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HandleOidcCallbackRequest
	GetCode() *string
	SetState(v string) *HandleOidcCallbackRequest
	GetState() *string
}

type HandleOidcCallbackRequest struct {
	// The code parameter carried in the OIDC callback request.
	//
	// example:
	//
	// code-xxxxxxxx
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The state parameter carried in the OIDC callback request.
	//
	// example:
	//
	// state-xxxxxxxx
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s HandleOidcCallbackRequest) String() string {
	return dara.Prettify(s)
}

func (s HandleOidcCallbackRequest) GoString() string {
	return s.String()
}

func (s *HandleOidcCallbackRequest) GetCode() *string {
	return s.Code
}

func (s *HandleOidcCallbackRequest) GetState() *string {
	return s.State
}

func (s *HandleOidcCallbackRequest) SetCode(v string) *HandleOidcCallbackRequest {
	s.Code = &v
	return s
}

func (s *HandleOidcCallbackRequest) SetState(v string) *HandleOidcCallbackRequest {
	s.State = &v
	return s
}

func (s *HandleOidcCallbackRequest) Validate() error {
	return dara.Validate(s)
}
