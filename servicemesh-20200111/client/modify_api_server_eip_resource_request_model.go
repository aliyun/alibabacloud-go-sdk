// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApiServerEipResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiServerEipId(v string) *ModifyApiServerEipResourceRequest
	GetApiServerEipId() *string
	SetOperation(v string) *ModifyApiServerEipResourceRequest
	GetOperation() *string
	SetServiceMeshId(v string) *ModifyApiServerEipResourceRequest
	GetServiceMeshId() *string
}

type ModifyApiServerEipResourceRequest struct {
	// The ID of the EIP.
	//
	// example:
	//
	// eip-bp1adu9jegmxnaoq****
	ApiServerEipId *string `json:"ApiServerEipId,omitempty" xml:"ApiServerEipId,omitempty"`
	// The type of the operation. Valid values:
	//
	// 	- `UnBindEip`: disassociates an EIP from the API server.
	//
	// 	- `BindEip`: associates an EIP with the API server.
	//
	// example:
	//
	// BindEip
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The ID of the Alibaba Cloud Service Mesh (ASM) instance.
	//
	// example:
	//
	// cb8963379255149cb98c8686f274x****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s ModifyApiServerEipResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApiServerEipResourceRequest) GoString() string {
	return s.String()
}

func (s *ModifyApiServerEipResourceRequest) GetApiServerEipId() *string {
	return s.ApiServerEipId
}

func (s *ModifyApiServerEipResourceRequest) GetOperation() *string {
	return s.Operation
}

func (s *ModifyApiServerEipResourceRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *ModifyApiServerEipResourceRequest) SetApiServerEipId(v string) *ModifyApiServerEipResourceRequest {
	s.ApiServerEipId = &v
	return s
}

func (s *ModifyApiServerEipResourceRequest) SetOperation(v string) *ModifyApiServerEipResourceRequest {
	s.Operation = &v
	return s
}

func (s *ModifyApiServerEipResourceRequest) SetServiceMeshId(v string) *ModifyApiServerEipResourceRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *ModifyApiServerEipResourceRequest) Validate() error {
	return dara.Validate(s)
}
