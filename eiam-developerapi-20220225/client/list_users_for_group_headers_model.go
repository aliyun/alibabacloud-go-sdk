// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersForGroupHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListUsersForGroupHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *ListUsersForGroupHeaders
	GetAuthorization() *string
}

type ListUsersForGroupHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. The value is in the Bearer ${access_token} format. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListUsersForGroupHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForGroupHeaders) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListUsersForGroupHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListUsersForGroupHeaders) SetCommonHeaders(v map[string]*string) *ListUsersForGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListUsersForGroupHeaders) SetAuthorization(v string) *ListUsersForGroupHeaders {
	s.Authorization = &v
	return s
}

func (s *ListUsersForGroupHeaders) Validate() error {
	return dara.Validate(s)
}
