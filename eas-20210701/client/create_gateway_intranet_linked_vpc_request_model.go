// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayIntranetLinkedVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *CreateGatewayIntranetLinkedVpcRequest
	GetAccountId() *string
	SetEnableAuthoritativeDns(v bool) *CreateGatewayIntranetLinkedVpcRequest
	GetEnableAuthoritativeDns() *bool
	SetVSwitchId(v string) *CreateGatewayIntranetLinkedVpcRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateGatewayIntranetLinkedVpcRequest
	GetVpcId() *string
}

type CreateGatewayIntranetLinkedVpcRequest struct {
	// The ID of the Alibaba Cloud account to which the VPC belongs.
	//
	// example:
	//
	// 19*****10
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Specifies whether to enable authoritative DNS parsing. The default value is false.
	//
	// example:
	//
	// true
	EnableAuthoritativeDns *bool `json:"EnableAuthoritativeDns,omitempty" xml:"EnableAuthoritativeDns,omitempty"`
	// The ID of the virtual switch. For more information, see [DescribeVpcs](https://help.aliyun.com/document_detail/448581.html).
	//
	// example:
	//
	// vsw-8vbqn2at0kljjxxxx****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC). For more information, see [DescribeVpcs](https://help.aliyun.com/document_detail/448581.html).
	//
	// example:
	//
	// vpc-uf66uio7md****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateGatewayIntranetLinkedVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayIntranetLinkedVpcRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayIntranetLinkedVpcRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateGatewayIntranetLinkedVpcRequest) GetEnableAuthoritativeDns() *bool {
	return s.EnableAuthoritativeDns
}

func (s *CreateGatewayIntranetLinkedVpcRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateGatewayIntranetLinkedVpcRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateGatewayIntranetLinkedVpcRequest) SetAccountId(v string) *CreateGatewayIntranetLinkedVpcRequest {
	s.AccountId = &v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcRequest) SetEnableAuthoritativeDns(v bool) *CreateGatewayIntranetLinkedVpcRequest {
	s.EnableAuthoritativeDns = &v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcRequest) SetVSwitchId(v string) *CreateGatewayIntranetLinkedVpcRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcRequest) SetVpcId(v string) *CreateGatewayIntranetLinkedVpcRequest {
	s.VpcId = &v
	return s
}

func (s *CreateGatewayIntranetLinkedVpcRequest) Validate() error {
	return dara.Validate(s)
}
