// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentJobCertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDeploymentJobCertResponseBodyData) *ListDeploymentJobCertResponseBody
	GetData() []*ListDeploymentJobCertResponseBodyData
	SetRequestId(v string) *ListDeploymentJobCertResponseBody
	GetRequestId() *string
}

type ListDeploymentJobCertResponseBody struct {
	// The response parameters.
	Data []*ListDeploymentJobCertResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeploymentJobCertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentJobCertResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeploymentJobCertResponseBody) GetData() []*ListDeploymentJobCertResponseBodyData {
	return s.Data
}

func (s *ListDeploymentJobCertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeploymentJobCertResponseBody) SetData(v []*ListDeploymentJobCertResponseBodyData) *ListDeploymentJobCertResponseBody {
	s.Data = v
	return s
}

func (s *ListDeploymentJobCertResponseBody) SetRequestId(v string) *ListDeploymentJobCertResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeploymentJobCertResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDeploymentJobCertResponseBodyData struct {
	// The algorithm of the certificate public key.
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The ID of the certificate.
	//
	// example:
	//
	// 11174100
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The instance ID of the certificate.
	//
	// example:
	//
	// cas-ivauto-2crxzi
	CertInstanceId *string `json:"CertInstanceId,omitempty" xml:"CertInstanceId,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// edkog.shop
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The type of the certificate order. Valid values:
	//
	// 	- **upload**: uploaded certificate.
	//
	// 	- **buy**: purchased certificate.
	//
	// 	- **free**: free certificate. This value is available only on the China site (aliyun.com).
	//
	// example:
	//
	// buy
	CertOrderType *string `json:"CertOrderType,omitempty" xml:"CertOrderType,omitempty"`
	// The type of the certificate.
	//
	// example:
	//
	// DV
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The common name of the certificate.
	//
	// example:
	//
	// vaultwebhook.vault-webhook.svc
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// Indicates whether the certificate is hosted. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsTrustee *bool `json:"IsTrustee,omitempty" xml:"IsTrustee,omitempty"`
	// The month in which the certificate is applied for.
	//
	// example:
	//
	// 12
	Month *int32 `json:"Month,omitempty" xml:"Month,omitempty"`
	// The end time of the validity period of the certificate. The value is a timestamp in seconds.
	//
	// example:
	//
	// 1683625266108
	NotAfterTime *int64 `json:"NotAfterTime,omitempty" xml:"NotAfterTime,omitempty"`
	// The start time of the validity period of the certificate. The value is a timestamp in seconds.
	//
	// example:
	//
	// 1683625266108
	NotBeforeTime *int64 `json:"NotBeforeTime,omitempty" xml:"NotBeforeTime,omitempty"`
	// The ID of the certificate order.
	//
	// >  If CertId is returned, this parameter is not returned.
	//
	// example:
	//
	// 6127067
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The subject alternative name (SAN) extensions of the certificate.
	Sans []*string `json:"Sans,omitempty" xml:"Sans,omitempty" type:"Repeated"`
	// The status code of the certificate. Valid values:
	//
	// 	- **payed**: paid and pending application
	//
	// 	- **checking**: being validated
	//
	// 	- **checkedFail**: validation failed
	//
	// 	- **revoked**: revoked
	//
	// 	- **revokeChecking**: revocation request being validated
	//
	// 	- **issued**: issued (excluding hosted certificates that are issued, certificates that are about to expire, expired certificates, and uploaded certificates)
	//
	// 	- **trustee**: hosted and issued
	//
	// 	- **upload**: uploaded (excluding certificates that are about to expire and expired certificates)
	//
	// 	- **willExpired**: about to expire (including certificates issued by using the Certificate Management Service console and uploaded certificates)
	//
	// 	- **expired**: expired (including certificates issued by using the Certificate Management Service console and uploaded certificates)
	//
	// 	- **validity**: valid (including certificates that are not expired or revoked)
	//
	// 	- **refund**: refunded
	//
	// 	- **closed**: closed
	//
	// example:
	//
	// issued
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ListDeploymentJobCertResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentJobCertResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeploymentJobCertResponseBodyData) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *ListDeploymentJobCertResponseBodyData) GetCertId() *int64 {
	return s.CertId
}

func (s *ListDeploymentJobCertResponseBodyData) GetCertInstanceId() *string {
	return s.CertInstanceId
}

func (s *ListDeploymentJobCertResponseBodyData) GetCertName() *string {
	return s.CertName
}

func (s *ListDeploymentJobCertResponseBodyData) GetCertOrderType() *string {
	return s.CertOrderType
}

func (s *ListDeploymentJobCertResponseBodyData) GetCertType() *string {
	return s.CertType
}

func (s *ListDeploymentJobCertResponseBodyData) GetCommonName() *string {
	return s.CommonName
}

func (s *ListDeploymentJobCertResponseBodyData) GetIsTrustee() *bool {
	return s.IsTrustee
}

func (s *ListDeploymentJobCertResponseBodyData) GetMonth() *int32 {
	return s.Month
}

func (s *ListDeploymentJobCertResponseBodyData) GetNotAfterTime() *int64 {
	return s.NotAfterTime
}

func (s *ListDeploymentJobCertResponseBodyData) GetNotBeforeTime() *int64 {
	return s.NotBeforeTime
}

func (s *ListDeploymentJobCertResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ListDeploymentJobCertResponseBodyData) GetSans() []*string {
	return s.Sans
}

func (s *ListDeploymentJobCertResponseBodyData) GetStatusCode() *string {
	return s.StatusCode
}

func (s *ListDeploymentJobCertResponseBodyData) SetAlgorithm(v string) *ListDeploymentJobCertResponseBodyData {
	s.Algorithm = &v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetCertId(v int64) *ListDeploymentJobCertResponseBodyData {
	s.CertId = &v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetCertInstanceId(v string) *ListDeploymentJobCertResponseBodyData {
	s.CertInstanceId = &v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetCertName(v string) *ListDeploymentJobCertResponseBodyData {
	s.CertName = &v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetCertOrderType(v string) *ListDeploymentJobCertResponseBodyData {
	s.CertOrderType = &v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetCertType(v string) *ListDeploymentJobCertResponseBodyData {
	s.CertType = &v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetCommonName(v string) *ListDeploymentJobCertResponseBodyData {
	s.CommonName = &v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetIsTrustee(v bool) *ListDeploymentJobCertResponseBodyData {
	s.IsTrustee = &v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetMonth(v int32) *ListDeploymentJobCertResponseBodyData {
	s.Month = &v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetNotAfterTime(v int64) *ListDeploymentJobCertResponseBodyData {
	s.NotAfterTime = &v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetNotBeforeTime(v int64) *ListDeploymentJobCertResponseBodyData {
	s.NotBeforeTime = &v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetOrderId(v int64) *ListDeploymentJobCertResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetSans(v []*string) *ListDeploymentJobCertResponseBodyData {
	s.Sans = v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) SetStatusCode(v string) *ListDeploymentJobCertResponseBodyData {
	s.StatusCode = &v
	return s
}

func (s *ListDeploymentJobCertResponseBodyData) Validate() error {
	return dara.Validate(s)
}
