// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnSMCertificateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateListModel(v *DescribeCdnSMCertificateListResponseBodyCertificateListModel) *DescribeCdnSMCertificateListResponseBody
	GetCertificateListModel() *DescribeCdnSMCertificateListResponseBodyCertificateListModel
	SetRequestId(v string) *DescribeCdnSMCertificateListResponseBody
	GetRequestId() *string
}

type DescribeCdnSMCertificateListResponseBody struct {
	// The type of the certificate information.
	CertificateListModel *DescribeCdnSMCertificateListResponseBodyCertificateListModel `json:"CertificateListModel,omitempty" xml:"CertificateListModel,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// DC0E34AC-0239-44A7-AB0E-800DE522C8DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnSMCertificateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnSMCertificateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnSMCertificateListResponseBody) GetCertificateListModel() *DescribeCdnSMCertificateListResponseBodyCertificateListModel {
	return s.CertificateListModel
}

func (s *DescribeCdnSMCertificateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnSMCertificateListResponseBody) SetCertificateListModel(v *DescribeCdnSMCertificateListResponseBodyCertificateListModel) *DescribeCdnSMCertificateListResponseBody {
	s.CertificateListModel = v
	return s
}

func (s *DescribeCdnSMCertificateListResponseBody) SetRequestId(v string) *DescribeCdnSMCertificateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnSMCertificateListResponseBody) Validate() error {
	if s.CertificateListModel != nil {
		if err := s.CertificateListModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnSMCertificateListResponseBodyCertificateListModel struct {
	// The list of certificates.
	CertList *DescribeCdnSMCertificateListResponseBodyCertificateListModelCertList `json:"CertList,omitempty" xml:"CertList,omitempty" type:"Struct"`
	// The number of certificates that are returned.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeCdnSMCertificateListResponseBodyCertificateListModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnSMCertificateListResponseBodyCertificateListModel) GoString() string {
	return s.String()
}

func (s *DescribeCdnSMCertificateListResponseBodyCertificateListModel) GetCertList() *DescribeCdnSMCertificateListResponseBodyCertificateListModelCertList {
	return s.CertList
}

func (s *DescribeCdnSMCertificateListResponseBodyCertificateListModel) GetCount() *int32 {
	return s.Count
}

func (s *DescribeCdnSMCertificateListResponseBodyCertificateListModel) SetCertList(v *DescribeCdnSMCertificateListResponseBodyCertificateListModelCertList) *DescribeCdnSMCertificateListResponseBodyCertificateListModel {
	s.CertList = v
	return s
}

func (s *DescribeCdnSMCertificateListResponseBodyCertificateListModel) SetCount(v int32) *DescribeCdnSMCertificateListResponseBodyCertificateListModel {
	s.Count = &v
	return s
}

func (s *DescribeCdnSMCertificateListResponseBodyCertificateListModel) Validate() error {
	if s.CertList != nil {
		if err := s.CertList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnSMCertificateListResponseBodyCertificateListModelCertList struct {
	Cert []*DescribeCdnSMCertificateListResponseBodyCertificateListModelCertListCert `json:"Cert,omitempty" xml:"Cert,omitempty" type:"Repeated"`
}

func (s DescribeCdnSMCertificateListResponseBodyCertificateListModelCertList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnSMCertificateListResponseBodyCertificateListModelCertList) GoString() string {
	return s.String()
}

func (s *DescribeCdnSMCertificateListResponseBodyCertificateListModelCertList) GetCert() []*DescribeCdnSMCertificateListResponseBodyCertificateListModelCertListCert {
	return s.Cert
}

func (s *DescribeCdnSMCertificateListResponseBodyCertificateListModelCertList) SetCert(v []*DescribeCdnSMCertificateListResponseBodyCertificateListModelCertListCert) *DescribeCdnSMCertificateListResponseBodyCertificateListModelCertList {
	s.Cert = v
	return s
}

func (s *DescribeCdnSMCertificateListResponseBodyCertificateListModelCertList) Validate() error {
	if s.Cert != nil {
		for _, item := range s.Cert {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCdnSMCertificateListResponseBodyCertificateListModelCertListCert struct {
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

func (s DescribeCdnSMCertificateListResponseBodyCertificateListModelCertListCert) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnSMCertificateListResponseBodyCertificateListModelCertListCert) GoString() string {
	return s.String()
}

func (s *DescribeCdnSMCertificateListResponseBodyCertificateListModelCertListCert) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *DescribeCdnSMCertificateListResponseBodyCertificateListModelCertListCert) GetCertName() *string {
	return s.CertName
}

func (s *DescribeCdnSMCertificateListResponseBodyCertificateListModelCertListCert) GetCommon() *string {
	return s.Common
}

func (s *DescribeCdnSMCertificateListResponseBodyCertificateListModelCertListCert) GetIssuer() *string {
	return s.Issuer
}

func (s *DescribeCdnSMCertificateListResponseBodyCertificateListModelCertListCert) SetCertIdentifier(v string) *DescribeCdnSMCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertIdentifier = &v
	return s
}

func (s *DescribeCdnSMCertificateListResponseBodyCertificateListModelCertListCert) SetCertName(v string) *DescribeCdnSMCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertName = &v
	return s
}

func (s *DescribeCdnSMCertificateListResponseBodyCertificateListModelCertListCert) SetCommon(v string) *DescribeCdnSMCertificateListResponseBodyCertificateListModelCertListCert {
	s.Common = &v
	return s
}

func (s *DescribeCdnSMCertificateListResponseBodyCertificateListModelCertListCert) SetIssuer(v string) *DescribeCdnSMCertificateListResponseBodyCertificateListModelCertListCert {
	s.Issuer = &v
	return s
}

func (s *DescribeCdnSMCertificateListResponseBodyCertificateListModelCertListCert) Validate() error {
	return dara.Validate(s)
}
