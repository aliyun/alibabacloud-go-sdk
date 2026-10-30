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
	// The maximum number of records per page that takes effect in this query.
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
	// The cursor for the next page. Empty if there is no next page.
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
	// Indicates whether the request is successful.
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
	// The actual SAML configuration is read on demand only by GetIdentityProvider. SCIM queries and preparations return the actual received status and URLs.
	Configuration *ListIdentityProvidersResponseBodyItemsConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Struct"`
	// The creation time in UTC format based on RFC 3339.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The event subscription callback URL. You must configure this URL in the external identity provider application to receive organization change events. An empty string is returned if the user pool is not activated.
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
	// The logon callback URL. You must configure this URL in the external identity provider application. An empty string is returned if the user pool is not activated.
	//
	// example:
	//
	// https://signin-cn-hangzhou.aliyunagentid.com/up-123456/dingtalk/callback
	LoginCallbackUrl *string `json:"loginCallbackUrl,omitempty" xml:"loginCallbackUrl,omitempty"`
	// Specifies whether workspace users are allowed to log on through this external identity provider.
	LoginEnabled *bool `json:"loginEnabled,omitempty" xml:"loginEnabled,omitempty"`
	// The application configuration of the external identity provider. Application secret configurations are not returned.
	Metadata *ListIdentityProvidersResponseBodyItemsMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	// Only SAML is currently supported. This field may not be returned if no valid configuration is saved. OIDC placeholder initialization is not supported.
	//
	// example:
	//
	// SAML
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The input supports either XML or complete manual entry. Queries return the standardized configuration and do not echo the original XML. This field can be omitted if it is not configured.
	Saml *ListIdentityProvidersResponseBodyItemsSaml `json:"saml,omitempty" xml:"saml,omitempty" type:"Struct"`
	// The binding status. Valid values:
	//
	// - CONFIGURED: The configuration is accepted and waiting for user pool activation.
	//
	// - SYNCING: Organization members are being synchronized.
	//
	// - SYNCED: Organization member synchronization is complete.
	//
	// - READY: The binding is in effect.
	//
	// - SYNC_FAILED: Organization member synchronization failed.
	//
	// - UPDATING: The configuration is being updated.
	//
	// - UPDATE_FAILED: Configuration update failed.
	//
	// - DISCONNECTING: The binding is being disconnected.
	//
	// - DISCONNECT_FAILED: Disconnection failed.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Specifies whether to enable organization member synchronization. If enabled, organization members are synchronized as workspace users by this external identity provider.
	SyncEnabled *bool `json:"syncEnabled,omitempty" xml:"syncEnabled,omitempty"`
	// Follows the department-based synchronization scope definition. If omitted during creation, the default value is AllAuthorized, and the first provider synchronization is determined by syncOnBind. If omitted or set to null during an update, the original value is retained. Updating the saved scope does not automatically trigger full synchronization.
	SyncScope *ListIdentityProvidersResponseBodyItemsSyncScope `json:"syncScope,omitempty" xml:"syncScope,omitempty" type:"Struct"`
	// The time of the last modification in UTC format based on RFC 3339.
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

func (s *ListIdentityProvidersResponseBodyItems) GetConfiguration() *ListIdentityProvidersResponseBodyItemsConfiguration {
	return s.Configuration
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

func (s *ListIdentityProvidersResponseBodyItems) GetProtocol() *string {
	return s.Protocol
}

func (s *ListIdentityProvidersResponseBodyItems) GetSaml() *ListIdentityProvidersResponseBodyItemsSaml {
	return s.Saml
}

func (s *ListIdentityProvidersResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListIdentityProvidersResponseBodyItems) GetSyncEnabled() *bool {
	return s.SyncEnabled
}

func (s *ListIdentityProvidersResponseBodyItems) GetSyncScope() *ListIdentityProvidersResponseBodyItemsSyncScope {
	return s.SyncScope
}

func (s *ListIdentityProvidersResponseBodyItems) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListIdentityProvidersResponseBodyItems) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListIdentityProvidersResponseBodyItems) SetConfiguration(v *ListIdentityProvidersResponseBodyItemsConfiguration) *ListIdentityProvidersResponseBodyItems {
	s.Configuration = v
	return s
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

func (s *ListIdentityProvidersResponseBodyItems) SetProtocol(v string) *ListIdentityProvidersResponseBodyItems {
	s.Protocol = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItems) SetSaml(v *ListIdentityProvidersResponseBodyItemsSaml) *ListIdentityProvidersResponseBodyItems {
	s.Saml = v
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

func (s *ListIdentityProvidersResponseBodyItems) SetSyncScope(v *ListIdentityProvidersResponseBodyItemsSyncScope) *ListIdentityProvidersResponseBodyItems {
	s.SyncScope = v
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
	if s.Configuration != nil {
		if err := s.Configuration.Validate(); err != nil {
			return err
		}
	}
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	if s.Saml != nil {
		if err := s.Saml.Validate(); err != nil {
			return err
		}
	}
	if s.SyncScope != nil {
		if err := s.SyncScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListIdentityProvidersResponseBodyItemsConfiguration struct {
	// HTTP-Redirect or HTTP-POST. Required for manual entry. Queries return the short name.
	//
	// example:
	//
	// HTTP-Redirect
	BindingType *string `json:"bindingType,omitempty" xml:"bindingType,omitempty"`
	// Required for manual entry. It must be non-empty and contain a maximum of 2,048 characters. Existing configurations cannot be changed. This is different from the local entity ID returned in SP queries.
	//
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// Required for manual entry. It must be an absolute HTTPS URL without user credentials or fragments.
	//
	// example:
	//
	// https://idp.example.com/sso
	LoginUrl *string `json:"loginUrl,omitempty" xml:"loginUrl,omitempty"`
	// The actual switch value is Enabled or Disabled. This is not equivalent to the expected loginEnabled value saved in AC.
	//
	// example:
	//
	// Enabled
	SsoStatus *string `json:"ssoStatus,omitempty" xml:"ssoStatus,omitempty"`
}

func (s ListIdentityProvidersResponseBodyItemsConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProvidersResponseBodyItemsConfiguration) GoString() string {
	return s.String()
}

func (s *ListIdentityProvidersResponseBodyItemsConfiguration) GetBindingType() *string {
	return s.BindingType
}

func (s *ListIdentityProvidersResponseBodyItemsConfiguration) GetEntityId() *string {
	return s.EntityId
}

func (s *ListIdentityProvidersResponseBodyItemsConfiguration) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *ListIdentityProvidersResponseBodyItemsConfiguration) GetSsoStatus() *string {
	return s.SsoStatus
}

func (s *ListIdentityProvidersResponseBodyItemsConfiguration) SetBindingType(v string) *ListIdentityProvidersResponseBodyItemsConfiguration {
	s.BindingType = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItemsConfiguration) SetEntityId(v string) *ListIdentityProvidersResponseBodyItemsConfiguration {
	s.EntityId = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItemsConfiguration) SetLoginUrl(v string) *ListIdentityProvidersResponseBodyItemsConfiguration {
	s.LoginUrl = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItemsConfiguration) SetSsoStatus(v string) *ListIdentityProvidersResponseBodyItemsConfiguration {
	s.SsoStatus = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItemsConfiguration) Validate() error {
	return dara.Validate(s)
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
	// The SAML binding URI in the historical metadata snapshot. Use saml.bindingType for CustomSSO updates.
	//
	// example:
	//
	// urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect
	Binding *string `json:"binding,omitempty" xml:"binding,omitempty"`
	// The certificate string in the historical metadata snapshot. Use saml.x509Certificates to edit the configuration. Do not submit the old metadata.
	//
	// example:
	//
	// MIICwDCCAaigAwIBAgIULlsHLfcADyBax5ZtFlj+YGzZVo0wDQYJKoZIhvcNAQELBQAwGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMB4XDTI2MDEwMTAwMDAwMFoXDTM2MDEwMTAwMDAwMFowGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAto0DN0W1CrPRzK9QZ9A1aApb4Lg8oGlp2c0PbN3LokXe/HzwxCMXx65IDnrenttxzTwp9BgUnUN/3LRzWaOL2UQ/vNYNW9LFLkI1T64UPyI7eyi28YlALhb41xkYhEp7rz+MjqifBzVZeTDUxBYtbK2hAhbYEL1+udwsXEx0KB7xuQ/rbI2OKeqAEAD1tfYg7T/SM2wrXNYhx+o1HtchM4/0I8Okkk27JQ4y3ofzC5sTySmFF1Y7XcpSd+5WSyazedhkAieo8AB01DkJV6nnKfH8sH0EBsnH3FP8FhqCLrIeEkjX9B9OcIRLGjiZqx4u/ryW1QhMldnv+pFfNjWy/wIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQC1hSQG8AwtGOoX2gKBK4Hi5XBrBreij8WlLTUWld0FINQ0Rg84f/27CjDthywj6wgxCesCuOVFhpxQa1UacsP/Gcai1aKABtLWImcaWECRYK5gDkwhTMFKM9mXA26hH1l4v/UR2iTpEK9H78egC4W7bs2LDTpPryuaNKTqP2O0tMbNkAHB8yecz5LxHXnxadL9Egow25g2RClvpifWqDnuS9TSbKkIWLcPLjocJj2K9ILpBwhXiXi/BiSgyYTLiYvVvih9vX4s+hKGZNJc0j92YLiOHDF0XH1oxgdVjfTmRiedeb2VZJZElptD92z0AqFR7z2d5B+OqpPA2gjuMLV5
	Certificates *string `json:"certificates,omitempty" xml:"certificates,omitempty"`
	// The CorpId of the DingTalk enterprise. Required when the binding type is DingTalk.
	//
	// example:
	//
	// dingexamplecorpid01
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// Required for manual entry. It must be non-empty and contain a maximum of 2,048 characters. Existing configurations cannot be changed. This is different from the local entity ID returned in SP queries.
	//
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// Required for manual entry. It must be an absolute HTTPS URL without user credentials or fragments.
	//
	// example:
	//
	// https://idp.example.com/sso
	LoginUrl *string `json:"loginUrl,omitempty" xml:"loginUrl,omitempty"`
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

func (s *ListIdentityProvidersResponseBodyItemsMetadata) GetBinding() *string {
	return s.Binding
}

func (s *ListIdentityProvidersResponseBodyItemsMetadata) GetCertificates() *string {
	return s.Certificates
}

func (s *ListIdentityProvidersResponseBodyItemsMetadata) GetCorpId() *string {
	return s.CorpId
}

func (s *ListIdentityProvidersResponseBodyItemsMetadata) GetEntityId() *string {
	return s.EntityId
}

func (s *ListIdentityProvidersResponseBodyItemsMetadata) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *ListIdentityProvidersResponseBodyItemsMetadata) SetAppId(v string) *ListIdentityProvidersResponseBodyItemsMetadata {
	s.AppId = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItemsMetadata) SetAppKey(v string) *ListIdentityProvidersResponseBodyItemsMetadata {
	s.AppKey = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItemsMetadata) SetBinding(v string) *ListIdentityProvidersResponseBodyItemsMetadata {
	s.Binding = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItemsMetadata) SetCertificates(v string) *ListIdentityProvidersResponseBodyItemsMetadata {
	s.Certificates = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItemsMetadata) SetCorpId(v string) *ListIdentityProvidersResponseBodyItemsMetadata {
	s.CorpId = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItemsMetadata) SetEntityId(v string) *ListIdentityProvidersResponseBodyItemsMetadata {
	s.EntityId = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItemsMetadata) SetLoginUrl(v string) *ListIdentityProvidersResponseBodyItemsMetadata {
	s.LoginUrl = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItemsMetadata) Validate() error {
	return dara.Validate(s)
}

type ListIdentityProvidersResponseBodyItemsSaml struct {
	// HTTP-Redirect or HTTP-POST. Required for manual entry. Queries return the short name.
	//
	// example:
	//
	// HTTP-Redirect
	BindingType *string `json:"bindingType,omitempty" xml:"bindingType,omitempty"`
	// Required for manual entry. It must be non-empty and contain a maximum of 2,048 characters. Existing configurations cannot be changed. This is different from the local entity ID returned in SP queries.
	//
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// Required for manual entry. It must be an absolute HTTPS URL without user credentials or fragments.
	//
	// example:
	//
	// https://idp.example.com/sso
	LoginUrl *string `json:"loginUrl,omitempty" xml:"loginUrl,omitempty"`
	// A non-empty array is required for manual entry. Use PEM X.509 certificate text. The current validity period is verified. A maximum of two certificates are allowed after deduplication.
	//
	// example:
	//
	// ["-----BEGIN CERTIFICATE-----\\\\nMIICwDCCAaigAwIBAgIULlsHLfcADyBax5ZtFlj+YGzZVo0wDQYJKoZIhvcNAQEL\\\\nBQAwGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMB4XDTI2MDEwMTAwMDAwMFoX\\\\nDTM2MDEwMTAwMDAwMFowGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMIIBIjAN\\\\nBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAto0DN0W1CrPRzK9QZ9A1aApb4Lg8\\\\noGlp2c0PbN3LokXe/HzwxCMXx65IDnrenttxzTwp9BgUnUN/3LRzWaOL2UQ/vNYN\\\\nW9LFLkI1T64UPyI7eyi28YlALhb41xkYhEp7rz+MjqifBzVZeTDUxBYtbK2hAhbY\\\\nEL1+udwsXEx0KB7xuQ/rbI2OKeqAEAD1tfYg7T/SM2wrXNYhx+o1HtchM4/0I8Ok\\\\nkk27JQ4y3ofzC5sTySmFF1Y7XcpSd+5WSyazedhkAieo8AB01DkJV6nnKfH8sH0E\\\\nBsnH3FP8FhqCLrIeEkjX9B9OcIRLGjiZqx4u/ryW1QhMldnv+pFfNjWy/wIDAQAB\\\\nMA0GCSqGSIb3DQEBCwUAA4IBAQC1hSQG8AwtGOoX2gKBK4Hi5XBrBreij8WlLTUW\\\\nld0FINQ0Rg84f/27CjDthywj6wgxCesCuOVFhpxQa1UacsP/Gcai1aKABtLWImca\\\\nWECRYK5gDkwhTMFKM9mXA26hH1l4v/UR2iTpEK9H78egC4W7bs2LDTpPryuaNKTq\\\\nP2O0tMbNkAHB8yecz5LxHXnxadL9Egow25g2RClvpifWqDnuS9TSbKkIWLcPLjoc\\\\nJj2K9ILpBwhXiXi/BiSgyYTLiYvVvih9vX4s+hKGZNJc0j92YLiOHDF0XH1oxgdV\\\\njfTmRiedeb2VZJZElptD92z0AqFR7z2d5B+OqpPA2gjuMLV5\\\\n-----END CERTIFICATE-----"]
	X509Certificates []*string `json:"x509Certificates,omitempty" xml:"x509Certificates,omitempty" type:"Repeated"`
}

func (s ListIdentityProvidersResponseBodyItemsSaml) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProvidersResponseBodyItemsSaml) GoString() string {
	return s.String()
}

func (s *ListIdentityProvidersResponseBodyItemsSaml) GetBindingType() *string {
	return s.BindingType
}

func (s *ListIdentityProvidersResponseBodyItemsSaml) GetEntityId() *string {
	return s.EntityId
}

func (s *ListIdentityProvidersResponseBodyItemsSaml) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *ListIdentityProvidersResponseBodyItemsSaml) GetX509Certificates() []*string {
	return s.X509Certificates
}

func (s *ListIdentityProvidersResponseBodyItemsSaml) SetBindingType(v string) *ListIdentityProvidersResponseBodyItemsSaml {
	s.BindingType = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItemsSaml) SetEntityId(v string) *ListIdentityProvidersResponseBodyItemsSaml {
	s.EntityId = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItemsSaml) SetLoginUrl(v string) *ListIdentityProvidersResponseBodyItemsSaml {
	s.LoginUrl = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItemsSaml) SetX509Certificates(v []*string) *ListIdentityProvidersResponseBodyItemsSaml {
	s.X509Certificates = v
	return s
}

func (s *ListIdentityProvidersResponseBodyItemsSaml) Validate() error {
	return dara.Validate(s)
}

type ListIdentityProvidersResponseBodyItemsSyncScope struct {
	// When the mode is SelectedDepartments, it must contain at least one non-empty department ID obtained from ListIdentityProviderDepartments. When the mode is AllAuthorized, pass an empty array or omit this field.
	//
	// example:
	//
	// ["10001"]
	DepartmentIds []*string `json:"departmentIds,omitempty" xml:"departmentIds,omitempty" type:"Repeated"`
	// Valid values:
	//
	// - AllAuthorized: The full authorization scope, limited by the application address book authorization.
	//
	// - SelectedDepartments: The selected departments and all their sub-departments.
	//
	// You must provide a valid mode when explicitly submitting syncScope. The backend verifies the mode.
	//
	// example:
	//
	// SelectedDepartments
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
}

func (s ListIdentityProvidersResponseBodyItemsSyncScope) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProvidersResponseBodyItemsSyncScope) GoString() string {
	return s.String()
}

func (s *ListIdentityProvidersResponseBodyItemsSyncScope) GetDepartmentIds() []*string {
	return s.DepartmentIds
}

func (s *ListIdentityProvidersResponseBodyItemsSyncScope) GetMode() *string {
	return s.Mode
}

func (s *ListIdentityProvidersResponseBodyItemsSyncScope) SetDepartmentIds(v []*string) *ListIdentityProvidersResponseBodyItemsSyncScope {
	s.DepartmentIds = v
	return s
}

func (s *ListIdentityProvidersResponseBodyItemsSyncScope) SetMode(v string) *ListIdentityProvidersResponseBodyItemsSyncScope {
	s.Mode = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItemsSyncScope) Validate() error {
	return dara.Validate(s)
}
