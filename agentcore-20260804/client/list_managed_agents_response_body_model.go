// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListManagedAgentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListManagedAgentsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListManagedAgentsResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListManagedAgentsResponseBodyItems) *ListManagedAgentsResponseBody
	GetItems() []*ListManagedAgentsResponseBodyItems
	SetMaxResults(v int32) *ListManagedAgentsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListManagedAgentsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListManagedAgentsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListManagedAgentsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListManagedAgentsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListManagedAgentsResponseBody
	GetTotalCount() *int64
}

type ListManagedAgentsResponseBody struct {
	// The business status code. A value of SUCCESS indicates success.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The list of managed agents.
	//
	// example:
	//
	// [{"agentId":"agent-xxxxxx","workspaceId":"ws-xxxxxx","agentName":"demo-agent","status":"Running","template":{"templateId":"tpl-xxxxxx","templateVersion":"v1.0.0"},"spec":{"replicas":1,"cpu":"2","memory":"4Gi"},"createTime":"2026-01-01T00:00:00Z","updateTime":"2026-01-01T00:00:00Z"}]
	Items []*ListManagedAgentsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of results returned for this request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The message returned for the request.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The token for the next page. An empty value indicates that no more pages are available.
	//
	// example:
	//
	// next-token-1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1a2b3c4d-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListManagedAgentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListManagedAgentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListManagedAgentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListManagedAgentsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListManagedAgentsResponseBody) GetItems() []*ListManagedAgentsResponseBodyItems {
	return s.Items
}

func (s *ListManagedAgentsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListManagedAgentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListManagedAgentsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListManagedAgentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListManagedAgentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListManagedAgentsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListManagedAgentsResponseBody) SetCode(v string) *ListManagedAgentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListManagedAgentsResponseBody) SetHttpStatusCode(v int32) *ListManagedAgentsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListManagedAgentsResponseBody) SetItems(v []*ListManagedAgentsResponseBodyItems) *ListManagedAgentsResponseBody {
	s.Items = v
	return s
}

func (s *ListManagedAgentsResponseBody) SetMaxResults(v int32) *ListManagedAgentsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListManagedAgentsResponseBody) SetMessage(v string) *ListManagedAgentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListManagedAgentsResponseBody) SetNextToken(v string) *ListManagedAgentsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListManagedAgentsResponseBody) SetRequestId(v string) *ListManagedAgentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListManagedAgentsResponseBody) SetSuccess(v bool) *ListManagedAgentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListManagedAgentsResponseBody) SetTotalCount(v int64) *ListManagedAgentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListManagedAgentsResponseBody) Validate() error {
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

type ListManagedAgentsResponseBodyItems struct {
	// The managed agent ID.
	//
	// example:
	//
	// agent-1
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The creation mode.
	//
	// example:
	//
	// Managed
	CreateMode *string `json:"createMode,omitempty" xml:"createMode,omitempty"`
	// The creation time in RFC 3339 format.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The deployment type.
	//
	// example:
	//
	// Managed
	DeployType *string `json:"deployType,omitempty" xml:"deployType,omitempty"`
	// The description of the managed agent.
	//
	// example:
	//
	// An agent for code review
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The effective specification version number.
	//
	// example:
	//
	// 1
	EffectiveSpecVersion *int64 `json:"effectiveSpecVersion,omitempty" xml:"effectiveSpecVersion,omitempty"`
	// The latest specification version number.
	//
	// example:
	//
	// 1
	LatestSpecVersion *int64 `json:"latestSpecVersion,omitempty" xml:"latestSpecVersion,omitempty"`
	// The managed agent name.
	//
	// example:
	//
	// my-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The runtime type.
	//
	// example:
	//
	// Managed
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// The status of the managed agent.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The update time in RFC 3339 format.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-1
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListManagedAgentsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListManagedAgentsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListManagedAgentsResponseBodyItems) GetAgentId() *string {
	return s.AgentId
}

func (s *ListManagedAgentsResponseBodyItems) GetCreateMode() *string {
	return s.CreateMode
}

func (s *ListManagedAgentsResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListManagedAgentsResponseBodyItems) GetDeployType() *string {
	return s.DeployType
}

func (s *ListManagedAgentsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListManagedAgentsResponseBodyItems) GetEffectiveSpecVersion() *int64 {
	return s.EffectiveSpecVersion
}

func (s *ListManagedAgentsResponseBodyItems) GetLatestSpecVersion() *int64 {
	return s.LatestSpecVersion
}

func (s *ListManagedAgentsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListManagedAgentsResponseBodyItems) GetRuntime() *string {
	return s.Runtime
}

func (s *ListManagedAgentsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListManagedAgentsResponseBodyItems) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListManagedAgentsResponseBodyItems) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListManagedAgentsResponseBodyItems) SetAgentId(v string) *ListManagedAgentsResponseBodyItems {
	s.AgentId = &v
	return s
}

func (s *ListManagedAgentsResponseBodyItems) SetCreateMode(v string) *ListManagedAgentsResponseBodyItems {
	s.CreateMode = &v
	return s
}

func (s *ListManagedAgentsResponseBodyItems) SetCreatedAt(v string) *ListManagedAgentsResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListManagedAgentsResponseBodyItems) SetDeployType(v string) *ListManagedAgentsResponseBodyItems {
	s.DeployType = &v
	return s
}

func (s *ListManagedAgentsResponseBodyItems) SetDescription(v string) *ListManagedAgentsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListManagedAgentsResponseBodyItems) SetEffectiveSpecVersion(v int64) *ListManagedAgentsResponseBodyItems {
	s.EffectiveSpecVersion = &v
	return s
}

func (s *ListManagedAgentsResponseBodyItems) SetLatestSpecVersion(v int64) *ListManagedAgentsResponseBodyItems {
	s.LatestSpecVersion = &v
	return s
}

func (s *ListManagedAgentsResponseBodyItems) SetName(v string) *ListManagedAgentsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListManagedAgentsResponseBodyItems) SetRuntime(v string) *ListManagedAgentsResponseBodyItems {
	s.Runtime = &v
	return s
}

func (s *ListManagedAgentsResponseBodyItems) SetStatus(v string) *ListManagedAgentsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListManagedAgentsResponseBodyItems) SetUpdatedAt(v string) *ListManagedAgentsResponseBodyItems {
	s.UpdatedAt = &v
	return s
}

func (s *ListManagedAgentsResponseBodyItems) SetWorkspaceId(v string) *ListManagedAgentsResponseBodyItems {
	s.WorkspaceId = &v
	return s
}

func (s *ListManagedAgentsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
