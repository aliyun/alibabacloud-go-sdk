// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AssociateDefaultFilterRequest struct {
	// The name of the filter.
	//
	// This parameter is required.
	FilterName *string `json:"FilterName,omitempty" xml:"FilterName,omitempty"`
}

func (s AssociateDefaultFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateDefaultFilterRequest) GoString() string {
	return s.String()
}

func (s *AssociateDefaultFilterRequest) SetFilterName(v string) *AssociateDefaultFilterRequest {
	s.FilterName = &v
	return s
}

type AssociateDefaultFilterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54673B22-2001-556A-B394-B8697AA9670B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateDefaultFilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateDefaultFilterResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateDefaultFilterResponseBody) SetRequestId(v string) *AssociateDefaultFilterResponseBody {
	s.RequestId = &v
	return s
}

type AssociateDefaultFilterResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateDefaultFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateDefaultFilterResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateDefaultFilterResponse) GoString() string {
	return s.String()
}

func (s *AssociateDefaultFilterResponse) SetHeaders(v map[string]*string) *AssociateDefaultFilterResponse {
	s.Headers = v
	return s
}

func (s *AssociateDefaultFilterResponse) SetStatusCode(v int32) *AssociateDefaultFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateDefaultFilterResponse) SetBody(v *AssociateDefaultFilterResponseBody) *AssociateDefaultFilterResponse {
	s.Body = v
	return s
}

type CreateFilterRequest struct {
	// The configurations of the filter.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "regions": [],
	//
	//   "tagFilters": [
	//
	//     [{ "type": "notContainTagKey", "tagKey": "xxx", "tagValue": "" }],
	//
	//     [{ "tagKey": "xxx", "tagValue": "xxx" }]
	//
	//   ],
	//
	//   "resourceTypes": [
	//
	//     "ACS::ECS::AutoSnapshotPolicy"
	//
	//   ]
	//
	// }
	FilterConfiguration *string `json:"FilterConfiguration,omitempty" xml:"FilterConfiguration,omitempty"`
	// The name of the filter.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS
	FilterName *string `json:"FilterName,omitempty" xml:"FilterName,omitempty"`
}

func (s CreateFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFilterRequest) GoString() string {
	return s.String()
}

func (s *CreateFilterRequest) SetFilterConfiguration(v string) *CreateFilterRequest {
	s.FilterConfiguration = &v
	return s
}

func (s *CreateFilterRequest) SetFilterName(v string) *CreateFilterRequest {
	s.FilterName = &v
	return s
}

type CreateFilterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EEF1EE1F-50F6-5494-B3DA-8F597DEB31BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFilterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFilterResponseBody) SetRequestId(v string) *CreateFilterResponseBody {
	s.RequestId = &v
	return s
}

type CreateFilterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFilterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFilterResponse) GoString() string {
	return s.String()
}

func (s *CreateFilterResponse) SetHeaders(v map[string]*string) *CreateFilterResponse {
	s.Headers = v
	return s
}

func (s *CreateFilterResponse) SetStatusCode(v int32) *CreateFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFilterResponse) SetBody(v *CreateFilterResponseBody) *CreateFilterResponse {
	s.Body = v
	return s
}

type CreateSavedQueryRequest struct {
	// The description of the template.
	//
	// The description must be 1 to 256 characters in length.
	//
	// example:
	//
	// Queries all resources on which you have permissions and sorts the resources by resource type and resource ID.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The query statement in the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// SELECT 	- FROM resources;
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The name of the template.
	//
	// 	- The name must be 1 to 64 characters in length.
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// Query of All Alibaba Cloud Resources
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateSavedQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSavedQueryRequest) GoString() string {
	return s.String()
}

func (s *CreateSavedQueryRequest) SetDescription(v string) *CreateSavedQueryRequest {
	s.Description = &v
	return s
}

func (s *CreateSavedQueryRequest) SetExpression(v string) *CreateSavedQueryRequest {
	s.Expression = &v
	return s
}

func (s *CreateSavedQueryRequest) SetName(v string) *CreateSavedQueryRequest {
	s.Name = &v
	return s
}

type CreateSavedQueryResponseBody struct {
	// The template ID.
	//
	// example:
	//
	// sq-GeAck****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EFA806B9-7F36-55AB-8B7A-D680C2C5EE57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSavedQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSavedQueryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSavedQueryResponseBody) SetQueryId(v string) *CreateSavedQueryResponseBody {
	s.QueryId = &v
	return s
}

func (s *CreateSavedQueryResponseBody) SetRequestId(v string) *CreateSavedQueryResponseBody {
	s.RequestId = &v
	return s
}

type CreateSavedQueryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSavedQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSavedQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSavedQueryResponse) GoString() string {
	return s.String()
}

func (s *CreateSavedQueryResponse) SetHeaders(v map[string]*string) *CreateSavedQueryResponse {
	s.Headers = v
	return s
}

func (s *CreateSavedQueryResponse) SetStatusCode(v int32) *CreateSavedQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSavedQueryResponse) SetBody(v *CreateSavedQueryResponseBody) *CreateSavedQueryResponse {
	s.Body = v
	return s
}

type DeleteFilterRequest struct {
	// The name of the filter.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS
	FilterName *string `json:"FilterName,omitempty" xml:"FilterName,omitempty"`
}

func (s DeleteFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFilterRequest) GoString() string {
	return s.String()
}

func (s *DeleteFilterRequest) SetFilterName(v string) *DeleteFilterRequest {
	s.FilterName = &v
	return s
}

type DeleteFilterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A4A63E3C-89EC-51F9-9934-C9AF1BCBAAA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFilterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFilterResponseBody) SetRequestId(v string) *DeleteFilterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFilterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFilterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFilterResponse) GoString() string {
	return s.String()
}

func (s *DeleteFilterResponse) SetHeaders(v map[string]*string) *DeleteFilterResponse {
	s.Headers = v
	return s
}

func (s *DeleteFilterResponse) SetStatusCode(v int32) *DeleteFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFilterResponse) SetBody(v *DeleteFilterResponseBody) *DeleteFilterResponse {
	s.Body = v
	return s
}

type DeleteSavedQueryRequest struct {
	// The ID of the template.
	//
	// You can call the [ListSavedQueries](~~ListSavedQueries~~) operation to obtain the template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sq-GeAck****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s DeleteSavedQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSavedQueryRequest) GoString() string {
	return s.String()
}

func (s *DeleteSavedQueryRequest) SetQueryId(v string) *DeleteSavedQueryRequest {
	s.QueryId = &v
	return s
}

type DeleteSavedQueryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D170D58E-6256-5344-8F5E-922EC9ECB7EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSavedQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSavedQueryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSavedQueryResponseBody) SetRequestId(v string) *DeleteSavedQueryResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSavedQueryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSavedQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSavedQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSavedQueryResponse) GoString() string {
	return s.String()
}

func (s *DeleteSavedQueryResponse) SetHeaders(v map[string]*string) *DeleteSavedQueryResponse {
	s.Headers = v
	return s
}

func (s *DeleteSavedQueryResponse) SetStatusCode(v int32) *DeleteSavedQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSavedQueryResponse) SetBody(v *DeleteSavedQueryResponseBody) *DeleteSavedQueryResponse {
	s.Body = v
	return s
}

type DisableMultiAccountResourceCenterResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4951F920-48DB-5731-96AA-3A7C8AE617D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableMultiAccountResourceCenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableMultiAccountResourceCenterResponseBody) GoString() string {
	return s.String()
}

func (s *DisableMultiAccountResourceCenterResponseBody) SetRequestId(v string) *DisableMultiAccountResourceCenterResponseBody {
	s.RequestId = &v
	return s
}

type DisableMultiAccountResourceCenterResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableMultiAccountResourceCenterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableMultiAccountResourceCenterResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableMultiAccountResourceCenterResponse) GoString() string {
	return s.String()
}

func (s *DisableMultiAccountResourceCenterResponse) SetHeaders(v map[string]*string) *DisableMultiAccountResourceCenterResponse {
	s.Headers = v
	return s
}

func (s *DisableMultiAccountResourceCenterResponse) SetStatusCode(v int32) *DisableMultiAccountResourceCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableMultiAccountResourceCenterResponse) SetBody(v *DisableMultiAccountResourceCenterResponseBody) *DisableMultiAccountResourceCenterResponse {
	s.Body = v
	return s
}

type DisableResourceCenterResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D170D58E-6256-5344-8F5E-922EC9ECB7EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableResourceCenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableResourceCenterResponseBody) GoString() string {
	return s.String()
}

func (s *DisableResourceCenterResponseBody) SetRequestId(v string) *DisableResourceCenterResponseBody {
	s.RequestId = &v
	return s
}

type DisableResourceCenterResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableResourceCenterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableResourceCenterResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableResourceCenterResponse) GoString() string {
	return s.String()
}

func (s *DisableResourceCenterResponse) SetHeaders(v map[string]*string) *DisableResourceCenterResponse {
	s.Headers = v
	return s
}

func (s *DisableResourceCenterResponse) SetStatusCode(v int32) *DisableResourceCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableResourceCenterResponse) SetBody(v *DisableResourceCenterResponseBody) *DisableResourceCenterResponse {
	s.Body = v
	return s
}

type DisassociateDefaultFilterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BCAB07BA-82FA-5DC0-9322-FB7ED726481D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisassociateDefaultFilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisassociateDefaultFilterResponseBody) GoString() string {
	return s.String()
}

func (s *DisassociateDefaultFilterResponseBody) SetRequestId(v string) *DisassociateDefaultFilterResponseBody {
	s.RequestId = &v
	return s
}

type DisassociateDefaultFilterResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisassociateDefaultFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisassociateDefaultFilterResponse) String() string {
	return tea.Prettify(s)
}

func (s DisassociateDefaultFilterResponse) GoString() string {
	return s.String()
}

func (s *DisassociateDefaultFilterResponse) SetHeaders(v map[string]*string) *DisassociateDefaultFilterResponse {
	s.Headers = v
	return s
}

func (s *DisassociateDefaultFilterResponse) SetStatusCode(v int32) *DisassociateDefaultFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *DisassociateDefaultFilterResponse) SetBody(v *DisassociateDefaultFilterResponseBody) *DisassociateDefaultFilterResponse {
	s.Body = v
	return s
}

type EnableMultiAccountResourceCenterResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 767038B7-2027-5508-858B-E213232D57D5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the feature. Valid values:
	//
	// 	- Pending: The feature is being enabled.
	//
	// 	- Enabled: The feature is enabled.
	//
	// example:
	//
	// Pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s EnableMultiAccountResourceCenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableMultiAccountResourceCenterResponseBody) GoString() string {
	return s.String()
}

func (s *EnableMultiAccountResourceCenterResponseBody) SetRequestId(v string) *EnableMultiAccountResourceCenterResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableMultiAccountResourceCenterResponseBody) SetStatus(v string) *EnableMultiAccountResourceCenterResponseBody {
	s.Status = &v
	return s
}

type EnableMultiAccountResourceCenterResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableMultiAccountResourceCenterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableMultiAccountResourceCenterResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableMultiAccountResourceCenterResponse) GoString() string {
	return s.String()
}

func (s *EnableMultiAccountResourceCenterResponse) SetHeaders(v map[string]*string) *EnableMultiAccountResourceCenterResponse {
	s.Headers = v
	return s
}

func (s *EnableMultiAccountResourceCenterResponse) SetStatusCode(v int32) *EnableMultiAccountResourceCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableMultiAccountResourceCenterResponse) SetBody(v *EnableMultiAccountResourceCenterResponseBody) *EnableMultiAccountResourceCenterResponse {
	s.Body = v
	return s
}

type EnableResourceCenterResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 45357BEF-AB50-5E4D-B05D-5A882A4BE924
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The activation status of the service. Valid values:
	//
	// 	- Pending: The service is being activated.
	//
	// 	- Enabled: The service is activated.
	//
	// example:
	//
	// Pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s EnableResourceCenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableResourceCenterResponseBody) GoString() string {
	return s.String()
}

func (s *EnableResourceCenterResponseBody) SetRequestId(v string) *EnableResourceCenterResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableResourceCenterResponseBody) SetStatus(v string) *EnableResourceCenterResponseBody {
	s.Status = &v
	return s
}

type EnableResourceCenterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableResourceCenterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableResourceCenterResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableResourceCenterResponse) GoString() string {
	return s.String()
}

func (s *EnableResourceCenterResponse) SetHeaders(v map[string]*string) *EnableResourceCenterResponse {
	s.Headers = v
	return s
}

func (s *EnableResourceCenterResponse) SetStatusCode(v int32) *EnableResourceCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableResourceCenterResponse) SetBody(v *EnableResourceCenterResponseBody) *EnableResourceCenterResponse {
	s.Body = v
	return s
}

type ExecuteMultiAccountSQLQueryRequest struct {
	// The SQL statement to be executed.
	//
	// The number of characters in the SQL statement must be less than 2,000.
	//
	// For more information about the SQL syntax, see [Basic SQL syntax](https://help.aliyun.com/document_detail/2539395.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// SELECT 	- FROM resources LIMIT 100;
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The search scope. The value of this parameter can be one of the following items:
	//
	// 	- ID of a resource directory: Resources within the management account and all members of the resource directory are searched.
	//
	// 	- ID of the Root folder: Resources within all members in the Root folder and the subfolders of the Root folder are searched.
	//
	// 	- ID of a folder: Resources within all members in the folder are searched.
	//
	// 	- ID of a member: Resources within the member are searched.
	//
	// 	- ID of a member/ID of a Resource group: Resources within the member in the resource group are searched.
	//
	// For more information about how to obtain the ID of a resource directory, the Root folder, a folder, a member, or a resource group, see [GetResourceDirectory](https://help.aliyun.com/document_detail/159995.html), [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html), [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html), [ListAccounts](https://help.aliyun.com/document_detail/160016.html), or [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// rd-r4****
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s ExecuteMultiAccountSQLQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteMultiAccountSQLQueryRequest) GoString() string {
	return s.String()
}

func (s *ExecuteMultiAccountSQLQueryRequest) SetExpression(v string) *ExecuteMultiAccountSQLQueryRequest {
	s.Expression = &v
	return s
}

func (s *ExecuteMultiAccountSQLQueryRequest) SetScope(v string) *ExecuteMultiAccountSQLQueryRequest {
	s.Scope = &v
	return s
}

type ExecuteMultiAccountSQLQueryResponseBody struct {
	// The columns.
	Columns    []*ExecuteMultiAccountSQLQueryResponseBodyColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	MaxResults *int32                                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 44C8A952-D6B0-5BC8-82D5-93BA02E26F2E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array of search results.
	Rows []interface{} `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
}

func (s ExecuteMultiAccountSQLQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteMultiAccountSQLQueryResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteMultiAccountSQLQueryResponseBody) SetColumns(v []*ExecuteMultiAccountSQLQueryResponseBodyColumns) *ExecuteMultiAccountSQLQueryResponseBody {
	s.Columns = v
	return s
}

func (s *ExecuteMultiAccountSQLQueryResponseBody) SetMaxResults(v int32) *ExecuteMultiAccountSQLQueryResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ExecuteMultiAccountSQLQueryResponseBody) SetNextToken(v string) *ExecuteMultiAccountSQLQueryResponseBody {
	s.NextToken = &v
	return s
}

func (s *ExecuteMultiAccountSQLQueryResponseBody) SetRequestId(v string) *ExecuteMultiAccountSQLQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteMultiAccountSQLQueryResponseBody) SetRows(v []interface{}) *ExecuteMultiAccountSQLQueryResponseBody {
	s.Rows = v
	return s
}

type ExecuteMultiAccountSQLQueryResponseBodyColumns struct {
	// The name of the column.
	//
	// example:
	//
	// resource_id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the column.
	//
	// example:
	//
	// varchar
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ExecuteMultiAccountSQLQueryResponseBodyColumns) String() string {
	return tea.Prettify(s)
}

func (s ExecuteMultiAccountSQLQueryResponseBodyColumns) GoString() string {
	return s.String()
}

func (s *ExecuteMultiAccountSQLQueryResponseBodyColumns) SetName(v string) *ExecuteMultiAccountSQLQueryResponseBodyColumns {
	s.Name = &v
	return s
}

func (s *ExecuteMultiAccountSQLQueryResponseBodyColumns) SetType(v string) *ExecuteMultiAccountSQLQueryResponseBodyColumns {
	s.Type = &v
	return s
}

type ExecuteMultiAccountSQLQueryResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteMultiAccountSQLQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteMultiAccountSQLQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteMultiAccountSQLQueryResponse) GoString() string {
	return s.String()
}

func (s *ExecuteMultiAccountSQLQueryResponse) SetHeaders(v map[string]*string) *ExecuteMultiAccountSQLQueryResponse {
	s.Headers = v
	return s
}

func (s *ExecuteMultiAccountSQLQueryResponse) SetStatusCode(v int32) *ExecuteMultiAccountSQLQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteMultiAccountSQLQueryResponse) SetBody(v *ExecuteMultiAccountSQLQueryResponseBody) *ExecuteMultiAccountSQLQueryResponse {
	s.Body = v
	return s
}

type ExecuteSQLQueryRequest struct {
	// The SQL statement to be executed.
	//
	// The number of characters in the SQL statement must be less than 2,000.
	//
	// For more information about the SQL syntax, see [Basic SQL syntax](https://help.aliyun.com/document_detail/2539395.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// SELECT 	- FROM resources LIMIT 100;
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The search scope.
	//
	// Set this parameter to the ID of a resource group.
	//
	// For information about how to obtain the ID of a resource group, see [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html).
	//
	// example:
	//
	// rg-acfmzawhxxc****
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s ExecuteSQLQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteSQLQueryRequest) GoString() string {
	return s.String()
}

func (s *ExecuteSQLQueryRequest) SetExpression(v string) *ExecuteSQLQueryRequest {
	s.Expression = &v
	return s
}

func (s *ExecuteSQLQueryRequest) SetScope(v string) *ExecuteSQLQueryRequest {
	s.Scope = &v
	return s
}

type ExecuteSQLQueryResponseBody struct {
	// The columns.
	Columns    []*ExecuteSQLQueryResponseBodyColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	MaxResults *int32                                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D696E6EF-3A6D-5770-801E-4982081FE4D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array of search results.
	Rows []interface{} `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
}

func (s ExecuteSQLQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteSQLQueryResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteSQLQueryResponseBody) SetColumns(v []*ExecuteSQLQueryResponseBodyColumns) *ExecuteSQLQueryResponseBody {
	s.Columns = v
	return s
}

func (s *ExecuteSQLQueryResponseBody) SetMaxResults(v int32) *ExecuteSQLQueryResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ExecuteSQLQueryResponseBody) SetNextToken(v string) *ExecuteSQLQueryResponseBody {
	s.NextToken = &v
	return s
}

func (s *ExecuteSQLQueryResponseBody) SetRequestId(v string) *ExecuteSQLQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteSQLQueryResponseBody) SetRows(v []interface{}) *ExecuteSQLQueryResponseBody {
	s.Rows = v
	return s
}

type ExecuteSQLQueryResponseBodyColumns struct {
	// The name of the column.
	//
	// example:
	//
	// resource_id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the column.
	//
	// example:
	//
	// varchar
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ExecuteSQLQueryResponseBodyColumns) String() string {
	return tea.Prettify(s)
}

func (s ExecuteSQLQueryResponseBodyColumns) GoString() string {
	return s.String()
}

func (s *ExecuteSQLQueryResponseBodyColumns) SetName(v string) *ExecuteSQLQueryResponseBodyColumns {
	s.Name = &v
	return s
}

func (s *ExecuteSQLQueryResponseBodyColumns) SetType(v string) *ExecuteSQLQueryResponseBodyColumns {
	s.Type = &v
	return s
}

type ExecuteSQLQueryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteSQLQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteSQLQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteSQLQueryResponse) GoString() string {
	return s.String()
}

func (s *ExecuteSQLQueryResponse) SetHeaders(v map[string]*string) *ExecuteSQLQueryResponse {
	s.Headers = v
	return s
}

func (s *ExecuteSQLQueryResponse) SetStatusCode(v int32) *ExecuteSQLQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteSQLQueryResponse) SetBody(v *ExecuteSQLQueryResponseBody) *ExecuteSQLQueryResponse {
	s.Body = v
	return s
}

type GetExampleQueryRequest struct {
	// The ID of the template.
	//
	// >  You can call the [ListExampleQueries](~~ListExampleQueries~~) operation to obtain the template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sq-0PfKy****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s GetExampleQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExampleQueryRequest) GoString() string {
	return s.String()
}

func (s *GetExampleQueryRequest) SetQueryId(v string) *GetExampleQueryRequest {
	s.QueryId = &v
	return s
}

type GetExampleQueryResponseBody struct {
	// The information about the sample query template.
	ExampleQuery *GetExampleQueryResponseBodyExampleQuery `json:"ExampleQuery,omitempty" xml:"ExampleQuery,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 36A3D9BE-B607-5993-B546-7E19EF65DC00
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetExampleQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExampleQueryResponseBody) GoString() string {
	return s.String()
}

func (s *GetExampleQueryResponseBody) SetExampleQuery(v *GetExampleQueryResponseBodyExampleQuery) *GetExampleQueryResponseBody {
	s.ExampleQuery = v
	return s
}

func (s *GetExampleQueryResponseBody) SetRequestId(v string) *GetExampleQueryResponseBody {
	s.RequestId = &v
	return s
}

type GetExampleQueryResponseBodyExampleQuery struct {
	// The description of the template.
	//
	// example:
	//
	// Queries all resources on which you have permissions and sorts the resources by resource type and resource ID.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The query statement in the template.
	//
	// example:
	//
	// SELECT
	//
	//   resource_id,
	//
	//   resource_name,
	//
	//   region_id,
	//
	//   zone_id,
	//
	//   resource_type,
	//
	//   account_id,
	//
	//   create_time,
	//
	//   resource_group_id,
	//
	//   tags,
	//
	//   ip_addresses,
	//
	//   vpc_id,
	//
	//   v_switch_id
	//
	// FROM
	//
	//   resources
	//
	// ORDER BY
	//
	//   resource_type,
	//
	//   resource_id
	//
	// LIMIT
	//
	//   1000 OFFSET 0;
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// Query of All Alibaba Cloud Resources
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// sq-0PfKy****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s GetExampleQueryResponseBodyExampleQuery) String() string {
	return tea.Prettify(s)
}

func (s GetExampleQueryResponseBodyExampleQuery) GoString() string {
	return s.String()
}

func (s *GetExampleQueryResponseBodyExampleQuery) SetDescription(v string) *GetExampleQueryResponseBodyExampleQuery {
	s.Description = &v
	return s
}

func (s *GetExampleQueryResponseBodyExampleQuery) SetExpression(v string) *GetExampleQueryResponseBodyExampleQuery {
	s.Expression = &v
	return s
}

func (s *GetExampleQueryResponseBodyExampleQuery) SetName(v string) *GetExampleQueryResponseBodyExampleQuery {
	s.Name = &v
	return s
}

func (s *GetExampleQueryResponseBodyExampleQuery) SetQueryId(v string) *GetExampleQueryResponseBodyExampleQuery {
	s.QueryId = &v
	return s
}

type GetExampleQueryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExampleQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExampleQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExampleQueryResponse) GoString() string {
	return s.String()
}

func (s *GetExampleQueryResponse) SetHeaders(v map[string]*string) *GetExampleQueryResponse {
	s.Headers = v
	return s
}

func (s *GetExampleQueryResponse) SetStatusCode(v int32) *GetExampleQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExampleQueryResponse) SetBody(v *GetExampleQueryResponseBody) *GetExampleQueryResponse {
	s.Body = v
	return s
}

type GetMultiAccountResourceCenterServiceStatusResponseBody struct {
	// The initialization status of the feature. Valid values:
	//
	// 	- Pending: The feature is being initialized.
	//
	// 	- Finished: The feature is initialized.
	//
	// example:
	//
	// Pending
	InitialStatus *string `json:"InitialStatus,omitempty" xml:"InitialStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 81671397-1425-51F1-A144-4799E01BEBFF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the feature. Valid values:
	//
	// 	- Enabled: The feature is enabled.
	//
	// 	- Disabled: The feature is disabled.
	//
	// example:
	//
	// Enabled
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
}

func (s GetMultiAccountResourceCenterServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceCenterServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceCenterServiceStatusResponseBody) SetInitialStatus(v string) *GetMultiAccountResourceCenterServiceStatusResponseBody {
	s.InitialStatus = &v
	return s
}

func (s *GetMultiAccountResourceCenterServiceStatusResponseBody) SetRequestId(v string) *GetMultiAccountResourceCenterServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultiAccountResourceCenterServiceStatusResponseBody) SetServiceStatus(v string) *GetMultiAccountResourceCenterServiceStatusResponseBody {
	s.ServiceStatus = &v
	return s
}

type GetMultiAccountResourceCenterServiceStatusResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMultiAccountResourceCenterServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMultiAccountResourceCenterServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceCenterServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceCenterServiceStatusResponse) SetHeaders(v map[string]*string) *GetMultiAccountResourceCenterServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetMultiAccountResourceCenterServiceStatusResponse) SetStatusCode(v int32) *GetMultiAccountResourceCenterServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultiAccountResourceCenterServiceStatusResponse) SetBody(v *GetMultiAccountResourceCenterServiceStatusResponseBody) *GetMultiAccountResourceCenterServiceStatusResponse {
	s.Body = v
	return s
}

type GetMultiAccountResourceConfigurationRequest struct {
	// The ID of the management account or member of the resource directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1619302****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-eb3hji****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The region ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The type of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::VPC::RouteTable
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetMultiAccountResourceConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceConfigurationRequest) SetAccountId(v string) *GetMultiAccountResourceConfigurationRequest {
	s.AccountId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationRequest) SetResourceId(v string) *GetMultiAccountResourceConfigurationRequest {
	s.ResourceId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationRequest) SetResourceRegionId(v string) *GetMultiAccountResourceConfigurationRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationRequest) SetResourceType(v string) *GetMultiAccountResourceConfigurationRequest {
	s.ResourceType = &v
	return s
}

type GetMultiAccountResourceConfigurationResponseBody struct {
	// The ID of the management account or member of the resource directory.
	//
	// example:
	//
	// 1619302****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The configurations of the resource.
	Configuration map[string]interface{} `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// The time when the resource was created.
	//
	// example:
	//
	// 2023-02-14T03:12:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the resource expires.
	//
	// example:
	//
	// 2023-09-18T07:04:21Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The attributes of the IP address.
	IpAddressAttributes []*GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes `json:"IpAddressAttributes,omitempty" xml:"IpAddressAttributes,omitempty" type:"Repeated"`
	// The IP addresses.
	//
	// > Whether this parameter is returned is determined by the Alibaba Cloud service to which the resource belongs.
	IpAddresses []*string `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	// The region ID of the resource.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B2DCC08B-C12A-5705-879C-5A1450016156
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the resource belongs.
	//
	// example:
	//
	// rg-acfmzy6d****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// m-eb3hji****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource.
	//
	// example:
	//
	// test_resource
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// ACS::VPC::RouteTable
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags of the resource.
	Tags []*GetMultiAccountResourceConfigurationResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The zone ID of the resource.
	//
	// example:
	//
	// cn-shanghai-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetMultiAccountResourceConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetAccountId(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.AccountId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetConfiguration(v map[string]interface{}) *GetMultiAccountResourceConfigurationResponseBody {
	s.Configuration = v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetCreateTime(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetExpireTime(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetIpAddressAttributes(v []*GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes) *GetMultiAccountResourceConfigurationResponseBody {
	s.IpAddressAttributes = v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetIpAddresses(v []*string) *GetMultiAccountResourceConfigurationResponseBody {
	s.IpAddresses = v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetRegionId(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetRequestId(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetResourceGroupId(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetResourceId(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.ResourceId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetResourceName(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.ResourceName = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetResourceType(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetTags(v []*GetMultiAccountResourceConfigurationResponseBodyTags) *GetMultiAccountResourceConfigurationResponseBody {
	s.Tags = v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetZoneId(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.ZoneId = &v
	return s
}

type GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes struct {
	// The IP address.
	//
	// example:
	//
	// 172.27.199.42
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The network type. Valid values:
	//
	// 	- **Public**: the Internet
	//
	// 	- **Private**: internal network
	//
	// example:
	//
	// Public
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The version.
	//
	// example:
	//
	// Ipv4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes) SetIpAddress(v string) *GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes {
	s.IpAddress = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes) SetNetworkType(v string) *GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes {
	s.NetworkType = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes) SetVersion(v string) *GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes {
	s.Version = &v
	return s
}

type GetMultiAccountResourceConfigurationResponseBodyTags struct {
	// The key of tag N.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// example:
	//
	// tag-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetMultiAccountResourceConfigurationResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceConfigurationResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceConfigurationResponseBodyTags) SetKey(v string) *GetMultiAccountResourceConfigurationResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBodyTags) SetValue(v string) *GetMultiAccountResourceConfigurationResponseBodyTags {
	s.Value = &v
	return s
}

type GetMultiAccountResourceConfigurationResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMultiAccountResourceConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMultiAccountResourceConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMultiAccountResourceConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceConfigurationResponse) SetHeaders(v map[string]*string) *GetMultiAccountResourceConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponse) SetStatusCode(v int32) *GetMultiAccountResourceConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponse) SetBody(v *GetMultiAccountResourceConfigurationResponseBody) *GetMultiAccountResourceConfigurationResponse {
	s.Body = v
	return s
}

type GetResourceCenterServiceStatusResponseBody struct {
	// The initialization status of the service. Valid values:
	//
	// 	- Pending: The service is being initialized.
	//
	// 	- Finished: The service is initialized.
	//
	// example:
	//
	// Pending
	InitialStatus *string `json:"InitialStatus,omitempty" xml:"InitialStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AD5F848D-CCDC-5464-93E1-4BA50A4826DD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the service. Valid values:
	//
	// 	- Enabled: The service is activated.
	//
	// 	- Disabled: The service is deactivated.
	//
	// example:
	//
	// Enabled
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
}

func (s GetResourceCenterServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceCenterServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceCenterServiceStatusResponseBody) SetInitialStatus(v string) *GetResourceCenterServiceStatusResponseBody {
	s.InitialStatus = &v
	return s
}

func (s *GetResourceCenterServiceStatusResponseBody) SetRequestId(v string) *GetResourceCenterServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceCenterServiceStatusResponseBody) SetServiceStatus(v string) *GetResourceCenterServiceStatusResponseBody {
	s.ServiceStatus = &v
	return s
}

type GetResourceCenterServiceStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceCenterServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceCenterServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceCenterServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetResourceCenterServiceStatusResponse) SetHeaders(v map[string]*string) *GetResourceCenterServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetResourceCenterServiceStatusResponse) SetStatusCode(v int32) *GetResourceCenterServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceCenterServiceStatusResponse) SetBody(v *GetResourceCenterServiceStatusResponseBody) *GetResourceCenterServiceStatusResponse {
	s.Body = v
	return s
}

type GetResourceConfigurationRequest struct {
	// The ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// eip-bp1kyg72m****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The region ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The type of the resource.
	//
	// For more information about the resource types supported by Resource Center, see [Services that work with Resource Center](https://help.aliyun.com/document_detail/477798.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::VPC::RouteTable
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetResourceConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationRequest) SetResourceId(v string) *GetResourceConfigurationRequest {
	s.ResourceId = &v
	return s
}

func (s *GetResourceConfigurationRequest) SetResourceRegionId(v string) *GetResourceConfigurationRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *GetResourceConfigurationRequest) SetResourceType(v string) *GetResourceConfigurationRequest {
	s.ResourceType = &v
	return s
}

type GetResourceConfigurationResponseBody struct {
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 151266687691****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The configurations of the resource.
	Configuration map[string]interface{} `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// The time when the resource was created.
	//
	// example:
	//
	// 2021-06-30T09:20:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the resource expires.
	//
	// example:
	//
	// 2021-07-30T09:20:08Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The attributes of the IP address.
	IpAddressAttributes []*GetResourceConfigurationResponseBodyIpAddressAttributes `json:"IpAddressAttributes,omitempty" xml:"IpAddressAttributes,omitempty" type:"Repeated"`
	// The IP addresses.
	//
	// > Whether this parameter is returned is determined by the Alibaba Cloud service to which the resource belongs.
	IpAddresses []*string `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	// The region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F1CE0D52-32DA-531A-87A4-B9A5B68D5D8E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the resource belongs.
	//
	// example:
	//
	// rg-acfmv4k****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// vtb-uf6978gdqbi****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource.
	//
	// example:
	//
	// group1
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// ACS::VPC::VSwitch
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags of the resource.
	Tags []*GetResourceConfigurationResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The zone ID of the resource.
	//
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetResourceConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationResponseBody) SetAccountId(v string) *GetResourceConfigurationResponseBody {
	s.AccountId = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetConfiguration(v map[string]interface{}) *GetResourceConfigurationResponseBody {
	s.Configuration = v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetCreateTime(v string) *GetResourceConfigurationResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetExpireTime(v string) *GetResourceConfigurationResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetIpAddressAttributes(v []*GetResourceConfigurationResponseBodyIpAddressAttributes) *GetResourceConfigurationResponseBody {
	s.IpAddressAttributes = v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetIpAddresses(v []*string) *GetResourceConfigurationResponseBody {
	s.IpAddresses = v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetRegionId(v string) *GetResourceConfigurationResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetRequestId(v string) *GetResourceConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetResourceGroupId(v string) *GetResourceConfigurationResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetResourceId(v string) *GetResourceConfigurationResponseBody {
	s.ResourceId = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetResourceName(v string) *GetResourceConfigurationResponseBody {
	s.ResourceName = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetResourceType(v string) *GetResourceConfigurationResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetTags(v []*GetResourceConfigurationResponseBodyTags) *GetResourceConfigurationResponseBody {
	s.Tags = v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetZoneId(v string) *GetResourceConfigurationResponseBody {
	s.ZoneId = &v
	return s
}

type GetResourceConfigurationResponseBodyIpAddressAttributes struct {
	// The IP address.
	//
	// example:
	//
	// 192.168.1.2
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The network type. Valid values:
	//
	// 	- **Public**: the Internet
	//
	// 	- **Private**: internal network
	//
	// example:
	//
	// Public
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The version.
	//
	// example:
	//
	// Ipv4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetResourceConfigurationResponseBodyIpAddressAttributes) String() string {
	return tea.Prettify(s)
}

func (s GetResourceConfigurationResponseBodyIpAddressAttributes) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationResponseBodyIpAddressAttributes) SetIpAddress(v string) *GetResourceConfigurationResponseBodyIpAddressAttributes {
	s.IpAddress = &v
	return s
}

func (s *GetResourceConfigurationResponseBodyIpAddressAttributes) SetNetworkType(v string) *GetResourceConfigurationResponseBodyIpAddressAttributes {
	s.NetworkType = &v
	return s
}

func (s *GetResourceConfigurationResponseBodyIpAddressAttributes) SetVersion(v string) *GetResourceConfigurationResponseBodyIpAddressAttributes {
	s.Version = &v
	return s
}

type GetResourceConfigurationResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetResourceConfigurationResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetResourceConfigurationResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationResponseBodyTags) SetKey(v string) *GetResourceConfigurationResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetResourceConfigurationResponseBodyTags) SetValue(v string) *GetResourceConfigurationResponseBodyTags {
	s.Value = &v
	return s
}

type GetResourceConfigurationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationResponse) SetHeaders(v map[string]*string) *GetResourceConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetResourceConfigurationResponse) SetStatusCode(v int32) *GetResourceConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceConfigurationResponse) SetBody(v *GetResourceConfigurationResponseBody) *GetResourceConfigurationResponse {
	s.Body = v
	return s
}

type GetResourceCountsRequest struct {
	// The filter conditions.
	Filter []*GetResourceCountsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The dimension by which resources are queried. Valid values:
	//
	// 	- ResourceType
	//
	// 	- Region
	//
	// 	- ResourceGroupId
	//
	// 	- TagKey
	//
	// 	- TagValue
	//
	// example:
	//
	// ResourceType
	GroupByKey *string `json:"GroupByKey,omitempty" xml:"GroupByKey,omitempty"`
}

func (s GetResourceCountsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceCountsRequest) GoString() string {
	return s.String()
}

func (s *GetResourceCountsRequest) SetFilter(v []*GetResourceCountsRequestFilter) *GetResourceCountsRequest {
	s.Filter = v
	return s
}

func (s *GetResourceCountsRequest) SetGroupByKey(v string) *GetResourceCountsRequest {
	s.GroupByKey = &v
	return s
}

type GetResourceCountsRequestFilter struct {
	// The key of the filter condition. For more information, see `Supported filter parameters`.
	//
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching mode.
	//
	// The value Equals indicates an equal match.
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The values of the filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s GetResourceCountsRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s GetResourceCountsRequestFilter) GoString() string {
	return s.String()
}

func (s *GetResourceCountsRequestFilter) SetKey(v string) *GetResourceCountsRequestFilter {
	s.Key = &v
	return s
}

func (s *GetResourceCountsRequestFilter) SetMatchType(v string) *GetResourceCountsRequestFilter {
	s.MatchType = &v
	return s
}

func (s *GetResourceCountsRequestFilter) SetValue(v []*string) *GetResourceCountsRequestFilter {
	s.Value = v
	return s
}

type GetResourceCountsResponseBody struct {
	// The filter conditions.
	Filters []*GetResourceCountsResponseBodyFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The dimension by which resources are queried.
	//
	// example:
	//
	// ResourceType
	GroupByKey *string `json:"GroupByKey,omitempty" xml:"GroupByKey,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6D98D9B0-318D-56A4-910C-93B5F945AF2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The numbers of resources.
	ResourceCounts []*GetResourceCountsResponseBodyResourceCounts `json:"ResourceCounts,omitempty" xml:"ResourceCounts,omitempty" type:"Repeated"`
}

func (s GetResourceCountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceCountsResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceCountsResponseBody) SetFilters(v []*GetResourceCountsResponseBodyFilters) *GetResourceCountsResponseBody {
	s.Filters = v
	return s
}

func (s *GetResourceCountsResponseBody) SetGroupByKey(v string) *GetResourceCountsResponseBody {
	s.GroupByKey = &v
	return s
}

func (s *GetResourceCountsResponseBody) SetRequestId(v string) *GetResourceCountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceCountsResponseBody) SetResourceCounts(v []*GetResourceCountsResponseBodyResourceCounts) *GetResourceCountsResponseBody {
	s.ResourceCounts = v
	return s
}

type GetResourceCountsResponseBodyFilters struct {
	// The key of the filter condition.
	//
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The values of the filter condition.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetResourceCountsResponseBodyFilters) String() string {
	return tea.Prettify(s)
}

func (s GetResourceCountsResponseBodyFilters) GoString() string {
	return s.String()
}

func (s *GetResourceCountsResponseBodyFilters) SetKey(v string) *GetResourceCountsResponseBodyFilters {
	s.Key = &v
	return s
}

func (s *GetResourceCountsResponseBodyFilters) SetValues(v []*string) *GetResourceCountsResponseBodyFilters {
	s.Values = v
	return s
}

type GetResourceCountsResponseBodyResourceCounts struct {
	// The number of resources.
	//
	// example:
	//
	// 2
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The group name.
	//
	// example:
	//
	// ACS::ECS::NetworkInterface
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s GetResourceCountsResponseBodyResourceCounts) String() string {
	return tea.Prettify(s)
}

func (s GetResourceCountsResponseBodyResourceCounts) GoString() string {
	return s.String()
}

func (s *GetResourceCountsResponseBodyResourceCounts) SetCount(v int64) *GetResourceCountsResponseBodyResourceCounts {
	s.Count = &v
	return s
}

func (s *GetResourceCountsResponseBodyResourceCounts) SetGroupName(v string) *GetResourceCountsResponseBodyResourceCounts {
	s.GroupName = &v
	return s
}

type GetResourceCountsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceCountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceCountsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceCountsResponse) GoString() string {
	return s.String()
}

func (s *GetResourceCountsResponse) SetHeaders(v map[string]*string) *GetResourceCountsResponse {
	s.Headers = v
	return s
}

func (s *GetResourceCountsResponse) SetStatusCode(v int32) *GetResourceCountsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceCountsResponse) SetBody(v *GetResourceCountsResponseBody) *GetResourceCountsResponse {
	s.Body = v
	return s
}

type GetSavedQueryRequest struct {
	// The template ID.
	//
	// >  You can call the [ListSavedQueries](~~ListSavedQueries~~) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sq-GeAck****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s GetSavedQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSavedQueryRequest) GoString() string {
	return s.String()
}

func (s *GetSavedQueryRequest) SetQueryId(v string) *GetSavedQueryRequest {
	s.QueryId = &v
	return s
}

type GetSavedQueryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6D98D9B0-318D-56A4-910C-93B5F945AF2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the template.
	SavedQuery *GetSavedQueryResponseBodySavedQuery `json:"SavedQuery,omitempty" xml:"SavedQuery,omitempty" type:"Struct"`
}

func (s GetSavedQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSavedQueryResponseBody) GoString() string {
	return s.String()
}

func (s *GetSavedQueryResponseBody) SetRequestId(v string) *GetSavedQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSavedQueryResponseBody) SetSavedQuery(v *GetSavedQueryResponseBodySavedQuery) *GetSavedQueryResponseBody {
	s.SavedQuery = v
	return s
}

type GetSavedQueryResponseBodySavedQuery struct {
	// The time when the template was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-10-30T01:43:16Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the template.
	//
	// example:
	//
	// Queries all resources on which you have permissions and sorts the resources by resource type and resource ID.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The query statement in the template.
	//
	// example:
	//
	// SELECT 	- FROM resources;
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// Query of All Alibaba Cloud Resources
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The template ID.
	//
	// example:
	//
	// sq-GeAck****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// The time when the template was updated. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-10-30T01:43:16Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetSavedQueryResponseBodySavedQuery) String() string {
	return tea.Prettify(s)
}

func (s GetSavedQueryResponseBodySavedQuery) GoString() string {
	return s.String()
}

func (s *GetSavedQueryResponseBodySavedQuery) SetCreateTime(v string) *GetSavedQueryResponseBodySavedQuery {
	s.CreateTime = &v
	return s
}

func (s *GetSavedQueryResponseBodySavedQuery) SetDescription(v string) *GetSavedQueryResponseBodySavedQuery {
	s.Description = &v
	return s
}

func (s *GetSavedQueryResponseBodySavedQuery) SetExpression(v string) *GetSavedQueryResponseBodySavedQuery {
	s.Expression = &v
	return s
}

func (s *GetSavedQueryResponseBodySavedQuery) SetName(v string) *GetSavedQueryResponseBodySavedQuery {
	s.Name = &v
	return s
}

func (s *GetSavedQueryResponseBodySavedQuery) SetQueryId(v string) *GetSavedQueryResponseBodySavedQuery {
	s.QueryId = &v
	return s
}

func (s *GetSavedQueryResponseBodySavedQuery) SetUpdateTime(v string) *GetSavedQueryResponseBodySavedQuery {
	s.UpdateTime = &v
	return s
}

type GetSavedQueryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSavedQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSavedQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSavedQueryResponse) GoString() string {
	return s.String()
}

func (s *GetSavedQueryResponse) SetHeaders(v map[string]*string) *GetSavedQueryResponse {
	s.Headers = v
	return s
}

func (s *GetSavedQueryResponse) SetStatusCode(v int32) *GetSavedQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSavedQueryResponse) SetBody(v *GetSavedQueryResponseBody) *GetSavedQueryResponse {
	s.Body = v
	return s
}

type ListExampleQueriesRequest struct {
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 50.
	//
	// Default value: 50.
	//
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListExampleQueriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExampleQueriesRequest) GoString() string {
	return s.String()
}

func (s *ListExampleQueriesRequest) SetMaxResults(v string) *ListExampleQueriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExampleQueriesRequest) SetNextToken(v string) *ListExampleQueriesRequest {
	s.NextToken = &v
	return s
}

type ListExampleQueriesResponseBody struct {
	// The information about the sample query templates.
	ExampleQueries []*ListExampleQueriesResponseBodyExampleQueries `json:"ExampleQueries,omitempty" xml:"ExampleQueries,omitempty" type:"Repeated"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D696E6EF-3A6D-5770-801E-4982081FE4D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListExampleQueriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExampleQueriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListExampleQueriesResponseBody) SetExampleQueries(v []*ListExampleQueriesResponseBodyExampleQueries) *ListExampleQueriesResponseBody {
	s.ExampleQueries = v
	return s
}

func (s *ListExampleQueriesResponseBody) SetMaxResults(v string) *ListExampleQueriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListExampleQueriesResponseBody) SetNextToken(v string) *ListExampleQueriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListExampleQueriesResponseBody) SetRequestId(v string) *ListExampleQueriesResponseBody {
	s.RequestId = &v
	return s
}

type ListExampleQueriesResponseBodyExampleQueries struct {
	// The description of the template.
	//
	// example:
	//
	// Queries all resources on which you have permissions and sorts the resources by resource type and resource ID.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// Query of All Alibaba Cloud Resources
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// sq-0PfKy****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s ListExampleQueriesResponseBodyExampleQueries) String() string {
	return tea.Prettify(s)
}

func (s ListExampleQueriesResponseBodyExampleQueries) GoString() string {
	return s.String()
}

func (s *ListExampleQueriesResponseBodyExampleQueries) SetDescription(v string) *ListExampleQueriesResponseBodyExampleQueries {
	s.Description = &v
	return s
}

func (s *ListExampleQueriesResponseBodyExampleQueries) SetName(v string) *ListExampleQueriesResponseBodyExampleQueries {
	s.Name = &v
	return s
}

func (s *ListExampleQueriesResponseBodyExampleQueries) SetQueryId(v string) *ListExampleQueriesResponseBodyExampleQueries {
	s.QueryId = &v
	return s
}

type ListExampleQueriesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExampleQueriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExampleQueriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExampleQueriesResponse) GoString() string {
	return s.String()
}

func (s *ListExampleQueriesResponse) SetHeaders(v map[string]*string) *ListExampleQueriesResponse {
	s.Headers = v
	return s
}

func (s *ListExampleQueriesResponse) SetStatusCode(v int32) *ListExampleQueriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExampleQueriesResponse) SetBody(v *ListExampleQueriesResponseBody) *ListExampleQueriesResponse {
	s.Body = v
	return s
}

type ListFiltersResponseBody struct {
	// The name of the default filter.
	DefaultFilterName *string `json:"DefaultFilterName,omitempty" xml:"DefaultFilterName,omitempty"`
	// The configurations of the filter.
	Filters []*ListFiltersResponseBodyFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// AA39FB9C-CB74-5E73-8DFE-3A2B096F0759
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFiltersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFiltersResponseBody) GoString() string {
	return s.String()
}

func (s *ListFiltersResponseBody) SetDefaultFilterName(v string) *ListFiltersResponseBody {
	s.DefaultFilterName = &v
	return s
}

func (s *ListFiltersResponseBody) SetFilters(v []*ListFiltersResponseBodyFilters) *ListFiltersResponseBody {
	s.Filters = v
	return s
}

func (s *ListFiltersResponseBody) SetRequestId(v string) *ListFiltersResponseBody {
	s.RequestId = &v
	return s
}

type ListFiltersResponseBodyFilters struct {
	// The configurations of the filter.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "regions": [],
	//
	//   "tagFilters": [
	//
	//     [{ "type": "notContainTagKey", "tagKey": "xxx", "tagValue": "" }],
	//
	//     [{ "tagKey": "xxx", "tagValue": "xxx" }]
	//
	//   ],
	//
	//   "resourceTypes": [
	//
	//     "ACS::ECS::AutoSnapshotPolicy"
	//
	//   ]
	//
	// }
	FilterConfiguration *string `json:"FilterConfiguration,omitempty" xml:"FilterConfiguration,omitempty"`
	// The name of the filter.
	FilterName *string `json:"FilterName,omitempty" xml:"FilterName,omitempty"`
}

func (s ListFiltersResponseBodyFilters) String() string {
	return tea.Prettify(s)
}

func (s ListFiltersResponseBodyFilters) GoString() string {
	return s.String()
}

func (s *ListFiltersResponseBodyFilters) SetFilterConfiguration(v string) *ListFiltersResponseBodyFilters {
	s.FilterConfiguration = &v
	return s
}

func (s *ListFiltersResponseBodyFilters) SetFilterName(v string) *ListFiltersResponseBodyFilters {
	s.FilterName = &v
	return s
}

type ListFiltersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFiltersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFiltersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFiltersResponse) GoString() string {
	return s.String()
}

func (s *ListFiltersResponse) SetHeaders(v map[string]*string) *ListFiltersResponse {
	s.Headers = v
	return s
}

func (s *ListFiltersResponse) SetStatusCode(v int32) *ListFiltersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFiltersResponse) SetBody(v *ListFiltersResponseBody) *ListFiltersResponse {
	s.Body = v
	return s
}

type ListMultiAccountResourceGroupsRequest struct {
	// The ID of the management account or member of the resource directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1394339739****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAS2Nboi3t4xGrdlG5/Ks/Q1xPG9jzviYEuZydevXIkgF
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The IDs of resource groups.
	ResourceGroupIds []*string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
}

func (s ListMultiAccountResourceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountResourceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceGroupsRequest) SetAccountId(v string) *ListMultiAccountResourceGroupsRequest {
	s.AccountId = &v
	return s
}

func (s *ListMultiAccountResourceGroupsRequest) SetMaxResults(v int32) *ListMultiAccountResourceGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMultiAccountResourceGroupsRequest) SetNextToken(v string) *ListMultiAccountResourceGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountResourceGroupsRequest) SetResourceGroupIds(v []*string) *ListMultiAccountResourceGroupsRequest {
	s.ResourceGroupIds = v
	return s
}

type ListMultiAccountResourceGroupsResponseBody struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAU5VsT9R1adMTuz9GzginZ3Y+7Y/5JATS+6q5GK9kT75
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0FF0A66E-781F-51EE-9531-928F197558F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resource groups.
	ResourceGroups []*ListMultiAccountResourceGroupsResponseBodyResourceGroups `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty" type:"Repeated"`
}

func (s ListMultiAccountResourceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountResourceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceGroupsResponseBody) SetNextToken(v string) *ListMultiAccountResourceGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBody) SetRequestId(v string) *ListMultiAccountResourceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBody) SetResourceGroups(v []*ListMultiAccountResourceGroupsResponseBodyResourceGroups) *ListMultiAccountResourceGroupsResponseBody {
	s.ResourceGroups = v
	return s
}

type ListMultiAccountResourceGroupsResponseBodyResourceGroups struct {
	// The ID of the management account or member of the resource directory.
	//
	// example:
	//
	// 1394339739****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the resource group was created.
	//
	// example:
	//
	// 2021-06-30T09:20:08Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the resource group.
	//
	// example:
	//
	// group1
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The unique identifier of the resource group.
	//
	// example:
	//
	// my-project
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the resource group. Valid values:
	//
	// 	- Creating: The resource group is being created.
	//
	// 	- OK: The resource group is created.
	//
	// 	- PendingDelete: The resource group is waiting to be deleted.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListMultiAccountResourceGroupsResponseBodyResourceGroups) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountResourceGroupsResponseBodyResourceGroups) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) SetAccountId(v string) *ListMultiAccountResourceGroupsResponseBodyResourceGroups {
	s.AccountId = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) SetCreateDate(v string) *ListMultiAccountResourceGroupsResponseBodyResourceGroups {
	s.CreateDate = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) SetDisplayName(v string) *ListMultiAccountResourceGroupsResponseBodyResourceGroups {
	s.DisplayName = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) SetId(v string) *ListMultiAccountResourceGroupsResponseBodyResourceGroups {
	s.Id = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) SetName(v string) *ListMultiAccountResourceGroupsResponseBodyResourceGroups {
	s.Name = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponseBodyResourceGroups) SetStatus(v string) *ListMultiAccountResourceGroupsResponseBodyResourceGroups {
	s.Status = &v
	return s
}

type ListMultiAccountResourceGroupsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMultiAccountResourceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMultiAccountResourceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountResourceGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListMultiAccountResourceGroupsResponse) SetHeaders(v map[string]*string) *ListMultiAccountResourceGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListMultiAccountResourceGroupsResponse) SetStatusCode(v int32) *ListMultiAccountResourceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultiAccountResourceGroupsResponse) SetBody(v *ListMultiAccountResourceGroupsResponseBody) *ListMultiAccountResourceGroupsResponse {
	s.Body = v
	return s
}

type ListMultiAccountTagKeysRequest struct {
	// The matching mode. Valid values:
	//
	// 	- Equals: equal match
	//
	// 	- Prefix: match by prefix
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// If the total number of entries returned for the current request exceeds the value of the `MaxResults` parameter, the entries are truncated. In this case, you can use the `token` to initiate another request and obtain the remaining entries.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The search scope. You can set the value to one of the following items:
	//
	// 	- ID of a resource directory: Resources within the management account and all members of the resource directory are searched. You can call the [GetResourceDirectory](https://help.aliyun.com/document_detail/159995.html) operation to obtain the ID.
	//
	// 	- ID of the Root folder: Resources within all members in the Root folder and the subfolders of the Root folder are searched. You can call the [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html) operation to obtain the ID.
	//
	// 	- ID of a folder: Resources within all members in the folder are searched. You can call the [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html) operation to obtain the ID.
	//
	// 	- ID of a member: Resources within the member are searched. You can call the [ListAccounts](https://help.aliyun.com/document_detail/160016.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rd-r4****
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The tag key.
	//
	// example:
	//
	// test_key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListMultiAccountTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListMultiAccountTagKeysRequest) SetMatchType(v string) *ListMultiAccountTagKeysRequest {
	s.MatchType = &v
	return s
}

func (s *ListMultiAccountTagKeysRequest) SetMaxResults(v int32) *ListMultiAccountTagKeysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMultiAccountTagKeysRequest) SetNextToken(v string) *ListMultiAccountTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountTagKeysRequest) SetScope(v string) *ListMultiAccountTagKeysRequest {
	s.Scope = &v
	return s
}

func (s *ListMultiAccountTagKeysRequest) SetTagKey(v string) *ListMultiAccountTagKeysRequest {
	s.TagKey = &v
	return s
}

type ListMultiAccountTagKeysResponseBody struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAARfZmVDe9NvRXloR5+8CK9nNJufMdRA7W1miLC1P****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// FA6086F9-6363-51A5-A507-88E3201EBCCB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag keys.
	TagKeys []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s ListMultiAccountTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultiAccountTagKeysResponseBody) SetNextToken(v string) *ListMultiAccountTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountTagKeysResponseBody) SetRequestId(v string) *ListMultiAccountTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultiAccountTagKeysResponseBody) SetTagKeys(v []*string) *ListMultiAccountTagKeysResponseBody {
	s.TagKeys = v
	return s
}

type ListMultiAccountTagKeysResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMultiAccountTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMultiAccountTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListMultiAccountTagKeysResponse) SetHeaders(v map[string]*string) *ListMultiAccountTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListMultiAccountTagKeysResponse) SetStatusCode(v int32) *ListMultiAccountTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultiAccountTagKeysResponse) SetBody(v *ListMultiAccountTagKeysResponseBody) *ListMultiAccountTagKeysResponse {
	s.Body = v
	return s
}

type ListMultiAccountTagValuesRequest struct {
	// The matching mode. Valid values:
	//
	// 	- Equals: equal match
	//
	// 	- Prefix: match by prefix
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// If the total number of entries returned for the current request exceeds the value of the `MaxResults` parameter, the entries are truncated. In this case, you can use the `token` to initiate another request and obtain the remaining entries.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The search scope. You can set the value to one of the following items:
	//
	// 	- ID of a resource directory: Resources within the management account and all members of the resource directory are searched. You can call the [GetResourceDirectory](https://help.aliyun.com/document_detail/159995.html) operation to obtain the ID.
	//
	// 	- ID of the Root folder: Resources within all members in the Root folder and the subfolders of the Root folder are searched. You can call the [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html) operation to obtain the ID.
	//
	// 	- ID of a folder: Resources within all members in the folder are searched. You can call the [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html) operation to obtain the ID.
	//
	// 	- ID of a member: Resources within the member are searched. You can call the [ListAccounts](https://help.aliyun.com/document_detail/160016.html) operation to obtain the ID.
	//
	// example:
	//
	// rd-r4****
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The tag key.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test_value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListMultiAccountTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListMultiAccountTagValuesRequest) SetMatchType(v string) *ListMultiAccountTagValuesRequest {
	s.MatchType = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) SetMaxResults(v int32) *ListMultiAccountTagValuesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) SetNextToken(v string) *ListMultiAccountTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) SetScope(v string) *ListMultiAccountTagValuesRequest {
	s.Scope = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) SetTagKey(v string) *ListMultiAccountTagValuesRequest {
	s.TagKey = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) SetTagValue(v string) *ListMultiAccountTagValuesRequest {
	s.TagValue = &v
	return s
}

type ListMultiAccountTagValuesResponseBody struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 36A3D9BE-B607-5993-B546-7E19EF65DC00
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag values.
	TagValues []*string `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
}

func (s ListMultiAccountTagValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultiAccountTagValuesResponseBody) SetNextToken(v string) *ListMultiAccountTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountTagValuesResponseBody) SetRequestId(v string) *ListMultiAccountTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultiAccountTagValuesResponseBody) SetTagValues(v []*string) *ListMultiAccountTagValuesResponseBody {
	s.TagValues = v
	return s
}

type ListMultiAccountTagValuesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMultiAccountTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMultiAccountTagValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMultiAccountTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListMultiAccountTagValuesResponse) SetHeaders(v map[string]*string) *ListMultiAccountTagValuesResponse {
	s.Headers = v
	return s
}

func (s *ListMultiAccountTagValuesResponse) SetStatusCode(v int32) *ListMultiAccountTagValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultiAccountTagValuesResponse) SetBody(v *ListMultiAccountTagValuesResponseBody) *ListMultiAccountTagValuesResponse {
	s.Body = v
	return s
}

type ListResourceTypesRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US: English
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The query conditions.
	Query []*string `json:"Query,omitempty" xml:"Query,omitempty" type:"Repeated"`
	// The resource type.
	//
	// For more information about the resource types that are supported by Resource Center, see [Services that work with Resource Center](https://help.aliyun.com/document_detail/477798.html).
	//
	// example:
	//
	// ACS::ACK::Cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListResourceTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesRequest) GoString() string {
	return s.String()
}

func (s *ListResourceTypesRequest) SetAcceptLanguage(v string) *ListResourceTypesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListResourceTypesRequest) SetQuery(v []*string) *ListResourceTypesRequest {
	s.Query = v
	return s
}

func (s *ListResourceTypesRequest) SetResourceType(v string) *ListResourceTypesRequest {
	s.ResourceType = &v
	return s
}

type ListResourceTypesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E5556E4C-479A-5BBB-B325-F07563E7E917
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resource types.
	ResourceTypes []*ListResourceTypesResponseBodyResourceTypes `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s ListResourceTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBody) SetRequestId(v string) *ListResourceTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceTypesResponseBody) SetResourceTypes(v []*ListResourceTypesResponseBodyResourceTypes) *ListResourceTypesResponseBody {
	s.ResourceTypes = v
	return s
}

type ListResourceTypesResponseBodyResourceTypes struct {
	// if can be null:
	// true
	Authorized  *bool                                                  `json:"Authorized,omitempty" xml:"Authorized,omitempty"`
	CodeMapping *ListResourceTypesResponseBodyResourceTypesCodeMapping `json:"CodeMapping,omitempty" xml:"CodeMapping,omitempty" type:"Struct"`
	// The supported filter conditions.
	FilterKeys []*string `json:"FilterKeys,omitempty" xml:"FilterKeys,omitempty" type:"Repeated"`
	// The name of the Alibaba Cloud service.
	//
	// example:
	//
	// Container Service for Kubernetes
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ACS::ACK::Cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The name of the resource type.
	//
	// example:
	//
	// Cluster
	ResourceTypeName *string `json:"ResourceTypeName,omitempty" xml:"ResourceTypeName,omitempty"`
}

func (s ListResourceTypesResponseBodyResourceTypes) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypes) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetAuthorized(v bool) *ListResourceTypesResponseBodyResourceTypes {
	s.Authorized = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetCodeMapping(v *ListResourceTypesResponseBodyResourceTypesCodeMapping) *ListResourceTypesResponseBodyResourceTypes {
	s.CodeMapping = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetFilterKeys(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.FilterKeys = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetProductName(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.ProductName = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetResourceType(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.ResourceType = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetResourceTypeName(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.ResourceTypeName = &v
	return s
}

type ListResourceTypesResponseBodyResourceTypesCodeMapping struct {
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	Tag           *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListResourceTypesResponseBodyResourceTypesCodeMapping) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypesCodeMapping) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypesCodeMapping) SetResourceGroup(v string) *ListResourceTypesResponseBodyResourceTypesCodeMapping {
	s.ResourceGroup = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesCodeMapping) SetTag(v string) *ListResourceTypesResponseBodyResourceTypesCodeMapping {
	s.Tag = &v
	return s
}

type ListResourceTypesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponse) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponse) SetHeaders(v map[string]*string) *ListResourceTypesResponse {
	s.Headers = v
	return s
}

func (s *ListResourceTypesResponse) SetStatusCode(v int32) *ListResourceTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceTypesResponse) SetBody(v *ListResourceTypesResponseBody) *ListResourceTypesResponse {
	s.Body = v
	return s
}

type ListSavedQueriesRequest struct {
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 50.
	//
	// Default value: 50.
	//
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListSavedQueriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSavedQueriesRequest) GoString() string {
	return s.String()
}

func (s *ListSavedQueriesRequest) SetMaxResults(v string) *ListSavedQueriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSavedQueriesRequest) SetNextToken(v string) *ListSavedQueriesRequest {
	s.NextToken = &v
	return s
}

type ListSavedQueriesResponseBody struct {
	// The maximum number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAARfZmVDe9NvRXloR5+8CK9nNJufMdRA7W1miLC1P****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D696E6EF-3A6D-5770-801E-4982081FE4D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the custom query templates.
	SavedQueries []*ListSavedQueriesResponseBodySavedQueries `json:"SavedQueries,omitempty" xml:"SavedQueries,omitempty" type:"Repeated"`
}

func (s ListSavedQueriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSavedQueriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSavedQueriesResponseBody) SetMaxResults(v string) *ListSavedQueriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSavedQueriesResponseBody) SetNextToken(v string) *ListSavedQueriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSavedQueriesResponseBody) SetRequestId(v string) *ListSavedQueriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSavedQueriesResponseBody) SetSavedQueries(v []*ListSavedQueriesResponseBodySavedQueries) *ListSavedQueriesResponseBody {
	s.SavedQueries = v
	return s
}

type ListSavedQueriesResponseBodySavedQueries struct {
	// The time when the template was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-01-13T05:50:35Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the template.
	//
	// example:
	//
	// Queries all resources on which you have permissions and sorts the resources by resource type and resource ID.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The template name.
	//
	// example:
	//
	// Query of All Alibaba Cloud Resources
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The template ID.
	//
	// example:
	//
	// sq-GeAck****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// The time when the template was updated. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-03-14 10:27:07
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListSavedQueriesResponseBodySavedQueries) String() string {
	return tea.Prettify(s)
}

func (s ListSavedQueriesResponseBodySavedQueries) GoString() string {
	return s.String()
}

func (s *ListSavedQueriesResponseBodySavedQueries) SetCreateTime(v string) *ListSavedQueriesResponseBodySavedQueries {
	s.CreateTime = &v
	return s
}

func (s *ListSavedQueriesResponseBodySavedQueries) SetDescription(v string) *ListSavedQueriesResponseBodySavedQueries {
	s.Description = &v
	return s
}

func (s *ListSavedQueriesResponseBodySavedQueries) SetName(v string) *ListSavedQueriesResponseBodySavedQueries {
	s.Name = &v
	return s
}

func (s *ListSavedQueriesResponseBodySavedQueries) SetQueryId(v string) *ListSavedQueriesResponseBodySavedQueries {
	s.QueryId = &v
	return s
}

func (s *ListSavedQueriesResponseBodySavedQueries) SetUpdateTime(v string) *ListSavedQueriesResponseBodySavedQueries {
	s.UpdateTime = &v
	return s
}

type ListSavedQueriesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSavedQueriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSavedQueriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSavedQueriesResponse) GoString() string {
	return s.String()
}

func (s *ListSavedQueriesResponse) SetHeaders(v map[string]*string) *ListSavedQueriesResponse {
	s.Headers = v
	return s
}

func (s *ListSavedQueriesResponse) SetStatusCode(v int32) *ListSavedQueriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSavedQueriesResponse) SetBody(v *ListSavedQueriesResponseBody) *ListSavedQueriesResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	// The matching mode. Valid values:
	//
	// 	- Equals: equal match
	//
	// 	- Prefix: match by prefix
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// If the total number of entries returned for the current request exceeds the value of the `MaxResults` parameter, the entries are truncated. In this case, you can use the `token` to initiate another request and obtain the remaining entries.
	//
	// example:
	//
	// AAAAAUYb00R0gHZBE8FVDeoh2ME93VeeEPUHs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The tag key.
	//
	// example:
	//
	// test_key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetMatchType(v string) *ListTagKeysRequest {
	s.MatchType = &v
	return s
}

func (s *ListTagKeysRequest) SetMaxResults(v int32) *ListTagKeysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetTagKey(v string) *ListTagKeysRequest {
	s.TagKey = &v
	return s
}

type ListTagKeysResponseBody struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAUDnubHKJbVTCdlIGYUPtsu3EoN3bfdgjDA****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 44C8A952-D6B0-5BC8-82D5-93BA02E26F2E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag keys.
	TagKeys []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetNextToken(v string) *ListTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetTagKeys(v []*string) *ListTagKeysResponseBody {
	s.TagKeys = v
	return s
}

type ListTagKeysResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetStatusCode(v int32) *ListTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagValuesRequest struct {
	// The matching mode. Valid values:
	//
	// 	- Equals: equal match
	//
	// 	- Prefix: match by prefix
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// If the total number of entries returned for the current request exceeds the value of the `MaxResults` parameter, the entries are truncated. In this case, you can use the `token` to initiate another request and obtain the remaining entries.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The tag key.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test_value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) SetMatchType(v string) *ListTagValuesRequest {
	s.MatchType = &v
	return s
}

func (s *ListTagValuesRequest) SetMaxResults(v int32) *ListTagValuesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetTagKey(v string) *ListTagValuesRequest {
	s.TagKey = &v
	return s
}

func (s *ListTagValuesRequest) SetTagValue(v string) *ListTagValuesRequest {
	s.TagValue = &v
	return s
}

type ListTagValuesResponseBody struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C1840B83-1193-5E83-AFA6-4B8D303E29F5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag values.
	TagValues []*string `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
}

func (s ListTagValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBody) SetNextToken(v string) *ListTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesResponseBody) SetRequestId(v string) *ListTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagValuesResponseBody) SetTagValues(v []*string) *ListTagValuesResponseBody {
	s.TagValues = v
	return s
}

type ListTagValuesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponse) SetHeaders(v map[string]*string) *ListTagValuesResponse {
	s.Headers = v
	return s
}

func (s *ListTagValuesResponse) SetStatusCode(v int32) *ListTagValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagValuesResponse) SetBody(v *ListTagValuesResponseBody) *ListTagValuesResponse {
	s.Body = v
	return s
}

type SearchMultiAccountResourcesRequest struct {
	// The filter conditions.
	Filter []*SearchMultiAccountResourcesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The maximum number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// If the total number of entries returned for the current request exceeds the value of the `MaxResults` parameter, the entries are truncated. In this case, you can use the token to initiate another request and obtain the remaining entries.``
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The search scope. You can set the value to one of the following items:
	//
	// 	- ID of a resource directory: Resources within the management account and all members of the resource directory are searched. You can call the [GetResourceDirectory](https://help.aliyun.com/document_detail/159995.html) operation to obtain the ID.
	//
	// 	- ID of the Root folder: Resources within all members in the Root folder and the subfolders of the Root folder are searched. You can call the [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html) operation to obtain the ID.
	//
	// 	- ID of a folder: Resources within all members in the folder are searched. You can call the [ListFoldersForParent](https://help.aliyun.com/document_detail/159997.html) operation to obtain the ID.
	//
	// 	- ID of a member: Resources within the member are searched. You can call the [ListAccounts](https://help.aliyun.com/document_detail/160016.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rd-r4****
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The method that is used to sort the entries returned.
	SortCriterion *SearchMultiAccountResourcesRequestSortCriterion `json:"SortCriterion,omitempty" xml:"SortCriterion,omitempty" type:"Struct"`
}

func (s SearchMultiAccountResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesRequest) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesRequest) SetFilter(v []*SearchMultiAccountResourcesRequestFilter) *SearchMultiAccountResourcesRequest {
	s.Filter = v
	return s
}

func (s *SearchMultiAccountResourcesRequest) SetMaxResults(v int32) *SearchMultiAccountResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchMultiAccountResourcesRequest) SetNextToken(v string) *SearchMultiAccountResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *SearchMultiAccountResourcesRequest) SetScope(v string) *SearchMultiAccountResourcesRequest {
	s.Scope = &v
	return s
}

func (s *SearchMultiAccountResourcesRequest) SetSortCriterion(v *SearchMultiAccountResourcesRequestSortCriterion) *SearchMultiAccountResourcesRequest {
	s.SortCriterion = v
	return s
}

type SearchMultiAccountResourcesRequestFilter struct {
	// The key of the filter condition. For more information, see `Supported filter parameters`.
	//
	// example:
	//
	// ResourceGroupId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching mode.
	//
	// The value Equals indicates an equal match.
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The values of the filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s SearchMultiAccountResourcesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesRequestFilter) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesRequestFilter) SetKey(v string) *SearchMultiAccountResourcesRequestFilter {
	s.Key = &v
	return s
}

func (s *SearchMultiAccountResourcesRequestFilter) SetMatchType(v string) *SearchMultiAccountResourcesRequestFilter {
	s.MatchType = &v
	return s
}

func (s *SearchMultiAccountResourcesRequestFilter) SetValue(v []*string) *SearchMultiAccountResourcesRequestFilter {
	s.Value = v
	return s
}

type SearchMultiAccountResourcesRequestSortCriterion struct {
	// The attribute based on which the entries are sorted.
	//
	// The value CreateTime indicates the creation time of resources.
	//
	// example:
	//
	// CreateTime
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The order in which the entries are sorted. Valid values:
	//
	// 	- ASC: The entries are sorted in ascending order. This value is the default value.
	//
	// 	- DESC: The entries are sorted in descending order.
	//
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s SearchMultiAccountResourcesRequestSortCriterion) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesRequestSortCriterion) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesRequestSortCriterion) SetKey(v string) *SearchMultiAccountResourcesRequestSortCriterion {
	s.Key = &v
	return s
}

func (s *SearchMultiAccountResourcesRequestSortCriterion) SetOrder(v string) *SearchMultiAccountResourcesRequestSortCriterion {
	s.Order = &v
	return s
}

type SearchMultiAccountResourcesResponseBody struct {
	// The filter conditions.
	Filters []*SearchMultiAccountResourcesResponseBodyFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// The pagination token that is used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EFA806B9-7F36-55AB-8B7A-D680C2C5EE57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resources.
	Resources []*SearchMultiAccountResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The search scope.
	//
	// 	- ID of a resource directory: Resources within the management account and all members of the resource directory are searched.
	//
	// 	- ID of the Root folder: Resources within all members in the Root folder and the subfolders of the Root folder are searched.
	//
	// 	- ID of a folder: Resources within all members in the folder are searched.
	//
	// 	- ID of a member: Resources within the member are searched.
	//
	// example:
	//
	// rd-r4****
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s SearchMultiAccountResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesResponseBody) SetFilters(v []*SearchMultiAccountResourcesResponseBodyFilters) *SearchMultiAccountResourcesResponseBody {
	s.Filters = v
	return s
}

func (s *SearchMultiAccountResourcesResponseBody) SetMaxResults(v int32) *SearchMultiAccountResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBody) SetNextToken(v string) *SearchMultiAccountResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBody) SetRequestId(v string) *SearchMultiAccountResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBody) SetResources(v []*SearchMultiAccountResourcesResponseBodyResources) *SearchMultiAccountResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *SearchMultiAccountResourcesResponseBody) SetScope(v string) *SearchMultiAccountResourcesResponseBody {
	s.Scope = &v
	return s
}

type SearchMultiAccountResourcesResponseBodyFilters struct {
	// The key of the filter condition.
	//
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching mode.
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The values of the filter condition.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s SearchMultiAccountResourcesResponseBodyFilters) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesResponseBodyFilters) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesResponseBodyFilters) SetKey(v string) *SearchMultiAccountResourcesResponseBodyFilters {
	s.Key = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyFilters) SetMatchType(v string) *SearchMultiAccountResourcesResponseBodyFilters {
	s.MatchType = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyFilters) SetValues(v []*string) *SearchMultiAccountResourcesResponseBodyFilters {
	s.Values = v
	return s
}

type SearchMultiAccountResourcesResponseBodyResources struct {
	// The ID of the management account or member of the resource directory.
	//
	// example:
	//
	// 151266687691****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the resource was created.
	//
	// >  Whether this parameter is returned is determined by the Alibaba Cloud service to which the resource belongs.
	//
	// example:
	//
	// 2021-06-30T09:20:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the resource expires.
	//
	// example:
	//
	// 2023-06-14T14:35:45Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The attributes of the IP address.
	IpAddressAttributes []*SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes `json:"IpAddressAttributes,omitempty" xml:"IpAddressAttributes,omitempty" type:"Repeated"`
	// The IP addresses.
	//
	// >  Whether this parameter is returned is determined by the Alibaba Cloud service to which the resource belongs.
	IpAddresses []*string `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// vtb-bp11lbh452fr8940s****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource name.
	//
	// example:
	//
	// group1
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ACS::VPC::RouteTable
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	Tags []*SearchMultiAccountResourcesResponseBodyResourcesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The zone ID.
	//
	// >  Whether this parameter is returned is determined by the Alibaba Cloud service to which the resource belongs.
	//
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s SearchMultiAccountResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetAccountId(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.AccountId = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetCreateTime(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.CreateTime = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetExpireTime(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.ExpireTime = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetIpAddressAttributes(v []*SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes) *SearchMultiAccountResourcesResponseBodyResources {
	s.IpAddressAttributes = v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetIpAddresses(v []*string) *SearchMultiAccountResourcesResponseBodyResources {
	s.IpAddresses = v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetRegionId(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.RegionId = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetResourceGroupId(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.ResourceGroupId = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetResourceId(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.ResourceId = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetResourceName(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.ResourceName = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetResourceType(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.ResourceType = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetTags(v []*SearchMultiAccountResourcesResponseBodyResourcesTags) *SearchMultiAccountResourcesResponseBodyResources {
	s.Tags = v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResources) SetZoneId(v string) *SearchMultiAccountResourcesResponseBodyResources {
	s.ZoneId = &v
	return s
}

type SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes struct {
	// The IP address.
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The network type. Valid values:
	//
	// 	- **Public**: the Internet
	//
	// 	- **Private**: internal network
	//
	// example:
	//
	// Public
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The version.
	//
	// example:
	//
	// Ipv4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes) SetIpAddress(v string) *SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes {
	s.IpAddress = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes) SetNetworkType(v string) *SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes {
	s.NetworkType = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes) SetVersion(v string) *SearchMultiAccountResourcesResponseBodyResourcesIpAddressAttributes {
	s.Version = &v
	return s
}

type SearchMultiAccountResourcesResponseBodyResourcesTags struct {
	// The key of tag N.
	//
	// example:
	//
	// test_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// example:
	//
	// test_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchMultiAccountResourcesResponseBodyResourcesTags) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesResponseBodyResourcesTags) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesResponseBodyResourcesTags) SetKey(v string) *SearchMultiAccountResourcesResponseBodyResourcesTags {
	s.Key = &v
	return s
}

func (s *SearchMultiAccountResourcesResponseBodyResourcesTags) SetValue(v string) *SearchMultiAccountResourcesResponseBodyResourcesTags {
	s.Value = &v
	return s
}

type SearchMultiAccountResourcesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchMultiAccountResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchMultiAccountResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchMultiAccountResourcesResponse) GoString() string {
	return s.String()
}

func (s *SearchMultiAccountResourcesResponse) SetHeaders(v map[string]*string) *SearchMultiAccountResourcesResponse {
	s.Headers = v
	return s
}

func (s *SearchMultiAccountResourcesResponse) SetStatusCode(v int32) *SearchMultiAccountResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchMultiAccountResourcesResponse) SetBody(v *SearchMultiAccountResourcesResponseBody) *SearchMultiAccountResourcesResponse {
	s.Body = v
	return s
}

type SearchResourcesRequest struct {
	// The filter conditions.
	Filter []*SearchResourcesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// If the total number of entries returned for the current request exceeds the value of the `MaxResults` parameter, the entries are truncated. In this case, you can use the `token` to initiate another request and obtain the remaining entries.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The method that is used to sort the entries returned.
	SortCriterion *SearchResourcesRequestSortCriterion `json:"SortCriterion,omitempty" xml:"SortCriterion,omitempty" type:"Struct"`
}

func (s SearchResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesRequest) GoString() string {
	return s.String()
}

func (s *SearchResourcesRequest) SetFilter(v []*SearchResourcesRequestFilter) *SearchResourcesRequest {
	s.Filter = v
	return s
}

func (s *SearchResourcesRequest) SetMaxResults(v int32) *SearchResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchResourcesRequest) SetNextToken(v string) *SearchResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *SearchResourcesRequest) SetResourceGroupId(v string) *SearchResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *SearchResourcesRequest) SetSortCriterion(v *SearchResourcesRequestSortCriterion) *SearchResourcesRequest {
	s.SortCriterion = v
	return s
}

type SearchResourcesRequestFilter struct {
	// The key of the filter condition. For more information, see `Supported filter parameters`.
	//
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching mode.
	//
	// The value Equals indicates an equal match.
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The values of the filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s SearchResourcesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesRequestFilter) GoString() string {
	return s.String()
}

func (s *SearchResourcesRequestFilter) SetKey(v string) *SearchResourcesRequestFilter {
	s.Key = &v
	return s
}

func (s *SearchResourcesRequestFilter) SetMatchType(v string) *SearchResourcesRequestFilter {
	s.MatchType = &v
	return s
}

func (s *SearchResourcesRequestFilter) SetValue(v []*string) *SearchResourcesRequestFilter {
	s.Value = v
	return s
}

type SearchResourcesRequestSortCriterion struct {
	// The attribute based on which the entries are sorted.
	//
	// The value CreateTime indicates the creation time of resources.
	//
	// example:
	//
	// CreateTime
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The order in which the entries are sorted. Valid values:
	//
	// 	- ASC: The entries are sorted in ascending order. This value is the default value.
	//
	// 	- DESC: The entries are sorted in descending order.
	//
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s SearchResourcesRequestSortCriterion) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesRequestSortCriterion) GoString() string {
	return s.String()
}

func (s *SearchResourcesRequestSortCriterion) SetKey(v string) *SearchResourcesRequestSortCriterion {
	s.Key = &v
	return s
}

func (s *SearchResourcesRequestSortCriterion) SetOrder(v string) *SearchResourcesRequestSortCriterion {
	s.Order = &v
	return s
}

type SearchResourcesResponseBody struct {
	// The filter conditions.
	Filters []*SearchResourcesResponseBodyFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D696E6EF-3A6D-5770-801E-4982081FE4D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resources.
	Resources []*SearchResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s SearchResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchResourcesResponseBody) SetFilters(v []*SearchResourcesResponseBodyFilters) *SearchResourcesResponseBody {
	s.Filters = v
	return s
}

func (s *SearchResourcesResponseBody) SetMaxResults(v int32) *SearchResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *SearchResourcesResponseBody) SetNextToken(v string) *SearchResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchResourcesResponseBody) SetRequestId(v string) *SearchResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchResourcesResponseBody) SetResources(v []*SearchResourcesResponseBodyResources) *SearchResourcesResponseBody {
	s.Resources = v
	return s
}

type SearchResourcesResponseBodyFilters struct {
	// The key of the filter condition.
	//
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching mode.
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The values of the filter condition.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s SearchResourcesResponseBodyFilters) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesResponseBodyFilters) GoString() string {
	return s.String()
}

func (s *SearchResourcesResponseBodyFilters) SetKey(v string) *SearchResourcesResponseBodyFilters {
	s.Key = &v
	return s
}

func (s *SearchResourcesResponseBodyFilters) SetMatchType(v string) *SearchResourcesResponseBodyFilters {
	s.MatchType = &v
	return s
}

func (s *SearchResourcesResponseBodyFilters) SetValues(v []*string) *SearchResourcesResponseBodyFilters {
	s.Values = v
	return s
}

type SearchResourcesResponseBodyResources struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 151266687691****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the resource was created.
	//
	// >  Whether this parameter is returned is determined by the Alibaba Cloud service to which the resource belongs.
	//
	// example:
	//
	// 2021-06-30T09:20:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the resource expires.
	//
	// example:
	//
	// 2021-07-30T09:20:08Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The attributes of the IP address.
	IpAddressAttributes []*SearchResourcesResponseBodyResourcesIpAddressAttributes `json:"IpAddressAttributes,omitempty" xml:"IpAddressAttributes,omitempty" type:"Repeated"`
	// The IP addresses.
	//
	// >  Whether this parameter is returned is determined by the Alibaba Cloud service to which the resource belongs.
	IpAddresses []*string `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// vtb-bp11lbh452fr8940s****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource name.
	//
	// example:
	//
	// group1
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ACS::VPC::RouteTable
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	Tags []*SearchResourcesResponseBodyResourcesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The zone ID.
	//
	// >  Whether this parameter is returned is determined by the Alibaba Cloud service to which the resource belongs.
	//
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s SearchResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *SearchResourcesResponseBodyResources) SetAccountId(v string) *SearchResourcesResponseBodyResources {
	s.AccountId = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetCreateTime(v string) *SearchResourcesResponseBodyResources {
	s.CreateTime = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetExpireTime(v string) *SearchResourcesResponseBodyResources {
	s.ExpireTime = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetIpAddressAttributes(v []*SearchResourcesResponseBodyResourcesIpAddressAttributes) *SearchResourcesResponseBodyResources {
	s.IpAddressAttributes = v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetIpAddresses(v []*string) *SearchResourcesResponseBodyResources {
	s.IpAddresses = v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetRegionId(v string) *SearchResourcesResponseBodyResources {
	s.RegionId = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetResourceGroupId(v string) *SearchResourcesResponseBodyResources {
	s.ResourceGroupId = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetResourceId(v string) *SearchResourcesResponseBodyResources {
	s.ResourceId = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetResourceName(v string) *SearchResourcesResponseBodyResources {
	s.ResourceName = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetResourceType(v string) *SearchResourcesResponseBodyResources {
	s.ResourceType = &v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetTags(v []*SearchResourcesResponseBodyResourcesTags) *SearchResourcesResponseBodyResources {
	s.Tags = v
	return s
}

func (s *SearchResourcesResponseBodyResources) SetZoneId(v string) *SearchResourcesResponseBodyResources {
	s.ZoneId = &v
	return s
}

type SearchResourcesResponseBodyResourcesIpAddressAttributes struct {
	// The IP address.
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The network type. Valid values:
	//
	// 	- **Public**: the Internet
	//
	// 	- **Private**: internal network
	//
	// example:
	//
	// Public
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The version.
	//
	// example:
	//
	// Ipv4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s SearchResourcesResponseBodyResourcesIpAddressAttributes) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesResponseBodyResourcesIpAddressAttributes) GoString() string {
	return s.String()
}

func (s *SearchResourcesResponseBodyResourcesIpAddressAttributes) SetIpAddress(v string) *SearchResourcesResponseBodyResourcesIpAddressAttributes {
	s.IpAddress = &v
	return s
}

func (s *SearchResourcesResponseBodyResourcesIpAddressAttributes) SetNetworkType(v string) *SearchResourcesResponseBodyResourcesIpAddressAttributes {
	s.NetworkType = &v
	return s
}

func (s *SearchResourcesResponseBodyResourcesIpAddressAttributes) SetVersion(v string) *SearchResourcesResponseBodyResourcesIpAddressAttributes {
	s.Version = &v
	return s
}

type SearchResourcesResponseBodyResourcesTags struct {
	// The key of tag N.
	//
	// example:
	//
	// test_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// example:
	//
	// test_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchResourcesResponseBodyResourcesTags) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesResponseBodyResourcesTags) GoString() string {
	return s.String()
}

func (s *SearchResourcesResponseBodyResourcesTags) SetKey(v string) *SearchResourcesResponseBodyResourcesTags {
	s.Key = &v
	return s
}

func (s *SearchResourcesResponseBodyResourcesTags) SetValue(v string) *SearchResourcesResponseBodyResourcesTags {
	s.Value = &v
	return s
}

type SearchResourcesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchResourcesResponse) GoString() string {
	return s.String()
}

func (s *SearchResourcesResponse) SetHeaders(v map[string]*string) *SearchResourcesResponse {
	s.Headers = v
	return s
}

func (s *SearchResourcesResponse) SetStatusCode(v int32) *SearchResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchResourcesResponse) SetBody(v *SearchResourcesResponseBody) *SearchResourcesResponse {
	s.Body = v
	return s
}

type UpdateFilterRequest struct {
	// The configurations of the filter.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "regions": [],
	//
	//   "tagFilters": [
	//
	//     [{ "type": "notContainTagKey", "tagKey": "xxx", "tagValue": "" }],
	//
	//     [{ "tagKey": "xxx", "tagValue": "xxx" }]
	//
	//   ],
	//
	//   "resourceTypes": [
	//
	//     "ACS::ECS::AutoSnapshotPolicy"
	//
	//   ]
	//
	// }
	FilterConfiguration *string `json:"FilterConfiguration,omitempty" xml:"FilterConfiguration,omitempty"`
	// The name of the filter.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS
	FilterName *string `json:"FilterName,omitempty" xml:"FilterName,omitempty"`
}

func (s UpdateFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFilterRequest) GoString() string {
	return s.String()
}

func (s *UpdateFilterRequest) SetFilterConfiguration(v string) *UpdateFilterRequest {
	s.FilterConfiguration = &v
	return s
}

func (s *UpdateFilterRequest) SetFilterName(v string) *UpdateFilterRequest {
	s.FilterName = &v
	return s
}

type UpdateFilterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3C5CDBF6-4DB7-53E9-ADDC-5919E3FACF6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFilterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFilterResponseBody) SetRequestId(v string) *UpdateFilterResponseBody {
	s.RequestId = &v
	return s
}

type UpdateFilterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFilterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFilterResponse) GoString() string {
	return s.String()
}

func (s *UpdateFilterResponse) SetHeaders(v map[string]*string) *UpdateFilterResponse {
	s.Headers = v
	return s
}

func (s *UpdateFilterResponse) SetStatusCode(v int32) *UpdateFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFilterResponse) SetBody(v *UpdateFilterResponseBody) *UpdateFilterResponse {
	s.Body = v
	return s
}

type UpdateSavedQueryRequest struct {
	// The description of the template.
	//
	// The description must be 1 to 256 characters in length.
	//
	// example:
	//
	// Queries all resources on which you have permissions and sorts the resources by resource type and resource ID.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The query statement in the template.
	//
	// example:
	//
	// SELECT 	- FROM resources;
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The name of the template.
	//
	// 	- The name must be 1 to 64 characters in length.
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must be unique.
	//
	// example:
	//
	// Query of All Alibaba Cloud Resources
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sq-GeAck****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s UpdateSavedQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSavedQueryRequest) GoString() string {
	return s.String()
}

func (s *UpdateSavedQueryRequest) SetDescription(v string) *UpdateSavedQueryRequest {
	s.Description = &v
	return s
}

func (s *UpdateSavedQueryRequest) SetExpression(v string) *UpdateSavedQueryRequest {
	s.Expression = &v
	return s
}

func (s *UpdateSavedQueryRequest) SetName(v string) *UpdateSavedQueryRequest {
	s.Name = &v
	return s
}

func (s *UpdateSavedQueryRequest) SetQueryId(v string) *UpdateSavedQueryRequest {
	s.QueryId = &v
	return s
}

type UpdateSavedQueryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D696E6EF-3A6D-5770-801E-4982081FE4D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSavedQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSavedQueryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSavedQueryResponseBody) SetRequestId(v string) *UpdateSavedQueryResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSavedQueryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSavedQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSavedQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSavedQueryResponse) GoString() string {
	return s.String()
}

func (s *UpdateSavedQueryResponse) SetHeaders(v map[string]*string) *UpdateSavedQueryResponse {
	s.Headers = v
	return s
}

func (s *UpdateSavedQueryResponse) SetStatusCode(v int32) *UpdateSavedQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSavedQueryResponse) SetBody(v *UpdateSavedQueryResponseBody) *UpdateSavedQueryResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("resourcecenter"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets a default filter.
//
// @param request - AssociateDefaultFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateDefaultFilterResponse
func (client *Client) AssociateDefaultFilterWithOptions(request *AssociateDefaultFilterRequest, runtime *util.RuntimeOptions) (_result *AssociateDefaultFilterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterName)) {
		query["FilterName"] = request.FilterName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociateDefaultFilter"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociateDefaultFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets a default filter.
//
// @param request - AssociateDefaultFilterRequest
//
// @return AssociateDefaultFilterResponse
func (client *Client) AssociateDefaultFilter(request *AssociateDefaultFilterRequest) (_result *AssociateDefaultFilterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateDefaultFilterResponse{}
	_body, _err := client.AssociateDefaultFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a filter.
//
// @param request - CreateFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFilterResponse
func (client *Client) CreateFilterWithOptions(request *CreateFilterRequest, runtime *util.RuntimeOptions) (_result *CreateFilterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterConfiguration)) {
		query["FilterConfiguration"] = request.FilterConfiguration
	}

	if !tea.BoolValue(util.IsUnset(request.FilterName)) {
		query["FilterName"] = request.FilterName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFilter"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a filter.
//
// @param request - CreateFilterRequest
//
// @return CreateFilterResponse
func (client *Client) CreateFilter(request *CreateFilterRequest) (_result *CreateFilterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFilterResponse{}
	_body, _err := client.CreateFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom query template.
//
// @param request - CreateSavedQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSavedQueryResponse
func (client *Client) CreateSavedQueryWithOptions(request *CreateSavedQueryRequest, runtime *util.RuntimeOptions) (_result *CreateSavedQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Expression)) {
		query["Expression"] = request.Expression
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSavedQuery"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSavedQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom query template.
//
// @param request - CreateSavedQueryRequest
//
// @return CreateSavedQueryResponse
func (client *Client) CreateSavedQuery(request *CreateSavedQueryRequest) (_result *CreateSavedQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSavedQueryResponse{}
	_body, _err := client.CreateSavedQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a filter.
//
// @param request - DeleteFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFilterResponse
func (client *Client) DeleteFilterWithOptions(request *DeleteFilterRequest, runtime *util.RuntimeOptions) (_result *DeleteFilterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterName)) {
		query["FilterName"] = request.FilterName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFilter"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a filter.
//
// @param request - DeleteFilterRequest
//
// @return DeleteFilterResponse
func (client *Client) DeleteFilter(request *DeleteFilterRequest) (_result *DeleteFilterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFilterResponse{}
	_body, _err := client.DeleteFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom query template.
//
// @param request - DeleteSavedQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSavedQueryResponse
func (client *Client) DeleteSavedQueryWithOptions(request *DeleteSavedQueryRequest, runtime *util.RuntimeOptions) (_result *DeleteSavedQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QueryId)) {
		query["QueryId"] = request.QueryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSavedQuery"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSavedQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom query template.
//
// @param request - DeleteSavedQueryRequest
//
// @return DeleteSavedQueryResponse
func (client *Client) DeleteSavedQuery(request *DeleteSavedQueryRequest) (_result *DeleteSavedQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSavedQueryResponse{}
	_body, _err := client.DeleteSavedQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the cross-account resource search feature by using the management account of a resource directory or a delegated administrator account of Resource Center.
//
// @param request - DisableMultiAccountResourceCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableMultiAccountResourceCenterResponse
func (client *Client) DisableMultiAccountResourceCenterWithOptions(runtime *util.RuntimeOptions) (_result *DisableMultiAccountResourceCenterResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DisableMultiAccountResourceCenter"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableMultiAccountResourceCenterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the cross-account resource search feature by using the management account of a resource directory or a delegated administrator account of Resource Center.
//
// @return DisableMultiAccountResourceCenterResponse
func (client *Client) DisableMultiAccountResourceCenter() (_result *DisableMultiAccountResourceCenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableMultiAccountResourceCenterResponse{}
	_body, _err := client.DisableMultiAccountResourceCenterWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deactivates the Resource Center service.
//
// @param request - DisableResourceCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableResourceCenterResponse
func (client *Client) DisableResourceCenterWithOptions(runtime *util.RuntimeOptions) (_result *DisableResourceCenterResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DisableResourceCenter"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableResourceCenterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deactivates the Resource Center service.
//
// @return DisableResourceCenterResponse
func (client *Client) DisableResourceCenter() (_result *DisableResourceCenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableResourceCenterResponse{}
	_body, _err := client.DisableResourceCenterWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels the default filter.
//
// @param request - DisassociateDefaultFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisassociateDefaultFilterResponse
func (client *Client) DisassociateDefaultFilterWithOptions(runtime *util.RuntimeOptions) (_result *DisassociateDefaultFilterResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DisassociateDefaultFilter"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisassociateDefaultFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the default filter.
//
// @return DisassociateDefaultFilterResponse
func (client *Client) DisassociateDefaultFilter() (_result *DisassociateDefaultFilterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisassociateDefaultFilterResponse{}
	_body, _err := client.DisassociateDefaultFilterWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the cross-account resource search feature by using the management account of a resource directory or a delegated administrator account of Resource Center.
//
// Description:
//
// If you have created a resource directory for your enterprise, you can enable the cross-account resource search feature by using the management account of the resource directory or a delegated administrator account of Resource Center to view the resources of members in the resource directory. For more information about a resource directory, see [Resource Directory overview](https://help.aliyun.com/document_detail/200506.html).
//
// @param request - EnableMultiAccountResourceCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableMultiAccountResourceCenterResponse
func (client *Client) EnableMultiAccountResourceCenterWithOptions(runtime *util.RuntimeOptions) (_result *EnableMultiAccountResourceCenterResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("EnableMultiAccountResourceCenter"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableMultiAccountResourceCenterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the cross-account resource search feature by using the management account of a resource directory or a delegated administrator account of Resource Center.
//
// Description:
//
// If you have created a resource directory for your enterprise, you can enable the cross-account resource search feature by using the management account of the resource directory or a delegated administrator account of Resource Center to view the resources of members in the resource directory. For more information about a resource directory, see [Resource Directory overview](https://help.aliyun.com/document_detail/200506.html).
//
// @return EnableMultiAccountResourceCenterResponse
func (client *Client) EnableMultiAccountResourceCenter() (_result *EnableMultiAccountResourceCenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableMultiAccountResourceCenterResponse{}
	_body, _err := client.EnableMultiAccountResourceCenterWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Activates the Resource Center service.
//
// @param request - EnableResourceCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableResourceCenterResponse
func (client *Client) EnableResourceCenterWithOptions(runtime *util.RuntimeOptions) (_result *EnableResourceCenterResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("EnableResourceCenter"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableResourceCenterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates the Resource Center service.
//
// @return EnableResourceCenterResponse
func (client *Client) EnableResourceCenter() (_result *EnableResourceCenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableResourceCenterResponse{}
	_body, _err := client.EnableResourceCenterWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Executes an SQL statement to query resources across accounts.
//
// @param request - ExecuteMultiAccountSQLQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteMultiAccountSQLQueryResponse
func (client *Client) ExecuteMultiAccountSQLQueryWithOptions(request *ExecuteMultiAccountSQLQueryRequest, runtime *util.RuntimeOptions) (_result *ExecuteMultiAccountSQLQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Expression)) {
		query["Expression"] = request.Expression
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteMultiAccountSQLQuery"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteMultiAccountSQLQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes an SQL statement to query resources across accounts.
//
// @param request - ExecuteMultiAccountSQLQueryRequest
//
// @return ExecuteMultiAccountSQLQueryResponse
func (client *Client) ExecuteMultiAccountSQLQuery(request *ExecuteMultiAccountSQLQueryRequest) (_result *ExecuteMultiAccountSQLQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteMultiAccountSQLQueryResponse{}
	_body, _err := client.ExecuteMultiAccountSQLQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Executes an SQL statement to query the resources that can be accessed within the current account.
//
// @param request - ExecuteSQLQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteSQLQueryResponse
func (client *Client) ExecuteSQLQueryWithOptions(request *ExecuteSQLQueryRequest, runtime *util.RuntimeOptions) (_result *ExecuteSQLQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Expression)) {
		query["Expression"] = request.Expression
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteSQLQuery"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteSQLQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes an SQL statement to query the resources that can be accessed within the current account.
//
// @param request - ExecuteSQLQueryRequest
//
// @return ExecuteSQLQueryResponse
func (client *Client) ExecuteSQLQuery(request *ExecuteSQLQueryRequest) (_result *ExecuteSQLQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteSQLQueryResponse{}
	_body, _err := client.ExecuteSQLQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a sample query template.
//
// @param request - GetExampleQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExampleQueryResponse
func (client *Client) GetExampleQueryWithOptions(request *GetExampleQueryRequest, runtime *util.RuntimeOptions) (_result *GetExampleQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QueryId)) {
		query["QueryId"] = request.QueryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetExampleQuery"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExampleQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a sample query template.
//
// @param request - GetExampleQueryRequest
//
// @return GetExampleQueryResponse
func (client *Client) GetExampleQuery(request *GetExampleQueryRequest) (_result *GetExampleQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetExampleQueryResponse{}
	_body, _err := client.GetExampleQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the cross-account resource search feature by using the management account of a resource directory or a delegated administrator account of Resource Center.
//
// @param request - GetMultiAccountResourceCenterServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultiAccountResourceCenterServiceStatusResponse
func (client *Client) GetMultiAccountResourceCenterServiceStatusWithOptions(runtime *util.RuntimeOptions) (_result *GetMultiAccountResourceCenterServiceStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetMultiAccountResourceCenterServiceStatus"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMultiAccountResourceCenterServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the cross-account resource search feature by using the management account of a resource directory or a delegated administrator account of Resource Center.
//
// @return GetMultiAccountResourceCenterServiceStatusResponse
func (client *Client) GetMultiAccountResourceCenterServiceStatus() (_result *GetMultiAccountResourceCenterServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMultiAccountResourceCenterServiceStatusResponse{}
	_body, _err := client.GetMultiAccountResourceCenterServiceStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of a resource within the management account or a member of a resource directory.
//
// @param request - GetMultiAccountResourceConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMultiAccountResourceConfigurationResponse
func (client *Client) GetMultiAccountResourceConfigurationWithOptions(request *GetMultiAccountResourceConfigurationRequest, runtime *util.RuntimeOptions) (_result *GetMultiAccountResourceConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceRegionId)) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMultiAccountResourceConfiguration"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMultiAccountResourceConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of a resource within the management account or a member of a resource directory.
//
// @param request - GetMultiAccountResourceConfigurationRequest
//
// @return GetMultiAccountResourceConfigurationResponse
func (client *Client) GetMultiAccountResourceConfiguration(request *GetMultiAccountResourceConfigurationRequest) (_result *GetMultiAccountResourceConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMultiAccountResourceConfigurationResponse{}
	_body, _err := client.GetMultiAccountResourceConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the Resource Center service.
//
// @param request - GetResourceCenterServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceCenterServiceStatusResponse
func (client *Client) GetResourceCenterServiceStatusWithOptions(runtime *util.RuntimeOptions) (_result *GetResourceCenterServiceStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetResourceCenterServiceStatus"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceCenterServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the Resource Center service.
//
// @return GetResourceCenterServiceStatusResponse
func (client *Client) GetResourceCenterServiceStatus() (_result *GetResourceCenterServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceCenterServiceStatusResponse{}
	_body, _err := client.GetResourceCenterServiceStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of a resource within the current account.
//
// @param request - GetResourceConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceConfigurationResponse
func (client *Client) GetResourceConfigurationWithOptions(request *GetResourceConfigurationRequest, runtime *util.RuntimeOptions) (_result *GetResourceConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceRegionId)) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceConfiguration"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of a resource within the current account.
//
// @param request - GetResourceConfigurationRequest
//
// @return GetResourceConfigurationResponse
func (client *Client) GetResourceConfiguration(request *GetResourceConfigurationRequest) (_result *GetResourceConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceConfigurationResponse{}
	_body, _err := client.GetResourceConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the numbers of resources on which the current account has access permissions.
//
// @param request - GetResourceCountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceCountsResponse
func (client *Client) GetResourceCountsWithOptions(request *GetResourceCountsRequest, runtime *util.RuntimeOptions) (_result *GetResourceCountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.GroupByKey)) {
		query["GroupByKey"] = request.GroupByKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceCounts"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceCountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the numbers of resources on which the current account has access permissions.
//
// @param request - GetResourceCountsRequest
//
// @return GetResourceCountsResponse
func (client *Client) GetResourceCounts(request *GetResourceCountsRequest) (_result *GetResourceCountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceCountsResponse{}
	_body, _err := client.GetResourceCountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a custom query template.
//
// @param request - GetSavedQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSavedQueryResponse
func (client *Client) GetSavedQueryWithOptions(request *GetSavedQueryRequest, runtime *util.RuntimeOptions) (_result *GetSavedQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QueryId)) {
		query["QueryId"] = request.QueryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSavedQuery"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSavedQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a custom query template.
//
// @param request - GetSavedQueryRequest
//
// @return GetSavedQueryResponse
func (client *Client) GetSavedQuery(request *GetSavedQueryRequest) (_result *GetSavedQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSavedQueryResponse{}
	_body, _err := client.GetSavedQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all sample query templates.
//
// @param request - ListExampleQueriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExampleQueriesResponse
func (client *Client) ListExampleQueriesWithOptions(request *ListExampleQueriesRequest, runtime *util.RuntimeOptions) (_result *ListExampleQueriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExampleQueries"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExampleQueriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all sample query templates.
//
// @param request - ListExampleQueriesRequest
//
// @return ListExampleQueriesResponse
func (client *Client) ListExampleQueries(request *ListExampleQueriesRequest) (_result *ListExampleQueriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExampleQueriesResponse{}
	_body, _err := client.ListExampleQueriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of filters.
//
// @param request - ListFiltersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFiltersResponse
func (client *Client) ListFiltersWithOptions(runtime *util.RuntimeOptions) (_result *ListFiltersResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListFilters"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFiltersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of filters.
//
// @return ListFiltersResponse
func (client *Client) ListFilters() (_result *ListFiltersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFiltersResponse{}
	_body, _err := client.ListFiltersWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resource groups within the management account or a member of a resource directory by using the management account of the resource directory or a delegated administrator account of Resource Center.
//
// @param request - ListMultiAccountResourceGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultiAccountResourceGroupsResponse
func (client *Client) ListMultiAccountResourceGroupsWithOptions(request *ListMultiAccountResourceGroupsRequest, runtime *util.RuntimeOptions) (_result *ListMultiAccountResourceGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupIds)) {
		query["ResourceGroupIds"] = request.ResourceGroupIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMultiAccountResourceGroups"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMultiAccountResourceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resource groups within the management account or a member of a resource directory by using the management account of the resource directory or a delegated administrator account of Resource Center.
//
// @param request - ListMultiAccountResourceGroupsRequest
//
// @return ListMultiAccountResourceGroupsResponse
func (client *Client) ListMultiAccountResourceGroups(request *ListMultiAccountResourceGroupsRequest) (_result *ListMultiAccountResourceGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMultiAccountResourceGroupsResponse{}
	_body, _err := client.ListMultiAccountResourceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tag keys of resources within the management account or a member of a resource directory by using the management account of the resource directory or a delegated administrator account of Resource Center.
//
// @param request - ListMultiAccountTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultiAccountTagKeysResponse
func (client *Client) ListMultiAccountTagKeysWithOptions(request *ListMultiAccountTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListMultiAccountTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MatchType)) {
		query["MatchType"] = request.MatchType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMultiAccountTagKeys"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMultiAccountTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tag keys of resources within the management account or a member of a resource directory by using the management account of the resource directory or a delegated administrator account of Resource Center.
//
// @param request - ListMultiAccountTagKeysRequest
//
// @return ListMultiAccountTagKeysResponse
func (client *Client) ListMultiAccountTagKeys(request *ListMultiAccountTagKeysRequest) (_result *ListMultiAccountTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMultiAccountTagKeysResponse{}
	_body, _err := client.ListMultiAccountTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tag values of resources within the management account or a member of a resource directory by using the management account of the resource directory or a delegated administrator account of Resource Center.
//
// @param request - ListMultiAccountTagValuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultiAccountTagValuesResponse
func (client *Client) ListMultiAccountTagValuesWithOptions(request *ListMultiAccountTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListMultiAccountTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MatchType)) {
		query["MatchType"] = request.MatchType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	if !tea.BoolValue(util.IsUnset(request.TagValue)) {
		query["TagValue"] = request.TagValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMultiAccountTagValues"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMultiAccountTagValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tag values of resources within the management account or a member of a resource directory by using the management account of the resource directory or a delegated administrator account of Resource Center.
//
// @param request - ListMultiAccountTagValuesRequest
//
// @return ListMultiAccountTagValuesResponse
func (client *Client) ListMultiAccountTagValues(request *ListMultiAccountTagValuesRequest) (_result *ListMultiAccountTagValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMultiAccountTagValuesResponse{}
	_body, _err := client.ListMultiAccountTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resource types supported by Resource Center.
//
// @param request - ListResourceTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceTypesResponse
func (client *Client) ListResourceTypesWithOptions(request *ListResourceTypesRequest, runtime *util.RuntimeOptions) (_result *ListResourceTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceTypes"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resource types supported by Resource Center.
//
// @param request - ListResourceTypesRequest
//
// @return ListResourceTypesResponse
func (client *Client) ListResourceTypes(request *ListResourceTypesRequest) (_result *ListResourceTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourceTypesResponse{}
	_body, _err := client.ListResourceTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all custom query templates.
//
// @param request - ListSavedQueriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSavedQueriesResponse
func (client *Client) ListSavedQueriesWithOptions(request *ListSavedQueriesRequest, runtime *util.RuntimeOptions) (_result *ListSavedQueriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSavedQueries"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSavedQueriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all custom query templates.
//
// @param request - ListSavedQueriesRequest
//
// @return ListSavedQueriesResponse
func (client *Client) ListSavedQueries(request *ListSavedQueriesRequest) (_result *ListSavedQueriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSavedQueriesResponse{}
	_body, _err := client.ListSavedQueriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tag keys of resources within the current account.
//
// @param request - ListTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MatchType)) {
		query["MatchType"] = request.MatchType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagKeys"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tag keys of resources within the current account.
//
// @param request - ListTagKeysRequest
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tag values of resources within the current account.
//
// @param request - ListTagValuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagValuesResponse
func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MatchType)) {
		query["MatchType"] = request.MatchType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	if !tea.BoolValue(util.IsUnset(request.TagValue)) {
		query["TagValue"] = request.TagValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagValues"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tag values of resources within the current account.
//
// @param request - ListTagValuesRequest
//
// @return ListTagValuesResponse
func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagValuesResponse{}
	_body, _err := client.ListTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Searches for resources within the management account or members of a resource directory.
//
// Description:
//
//   You can use this operation to search for only resources whose types are supported by Resource Center in services that work with Resource Center. For more information about the services and the resource types that are supported by Resource Center, see [Services that work with Resource Center](https://help.aliyun.com/document_detail/477798.html).
//
// 	- Before you use a RAM user or a RAM role to call the operation, you must make sure that the RAM user or RAM role is granted the required permissions. For more information, see [Grant a RAM user the permissions to use Resource Center](https://help.aliyun.com/document_detail/600556.html).
//
// 	- By default, the operation returns a maximum of 20 entries. You can configure the `MaxResults` parameter to specify the maximum number of entries to return.
//
// 	- If the response does not contain the `NextToken` parameter, all entries are returned. Otherwise, more entries exist. If you want to obtain the entries, you can call the operation again to initiate another query request. In the request, set the `NextToken` parameter to the value of `NextToken` in the last response of the operation. If you do not configure the `NextToken` parameter, entries on the first page are returned by default.
//
// 	- You can specify one or more filter conditions to narrow the search scope. For more information about supported filter parameters and matching methods, see the Supported filter parameters section. Multiple filter conditions have logical `AND` relations. Only resources that meet all filter conditions are returned. The values of a filter condition have logical `OR` relations. Resources that meet any value of the filter condition are returned.
//
// 	- You can visit [Sample Code Center](https://api.alibabacloud.com/api-tools/demo/ResourceCenter) to view more sample queries.
//
// @param request - SearchMultiAccountResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchMultiAccountResourcesResponse
func (client *Client) SearchMultiAccountResourcesWithOptions(request *SearchMultiAccountResourcesRequest, runtime *util.RuntimeOptions) (_result *SearchMultiAccountResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.SortCriterion)) {
		query["SortCriterion"] = request.SortCriterion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchMultiAccountResources"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchMultiAccountResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches for resources within the management account or members of a resource directory.
//
// Description:
//
//   You can use this operation to search for only resources whose types are supported by Resource Center in services that work with Resource Center. For more information about the services and the resource types that are supported by Resource Center, see [Services that work with Resource Center](https://help.aliyun.com/document_detail/477798.html).
//
// 	- Before you use a RAM user or a RAM role to call the operation, you must make sure that the RAM user or RAM role is granted the required permissions. For more information, see [Grant a RAM user the permissions to use Resource Center](https://help.aliyun.com/document_detail/600556.html).
//
// 	- By default, the operation returns a maximum of 20 entries. You can configure the `MaxResults` parameter to specify the maximum number of entries to return.
//
// 	- If the response does not contain the `NextToken` parameter, all entries are returned. Otherwise, more entries exist. If you want to obtain the entries, you can call the operation again to initiate another query request. In the request, set the `NextToken` parameter to the value of `NextToken` in the last response of the operation. If you do not configure the `NextToken` parameter, entries on the first page are returned by default.
//
// 	- You can specify one or more filter conditions to narrow the search scope. For more information about supported filter parameters and matching methods, see the Supported filter parameters section. Multiple filter conditions have logical `AND` relations. Only resources that meet all filter conditions are returned. The values of a filter condition have logical `OR` relations. Resources that meet any value of the filter condition are returned.
//
// 	- You can visit [Sample Code Center](https://api.alibabacloud.com/api-tools/demo/ResourceCenter) to view more sample queries.
//
// @param request - SearchMultiAccountResourcesRequest
//
// @return SearchMultiAccountResourcesResponse
func (client *Client) SearchMultiAccountResources(request *SearchMultiAccountResourcesRequest) (_result *SearchMultiAccountResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchMultiAccountResourcesResponse{}
	_body, _err := client.SearchMultiAccountResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Search for resources that you can access within the current account.
//
// Description:
//
//   You can use this operation to search for only resources whose types are supported by Resource Center in services that work with Resource Center. For more information about the services and the resource types that are supported by Resource Center, see [Services that work with Resource Center](https://help.aliyun.com/document_detail/477798.html).
//
// 	- By default, the operation returns a maximum of 20 entries. You can configure the `MaxResults` parameter to specify the maximum number of entries to return.
//
// 	- If the response does not contain the `NextToken` parameter, all entries are returned. Otherwise, more entries exist. If you want to obtain the entries, you can call the operation again to initiate another query request. In the request, set the `NextToken` parameter to the value of `NextToken` in the last response of the operation. If you do not configure the `NextToken` parameter, entries on the first page are returned by default.
//
// 	- You can specify one or more filter conditions to narrow the search scope. For more information about supported filter parameters and matching methods, see the Supported filter parameters section. Multiple filter conditions have logical `AND` relations. Only resources that meet all filter conditions are returned. The values of a filter condition have logical `OR` relations. Resources that meet any value of the filter condition are returned.
//
// 	- You can visit [Sample Code Center](https://api.aliyun.com/api-tools/demo/ResourceCenter) to view more sample queries.
//
// @param request - SearchResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchResourcesResponse
func (client *Client) SearchResourcesWithOptions(request *SearchResourcesRequest, runtime *util.RuntimeOptions) (_result *SearchResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SortCriterion)) {
		query["SortCriterion"] = request.SortCriterion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchResources"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Search for resources that you can access within the current account.
//
// Description:
//
//   You can use this operation to search for only resources whose types are supported by Resource Center in services that work with Resource Center. For more information about the services and the resource types that are supported by Resource Center, see [Services that work with Resource Center](https://help.aliyun.com/document_detail/477798.html).
//
// 	- By default, the operation returns a maximum of 20 entries. You can configure the `MaxResults` parameter to specify the maximum number of entries to return.
//
// 	- If the response does not contain the `NextToken` parameter, all entries are returned. Otherwise, more entries exist. If you want to obtain the entries, you can call the operation again to initiate another query request. In the request, set the `NextToken` parameter to the value of `NextToken` in the last response of the operation. If you do not configure the `NextToken` parameter, entries on the first page are returned by default.
//
// 	- You can specify one or more filter conditions to narrow the search scope. For more information about supported filter parameters and matching methods, see the Supported filter parameters section. Multiple filter conditions have logical `AND` relations. Only resources that meet all filter conditions are returned. The values of a filter condition have logical `OR` relations. Resources that meet any value of the filter condition are returned.
//
// 	- You can visit [Sample Code Center](https://api.aliyun.com/api-tools/demo/ResourceCenter) to view more sample queries.
//
// @param request - SearchResourcesRequest
//
// @return SearchResourcesResponse
func (client *Client) SearchResources(request *SearchResourcesRequest) (_result *SearchResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchResourcesResponse{}
	_body, _err := client.SearchResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a filter.
//
// @param request - UpdateFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFilterResponse
func (client *Client) UpdateFilterWithOptions(request *UpdateFilterRequest, runtime *util.RuntimeOptions) (_result *UpdateFilterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterConfiguration)) {
		query["FilterConfiguration"] = request.FilterConfiguration
	}

	if !tea.BoolValue(util.IsUnset(request.FilterName)) {
		query["FilterName"] = request.FilterName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFilter"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a filter.
//
// @param request - UpdateFilterRequest
//
// @return UpdateFilterResponse
func (client *Client) UpdateFilter(request *UpdateFilterRequest) (_result *UpdateFilterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFilterResponse{}
	_body, _err := client.UpdateFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a custom query template.
//
// @param request - UpdateSavedQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSavedQueryResponse
func (client *Client) UpdateSavedQueryWithOptions(request *UpdateSavedQueryRequest, runtime *util.RuntimeOptions) (_result *UpdateSavedQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Expression)) {
		query["Expression"] = request.Expression
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.QueryId)) {
		query["QueryId"] = request.QueryId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSavedQuery"),
		Version:     tea.String("2022-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSavedQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a custom query template.
//
// @param request - UpdateSavedQueryRequest
//
// @return UpdateSavedQueryResponse
func (client *Client) UpdateSavedQuery(request *UpdateSavedQueryRequest) (_result *UpdateSavedQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSavedQueryResponse{}
	_body, _err := client.UpdateSavedQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
