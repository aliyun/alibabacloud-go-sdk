// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshProxyStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceMeshId(v string) *DescribeServiceMeshProxyStatusRequest
	GetServiceMeshId() *string
}

type DescribeServiceMeshProxyStatusRequest struct {
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca04bc38979214bf2882be79d39b4****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeServiceMeshProxyStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshProxyStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshProxyStatusRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeServiceMeshProxyStatusRequest) SetServiceMeshId(v string) *DescribeServiceMeshProxyStatusRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeServiceMeshProxyStatusRequest) Validate() error {
	return dara.Validate(s)
}
