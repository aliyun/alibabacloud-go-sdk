// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBasicAccelerateIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBasicAccelerateIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBasicAccelerateIpsResponse
	GetStatusCode() *int32
	SetBody(v *ListBasicAccelerateIpsResponseBody) *ListBasicAccelerateIpsResponse
	GetBody() *ListBasicAccelerateIpsResponseBody
}

type ListBasicAccelerateIpsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBasicAccelerateIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBasicAccelerateIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBasicAccelerateIpsResponse) GoString() string {
	return s.String()
}

func (s *ListBasicAccelerateIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBasicAccelerateIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBasicAccelerateIpsResponse) GetBody() *ListBasicAccelerateIpsResponseBody {
	return s.Body
}

func (s *ListBasicAccelerateIpsResponse) SetHeaders(v map[string]*string) *ListBasicAccelerateIpsResponse {
	s.Headers = v
	return s
}

func (s *ListBasicAccelerateIpsResponse) SetStatusCode(v int32) *ListBasicAccelerateIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBasicAccelerateIpsResponse) SetBody(v *ListBasicAccelerateIpsResponseBody) *ListBasicAccelerateIpsResponse {
	s.Body = v
	return s
}

func (s *ListBasicAccelerateIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
