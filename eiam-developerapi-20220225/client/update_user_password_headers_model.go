// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserPasswordHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateUserPasswordHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *UpdateUserPasswordHeaders
	GetAuthorization() *string
}

type UpdateUserPasswordHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateUserPasswordHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserPasswordHeaders) GoString() string {
	return s.String()
}

func (s *UpdateUserPasswordHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateUserPasswordHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *UpdateUserPasswordHeaders) SetCommonHeaders(v map[string]*string) *UpdateUserPasswordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateUserPasswordHeaders) SetAuthorization(v string) *UpdateUserPasswordHeaders {
	s.Authorization = &v
	return s
}

func (s *UpdateUserPasswordHeaders) Validate() error {
	return dara.Validate(s)
}
