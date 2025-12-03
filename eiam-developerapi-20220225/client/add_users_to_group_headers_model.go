// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUsersToGroupHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddUsersToGroupHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *AddUsersToGroupHeaders
	GetAuthorization() *string
}

type AddUsersToGroupHeaders struct {
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

func (s AddUsersToGroupHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddUsersToGroupHeaders) GoString() string {
	return s.String()
}

func (s *AddUsersToGroupHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddUsersToGroupHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *AddUsersToGroupHeaders) SetCommonHeaders(v map[string]*string) *AddUsersToGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddUsersToGroupHeaders) SetAuthorization(v string) *AddUsersToGroupHeaders {
	s.Authorization = &v
	return s
}

func (s *AddUsersToGroupHeaders) Validate() error {
	return dara.Validate(s)
}
