// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGraphSchemaDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGraphName(v string) *GetGraphSchemaDetailRequest
	GetGraphName() *string
	SetTenantId(v string) *GetGraphSchemaDetailRequest
	GetTenantId() *string
}

type GetGraphSchemaDetailRequest struct {
	// The name of the graph.
	//
	// This parameter is required.
	//
	// example:
	//
	// crm_graph
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
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

func (s *GetGraphSchemaDetailRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *GetGraphSchemaDetailRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetGraphSchemaDetailRequest) SetGraphName(v string) *GetGraphSchemaDetailRequest {
	s.GraphName = &v
	return s
}

func (s *GetGraphSchemaDetailRequest) SetTenantId(v string) *GetGraphSchemaDetailRequest {
	s.TenantId = &v
	return s
}

func (s *GetGraphSchemaDetailRequest) Validate() error {
	return dara.Validate(s)
}
