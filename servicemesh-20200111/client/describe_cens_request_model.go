// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCensRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceMeshId(v string) *DescribeCensRequest
	GetServiceMeshId() *string
}

type DescribeCensRequest struct {
	// The ID of the ASM instance.
	//
	// example:
	//
	// ce134b0727aa2492db69f6c3880e1****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeCensRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCensRequest) GoString() string {
	return s.String()
}

func (s *DescribeCensRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeCensRequest) SetServiceMeshId(v string) *DescribeCensRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeCensRequest) Validate() error {
	return dara.Validate(s)
}
