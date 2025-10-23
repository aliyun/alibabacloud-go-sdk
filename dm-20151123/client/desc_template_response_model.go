// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DescTemplateResponseBody) *DescTemplateResponse
	GetBody() *DescTemplateResponseBody
}

type DescTemplateResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DescTemplateResponse) GoString() string {
	return s.String()
}

func (s *DescTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescTemplateResponse) GetBody() *DescTemplateResponseBody {
	return s.Body
}

func (s *DescTemplateResponse) SetHeaders(v map[string]*string) *DescTemplateResponse {
	s.Headers = v
	return s
}

func (s *DescTemplateResponse) SetStatusCode(v int32) *DescTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescTemplateResponse) SetBody(v *DescTemplateResponseBody) *DescTemplateResponse {
	s.Body = v
	return s
}

func (s *DescTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
