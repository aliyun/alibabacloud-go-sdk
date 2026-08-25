// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateModelConnectionResponseBody
	GetCode() *string
	SetData(v *UpdateModelConnectionResponseBodyData) *UpdateModelConnectionResponseBody
	GetData() *UpdateModelConnectionResponseBodyData
	SetHttpStatusCode(v int32) *UpdateModelConnectionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateModelConnectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateModelConnectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateModelConnectionResponseBody
	GetSuccess() *bool
}

type UpdateModelConnectionResponseBody struct {
	Code           *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data           *UpdateModelConnectionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string                                `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateModelConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateModelConnectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateModelConnectionResponseBody) GetData() *UpdateModelConnectionResponseBodyData {
	return s.Data
}

func (s *UpdateModelConnectionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateModelConnectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateModelConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateModelConnectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateModelConnectionResponseBody) SetCode(v string) *UpdateModelConnectionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateModelConnectionResponseBody) SetData(v *UpdateModelConnectionResponseBodyData) *UpdateModelConnectionResponseBody {
	s.Data = v
	return s
}

func (s *UpdateModelConnectionResponseBody) SetHttpStatusCode(v int32) *UpdateModelConnectionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateModelConnectionResponseBody) SetMessage(v string) *UpdateModelConnectionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateModelConnectionResponseBody) SetRequestId(v string) *UpdateModelConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateModelConnectionResponseBody) SetSuccess(v bool) *UpdateModelConnectionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateModelConnectionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateModelConnectionResponseBodyData struct {
	ApiKeyCount          *int32  `json:"apiKeyCount,omitempty" xml:"apiKeyCount,omitempty"`
	ConnectionId         *string `json:"connectionId,omitempty" xml:"connectionId,omitempty"`
	CreatedAt            *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CredentialConfigured *bool   `json:"credentialConfigured,omitempty" xml:"credentialConfigured,omitempty"`
	Description          *string `json:"description,omitempty" xml:"description,omitempty"`
	Endpoint             *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Name                 *string `json:"name,omitempty" xml:"name,omitempty"`
	Protocol             *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	ProviderType         *string `json:"providerType,omitempty" xml:"providerType,omitempty"`
	Status               *string `json:"status,omitempty" xml:"status,omitempty"`
	StatusReason         *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	UpdatedAt            *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	WorkspaceId          *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s UpdateModelConnectionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelConnectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateModelConnectionResponseBodyData) GetApiKeyCount() *int32 {
	return s.ApiKeyCount
}

func (s *UpdateModelConnectionResponseBodyData) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *UpdateModelConnectionResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateModelConnectionResponseBodyData) GetCredentialConfigured() *bool {
	return s.CredentialConfigured
}

func (s *UpdateModelConnectionResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *UpdateModelConnectionResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *UpdateModelConnectionResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateModelConnectionResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateModelConnectionResponseBodyData) GetProviderType() *string {
	return s.ProviderType
}

func (s *UpdateModelConnectionResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateModelConnectionResponseBodyData) GetStatusReason() *string {
	return s.StatusReason
}

func (s *UpdateModelConnectionResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateModelConnectionResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateModelConnectionResponseBodyData) SetApiKeyCount(v int32) *UpdateModelConnectionResponseBodyData {
	s.ApiKeyCount = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) SetConnectionId(v string) *UpdateModelConnectionResponseBodyData {
	s.ConnectionId = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) SetCreatedAt(v string) *UpdateModelConnectionResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) SetCredentialConfigured(v bool) *UpdateModelConnectionResponseBodyData {
	s.CredentialConfigured = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) SetDescription(v string) *UpdateModelConnectionResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) SetEndpoint(v string) *UpdateModelConnectionResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) SetName(v string) *UpdateModelConnectionResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) SetProtocol(v string) *UpdateModelConnectionResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) SetProviderType(v string) *UpdateModelConnectionResponseBodyData {
	s.ProviderType = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) SetStatus(v string) *UpdateModelConnectionResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) SetStatusReason(v string) *UpdateModelConnectionResponseBodyData {
	s.StatusReason = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) SetUpdatedAt(v string) *UpdateModelConnectionResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) SetWorkspaceId(v string) *UpdateModelConnectionResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateModelConnectionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
