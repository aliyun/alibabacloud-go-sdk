// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableUserHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DisableUserHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *DisableUserHeaders
	GetAuthorization() *string
}

type DisableUserHeaders struct {
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

func (s DisableUserHeaders) String() string {
	return dara.Prettify(s)
}

func (s DisableUserHeaders) GoString() string {
	return s.String()
}

func (s *DisableUserHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DisableUserHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *DisableUserHeaders) SetCommonHeaders(v map[string]*string) *DisableUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DisableUserHeaders) SetAuthorization(v string) *DisableUserHeaders {
	s.Authorization = &v
	return s
}

func (s *DisableUserHeaders) Validate() error {
	return dara.Validate(s)
}
