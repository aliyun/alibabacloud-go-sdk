// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAccountResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAccountId(v string) *ListCloudAccountResourcesRequest
	GetCloudAccountId() *string
	SetCloudAccountResourceEntityType(v string) *ListCloudAccountResourcesRequest
	GetCloudAccountResourceEntityType() *string
	SetCloudAccountResourceNameContains(v string) *ListCloudAccountResourcesRequest
	GetCloudAccountResourceNameContains() *string
	SetCloudAccountResourceNamePrefix(v string) *ListCloudAccountResourcesRequest
	GetCloudAccountResourceNamePrefix() *string
	SetInstanceId(v string) *ListCloudAccountResourcesRequest
	GetInstanceId() *string
	SetMaxResults(v int64) *ListCloudAccountResourcesRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListCloudAccountResourcesRequest
	GetNextToken() *string
}

type ListCloudAccountResourcesRequest struct {
	// The ID of the cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789012xxxx
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// Filters resources by entity type. Valid values: ram_policy.
	//
	// example:
	//
	// ram_policy
	CloudAccountResourceEntityType *string `json:"CloudAccountResourceEntityType,omitempty" xml:"CloudAccountResourceEntityType,omitempty"`
	// Performs a fuzzy match on the resource name, which is equivalent to LIKE \\"%value%\\".
	//
	// example:
	//
	// sso
	CloudAccountResourceNameContains *string `json:"CloudAccountResourceNameContains,omitempty" xml:"CloudAccountResourceNameContains,omitempty"`
	// Matches resources by resource name prefix.
	//
	// example:
	//
	// eiam-
	CloudAccountResourceNamePrefix *string `json:"CloudAccountResourceNamePrefix,omitempty" xml:"CloudAccountResourceNamePrefix,omitempty"`
	// The ID of the IDaaS EIAM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries to return per page for paging in a paged query. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token. The value is the NextToken parameter value returned by the previous API call.
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListCloudAccountResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccountResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListCloudAccountResourcesRequest) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *ListCloudAccountResourcesRequest) GetCloudAccountResourceEntityType() *string {
	return s.CloudAccountResourceEntityType
}

func (s *ListCloudAccountResourcesRequest) GetCloudAccountResourceNameContains() *string {
	return s.CloudAccountResourceNameContains
}

func (s *ListCloudAccountResourcesRequest) GetCloudAccountResourceNamePrefix() *string {
	return s.CloudAccountResourceNamePrefix
}

func (s *ListCloudAccountResourcesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCloudAccountResourcesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListCloudAccountResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCloudAccountResourcesRequest) SetCloudAccountId(v string) *ListCloudAccountResourcesRequest {
	s.CloudAccountId = &v
	return s
}

func (s *ListCloudAccountResourcesRequest) SetCloudAccountResourceEntityType(v string) *ListCloudAccountResourcesRequest {
	s.CloudAccountResourceEntityType = &v
	return s
}

func (s *ListCloudAccountResourcesRequest) SetCloudAccountResourceNameContains(v string) *ListCloudAccountResourcesRequest {
	s.CloudAccountResourceNameContains = &v
	return s
}

func (s *ListCloudAccountResourcesRequest) SetCloudAccountResourceNamePrefix(v string) *ListCloudAccountResourcesRequest {
	s.CloudAccountResourceNamePrefix = &v
	return s
}

func (s *ListCloudAccountResourcesRequest) SetInstanceId(v string) *ListCloudAccountResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCloudAccountResourcesRequest) SetMaxResults(v int64) *ListCloudAccountResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCloudAccountResourcesRequest) SetNextToken(v string) *ListCloudAccountResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListCloudAccountResourcesRequest) Validate() error {
	return dara.Validate(s)
}
