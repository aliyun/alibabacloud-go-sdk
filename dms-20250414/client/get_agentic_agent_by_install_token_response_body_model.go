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
	Data         *GetAgenticAgentByInstallTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *string                                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
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
	ActiveApiKeyPrefixes []*string `json:"ActiveApiKeyPrefixes,omitempty" xml:"ActiveApiKeyPrefixes,omitempty" type:"Repeated"`
	AgentId              *string   `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentName            *string   `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	AgentType            *string   `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	CreatedAt            *string   `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	CreationType         *string   `json:"CreationType,omitempty" xml:"CreationType,omitempty"`
	Description          *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	OwnerId              *string   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Status               *string   `json:"Status,omitempty" xml:"Status,omitempty"`
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
