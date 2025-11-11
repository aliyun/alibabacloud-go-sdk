// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotTopicsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHotTopicsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHotTopicsResponse
	GetStatusCode() *int32
	SetBody(v *ListHotTopicsResponseBody) *ListHotTopicsResponse
	GetBody() *ListHotTopicsResponseBody
}

type ListHotTopicsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotTopicsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotTopicsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHotTopicsResponse) GoString() string {
	return s.String()
}

func (s *ListHotTopicsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHotTopicsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHotTopicsResponse) GetBody() *ListHotTopicsResponseBody {
	return s.Body
}

func (s *ListHotTopicsResponse) SetHeaders(v map[string]*string) *ListHotTopicsResponse {
	s.Headers = v
	return s
}

func (s *ListHotTopicsResponse) SetStatusCode(v int32) *ListHotTopicsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotTopicsResponse) SetBody(v *ListHotTopicsResponseBody) *ListHotTopicsResponse {
	s.Body = v
	return s
}

func (s *ListHotTopicsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
