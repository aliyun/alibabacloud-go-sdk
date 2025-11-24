// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegisteredServiceNamespacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceMeshId(v string) *GetRegisteredServiceNamespacesRequest
	GetServiceMeshId() *string
}

type GetRegisteredServiceNamespacesRequest struct {
	// The ID of the ASM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cb8963379255149cb98c8686f274x****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s GetRegisteredServiceNamespacesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRegisteredServiceNamespacesRequest) GoString() string {
	return s.String()
}

func (s *GetRegisteredServiceNamespacesRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *GetRegisteredServiceNamespacesRequest) SetServiceMeshId(v string) *GetRegisteredServiceNamespacesRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *GetRegisteredServiceNamespacesRequest) Validate() error {
	return dara.Validate(s)
}
