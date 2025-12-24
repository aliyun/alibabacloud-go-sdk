// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateFileImportTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateFileImportTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateFileImportTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GenerateFileImportTemplateResponseBody) *GenerateFileImportTemplateResponse
	GetBody() *GenerateFileImportTemplateResponseBody
}

type GenerateFileImportTemplateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateFileImportTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateFileImportTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateFileImportTemplateResponse) GoString() string {
	return s.String()
}

func (s *GenerateFileImportTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateFileImportTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateFileImportTemplateResponse) GetBody() *GenerateFileImportTemplateResponseBody {
	return s.Body
}

func (s *GenerateFileImportTemplateResponse) SetHeaders(v map[string]*string) *GenerateFileImportTemplateResponse {
	s.Headers = v
	return s
}

func (s *GenerateFileImportTemplateResponse) SetStatusCode(v int32) *GenerateFileImportTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateFileImportTemplateResponse) SetBody(v *GenerateFileImportTemplateResponseBody) *GenerateFileImportTemplateResponse {
	s.Body = v
	return s
}

func (s *GenerateFileImportTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
