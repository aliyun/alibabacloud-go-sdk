// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAdvancedQueryTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAdvancedQueryTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAdvancedQueryTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAdvancedQueryTemplateResponseBody) *UpdateAdvancedQueryTemplateResponse
	GetBody() *UpdateAdvancedQueryTemplateResponseBody
}

type UpdateAdvancedQueryTemplateResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAdvancedQueryTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAdvancedQueryTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdvancedQueryTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateAdvancedQueryTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAdvancedQueryTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAdvancedQueryTemplateResponse) GetBody() *UpdateAdvancedQueryTemplateResponseBody {
	return s.Body
}

func (s *UpdateAdvancedQueryTemplateResponse) SetHeaders(v map[string]*string) *UpdateAdvancedQueryTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateAdvancedQueryTemplateResponse) SetStatusCode(v int32) *UpdateAdvancedQueryTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAdvancedQueryTemplateResponse) SetBody(v *UpdateAdvancedQueryTemplateResponseBody) *UpdateAdvancedQueryTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateAdvancedQueryTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
