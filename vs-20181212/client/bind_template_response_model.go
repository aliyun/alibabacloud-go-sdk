// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindTemplateResponse
	GetStatusCode() *int32
	SetBody(v *BindTemplateResponseBody) *BindTemplateResponse
	GetBody() *BindTemplateResponseBody
}

type BindTemplateResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s BindTemplateResponse) GoString() string {
	return s.String()
}

func (s *BindTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindTemplateResponse) GetBody() *BindTemplateResponseBody {
	return s.Body
}

func (s *BindTemplateResponse) SetHeaders(v map[string]*string) *BindTemplateResponse {
	s.Headers = v
	return s
}

func (s *BindTemplateResponse) SetStatusCode(v int32) *BindTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *BindTemplateResponse) SetBody(v *BindTemplateResponseBody) *BindTemplateResponse {
	s.Body = v
	return s
}

func (s *BindTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
