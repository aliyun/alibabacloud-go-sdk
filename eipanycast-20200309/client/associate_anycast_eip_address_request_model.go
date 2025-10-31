// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateAnycastEipAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnycastId(v string) *AssociateAnycastEipAddressRequest
	GetAnycastId() *string
	SetAssociationMode(v string) *AssociateAnycastEipAddressRequest
	GetAssociationMode() *string
	SetBindInstanceId(v string) *AssociateAnycastEipAddressRequest
	GetBindInstanceId() *string
	SetBindInstanceRegionId(v string) *AssociateAnycastEipAddressRequest
	GetBindInstanceRegionId() *string
	SetBindInstanceType(v string) *AssociateAnycastEipAddressRequest
	GetBindInstanceType() *string
	SetClientToken(v string) *AssociateAnycastEipAddressRequest
	GetClientToken() *string
	SetDryRun(v bool) *AssociateAnycastEipAddressRequest
	GetDryRun() *bool
	SetPopLocations(v []*AssociateAnycastEipAddressRequestPopLocations) *AssociateAnycastEipAddressRequest
	GetPopLocations() []*AssociateAnycastEipAddressRequestPopLocations
	SetPrivateIpAddress(v string) *AssociateAnycastEipAddressRequest
	GetPrivateIpAddress() *string
}

type AssociateAnycastEipAddressRequest struct {
	// The ID of the Anycast EIP.
	//
	// This parameter is required.
	//
	// example:
	//
	// aeip-bp1ix34fralt4ykf3****
	AnycastId *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	// The association mode. Valid values:
	//
	// 	- **Default**: the default mode. In this mode, the endpoint to be associated serves as the default origin server.
	//
	// 	- **Normal**: the standard mode. In this mode, the endpoint to be associated serves as a standard origin server.
	//
	// > You can associate endpoints in multiple regions with an Anycast EIP. However, only one endpoint can serve as the default origin server. Others serve as standard origin servers. If you do not specify or add an access point, requests are forwarded to the default origin server.\\
	//
	//
	// 	- If this is your first time to associate an Anycast EIP with an endpoint, set the value to **Default**.
	//
	// 	- If not, you can also set the value to **Default**, which specifies a new default origin server. In this case, the previous origin server functions as a standard origin server.
	//
	// example:
	//
	// Default
	AssociationMode *string `json:"AssociationMode,omitempty" xml:"AssociationMode,omitempty"`
	// The ID of the endpoint with which you want to associate the Anycast EIP.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-d7oxbixhxv1uupnon****
	BindInstanceId *string `json:"BindInstanceId,omitempty" xml:"BindInstanceId,omitempty"`
	// The ID of the region where the endpoint is deployed.
	//
	// You can associate Anycast EIPs only with endpoints in specific regions. You can call the [DescribeAnycastServerRegions](https://help.aliyun.com/document_detail/171939.html) operation to query the region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// us-west-1
	BindInstanceRegionId *string `json:"BindInstanceRegionId,omitempty" xml:"BindInstanceRegionId,omitempty"`
	// The type of endpoint with which you want to associate the Anycast EIP. Valid values:
	//
	// 	- **SlbInstance**: internal-facing Server Load Balancer (SLB) instance that is deployed in a virtual private cloud (VPC)
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
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false**(default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The information about the access points in associated access areas when you associate an Anycast EIP with an endpoint.
	//
	// If this is your first time to associate an Anycast EIP with an endpoint, ignore this parameter. The system automatically associates all access areas.
	//
	// You can call the [DescribeAnycastPopLocations](https://help.aliyun.com/document_detail/171938.html) operation to query information about access points in supported access areas.
	PopLocations []*AssociateAnycastEipAddressRequestPopLocations `json:"PopLocations,omitempty" xml:"PopLocations,omitempty" type:"Repeated"`
	// The secondary private IP address of the ENI with which you want to associate the Anycast EIP.
	//
	// This parameter is valid only when you set **BindInstanceType*	- to **NetworkInterface**. If you do not set this parameter, the primary private IP address of the ENI is used.
	//
	// example:
	//
	// 192.168.XX.XX
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
}

func (s AssociateAnycastEipAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateAnycastEipAddressRequest) GoString() string {
	return s.String()
}

func (s *AssociateAnycastEipAddressRequest) GetAnycastId() *string {
	return s.AnycastId
}

func (s *AssociateAnycastEipAddressRequest) GetAssociationMode() *string {
	return s.AssociationMode
}

func (s *AssociateAnycastEipAddressRequest) GetBindInstanceId() *string {
	return s.BindInstanceId
}

func (s *AssociateAnycastEipAddressRequest) GetBindInstanceRegionId() *string {
	return s.BindInstanceRegionId
}

func (s *AssociateAnycastEipAddressRequest) GetBindInstanceType() *string {
	return s.BindInstanceType
}

func (s *AssociateAnycastEipAddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AssociateAnycastEipAddressRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AssociateAnycastEipAddressRequest) GetPopLocations() []*AssociateAnycastEipAddressRequestPopLocations {
	return s.PopLocations
}

func (s *AssociateAnycastEipAddressRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *AssociateAnycastEipAddressRequest) SetAnycastId(v string) *AssociateAnycastEipAddressRequest {
	s.AnycastId = &v
	return s
}

func (s *AssociateAnycastEipAddressRequest) SetAssociationMode(v string) *AssociateAnycastEipAddressRequest {
	s.AssociationMode = &v
	return s
}

func (s *AssociateAnycastEipAddressRequest) SetBindInstanceId(v string) *AssociateAnycastEipAddressRequest {
	s.BindInstanceId = &v
	return s
}

func (s *AssociateAnycastEipAddressRequest) SetBindInstanceRegionId(v string) *AssociateAnycastEipAddressRequest {
	s.BindInstanceRegionId = &v
	return s
}

func (s *AssociateAnycastEipAddressRequest) SetBindInstanceType(v string) *AssociateAnycastEipAddressRequest {
	s.BindInstanceType = &v
	return s
}

func (s *AssociateAnycastEipAddressRequest) SetClientToken(v string) *AssociateAnycastEipAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateAnycastEipAddressRequest) SetDryRun(v bool) *AssociateAnycastEipAddressRequest {
	s.DryRun = &v
	return s
}

func (s *AssociateAnycastEipAddressRequest) SetPopLocations(v []*AssociateAnycastEipAddressRequestPopLocations) *AssociateAnycastEipAddressRequest {
	s.PopLocations = v
	return s
}

func (s *AssociateAnycastEipAddressRequest) SetPrivateIpAddress(v string) *AssociateAnycastEipAddressRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *AssociateAnycastEipAddressRequest) Validate() error {
	if s.PopLocations != nil {
		for _, item := range s.PopLocations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AssociateAnycastEipAddressRequestPopLocations struct {
	// The information about the access points in associated access areas when you associate an Anycast EIP with an endpoint.
	//
	// If this is your first time to associate an Anycast EIP with an endpoint, ignore this parameter. The system automatically associates all access areas.
	//
	// You can call the [DescribeAnycastPopLocations](https://help.aliyun.com/document_detail/171938.html) operation to query information about access points in supported access areas.
	//
	// example:
	//
	// us-west-1-pop
	PopLocation *string `json:"PopLocation,omitempty" xml:"PopLocation,omitempty"`
}

func (s AssociateAnycastEipAddressRequestPopLocations) String() string {
	return dara.Prettify(s)
}

func (s AssociateAnycastEipAddressRequestPopLocations) GoString() string {
	return s.String()
}

func (s *AssociateAnycastEipAddressRequestPopLocations) GetPopLocation() *string {
	return s.PopLocation
}

func (s *AssociateAnycastEipAddressRequestPopLocations) SetPopLocation(v string) *AssociateAnycastEipAddressRequestPopLocations {
	s.PopLocation = &v
	return s
}

func (s *AssociateAnycastEipAddressRequestPopLocations) Validate() error {
	return dara.Validate(s)
}
