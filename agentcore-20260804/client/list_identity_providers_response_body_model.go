// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentityProvidersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListIdentityProvidersResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListIdentityProvidersResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListIdentityProvidersResponseBodyItems) *ListIdentityProvidersResponseBody
	GetItems() []*ListIdentityProvidersResponseBodyItems
	SetMaxResults(v int32) *ListIdentityProvidersResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListIdentityProvidersResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListIdentityProvidersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListIdentityProvidersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListIdentityProvidersResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListIdentityProvidersResponseBody
	GetTotalCount() *int64
}

type ListIdentityProvidersResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The list of external identity providers.
	Items []*ListIdentityProvidersResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of records per page that takes effect for this query.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message. An error description is returned if the request fails.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token for the next page. This parameter is empty if no more pages exist.
	//
	// example:
	//
	// aWRlbnRpdHktcHJvdmlkZXItb2Zmc2V0OjEw
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of external identity providers that match the query conditions.
	//
	// example:
	//
	// 42
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListIdentityProvidersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProvidersResponseBody) GoString() string {
	return s.String()
}

func (s *ListIdentityProvidersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListIdentityProvidersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListIdentityProvidersResponseBody) GetItems() []*ListIdentityProvidersResponseBodyItems {
	return s.Items
}

func (s *ListIdentityProvidersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIdentityProvidersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListIdentityProvidersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIdentityProvidersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIdentityProvidersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListIdentityProvidersResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListIdentityProvidersResponseBody) SetCode(v string) *ListIdentityProvidersResponseBody {
	s.Code = &v
	return s
}

func (s *ListIdentityProvidersResponseBody) SetHttpStatusCode(v int32) *ListIdentityProvidersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListIdentityProvidersResponseBody) SetItems(v []*ListIdentityProvidersResponseBodyItems) *ListIdentityProvidersResponseBody {
	s.Items = v
	return s
}

func (s *ListIdentityProvidersResponseBody) SetMaxResults(v int32) *ListIdentityProvidersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIdentityProvidersResponseBody) SetMessage(v string) *ListIdentityProvidersResponseBody {
	s.Message = &v
	return s
}

func (s *ListIdentityProvidersResponseBody) SetNextToken(v string) *ListIdentityProvidersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIdentityProvidersResponseBody) SetRequestId(v string) *ListIdentityProvidersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIdentityProvidersResponseBody) SetSuccess(v bool) *ListIdentityProvidersResponseBody {
	s.Success = &v
	return s
}

func (s *ListIdentityProvidersResponseBody) SetTotalCount(v int64) *ListIdentityProvidersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIdentityProvidersResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIdentityProvidersResponseBodyItems struct {
	// The creation time in UTC, formatted according to RFC 3339.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The event subscription callback URL. Configure this URL in the external identity provider application to receive organization change events. An empty string is returned if the user pool has not been provisioned.
	//
	// example:
	//
	// http://auth.cn-hangzhou.agentteams.aliyuncs.com/roa/dingtalk/event/up-123456
	EventSubscriptionCallbackUrl *string `json:"eventSubscriptionCallbackUrl,omitempty" xml:"eventSubscriptionCallbackUrl,omitempty"`
	// The type of the external identity provider. Valid values:
	//
	// - DingTalk
	//
	// - Feishu
	//
	// example:
	//
	// DingTalk
	IdentityProviderType *string `json:"identityProviderType,omitempty" xml:"identityProviderType,omitempty"`
	// The logon callback URL. Configure this URL in the external identity provider application. An empty string is returned if the user pool has not been provisioned.
	//
	// example:
	//
	// https://signin-cn-hangzhou.aliyunagentid.com/up-123456/dingtalk/callback
	LoginCallbackUrl *string `json:"loginCallbackUrl,omitempty" xml:"loginCallbackUrl,omitempty"`
	// Indicates whether workspace users are allowed to log on through this external identity provider.
	LoginEnabled *bool `json:"loginEnabled,omitempty" xml:"loginEnabled,omitempty"`
	// The application configuration of the external identity provider. Application secret configurations are not returned.
	Metadata *ListIdentityProvidersResponseBodyItemsMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
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
	// Indicates whether organization member synchronization is enabled. When enabled, organization members are synchronized from this external identity provider as workspace users.
	SyncEnabled *bool `json:"syncEnabled,omitempty" xml:"syncEnabled,omitempty"`
	// The last modification time in UTC, formatted according to RFC 3339.
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

func (s ListIdentityProvidersResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProvidersResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListIdentityProvidersResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListIdentityProvidersResponseBodyItems) GetEventSubscriptionCallbackUrl() *string {
	return s.EventSubscriptionCallbackUrl
}

func (s *ListIdentityProvidersResponseBodyItems) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *ListIdentityProvidersResponseBodyItems) GetLoginCallbackUrl() *string {
	return s.LoginCallbackUrl
}

func (s *ListIdentityProvidersResponseBodyItems) GetLoginEnabled() *bool {
	return s.LoginEnabled
}

func (s *ListIdentityProvidersResponseBodyItems) GetMetadata() *ListIdentityProvidersResponseBodyItemsMetadata {
	return s.Metadata
}

func (s *ListIdentityProvidersResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListIdentityProvidersResponseBodyItems) GetSyncEnabled() *bool {
	return s.SyncEnabled
}

func (s *ListIdentityProvidersResponseBodyItems) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListIdentityProvidersResponseBodyItems) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListIdentityProvidersResponseBodyItems) SetCreatedAt(v string) *ListIdentityProvidersResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItems) SetEventSubscriptionCallbackUrl(v string) *ListIdentityProvidersResponseBodyItems {
	s.EventSubscriptionCallbackUrl = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItems) SetIdentityProviderType(v string) *ListIdentityProvidersResponseBodyItems {
	s.IdentityProviderType = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItems) SetLoginCallbackUrl(v string) *ListIdentityProvidersResponseBodyItems {
	s.LoginCallbackUrl = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItems) SetLoginEnabled(v bool) *ListIdentityProvidersResponseBodyItems {
	s.LoginEnabled = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItems) SetMetadata(v *ListIdentityProvidersResponseBodyItemsMetadata) *ListIdentityProvidersResponseBodyItems {
	s.Metadata = v
	return s
}

func (s *ListIdentityProvidersResponseBodyItems) SetStatus(v string) *ListIdentityProvidersResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItems) SetSyncEnabled(v bool) *ListIdentityProvidersResponseBodyItems {
	s.SyncEnabled = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItems) SetUpdatedAt(v string) *ListIdentityProvidersResponseBodyItems {
	s.UpdatedAt = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItems) SetWorkspaceId(v string) *ListIdentityProvidersResponseBodyItems {
	s.WorkspaceId = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItems) Validate() error {
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListIdentityProvidersResponseBodyItemsMetadata struct {
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

func (s ListIdentityProvidersResponseBodyItemsMetadata) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProvidersResponseBodyItemsMetadata) GoString() string {
	return s.String()
}

func (s *ListIdentityProvidersResponseBodyItemsMetadata) GetAppId() *string {
	return s.AppId
}

func (s *ListIdentityProvidersResponseBodyItemsMetadata) GetAppKey() *string {
	return s.AppKey
}

func (s *ListIdentityProvidersResponseBodyItemsMetadata) GetCorpId() *string {
	return s.CorpId
}

func (s *ListIdentityProvidersResponseBodyItemsMetadata) SetAppId(v string) *ListIdentityProvidersResponseBodyItemsMetadata {
	s.AppId = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItemsMetadata) SetAppKey(v string) *ListIdentityProvidersResponseBodyItemsMetadata {
	s.AppKey = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItemsMetadata) SetCorpId(v string) *ListIdentityProvidersResponseBodyItemsMetadata {
	s.CorpId = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItemsMetadata) Validate() error {
	return dara.Validate(s)
}
