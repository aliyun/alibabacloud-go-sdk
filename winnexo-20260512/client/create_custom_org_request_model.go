// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomOrgRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreateCustomOrgRequest
	GetAppToken() *string
	SetCorpId(v string) *CreateCustomOrgRequest
	GetCorpId() *string
	SetCorpName(v string) *CreateCustomOrgRequest
	GetCorpName() *string
	SetOpsUserCode(v string) *CreateCustomOrgRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateCustomOrgRequest
	GetOpsUserId() *string
	SetTenantId(v string) *CreateCustomOrgRequest
	GetTenantId() *string
}

type CreateCustomOrgRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The corpId of the activated enterprise.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleCorpId
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// The organization name.
	//
	// example:
	//
	// string_value
	CorpName *string `json:"corpName,omitempty" xml:"corpName,omitempty"`
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
	// The tenant ID.
	//
	// example:
	//
	// 692318833855074
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateCustomOrgRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomOrgRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomOrgRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreateCustomOrgRequest) GetCorpId() *string {
	return s.CorpId
}

func (s *CreateCustomOrgRequest) GetCorpName() *string {
	return s.CorpName
}

func (s *CreateCustomOrgRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateCustomOrgRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreateCustomOrgRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateCustomOrgRequest) SetAppToken(v string) *CreateCustomOrgRequest {
	s.AppToken = &v
	return s
}

func (s *CreateCustomOrgRequest) SetCorpId(v string) *CreateCustomOrgRequest {
	s.CorpId = &v
	return s
}

func (s *CreateCustomOrgRequest) SetCorpName(v string) *CreateCustomOrgRequest {
	s.CorpName = &v
	return s
}

func (s *CreateCustomOrgRequest) SetOpsUserCode(v string) *CreateCustomOrgRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateCustomOrgRequest) SetOpsUserId(v string) *CreateCustomOrgRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreateCustomOrgRequest) SetTenantId(v string) *CreateCustomOrgRequest {
	s.TenantId = &v
	return s
}

func (s *CreateCustomOrgRequest) Validate() error {
	return dara.Validate(s)
}
