// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagOptionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTagOptionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTagOptionsResponse
	GetStatusCode() *int32
	SetBody(v *ListTagOptionsResponseBody) *ListTagOptionsResponse
	GetBody() *ListTagOptionsResponseBody
}

type ListTagOptionsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagOptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagOptionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTagOptionsResponse) GoString() string {
	return s.String()
}

func (s *ListTagOptionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTagOptionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTagOptionsResponse) GetBody() *ListTagOptionsResponseBody {
	return s.Body
}

func (s *ListTagOptionsResponse) SetHeaders(v map[string]*string) *ListTagOptionsResponse {
	s.Headers = v
	return s
}

func (s *ListTagOptionsResponse) SetStatusCode(v int32) *ListTagOptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagOptionsResponse) SetBody(v *ListTagOptionsResponseBody) *ListTagOptionsResponse {
	s.Body = v
	return s
}

func (s *ListTagOptionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
