// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdentityProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateIdentityProviderResponseBody
	GetCode() *string
	SetData(v *CreateIdentityProviderResponseBodyData) *CreateIdentityProviderResponseBody
	GetData() *CreateIdentityProviderResponseBodyData
	SetHttpStatusCode(v int32) *CreateIdentityProviderResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateIdentityProviderResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateIdentityProviderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateIdentityProviderResponseBody
	GetSuccess() *bool
}

type CreateIdentityProviderResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The binding information of the external identity provider.
	Data *CreateIdentityProviderResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message. An error description is returned if the request fails.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateIdentityProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateIdentityProviderResponseBody) GetData() *CreateIdentityProviderResponseBodyData {
	return s.Data
}

func (s *CreateIdentityProviderResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateIdentityProviderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateIdentityProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIdentityProviderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateIdentityProviderResponseBody) SetCode(v string) *CreateIdentityProviderResponseBody {
	s.Code = &v
	return s
}

func (s *CreateIdentityProviderResponseBody) SetData(v *CreateIdentityProviderResponseBodyData) *CreateIdentityProviderResponseBody {
	s.Data = v
	return s
}

func (s *CreateIdentityProviderResponseBody) SetHttpStatusCode(v int32) *CreateIdentityProviderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateIdentityProviderResponseBody) SetMessage(v string) *CreateIdentityProviderResponseBody {
	s.Message = &v
	return s
}

func (s *CreateIdentityProviderResponseBody) SetRequestId(v string) *CreateIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIdentityProviderResponseBody) SetSuccess(v bool) *CreateIdentityProviderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateIdentityProviderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateIdentityProviderResponseBodyData struct {
	// The type of the external identity provider. Valid values: DingTalk, Feishu.
	//
	// example:
	//
	// DingTalk
	IdentityProviderType *string `json:"identityProviderType,omitempty" xml:"identityProviderType,omitempty"`
	// Specifies whether workspace users are allowed to log on through this external identity provider.
	LoginEnabled *bool `json:"loginEnabled,omitempty" xml:"loginEnabled,omitempty"`
	// The binding status. Valid values:
	//
	// - CONFIGURED: The configuration has been accepted and is waiting for user pool provisioning.
	//
	// - SYNCING: Organization members are being synchronized.
	//
	// - SYNCED: Organization member synchronization is complete.
	//
	// - READY: The binding is active.
	//
	// - SYNC_FAILED: Organization member synchronization failed.
	//
	// - UPDATING: The configuration is being updated.
	//
	// - UPDATE_FAILED: The configuration update failed.
	//
	// - DISCONNECTING: The binding is being removed.
	//
	// - DISCONNECT_FAILED: The unbinding failed.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Specifies whether to enable organization member synchronization. After this feature is enabled, the external identity provider synchronizes organization members as workspace users.
	SyncEnabled *bool `json:"syncEnabled,omitempty" xml:"syncEnabled,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateIdentityProviderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderResponseBodyData) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *CreateIdentityProviderResponseBodyData) GetLoginEnabled() *bool {
	return s.LoginEnabled
}

func (s *CreateIdentityProviderResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateIdentityProviderResponseBodyData) GetSyncEnabled() *bool {
	return s.SyncEnabled
}

func (s *CreateIdentityProviderResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateIdentityProviderResponseBodyData) SetIdentityProviderType(v string) *CreateIdentityProviderResponseBodyData {
	s.IdentityProviderType = &v
	return s
}

func (s *CreateIdentityProviderResponseBodyData) SetLoginEnabled(v bool) *CreateIdentityProviderResponseBodyData {
	s.LoginEnabled = &v
	return s
}

func (s *CreateIdentityProviderResponseBodyData) SetStatus(v string) *CreateIdentityProviderResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateIdentityProviderResponseBodyData) SetSyncEnabled(v bool) *CreateIdentityProviderResponseBodyData {
	s.SyncEnabled = &v
	return s
}

func (s *CreateIdentityProviderResponseBodyData) SetWorkspaceId(v string) *CreateIdentityProviderResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreateIdentityProviderResponseBodyData) Validate() error {
	return dara.Validate(s)
}
