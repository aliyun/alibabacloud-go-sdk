// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProduceEditingProjectVideoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ProduceEditingProjectVideoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ProduceEditingProjectVideoResponse
	GetStatusCode() *int32
	SetBody(v *ProduceEditingProjectVideoResponseBody) *ProduceEditingProjectVideoResponse
	GetBody() *ProduceEditingProjectVideoResponseBody
}

type ProduceEditingProjectVideoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProduceEditingProjectVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ProduceEditingProjectVideoResponse) String() string {
	return dara.Prettify(s)
}

func (s ProduceEditingProjectVideoResponse) GoString() string {
	return s.String()
}

func (s *ProduceEditingProjectVideoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ProduceEditingProjectVideoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ProduceEditingProjectVideoResponse) GetBody() *ProduceEditingProjectVideoResponseBody {
	return s.Body
}

func (s *ProduceEditingProjectVideoResponse) SetHeaders(v map[string]*string) *ProduceEditingProjectVideoResponse {
	s.Headers = v
	return s
}

func (s *ProduceEditingProjectVideoResponse) SetStatusCode(v int32) *ProduceEditingProjectVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *ProduceEditingProjectVideoResponse) SetBody(v *ProduceEditingProjectVideoResponseBody) *ProduceEditingProjectVideoResponse {
	s.Body = v
	return s
}

func (s *ProduceEditingProjectVideoResponse) Validate() error {
	return dara.Validate(s)
}
