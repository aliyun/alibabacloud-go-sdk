// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetGroupHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *GetGroupHeaders
	GetAuthorization() *string
}

type GetGroupHeaders struct {
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

func (s GetGroupHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetGroupHeaders) GoString() string {
	return s.String()
}

func (s *GetGroupHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetGroupHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetGroupHeaders) SetCommonHeaders(v map[string]*string) *GetGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetGroupHeaders) SetAuthorization(v string) *GetGroupHeaders {
	s.Authorization = &v
	return s
}

func (s *GetGroupHeaders) Validate() error {
	return dara.Validate(s)
}
