// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientCertificateStatusForSerialNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateStatus(v []*DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus) *DescribeClientCertificateStatusForSerialNumberResponseBody
	GetCertificateStatus() []*DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus
	SetRequestId(v string) *DescribeClientCertificateStatusForSerialNumberResponseBody
	GetRequestId() *string
}

type DescribeClientCertificateStatusForSerialNumberResponseBody struct {
	// The object.
	CertificateStatus []*DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus `json:"CertificateStatus,omitempty" xml:"CertificateStatus,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClientCertificateStatusForSerialNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientCertificateStatusForSerialNumberResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateStatusForSerialNumberResponseBody) GetCertificateStatus() []*DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus {
	return s.CertificateStatus
}

func (s *DescribeClientCertificateStatusForSerialNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClientCertificateStatusForSerialNumberResponseBody) SetCertificateStatus(v []*DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus) *DescribeClientCertificateStatusForSerialNumberResponseBody {
	s.CertificateStatus = v
	return s
}

func (s *DescribeClientCertificateStatusForSerialNumberResponseBody) SetRequestId(v string) *DescribeClientCertificateStatusForSerialNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClientCertificateStatusForSerialNumberResponseBody) Validate() error {
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

type DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus struct {
	// The date on which the certificate was revoked.
	//
	// >  This parameter is returned only when the value of the **Status*	- parameter is **revoked**. The value revoked indicates that the certificate is revoked.
	//
	// example:
	//
	// 2021-01-01T00:00
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

func (s DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus) GetRevokeTime() *int64 {
	return s.RevokeTime
}

func (s *DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus) SetRevokeTime(v int64) *DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus {
	s.RevokeTime = &v
	return s
}

func (s *DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus) SetSerialNumber(v string) *DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus {
	s.SerialNumber = &v
	return s
}

func (s *DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus) SetStatus(v string) *DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus {
	s.Status = &v
	return s
}

func (s *DescribeClientCertificateStatusForSerialNumberResponseBodyCertificateStatus) Validate() error {
	return dara.Validate(s)
}
