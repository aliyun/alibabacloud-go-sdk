// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateVpnGatewayWithCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateId(v string) *DissociateVpnGatewayWithCertificateRequest
	GetCertificateId() *string
	SetCertificateType(v string) *DissociateVpnGatewayWithCertificateRequest
	GetCertificateType() *string
	SetClientToken(v string) *DissociateVpnGatewayWithCertificateRequest
	GetClientToken() *string
	SetDryRun(v bool) *DissociateVpnGatewayWithCertificateRequest
	GetDryRun() *bool
	SetRegionId(v string) *DissociateVpnGatewayWithCertificateRequest
	GetRegionId() *string
	SetVpnGatewayId(v string) *DissociateVpnGatewayWithCertificateRequest
	GetVpnGatewayId() *string
}

type DissociateVpnGatewayWithCertificateRequest struct {
	// The ID of the certificate.
	//
	// >  The certificate ID refers to the ID generated after the SSL certificate is associated with the VPN gateway. It is not the ID of the SSL certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6bfe4218-ea1d****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The certificate type. Valid values:
	//
	// 	- **Encryption**
	//
	// 	- **Signature**
	//
	// This parameter is required.
	//
	// example:
	//
	// Encryption
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e*******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request passes the dry run, the `DryRunOperation` error code is returned. Otherwise, an error message is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
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
	// This parameter is required.
	//
	// example:
	//
	// vpn-bp1q8bgx4xnkm2ogj****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s DissociateVpnGatewayWithCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s DissociateVpnGatewayWithCertificateRequest) GoString() string {
	return s.String()
}

func (s *DissociateVpnGatewayWithCertificateRequest) GetCertificateId() *string {
	return s.CertificateId
}

func (s *DissociateVpnGatewayWithCertificateRequest) GetCertificateType() *string {
	return s.CertificateType
}

func (s *DissociateVpnGatewayWithCertificateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DissociateVpnGatewayWithCertificateRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DissociateVpnGatewayWithCertificateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DissociateVpnGatewayWithCertificateRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *DissociateVpnGatewayWithCertificateRequest) SetCertificateId(v string) *DissociateVpnGatewayWithCertificateRequest {
	s.CertificateId = &v
	return s
}

func (s *DissociateVpnGatewayWithCertificateRequest) SetCertificateType(v string) *DissociateVpnGatewayWithCertificateRequest {
	s.CertificateType = &v
	return s
}

func (s *DissociateVpnGatewayWithCertificateRequest) SetClientToken(v string) *DissociateVpnGatewayWithCertificateRequest {
	s.ClientToken = &v
	return s
}

func (s *DissociateVpnGatewayWithCertificateRequest) SetDryRun(v bool) *DissociateVpnGatewayWithCertificateRequest {
	s.DryRun = &v
	return s
}

func (s *DissociateVpnGatewayWithCertificateRequest) SetRegionId(v string) *DissociateVpnGatewayWithCertificateRequest {
	s.RegionId = &v
	return s
}

func (s *DissociateVpnGatewayWithCertificateRequest) SetVpnGatewayId(v string) *DissociateVpnGatewayWithCertificateRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *DissociateVpnGatewayWithCertificateRequest) Validate() error {
	return dara.Validate(s)
}
