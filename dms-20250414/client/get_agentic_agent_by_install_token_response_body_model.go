// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgenticAgentByInstallTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetAgenticAgentByInstallTokenResponseBodyData) *GetAgenticAgentByInstallTokenResponseBody
	GetData() *GetAgenticAgentByInstallTokenResponseBodyData
	SetErrorCode(v string) *GetAgenticAgentByInstallTokenResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetAgenticAgentByInstallTokenResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetAgenticAgentByInstallTokenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAgenticAgentByInstallTokenResponseBody
	GetSuccess() *bool
}

type GetAgenticAgentByInstallTokenResponseBody struct {
	// The agent information associated with the install token after redemption. This parameter is empty if the token has not been redeemed, or if the associated credential or agent does not exist.
	Data *GetAgenticAgentByInstallTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned when the request fails. Programs use this code to determine the failure type. This parameter is empty when the request succeeds.
	//
	// example:
	//
	// INSTALL_TOKEN_INVALID
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the request fails. This message helps with manual troubleshooting. This parameter is empty when the request succeeds.
	//
	// example:
	//
	// install_token not found
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The unique identifier of the request, used for troubleshooting and log correlation.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6C-2C4DD51BD5E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. A value of true indicates that the query succeeded. A value of false indicates that the query failed. In this case, you can use ErrorCode and ErrorMessage to identify the cause.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgenticAgentByInstallTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgenticAgentByInstallTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgenticAgentByInstallTokenResponseBody) GetData() *GetAgenticAgentByInstallTokenResponseBodyData {
	return s.Data
}

func (s *GetAgenticAgentByInstallTokenResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetAgenticAgentByInstallTokenResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetAgenticAgentByInstallTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgenticAgentByInstallTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAgenticAgentByInstallTokenResponseBody) SetData(v *GetAgenticAgentByInstallTokenResponseBodyData) *GetAgenticAgentByInstallTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetAgenticAgentByInstallTokenResponseBody) SetErrorCode(v string) *GetAgenticAgentByInstallTokenResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAgenticAgentByInstallTokenResponseBody) SetErrorMessage(v string) *GetAgenticAgentByInstallTokenResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAgenticAgentByInstallTokenResponseBody) SetRequestId(v string) *GetAgenticAgentByInstallTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgenticAgentByInstallTokenResponseBody) SetSuccess(v bool) *GetAgenticAgentByInstallTokenResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgenticAgentByInstallTokenResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgenticAgentByInstallTokenResponseBodyData struct {
	// The prefix list of visible access tokens that are neither revoked nor expired under the agent. This parameter is empty if no access tokens are returned or no access tokens meet the conditional criteria. The prefix is used only to identify the credential and does not contain the complete secret.
	ActiveApiKeyPrefixes []*string `json:"ActiveApiKeyPrefixes,omitempty" xml:"ActiveApiKeyPrefixes,omitempty" type:"Repeated"`
	// The unique identifier of the agent.
	//
	// example:
	//
	// agent-7f3c9a2b
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The name of the agent.
	//
	// example:
	//
	// my-data-agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The permission type of the agent. Valid values: HUMAN_BOUND (inherits user permissions), PERMISSION_NARROW (narrowed down from user permissions), and STANDALONE (independent permissions). The agent generated through the install token redemption process is of the HUMAN_BOUND type.
	//
	// example:
	//
	// HUMAN_BOUND
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// The creation time of the agent, in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2026-07-13 10:30:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The creation method of the agent. A value of auto indicates automatic creation by the system, and a value of manual indicates manual creation by a user. The agent generated through the install token redemption process is auto.
	//
	// example:
	//
	// auto
	CreationType *string `json:"CreationType,omitempty" xml:"CreationType,omitempty"`
	// The description of the agent. This parameter is empty if it is not set.
	//
	// example:
	//
	// Agent for data analytics
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the user to whom the agent belongs.
	//
	// example:
	//
	// 27400000000000001
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The status of the agent. A value of active indicates that the agent is enabled, and a value of disabled indicates that the agent is disabled.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAgenticAgentByInstallTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgenticAgentByInstallTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgenticAgentByInstallTokenResponseBodyData) GetActiveApiKeyPrefixes() []*string {
	return s.ActiveApiKeyPrefixes
}

func (s *GetAgenticAgentByInstallTokenResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *GetAgenticAgentByInstallTokenResponseBodyData) GetAgentName() *string {
	return s.AgentName
}

func (s *GetAgenticAgentByInstallTokenResponseBodyData) GetAgentType() *string {
	return s.AgentType
}

func (s *GetAgenticAgentByInstallTokenResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetAgenticAgentByInstallTokenResponseBodyData) GetCreationType() *string {
	return s.CreationType
}

func (s *GetAgenticAgentByInstallTokenResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetAgenticAgentByInstallTokenResponseBodyData) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetAgenticAgentByInstallTokenResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetAgenticAgentByInstallTokenResponseBodyData) SetActiveApiKeyPrefixes(v []*string) *GetAgenticAgentByInstallTokenResponseBodyData {
	s.ActiveApiKeyPrefixes = v
	return s
}

func (s *GetAgenticAgentByInstallTokenResponseBodyData) SetAgentId(v string) *GetAgenticAgentByInstallTokenResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *GetAgenticAgentByInstallTokenResponseBodyData) SetAgentName(v string) *GetAgenticAgentByInstallTokenResponseBodyData {
	s.AgentName = &v
	return s
}

func (s *GetAgenticAgentByInstallTokenResponseBodyData) SetAgentType(v string) *GetAgenticAgentByInstallTokenResponseBodyData {
	s.AgentType = &v
	return s
}

func (s *GetAgenticAgentByInstallTokenResponseBodyData) SetCreatedAt(v string) *GetAgenticAgentByInstallTokenResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetAgenticAgentByInstallTokenResponseBodyData) SetCreationType(v string) *GetAgenticAgentByInstallTokenResponseBodyData {
	s.CreationType = &v
	return s
}

func (s *GetAgenticAgentByInstallTokenResponseBodyData) SetDescription(v string) *GetAgenticAgentByInstallTokenResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetAgenticAgentByInstallTokenResponseBodyData) SetOwnerId(v string) *GetAgenticAgentByInstallTokenResponseBodyData {
	s.OwnerId = &v
	return s
}

func (s *GetAgenticAgentByInstallTokenResponseBodyData) SetStatus(v string) *GetAgenticAgentByInstallTokenResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAgenticAgentByInstallTokenResponseBodyData) Validate() error {
	return dara.Validate(s)
}
