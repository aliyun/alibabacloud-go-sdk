// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTopicsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTopicsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTopicsResponse
	GetStatusCode() *int32
	SetBody(v *ListTopicsResponseBody) *ListTopicsResponse
	GetBody() *ListTopicsResponseBody
}

type ListTopicsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTopicsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTopicsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTopicsResponse) GoString() string {
	return s.String()
}

func (s *ListTopicsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTopicsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTopicsResponse) GetBody() *ListTopicsResponseBody {
	return s.Body
}

func (s *ListTopicsResponse) SetHeaders(v map[string]*string) *ListTopicsResponse {
	s.Headers = v
	return s
}

func (s *ListTopicsResponse) SetStatusCode(v int32) *ListTopicsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTopicsResponse) SetBody(v *ListTopicsResponseBody) *ListTopicsResponse {
	s.Body = v
	return s
}

func (s *ListTopicsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
