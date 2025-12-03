// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateGroupHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *CreateGroupHeaders
	GetAuthorization() *string
}

type CreateGroupHeaders struct {
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

func (s CreateGroupHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupHeaders) GoString() string {
	return s.String()
}

func (s *CreateGroupHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateGroupHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *CreateGroupHeaders) SetCommonHeaders(v map[string]*string) *CreateGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateGroupHeaders) SetAuthorization(v string) *CreateGroupHeaders {
	s.Authorization = &v
	return s
}

func (s *CreateGroupHeaders) Validate() error {
	return dara.Validate(s)
}
