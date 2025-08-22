// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnSMCertificateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateListModel(v *DescribeDcdnSMCertificateListResponseBodyCertificateListModel) *DescribeDcdnSMCertificateListResponseBody
	GetCertificateListModel() *DescribeDcdnSMCertificateListResponseBodyCertificateListModel
	SetRequestId(v string) *DescribeDcdnSMCertificateListResponseBody
	GetRequestId() *string
}

type DescribeDcdnSMCertificateListResponseBody struct {
	// The type of the certificate information.
	CertificateListModel *DescribeDcdnSMCertificateListResponseBodyCertificateListModel `json:"CertificateListModel,omitempty" xml:"CertificateListModel,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// DC0E34AC-0239-44A7-AB0E-800DE522C8DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnSMCertificateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSMCertificateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSMCertificateListResponseBody) GetCertificateListModel() *DescribeDcdnSMCertificateListResponseBodyCertificateListModel {
	return s.CertificateListModel
}

func (s *DescribeDcdnSMCertificateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnSMCertificateListResponseBody) SetCertificateListModel(v *DescribeDcdnSMCertificateListResponseBodyCertificateListModel) *DescribeDcdnSMCertificateListResponseBody {
	s.CertificateListModel = v
	return s
}

func (s *DescribeDcdnSMCertificateListResponseBody) SetRequestId(v string) *DescribeDcdnSMCertificateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnSMCertificateListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnSMCertificateListResponseBodyCertificateListModel struct {
	// A list of certificates.
	CertList []*DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList `json:"CertList,omitempty" xml:"CertList,omitempty" type:"Repeated"`
	// The number of certificates that are returned.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeDcdnSMCertificateListResponseBodyCertificateListModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSMCertificateListResponseBodyCertificateListModel) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSMCertificateListResponseBodyCertificateListModel) GetCertList() []*DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList {
	return s.CertList
}

func (s *DescribeDcdnSMCertificateListResponseBodyCertificateListModel) GetCount() *int32 {
	return s.Count
}

func (s *DescribeDcdnSMCertificateListResponseBodyCertificateListModel) SetCertList(v []*DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList) *DescribeDcdnSMCertificateListResponseBodyCertificateListModel {
	s.CertList = v
	return s
}

func (s *DescribeDcdnSMCertificateListResponseBodyCertificateListModel) SetCount(v int32) *DescribeDcdnSMCertificateListResponseBodyCertificateListModel {
	s.Count = &v
	return s
}

func (s *DescribeDcdnSMCertificateListResponseBodyCertificateListModel) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList struct {
	// The ID of the certificate.
	//
	// example:
	//
	// yourCertldentifier
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// yourCertName
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The common name of the certificate.
	//
	// example:
	//
	// yourCertCommon
	Common *string `json:"Common,omitempty" xml:"Common,omitempty"`
	// The certificate authority (CA) that issued the certificate.
	//
	// example:
	//
	// yourCertIssuer
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
}

func (s DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList) GetCertName() *string {
	return s.CertName
}

func (s *DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList) GetCommon() *string {
	return s.Common
}

func (s *DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList) GetIssuer() *string {
	return s.Issuer
}

func (s *DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList) SetCertIdentifier(v string) *DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList {
	s.CertIdentifier = &v
	return s
}

func (s *DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList) SetCertName(v string) *DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList {
	s.CertName = &v
	return s
}

func (s *DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList) SetCommon(v string) *DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList {
	s.Common = &v
	return s
}

func (s *DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList) SetIssuer(v string) *DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList {
	s.Issuer = &v
	return s
}

func (s *DescribeDcdnSMCertificateListResponseBodyCertificateListModelCertList) Validate() error {
	return dara.Validate(s)
}
