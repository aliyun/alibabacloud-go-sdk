// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizationRulesForGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ListAuthorizationRulesForGroupRequest
	GetGroupId() *string
	SetInstanceId(v string) *ListAuthorizationRulesForGroupRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListAuthorizationRulesForGroupRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAuthorizationRulesForGroupRequest
	GetNextToken() *string
}

type ListAuthorizationRulesForGroupRequest struct {
	// 组标识。
	//
	// This parameter is required.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
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

func (s ListAuthorizationRulesForGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationRulesForGroupRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizationRulesForGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListAuthorizationRulesForGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAuthorizationRulesForGroupRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAuthorizationRulesForGroupRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAuthorizationRulesForGroupRequest) SetGroupId(v string) *ListAuthorizationRulesForGroupRequest {
	s.GroupId = &v
	return s
}

func (s *ListAuthorizationRulesForGroupRequest) SetInstanceId(v string) *ListAuthorizationRulesForGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAuthorizationRulesForGroupRequest) SetMaxResults(v int32) *ListAuthorizationRulesForGroupRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAuthorizationRulesForGroupRequest) SetNextToken(v string) *ListAuthorizationRulesForGroupRequest {
	s.NextToken = &v
	return s
}

func (s *ListAuthorizationRulesForGroupRequest) Validate() error {
	return dara.Validate(s)
}
