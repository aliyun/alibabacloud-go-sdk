// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceServerScopesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListResourceServerScopesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListResourceServerScopesResponseBody
	GetNextToken() *string
	SetPreviousToken(v string) *ListResourceServerScopesResponseBody
	GetPreviousToken() *string
	SetRequestId(v string) *ListResourceServerScopesResponseBody
	GetRequestId() *string
	SetResourceServerScopes(v []*ListResourceServerScopesResponseBodyResourceServerScopes) *ListResourceServerScopesResponseBody
	GetResourceServerScopes() []*ListResourceServerScopesResponseBodyResourceServerScopes
	SetTotalCount(v int32) *ListResourceServerScopesResponseBody
	GetTotalCount() *int32
}

type ListResourceServerScopesResponseBody struct {
	// The number of entries per page in a paged query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The pagination token for the previous page.
	//
	// example:
	//
	// PTxxxexample
	PreviousToken *string `json:"PreviousToken,omitempty" xml:"PreviousToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of scope permissions under the ResourceServer.
	ResourceServerScopes []*ListResourceServerScopesResponseBodyResourceServerScopes `json:"ResourceServerScopes,omitempty" xml:"ResourceServerScopes,omitempty" type:"Repeated"`
	// The total number of entries in the list.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceServerScopesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceServerScopesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceServerScopesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceServerScopesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceServerScopesResponseBody) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListResourceServerScopesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceServerScopesResponseBody) GetResourceServerScopes() []*ListResourceServerScopesResponseBodyResourceServerScopes {
	return s.ResourceServerScopes
}

func (s *ListResourceServerScopesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListResourceServerScopesResponseBody) SetMaxResults(v int32) *ListResourceServerScopesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListResourceServerScopesResponseBody) SetNextToken(v string) *ListResourceServerScopesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResourceServerScopesResponseBody) SetPreviousToken(v string) *ListResourceServerScopesResponseBody {
	s.PreviousToken = &v
	return s
}

func (s *ListResourceServerScopesResponseBody) SetRequestId(v string) *ListResourceServerScopesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceServerScopesResponseBody) SetResourceServerScopes(v []*ListResourceServerScopesResponseBodyResourceServerScopes) *ListResourceServerScopesResponseBody {
	s.ResourceServerScopes = v
	return s
}

func (s *ListResourceServerScopesResponseBody) SetTotalCount(v int32) *ListResourceServerScopesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListResourceServerScopesResponseBody) Validate() error {
	if s.ResourceServerScopes != nil {
		for _, item := range s.ResourceServerScopes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceServerScopesResponseBodyResourceServerScopes struct {
	// The application ID.
	//
	// example:
	//
	// app_xxxxxxxxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The authorization type.
	//
	// example:
	//
	// authorize_required
	AuthorizationType *string `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_xxxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The scope permission ID.
	//
	// example:
	//
	// rss_xxxxxxxxxxx
	ResourceServerScopeId *string `json:"ResourceServerScopeId,omitempty" xml:"ResourceServerScopeId,omitempty"`
	// The scope permission name.
	//
	// example:
	//
	// 读取全部用户
	ResourceServerScopeName *string `json:"ResourceServerScopeName,omitempty" xml:"ResourceServerScopeName,omitempty"`
	// The scope permission type.
	//
	// example:
	//
	// urn:alibaba:idaas:resourceserver:scope:delegated
	ResourceServerScopeType *string `json:"ResourceServerScopeType,omitempty" xml:"ResourceServerScopeType,omitempty"`
	// The scope permission value.
	//
	// example:
	//
	// User:Write:ALL
	ResourceServerScopeValue *string `json:"ResourceServerScopeValue,omitempty" xml:"ResourceServerScopeValue,omitempty"`
}

func (s ListResourceServerScopesResponseBodyResourceServerScopes) String() string {
	return dara.Prettify(s)
}

func (s ListResourceServerScopesResponseBodyResourceServerScopes) GoString() string {
	return s.String()
}

func (s *ListResourceServerScopesResponseBodyResourceServerScopes) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListResourceServerScopesResponseBodyResourceServerScopes) GetAuthorizationType() *string {
	return s.AuthorizationType
}

func (s *ListResourceServerScopesResponseBodyResourceServerScopes) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListResourceServerScopesResponseBodyResourceServerScopes) GetResourceServerScopeId() *string {
	return s.ResourceServerScopeId
}

func (s *ListResourceServerScopesResponseBodyResourceServerScopes) GetResourceServerScopeName() *string {
	return s.ResourceServerScopeName
}

func (s *ListResourceServerScopesResponseBodyResourceServerScopes) GetResourceServerScopeType() *string {
	return s.ResourceServerScopeType
}

func (s *ListResourceServerScopesResponseBodyResourceServerScopes) GetResourceServerScopeValue() *string {
	return s.ResourceServerScopeValue
}

func (s *ListResourceServerScopesResponseBodyResourceServerScopes) SetApplicationId(v string) *ListResourceServerScopesResponseBodyResourceServerScopes {
	s.ApplicationId = &v
	return s
}

func (s *ListResourceServerScopesResponseBodyResourceServerScopes) SetAuthorizationType(v string) *ListResourceServerScopesResponseBodyResourceServerScopes {
	s.AuthorizationType = &v
	return s
}

func (s *ListResourceServerScopesResponseBodyResourceServerScopes) SetInstanceId(v string) *ListResourceServerScopesResponseBodyResourceServerScopes {
	s.InstanceId = &v
	return s
}

func (s *ListResourceServerScopesResponseBodyResourceServerScopes) SetResourceServerScopeId(v string) *ListResourceServerScopesResponseBodyResourceServerScopes {
	s.ResourceServerScopeId = &v
	return s
}

func (s *ListResourceServerScopesResponseBodyResourceServerScopes) SetResourceServerScopeName(v string) *ListResourceServerScopesResponseBodyResourceServerScopes {
	s.ResourceServerScopeName = &v
	return s
}

func (s *ListResourceServerScopesResponseBodyResourceServerScopes) SetResourceServerScopeType(v string) *ListResourceServerScopesResponseBodyResourceServerScopes {
	s.ResourceServerScopeType = &v
	return s
}

func (s *ListResourceServerScopesResponseBodyResourceServerScopes) SetResourceServerScopeValue(v string) *ListResourceServerScopesResponseBodyResourceServerScopes {
	s.ResourceServerScopeValue = &v
	return s
}

func (s *ListResourceServerScopesResponseBodyResourceServerScopes) Validate() error {
	return dara.Validate(s)
}
