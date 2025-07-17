// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iImportCertificateAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateFileObject(v io.Reader) *ImportCertificateAdvanceRequest
	GetCertificateFileObject() io.Reader
	SetDescription(v string) *ImportCertificateAdvanceRequest
	GetDescription() *string
	SetName(v string) *ImportCertificateAdvanceRequest
	GetName() *string
	SetProjectId(v int64) *ImportCertificateAdvanceRequest
	GetProjectId() *int64
}

type ImportCertificateAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// -
	CertificateFileObject io.Reader `json:"CertificateFile,omitempty" xml:"CertificateFile,omitempty"`
	// example:
	//
	// This is a file
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ca1.crt
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 106560
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ImportCertificateAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportCertificateAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ImportCertificateAdvanceRequest) GetCertificateFileObject() io.Reader {
	return s.CertificateFileObject
}

func (s *ImportCertificateAdvanceRequest) GetDescription() *string {
	return s.Description
}

func (s *ImportCertificateAdvanceRequest) GetName() *string {
	return s.Name
}

func (s *ImportCertificateAdvanceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ImportCertificateAdvanceRequest) SetCertificateFileObject(v io.Reader) *ImportCertificateAdvanceRequest {
	s.CertificateFileObject = v
	return s
}

func (s *ImportCertificateAdvanceRequest) SetDescription(v string) *ImportCertificateAdvanceRequest {
	s.Description = &v
	return s
}

func (s *ImportCertificateAdvanceRequest) SetName(v string) *ImportCertificateAdvanceRequest {
	s.Name = &v
	return s
}

func (s *ImportCertificateAdvanceRequest) SetProjectId(v int64) *ImportCertificateAdvanceRequest {
	s.ProjectId = &v
	return s
}

func (s *ImportCertificateAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
