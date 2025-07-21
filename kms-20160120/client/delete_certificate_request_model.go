// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateId(v string) *DeleteCertificateRequest
	GetCertificateId() *string
}

type DeleteCertificateRequest struct {
	// The ID of the certificate. It is the globally unique identifier (GUID) of the certificate in Alibaba Cloud Certificate Manager.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9a28de48-8d8b-484d-a766-dec4****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s DeleteCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteCertificateRequest) GetCertificateId() *string {
	return s.CertificateId
}

func (s *DeleteCertificateRequest) SetCertificateId(v string) *DeleteCertificateRequest {
	s.CertificateId = &v
	return s
}

func (s *DeleteCertificateRequest) Validate() error {
	return dara.Validate(s)
}
