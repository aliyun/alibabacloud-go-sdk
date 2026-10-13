// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseSyncConnectionAuthorizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetContextDatabaseSyncConnectionAuthorizationResponseBody
	GetCode() *string
	SetData(v *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) *GetContextDatabaseSyncConnectionAuthorizationResponseBody
	GetData() *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData
	SetHttpStatusCode(v int64) *GetContextDatabaseSyncConnectionAuthorizationResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *GetContextDatabaseSyncConnectionAuthorizationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetContextDatabaseSyncConnectionAuthorizationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetContextDatabaseSyncConnectionAuthorizationResponseBody
	GetSuccess() *bool
}

type GetContextDatabaseSyncConnectionAuthorizationResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Sample value
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 1
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message.
	//
	// example:
	//
	// Sample value
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// id-001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetContextDatabaseSyncConnectionAuthorizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncConnectionAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBody) GetData() *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData {
	return s.Data
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBody) SetCode(v string) *GetContextDatabaseSyncConnectionAuthorizationResponseBody {
	s.Code = &v
	return s
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBody) SetData(v *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) *GetContextDatabaseSyncConnectionAuthorizationResponseBody {
	s.Data = v
	return s
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBody) SetHttpStatusCode(v int64) *GetContextDatabaseSyncConnectionAuthorizationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBody) SetMessage(v string) *GetContextDatabaseSyncConnectionAuthorizationResponseBody {
	s.Message = &v
	return s
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBody) SetRequestId(v string) *GetContextDatabaseSyncConnectionAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBody) SetSuccess(v bool) *GetContextDatabaseSyncConnectionAuthorizationResponseBody {
	s.Success = &v
	return s
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetContextDatabaseSyncConnectionAuthorizationResponseBodyData struct {
	// The authorization method. Valid values: app, pasted-url, static-secret, and none.
	//
	// example:
	//
	// app
	AuthKind *string `json:"AuthKind,omitempty" xml:"AuthKind,omitempty"`
	// The authorization status.
	//
	// example:
	//
	// ready
	AuthState *string `json:"AuthState,omitempty" xml:"AuthState,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The display name.
	//
	// example:
	//
	// R&D team DingTalk data source
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The connection ID, which is a system-generated UUID.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The plugin ID.
	//
	// example:
	//
	// dingtalk
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The remediation suggestion.
	//
	// example:
	//
	// Re-authorize the data source
	Remediation *string `json:"Remediation,omitempty" xml:"Remediation,omitempty"`
	// The authorization scope report.
	//
	// example:
	//
	// {}
	ScopeReport map[string]interface{} `json:"ScopeReport,omitempty" xml:"ScopeReport,omitempty"`
	// The credential fingerprint.
	//
	// example:
	//
	// sha256:4b3b7b2e5f6a1c9d8e0f1234567890abcdef1234567890abcdef1234567890ab
	SecretFingerprint *string `json:"SecretFingerprint,omitempty" xml:"SecretFingerprint,omitempty"`
	// The credential update time.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	SecretUpdatedAt *string `json:"SecretUpdatedAt,omitempty" xml:"SecretUpdatedAt,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	// The verification time.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	VerifiedAt *string `json:"VerifiedAt,omitempty" xml:"VerifiedAt,omitempty"`
}

func (s GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) GetAuthKind() *string {
	return s.AuthKind
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) GetAuthState() *string {
	return s.AuthState
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) GetPluginId() *string {
	return s.PluginId
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) GetRemediation() *string {
	return s.Remediation
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) GetScopeReport() map[string]interface{} {
	return s.ScopeReport
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) GetSecretFingerprint() *string {
	return s.SecretFingerprint
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) GetSecretUpdatedAt() *string {
	return s.SecretUpdatedAt
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) GetVerifiedAt() *string {
	return s.VerifiedAt
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) SetAuthKind(v string) *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData {
	s.AuthKind = &v
	return s
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) SetAuthState(v string) *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData {
	s.AuthState = &v
	return s
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) SetCreatedAt(v string) *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) SetDisplayName(v string) *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) SetId(v string) *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) SetPluginId(v string) *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData {
	s.PluginId = &v
	return s
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) SetRemediation(v string) *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData {
	s.Remediation = &v
	return s
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) SetScopeReport(v map[string]interface{}) *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData {
	s.ScopeReport = v
	return s
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) SetSecretFingerprint(v string) *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData {
	s.SecretFingerprint = &v
	return s
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) SetSecretUpdatedAt(v string) *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData {
	s.SecretUpdatedAt = &v
	return s
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) SetUpdatedAt(v string) *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) SetVerifiedAt(v string) *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData {
	s.VerifiedAt = &v
	return s
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
