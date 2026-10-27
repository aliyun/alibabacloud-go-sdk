// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentityProviderDepartmentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListIdentityProviderDepartmentsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListIdentityProviderDepartmentsResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListIdentityProviderDepartmentsResponseBodyItems) *ListIdentityProviderDepartmentsResponseBody
	GetItems() []*ListIdentityProviderDepartmentsResponseBodyItems
	SetMaxResults(v int32) *ListIdentityProviderDepartmentsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListIdentityProviderDepartmentsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListIdentityProviderDepartmentsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListIdentityProviderDepartmentsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListIdentityProviderDepartmentsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListIdentityProviderDepartmentsResponseBody
	GetTotalCount() *int64
}

type ListIdentityProviderDepartmentsResponseBody struct {
	// The business status code of the API request. SUCCESS indicates that the request is successful, but does not indicate that the asynchronous synchronization task is completed.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code of the API request, which does not indicate the execution result of the asynchronous task.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The direct child departments on the current page under the current parent department. You must make another query to obtain the details of the child departments.
	//
	// example:
	//
	// [{"id":"10001","name":"R&D Department","parentId":"1","memberCount":12,"hasChildren":true}]
	Items []*ListIdentityProviderDepartmentsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The number of items per page used for this query.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message of the API request.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The token for the next page. An empty or missing value indicates that there are no subsequent pages at the current hierarchy level of the parent department, but does not indicate that the entire department tree has been returned.
	//
	// example:
	//
	// example-next-page-token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The unique identifier of the request. Provide this value when troubleshooting issues.
	//
	// example:
	//
	// 01234567-89AB-CDEF-0123-456789ABCDEF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the API request is successful. To check whether an asynchronous task is completed, query the task status.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of departments. This operation does not provide this statistic and may not return this field. Do not treat a missing value as 0.
	//
	// example:
	//
	// null
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListIdentityProviderDepartmentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProviderDepartmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIdentityProviderDepartmentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListIdentityProviderDepartmentsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListIdentityProviderDepartmentsResponseBody) GetItems() []*ListIdentityProviderDepartmentsResponseBodyItems {
	return s.Items
}

func (s *ListIdentityProviderDepartmentsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIdentityProviderDepartmentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListIdentityProviderDepartmentsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIdentityProviderDepartmentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIdentityProviderDepartmentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListIdentityProviderDepartmentsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListIdentityProviderDepartmentsResponseBody) SetCode(v string) *ListIdentityProviderDepartmentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListIdentityProviderDepartmentsResponseBody) SetHttpStatusCode(v int32) *ListIdentityProviderDepartmentsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListIdentityProviderDepartmentsResponseBody) SetItems(v []*ListIdentityProviderDepartmentsResponseBodyItems) *ListIdentityProviderDepartmentsResponseBody {
	s.Items = v
	return s
}

func (s *ListIdentityProviderDepartmentsResponseBody) SetMaxResults(v int32) *ListIdentityProviderDepartmentsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIdentityProviderDepartmentsResponseBody) SetMessage(v string) *ListIdentityProviderDepartmentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListIdentityProviderDepartmentsResponseBody) SetNextToken(v string) *ListIdentityProviderDepartmentsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIdentityProviderDepartmentsResponseBody) SetRequestId(v string) *ListIdentityProviderDepartmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIdentityProviderDepartmentsResponseBody) SetSuccess(v bool) *ListIdentityProviderDepartmentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListIdentityProviderDepartmentsResponseBody) SetTotalCount(v int64) *ListIdentityProviderDepartmentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIdentityProviderDepartmentsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIdentityProviderDepartmentsResponseBodyItems struct {
	// Indicates whether the department has child departments. A value of true indicates that the department has child departments and an expansion entry can be displayed accordingly. A value of false indicates a leaf node. If this field is not returned, the status is unknown. You must use the ID of the current node to make another query to obtain the details of the child departments.
	//
	// example:
	//
	// true
	HasChildren *bool `json:"hasChildren,omitempty" xml:"hasChildren,omitempty"`
	// The unique identifier of the department returned by the identity provider. When querying the child departments of this node, use this value as the parentDepartmentId.
	//
	// example:
	//
	// 10001
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The number of members in the department returned by the identity provider. This field is not returned when the number is unknown. Do not treat a missing value as 0.
	//
	// example:
	//
	// 12
	MemberCount *int64 `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	// The name of the department returned by the identity provider.
	//
	// example:
	//
	// R&D Department
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the parent department to which this department belongs. This value describes the parent-child relationship and is not used to expand the current node.
	//
	// example:
	//
	// 1
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
}

func (s ListIdentityProviderDepartmentsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProviderDepartmentsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListIdentityProviderDepartmentsResponseBodyItems) GetHasChildren() *bool {
	return s.HasChildren
}

func (s *ListIdentityProviderDepartmentsResponseBodyItems) GetId() *string {
	return s.Id
}

func (s *ListIdentityProviderDepartmentsResponseBodyItems) GetMemberCount() *int64 {
	return s.MemberCount
}

func (s *ListIdentityProviderDepartmentsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListIdentityProviderDepartmentsResponseBodyItems) GetParentId() *string {
	return s.ParentId
}

func (s *ListIdentityProviderDepartmentsResponseBodyItems) SetHasChildren(v bool) *ListIdentityProviderDepartmentsResponseBodyItems {
	s.HasChildren = &v
	return s
}

func (s *ListIdentityProviderDepartmentsResponseBodyItems) SetId(v string) *ListIdentityProviderDepartmentsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *ListIdentityProviderDepartmentsResponseBodyItems) SetMemberCount(v int64) *ListIdentityProviderDepartmentsResponseBodyItems {
	s.MemberCount = &v
	return s
}

func (s *ListIdentityProviderDepartmentsResponseBodyItems) SetName(v string) *ListIdentityProviderDepartmentsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListIdentityProviderDepartmentsResponseBodyItems) SetParentId(v string) *ListIdentityProviderDepartmentsResponseBodyItems {
	s.ParentId = &v
	return s
}

func (s *ListIdentityProviderDepartmentsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
