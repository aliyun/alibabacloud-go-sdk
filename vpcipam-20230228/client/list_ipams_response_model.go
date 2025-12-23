// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIpamsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIpamsResponse
	GetStatusCode() *int32
	SetBody(v *ListIpamsResponseBody) *ListIpamsResponse
	GetBody() *ListIpamsResponseBody
}

type ListIpamsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIpamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIpamsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIpamsResponse) GoString() string {
	return s.String()
}

func (s *ListIpamsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIpamsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIpamsResponse) GetBody() *ListIpamsResponseBody {
	return s.Body
}

func (s *ListIpamsResponse) SetHeaders(v map[string]*string) *ListIpamsResponse {
	s.Headers = v
	return s
}

func (s *ListIpamsResponse) SetStatusCode(v int32) *ListIpamsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpamsResponse) SetBody(v *ListIpamsResponseBody) *ListIpamsResponse {
	s.Body = v
	return s
}

func (s *ListIpamsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
