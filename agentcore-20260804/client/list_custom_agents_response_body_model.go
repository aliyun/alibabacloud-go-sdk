// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomAgentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCustomAgentsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListCustomAgentsResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListCustomAgentsResponseBodyItems) *ListCustomAgentsResponseBody
	GetItems() []*ListCustomAgentsResponseBodyItems
	SetMaxResults(v int32) *ListCustomAgentsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListCustomAgentsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListCustomAgentsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListCustomAgentsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCustomAgentsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListCustomAgentsResponseBody
	GetTotalCount() *int64
}

type ListCustomAgentsResponseBody struct {
	// The result code of the request. The value is SUCCESS if the request is successful.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code of the request. The value is 200 if the request is successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The list of custom agent summaries on the current page. To obtain the complete configuration and the execution results of the latest version, call GetCustomAgent.
	Items []*ListCustomAgentsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of entries returned per page for this query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The description of the processing result for this request. The value is success if the request is successful.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The opaque token used to query the next page. If this value is empty, no more pages are available. Pass this value as is in subsequent requests.
	//
	// example:
	//
	// djE6YWdlbnRjb3JlLmFnZW50Lmxpc3QuY3VzdG9tOjIw
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The unique identifier of the request. You can use this ID to locate and troubleshoot issues.
	//
	// example:
	//
	// 550E8400-E29B-41D4-A716-446655440000
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was processed successfully. A value of true indicates that the request was processed successfully. A successful creation, update, or deletion does not indicate that the background operation is complete.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of custom agents that meet the query conditions in the specified workspace.
	//
	// example:
	//
	// 40
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListCustomAgentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomAgentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCustomAgentsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCustomAgentsResponseBody) GetItems() []*ListCustomAgentsResponseBodyItems {
	return s.Items
}

func (s *ListCustomAgentsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCustomAgentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCustomAgentsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCustomAgentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomAgentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCustomAgentsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCustomAgentsResponseBody) SetCode(v string) *ListCustomAgentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCustomAgentsResponseBody) SetHttpStatusCode(v int32) *ListCustomAgentsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCustomAgentsResponseBody) SetItems(v []*ListCustomAgentsResponseBodyItems) *ListCustomAgentsResponseBody {
	s.Items = v
	return s
}

func (s *ListCustomAgentsResponseBody) SetMaxResults(v int32) *ListCustomAgentsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCustomAgentsResponseBody) SetMessage(v string) *ListCustomAgentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCustomAgentsResponseBody) SetNextToken(v string) *ListCustomAgentsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCustomAgentsResponseBody) SetRequestId(v string) *ListCustomAgentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomAgentsResponseBody) SetSuccess(v bool) *ListCustomAgentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListCustomAgentsResponseBody) SetTotalCount(v int64) *ListCustomAgentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCustomAgentsResponseBody) Validate() error {
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

type ListCustomAgentsResponseBodyItems struct {
	// The unique identifier of the custom agent. You can obtain this value from the response of CreateCustomAgent or ListCustomAgents.
	//
	// example:
	//
	// agent-550e8400e29b41d4a716446655440000
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The creation method of the agent. The value is CUSTOM for custom agents.
	//
	// example:
	//
	// CUSTOM
	CreateMode *string `json:"createMode,omitempty" xml:"createMode,omitempty"`
	// The time when the agent was created. The time is in ISO 8601 format. The trailing Z indicates the UTC time zone. Fractional seconds may be included.
	//
	// example:
	//
	// 2026-10-10T06:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The deployment method of the agent. Custom agents are deployed and managed by the platform. The value is MANAGED for custom agents.
	//
	// example:
	//
	// MANAGED
	DeployType *string `json:"deployType,omitempty" xml:"deployType,omitempty"`
	// The description of the agent.
	//
	// example:
	//
	// Custom agent that provides conversational services
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The latest configuration version number. Updates with the same configuration may retain the original version number. This field does not indicate that the configuration has taken effect.
	//
	// example:
	//
	// 1
	LatestSpecVersion *int64 `json:"latestSpecVersion,omitempty" xml:"latestSpecVersion,omitempty"`
	// The name of the custom agent.
	//
	// example:
	//
	// my-custom-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The list of invocation protocol configurations for the agent.
	Protocol []*ListCustomAgentsResponseBodyItemsProtocol `json:"protocol,omitempty" xml:"protocol,omitempty" type:"Repeated"`
	// The compute specifications of the agent. Valid values: INTERACTIVE, STANDARD, and HEAVY.
	//
	// example:
	//
	// STANDARD
	Runtime *string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// The overall status of the agent. Valid values:
	//
	// - Creating: The agent is being created.
	//
	// - Running: The agent is running.
	//
	// - Failed: The agent failed.
	//
	// - Updating: The agent is being updated.
	//
	// - Deleting: The agent is being deleted.
	//
	// - Deleted: The agent is deleted.
	//
	// To check the execution progress of the latest configuration, call GetCustomAgent.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the agent was last updated. The time is in ISO 8601 format. The trailing Z indicates the UTC time zone. Fractional seconds may be included.
	//
	// example:
	//
	// 2026-10-10T06:05:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The ID of the workspace to which the custom agent belongs.
	//
	// example:
	//
	// ws-a1b2c3d4e5f6
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListCustomAgentsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListCustomAgentsResponseBodyItems) GetAgentId() *string {
	return s.AgentId
}

func (s *ListCustomAgentsResponseBodyItems) GetCreateMode() *string {
	return s.CreateMode
}

func (s *ListCustomAgentsResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListCustomAgentsResponseBodyItems) GetDeployType() *string {
	return s.DeployType
}

func (s *ListCustomAgentsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListCustomAgentsResponseBodyItems) GetLatestSpecVersion() *int64 {
	return s.LatestSpecVersion
}

func (s *ListCustomAgentsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListCustomAgentsResponseBodyItems) GetProtocol() []*ListCustomAgentsResponseBodyItemsProtocol {
	return s.Protocol
}

func (s *ListCustomAgentsResponseBodyItems) GetRuntime() *string {
	return s.Runtime
}

func (s *ListCustomAgentsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListCustomAgentsResponseBodyItems) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListCustomAgentsResponseBodyItems) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListCustomAgentsResponseBodyItems) SetAgentId(v string) *ListCustomAgentsResponseBodyItems {
	s.AgentId = &v
	return s
}

func (s *ListCustomAgentsResponseBodyItems) SetCreateMode(v string) *ListCustomAgentsResponseBodyItems {
	s.CreateMode = &v
	return s
}

func (s *ListCustomAgentsResponseBodyItems) SetCreatedAt(v string) *ListCustomAgentsResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListCustomAgentsResponseBodyItems) SetDeployType(v string) *ListCustomAgentsResponseBodyItems {
	s.DeployType = &v
	return s
}

func (s *ListCustomAgentsResponseBodyItems) SetDescription(v string) *ListCustomAgentsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListCustomAgentsResponseBodyItems) SetLatestSpecVersion(v int64) *ListCustomAgentsResponseBodyItems {
	s.LatestSpecVersion = &v
	return s
}

func (s *ListCustomAgentsResponseBodyItems) SetName(v string) *ListCustomAgentsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListCustomAgentsResponseBodyItems) SetProtocol(v []*ListCustomAgentsResponseBodyItemsProtocol) *ListCustomAgentsResponseBodyItems {
	s.Protocol = v
	return s
}

func (s *ListCustomAgentsResponseBodyItems) SetRuntime(v string) *ListCustomAgentsResponseBodyItems {
	s.Runtime = &v
	return s
}

func (s *ListCustomAgentsResponseBodyItems) SetStatus(v string) *ListCustomAgentsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListCustomAgentsResponseBodyItems) SetUpdatedAt(v string) *ListCustomAgentsResponseBodyItems {
	s.UpdatedAt = &v
	return s
}

func (s *ListCustomAgentsResponseBodyItems) SetWorkspaceId(v string) *ListCustomAgentsResponseBodyItems {
	s.WorkspaceId = &v
	return s
}

func (s *ListCustomAgentsResponseBodyItems) Validate() error {
	if s.Protocol != nil {
		for _, item := range s.Protocol {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomAgentsResponseBodyItemsProtocol struct {
	// The request header configuration of the current protocol. The key is the header name and the value is a string.
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	// The list of API paths corresponding to the current protocol.
	Paths []*ListCustomAgentsResponseBodyItemsProtocolPaths `json:"paths,omitempty" xml:"paths,omitempty" type:"Repeated"`
	// The invocation protocol type. Valid values:
	//
	// - agui: AG-UI protocol.
	//
	// - openai-chat-completion: OpenAI Chat Completions compatible protocol.
	//
	// - custom: Custom protocol.
	//
	// This parameter is required.
	//
	// example:
	//
	// agui
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListCustomAgentsResponseBodyItemsProtocol) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentsResponseBodyItemsProtocol) GoString() string {
	return s.String()
}

func (s *ListCustomAgentsResponseBodyItemsProtocol) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomAgentsResponseBodyItemsProtocol) GetPaths() []*ListCustomAgentsResponseBodyItemsProtocolPaths {
	return s.Paths
}

func (s *ListCustomAgentsResponseBodyItemsProtocol) GetType() *string {
	return s.Type
}

func (s *ListCustomAgentsResponseBodyItemsProtocol) SetHeaders(v map[string]*string) *ListCustomAgentsResponseBodyItemsProtocol {
	s.Headers = v
	return s
}

func (s *ListCustomAgentsResponseBodyItemsProtocol) SetPaths(v []*ListCustomAgentsResponseBodyItemsProtocolPaths) *ListCustomAgentsResponseBodyItemsProtocol {
	s.Paths = v
	return s
}

func (s *ListCustomAgentsResponseBodyItemsProtocol) SetType(v string) *ListCustomAgentsResponseBodyItemsProtocol {
	s.Type = &v
	return s
}

func (s *ListCustomAgentsResponseBodyItemsProtocol) Validate() error {
	if s.Paths != nil {
		for _, item := range s.Paths {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomAgentsResponseBodyItemsProtocolPaths struct {
	// The description of the current API path.
	//
	// example:
	//
	// AG-UI conversation API
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The API path corresponding to the current protocol.
	//
	// This parameter is required.
	//
	// example:
	//
	// /agui
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListCustomAgentsResponseBodyItemsProtocolPaths) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentsResponseBodyItemsProtocolPaths) GoString() string {
	return s.String()
}

func (s *ListCustomAgentsResponseBodyItemsProtocolPaths) GetDescription() *string {
	return s.Description
}

func (s *ListCustomAgentsResponseBodyItemsProtocolPaths) GetName() *string {
	return s.Name
}

func (s *ListCustomAgentsResponseBodyItemsProtocolPaths) SetDescription(v string) *ListCustomAgentsResponseBodyItemsProtocolPaths {
	s.Description = &v
	return s
}

func (s *ListCustomAgentsResponseBodyItemsProtocolPaths) SetName(v string) *ListCustomAgentsResponseBodyItemsProtocolPaths {
	s.Name = &v
	return s
}

func (s *ListCustomAgentsResponseBodyItemsProtocolPaths) Validate() error {
	return dara.Validate(s)
}
