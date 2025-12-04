// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAdvancedQueryTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAdvancedQueryTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAdvancedQueryTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateAdvancedQueryTemplateResponseBody) *CreateAdvancedQueryTemplateResponse
	GetBody() *CreateAdvancedQueryTemplateResponseBody
}

type CreateAdvancedQueryTemplateResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAdvancedQueryTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAdvancedQueryTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAdvancedQueryTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateAdvancedQueryTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAdvancedQueryTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAdvancedQueryTemplateResponse) GetBody() *CreateAdvancedQueryTemplateResponseBody {
	return s.Body
}

func (s *CreateAdvancedQueryTemplateResponse) SetHeaders(v map[string]*string) *CreateAdvancedQueryTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateAdvancedQueryTemplateResponse) SetStatusCode(v int32) *CreateAdvancedQueryTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAdvancedQueryTemplateResponse) SetBody(v *CreateAdvancedQueryTemplateResponseBody) *CreateAdvancedQueryTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateAdvancedQueryTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
