// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshAdditionalStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckMode(v string) *DescribeServiceMeshAdditionalStatusRequest
	GetCheckMode() *string
	SetServiceMeshId(v string) *DescribeServiceMeshAdditionalStatusRequest
	GetServiceMeshId() *string
}

type DescribeServiceMeshAdditionalStatusRequest struct {
	// The check mode of the ASM instance. Valid values:
	//
	// 	- `normal`: checks the Server Load Balancer (SLB) instances created for exposing the API server and Istio Pilot, audit logs, and installation of Logtail for clusters on the data plane.
	//
	// 	- `full`: checks control plane logs, access logs, security groups, and the elastic IP addresses (EIPs) of the API server in addition to the check items in normal mode.
	//
	// example:
	//
	// full
	CheckMode *string `json:"CheckMode,omitempty" xml:"CheckMode,omitempty"`
	// The ID of the ASM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca04bc38979214bf2882be79d39b4****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeServiceMeshAdditionalStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshAdditionalStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshAdditionalStatusRequest) GetCheckMode() *string {
	return s.CheckMode
}

func (s *DescribeServiceMeshAdditionalStatusRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeServiceMeshAdditionalStatusRequest) SetCheckMode(v string) *DescribeServiceMeshAdditionalStatusRequest {
	s.CheckMode = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusRequest) SetServiceMeshId(v string) *DescribeServiceMeshAdditionalStatusRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusRequest) Validate() error {
	return dara.Validate(s)
}
