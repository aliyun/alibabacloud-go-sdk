// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateAgentResponseBodyData) *CreateAgentResponseBody
	GetData() *CreateAgentResponseBodyData
	SetErrorCode(v string) *CreateAgentResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateAgentResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAgentResponseBody
	GetSuccess() *bool
}

type CreateAgentResponseBody struct {
	// The agent information and the automatically issued API key returned after the agent is created.
	Data *CreateAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The status code of the request result. A value of success indicates success. A specific error code is returned upon failure.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the request fails. This parameter is empty when the request succeeds.
	//
	// example:
	//
	// agentName must not be blank
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The unique ID of the request. You can use this ID for troubleshooting and tracing.
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentResponseBody) GetData() *CreateAgentResponseBodyData {
	return s.Data
}

func (s *CreateAgentResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateAgentResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAgentResponseBody) SetData(v *CreateAgentResponseBodyData) *CreateAgentResponseBody {
	s.Data = v
	return s
}

func (s *CreateAgentResponseBody) SetErrorCode(v string) *CreateAgentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateAgentResponseBody) SetErrorMessage(v string) *CreateAgentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateAgentResponseBody) SetRequestId(v string) *CreateAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgentResponseBody) SetSuccess(v bool) *CreateAgentResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentResponseBodyData struct {
	// The globally unique ID of the agent.
	//
	// example:
	//
	// agt-1a2b3c4d5e6f
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The agent name.
	//
	// example:
	//
	// order-analysis-agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The permission inheritance type of the agent. Valid values: HUMAN_BOUND (inherits user permissions), PERMISSION_NARROW (narrows permissions), STANDALONE (operates as an independent identity principal without inheriting permissions from other principals).
	//
	// example:
	//
	// HUMAN_BOUND
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// The automatically issued API key for the new agent. The plaintext secret is returned only once in this response.
	ApiKey *CreateAgentResponseBodyDataApiKey `json:"ApiKey,omitempty" xml:"ApiKey,omitempty" type:"Struct"`
	// The time when the agent was created. The value is a time string in RFC 3339 format.
	//
	// example:
	//
	// 2025-12-11T14:04:32Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The creation method of the agent. Valid values: manual (manually created in the console), auto (automatic creation by the system). Agents created by this operation are always manual.
	//
	// example:
	//
	// manual
	CreationType *string `json:"CreationType,omitempty" xml:"CreationType,omitempty"`
	// The description of the agent.
	//
	// example:
	//
	// An agent for querying and analyzing order data
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The user ID of the agent owner, which is the current user who initiated the creation request.
	//
	// example:
	//
	// usr-1a2b3c4d
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The status of the agent. Valid values: active (enabled), disabled (disabled), deleted (deleted). A newly created agent is always active.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAgentResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateAgentResponseBodyData) GetAgentName() *string {
	return s.AgentName
}

func (s *CreateAgentResponseBodyData) GetAgentType() *string {
	return s.AgentType
}

func (s *CreateAgentResponseBodyData) GetApiKey() *CreateAgentResponseBodyDataApiKey {
	return s.ApiKey
}

func (s *CreateAgentResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateAgentResponseBodyData) GetCreationType() *string {
	return s.CreationType
}

func (s *CreateAgentResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateAgentResponseBodyData) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateAgentResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateAgentResponseBodyData) SetAgentId(v string) *CreateAgentResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *CreateAgentResponseBodyData) SetAgentName(v string) *CreateAgentResponseBodyData {
	s.AgentName = &v
	return s
}

func (s *CreateAgentResponseBodyData) SetAgentType(v string) *CreateAgentResponseBodyData {
	s.AgentType = &v
	return s
}

func (s *CreateAgentResponseBodyData) SetApiKey(v *CreateAgentResponseBodyDataApiKey) *CreateAgentResponseBodyData {
	s.ApiKey = v
	return s
}

func (s *CreateAgentResponseBodyData) SetCreatedAt(v string) *CreateAgentResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateAgentResponseBodyData) SetCreationType(v string) *CreateAgentResponseBodyData {
	s.CreationType = &v
	return s
}

func (s *CreateAgentResponseBodyData) SetDescription(v string) *CreateAgentResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateAgentResponseBodyData) SetOwnerId(v string) *CreateAgentResponseBodyData {
	s.OwnerId = &v
	return s
}

func (s *CreateAgentResponseBodyData) SetStatus(v string) *CreateAgentResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateAgentResponseBodyData) Validate() error {
	if s.ApiKey != nil {
		if err := s.ApiKey.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentResponseBodyDataApiKey struct {
	// The ID of the agent to which the API key belongs.
	//
	// example:
	//
	// agt-1a2b3c4d5e6f
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The name of the agent to which the API key belongs.
	//
	// example:
	//
	// order-analysis-agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The permission inheritance type of the agent to which the API key belongs.
	//
	// example:
	//
	// HUMAN_BOUND
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// The time when the API key was created. The value is a time string in RFC 3339 format.
	//
	// example:
	//
	// 2025-12-11T14:04:32Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The user ID of the user who created the API key.
	//
	// example:
	//
	// usr-1a2b3c4d
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The name of the user who created the API key.
	//
	// example:
	//
	// John Doe
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The description of the API key.
	//
	// example:
	//
	// Access Token automatically issued when the Agent is created
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The expiration time of the API key. The value is a time string in RFC 3339 format.
	//
	// example:
	//
	// 2026-12-11T14:04:32Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The primary key ID of the API key.
	//
	// example:
	//
	// 1001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the API key has been revoked.
	IsRevoked *bool `json:"IsRevoked,omitempty" xml:"IsRevoked,omitempty"`
	// The non-sensitive visible prefix of the API key plaintext, used to identify the credential. The plaintext secret is not returned again.
	//
	// example:
	//
	// dms_sk_1a2b
	KeyPrefix *string `json:"KeyPrefix,omitempty" xml:"KeyPrefix,omitempty"`
	// The time when the API key was last used. The value is a time string in RFC 3339 format. This parameter is empty if the API key has never been used.
	//
	// example:
	//
	// 2025-12-12T09:30:00Z
	LastUsedTime *string `json:"LastUsedTime,omitempty" xml:"LastUsedTime,omitempty"`
	// The name of the API key.
	//
	// example:
	//
	// order-analysis-agent-default
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The plaintext secret of the API key. This value is returned only once in this creation response. Store it securely. Subsequent API calls do not return the plaintext secret again.
	//
	// example:
	//
	// dms_sk_1a2b3c4d****
	Secret *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
	// The credential source. Valid values: console (issued from the console), oauth (issued through the OAuth flow), install_token (issued through the install-and-authenticate flow). The API key automatically issued by this operation is always console.
	//
	// example:
	//
	// console
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreateAgentResponseBodyDataApiKey) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentResponseBodyDataApiKey) GoString() string {
	return s.String()
}

func (s *CreateAgentResponseBodyDataApiKey) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateAgentResponseBodyDataApiKey) GetAgentName() *string {
	return s.AgentName
}

func (s *CreateAgentResponseBodyDataApiKey) GetAgentType() *string {
	return s.AgentType
}

func (s *CreateAgentResponseBodyDataApiKey) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateAgentResponseBodyDataApiKey) GetCreatorId() *string {
	return s.CreatorId
}

func (s *CreateAgentResponseBodyDataApiKey) GetCreatorName() *string {
	return s.CreatorName
}

func (s *CreateAgentResponseBodyDataApiKey) GetDescription() *string {
	return s.Description
}

func (s *CreateAgentResponseBodyDataApiKey) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *CreateAgentResponseBodyDataApiKey) GetId() *int64 {
	return s.Id
}

func (s *CreateAgentResponseBodyDataApiKey) GetIsRevoked() *bool {
	return s.IsRevoked
}

func (s *CreateAgentResponseBodyDataApiKey) GetKeyPrefix() *string {
	return s.KeyPrefix
}

func (s *CreateAgentResponseBodyDataApiKey) GetLastUsedTime() *string {
	return s.LastUsedTime
}

func (s *CreateAgentResponseBodyDataApiKey) GetName() *string {
	return s.Name
}

func (s *CreateAgentResponseBodyDataApiKey) GetSecret() *string {
	return s.Secret
}

func (s *CreateAgentResponseBodyDataApiKey) GetSource() *string {
	return s.Source
}

func (s *CreateAgentResponseBodyDataApiKey) SetAgentId(v string) *CreateAgentResponseBodyDataApiKey {
	s.AgentId = &v
	return s
}

func (s *CreateAgentResponseBodyDataApiKey) SetAgentName(v string) *CreateAgentResponseBodyDataApiKey {
	s.AgentName = &v
	return s
}

func (s *CreateAgentResponseBodyDataApiKey) SetAgentType(v string) *CreateAgentResponseBodyDataApiKey {
	s.AgentType = &v
	return s
}

func (s *CreateAgentResponseBodyDataApiKey) SetCreatedAt(v string) *CreateAgentResponseBodyDataApiKey {
	s.CreatedAt = &v
	return s
}

func (s *CreateAgentResponseBodyDataApiKey) SetCreatorId(v string) *CreateAgentResponseBodyDataApiKey {
	s.CreatorId = &v
	return s
}

func (s *CreateAgentResponseBodyDataApiKey) SetCreatorName(v string) *CreateAgentResponseBodyDataApiKey {
	s.CreatorName = &v
	return s
}

func (s *CreateAgentResponseBodyDataApiKey) SetDescription(v string) *CreateAgentResponseBodyDataApiKey {
	s.Description = &v
	return s
}

func (s *CreateAgentResponseBodyDataApiKey) SetExpireTime(v string) *CreateAgentResponseBodyDataApiKey {
	s.ExpireTime = &v
	return s
}

func (s *CreateAgentResponseBodyDataApiKey) SetId(v int64) *CreateAgentResponseBodyDataApiKey {
	s.Id = &v
	return s
}

func (s *CreateAgentResponseBodyDataApiKey) SetIsRevoked(v bool) *CreateAgentResponseBodyDataApiKey {
	s.IsRevoked = &v
	return s
}

func (s *CreateAgentResponseBodyDataApiKey) SetKeyPrefix(v string) *CreateAgentResponseBodyDataApiKey {
	s.KeyPrefix = &v
	return s
}

func (s *CreateAgentResponseBodyDataApiKey) SetLastUsedTime(v string) *CreateAgentResponseBodyDataApiKey {
	s.LastUsedTime = &v
	return s
}

func (s *CreateAgentResponseBodyDataApiKey) SetName(v string) *CreateAgentResponseBodyDataApiKey {
	s.Name = &v
	return s
}

func (s *CreateAgentResponseBodyDataApiKey) SetSecret(v string) *CreateAgentResponseBodyDataApiKey {
	s.Secret = &v
	return s
}

func (s *CreateAgentResponseBodyDataApiKey) SetSource(v string) *CreateAgentResponseBodyDataApiKey {
	s.Source = &v
	return s
}

func (s *CreateAgentResponseBodyDataApiKey) Validate() error {
	return dara.Validate(s)
}
