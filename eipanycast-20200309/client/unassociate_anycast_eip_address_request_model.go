// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateAnycastEipAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnycastId(v string) *UnassociateAnycastEipAddressRequest
	GetAnycastId() *string
	SetBindInstanceId(v string) *UnassociateAnycastEipAddressRequest
	GetBindInstanceId() *string
	SetBindInstanceRegionId(v string) *UnassociateAnycastEipAddressRequest
	GetBindInstanceRegionId() *string
	SetBindInstanceType(v string) *UnassociateAnycastEipAddressRequest
	GetBindInstanceType() *string
	SetClientToken(v string) *UnassociateAnycastEipAddressRequest
	GetClientToken() *string
	SetDryRun(v string) *UnassociateAnycastEipAddressRequest
	GetDryRun() *string
	SetPrivateIpAddress(v string) *UnassociateAnycastEipAddressRequest
	GetPrivateIpAddress() *string
}

type UnassociateAnycastEipAddressRequest struct {
	// The ID of the Anycast EIP.
	//
	// This parameter is required.
	//
	// example:
	//
	// aeip-2zeerraiwb7ujsxdc****
	AnycastId *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	// The ID of the endpoint from which you want to disassociate the Anycast EIP.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-2zebb08phyczzawe****
	BindInstanceId *string `json:"BindInstanceId,omitempty" xml:"BindInstanceId,omitempty"`
	// The region where the endpoint is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// us-west-1
	BindInstanceRegionId *string `json:"BindInstanceRegionId,omitempty" xml:"BindInstanceRegionId,omitempty"`
	// The type of endpoint from which you want to disassociate the Anycast EIP. Valid values:
	//
	// 	- **SlbInstance**: an internal-facing Server Load Balancer (SLB) instance that is deployed in a virtual private cloud (VPC)
	//
	// 	- **NetworkInterface**: elastic network interface (ENI)
	//
	// This parameter is required.
	//
	// example:
	//
	// SlbInstance
	BindInstanceType *string `json:"BindInstanceType,omitempty" xml:"BindInstanceType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The secondary private IP address of the ENI from which you want to disassociate the Anycast EIP.
	//
	// This parameter is valid only when you set **BindInstanceType*	- to **NetworkInterface**. If you do not specify this parameter, the primary private IP address of the ENI is used.
	//
	// example:
	//
	// 192.168.XX.XX
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
}

func (s UnassociateAnycastEipAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s UnassociateAnycastEipAddressRequest) GoString() string {
	return s.String()
}

func (s *UnassociateAnycastEipAddressRequest) GetAnycastId() *string {
	return s.AnycastId
}

func (s *UnassociateAnycastEipAddressRequest) GetBindInstanceId() *string {
	return s.BindInstanceId
}

func (s *UnassociateAnycastEipAddressRequest) GetBindInstanceRegionId() *string {
	return s.BindInstanceRegionId
}

func (s *UnassociateAnycastEipAddressRequest) GetBindInstanceType() *string {
	return s.BindInstanceType
}

func (s *UnassociateAnycastEipAddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UnassociateAnycastEipAddressRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *UnassociateAnycastEipAddressRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *UnassociateAnycastEipAddressRequest) SetAnycastId(v string) *UnassociateAnycastEipAddressRequest {
	s.AnycastId = &v
	return s
}

func (s *UnassociateAnycastEipAddressRequest) SetBindInstanceId(v string) *UnassociateAnycastEipAddressRequest {
	s.BindInstanceId = &v
	return s
}

func (s *UnassociateAnycastEipAddressRequest) SetBindInstanceRegionId(v string) *UnassociateAnycastEipAddressRequest {
	s.BindInstanceRegionId = &v
	return s
}

func (s *UnassociateAnycastEipAddressRequest) SetBindInstanceType(v string) *UnassociateAnycastEipAddressRequest {
	s.BindInstanceType = &v
	return s
}

func (s *UnassociateAnycastEipAddressRequest) SetClientToken(v string) *UnassociateAnycastEipAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *UnassociateAnycastEipAddressRequest) SetDryRun(v string) *UnassociateAnycastEipAddressRequest {
	s.DryRun = &v
	return s
}

func (s *UnassociateAnycastEipAddressRequest) SetPrivateIpAddress(v string) *UnassociateAnycastEipAddressRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *UnassociateAnycastEipAddressRequest) Validate() error {
	return dara.Validate(s)
}
