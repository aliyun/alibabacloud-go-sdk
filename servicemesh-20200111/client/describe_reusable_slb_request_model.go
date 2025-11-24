// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReusableSlbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetK8sClusterId(v string) *DescribeReusableSlbRequest
	GetK8sClusterId() *string
	SetLbType(v string) *DescribeReusableSlbRequest
	GetLbType() *string
	SetNetworkType(v string) *DescribeReusableSlbRequest
	GetNetworkType() *string
	SetServiceMeshId(v string) *DescribeReusableSlbRequest
	GetServiceMeshId() *string
}

type DescribeReusableSlbRequest struct {
	// The ID of the cluster on the data plane. If you specify this parameter, you cannot specify ServiceMeshId. This parameter and ServiceMeshId cannot be left empty at the same time.
	//
	// example:
	//
	// ca2cfe41fefeb489d9b9dba18a7c5****
	K8sClusterId *string `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	// Load balancer type, value:
	//
	// 	- `clb`: Classic Load Balancer
	//
	// 	- `nlb`: Network Load Balancer
	//
	// default is `clb`.
	//
	// example:
	//
	// clb
	LbType *string `json:"LbType,omitempty" xml:"LbType,omitempty"`
	// The network type of the SLB instance. Valid values:
	//
	// 	- `intranet`
	//
	// 	- `internet`
	//
	// This parameter is required.
	//
	// example:
	//
	// internet
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The ID of the Service Mesh (ASM) instance. If you specify this parameter, you cannot specify K8sClusterId. This parameter and K8sClusterId cannot be left empty at the same time.
	//
	// example:
	//
	// cb8963379255149cb98c8686f274x****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeReusableSlbRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeReusableSlbRequest) GoString() string {
	return s.String()
}

func (s *DescribeReusableSlbRequest) GetK8sClusterId() *string {
	return s.K8sClusterId
}

func (s *DescribeReusableSlbRequest) GetLbType() *string {
	return s.LbType
}

func (s *DescribeReusableSlbRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeReusableSlbRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeReusableSlbRequest) SetK8sClusterId(v string) *DescribeReusableSlbRequest {
	s.K8sClusterId = &v
	return s
}

func (s *DescribeReusableSlbRequest) SetLbType(v string) *DescribeReusableSlbRequest {
	s.LbType = &v
	return s
}

func (s *DescribeReusableSlbRequest) SetNetworkType(v string) *DescribeReusableSlbRequest {
	s.NetworkType = &v
	return s
}

func (s *DescribeReusableSlbRequest) SetServiceMeshId(v string) *DescribeReusableSlbRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeReusableSlbRequest) Validate() error {
	return dara.Validate(s)
}
