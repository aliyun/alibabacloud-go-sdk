// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateUploadTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateUploadTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateUploadTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ValidateUploadTemplateResponseBody) *ValidateUploadTemplateResponse
	GetBody() *ValidateUploadTemplateResponseBody
}

type ValidateUploadTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateUploadTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateUploadTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateUploadTemplateResponse) GoString() string {
	return s.String()
}

func (s *ValidateUploadTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateUploadTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateUploadTemplateResponse) GetBody() *ValidateUploadTemplateResponseBody {
	return s.Body
}

func (s *ValidateUploadTemplateResponse) SetHeaders(v map[string]*string) *ValidateUploadTemplateResponse {
	s.Headers = v
	return s
}

func (s *ValidateUploadTemplateResponse) SetStatusCode(v int32) *ValidateUploadTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateUploadTemplateResponse) SetBody(v *ValidateUploadTemplateResponseBody) *ValidateUploadTemplateResponse {
	s.Body = v
	return s
}

func (s *ValidateUploadTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
