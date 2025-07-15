// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateVpnGatewayWithCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateId(v string) *AssociateVpnGatewayWithCertificateRequest
	GetCertificateId() *string
	SetCertificateType(v string) *AssociateVpnGatewayWithCertificateRequest
	GetCertificateType() *string
	SetClientToken(v string) *AssociateVpnGatewayWithCertificateRequest
	GetClientToken() *string
	SetDryRun(v bool) *AssociateVpnGatewayWithCertificateRequest
	GetDryRun() *bool
	SetRegionId(v string) *AssociateVpnGatewayWithCertificateRequest
	GetRegionId() *string
	SetVpnGatewayId(v string) *AssociateVpnGatewayWithCertificateRequest
	GetVpnGatewayId() *string
}

type AssociateVpnGatewayWithCertificateRequest struct {
	// The ID of the certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6bfe4218-ea1d****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The type of the certificate. Valid values:
	//
	// 	- **Encryption**
	//
	// 	- **Signature**
	//
	// This parameter is required.
	//
	// example:
	//
	// Signature
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 0c593ea1-3bea****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request passes the dry run, a request ID is returned. Otherwise, an error message is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the VPN gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the VPN gateway.
	//
	// > You can associate only VPN gateways of the SM type with certificates.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpn-bp1q8bgx4xnkm2ogj****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s AssociateVpnGatewayWithCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateVpnGatewayWithCertificateRequest) GoString() string {
	return s.String()
}

func (s *AssociateVpnGatewayWithCertificateRequest) GetCertificateId() *string {
	return s.CertificateId
}

func (s *AssociateVpnGatewayWithCertificateRequest) GetCertificateType() *string {
	return s.CertificateType
}

func (s *AssociateVpnGatewayWithCertificateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AssociateVpnGatewayWithCertificateRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AssociateVpnGatewayWithCertificateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociateVpnGatewayWithCertificateRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *AssociateVpnGatewayWithCertificateRequest) SetCertificateId(v string) *AssociateVpnGatewayWithCertificateRequest {
	s.CertificateId = &v
	return s
}

func (s *AssociateVpnGatewayWithCertificateRequest) SetCertificateType(v string) *AssociateVpnGatewayWithCertificateRequest {
	s.CertificateType = &v
	return s
}

func (s *AssociateVpnGatewayWithCertificateRequest) SetClientToken(v string) *AssociateVpnGatewayWithCertificateRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateVpnGatewayWithCertificateRequest) SetDryRun(v bool) *AssociateVpnGatewayWithCertificateRequest {
	s.DryRun = &v
	return s
}

func (s *AssociateVpnGatewayWithCertificateRequest) SetRegionId(v string) *AssociateVpnGatewayWithCertificateRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateVpnGatewayWithCertificateRequest) SetVpnGatewayId(v string) *AssociateVpnGatewayWithCertificateRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *AssociateVpnGatewayWithCertificateRequest) Validate() error {
	return dara.Validate(s)
}
