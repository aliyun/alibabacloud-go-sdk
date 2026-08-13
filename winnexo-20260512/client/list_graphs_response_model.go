// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGraphsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGraphsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGraphsResponse
	GetStatusCode() *int32
	SetBody(v *ListGraphsResponseBody) *ListGraphsResponse
	GetBody() *ListGraphsResponseBody
}

type ListGraphsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGraphsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGraphsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGraphsResponse) GoString() string {
	return s.String()
}

func (s *ListGraphsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGraphsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGraphsResponse) GetBody() *ListGraphsResponseBody {
	return s.Body
}

func (s *ListGraphsResponse) SetHeaders(v map[string]*string) *ListGraphsResponse {
	s.Headers = v
	return s
}

func (s *ListGraphsResponse) SetStatusCode(v int32) *ListGraphsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGraphsResponse) SetBody(v *ListGraphsResponseBody) *ListGraphsResponse {
	s.Body = v
	return s
}

func (s *ListGraphsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
