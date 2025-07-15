// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpnCertificateAssociationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListVpnCertificateAssociationsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListVpnCertificateAssociationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListVpnCertificateAssociationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListVpnCertificateAssociationsResponseBody
	GetTotalCount() *int32
	SetVpnCertificateRelations(v []*ListVpnCertificateAssociationsResponseBodyVpnCertificateRelations) *ListVpnCertificateAssociationsResponseBody
	GetVpnCertificateRelations() []*ListVpnCertificateAssociationsResponseBodyVpnCertificateRelations
}

type ListVpnCertificateAssociationsResponseBody struct {
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If **NextToken*	- is not empty, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// caeba0bbb2be****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 197AF2BD-547F-470C-B29A-8400400233EB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The association information.
	VpnCertificateRelations []*ListVpnCertificateAssociationsResponseBodyVpnCertificateRelations `json:"VpnCertificateRelations,omitempty" xml:"VpnCertificateRelations,omitempty" type:"Repeated"`
}

func (s ListVpnCertificateAssociationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVpnCertificateAssociationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpnCertificateAssociationsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVpnCertificateAssociationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpnCertificateAssociationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVpnCertificateAssociationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVpnCertificateAssociationsResponseBody) GetVpnCertificateRelations() []*ListVpnCertificateAssociationsResponseBodyVpnCertificateRelations {
	return s.VpnCertificateRelations
}

func (s *ListVpnCertificateAssociationsResponseBody) SetMaxResults(v int32) *ListVpnCertificateAssociationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVpnCertificateAssociationsResponseBody) SetNextToken(v string) *ListVpnCertificateAssociationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVpnCertificateAssociationsResponseBody) SetRequestId(v string) *ListVpnCertificateAssociationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpnCertificateAssociationsResponseBody) SetTotalCount(v int32) *ListVpnCertificateAssociationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVpnCertificateAssociationsResponseBody) SetVpnCertificateRelations(v []*ListVpnCertificateAssociationsResponseBodyVpnCertificateRelations) *ListVpnCertificateAssociationsResponseBody {
	s.VpnCertificateRelations = v
	return s
}

func (s *ListVpnCertificateAssociationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListVpnCertificateAssociationsResponseBodyVpnCertificateRelations struct {
	// The time when the Anycast EIP was associated.
	//
	// The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-12-29T09:30:29Z
	AssociationTime *string `json:"AssociationTime,omitempty" xml:"AssociationTime,omitempty"`
	// The certificate ID.
	//
	// example:
	//
	// 6bfe4218-ea1d****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The type of the certificate.
	//
	// 	- **Encryption**
	//
	// 	- **Signature**
	//
	// example:
	//
	// Signature
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The ID of the region where the VPN gateway is created.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the VPN gateway.
	//
	// example:
	//
	// vpn-bp1usbiorilk51760****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s ListVpnCertificateAssociationsResponseBodyVpnCertificateRelations) String() string {
	return dara.Prettify(s)
}

func (s ListVpnCertificateAssociationsResponseBodyVpnCertificateRelations) GoString() string {
	return s.String()
}

func (s *ListVpnCertificateAssociationsResponseBodyVpnCertificateRelations) GetAssociationTime() *string {
	return s.AssociationTime
}

func (s *ListVpnCertificateAssociationsResponseBodyVpnCertificateRelations) GetCertificateId() *string {
	return s.CertificateId
}

func (s *ListVpnCertificateAssociationsResponseBodyVpnCertificateRelations) GetCertificateType() *string {
	return s.CertificateType
}

func (s *ListVpnCertificateAssociationsResponseBodyVpnCertificateRelations) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpnCertificateAssociationsResponseBodyVpnCertificateRelations) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *ListVpnCertificateAssociationsResponseBodyVpnCertificateRelations) SetAssociationTime(v string) *ListVpnCertificateAssociationsResponseBodyVpnCertificateRelations {
	s.AssociationTime = &v
	return s
}

func (s *ListVpnCertificateAssociationsResponseBodyVpnCertificateRelations) SetCertificateId(v string) *ListVpnCertificateAssociationsResponseBodyVpnCertificateRelations {
	s.CertificateId = &v
	return s
}

func (s *ListVpnCertificateAssociationsResponseBodyVpnCertificateRelations) SetCertificateType(v string) *ListVpnCertificateAssociationsResponseBodyVpnCertificateRelations {
	s.CertificateType = &v
	return s
}

func (s *ListVpnCertificateAssociationsResponseBodyVpnCertificateRelations) SetRegionId(v string) *ListVpnCertificateAssociationsResponseBodyVpnCertificateRelations {
	s.RegionId = &v
	return s
}

func (s *ListVpnCertificateAssociationsResponseBodyVpnCertificateRelations) SetVpnGatewayId(v string) *ListVpnCertificateAssociationsResponseBodyVpnCertificateRelations {
	s.VpnGatewayId = &v
	return s
}

func (s *ListVpnCertificateAssociationsResponseBodyVpnCertificateRelations) Validate() error {
	return dara.Validate(s)
}
