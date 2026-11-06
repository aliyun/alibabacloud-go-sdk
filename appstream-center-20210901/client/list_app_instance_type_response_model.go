// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppInstanceTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppInstanceTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppInstanceTypeResponse
	GetStatusCode() *int32
	SetBody(v *ListAppInstanceTypeResponseBody) *ListAppInstanceTypeResponse
	GetBody() *ListAppInstanceTypeResponseBody
}

type ListAppInstanceTypeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppInstanceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppInstanceTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceTypeResponse) GoString() string {
	return s.String()
}

func (s *ListAppInstanceTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppInstanceTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppInstanceTypeResponse) GetBody() *ListAppInstanceTypeResponseBody {
	return s.Body
}

func (s *ListAppInstanceTypeResponse) SetHeaders(v map[string]*string) *ListAppInstanceTypeResponse {
	s.Headers = v
	return s
}

func (s *ListAppInstanceTypeResponse) SetStatusCode(v int32) *ListAppInstanceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppInstanceTypeResponse) SetBody(v *ListAppInstanceTypeResponseBody) *ListAppInstanceTypeResponse {
	s.Body = v
	return s
}

func (s *ListAppInstanceTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
