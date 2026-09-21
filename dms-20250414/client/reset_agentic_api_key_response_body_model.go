// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAgenticApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ResetAgenticApiKeyResponseBodyData) *ResetAgenticApiKeyResponseBody
	GetData() *ResetAgenticApiKeyResponseBodyData
	SetErrorCode(v string) *ResetAgenticApiKeyResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ResetAgenticApiKeyResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ResetAgenticApiKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ResetAgenticApiKeyResponseBody
	GetSuccess() *bool
}

type ResetAgenticApiKeyResponseBody struct {
	// The Access Token information returned after a successful reset. The Secret field contains the new plaintext Secret, which is returned only once in this response.
	Data *ResetAgenticApiKeyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned when the request fails. You can use this code to programmatically determine the failure type. This value is empty when the request succeeds.
	//
	// example:
	//
	// ACCESS_TOKEN_NOT_FOUND
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the request fails. This message helps you locate the issue. This value is empty when the request succeeds.
	//
	// example:
	//
	// access token not found: 1024
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The unique request ID, which is used for troubleshooting and log correlation.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6C-2C4DD51BD5E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. A value of true indicates that the reset was successful. A value of false indicates a failure. In this case, check ErrorCode and ErrorMessage to identify the cause.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResetAgenticApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetAgenticApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAgenticApiKeyResponseBody) GetData() *ResetAgenticApiKeyResponseBodyData {
	return s.Data
}

func (s *ResetAgenticApiKeyResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ResetAgenticApiKeyResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ResetAgenticApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetAgenticApiKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResetAgenticApiKeyResponseBody) SetData(v *ResetAgenticApiKeyResponseBodyData) *ResetAgenticApiKeyResponseBody {
	s.Data = v
	return s
}

func (s *ResetAgenticApiKeyResponseBody) SetErrorCode(v string) *ResetAgenticApiKeyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ResetAgenticApiKeyResponseBody) SetErrorMessage(v string) *ResetAgenticApiKeyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ResetAgenticApiKeyResponseBody) SetRequestId(v string) *ResetAgenticApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetAgenticApiKeyResponseBody) SetSuccess(v bool) *ResetAgenticApiKeyResponseBody {
	s.Success = &v
	return s
}

func (s *ResetAgenticApiKeyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResetAgenticApiKeyResponseBodyData struct {
	// The ID of the Agent to which the Access Token belongs.
	//
	// example:
	//
	// agent-7f3c9a2b
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The name of the Agent to which the Access Token belongs.
	//
	// example:
	//
	// my-data-agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The type of the Agent to which the Access Token belongs. Valid values:
	//
	// - HUMAN_BOUND: fully inherits the permissions of the associated user.
	//
	// - PERMISSION_NARROW: narrows the permissions based on the associated user\\"s permission baseline.
	//
	// - AGENT_BOUND: inherits the permissions of the parent Agent.
	//
	// - STANDALONE: holds permissions as an independent identity principal.
	//
	// example:
	//
	// STANDALONE
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// The time when the Access Token was created, in the yyyy-MM-dd HH:mm:ss format (UTC+8). This value remains unchanged after the reset.
	//
	// example:
	//
	// 2026-05-13 08:00:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The user ID of the Access Token creator.
	//
	// example:
	//
	// 27400000000000001
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The display name of the Access Token creator.
	//
	// example:
	//
	// alice
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The description of the Access Token. This value remains unchanged after the reset.
	//
	// example:
	//
	// Access token used by the data analysis agent
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The expiration time of the Access Token, in the yyyy-MM-dd HH:mm:ss format (UTC+8). If ExpireAfterSeconds is specified, the expiration time is recalculated from the time of the reset. If ExpireAfterSeconds is not specified, the original expiration time is retained.
	//
	// example:
	//
	// 2027-05-13 10:20:30
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the reset Access Token. This value remains unchanged after the reset.
	//
	// example:
	//
	// 1024
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the Access Token has been revoked. An Access Token returned after a successful reset is always in the non-revoked state (false).
	IsRevoked *bool `json:"IsRevoked,omitempty" xml:"IsRevoked,omitempty"`
	// The visible prefix of the Access Token, which is used to identify the Access Token without exposing the full Secret. This value remains unchanged after the reset.
	//
	// example:
	//
	// dms_sk_3f9a
	KeyPrefix *string `json:"KeyPrefix,omitempty" xml:"KeyPrefix,omitempty"`
	// The time when the Access Token was last used, in the yyyy-MM-dd HH:mm:ss format (UTC+8). This value is empty if the Access Token has never been used.
	//
	// example:
	//
	// 2026-05-13 09:15:00
	LastUsedTime *string `json:"LastUsedTime,omitempty" xml:"LastUsedTime,omitempty"`
	// The name of the Access Token. This value remains unchanged after the reset.
	//
	// example:
	//
	// prod-readonly-key
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The new plaintext Secret generated by this reset. This value is returned only once in this response and will not be returned by any subsequent operation. Store it securely right away. The old Secret becomes invalid immediately after the reset.
	//
	// example:
	//
	// dms_sk_3f9a1c8e5b7d4062a1f6c9e2b8d05a3f
	Secret *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
	// The credential source of the Access Token. The reset operation supports only Access Tokens issued by the console. Therefore, the value is always console.
	//
	// example:
	//
	// console
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The reminder information related to this reset, such as a notice that the new Secret is returned only once and must be stored immediately. This value is empty if no reminder exists.
	//
	// example:
	//
	// The new secret is shown only once. Please store it securely now.
	Warning *string `json:"Warning,omitempty" xml:"Warning,omitempty"`
}

func (s ResetAgenticApiKeyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ResetAgenticApiKeyResponseBodyData) GoString() string {
	return s.String()
}

func (s *ResetAgenticApiKeyResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *ResetAgenticApiKeyResponseBodyData) GetAgentName() *string {
	return s.AgentName
}

func (s *ResetAgenticApiKeyResponseBodyData) GetAgentType() *string {
	return s.AgentType
}

func (s *ResetAgenticApiKeyResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ResetAgenticApiKeyResponseBodyData) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ResetAgenticApiKeyResponseBodyData) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ResetAgenticApiKeyResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ResetAgenticApiKeyResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ResetAgenticApiKeyResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ResetAgenticApiKeyResponseBodyData) GetIsRevoked() *bool {
	return s.IsRevoked
}

func (s *ResetAgenticApiKeyResponseBodyData) GetKeyPrefix() *string {
	return s.KeyPrefix
}

func (s *ResetAgenticApiKeyResponseBodyData) GetLastUsedTime() *string {
	return s.LastUsedTime
}

func (s *ResetAgenticApiKeyResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ResetAgenticApiKeyResponseBodyData) GetSecret() *string {
	return s.Secret
}

func (s *ResetAgenticApiKeyResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *ResetAgenticApiKeyResponseBodyData) GetWarning() *string {
	return s.Warning
}

func (s *ResetAgenticApiKeyResponseBodyData) SetAgentId(v string) *ResetAgenticApiKeyResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *ResetAgenticApiKeyResponseBodyData) SetAgentName(v string) *ResetAgenticApiKeyResponseBodyData {
	s.AgentName = &v
	return s
}

func (s *ResetAgenticApiKeyResponseBodyData) SetAgentType(v string) *ResetAgenticApiKeyResponseBodyData {
	s.AgentType = &v
	return s
}

func (s *ResetAgenticApiKeyResponseBodyData) SetCreatedAt(v string) *ResetAgenticApiKeyResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *ResetAgenticApiKeyResponseBodyData) SetCreatorId(v string) *ResetAgenticApiKeyResponseBodyData {
	s.CreatorId = &v
	return s
}

func (s *ResetAgenticApiKeyResponseBodyData) SetCreatorName(v string) *ResetAgenticApiKeyResponseBodyData {
	s.CreatorName = &v
	return s
}

func (s *ResetAgenticApiKeyResponseBodyData) SetDescription(v string) *ResetAgenticApiKeyResponseBodyData {
	s.Description = &v
	return s
}

func (s *ResetAgenticApiKeyResponseBodyData) SetExpireTime(v string) *ResetAgenticApiKeyResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *ResetAgenticApiKeyResponseBodyData) SetId(v int64) *ResetAgenticApiKeyResponseBodyData {
	s.Id = &v
	return s
}

func (s *ResetAgenticApiKeyResponseBodyData) SetIsRevoked(v bool) *ResetAgenticApiKeyResponseBodyData {
	s.IsRevoked = &v
	return s
}

func (s *ResetAgenticApiKeyResponseBodyData) SetKeyPrefix(v string) *ResetAgenticApiKeyResponseBodyData {
	s.KeyPrefix = &v
	return s
}

func (s *ResetAgenticApiKeyResponseBodyData) SetLastUsedTime(v string) *ResetAgenticApiKeyResponseBodyData {
	s.LastUsedTime = &v
	return s
}

func (s *ResetAgenticApiKeyResponseBodyData) SetName(v string) *ResetAgenticApiKeyResponseBodyData {
	s.Name = &v
	return s
}

func (s *ResetAgenticApiKeyResponseBodyData) SetSecret(v string) *ResetAgenticApiKeyResponseBodyData {
	s.Secret = &v
	return s
}

func (s *ResetAgenticApiKeyResponseBodyData) SetSource(v string) *ResetAgenticApiKeyResponseBodyData {
	s.Source = &v
	return s
}

func (s *ResetAgenticApiKeyResponseBodyData) SetWarning(v string) *ResetAgenticApiKeyResponseBodyData {
	s.Warning = &v
	return s
}

func (s *ResetAgenticApiKeyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
