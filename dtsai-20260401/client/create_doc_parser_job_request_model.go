// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDocParserJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileFormat(v string) *CreateDocParserJobRequest
	GetFileFormat() *string
	SetFileName(v string) *CreateDocParserJobRequest
	GetFileName() *string
	SetFileUrl(v string) *CreateDocParserJobRequest
	GetFileUrl() *string
	SetOutputFormat(v string) *CreateDocParserJobRequest
	GetOutputFormat() *string
	SetRegionId(v string) *CreateDocParserJobRequest
	GetRegionId() *string
}

type CreateDocParserJobRequest struct {
	// This parameter is required.
	FileFormat *string `json:"FileFormat,omitempty" xml:"FileFormat,omitempty"`
	// This parameter is required.
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// This parameter is required.
	OutputFormat *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateDocParserJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDocParserJobRequest) GoString() string {
	return s.String()
}

func (s *CreateDocParserJobRequest) GetFileFormat() *string {
	return s.FileFormat
}

func (s *CreateDocParserJobRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateDocParserJobRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *CreateDocParserJobRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *CreateDocParserJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDocParserJobRequest) SetFileFormat(v string) *CreateDocParserJobRequest {
	s.FileFormat = &v
	return s
}

func (s *CreateDocParserJobRequest) SetFileName(v string) *CreateDocParserJobRequest {
	s.FileName = &v
	return s
}

func (s *CreateDocParserJobRequest) SetFileUrl(v string) *CreateDocParserJobRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateDocParserJobRequest) SetOutputFormat(v string) *CreateDocParserJobRequest {
	s.OutputFormat = &v
	return s
}

func (s *CreateDocParserJobRequest) SetRegionId(v string) *CreateDocParserJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDocParserJobRequest) Validate() error {
	return dara.Validate(s)
}
