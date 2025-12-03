// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGroupHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteGroupHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *DeleteGroupHeaders
	GetAuthorization() *string
}

type DeleteGroupHeaders struct {
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

func (s DeleteGroupHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteGroupHeaders) GoString() string {
	return s.String()
}

func (s *DeleteGroupHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteGroupHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *DeleteGroupHeaders) SetCommonHeaders(v map[string]*string) *DeleteGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteGroupHeaders) SetAuthorization(v string) *DeleteGroupHeaders {
	s.Authorization = &v
	return s
}

func (s *DeleteGroupHeaders) Validate() error {
	return dara.Validate(s)
}
