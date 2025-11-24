// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceMeshId(v string) *DescribeServiceMeshLogsRequest
	GetServiceMeshId() *string
}

type DescribeServiceMeshLogsRequest struct {
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca04bc38979214bf2882be79d39b4****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeServiceMeshLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshLogsRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeServiceMeshLogsRequest) SetServiceMeshId(v string) *DescribeServiceMeshLogsRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeServiceMeshLogsRequest) Validate() error {
	return dara.Validate(s)
}
