// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByEmailHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetUserIdByEmailHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *GetUserIdByEmailHeaders
	GetAuthorization() *string
}

type GetUserIdByEmailHeaders struct {
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

func (s GetUserIdByEmailHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByEmailHeaders) GoString() string {
	return s.String()
}

func (s *GetUserIdByEmailHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetUserIdByEmailHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetUserIdByEmailHeaders) SetCommonHeaders(v map[string]*string) *GetUserIdByEmailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserIdByEmailHeaders) SetAuthorization(v string) *GetUserIdByEmailHeaders {
	s.Authorization = &v
	return s
}

func (s *GetUserIdByEmailHeaders) Validate() error {
	return dara.Validate(s)
}
