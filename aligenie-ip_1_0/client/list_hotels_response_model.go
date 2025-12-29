// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHotelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHotelsResponse
	GetStatusCode() *int32
	SetBody(v *ListHotelsResponseBody) *ListHotelsResponse
	GetBody() *ListHotelsResponseBody
}

type ListHotelsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHotelsResponse) GoString() string {
	return s.String()
}

func (s *ListHotelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHotelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHotelsResponse) GetBody() *ListHotelsResponseBody {
	return s.Body
}

func (s *ListHotelsResponse) SetHeaders(v map[string]*string) *ListHotelsResponse {
	s.Headers = v
	return s
}

func (s *ListHotelsResponse) SetStatusCode(v int32) *ListHotelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelsResponse) SetBody(v *ListHotelsResponseBody) *ListHotelsResponse {
	s.Body = v
	return s
}

func (s *ListHotelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
