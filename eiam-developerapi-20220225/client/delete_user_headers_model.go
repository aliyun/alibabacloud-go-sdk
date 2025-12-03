// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteUserHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *DeleteUserHeaders
	GetAuthorization() *string
}

type DeleteUserHeaders struct {
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

func (s DeleteUserHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserHeaders) GoString() string {
	return s.String()
}

func (s *DeleteUserHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteUserHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *DeleteUserHeaders) SetCommonHeaders(v map[string]*string) *DeleteUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteUserHeaders) SetAuthorization(v string) *DeleteUserHeaders {
	s.Authorization = &v
	return s
}

func (s *DeleteUserHeaders) Validate() error {
	return dara.Validate(s)
}
