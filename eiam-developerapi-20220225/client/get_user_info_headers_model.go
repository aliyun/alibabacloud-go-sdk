// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetUserInfoHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *GetUserInfoHeaders
	GetAuthorization() *string
}

type GetUserInfoHeaders struct {
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

func (s GetUserInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetUserInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetUserInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetUserInfoHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetUserInfoHeaders) SetCommonHeaders(v map[string]*string) *GetUserInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserInfoHeaders) SetAuthorization(v string) *GetUserInfoHeaders {
	s.Authorization = &v
	return s
}

func (s *GetUserInfoHeaders) Validate() error {
	return dara.Validate(s)
}
