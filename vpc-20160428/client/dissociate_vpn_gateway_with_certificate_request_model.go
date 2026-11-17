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
	// > The certificate ID refers to the ID generated after the SSL certificate is associated with the VPN gateway, not the ID of the SSL certificate itself.
	//
	// <props="china">You can call the [ListVpnCertificateAssociations](https://help.aliyun.com/document_detail/2521961.html) operation to query the certificate ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6bfe4218-ea1d****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The type of the certificate. Valid values:
	//
	// - **Encryption**: encryption certificate.
	//
	// - **Signature**: signing certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// Encryption
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, and you must make sure that the value is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e*******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run on the request. Valid values:
	//
	// - **true**: Performs only a dry run and does not perform the dissociation operation. The system checks whether required parameters are specified, the request format is valid, and the instance status is valid. If the request fails the dry run, an error is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): Sends a normal request. If the request passes the dry run, the VPN gateway is dissociated from the SSL certificate.
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
