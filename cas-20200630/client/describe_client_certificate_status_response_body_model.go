// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientCertificateStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateStatus(v []*DescribeClientCertificateStatusResponseBodyCertificateStatus) *DescribeClientCertificateStatusResponseBody
	GetCertificateStatus() []*DescribeClientCertificateStatusResponseBodyCertificateStatus
	SetRequestId(v string) *DescribeClientCertificateStatusResponseBody
	GetRequestId() *string
}

type DescribeClientCertificateStatusResponseBody struct {
	// An array that consists of the status information about the certificates.
	CertificateStatus []*DescribeClientCertificateStatusResponseBodyCertificateStatus `json:"CertificateStatus,omitempty" xml:"CertificateStatus,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClientCertificateStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientCertificateStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateStatusResponseBody) GetCertificateStatus() []*DescribeClientCertificateStatusResponseBodyCertificateStatus {
	return s.CertificateStatus
}

func (s *DescribeClientCertificateStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClientCertificateStatusResponseBody) SetCertificateStatus(v []*DescribeClientCertificateStatusResponseBodyCertificateStatus) *DescribeClientCertificateStatusResponseBody {
	s.CertificateStatus = v
	return s
}

func (s *DescribeClientCertificateStatusResponseBody) SetRequestId(v string) *DescribeClientCertificateStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClientCertificateStatusResponseBody) Validate() error {
	if s.CertificateStatus != nil {
		for _, item := range s.CertificateStatus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClientCertificateStatusResponseBodyCertificateStatus struct {
	// The date on which the certificate was revoked.
	//
	// >  This parameter is returned only when the value of the **Status*	- parameter is **revoked**. The value revoked indicates that the certificate is revoked.
	//
	// example:
	//
	// 2021-01-01T00:00Z
	RevokeTime *int64 `json:"RevokeTime,omitempty" xml:"RevokeTime,omitempty"`
	// The serial number of the certificate.
	//
	// example:
	//
	// b67e53ebcea9b77d65b0c3236646d715****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The status of the certificate. Valid values:
	//
	// 	- **good**: The certificate is not revoked.
	//
	// 	- **revoked**: The certificate is revoked.
	//
	// 	- **unknown**: The server cannot determine the status of the certificate.
	//
	// example:
	//
	// good
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeClientCertificateStatusResponseBodyCertificateStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientCertificateStatusResponseBodyCertificateStatus) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateStatusResponseBodyCertificateStatus) GetRevokeTime() *int64 {
	return s.RevokeTime
}

func (s *DescribeClientCertificateStatusResponseBodyCertificateStatus) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeClientCertificateStatusResponseBodyCertificateStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeClientCertificateStatusResponseBodyCertificateStatus) SetRevokeTime(v int64) *DescribeClientCertificateStatusResponseBodyCertificateStatus {
	s.RevokeTime = &v
	return s
}

func (s *DescribeClientCertificateStatusResponseBodyCertificateStatus) SetSerialNumber(v string) *DescribeClientCertificateStatusResponseBodyCertificateStatus {
	s.SerialNumber = &v
	return s
}

func (s *DescribeClientCertificateStatusResponseBodyCertificateStatus) SetStatus(v string) *DescribeClientCertificateStatusResponseBodyCertificateStatus {
	s.Status = &v
	return s
}

func (s *DescribeClientCertificateStatusResponseBodyCertificateStatus) Validate() error {
	return dara.Validate(s)
}
