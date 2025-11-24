// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyServiceMeshNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ModifyServiceMeshNameRequest
	GetName() *string
	SetServiceMeshId(v string) *ModifyServiceMeshNameRequest
	GetServiceMeshId() *string
}

type ModifyServiceMeshNameRequest struct {
	// The new name of the ASM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-mesh
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cdd30a90a7cea480ebcc7ff****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s ModifyServiceMeshNameRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyServiceMeshNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyServiceMeshNameRequest) GetName() *string {
	return s.Name
}

func (s *ModifyServiceMeshNameRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *ModifyServiceMeshNameRequest) SetName(v string) *ModifyServiceMeshNameRequest {
	s.Name = &v
	return s
}

func (s *ModifyServiceMeshNameRequest) SetServiceMeshId(v string) *ModifyServiceMeshNameRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *ModifyServiceMeshNameRequest) Validate() error {
	return dara.Validate(s)
}
