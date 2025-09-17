// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *GetResourcesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *GetResourcesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *GetResourcesResponseBody
	GetRequestId() *string
	SetResource(v *GetResourcesResponseBodyResource) *GetResourcesResponseBody
	GetResource() *GetResourcesResponseBodyResource
	SetResources(v []*GetResourcesResponseBodyResources) *GetResourcesResponseBody
	GetResources() []*GetResourcesResponseBodyResources
	SetTotalCount(v int32) *GetResourcesResponseBody
	GetTotalCount() *int32
}

type GetResourcesResponseBody struct {
	// The maximum number of entries returned. Return result of the List operation.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists. Return result of the List operation.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The specified resource. Return result of the Get operation.
	Resource *GetResourcesResponseBodyResource `json:"resource,omitempty" xml:"resource,omitempty" type:"Struct"`
	// The resource list. Return result of the List operation.
	Resources []*GetResourcesResponseBodyResources `json:"resources,omitempty" xml:"resources,omitempty" type:"Repeated"`
	// The total number of entries returned. Return result of the List operation.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourcesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourcesResponseBody) GetResource() *GetResourcesResponseBodyResource {
	return s.Resource
}

func (s *GetResourcesResponseBody) GetResources() []*GetResourcesResponseBodyResources {
	return s.Resources
}

func (s *GetResourcesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetResourcesResponseBody) SetMaxResults(v int32) *GetResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetResourcesResponseBody) SetNextToken(v string) *GetResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetResourcesResponseBody) SetRequestId(v string) *GetResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourcesResponseBody) SetResource(v *GetResourcesResponseBodyResource) *GetResourcesResponseBody {
	s.Resource = v
	return s
}

func (s *GetResourcesResponseBody) SetResources(v []*GetResourcesResponseBodyResources) *GetResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *GetResourcesResponseBody) SetTotalCount(v int32) *GetResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetResourcesResponseBodyResource struct {
	// The resource properties in the JSON format.
	//
	// example:
	//
	// {"Status":"Available","Description":"","AccountPrivilege":"RoleReadWrite","InstanceId":"r-8vbf5abe31c9c4d4","RegionId":"cn-zhangjiakou","AccountType":"Normal","TypeInfo":{},"AccountName":"cctest"}
	ResourceAttributes map[string]interface{} `json:"resourceAttributes,omitempty" xml:"resourceAttributes,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// cctest
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
}

func (s GetResourcesResponseBodyResource) String() string {
	return dara.Prettify(s)
}

func (s GetResourcesResponseBodyResource) GoString() string {
	return s.String()
}

func (s *GetResourcesResponseBodyResource) GetResourceAttributes() map[string]interface{} {
	return s.ResourceAttributes
}

func (s *GetResourcesResponseBodyResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetResourcesResponseBodyResource) SetResourceAttributes(v map[string]interface{}) *GetResourcesResponseBodyResource {
	s.ResourceAttributes = v
	return s
}

func (s *GetResourcesResponseBodyResource) SetResourceId(v string) *GetResourcesResponseBodyResource {
	s.ResourceId = &v
	return s
}

func (s *GetResourcesResponseBodyResource) Validate() error {
	return dara.Validate(s)
}

type GetResourcesResponseBodyResources struct {
	// The resource properties in the JSON format.
	//
	// example:
	//
	// {"Status":"Available","Description":"","AccountPrivilege":"RoleReadWrite","InstanceId":"r-8vbf5abe31c9c4d4","RegionId":"cn-zhangjiakou","AccountType":"Normal","TypeInfo":{},"AccountName":"cctest"}
	ResourceAttributes map[string]interface{} `json:"resourceAttributes,omitempty" xml:"resourceAttributes,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// cctest
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
}

func (s GetResourcesResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s GetResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *GetResourcesResponseBodyResources) GetResourceAttributes() map[string]interface{} {
	return s.ResourceAttributes
}

func (s *GetResourcesResponseBodyResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetResourcesResponseBodyResources) SetResourceAttributes(v map[string]interface{}) *GetResourcesResponseBodyResources {
	s.ResourceAttributes = v
	return s
}

func (s *GetResourcesResponseBodyResources) SetResourceId(v string) *GetResourcesResponseBodyResources {
	s.ResourceId = &v
	return s
}

func (s *GetResourcesResponseBodyResources) Validate() error {
	return dara.Validate(s)
}
