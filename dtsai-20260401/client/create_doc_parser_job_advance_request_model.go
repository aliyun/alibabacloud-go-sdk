// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iCreateDocParserJobAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileFormat(v string) *CreateDocParserJobAdvanceRequest
	GetFileFormat() *string
	SetFileName(v string) *CreateDocParserJobAdvanceRequest
	GetFileName() *string
	SetFileUrlObject(v io.Reader) *CreateDocParserJobAdvanceRequest
	GetFileUrlObject() io.Reader
	SetOutputFormat(v string) *CreateDocParserJobAdvanceRequest
	GetOutputFormat() *string
	SetRegionId(v string) *CreateDocParserJobAdvanceRequest
	GetRegionId() *string
}

type CreateDocParserJobAdvanceRequest struct {
	// This parameter is required.
	FileFormat *string `json:"FileFormat,omitempty" xml:"FileFormat,omitempty"`
	// This parameter is required.
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	FileUrlObject io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// This parameter is required.
	OutputFormat *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateDocParserJobAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDocParserJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDocParserJobAdvanceRequest) GetFileFormat() *string {
	return s.FileFormat
}

func (s *CreateDocParserJobAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateDocParserJobAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *CreateDocParserJobAdvanceRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *CreateDocParserJobAdvanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDocParserJobAdvanceRequest) SetFileFormat(v string) *CreateDocParserJobAdvanceRequest {
	s.FileFormat = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetFileName(v string) *CreateDocParserJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetFileUrlObject(v io.Reader) *CreateDocParserJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetOutputFormat(v string) *CreateDocParserJobAdvanceRequest {
	s.OutputFormat = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) SetRegionId(v string) *CreateDocParserJobAdvanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDocParserJobAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
