// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCredentialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListCredentialsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListCredentialsRequest
	GetMaxResults() *int32
	SetNameLike(v string) *ListCredentialsRequest
	GetNameLike() *string
	SetNextToken(v string) *ListCredentialsRequest
	GetNextToken() *string
}

type ListCredentialsRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NameLike   *string `json:"NameLike,omitempty" xml:"NameLike,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListCredentialsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialsRequest) GoString() string {
	return s.String()
}

func (s *ListCredentialsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCredentialsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCredentialsRequest) GetNameLike() *string {
	return s.NameLike
}

func (s *ListCredentialsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCredentialsRequest) SetInstanceId(v string) *ListCredentialsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCredentialsRequest) SetMaxResults(v int32) *ListCredentialsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCredentialsRequest) SetNameLike(v string) *ListCredentialsRequest {
	s.NameLike = &v
	return s
}

func (s *ListCredentialsRequest) SetNextToken(v string) *ListCredentialsRequest {
	s.NextToken = &v
	return s
}

func (s *ListCredentialsRequest) Validate() error {
	return dara.Validate(s)
}
