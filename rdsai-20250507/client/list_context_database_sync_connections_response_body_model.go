// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncConnectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListContextDatabaseSyncConnectionsResponseBody
	GetCode() *string
	SetData(v *ListContextDatabaseSyncConnectionsResponseBodyData) *ListContextDatabaseSyncConnectionsResponseBody
	GetData() *ListContextDatabaseSyncConnectionsResponseBodyData
	SetHttpStatusCode(v int64) *ListContextDatabaseSyncConnectionsResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *ListContextDatabaseSyncConnectionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListContextDatabaseSyncConnectionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListContextDatabaseSyncConnectionsResponseBody
	GetSuccess() *bool
}

type ListContextDatabaseSyncConnectionsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListContextDatabaseSyncConnectionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 1
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// 示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// id-001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListContextDatabaseSyncConnectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncConnectionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListContextDatabaseSyncConnectionsResponseBody) GetData() *ListContextDatabaseSyncConnectionsResponseBodyData {
	return s.Data
}

func (s *ListContextDatabaseSyncConnectionsResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ListContextDatabaseSyncConnectionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListContextDatabaseSyncConnectionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListContextDatabaseSyncConnectionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListContextDatabaseSyncConnectionsResponseBody) SetCode(v string) *ListContextDatabaseSyncConnectionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListContextDatabaseSyncConnectionsResponseBody) SetData(v *ListContextDatabaseSyncConnectionsResponseBodyData) *ListContextDatabaseSyncConnectionsResponseBody {
	s.Data = v
	return s
}

func (s *ListContextDatabaseSyncConnectionsResponseBody) SetHttpStatusCode(v int64) *ListContextDatabaseSyncConnectionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListContextDatabaseSyncConnectionsResponseBody) SetMessage(v string) *ListContextDatabaseSyncConnectionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListContextDatabaseSyncConnectionsResponseBody) SetRequestId(v string) *ListContextDatabaseSyncConnectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContextDatabaseSyncConnectionsResponseBody) SetSuccess(v bool) *ListContextDatabaseSyncConnectionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListContextDatabaseSyncConnectionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListContextDatabaseSyncConnectionsResponseBodyData struct {
	// The list of datasource configurations.
	//
	// example:
	//
	// [{"Id":"conn-as1llqmkol****", "PluginId":"dingtalk", "DisplayName":"R&D Team DingTalk Datasource", "AuthKind":"app"}]
	Connections []*ListContextDatabaseSyncConnectionsResponseBodyDataConnections `json:"Connections,omitempty" xml:"Connections,omitempty" type:"Repeated"`
	// The page number. Pages start from 1.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListContextDatabaseSyncConnectionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncConnectionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyData) GetConnections() []*ListContextDatabaseSyncConnectionsResponseBodyDataConnections {
	return s.Connections
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyData) SetConnections(v []*ListContextDatabaseSyncConnectionsResponseBodyDataConnections) *ListContextDatabaseSyncConnectionsResponseBodyData {
	s.Connections = v
	return s
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyData) SetPage(v int32) *ListContextDatabaseSyncConnectionsResponseBodyData {
	s.Page = &v
	return s
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyData) SetPageSize(v int32) *ListContextDatabaseSyncConnectionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyData) SetTotal(v int32) *ListContextDatabaseSyncConnectionsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyData) Validate() error {
	if s.Connections != nil {
		for _, item := range s.Connections {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListContextDatabaseSyncConnectionsResponseBodyDataConnections struct {
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
	// The time when the datasource configuration was created.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The display name.
	//
	// example:
	//
	// R&D Team DingTalk Datasource
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The connection ID, which is a UUID generated by the system.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The plug-in ID.
	//
	// example:
	//
	// dingtalk
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The remediation suggestion.
	//
	// example:
	//
	// Re-authorize the datasource configuration
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
	// The time when the credential was last updated.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	SecretUpdatedAt *string `json:"SecretUpdatedAt,omitempty" xml:"SecretUpdatedAt,omitempty"`
	// The time when the datasource configuration was last updated.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	// The time when the datasource configuration was last verified.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	VerifiedAt *string `json:"VerifiedAt,omitempty" xml:"VerifiedAt,omitempty"`
}

func (s ListContextDatabaseSyncConnectionsResponseBodyDataConnections) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncConnectionsResponseBodyDataConnections) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyDataConnections) GetAuthKind() *string {
	return s.AuthKind
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyDataConnections) GetAuthState() *string {
	return s.AuthState
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyDataConnections) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyDataConnections) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyDataConnections) GetId() *string {
	return s.Id
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyDataConnections) GetPluginId() *string {
	return s.PluginId
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyDataConnections) GetRemediation() *string {
	return s.Remediation
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyDataConnections) GetScopeReport() map[string]interface{} {
	return s.ScopeReport
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyDataConnections) GetSecretFingerprint() *string {
	return s.SecretFingerprint
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyDataConnections) GetSecretUpdatedAt() *string {
	return s.SecretUpdatedAt
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyDataConnections) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyDataConnections) GetVerifiedAt() *string {
	return s.VerifiedAt
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyDataConnections) SetAuthKind(v string) *ListContextDatabaseSyncConnectionsResponseBodyDataConnections {
	s.AuthKind = &v
	return s
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyDataConnections) SetAuthState(v string) *ListContextDatabaseSyncConnectionsResponseBodyDataConnections {
	s.AuthState = &v
	return s
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyDataConnections) SetCreatedAt(v string) *ListContextDatabaseSyncConnectionsResponseBodyDataConnections {
	s.CreatedAt = &v
	return s
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyDataConnections) SetDisplayName(v string) *ListContextDatabaseSyncConnectionsResponseBodyDataConnections {
	s.DisplayName = &v
	return s
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyDataConnections) SetId(v string) *ListContextDatabaseSyncConnectionsResponseBodyDataConnections {
	s.Id = &v
	return s
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyDataConnections) SetPluginId(v string) *ListContextDatabaseSyncConnectionsResponseBodyDataConnections {
	s.PluginId = &v
	return s
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyDataConnections) SetRemediation(v string) *ListContextDatabaseSyncConnectionsResponseBodyDataConnections {
	s.Remediation = &v
	return s
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyDataConnections) SetScopeReport(v map[string]interface{}) *ListContextDatabaseSyncConnectionsResponseBodyDataConnections {
	s.ScopeReport = v
	return s
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyDataConnections) SetSecretFingerprint(v string) *ListContextDatabaseSyncConnectionsResponseBodyDataConnections {
	s.SecretFingerprint = &v
	return s
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyDataConnections) SetSecretUpdatedAt(v string) *ListContextDatabaseSyncConnectionsResponseBodyDataConnections {
	s.SecretUpdatedAt = &v
	return s
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyDataConnections) SetUpdatedAt(v string) *ListContextDatabaseSyncConnectionsResponseBodyDataConnections {
	s.UpdatedAt = &v
	return s
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyDataConnections) SetVerifiedAt(v string) *ListContextDatabaseSyncConnectionsResponseBodyDataConnections {
	s.VerifiedAt = &v
	return s
}

func (s *ListContextDatabaseSyncConnectionsResponseBodyDataConnections) Validate() error {
	return dara.Validate(s)
}
