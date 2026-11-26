// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetSourceRequest
	GetAppToken() *string
	SetIncludeDetails(v bool) *GetSourceRequest
	GetIncludeDetails() *bool
	SetOpsUserCode(v string) *GetSourceRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetSourceRequest
	GetOpsUserId() *string
	SetSourceId(v string) *GetSourceRequest
	GetSourceId() *string
	SetTenantId(v string) *GetSourceRequest
	GetTenantId() *string
}

type GetSourceRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// Specifies whether to return large detail fields (settings, notes, structuredTables, and unstructuredDocs). Default value: false. Only metadata is returned.
	//
	// example:
	//
	// false
	IncludeDetails *bool `json:"includeDetails,omitempty" xml:"includeDetails,omitempty"`
	// 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The primary ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The ID of the tenant to which the task belongs.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSourceRequest) GoString() string {
	return s.String()
}

func (s *GetSourceRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetSourceRequest) GetIncludeDetails() *bool {
	return s.IncludeDetails
}

func (s *GetSourceRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetSourceRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *GetSourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetSourceRequest) SetAppToken(v string) *GetSourceRequest {
	s.AppToken = &v
	return s
}

func (s *GetSourceRequest) SetIncludeDetails(v bool) *GetSourceRequest {
	s.IncludeDetails = &v
	return s
}

func (s *GetSourceRequest) SetOpsUserCode(v string) *GetSourceRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetSourceRequest) SetOpsUserId(v string) *GetSourceRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetSourceRequest) SetSourceId(v string) *GetSourceRequest {
	s.SourceId = &v
	return s
}

func (s *GetSourceRequest) SetTenantId(v string) *GetSourceRequest {
	s.TenantId = &v
	return s
}

func (s *GetSourceRequest) Validate() error {
	return dara.Validate(s)
}
