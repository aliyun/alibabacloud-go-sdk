// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApiInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApiInfosResponse
	GetStatusCode() *int32
	SetBody(v *ListApiInfosResponseBody) *ListApiInfosResponse
	GetBody() *ListApiInfosResponseBody
}

type ListApiInfosResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApiInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApiInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApiInfosResponse) GoString() string {
	return s.String()
}

func (s *ListApiInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApiInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApiInfosResponse) GetBody() *ListApiInfosResponseBody {
	return s.Body
}

func (s *ListApiInfosResponse) SetHeaders(v map[string]*string) *ListApiInfosResponse {
	s.Headers = v
	return s
}

func (s *ListApiInfosResponse) SetStatusCode(v int32) *ListApiInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApiInfosResponse) SetBody(v *ListApiInfosResponseBody) *ListApiInfosResponse {
	s.Body = v
	return s
}

func (s *ListApiInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
