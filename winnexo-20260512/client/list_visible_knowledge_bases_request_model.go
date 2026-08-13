// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVisibleKnowledgeBasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperatingObjectName(v string) *ListVisibleKnowledgeBasesRequest
	GetOperatingObjectName() *string
	SetTenantId(v string) *ListVisibleKnowledgeBasesRequest
	GetTenantId() *string
}

type ListVisibleKnowledgeBasesRequest struct {
	// 数字员工名称（运营对象 name）
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListVisibleKnowledgeBasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVisibleKnowledgeBasesRequest) GoString() string {
	return s.String()
}

func (s *ListVisibleKnowledgeBasesRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListVisibleKnowledgeBasesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListVisibleKnowledgeBasesRequest) SetOperatingObjectName(v string) *ListVisibleKnowledgeBasesRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ListVisibleKnowledgeBasesRequest) SetTenantId(v string) *ListVisibleKnowledgeBasesRequest {
	s.TenantId = &v
	return s
}

func (s *ListVisibleKnowledgeBasesRequest) Validate() error {
	return dara.Validate(s)
}
