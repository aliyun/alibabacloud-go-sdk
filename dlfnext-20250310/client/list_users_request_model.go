// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListUsersRequest
	GetMaxResults() *int32
	SetPageToken(v string) *ListUsersRequest
	GetPageToken() *string
	SetType(v string) *ListUsersRequest
	GetType() *string
	SetUserName(v string) *ListUsersRequest
	GetUserName() *string
}

type ListUsersRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	PageToken  *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
	UserName   *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s ListUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUsersRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *ListUsersRequest) GetType() *string {
	return s.Type
}

func (s *ListUsersRequest) GetUserName() *string {
	return s.UserName
}

func (s *ListUsersRequest) SetMaxResults(v int32) *ListUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUsersRequest) SetPageToken(v string) *ListUsersRequest {
	s.PageToken = &v
	return s
}

func (s *ListUsersRequest) SetType(v string) *ListUsersRequest {
	s.Type = &v
	return s
}

func (s *ListUsersRequest) SetUserName(v string) *ListUsersRequest {
	s.UserName = &v
	return s
}

func (s *ListUsersRequest) Validate() error {
	return dara.Validate(s)
}
