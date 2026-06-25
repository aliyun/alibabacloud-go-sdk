// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayIntranetLinkedVpcPeerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVpcId(v string) *ListGatewayIntranetLinkedVpcPeerRequest
	GetVpcId() *string
}

type ListGatewayIntranetLinkedVpcPeerRequest struct {
	// The ID of the attached virtual private cloud (VPC). For more information, see [ListGatewayIntranetLinkedVpc](https://help.aliyun.com/document_detail/2621223.html).
	//
	// - Specify a VPC ID to query only the VPC peers for that VPC.
	//
	// - If you do not specify a VPC ID, all VPC peers are returned.
	//
	// example:
	//
	// vpc-2zetuli9ws0qgjd******
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListGatewayIntranetLinkedVpcPeerRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayIntranetLinkedVpcPeerRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayIntranetLinkedVpcPeerRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListGatewayIntranetLinkedVpcPeerRequest) SetVpcId(v string) *ListGatewayIntranetLinkedVpcPeerRequest {
	s.VpcId = &v
	return s
}

func (s *ListGatewayIntranetLinkedVpcPeerRequest) Validate() error {
	return dara.Validate(s)
}
