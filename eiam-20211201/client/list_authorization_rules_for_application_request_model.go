// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizationRulesForApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListAuthorizationRulesForApplicationRequest
	GetApplicationId() *string
	SetInstanceId(v string) *ListAuthorizationRulesForApplicationRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListAuthorizationRulesForApplicationRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAuthorizationRulesForApplicationRequest
	GetNextToken() *string
}

type ListAuthorizationRulesForApplicationRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries per page.
	//
	// - Default value: 20.
	//
	// - Maximum value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the starting position of the next page.
	//
	// - If you do not specify this parameter, the query starts from the first page.
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListAuthorizationRulesForApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationRulesForApplicationRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizationRulesForApplicationRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListAuthorizationRulesForApplicationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAuthorizationRulesForApplicationRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAuthorizationRulesForApplicationRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAuthorizationRulesForApplicationRequest) SetApplicationId(v string) *ListAuthorizationRulesForApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListAuthorizationRulesForApplicationRequest) SetInstanceId(v string) *ListAuthorizationRulesForApplicationRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAuthorizationRulesForApplicationRequest) SetMaxResults(v int32) *ListAuthorizationRulesForApplicationRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAuthorizationRulesForApplicationRequest) SetNextToken(v string) *ListAuthorizationRulesForApplicationRequest {
	s.NextToken = &v
	return s
}

func (s *ListAuthorizationRulesForApplicationRequest) Validate() error {
	return dara.Validate(s)
}
