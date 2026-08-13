// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGraphSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGraphName(v string) *GetGraphSchemaRequest
	GetGraphName() *string
	SetTenantId(v string) *GetGraphSchemaRequest
	GetTenantId() *string
}

type GetGraphSchemaRequest struct {
	// 图谱名称，可先调用 listGraphs 获取
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// This parameter is required.
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetGraphSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGraphSchemaRequest) GoString() string {
	return s.String()
}

func (s *GetGraphSchemaRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *GetGraphSchemaRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetGraphSchemaRequest) SetGraphName(v string) *GetGraphSchemaRequest {
	s.GraphName = &v
	return s
}

func (s *GetGraphSchemaRequest) SetTenantId(v string) *GetGraphSchemaRequest {
	s.TenantId = &v
	return s
}

func (s *GetGraphSchemaRequest) Validate() error {
	return dara.Validate(s)
}
