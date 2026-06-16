// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizationRulesForUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListAuthorizationRulesForUserRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListAuthorizationRulesForUserRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAuthorizationRulesForUserRequest
	GetNextToken() *string
	SetUserId(v string) *ListAuthorizationRulesForUserRequest
	GetUserId() *string
}

type ListAuthorizationRulesForUserRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries to return on a single page.
	//
	// - The default value is 20.
	//
	// - The maximum value is 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that specifies the start of the next page for a paged query.
	//
	// - If you do not specify this parameter, the query starts from the first page.
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListAuthorizationRulesForUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationRulesForUserRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizationRulesForUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAuthorizationRulesForUserRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAuthorizationRulesForUserRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAuthorizationRulesForUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListAuthorizationRulesForUserRequest) SetInstanceId(v string) *ListAuthorizationRulesForUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAuthorizationRulesForUserRequest) SetMaxResults(v int32) *ListAuthorizationRulesForUserRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAuthorizationRulesForUserRequest) SetNextToken(v string) *ListAuthorizationRulesForUserRequest {
	s.NextToken = &v
	return s
}

func (s *ListAuthorizationRulesForUserRequest) SetUserId(v string) *ListAuthorizationRulesForUserRequest {
	s.UserId = &v
	return s
}

func (s *ListAuthorizationRulesForUserRequest) Validate() error {
	return dara.Validate(s)
}
