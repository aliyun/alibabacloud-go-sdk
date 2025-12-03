// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByUsernameHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetUserIdByUsernameHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *GetUserIdByUsernameHeaders
	GetAuthorization() *string
}

type GetUserIdByUsernameHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer ${access_token}. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetUserIdByUsernameHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByUsernameHeaders) GoString() string {
	return s.String()
}

func (s *GetUserIdByUsernameHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetUserIdByUsernameHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetUserIdByUsernameHeaders) SetCommonHeaders(v map[string]*string) *GetUserIdByUsernameHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserIdByUsernameHeaders) SetAuthorization(v string) *GetUserIdByUsernameHeaders {
	s.Authorization = &v
	return s
}

func (s *GetUserIdByUsernameHeaders) Validate() error {
	return dara.Validate(s)
}
