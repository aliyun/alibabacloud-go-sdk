// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCCMVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceMeshId(v string) *DescribeCCMVersionRequest
	GetServiceMeshId() *string
}

type DescribeCCMVersionRequest struct {
	// The ID of the Service Mesh (ASM) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c08ba3fd1e6484b0f8cc1ad8fe10d****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeCCMVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCCMVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeCCMVersionRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeCCMVersionRequest) SetServiceMeshId(v string) *DescribeCCMVersionRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeCCMVersionRequest) Validate() error {
	return dara.Validate(s)
}
