// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoolsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPoolsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPoolsResponse
	GetStatusCode() *int32
	SetBody(v *ListPoolsResponseBody) *ListPoolsResponse
	GetBody() *ListPoolsResponseBody
}

type ListPoolsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPoolsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPoolsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPoolsResponse) GoString() string {
	return s.String()
}

func (s *ListPoolsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPoolsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPoolsResponse) GetBody() *ListPoolsResponseBody {
	return s.Body
}

func (s *ListPoolsResponse) SetHeaders(v map[string]*string) *ListPoolsResponse {
	s.Headers = v
	return s
}

func (s *ListPoolsResponse) SetStatusCode(v int32) *ListPoolsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPoolsResponse) SetBody(v *ListPoolsResponseBody) *ListPoolsResponse {
	s.Body = v
	return s
}

func (s *ListPoolsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
