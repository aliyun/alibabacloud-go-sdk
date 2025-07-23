// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserCertificateDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertFilter(v bool) *GetUserCertificateDetailRequest
	GetCertFilter() *bool
	SetCertId(v int64) *GetUserCertificateDetailRequest
	GetCertId() *int64
}

type GetUserCertificateDetailRequest struct {
	// Specifies whether to filter return results. Valid values: true and false. Default value: false. **true*	- specifies that the Cert, Key, EncryptCert, EncryptPrivateKey, SignCert, and SignPrivateKey parameters are not returned. **false*	- specifies that the parameters are returned.
	//
	// example:
	//
	// false
	CertFilter *bool `json:"CertFilter,omitempty" xml:"CertFilter,omitempty"`
	// The ID of the certificate.
	//
	// >  You can call the [ListUserCertificateOrder](https://help.aliyun.com/document_detail/455804.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6055048
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
}

func (s GetUserCertificateDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserCertificateDetailRequest) GoString() string {
	return s.String()
}

func (s *GetUserCertificateDetailRequest) GetCertFilter() *bool {
	return s.CertFilter
}

func (s *GetUserCertificateDetailRequest) GetCertId() *int64 {
	return s.CertId
}

func (s *GetUserCertificateDetailRequest) SetCertFilter(v bool) *GetUserCertificateDetailRequest {
	s.CertFilter = &v
	return s
}

func (s *GetUserCertificateDetailRequest) SetCertId(v int64) *GetUserCertificateDetailRequest {
	s.CertId = &v
	return s
}

func (s *GetUserCertificateDetailRequest) Validate() error {
	return dara.Validate(s)
}
