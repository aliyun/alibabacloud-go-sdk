// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWarehouseCertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertIdentifier(v string) *DescribeWarehouseCertResponseBody
	GetCertIdentifier() *string
	SetCertStatus(v string) *DescribeWarehouseCertResponseBody
	GetCertStatus() *string
	SetCertType(v string) *DescribeWarehouseCertResponseBody
	GetCertType() *string
	SetCommonName(v string) *DescribeWarehouseCertResponseBody
	GetCommonName() *string
	SetContent(v string) *DescribeWarehouseCertResponseBody
	GetContent() *string
	SetFingerprint(v string) *DescribeWarehouseCertResponseBody
	GetFingerprint() *string
	SetIssuer(v string) *DescribeWarehouseCertResponseBody
	GetIssuer() *string
	SetIssuerIdentifier(v string) *DescribeWarehouseCertResponseBody
	GetIssuerIdentifier() *string
	SetPrivateCaInstanceId(v string) *DescribeWarehouseCertResponseBody
	GetPrivateCaInstanceId() *string
	SetPrivateCaRegionId(v string) *DescribeWarehouseCertResponseBody
	GetPrivateCaRegionId() *string
	SetRequestId(v string) *DescribeWarehouseCertResponseBody
	GetRequestId() *string
	SetWarehouseInstanceId(v string) *DescribeWarehouseCertResponseBody
	GetWarehouseInstanceId() *string
}

type DescribeWarehouseCertResponseBody struct {
	// The unique identifier of the certificate.
	//
	// example:
	//
	// 1ef1da5f-38ed-69b3-****-037781890265
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The status of the certificate.
	//
	// example:
	//
	// issued
	CertStatus *string `json:"CertStatus,omitempty" xml:"CertStatus,omitempty"`
	// The type of the certificate.
	//
	// example:
	//
	// OV
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The common name of the certificate subject. This field is empty if the certificate subject does not include a common name (CN).
	//
	// example:
	//
	// aliyundoc.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The content of the certificate, including the certificate chain.
	//
	// example:
	//
	// ---BEGIN CERTIFICATE----- MIIF...... -----END CERTIFICATE-----
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The fingerprint of the certificate content.
	//
	// example:
	//
	// C1291AF83F48170E48140FDFE5DADC19FE51F261
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The common name (or organization name) of the issuer.
	//
	// example:
	//
	// Digicert
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The unique identifier of the issuer certificate.
	//
	// example:
	//
	// 1ef1da5f-38ed-69b3-****-037781890265
	IssuerIdentifier *string `json:"IssuerIdentifier,omitempty" xml:"IssuerIdentifier,omitempty"`
	// The instance ID of the private CA instance associated with the certificate.
	//
	// example:
	//
	// 1ef1da5f-38ed-69b3-****-037781890265
	PrivateCaInstanceId *string `json:"PrivateCaInstanceId,omitempty" xml:"PrivateCaInstanceId,omitempty"`
	// The region ID of the private CA instance associated with the certificate.
	//
	// example:
	//
	// cn-hangzhou
	PrivateCaRegionId *string `json:"PrivateCaRegionId,omitempty" xml:"PrivateCaRegionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The instance ID of the certificate warehouse.
	//
	// example:
	//
	// 66
	WarehouseInstanceId *string `json:"WarehouseInstanceId,omitempty" xml:"WarehouseInstanceId,omitempty"`
}

func (s DescribeWarehouseCertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWarehouseCertResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWarehouseCertResponseBody) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *DescribeWarehouseCertResponseBody) GetCertStatus() *string {
	return s.CertStatus
}

func (s *DescribeWarehouseCertResponseBody) GetCertType() *string {
	return s.CertType
}

func (s *DescribeWarehouseCertResponseBody) GetCommonName() *string {
	return s.CommonName
}

func (s *DescribeWarehouseCertResponseBody) GetContent() *string {
	return s.Content
}

func (s *DescribeWarehouseCertResponseBody) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *DescribeWarehouseCertResponseBody) GetIssuer() *string {
	return s.Issuer
}

func (s *DescribeWarehouseCertResponseBody) GetIssuerIdentifier() *string {
	return s.IssuerIdentifier
}

func (s *DescribeWarehouseCertResponseBody) GetPrivateCaInstanceId() *string {
	return s.PrivateCaInstanceId
}

func (s *DescribeWarehouseCertResponseBody) GetPrivateCaRegionId() *string {
	return s.PrivateCaRegionId
}

func (s *DescribeWarehouseCertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWarehouseCertResponseBody) GetWarehouseInstanceId() *string {
	return s.WarehouseInstanceId
}

func (s *DescribeWarehouseCertResponseBody) SetCertIdentifier(v string) *DescribeWarehouseCertResponseBody {
	s.CertIdentifier = &v
	return s
}

func (s *DescribeWarehouseCertResponseBody) SetCertStatus(v string) *DescribeWarehouseCertResponseBody {
	s.CertStatus = &v
	return s
}

func (s *DescribeWarehouseCertResponseBody) SetCertType(v string) *DescribeWarehouseCertResponseBody {
	s.CertType = &v
	return s
}

func (s *DescribeWarehouseCertResponseBody) SetCommonName(v string) *DescribeWarehouseCertResponseBody {
	s.CommonName = &v
	return s
}

func (s *DescribeWarehouseCertResponseBody) SetContent(v string) *DescribeWarehouseCertResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeWarehouseCertResponseBody) SetFingerprint(v string) *DescribeWarehouseCertResponseBody {
	s.Fingerprint = &v
	return s
}

func (s *DescribeWarehouseCertResponseBody) SetIssuer(v string) *DescribeWarehouseCertResponseBody {
	s.Issuer = &v
	return s
}

func (s *DescribeWarehouseCertResponseBody) SetIssuerIdentifier(v string) *DescribeWarehouseCertResponseBody {
	s.IssuerIdentifier = &v
	return s
}

func (s *DescribeWarehouseCertResponseBody) SetPrivateCaInstanceId(v string) *DescribeWarehouseCertResponseBody {
	s.PrivateCaInstanceId = &v
	return s
}

func (s *DescribeWarehouseCertResponseBody) SetPrivateCaRegionId(v string) *DescribeWarehouseCertResponseBody {
	s.PrivateCaRegionId = &v
	return s
}

func (s *DescribeWarehouseCertResponseBody) SetRequestId(v string) *DescribeWarehouseCertResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWarehouseCertResponseBody) SetWarehouseInstanceId(v string) *DescribeWarehouseCertResponseBody {
	s.WarehouseInstanceId = &v
	return s
}

func (s *DescribeWarehouseCertResponseBody) Validate() error {
	return dara.Validate(s)
}
