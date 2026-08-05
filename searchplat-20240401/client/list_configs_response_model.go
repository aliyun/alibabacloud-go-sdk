// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListConfigsResponseBody) *ListConfigsResponse
	GetBody() *ListConfigsResponseBody
}

type ListConfigsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConfigsResponse) GetBody() *ListConfigsResponseBody {
	return s.Body
}

func (s *ListConfigsResponse) SetHeaders(v map[string]*string) *ListConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListConfigsResponse) SetStatusCode(v int32) *ListConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConfigsResponse) SetBody(v *ListConfigsResponseBody) *ListConfigsResponse {
	s.Body = v
	return s
}

func (s *ListConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
