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
	// The certificate ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6bfe4218-ea1d****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The certificate type. Valid values:
	//
	// - **Encryption**: specifies the SSL certificate as the encryption certificate.
	//
	// - **Signature**: specifies the SSL certificate as the signing certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// Signature
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 0c593ea1-3bea****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without actually associating the certificate. The system checks the required parameters, request syntax, and instance status. If the check fails, the corresponding error is returned. If the check succeeds, the corresponding request ID is returned.
	//
	// - **false*	- (default): sends the request. After the request passes the check, the VPN gateway is associated with the certificate.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the VPN gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance ID of the VPN gateway.
	//
	// > Only ShangMi (SM) VPN gateways support attaching certificates.
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
