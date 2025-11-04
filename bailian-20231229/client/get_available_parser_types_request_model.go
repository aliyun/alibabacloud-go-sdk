// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAvailableParserTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileType(v string) *GetAvailableParserTypesRequest
	GetFileType() *string
}

type GetAvailableParserTypesRequest struct {
	// The file type. Valid values: pdf, docx, and doc.
	//
	// This parameter is required.
	//
	// example:
	//
	// pdf
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
}

func (s GetAvailableParserTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAvailableParserTypesRequest) GoString() string {
	return s.String()
}

func (s *GetAvailableParserTypesRequest) GetFileType() *string {
	return s.FileType
}

func (s *GetAvailableParserTypesRequest) SetFileType(v string) *GetAvailableParserTypesRequest {
	s.FileType = &v
	return s
}

func (s *GetAvailableParserTypesRequest) Validate() error {
	return dara.Validate(s)
}
