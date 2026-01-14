// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBasicAcceleratorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBasicAcceleratorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBasicAcceleratorsResponse
	GetStatusCode() *int32
	SetBody(v *ListBasicAcceleratorsResponseBody) *ListBasicAcceleratorsResponse
	GetBody() *ListBasicAcceleratorsResponseBody
}

type ListBasicAcceleratorsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBasicAcceleratorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBasicAcceleratorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBasicAcceleratorsResponse) GoString() string {
	return s.String()
}

func (s *ListBasicAcceleratorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBasicAcceleratorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBasicAcceleratorsResponse) GetBody() *ListBasicAcceleratorsResponseBody {
	return s.Body
}

func (s *ListBasicAcceleratorsResponse) SetHeaders(v map[string]*string) *ListBasicAcceleratorsResponse {
	s.Headers = v
	return s
}

func (s *ListBasicAcceleratorsResponse) SetStatusCode(v int32) *ListBasicAcceleratorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBasicAcceleratorsResponse) SetBody(v *ListBasicAcceleratorsResponseBody) *ListBasicAcceleratorsResponse {
	s.Body = v
	return s
}

func (s *ListBasicAcceleratorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
