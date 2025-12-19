// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValuateTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValuateTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValuateTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ValuateTemplateResponseBody) *ValuateTemplateResponse
	GetBody() *ValuateTemplateResponseBody
}

type ValuateTemplateResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValuateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValuateTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ValuateTemplateResponse) GoString() string {
	return s.String()
}

func (s *ValuateTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValuateTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValuateTemplateResponse) GetBody() *ValuateTemplateResponseBody {
	return s.Body
}

func (s *ValuateTemplateResponse) SetHeaders(v map[string]*string) *ValuateTemplateResponse {
	s.Headers = v
	return s
}

func (s *ValuateTemplateResponse) SetStatusCode(v int32) *ValuateTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ValuateTemplateResponse) SetBody(v *ValuateTemplateResponseBody) *ValuateTemplateResponse {
	s.Body = v
	return s
}

func (s *ValuateTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
