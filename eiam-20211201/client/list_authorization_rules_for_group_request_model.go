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
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of records to return on each page.
	//
	// - If this parameter is not specified, the default value is 20.
	//
	// - The maximum value is 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that specifies the start of the next page for a paged query.
	//
	// - If this parameter is not specified, the query starts from the first page.
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
