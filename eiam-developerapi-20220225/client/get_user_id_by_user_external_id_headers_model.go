// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByUserExternalIdHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetUserIdByUserExternalIdHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *GetUserIdByUserExternalIdHeaders
	GetAuthorization() *string
}

type GetUserIdByUserExternalIdHeaders struct {
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

func (s GetUserIdByUserExternalIdHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByUserExternalIdHeaders) GoString() string {
	return s.String()
}

func (s *GetUserIdByUserExternalIdHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetUserIdByUserExternalIdHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetUserIdByUserExternalIdHeaders) SetCommonHeaders(v map[string]*string) *GetUserIdByUserExternalIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserIdByUserExternalIdHeaders) SetAuthorization(v string) *GetUserIdByUserExternalIdHeaders {
	s.Authorization = &v
	return s
}

func (s *GetUserIdByUserExternalIdHeaders) Validate() error {
	return dara.Validate(s)
}
