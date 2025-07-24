// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunApiTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunApiTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunApiTemplateResponse
	GetStatusCode() *int32
	SetBody(v *RunApiTemplateResponseBody) *RunApiTemplateResponse
	GetBody() *RunApiTemplateResponseBody
}

type RunApiTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunApiTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunApiTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s RunApiTemplateResponse) GoString() string {
	return s.String()
}

func (s *RunApiTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunApiTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunApiTemplateResponse) GetBody() *RunApiTemplateResponseBody {
	return s.Body
}

func (s *RunApiTemplateResponse) SetHeaders(v map[string]*string) *RunApiTemplateResponse {
	s.Headers = v
	return s
}

func (s *RunApiTemplateResponse) SetStatusCode(v int32) *RunApiTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *RunApiTemplateResponse) SetBody(v *RunApiTemplateResponseBody) *RunApiTemplateResponse {
	s.Body = v
	return s
}

func (s *RunApiTemplateResponse) Validate() error {
	return dara.Validate(s)
}
