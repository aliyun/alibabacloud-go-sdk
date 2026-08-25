// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveExternalSAMLIdPCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateId(v string) *RemoveExternalSAMLIdPCertificateRequest
	GetCertificateId() *string
	SetDirectoryId(v string) *RemoveExternalSAMLIdPCertificateRequest
	GetDirectoryId() *string
}

type RemoveExternalSAMLIdPCertificateRequest struct {
	// The ID of the certificate.
	//
	// You can call the [ListExternalSAMLIdPCertificates](https://help.aliyun.com/document_detail/341629.html) operation to query the IDs of certificates.
	//
	// example:
	//
	// idp-c-00dt9gnl7fmjaw9c****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s RemoveExternalSAMLIdPCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveExternalSAMLIdPCertificateRequest) GoString() string {
	return s.String()
}

func (s *RemoveExternalSAMLIdPCertificateRequest) GetCertificateId() *string {
	return s.CertificateId
}

func (s *RemoveExternalSAMLIdPCertificateRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *RemoveExternalSAMLIdPCertificateRequest) SetCertificateId(v string) *RemoveExternalSAMLIdPCertificateRequest {
	s.CertificateId = &v
	return s
}

func (s *RemoveExternalSAMLIdPCertificateRequest) SetDirectoryId(v string) *RemoveExternalSAMLIdPCertificateRequest {
	s.DirectoryId = &v
	return s
}

func (s *RemoveExternalSAMLIdPCertificateRequest) Validate() error {
	return dara.Validate(s)
}
