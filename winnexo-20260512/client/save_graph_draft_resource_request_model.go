// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveGraphDraftResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetElementType(v string) *SaveGraphDraftResourceRequest
	GetElementType() *string
	SetGraphName(v string) *SaveGraphDraftResourceRequest
	GetGraphName() *string
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
	// 资源小类：resourceType=object 时固定 object_type；resourceType=element 时为 indicator / logic / process / rule / analysis 之一
	//
	// This parameter is required.
	//
	// example:
	//
	// object_type
	ElementType *string `json:"elementType,omitempty" xml:"elementType,omitempty"`
	// 图谱名称，须已存在（active 记录）
	//
	// This parameter is required.
	//
	// example:
	//
	// crm_graph
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// 资源名（创建后不可改名，底层校验）
	//
	// This parameter is required.
	//
	// example:
	//
	// customer
	ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	// 资源大类：object（对象）/ element（业务元素）
	//
	// This parameter is required.
	//
	// example:
	//
	// object
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 单资源 YAML 文本
	//
	// This parameter is required.
	//
	// example:
	//
	// name: customer\\ndisplay_name: 客户
	YamlEdit *string `json:"yamlEdit,omitempty" xml:"yamlEdit,omitempty"`
}

func (s SaveGraphDraftResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveGraphDraftResourceRequest) GoString() string {
	return s.String()
}

func (s *SaveGraphDraftResourceRequest) GetElementType() *string {
	return s.ElementType
}

func (s *SaveGraphDraftResourceRequest) GetGraphName() *string {
	return s.GraphName
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

func (s *SaveGraphDraftResourceRequest) SetElementType(v string) *SaveGraphDraftResourceRequest {
	s.ElementType = &v
	return s
}

func (s *SaveGraphDraftResourceRequest) SetGraphName(v string) *SaveGraphDraftResourceRequest {
	s.GraphName = &v
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
