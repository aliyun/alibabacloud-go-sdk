// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribleCertListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertList(v []*DescribleCertListResponseBodyCertList) *DescribleCertListResponseBody
	GetCertList() []*DescribleCertListResponseBodyCertList
	SetRequestId(v string) *DescribleCertListResponseBody
	GetRequestId() *string
}

type DescribleCertListResponseBody struct {
	CertList []*DescribleCertListResponseBodyCertList `json:"CertList,omitempty" xml:"CertList,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribleCertListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribleCertListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribleCertListResponseBody) GetCertList() []*DescribleCertListResponseBodyCertList {
	return s.CertList
}

func (s *DescribleCertListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribleCertListResponseBody) SetCertList(v []*DescribleCertListResponseBodyCertList) *DescribleCertListResponseBody {
	s.CertList = v
	return s
}

func (s *DescribleCertListResponseBody) SetRequestId(v string) *DescribleCertListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribleCertListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribleCertListResponseBodyCertList struct {
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// example:
	//
	// www.aliyun.com
	Common *string `json:"Common,omitempty" xml:"Common,omitempty"`
	// example:
	//
	// false
	DomainRelated *bool `json:"DomainRelated,omitempty" xml:"DomainRelated,omitempty"`
	// example:
	//
	// 2020-09-23
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 123
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// DigiCert Inc
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// example:
	//
	// testCertName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2019-09-24
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribleCertListResponseBodyCertList) String() string {
	return dara.Prettify(s)
}

func (s DescribleCertListResponseBodyCertList) GoString() string {
	return s.String()
}

func (s *DescribleCertListResponseBodyCertList) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *DescribleCertListResponseBodyCertList) GetCommon() *string {
	return s.Common
}

func (s *DescribleCertListResponseBodyCertList) GetDomainRelated() *bool {
	return s.DomainRelated
}

func (s *DescribleCertListResponseBodyCertList) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribleCertListResponseBodyCertList) GetId() *int32 {
	return s.Id
}

func (s *DescribleCertListResponseBodyCertList) GetIssuer() *string {
	return s.Issuer
}

func (s *DescribleCertListResponseBodyCertList) GetName() *string {
	return s.Name
}

func (s *DescribleCertListResponseBodyCertList) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribleCertListResponseBodyCertList) SetCertIdentifier(v string) *DescribleCertListResponseBodyCertList {
	s.CertIdentifier = &v
	return s
}

func (s *DescribleCertListResponseBodyCertList) SetCommon(v string) *DescribleCertListResponseBodyCertList {
	s.Common = &v
	return s
}

func (s *DescribleCertListResponseBodyCertList) SetDomainRelated(v bool) *DescribleCertListResponseBodyCertList {
	s.DomainRelated = &v
	return s
}

func (s *DescribleCertListResponseBodyCertList) SetEndDate(v string) *DescribleCertListResponseBodyCertList {
	s.EndDate = &v
	return s
}

func (s *DescribleCertListResponseBodyCertList) SetId(v int32) *DescribleCertListResponseBodyCertList {
	s.Id = &v
	return s
}

func (s *DescribleCertListResponseBodyCertList) SetIssuer(v string) *DescribleCertListResponseBodyCertList {
	s.Issuer = &v
	return s
}

func (s *DescribleCertListResponseBodyCertList) SetName(v string) *DescribleCertListResponseBodyCertList {
	s.Name = &v
	return s
}

func (s *DescribleCertListResponseBodyCertList) SetStartDate(v string) *DescribleCertListResponseBodyCertList {
	s.StartDate = &v
	return s
}

func (s *DescribleCertListResponseBodyCertList) Validate() error {
	return dara.Validate(s)
}
