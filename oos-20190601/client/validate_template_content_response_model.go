// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateTemplateContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateTemplateContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateTemplateContentResponse
	GetStatusCode() *int32
	SetBody(v *ValidateTemplateContentResponseBody) *ValidateTemplateContentResponse
	GetBody() *ValidateTemplateContentResponseBody
}

type ValidateTemplateContentResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateTemplateContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateTemplateContentResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateTemplateContentResponse) GoString() string {
	return s.String()
}

func (s *ValidateTemplateContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateTemplateContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateTemplateContentResponse) GetBody() *ValidateTemplateContentResponseBody {
	return s.Body
}

func (s *ValidateTemplateContentResponse) SetHeaders(v map[string]*string) *ValidateTemplateContentResponse {
	s.Headers = v
	return s
}

func (s *ValidateTemplateContentResponse) SetStatusCode(v int32) *ValidateTemplateContentResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateTemplateContentResponse) SetBody(v *ValidateTemplateContentResponseBody) *ValidateTemplateContentResponse {
	s.Body = v
	return s
}

func (s *ValidateTemplateContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
