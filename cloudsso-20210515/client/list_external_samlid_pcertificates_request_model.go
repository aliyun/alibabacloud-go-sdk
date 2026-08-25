// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExternalSAMLIdPCertificatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListExternalSAMLIdPCertificatesRequest
	GetDirectoryId() *string
}

type ListExternalSAMLIdPCertificatesRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s ListExternalSAMLIdPCertificatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExternalSAMLIdPCertificatesRequest) GoString() string {
	return s.String()
}

func (s *ListExternalSAMLIdPCertificatesRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListExternalSAMLIdPCertificatesRequest) SetDirectoryId(v string) *ListExternalSAMLIdPCertificatesRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListExternalSAMLIdPCertificatesRequest) Validate() error {
	return dara.Validate(s)
}
