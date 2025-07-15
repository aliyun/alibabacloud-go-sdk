// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpnCertificateAssociationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateId(v []*string) *ListVpnCertificateAssociationsRequest
	GetCertificateId() []*string
	SetCertificateType(v string) *ListVpnCertificateAssociationsRequest
	GetCertificateType() *string
	SetMaxResults(v int32) *ListVpnCertificateAssociationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListVpnCertificateAssociationsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListVpnCertificateAssociationsRequest
	GetRegionId() *string
	SetVpnGatewayId(v []*string) *ListVpnCertificateAssociationsRequest
	GetVpnGatewayId() []*string
}

type ListVpnCertificateAssociationsRequest struct {
	// The list of certificate IDs.
	//
	// You can query the association between at most 20 SSL certificates and VPN gateways.
	//
	// example:
	//
	// 6bfe4218-ea1d****
	CertificateId []*string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty" type:"Repeated"`
	// The certificate type. Valid values:
	//
	// 	- **Encryption**
	//
	// 	- **Signature**
	//
	// example:
	//
	// Signature
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The number of entries to return on each page. Valid values: **1*	- to **20**. Default value: **1**.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of **NextToken**.
	//
	// example:
	//
	// caeba0bbb2be0****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	// The list of VPN gateway IDs.
	//
	// You can query the association between at most 20 VPN gateways and SSL certificates.
	//
	// example:
	//
	// vpn-bp1q8bgx4xnkm****
	VpnGatewayId []*string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty" type:"Repeated"`
}

func (s ListVpnCertificateAssociationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVpnCertificateAssociationsRequest) GoString() string {
	return s.String()
}

func (s *ListVpnCertificateAssociationsRequest) GetCertificateId() []*string {
	return s.CertificateId
}

func (s *ListVpnCertificateAssociationsRequest) GetCertificateType() *string {
	return s.CertificateType
}

func (s *ListVpnCertificateAssociationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVpnCertificateAssociationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpnCertificateAssociationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpnCertificateAssociationsRequest) GetVpnGatewayId() []*string {
	return s.VpnGatewayId
}

func (s *ListVpnCertificateAssociationsRequest) SetCertificateId(v []*string) *ListVpnCertificateAssociationsRequest {
	s.CertificateId = v
	return s
}

func (s *ListVpnCertificateAssociationsRequest) SetCertificateType(v string) *ListVpnCertificateAssociationsRequest {
	s.CertificateType = &v
	return s
}

func (s *ListVpnCertificateAssociationsRequest) SetMaxResults(v int32) *ListVpnCertificateAssociationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVpnCertificateAssociationsRequest) SetNextToken(v string) *ListVpnCertificateAssociationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListVpnCertificateAssociationsRequest) SetRegionId(v string) *ListVpnCertificateAssociationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpnCertificateAssociationsRequest) SetVpnGatewayId(v []*string) *ListVpnCertificateAssociationsRequest {
	s.VpnGatewayId = v
	return s
}

func (s *ListVpnCertificateAssociationsRequest) Validate() error {
	return dara.Validate(s)
}
