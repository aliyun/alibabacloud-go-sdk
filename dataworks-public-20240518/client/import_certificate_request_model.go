// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateFile(v string) *ImportCertificateRequest
	GetCertificateFile() *string
	SetDescription(v string) *ImportCertificateRequest
	GetDescription() *string
	SetName(v string) *ImportCertificateRequest
	GetName() *string
	SetProjectId(v int64) *ImportCertificateRequest
	GetProjectId() *int64
}

type ImportCertificateRequest struct {
	// The certificate file to upload. Upload method: Upload the file by creating an InputStream.
	//
	// This parameter is required.
	//
	// example:
	//
	// -
	CertificateFile *string `json:"CertificateFile,omitempty" xml:"CertificateFile,omitempty"`
	// The description of the task.
	//
	// example:
	//
	// This is a file
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The certificate file name. In a project workspace, certificate file names must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca1.crt
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the workspace to which the certificate file belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 106560
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ImportCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportCertificateRequest) GoString() string {
	return s.String()
}

func (s *ImportCertificateRequest) GetCertificateFile() *string {
	return s.CertificateFile
}

func (s *ImportCertificateRequest) GetDescription() *string {
	return s.Description
}

func (s *ImportCertificateRequest) GetName() *string {
	return s.Name
}

func (s *ImportCertificateRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ImportCertificateRequest) SetCertificateFile(v string) *ImportCertificateRequest {
	s.CertificateFile = &v
	return s
}

func (s *ImportCertificateRequest) SetDescription(v string) *ImportCertificateRequest {
	s.Description = &v
	return s
}

func (s *ImportCertificateRequest) SetName(v string) *ImportCertificateRequest {
	s.Name = &v
	return s
}

func (s *ImportCertificateRequest) SetProjectId(v int64) *ImportCertificateRequest {
	s.ProjectId = &v
	return s
}

func (s *ImportCertificateRequest) Validate() error {
	return dara.Validate(s)
}
