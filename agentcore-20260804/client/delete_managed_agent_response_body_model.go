// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteManagedAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteManagedAgentResponseBody
	GetCode() *string
	SetData(v *DeleteManagedAgentResponseBodyData) *DeleteManagedAgentResponseBody
	GetData() *DeleteManagedAgentResponseBodyData
	SetHttpStatusCode(v int32) *DeleteManagedAgentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteManagedAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteManagedAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteManagedAgentResponseBody
	GetSuccess() *bool
}

type DeleteManagedAgentResponseBody struct {
	// The business status code. The value SUCCESS indicates success.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The summary information of the managed agent after deletion.
	Data *DeleteManagedAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteManagedAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteManagedAgentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteManagedAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteManagedAgentResponseBody) GetData() *DeleteManagedAgentResponseBodyData {
	return s.Data
}

func (s *DeleteManagedAgentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteManagedAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteManagedAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteManagedAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteManagedAgentResponseBody) SetCode(v string) *DeleteManagedAgentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteManagedAgentResponseBody) SetData(v *DeleteManagedAgentResponseBodyData) *DeleteManagedAgentResponseBody {
	s.Data = v
	return s
}

func (s *DeleteManagedAgentResponseBody) SetHttpStatusCode(v int32) *DeleteManagedAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteManagedAgentResponseBody) SetMessage(v string) *DeleteManagedAgentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteManagedAgentResponseBody) SetRequestId(v string) *DeleteManagedAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteManagedAgentResponseBody) SetSuccess(v bool) *DeleteManagedAgentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteManagedAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteManagedAgentResponseBodyData struct {
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
	// The name of the managed agent.
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

func (s DeleteManagedAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteManagedAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteManagedAgentResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *DeleteManagedAgentResponseBodyData) GetCreateMode() *string {
	return s.CreateMode
}

func (s *DeleteManagedAgentResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DeleteManagedAgentResponseBodyData) GetDeployType() *string {
	return s.DeployType
}

func (s *DeleteManagedAgentResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DeleteManagedAgentResponseBodyData) GetEffectiveSpecVersion() *int64 {
	return s.EffectiveSpecVersion
}

func (s *DeleteManagedAgentResponseBodyData) GetLatestSpecVersion() *int64 {
	return s.LatestSpecVersion
}

func (s *DeleteManagedAgentResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DeleteManagedAgentResponseBodyData) GetRuntime() *string {
	return s.Runtime
}

func (s *DeleteManagedAgentResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DeleteManagedAgentResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *DeleteManagedAgentResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteManagedAgentResponseBodyData) SetAgentId(v string) *DeleteManagedAgentResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *DeleteManagedAgentResponseBodyData) SetCreateMode(v string) *DeleteManagedAgentResponseBodyData {
	s.CreateMode = &v
	return s
}

func (s *DeleteManagedAgentResponseBodyData) SetCreatedAt(v string) *DeleteManagedAgentResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *DeleteManagedAgentResponseBodyData) SetDeployType(v string) *DeleteManagedAgentResponseBodyData {
	s.DeployType = &v
	return s
}

func (s *DeleteManagedAgentResponseBodyData) SetDescription(v string) *DeleteManagedAgentResponseBodyData {
	s.Description = &v
	return s
}

func (s *DeleteManagedAgentResponseBodyData) SetEffectiveSpecVersion(v int64) *DeleteManagedAgentResponseBodyData {
	s.EffectiveSpecVersion = &v
	return s
}

func (s *DeleteManagedAgentResponseBodyData) SetLatestSpecVersion(v int64) *DeleteManagedAgentResponseBodyData {
	s.LatestSpecVersion = &v
	return s
}

func (s *DeleteManagedAgentResponseBodyData) SetName(v string) *DeleteManagedAgentResponseBodyData {
	s.Name = &v
	return s
}

func (s *DeleteManagedAgentResponseBodyData) SetRuntime(v string) *DeleteManagedAgentResponseBodyData {
	s.Runtime = &v
	return s
}

func (s *DeleteManagedAgentResponseBodyData) SetStatus(v string) *DeleteManagedAgentResponseBodyData {
	s.Status = &v
	return s
}

func (s *DeleteManagedAgentResponseBodyData) SetUpdatedAt(v string) *DeleteManagedAgentResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *DeleteManagedAgentResponseBodyData) SetWorkspaceId(v string) *DeleteManagedAgentResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteManagedAgentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
