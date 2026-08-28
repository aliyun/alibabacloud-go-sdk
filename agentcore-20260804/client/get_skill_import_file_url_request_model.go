// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillImportFileUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentType(v string) *GetSkillImportFileUrlRequest
	GetContentType() *string
}

type GetSkillImportFileUrlRequest struct {
	// The Content-Type of the upload file. Default value: application/zip.
	//
	// example:
	//
	// application/zip
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
}

func (s GetSkillImportFileUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillImportFileUrlRequest) GoString() string {
	return s.String()
}

func (s *GetSkillImportFileUrlRequest) GetContentType() *string {
	return s.ContentType
}

func (s *GetSkillImportFileUrlRequest) SetContentType(v string) *GetSkillImportFileUrlRequest {
	s.ContentType = &v
	return s
}

func (s *GetSkillImportFileUrlRequest) Validate() error {
	return dara.Validate(s)
}
