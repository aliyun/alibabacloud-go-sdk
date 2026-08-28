// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExternalAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteExternalAgentResponseBody
	GetCode() *string
	SetData(v *DeleteExternalAgentResponseBodyData) *DeleteExternalAgentResponseBody
	GetData() *DeleteExternalAgentResponseBodyData
	SetHttpStatusCode(v int32) *DeleteExternalAgentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteExternalAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteExternalAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteExternalAgentResponseBody
	GetSuccess() *bool
}

type DeleteExternalAgentResponseBody struct {
	// The business status code. The value SUCCESS indicates success.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The summary information of the external agent after deletion.
	Data *DeleteExternalAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code. The value 200 indicates success.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The result message of the request.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1a2b3c4d-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteExternalAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteExternalAgentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExternalAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteExternalAgentResponseBody) GetData() *DeleteExternalAgentResponseBodyData {
	return s.Data
}

func (s *DeleteExternalAgentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteExternalAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteExternalAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteExternalAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteExternalAgentResponseBody) SetCode(v string) *DeleteExternalAgentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteExternalAgentResponseBody) SetData(v *DeleteExternalAgentResponseBodyData) *DeleteExternalAgentResponseBody {
	s.Data = v
	return s
}

func (s *DeleteExternalAgentResponseBody) SetHttpStatusCode(v int32) *DeleteExternalAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteExternalAgentResponseBody) SetMessage(v string) *DeleteExternalAgentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteExternalAgentResponseBody) SetRequestId(v string) *DeleteExternalAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExternalAgentResponseBody) SetSuccess(v bool) *DeleteExternalAgentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteExternalAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteExternalAgentResponseBodyData struct {
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
	// The description of the external agent.
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
	// The name of the external agent.
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
	// The status of the external agent. Valid values:
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

func (s DeleteExternalAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteExternalAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteExternalAgentResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *DeleteExternalAgentResponseBodyData) GetCreateMode() *string {
	return s.CreateMode
}

func (s *DeleteExternalAgentResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DeleteExternalAgentResponseBodyData) GetDeployType() *string {
	return s.DeployType
}

func (s *DeleteExternalAgentResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DeleteExternalAgentResponseBodyData) GetEffectiveSpecVersion() *int64 {
	return s.EffectiveSpecVersion
}

func (s *DeleteExternalAgentResponseBodyData) GetLatestSpecVersion() *int64 {
	return s.LatestSpecVersion
}

func (s *DeleteExternalAgentResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DeleteExternalAgentResponseBodyData) GetRuntime() *string {
	return s.Runtime
}

func (s *DeleteExternalAgentResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DeleteExternalAgentResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *DeleteExternalAgentResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteExternalAgentResponseBodyData) SetAgentId(v string) *DeleteExternalAgentResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *DeleteExternalAgentResponseBodyData) SetCreateMode(v string) *DeleteExternalAgentResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *DeleteExternalAgentResponseBodyData) SetCreatedAt(v string) *DeleteExternalAgentResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *DeleteExternalAgentResponseBodyData) SetDeployType(v string) *DeleteExternalAgentResponseBodyData {
	s.DeployType = &v
	return s
}

func (s *DeleteExternalAgentResponseBodyData) SetDescription(v string) *DeleteExternalAgentResponseBodyData {
	s.Description = &v
	return s
}

func (s *DeleteExternalAgentResponseBodyData) SetEffectiveSpecVersion(v int64) *DeleteExternalAgentResponseBodyData {
	s.EffectiveSpecVersion = &v
	return s
}

func (s *DeleteExternalAgentResponseBodyData) SetLatestSpecVersion(v int64) *DeleteExternalAgentResponseBodyData {
	s.LatestSpecVersion = &v
	return s
}

func (s *DeleteExternalAgentResponseBodyData) SetName(v string) *DeleteExternalAgentResponseBodyData {
	s.Name = &v
	return s
}

func (s *DeleteExternalAgentResponseBodyData) SetRuntime(v string) *DeleteExternalAgentResponseBodyData {
	s.Runtime = &v
	return s
}

func (s *DeleteExternalAgentResponseBodyData) SetStatus(v string) *DeleteExternalAgentResponseBodyData {
	s.Status = &v
	return s
}

func (s *DeleteExternalAgentResponseBodyData) SetUpdatedAt(v string) *DeleteExternalAgentResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *DeleteExternalAgentResponseBodyData) SetWorkspaceId(v string) *DeleteExternalAgentResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteExternalAgentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
