// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceMeshResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateServiceMeshResponseBody
	GetRequestId() *string
	SetServiceMeshId(v string) *CreateServiceMeshResponseBody
	GetServiceMeshId() *string
}

type CreateServiceMeshResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BD65C0AD-D3C6-48D3-8D93-38D2015C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the ASM instance.
	//
	// example:
	//
	// c08ba3fd1e6484b0f8cc1ad8fe10d****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s CreateServiceMeshResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceMeshResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceMeshResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceMeshResponseBody) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *CreateServiceMeshResponseBody) SetRequestId(v string) *CreateServiceMeshResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceMeshResponseBody) SetServiceMeshId(v string) *CreateServiceMeshResponseBody {
	s.ServiceMeshId = &v
	return s
}

func (s *CreateServiceMeshResponseBody) Validate() error {
	return dara.Validate(s)
}
