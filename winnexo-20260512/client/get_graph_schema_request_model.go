// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGraphSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetGraphSchemaRequest
	GetAppToken() *string
	SetGraphName(v string) *GetGraphSchemaRequest
	GetGraphName() *string
	SetOpsUserCode(v string) *GetGraphSchemaRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetGraphSchemaRequest
	GetOpsUserId() *string
	SetTenantId(v string) *GetGraphSchemaRequest
	GetTenantId() *string
}

type GetGraphSchemaRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The graph name. Call listGraphs first to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
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
	// The effective tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 21577
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetGraphSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGraphSchemaRequest) GoString() string {
	return s.String()
}

func (s *GetGraphSchemaRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetGraphSchemaRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *GetGraphSchemaRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetGraphSchemaRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetGraphSchemaRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetGraphSchemaRequest) SetAppToken(v string) *GetGraphSchemaRequest {
	s.AppToken = &v
	return s
}

func (s *GetGraphSchemaRequest) SetGraphName(v string) *GetGraphSchemaRequest {
	s.GraphName = &v
	return s
}

func (s *GetGraphSchemaRequest) SetOpsUserCode(v string) *GetGraphSchemaRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetGraphSchemaRequest) SetOpsUserId(v string) *GetGraphSchemaRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetGraphSchemaRequest) SetTenantId(v string) *GetGraphSchemaRequest {
	s.TenantId = &v
	return s
}

func (s *GetGraphSchemaRequest) Validate() error {
	return dara.Validate(s)
}
