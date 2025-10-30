// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeDownStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNodeDownStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNodeDownStreamResponse
	GetStatusCode() *int32
	SetBody(v *ListNodeDownStreamResponseBody) *ListNodeDownStreamResponse
	GetBody() *ListNodeDownStreamResponseBody
}

type ListNodeDownStreamResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodeDownStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodeDownStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDownStreamResponse) GoString() string {
	return s.String()
}

func (s *ListNodeDownStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNodeDownStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNodeDownStreamResponse) GetBody() *ListNodeDownStreamResponseBody {
	return s.Body
}

func (s *ListNodeDownStreamResponse) SetHeaders(v map[string]*string) *ListNodeDownStreamResponse {
	s.Headers = v
	return s
}

func (s *ListNodeDownStreamResponse) SetStatusCode(v int32) *ListNodeDownStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeDownStreamResponse) SetBody(v *ListNodeDownStreamResponseBody) *ListNodeDownStreamResponse {
	s.Body = v
	return s
}

func (s *ListNodeDownStreamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
