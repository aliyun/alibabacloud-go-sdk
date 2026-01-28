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
	// 应用 ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 分页查询时每页行数。默认值为20，最大值为100。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 查询凭证（Token），取值为上一次API调用返回的NextToken参数值。
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
