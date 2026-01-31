// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCertificateDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateId(v int64) *GetCertificateDetailRequest
	GetCertificateId() *int64
}

type GetCertificateDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 18594156
	CertificateId *int64 `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s GetCertificateDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCertificateDetailRequest) GoString() string {
	return s.String()
}

func (s *GetCertificateDetailRequest) GetCertificateId() *int64 {
	return s.CertificateId
}

func (s *GetCertificateDetailRequest) SetCertificateId(v int64) *GetCertificateDetailRequest {
	s.CertificateId = &v
	return s
}

func (s *GetCertificateDetailRequest) Validate() error {
	return dara.Validate(s)
}
