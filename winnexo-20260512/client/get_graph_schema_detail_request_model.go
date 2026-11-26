// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGraphSchemaDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetGraphSchemaDetailRequest
	GetAppToken() *string
	SetGraphName(v string) *GetGraphSchemaDetailRequest
	GetGraphName() *string
	SetOpsUserCode(v string) *GetGraphSchemaDetailRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetGraphSchemaDetailRequest
	GetOpsUserId() *string
	SetTenantId(v string) *GetGraphSchemaDetailRequest
	GetTenantId() *string
}

type GetGraphSchemaDetailRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The name of the graph.
	//
	// This parameter is required.
	//
	// example:
	//
	// crm_graph
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
	// The tenant ID. This is a common parameter. You can pass this parameter explicitly by using `--tenant-id` in winnexo-cli.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetGraphSchemaDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGraphSchemaDetailRequest) GoString() string {
	return s.String()
}

func (s *GetGraphSchemaDetailRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetGraphSchemaDetailRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *GetGraphSchemaDetailRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetGraphSchemaDetailRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetGraphSchemaDetailRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetGraphSchemaDetailRequest) SetAppToken(v string) *GetGraphSchemaDetailRequest {
	s.AppToken = &v
	return s
}

func (s *GetGraphSchemaDetailRequest) SetGraphName(v string) *GetGraphSchemaDetailRequest {
	s.GraphName = &v
	return s
}

func (s *GetGraphSchemaDetailRequest) SetOpsUserCode(v string) *GetGraphSchemaDetailRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetGraphSchemaDetailRequest) SetOpsUserId(v string) *GetGraphSchemaDetailRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetGraphSchemaDetailRequest) SetTenantId(v string) *GetGraphSchemaDetailRequest {
	s.TenantId = &v
	return s
}

func (s *GetGraphSchemaDetailRequest) Validate() error {
	return dara.Validate(s)
}
