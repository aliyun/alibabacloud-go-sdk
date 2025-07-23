// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadUserCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertId(v int64) *UploadUserCertificateResponseBody
	GetCertId() *int64
	SetRequestId(v string) *UploadUserCertificateResponseBody
	GetRequestId() *string
	SetResourceId(v string) *UploadUserCertificateResponseBody
	GetResourceId() *string
}

type UploadUserCertificateResponseBody struct {
	// The ID of the certificate.
	//
	// example:
	//
	// 12345
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BDB81BA2-E1F5-4D08-A2DD-4BE2BF44C90E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// cas-upload-j2ofdb
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s UploadUserCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadUserCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *UploadUserCertificateResponseBody) GetCertId() *int64 {
	return s.CertId
}

func (s *UploadUserCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadUserCertificateResponseBody) GetResourceId() *string {
	return s.ResourceId
}

func (s *UploadUserCertificateResponseBody) SetCertId(v int64) *UploadUserCertificateResponseBody {
	s.CertId = &v
	return s
}

func (s *UploadUserCertificateResponseBody) SetRequestId(v string) *UploadUserCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadUserCertificateResponseBody) SetResourceId(v string) *UploadUserCertificateResponseBody {
	s.ResourceId = &v
	return s
}

func (s *UploadUserCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
