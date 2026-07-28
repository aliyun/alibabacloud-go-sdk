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
	// You can query the associations between up to 20 SSL certificates and VPN gateway instances at a time.
	//
	// example:
	//
	// 6bfe4218-ea1d****
	CertificateId []*string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty" type:"Repeated"`
	// The certificate type. Valid values:
	//
	// - **Encryption**: encryption certificate.
	//
	// - **Signature**: signing certificate.
	//
	// example:
	//
	// Signature
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **20**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. Valid values:
	//
	// - If this is the first query or no next query exists, leave this parameter empty.
	//
	// - If a next query exists, set this parameter to the **NextToken*	- value returned by the previous API call.
	//
	// example:
	//
	// caeba0bbb2be0****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the VPN gateway.
	//
	// You can call [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of VPN gateway instance IDs.
	//
	// You can query the associations between up to 20 VPN gateway instances and SSL certificates at a time.
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
