// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearOperatingObjectFavoritesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ClearOperatingObjectFavoritesRequest
	GetAppToken() *string
	SetGraphName(v string) *ClearOperatingObjectFavoritesRequest
	GetGraphName() *string
	SetObjectType(v string) *ClearOperatingObjectFavoritesRequest
	GetObjectType() *string
	SetOperatingObjectName(v string) *ClearOperatingObjectFavoritesRequest
	GetOperatingObjectName() *string
	SetOpsUserCode(v string) *ClearOperatingObjectFavoritesRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ClearOperatingObjectFavoritesRequest
	GetOpsUserId() *string
	SetTenantId(v string) *ClearOperatingObjectFavoritesRequest
	GetTenantId() *string
}

type ClearOperatingObjectFavoritesRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The graph name. You can call listGraphs to obtain the value.
	//
	// This parameter is required.
	//
	// example:
	//
	// crm
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The object type, such as customer. This parameter has a value when type is set to mention.
	//
	// This parameter is required.
	//
	// example:
	//
	// contract
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// The operating object name, such as customer_1.
	//
	// This parameter is required.
	//
	// example:
	//
	// customer_assistant
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
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
	// The tenant ID to take effect.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ClearOperatingObjectFavoritesRequest) String() string {
	return dara.Prettify(s)
}

func (s ClearOperatingObjectFavoritesRequest) GoString() string {
	return s.String()
}

func (s *ClearOperatingObjectFavoritesRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ClearOperatingObjectFavoritesRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *ClearOperatingObjectFavoritesRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *ClearOperatingObjectFavoritesRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ClearOperatingObjectFavoritesRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ClearOperatingObjectFavoritesRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ClearOperatingObjectFavoritesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ClearOperatingObjectFavoritesRequest) SetAppToken(v string) *ClearOperatingObjectFavoritesRequest {
	s.AppToken = &v
	return s
}

func (s *ClearOperatingObjectFavoritesRequest) SetGraphName(v string) *ClearOperatingObjectFavoritesRequest {
	s.GraphName = &v
	return s
}

func (s *ClearOperatingObjectFavoritesRequest) SetObjectType(v string) *ClearOperatingObjectFavoritesRequest {
	s.ObjectType = &v
	return s
}

func (s *ClearOperatingObjectFavoritesRequest) SetOperatingObjectName(v string) *ClearOperatingObjectFavoritesRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ClearOperatingObjectFavoritesRequest) SetOpsUserCode(v string) *ClearOperatingObjectFavoritesRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ClearOperatingObjectFavoritesRequest) SetOpsUserId(v string) *ClearOperatingObjectFavoritesRequest {
	s.OpsUserId = &v
	return s
}

func (s *ClearOperatingObjectFavoritesRequest) SetTenantId(v string) *ClearOperatingObjectFavoritesRequest {
	s.TenantId = &v
	return s
}

func (s *ClearOperatingObjectFavoritesRequest) Validate() error {
	return dara.Validate(s)
}
