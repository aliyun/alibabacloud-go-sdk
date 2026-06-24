// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOpenSourcePermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListOpenSourcePermissionsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListOpenSourcePermissionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListOpenSourcePermissionsRequest
	GetNextToken() *string
	SetUserName(v string) *ListOpenSourcePermissionsRequest
	GetUserName() *string
}

type ListOpenSourcePermissionsRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-v0h1kb9nu***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of results to return.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The username.
	//
	// This parameter is required.
	//
	// example:
	//
	// myName
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListOpenSourcePermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOpenSourcePermissionsRequest) GoString() string {
	return s.String()
}

func (s *ListOpenSourcePermissionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOpenSourcePermissionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListOpenSourcePermissionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOpenSourcePermissionsRequest) GetUserName() *string {
	return s.UserName
}

func (s *ListOpenSourcePermissionsRequest) SetInstanceId(v string) *ListOpenSourcePermissionsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOpenSourcePermissionsRequest) SetMaxResults(v int32) *ListOpenSourcePermissionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListOpenSourcePermissionsRequest) SetNextToken(v string) *ListOpenSourcePermissionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListOpenSourcePermissionsRequest) SetUserName(v string) *ListOpenSourcePermissionsRequest {
	s.UserName = &v
	return s
}

func (s *ListOpenSourcePermissionsRequest) Validate() error {
	return dara.Validate(s)
}
