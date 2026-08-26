// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateId(v int64) *RevokeCertificateRequest
	GetCertificateId() *int64
	SetInstanceId(v string) *RevokeCertificateRequest
	GetInstanceId() *string
}

type RevokeCertificateRequest struct {
	// The ID of the certificate to revoke.
	//
	// example:
	//
	// 51001
	CertificateId *int64 `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The ID of the instance. This parameter accepts only V2.0 subscription certificate instances (certificates purchased from a CA such as DigiCert or GeoTrust, with an InstanceId in the format of cas-cn-*). Uploaded certificates (cas-upload-*) are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// cas-cn-68n1mm16****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RevokeCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeCertificateRequest) GoString() string {
	return s.String()
}

func (s *RevokeCertificateRequest) GetCertificateId() *int64 {
	return s.CertificateId
}

func (s *RevokeCertificateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RevokeCertificateRequest) SetCertificateId(v int64) *RevokeCertificateRequest {
	s.CertificateId = &v
	return s
}

func (s *RevokeCertificateRequest) SetInstanceId(v string) *RevokeCertificateRequest {
	s.InstanceId = &v
	return s
}

func (s *RevokeCertificateRequest) Validate() error {
	return dara.Validate(s)
}
