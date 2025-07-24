// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationConfigsResponseBody) *ListApplicationConfigsResponse
	GetBody() *ListApplicationConfigsResponseBody
}

type ListApplicationConfigsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationConfigsResponse) GetBody() *ListApplicationConfigsResponseBody {
	return s.Body
}

func (s *ListApplicationConfigsResponse) SetHeaders(v map[string]*string) *ListApplicationConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationConfigsResponse) SetStatusCode(v int32) *ListApplicationConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationConfigsResponse) SetBody(v *ListApplicationConfigsResponseBody) *ListApplicationConfigsResponse {
	s.Body = v
	return s
}

func (s *ListApplicationConfigsResponse) Validate() error {
	return dara.Validate(s)
}
