// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConsumerAuthorizationRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiNameLike(v string) *QueryConsumerAuthorizationRulesRequest
	GetApiNameLike() *string
	SetConsumerId(v string) *QueryConsumerAuthorizationRulesRequest
	GetConsumerId() *string
	SetConsumerNameLike(v string) *QueryConsumerAuthorizationRulesRequest
	GetConsumerNameLike() *string
	SetEnvironmentId(v string) *QueryConsumerAuthorizationRulesRequest
	GetEnvironmentId() *string
	SetGroupByApi(v bool) *QueryConsumerAuthorizationRulesRequest
	GetGroupByApi() *bool
	SetPageNumber(v int32) *QueryConsumerAuthorizationRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryConsumerAuthorizationRulesRequest
	GetPageSize() *int32
	SetParentResourceId(v string) *QueryConsumerAuthorizationRulesRequest
	GetParentResourceId() *string
	SetResourceId(v string) *QueryConsumerAuthorizationRulesRequest
	GetResourceId() *string
	SetResourceType(v string) *QueryConsumerAuthorizationRulesRequest
	GetResourceType() *string
	SetResourceTypes(v string) *QueryConsumerAuthorizationRulesRequest
	GetResourceTypes() *string
}

type QueryConsumerAuthorizationRulesRequest struct {
	// The API name.
	//
	// example:
	//
	// api-xx
	ApiNameLike *string `json:"apiNameLike,omitempty" xml:"apiNameLike,omitempty"`
	// The consumer ID.
	//
	// example:
	//
	// cs-ct21c16m1hkp64hk6qmg
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// The consumer name.
	//
	// example:
	//
	// consumer-xxx
	ConsumerNameLike *string `json:"consumerNameLike,omitempty" xml:"consumerNameLike,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// env-cpqnr6tlhtgubc***
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// Specifies whether to group the results by API.
	//
	// example:
	//
	// true
	GroupByApi *bool `json:"groupByApi,omitempty" xml:"groupByApi,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The parent resource ID.
	//
	// example:
	//
	// ha-cn-li942gy8p01
	ParentResourceId *string `json:"parentResourceId,omitempty" xml:"parentResourceId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// ha-cn-li942gy8p03
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// HttpApiRoute
	ResourceType  *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	ResourceTypes *string `json:"resourceTypes,omitempty" xml:"resourceTypes,omitempty"`
}

func (s QueryConsumerAuthorizationRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryConsumerAuthorizationRulesRequest) GoString() string {
	return s.String()
}

func (s *QueryConsumerAuthorizationRulesRequest) GetApiNameLike() *string {
	return s.ApiNameLike
}

func (s *QueryConsumerAuthorizationRulesRequest) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *QueryConsumerAuthorizationRulesRequest) GetConsumerNameLike() *string {
	return s.ConsumerNameLike
}

func (s *QueryConsumerAuthorizationRulesRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *QueryConsumerAuthorizationRulesRequest) GetGroupByApi() *bool {
	return s.GroupByApi
}

func (s *QueryConsumerAuthorizationRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryConsumerAuthorizationRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryConsumerAuthorizationRulesRequest) GetParentResourceId() *string {
	return s.ParentResourceId
}

func (s *QueryConsumerAuthorizationRulesRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *QueryConsumerAuthorizationRulesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *QueryConsumerAuthorizationRulesRequest) GetResourceTypes() *string {
	return s.ResourceTypes
}

func (s *QueryConsumerAuthorizationRulesRequest) SetApiNameLike(v string) *QueryConsumerAuthorizationRulesRequest {
	s.ApiNameLike = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesRequest) SetConsumerId(v string) *QueryConsumerAuthorizationRulesRequest {
	s.ConsumerId = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesRequest) SetConsumerNameLike(v string) *QueryConsumerAuthorizationRulesRequest {
	s.ConsumerNameLike = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesRequest) SetEnvironmentId(v string) *QueryConsumerAuthorizationRulesRequest {
	s.EnvironmentId = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesRequest) SetGroupByApi(v bool) *QueryConsumerAuthorizationRulesRequest {
	s.GroupByApi = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesRequest) SetPageNumber(v int32) *QueryConsumerAuthorizationRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesRequest) SetPageSize(v int32) *QueryConsumerAuthorizationRulesRequest {
	s.PageSize = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesRequest) SetParentResourceId(v string) *QueryConsumerAuthorizationRulesRequest {
	s.ParentResourceId = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesRequest) SetResourceId(v string) *QueryConsumerAuthorizationRulesRequest {
	s.ResourceId = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesRequest) SetResourceType(v string) *QueryConsumerAuthorizationRulesRequest {
	s.ResourceType = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesRequest) SetResourceTypes(v string) *QueryConsumerAuthorizationRulesRequest {
	s.ResourceTypes = &v
	return s
}

func (s *QueryConsumerAuthorizationRulesRequest) Validate() error {
	return dara.Validate(s)
}
