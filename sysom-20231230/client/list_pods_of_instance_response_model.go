// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPodsOfInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPodsOfInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPodsOfInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ListPodsOfInstanceResponseBody) *ListPodsOfInstanceResponse
	GetBody() *ListPodsOfInstanceResponseBody
}

type ListPodsOfInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPodsOfInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPodsOfInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPodsOfInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListPodsOfInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPodsOfInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPodsOfInstanceResponse) GetBody() *ListPodsOfInstanceResponseBody {
	return s.Body
}

func (s *ListPodsOfInstanceResponse) SetHeaders(v map[string]*string) *ListPodsOfInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListPodsOfInstanceResponse) SetStatusCode(v int32) *ListPodsOfInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPodsOfInstanceResponse) SetBody(v *ListPodsOfInstanceResponseBody) *ListPodsOfInstanceResponse {
	s.Body = v
	return s
}

func (s *ListPodsOfInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
