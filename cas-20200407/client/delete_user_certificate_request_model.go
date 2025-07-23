// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertId(v int64) *DeleteUserCertificateRequest
	GetCertId() *int64
}

type DeleteUserCertificateRequest struct {
	// The ID of the certificate.
	//
	// >  You can call the [ListUserCertificateOrder](https://help.aliyun.com/document_detail/455804.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7562353
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
}

func (s DeleteUserCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserCertificateRequest) GetCertId() *int64 {
	return s.CertId
}

func (s *DeleteUserCertificateRequest) SetCertId(v int64) *DeleteUserCertificateRequest {
	s.CertId = &v
	return s
}

func (s *DeleteUserCertificateRequest) Validate() error {
	return dara.Validate(s)
}
