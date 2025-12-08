// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignCertificateCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertCount(v int32) *AssignCertificateCountResponseBody
	GetCertCount() *int32
	SetCurrentYearFreeCertCount(v int32) *AssignCertificateCountResponseBody
	GetCurrentYearFreeCertCount() *int32
	SetRequestId(v string) *AssignCertificateCountResponseBody
	GetRequestId() *string
}

type AssignCertificateCountResponseBody struct {
	// example:
	//
	// 2
	CertCount *int32 `json:"CertCount,omitempty" xml:"CertCount,omitempty"`
	// example:
	//
	// 0
	CurrentYearFreeCertCount *int32 `json:"CurrentYearFreeCertCount,omitempty" xml:"CurrentYearFreeCertCount,omitempty"`
	// example:
	//
	// E77C4794-F24F-58CB-9929-F0F0C0EDE7B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssignCertificateCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssignCertificateCountResponseBody) GoString() string {
	return s.String()
}

func (s *AssignCertificateCountResponseBody) GetCertCount() *int32 {
	return s.CertCount
}

func (s *AssignCertificateCountResponseBody) GetCurrentYearFreeCertCount() *int32 {
	return s.CurrentYearFreeCertCount
}

func (s *AssignCertificateCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssignCertificateCountResponseBody) SetCertCount(v int32) *AssignCertificateCountResponseBody {
	s.CertCount = &v
	return s
}

func (s *AssignCertificateCountResponseBody) SetCurrentYearFreeCertCount(v int32) *AssignCertificateCountResponseBody {
	s.CurrentYearFreeCertCount = &v
	return s
}

func (s *AssignCertificateCountResponseBody) SetRequestId(v string) *AssignCertificateCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssignCertificateCountResponseBody) Validate() error {
	return dara.Validate(s)
}
