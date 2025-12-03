// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListGroupsHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *ListGroupsHeaders
	GetAuthorization() *string
}

type ListGroupsHeaders struct {
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

func (s ListGroupsHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsHeaders) GoString() string {
	return s.String()
}

func (s *ListGroupsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListGroupsHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListGroupsHeaders) SetCommonHeaders(v map[string]*string) *ListGroupsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListGroupsHeaders) SetAuthorization(v string) *ListGroupsHeaders {
	s.Authorization = &v
	return s
}

func (s *ListGroupsHeaders) Validate() error {
	return dara.Validate(s)
}
