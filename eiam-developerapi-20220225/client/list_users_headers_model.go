// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListUsersHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *ListUsersHeaders
	GetAuthorization() *string
}

type ListUsersHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. The value is in the Bearer ${access_token} format. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer AT8csE2seYxxxxxij
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListUsersHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListUsersHeaders) GoString() string {
	return s.String()
}

func (s *ListUsersHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListUsersHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListUsersHeaders) SetCommonHeaders(v map[string]*string) *ListUsersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListUsersHeaders) SetAuthorization(v string) *ListUsersHeaders {
	s.Authorization = &v
	return s
}

func (s *ListUsersHeaders) Validate() error {
	return dara.Validate(s)
}
