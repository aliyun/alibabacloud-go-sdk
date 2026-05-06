// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSampleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSampleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSampleResponse
	GetStatusCode() *int32
	SetBody(v *ListSampleResponseBody) *ListSampleResponse
	GetBody() *ListSampleResponseBody
}

type ListSampleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSampleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSampleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSampleResponse) GoString() string {
	return s.String()
}

func (s *ListSampleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSampleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSampleResponse) GetBody() *ListSampleResponseBody {
	return s.Body
}

func (s *ListSampleResponse) SetHeaders(v map[string]*string) *ListSampleResponse {
	s.Headers = v
	return s
}

func (s *ListSampleResponse) SetStatusCode(v int32) *ListSampleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSampleResponse) SetBody(v *ListSampleResponseBody) *ListSampleResponse {
	s.Body = v
	return s
}

func (s *ListSampleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
