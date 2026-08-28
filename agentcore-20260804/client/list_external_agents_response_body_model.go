// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExternalAgentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListExternalAgentsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListExternalAgentsResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListExternalAgentsResponseBodyItems) *ListExternalAgentsResponseBody
	GetItems() []*ListExternalAgentsResponseBodyItems
	SetMaxResults(v int32) *ListExternalAgentsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListExternalAgentsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListExternalAgentsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListExternalAgentsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListExternalAgentsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListExternalAgentsResponseBody
	GetTotalCount() *int64
}

type ListExternalAgentsResponseBody struct {
	// The business status code. The value SUCCESS indicates success.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code. The value 200 indicates success.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The list of external agents.
	//
	// example:
	//
	// [{"agentId":"agent-1","workspaceId":"ws-1","name":"my-external-agent","description":"A code review agent running in the user environment","status":"Running","latestSpecVersion":1,"effectiveSpecVersion":1,"createMode":"CUSTOM","runtime":"qwenpaw","deployType":"SELF_HOSTED","createdAt":"2026-01-01T00:00:00Z","updatedAt":"2026-01-01T00:00:00Z"}]
	Items []*ListExternalAgentsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of entries returned per page for this request.
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
	// The token for the next page. An empty value indicates the last page.
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
	// The total number of records that match the query conditions.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListExternalAgentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExternalAgentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExternalAgentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListExternalAgentsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListExternalAgentsResponseBody) GetItems() []*ListExternalAgentsResponseBodyItems {
	return s.Items
}

func (s *ListExternalAgentsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExternalAgentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListExternalAgentsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExternalAgentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExternalAgentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListExternalAgentsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListExternalAgentsResponseBody) SetCode(v string) *ListExternalAgentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListExternalAgentsResponseBody) SetHttpStatusCode(v int32) *ListExternalAgentsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListExternalAgentsResponseBody) SetItems(v []*ListExternalAgentsResponseBodyItems) *ListExternalAgentsResponseBody {
	s.Items = v
	return s
}

func (s *ListExternalAgentsResponseBody) SetMaxResults(v int32) *ListExternalAgentsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListExternalAgentsResponseBody) SetMessage(v string) *ListExternalAgentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListExternalAgentsResponseBody) SetNextToken(v string) *ListExternalAgentsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListExternalAgentsResponseBody) SetRequestId(v string) *ListExternalAgentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExternalAgentsResponseBody) SetSuccess(v bool) *ListExternalAgentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListExternalAgentsResponseBody) SetTotalCount(v int64) *ListExternalAgentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListExternalAgentsResponseBody) Validate() error {
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

type ListExternalAgentsResponseBodyItems struct {
	// The external agent ID.
	//
	// example:
	//
	// agent-1
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The creation mode.
	//
	// example:
	//
	// CUSTOM
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
	// SELF_HOSTED
	DeployType *string `json:"deployType,omitempty" xml:"deployType,omitempty"`
	// The external agent description.
	//
	// example:
	//
	// A code review agent running in the user environment
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The currently effective specification version number.
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
	// The external agent name.
	//
	// example:
	//
	// my-external-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The runtime type reported by the external agent.
	//
	// example:
	//
	// qwenpaw
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// The external agent status. Valid values:
	//
	// - Creating: The agent is being created.
	//
	// - Running: The agent is running.
	//
	// - Failed: The agent has failed.
	//
	// - Updating: The agent is being updated.
	//
	// - Deleting: The agent is being deleted.
	//
	// - Deleted: The agent has been deleted.
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

func (s ListExternalAgentsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListExternalAgentsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListExternalAgentsResponseBodyItems) GetAgentId() *string {
	return s.AgentId
}

func (s *ListExternalAgentsResponseBodyItems) GetCreateMode() *string {
	return s.CreateMode
}

func (s *ListExternalAgentsResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListExternalAgentsResponseBodyItems) GetDeployType() *string {
	return s.DeployType
}

func (s *ListExternalAgentsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListExternalAgentsResponseBodyItems) GetEffectiveSpecVersion() *int64 {
	return s.EffectiveSpecVersion
}

func (s *ListExternalAgentsResponseBodyItems) GetLatestSpecVersion() *int64 {
	return s.LatestSpecVersion
}

func (s *ListExternalAgentsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListExternalAgentsResponseBodyItems) GetRuntime() *string {
	return s.Runtime
}

func (s *ListExternalAgentsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListExternalAgentsResponseBodyItems) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListExternalAgentsResponseBodyItems) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListExternalAgentsResponseBodyItems) SetAgentId(v string) *ListExternalAgentsResponseBodyItems {
	s.AgentId = &v
	return s
}

func (s *ListExternalAgentsResponseBodyItems) SetCreateMode(v string) *ListExternalAgentsResponseBodyItems {
	s.CreateMode = &v
	return s
}

func (s *ListExternalAgentsResponseBodyItems) SetCreatedAt(v string) *ListExternalAgentsResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListExternalAgentsResponseBodyItems) SetDeployType(v string) *ListExternalAgentsResponseBodyItems {
	s.DeployType = &v
	return s
}

func (s *ListExternalAgentsResponseBodyItems) SetDescription(v string) *ListExternalAgentsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListExternalAgentsResponseBodyItems) SetEffectiveSpecVersion(v int64) *ListExternalAgentsResponseBodyItems {
	s.EffectiveSpecVersion = &v
	return s
}

func (s *ListExternalAgentsResponseBodyItems) SetLatestSpecVersion(v int64) *ListExternalAgentsResponseBodyItems {
	s.LatestSpecVersion = &v
	return s
}

func (s *ListExternalAgentsResponseBodyItems) SetName(v string) *ListExternalAgentsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListExternalAgentsResponseBodyItems) SetRuntime(v string) *ListExternalAgentsResponseBodyItems {
	s.Runtime = &v
	return s
}

func (s *ListExternalAgentsResponseBodyItems) SetStatus(v string) *ListExternalAgentsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListExternalAgentsResponseBodyItems) SetUpdatedAt(v string) *ListExternalAgentsResponseBodyItems {
	s.UpdatedAt = &v
	return s
}

func (s *ListExternalAgentsResponseBodyItems) SetWorkspaceId(v string) *ListExternalAgentsResponseBodyItems {
	s.WorkspaceId = &v
	return s
}

func (s *ListExternalAgentsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
