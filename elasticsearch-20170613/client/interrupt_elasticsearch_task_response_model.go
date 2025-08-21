// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInterruptElasticsearchTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InterruptElasticsearchTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InterruptElasticsearchTaskResponse
	GetStatusCode() *int32
	SetBody(v *InterruptElasticsearchTaskResponseBody) *InterruptElasticsearchTaskResponse
	GetBody() *InterruptElasticsearchTaskResponseBody
}

type InterruptElasticsearchTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InterruptElasticsearchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InterruptElasticsearchTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s InterruptElasticsearchTaskResponse) GoString() string {
	return s.String()
}

func (s *InterruptElasticsearchTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InterruptElasticsearchTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InterruptElasticsearchTaskResponse) GetBody() *InterruptElasticsearchTaskResponseBody {
	return s.Body
}

func (s *InterruptElasticsearchTaskResponse) SetHeaders(v map[string]*string) *InterruptElasticsearchTaskResponse {
	s.Headers = v
	return s
}

func (s *InterruptElasticsearchTaskResponse) SetStatusCode(v int32) *InterruptElasticsearchTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *InterruptElasticsearchTaskResponse) SetBody(v *InterruptElasticsearchTaskResponseBody) *InterruptElasticsearchTaskResponse {
	s.Body = v
	return s
}

func (s *InterruptElasticsearchTaskResponse) Validate() error {
	return dara.Validate(s)
}
