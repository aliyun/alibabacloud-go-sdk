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
	Data         *CreateAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *string                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
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
	AgentId      *string                            `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentName    *string                            `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	AgentType    *string                            `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	ApiKey       *CreateAgentResponseBodyDataApiKey `json:"ApiKey,omitempty" xml:"ApiKey,omitempty" type:"Struct"`
	CreatedAt    *string                            `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	CreationType *string                            `json:"CreationType,omitempty" xml:"CreationType,omitempty"`
	Description  *string                            `json:"Description,omitempty" xml:"Description,omitempty"`
	OwnerId      *string                            `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Status       *string                            `json:"Status,omitempty" xml:"Status,omitempty"`
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
	AgentId      *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentName    *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	AgentType    *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	CreatedAt    *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	CreatorId    *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	CreatorName  *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpireTime   *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IsRevoked    *bool   `json:"IsRevoked,omitempty" xml:"IsRevoked,omitempty"`
	KeyPrefix    *string `json:"KeyPrefix,omitempty" xml:"KeyPrefix,omitempty"`
	LastUsedTime *string `json:"LastUsedTime,omitempty" xml:"LastUsedTime,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Secret       *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
	Source       *string `json:"Source,omitempty" xml:"Source,omitempty"`
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
