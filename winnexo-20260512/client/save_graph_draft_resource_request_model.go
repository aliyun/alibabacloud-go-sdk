// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveGraphDraftResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *SaveGraphDraftResourceRequest
	GetAppToken() *string
	SetElementType(v string) *SaveGraphDraftResourceRequest
	GetElementType() *string
	SetGraphName(v string) *SaveGraphDraftResourceRequest
	GetGraphName() *string
	SetOpsUserCode(v string) *SaveGraphDraftResourceRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *SaveGraphDraftResourceRequest
	GetOpsUserId() *string
	SetResourceName(v string) *SaveGraphDraftResourceRequest
	GetResourceName() *string
	SetResourceType(v string) *SaveGraphDraftResourceRequest
	GetResourceType() *string
	SetTenantId(v string) *SaveGraphDraftResourceRequest
	GetTenantId() *string
	SetYamlEdit(v string) *SaveGraphDraftResourceRequest
	GetYamlEdit() *string
}

type SaveGraphDraftResourceRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The element type. Currently, only text is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// object_type
	ElementType *string `json:"elementType,omitempty" xml:"elementType,omitempty"`
	// The graph name.
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
	// The resource name.
	//
	// This parameter is required.
	//
	// example:
	//
	// customer
	ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	// The resource type.
	//
	// This parameter is set to **instance**, which indicates that the resource type is instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// object
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The original YAML text of the graph schema trimmed by READ permissions, with $ref references within the authorized subgraph retained.
	//
	// This parameter is required.
	//
	// example:
	//
	// name: customer\ndisplay_name: Customer
	YamlEdit *string `json:"yamlEdit,omitempty" xml:"yamlEdit,omitempty"`
}

func (s SaveGraphDraftResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveGraphDraftResourceRequest) GoString() string {
	return s.String()
}

func (s *SaveGraphDraftResourceRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *SaveGraphDraftResourceRequest) GetElementType() *string {
	return s.ElementType
}

func (s *SaveGraphDraftResourceRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *SaveGraphDraftResourceRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *SaveGraphDraftResourceRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *SaveGraphDraftResourceRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *SaveGraphDraftResourceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *SaveGraphDraftResourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SaveGraphDraftResourceRequest) GetYamlEdit() *string {
	return s.YamlEdit
}

func (s *SaveGraphDraftResourceRequest) SetAppToken(v string) *SaveGraphDraftResourceRequest {
	s.AppToken = &v
	return s
}

func (s *SaveGraphDraftResourceRequest) SetElementType(v string) *SaveGraphDraftResourceRequest {
	s.ElementType = &v
	return s
}

func (s *SaveGraphDraftResourceRequest) SetGraphName(v string) *SaveGraphDraftResourceRequest {
	s.GraphName = &v
	return s
}

func (s *SaveGraphDraftResourceRequest) SetOpsUserCode(v string) *SaveGraphDraftResourceRequest {
	s.OpsUserCode = &v
	return s
}

func (s *SaveGraphDraftResourceRequest) SetOpsUserId(v string) *SaveGraphDraftResourceRequest {
	s.OpsUserId = &v
	return s
}

func (s *SaveGraphDraftResourceRequest) SetResourceName(v string) *SaveGraphDraftResourceRequest {
	s.ResourceName = &v
	return s
}

func (s *SaveGraphDraftResourceRequest) SetResourceType(v string) *SaveGraphDraftResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *SaveGraphDraftResourceRequest) SetTenantId(v string) *SaveGraphDraftResourceRequest {
	s.TenantId = &v
	return s
}

func (s *SaveGraphDraftResourceRequest) SetYamlEdit(v string) *SaveGraphDraftResourceRequest {
	s.YamlEdit = &v
	return s
}

func (s *SaveGraphDraftResourceRequest) Validate() error {
	return dara.Validate(s)
}
