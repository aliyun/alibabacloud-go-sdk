// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateUserHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *CreateUserHeaders
	GetAuthorization() *string
}

type CreateUserHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer ${access_token}. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer AT8csE2seYxxxxxij
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CreateUserHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateUserHeaders) GoString() string {
	return s.String()
}

func (s *CreateUserHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateUserHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *CreateUserHeaders) SetCommonHeaders(v map[string]*string) *CreateUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateUserHeaders) SetAuthorization(v string) *CreateUserHeaders {
	s.Authorization = &v
	return s
}

func (s *CreateUserHeaders) Validate() error {
	return dara.Validate(s)
}
