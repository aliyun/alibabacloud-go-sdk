// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetIdentityProviderResponseBody
	GetCode() *string
	SetData(v *GetIdentityProviderResponseBodyData) *GetIdentityProviderResponseBody
	GetData() *GetIdentityProviderResponseBodyData
	SetHttpStatusCode(v int32) *GetIdentityProviderResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetIdentityProviderResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetIdentityProviderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetIdentityProviderResponseBody
	GetSuccess() *bool
}

type GetIdentityProviderResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The binding details of the external identity provider.
	Data *GetIdentityProviderResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s GetIdentityProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetIdentityProviderResponseBody) GetData() *GetIdentityProviderResponseBodyData {
	return s.Data
}

func (s *GetIdentityProviderResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetIdentityProviderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetIdentityProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIdentityProviderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetIdentityProviderResponseBody) SetCode(v string) *GetIdentityProviderResponseBody {
	s.Code = &v
	return s
}

func (s *GetIdentityProviderResponseBody) SetData(v *GetIdentityProviderResponseBodyData) *GetIdentityProviderResponseBody {
	s.Data = v
	return s
}

func (s *GetIdentityProviderResponseBody) SetHttpStatusCode(v int32) *GetIdentityProviderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetIdentityProviderResponseBody) SetMessage(v string) *GetIdentityProviderResponseBody {
	s.Message = &v
	return s
}

func (s *GetIdentityProviderResponseBody) SetRequestId(v string) *GetIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIdentityProviderResponseBody) SetSuccess(v bool) *GetIdentityProviderResponseBody {
	s.Success = &v
	return s
}

func (s *GetIdentityProviderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIdentityProviderResponseBodyData struct {
	// The creation time in UTC, formatted according to RFC 3339.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The event subscription callback URL. Configure this URL in the application on the external identity provider side to receive organization change events. An empty string is returned if the user pool has not been provisioned.
	//
	// example:
	//
	// http://auth.cn-hangzhou.agentteams.aliyuncs.com/roa/dingtalk/event/up-123456
	EventSubscriptionCallbackUrl *string `json:"eventSubscriptionCallbackUrl,omitempty" xml:"eventSubscriptionCallbackUrl,omitempty"`
	// The type of the external identity provider. Valid values: DingTalk, Feishu.
	//
	// example:
	//
	// DingTalk
	IdentityProviderType *string `json:"identityProviderType,omitempty" xml:"identityProviderType,omitempty"`
	// The logon callback URL. Configure this URL in the application on the external identity provider side. An empty string is returned if the user pool has not been provisioned.
	//
	// example:
	//
	// https://signin-cn-hangzhou.aliyunagentid.com/up-123456/dingtalk/callback
	LoginCallbackUrl *string `json:"loginCallbackUrl,omitempty" xml:"loginCallbackUrl,omitempty"`
	// Indicates whether workspace users are allowed to log on through this external identity provider.
	LoginEnabled *bool `json:"loginEnabled,omitempty" xml:"loginEnabled,omitempty"`
	// The application configuration of the external identity provider. Application secret configurations are not returned.
	Metadata *GetIdentityProviderResponseBodyDataMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	// The status. Valid values:
	//
	// - CONFIGURED: The configuration has been accepted and is waiting for the user pool to be provisioned.
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
	// Indicates whether organization member synchronization is enabled. When enabled, the external identity provider synchronizes organization members as workspace users.
	SyncEnabled *bool `json:"syncEnabled,omitempty" xml:"syncEnabled,omitempty"`
	// The time of the last modification in UTC, formatted according to RFC 3339.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetIdentityProviderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetIdentityProviderResponseBodyData) GetEventSubscriptionCallbackUrl() *string {
	return s.EventSubscriptionCallbackUrl
}

func (s *GetIdentityProviderResponseBodyData) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *GetIdentityProviderResponseBodyData) GetLoginCallbackUrl() *string {
	return s.LoginCallbackUrl
}

func (s *GetIdentityProviderResponseBodyData) GetLoginEnabled() *bool {
	return s.LoginEnabled
}

func (s *GetIdentityProviderResponseBodyData) GetMetadata() *GetIdentityProviderResponseBodyDataMetadata {
	return s.Metadata
}

func (s *GetIdentityProviderResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetIdentityProviderResponseBodyData) GetSyncEnabled() *bool {
	return s.SyncEnabled
}

func (s *GetIdentityProviderResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetIdentityProviderResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetIdentityProviderResponseBodyData) SetCreatedAt(v string) *GetIdentityProviderResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetEventSubscriptionCallbackUrl(v string) *GetIdentityProviderResponseBodyData {
	s.EventSubscriptionCallbackUrl = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetIdentityProviderType(v string) *GetIdentityProviderResponseBodyData {
	s.IdentityProviderType = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetLoginCallbackUrl(v string) *GetIdentityProviderResponseBodyData {
	s.LoginCallbackUrl = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetLoginEnabled(v bool) *GetIdentityProviderResponseBodyData {
	s.LoginEnabled = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetMetadata(v *GetIdentityProviderResponseBodyDataMetadata) *GetIdentityProviderResponseBodyData {
	s.Metadata = v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetStatus(v string) *GetIdentityProviderResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetSyncEnabled(v bool) *GetIdentityProviderResponseBodyData {
	s.SyncEnabled = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetUpdatedAt(v string) *GetIdentityProviderResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetWorkspaceId(v string) *GetIdentityProviderResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) Validate() error {
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIdentityProviderResponseBodyDataMetadata struct {
	// The App ID of the Lark application. Required when the binding type is Feishu.
	//
	// example:
	//
	// cli_exampleappid01
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// The AppKey of the DingTalk application. Required when the binding type is DingTalk.
	//
	// example:
	//
	// dingexampleappkey01
	AppKey *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	// The CorpId of the DingTalk enterprise. Required when the binding type is DingTalk.
	//
	// example:
	//
	// dingexamplecorpid01
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s GetIdentityProviderResponseBodyDataMetadata) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyDataMetadata) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyDataMetadata) GetAppId() *string {
	return s.AppId
}

func (s *GetIdentityProviderResponseBodyDataMetadata) GetAppKey() *string {
	return s.AppKey
}

func (s *GetIdentityProviderResponseBodyDataMetadata) GetCorpId() *string {
	return s.CorpId
}

func (s *GetIdentityProviderResponseBodyDataMetadata) SetAppId(v string) *GetIdentityProviderResponseBodyDataMetadata {
	s.AppId = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataMetadata) SetAppKey(v string) *GetIdentityProviderResponseBodyDataMetadata {
	s.AppKey = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataMetadata) SetCorpId(v string) *GetIdentityProviderResponseBodyDataMetadata {
	s.CorpId = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataMetadata) Validate() error {
	return dara.Validate(s)
}
