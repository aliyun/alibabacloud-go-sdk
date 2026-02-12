// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainCredentialHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ObtainCredentialHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *ObtainCredentialHeaders
	GetAuthorization() *string
}

type ObtainCredentialHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ObtainCredentialHeaders) String() string {
	return dara.Prettify(s)
}

func (s ObtainCredentialHeaders) GoString() string {
	return s.String()
}

func (s *ObtainCredentialHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ObtainCredentialHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ObtainCredentialHeaders) SetCommonHeaders(v map[string]*string) *ObtainCredentialHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ObtainCredentialHeaders) SetAuthorization(v string) *ObtainCredentialHeaders {
	s.Authorization = &v
	return s
}

func (s *ObtainCredentialHeaders) Validate() error {
	return dara.Validate(s)
}
