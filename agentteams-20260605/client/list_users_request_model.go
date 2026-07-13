// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListUsersRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListUsersRequest
	GetMaxResults() *int32
	SetNameLike(v string) *ListUsersRequest
	GetNameLike() *string
	SetNextToken(v string) *ListUsersRequest
	GetNextToken() *string
}

type ListUsersRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NameLike   *string `json:"NameLike,omitempty" xml:"NameLike,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUsersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUsersRequest) GetNameLike() *string {
	return s.NameLike
}

func (s *ListUsersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUsersRequest) SetInstanceId(v string) *ListUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUsersRequest) SetMaxResults(v int32) *ListUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUsersRequest) SetNameLike(v string) *ListUsersRequest {
	s.NameLike = &v
	return s
}

func (s *ListUsersRequest) SetNextToken(v string) *ListUsersRequest {
	s.NextToken = &v
	return s
}

func (s *ListUsersRequest) Validate() error {
	return dara.Validate(s)
}
