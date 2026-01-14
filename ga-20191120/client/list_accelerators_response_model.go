// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAcceleratorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAcceleratorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAcceleratorsResponse
	GetStatusCode() *int32
	SetBody(v *ListAcceleratorsResponseBody) *ListAcceleratorsResponse
	GetBody() *ListAcceleratorsResponseBody
}

type ListAcceleratorsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAcceleratorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAcceleratorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAcceleratorsResponse) GoString() string {
	return s.String()
}

func (s *ListAcceleratorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAcceleratorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAcceleratorsResponse) GetBody() *ListAcceleratorsResponseBody {
	return s.Body
}

func (s *ListAcceleratorsResponse) SetHeaders(v map[string]*string) *ListAcceleratorsResponse {
	s.Headers = v
	return s
}

func (s *ListAcceleratorsResponse) SetStatusCode(v int32) *ListAcceleratorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAcceleratorsResponse) SetBody(v *ListAcceleratorsResponseBody) *ListAcceleratorsResponse {
	s.Body = v
	return s
}

func (s *ListAcceleratorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
