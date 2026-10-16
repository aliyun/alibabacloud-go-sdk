// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCustomAgentResponseBody
	GetCode() *string
	SetData(v *DeleteCustomAgentResponseBodyData) *DeleteCustomAgentResponseBody
	GetData() *DeleteCustomAgentResponseBodyData
	SetHttpStatusCode(v int32) *DeleteCustomAgentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteCustomAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCustomAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCustomAgentResponseBody
	GetSuccess() *bool
}

type DeleteCustomAgentResponseBody struct {
	// The result code of the request. The value is SUCCESS if the request is successful.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The summary of the custom agent whose deletion is accepted. The runtime resources are subsequently cleaned up in the background.
	Data *DeleteCustomAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code of the request. The value is 200 if the request is successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The processing result message of the request. The value is success if the request is successful.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The unique identifier of the request. You can use this ID to locate and troubleshoot issues.
	//
	// example:
	//
	// 550E8400-E29B-41D4-A716-446655440000
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful. A value of true indicates success. A successful creation, update, or deletion does not mean that the background operation is complete.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteCustomAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomAgentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCustomAgentResponseBody) GetData() *DeleteCustomAgentResponseBodyData {
	return s.Data
}

func (s *DeleteCustomAgentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteCustomAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCustomAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCustomAgentResponseBody) SetCode(v string) *DeleteCustomAgentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCustomAgentResponseBody) SetData(v *DeleteCustomAgentResponseBodyData) *DeleteCustomAgentResponseBody {
	s.Data = v
	return s
}

func (s *DeleteCustomAgentResponseBody) SetHttpStatusCode(v int32) *DeleteCustomAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteCustomAgentResponseBody) SetMessage(v string) *DeleteCustomAgentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomAgentResponseBody) SetRequestId(v string) *DeleteCustomAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomAgentResponseBody) SetSuccess(v bool) *DeleteCustomAgentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCustomAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteCustomAgentResponseBodyData struct {
	// The unique identifier of the custom agent. You can obtain it from the response of CreateCustomAgent or ListCustomAgents.
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
	// The time when the agent was created. The time is in ISO 8601 format. A trailing Z indicates the UTC time zone. Fractional seconds may be included.
	//
	// example:
	//
	// 2026-10-10T06:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The deployment method of the agent. Custom agents are deployed and managed by the platform. The value is MANAGED.
	//
	// example:
	//
	// MANAGED
	DeployType *string `json:"deployType,omitempty" xml:"deployType,omitempty"`
	// The description of the agent\\"s purpose.
	//
	// example:
	//
	// Custom agent that provides conversation services
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
	// - Failed: The agent is in a failed state.
	//
	// - Updating: The agent is being updated.
	//
	// - Deleting: The agent is being deleted.
	//
	// - Deleted: The agent has been deleted.
	//
	// Call GetCustomAgent to retrieve the latest configuration execution progress.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the agent was last updated. The time is in ISO 8601 format. A trailing Z indicates the UTC time zone. Fractional seconds may be included.
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

func (s DeleteCustomAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteCustomAgentResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *DeleteCustomAgentResponseBodyData) GetCreateMode() *string {
	return s.CreateMode
}

func (s *DeleteCustomAgentResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DeleteCustomAgentResponseBodyData) GetDeployType() *string {
	return s.DeployType
}

func (s *DeleteCustomAgentResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DeleteCustomAgentResponseBodyData) GetLatestSpecVersion() *int64 {
	return s.LatestSpecVersion
}

func (s *DeleteCustomAgentResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DeleteCustomAgentResponseBodyData) GetRuntime() *string {
	return s.Runtime
}

func (s *DeleteCustomAgentResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DeleteCustomAgentResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *DeleteCustomAgentResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteCustomAgentResponseBodyData) SetAgentId(v string) *DeleteCustomAgentResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *DeleteCustomAgentResponseBodyData) SetCreateMode(v string) *DeleteCustomAgentResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *DeleteCustomAgentResponseBodyData) SetCreatedAt(v string) *DeleteCustomAgentResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *DeleteCustomAgentResponseBodyData) SetDeployType(v string) *DeleteCustomAgentResponseBodyData {
	s.DeployType = &v
	return s
}

func (s *DeleteCustomAgentResponseBodyData) SetDescription(v string) *DeleteCustomAgentResponseBodyData {
	s.Description = &v
	return s
}

func (s *DeleteCustomAgentResponseBodyData) SetLatestSpecVersion(v int64) *DeleteCustomAgentResponseBodyData {
	s.LatestSpecVersion = &v
	return s
}

func (s *DeleteCustomAgentResponseBodyData) SetName(v string) *DeleteCustomAgentResponseBodyData {
	s.Name = &v
	return s
}

func (s *DeleteCustomAgentResponseBodyData) SetRuntime(v string) *DeleteCustomAgentResponseBodyData {
	s.Runtime = &v
	return s
}

func (s *DeleteCustomAgentResponseBodyData) SetStatus(v string) *DeleteCustomAgentResponseBodyData {
	s.Status = &v
	return s
}

func (s *DeleteCustomAgentResponseBodyData) SetUpdatedAt(v string) *DeleteCustomAgentResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *DeleteCustomAgentResponseBodyData) SetWorkspaceId(v string) *DeleteCustomAgentResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteCustomAgentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
