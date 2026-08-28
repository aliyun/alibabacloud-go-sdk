// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListK8sClusterSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVpcId(v string) *ListK8sClusterSourcesRequest
	GetVpcId() *string
}

type ListK8sClusterSourcesRequest struct {
	// example:
	//
	// vpc-xxxx
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListK8sClusterSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListK8sClusterSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListK8sClusterSourcesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListK8sClusterSourcesRequest) SetVpcId(v string) *ListK8sClusterSourcesRequest {
	s.VpcId = &v
	return s
}

func (s *ListK8sClusterSourcesRequest) Validate() error {
	return dara.Validate(s)
}
