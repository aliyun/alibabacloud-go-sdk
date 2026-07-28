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
	// The number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. Valid values:
	//
	// - If **NextToken*	- is empty, no next query exists.
	//
	// - If **NextToken*	- is returned, the value indicates the token for the next query.
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
	// The total number of associations.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of associations.
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
	if s.VpnCertificateRelations != nil {
		for _, item := range s.VpnCertificateRelations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVpnCertificateAssociationsResponseBodyVpnCertificateRelations struct {
	// The time when the association was created.
	//
	// The time is displayed in UTC in the YYYY-MM-DDThh:mm:ssZ format.
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
	// The certificate type.
	//
	// - **Encryption**: encryption certificate.
	//
	// - **Signature**: signing certificate.
	//
	// example:
	//
	// Signature
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The region ID of the VPN gateway.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The VPN gateway instance ID.
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
