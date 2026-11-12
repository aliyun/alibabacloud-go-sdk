// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAgentVersionResponseBody
	GetCode() *string
	SetData(v *GetAgentVersionResponseBodyData) *GetAgentVersionResponseBody
	GetData() *GetAgentVersionResponseBodyData
	SetHttpStatusCode(v int32) *GetAgentVersionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAgentVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAgentVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAgentVersionResponseBody
	GetSuccess() *bool
}

type GetAgentVersionResponseBody struct {
	// The business status code. The value is SUCCESS if the request is successful, or an error code if the request fails.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The details of the fixed agent version. The configuration field contains the agent configuration at the time the version was saved.
	Data *GetAgentVersionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message. The value is success if the request is successful, or an error message if the request fails.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID, which is used to locate and troubleshoot the request.
	//
	// example:
	//
	// 5E8F3C2A-1B4D-4F6E-9A7C-0D2B8E6F4A1C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetAgentVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgentVersionResponseBody) GetData() *GetAgentVersionResponseBodyData {
	return s.Data
}

func (s *GetAgentVersionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAgentVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgentVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAgentVersionResponseBody) SetCode(v string) *GetAgentVersionResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentVersionResponseBody) SetData(v *GetAgentVersionResponseBodyData) *GetAgentVersionResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentVersionResponseBody) SetHttpStatusCode(v int32) *GetAgentVersionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAgentVersionResponseBody) SetMessage(v string) *GetAgentVersionResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentVersionResponseBody) SetRequestId(v string) *GetAgentVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentVersionResponseBody) SetSuccess(v bool) *GetAgentVersionResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentVersionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentVersionResponseBodyData struct {
	// The ID of the managed or custom agent to which the version belongs.
	//
	// example:
	//
	// agent-0123456789abcdef0123456789abcdef
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The system-generated version number that increments for the agent, such as v1 and v2. This is used when querying or deleting versions, and when server endpoints reference the version.
	//
	// example:
	//
	// v1
	AgentVersion *string `json:"agentVersion,omitempty" xml:"agentVersion,omitempty"`
	// Indicates only whether the version template and artifact archive are complete, and does not represent the deployment or running status. Valid values: READY indicates that the archive is complete. INVALID indicates that the archive is missing or inconsistent.
	//
	// example:
	//
	// READY
	ArtifactStatus *string `json:"artifactStatus,omitempty" xml:"artifactStatus,omitempty"`
	// For managed agents, the structure is consistent with the data structure of GetManagedAgent. For custom agents, the structure is consistent with the data structure of GetCustomAgent. The data masking rules are the same.
	//
	// example:
	//
	// {"agentId":"agent-0123456789abcdef0123456789abcdef","workspaceId":"ws-0123456789abcdef01234","name":"customer-service","description":"Customer service agent","status":"Running","latestSpecVersion":2,"instruction":"You are a customer service assistant."}
	Configuration map[string]interface{} `json:"configuration,omitempty" xml:"configuration,omitempty"`
	// The time when the version was saved, in ISO 8601 UTC format.
	//
	// example:
	//
	// 2026-08-12T01:02:03Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The UID of the Alibaba Cloud account of the caller who saved the version. If the API is called by a RAM user, this is the UID of the Alibaba Cloud account to which the RAM user belongs.
	//
	// example:
	//
	// 1234567890123456
	CreatedBy *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// The version description provided when the version was saved. It can be up to 1,024 characters in length. If no description was provided when the version was saved, this field is empty.
	//
	// example:
	//
	// Stable version for the production environment
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The effective release sequence number of the agent when the version was saved, which is the effectiveSpecVersion at that time. This value does not change after the version is saved.
	//
	// example:
	//
	// 2
	TargetSpecVersion *int64 `json:"targetSpecVersion,omitempty" xml:"targetSpecVersion,omitempty"`
	// The display name specified by the user when the version was saved. For earlier saved versions where no name was recorded, the agentVersion is returned.
	//
	// example:
	//
	// Production stable version V1.0
	VersionName *string `json:"versionName,omitempty" xml:"versionName,omitempty"`
	// The ID of the AgentCore workspace where the agent that owns this version resides.
	//
	// example:
	//
	// ws-0123456789abcdef01234
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetAgentVersionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentVersionResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *GetAgentVersionResponseBodyData) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *GetAgentVersionResponseBodyData) GetArtifactStatus() *string {
	return s.ArtifactStatus
}

func (s *GetAgentVersionResponseBodyData) GetConfiguration() map[string]interface{} {
	return s.Configuration
}

func (s *GetAgentVersionResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetAgentVersionResponseBodyData) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *GetAgentVersionResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetAgentVersionResponseBodyData) GetTargetSpecVersion() *int64 {
	return s.TargetSpecVersion
}

func (s *GetAgentVersionResponseBodyData) GetVersionName() *string {
	return s.VersionName
}

func (s *GetAgentVersionResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetAgentVersionResponseBodyData) SetAgentId(v string) *GetAgentVersionResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *GetAgentVersionResponseBodyData) SetAgentVersion(v string) *GetAgentVersionResponseBodyData {
	s.AgentVersion = &v
	return s
}

func (s *GetAgentVersionResponseBodyData) SetArtifactStatus(v string) *GetAgentVersionResponseBodyData {
	s.ArtifactStatus = &v
	return s
}

func (s *GetAgentVersionResponseBodyData) SetConfiguration(v map[string]interface{}) *GetAgentVersionResponseBodyData {
	s.Configuration = v
	return s
}

func (s *GetAgentVersionResponseBodyData) SetCreatedAt(v string) *GetAgentVersionResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetAgentVersionResponseBodyData) SetCreatedBy(v string) *GetAgentVersionResponseBodyData {
	s.CreatedBy = &v
	return s
}

func (s *GetAgentVersionResponseBodyData) SetDescription(v string) *GetAgentVersionResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetAgentVersionResponseBodyData) SetTargetSpecVersion(v int64) *GetAgentVersionResponseBodyData {
	s.TargetSpecVersion = &v
	return s
}

func (s *GetAgentVersionResponseBodyData) SetVersionName(v string) *GetAgentVersionResponseBodyData {
	s.VersionName = &v
	return s
}

func (s *GetAgentVersionResponseBodyData) SetWorkspaceId(v string) *GetAgentVersionResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetAgentVersionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
