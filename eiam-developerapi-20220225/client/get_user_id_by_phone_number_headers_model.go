// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByPhoneNumberHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetUserIdByPhoneNumberHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *GetUserIdByPhoneNumberHeaders
	GetAuthorization() *string
}

type GetUserIdByPhoneNumberHeaders struct {
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

func (s GetUserIdByPhoneNumberHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByPhoneNumberHeaders) GoString() string {
	return s.String()
}

func (s *GetUserIdByPhoneNumberHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetUserIdByPhoneNumberHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetUserIdByPhoneNumberHeaders) SetCommonHeaders(v map[string]*string) *GetUserIdByPhoneNumberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserIdByPhoneNumberHeaders) SetAuthorization(v string) *GetUserIdByPhoneNumberHeaders {
	s.Authorization = &v
	return s
}

func (s *GetUserIdByPhoneNumberHeaders) Validate() error {
	return dara.Validate(s)
}
