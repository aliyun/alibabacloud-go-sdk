// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsForUserHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListGroupsForUserHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *ListGroupsForUserHeaders
	GetAuthorization() *string
}

type ListGroupsForUserHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListGroupsForUserHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForUserHeaders) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListGroupsForUserHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListGroupsForUserHeaders) SetCommonHeaders(v map[string]*string) *ListGroupsForUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListGroupsForUserHeaders) SetAuthorization(v string) *ListGroupsForUserHeaders {
	s.Authorization = &v
	return s
}

func (s *ListGroupsForUserHeaders) Validate() error {
	return dara.Validate(s)
}
