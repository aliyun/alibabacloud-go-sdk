// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserExclusiveCredentialHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateUserExclusiveCredentialHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *CreateUserExclusiveCredentialHeaders
	GetAuthorization() *string
}

type CreateUserExclusiveCredentialHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CreateUserExclusiveCredentialHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateUserExclusiveCredentialHeaders) GoString() string {
	return s.String()
}

func (s *CreateUserExclusiveCredentialHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateUserExclusiveCredentialHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *CreateUserExclusiveCredentialHeaders) SetCommonHeaders(v map[string]*string) *CreateUserExclusiveCredentialHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateUserExclusiveCredentialHeaders) SetAuthorization(v string) *CreateUserExclusiveCredentialHeaders {
	s.Authorization = &v
	return s
}

func (s *CreateUserExclusiveCredentialHeaders) Validate() error {
	return dara.Validate(s)
}
