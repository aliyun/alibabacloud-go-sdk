// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVodTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVodTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ListVodTemplateResponseBody) *ListVodTemplateResponse
	GetBody() *ListVodTemplateResponseBody
}

type ListVodTemplateResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVodTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVodTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVodTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListVodTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVodTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVodTemplateResponse) GetBody() *ListVodTemplateResponseBody {
	return s.Body
}

func (s *ListVodTemplateResponse) SetHeaders(v map[string]*string) *ListVodTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListVodTemplateResponse) SetStatusCode(v int32) *ListVodTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVodTemplateResponse) SetBody(v *ListVodTemplateResponseBody) *ListVodTemplateResponse {
	s.Body = v
	return s
}

func (s *ListVodTemplateResponse) Validate() error {
	return dara.Validate(s)
}
