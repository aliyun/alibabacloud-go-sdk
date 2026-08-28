// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShareCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateId(v int64) *ShareCertificateRequest
	GetCertificateId() *int64
	SetTargetUserId(v int64) *ShareCertificateRequest
	GetTargetUserId() *int64
}

type ShareCertificateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 23279004
	CertificateId *int64 `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1622883344556677
	TargetUserId *int64 `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty"`
}

func (s ShareCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s ShareCertificateRequest) GoString() string {
	return s.String()
}

func (s *ShareCertificateRequest) GetCertificateId() *int64 {
	return s.CertificateId
}

func (s *ShareCertificateRequest) GetTargetUserId() *int64 {
	return s.TargetUserId
}

func (s *ShareCertificateRequest) SetCertificateId(v int64) *ShareCertificateRequest {
	s.CertificateId = &v
	return s
}

func (s *ShareCertificateRequest) SetTargetUserId(v int64) *ShareCertificateRequest {
	s.TargetUserId = &v
	return s
}

func (s *ShareCertificateRequest) Validate() error {
	return dara.Validate(s)
}
