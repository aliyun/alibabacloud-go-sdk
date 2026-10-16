// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScanSbomExportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *CreateScanSbomExportRequest
	GetFileName() *string
	SetFormat(v string) *CreateScanSbomExportRequest
	GetFormat() *string
}

type CreateScanSbomExportRequest struct {
	// The file name (including the extension) for the downloaded file. The file name is signed into the Content-Disposition header of the URL. If you leave this parameter empty, the backend generates a file name in the format of project name-sbom-<format>.<extension>. If the file name contains control characters or path separators (/ or \\), or if the length exceeds 255 characters, a 400 InvalidFileName error is returned.
	//
	// example:
	//
	// test-sbom-cyclonedx.cdx.json
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// Valid values:
	//
	// 	- cyclonedx-json
	//
	// 	- spdx-json
	//
	// This parameter is required.
	//
	// example:
	//
	// cyclonedx-json
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
}

func (s CreateScanSbomExportRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScanSbomExportRequest) GoString() string {
	return s.String()
}

func (s *CreateScanSbomExportRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateScanSbomExportRequest) GetFormat() *string {
	return s.Format
}

func (s *CreateScanSbomExportRequest) SetFileName(v string) *CreateScanSbomExportRequest {
	s.FileName = &v
	return s
}

func (s *CreateScanSbomExportRequest) SetFormat(v string) *CreateScanSbomExportRequest {
	s.Format = &v
	return s
}

func (s *CreateScanSbomExportRequest) Validate() error {
	return dara.Validate(s)
}
