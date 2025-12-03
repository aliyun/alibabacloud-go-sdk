// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUsersFromGroupHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *RemoveUsersFromGroupHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *RemoveUsersFromGroupHeaders
	GetAuthorization() *string
}

type RemoveUsersFromGroupHeaders struct {
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

func (s RemoveUsersFromGroupHeaders) String() string {
	return dara.Prettify(s)
}

func (s RemoveUsersFromGroupHeaders) GoString() string {
	return s.String()
}

func (s *RemoveUsersFromGroupHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *RemoveUsersFromGroupHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *RemoveUsersFromGroupHeaders) SetCommonHeaders(v map[string]*string) *RemoveUsersFromGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveUsersFromGroupHeaders) SetAuthorization(v string) *RemoveUsersFromGroupHeaders {
	s.Authorization = &v
	return s
}

func (s *RemoveUsersFromGroupHeaders) Validate() error {
	return dara.Validate(s)
}
