// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFileResponse
	GetStatusCode() *int32
	SetBody(v *ListFileResponseBody) *ListFileResponse
	GetBody() *ListFileResponseBody
}

type ListFileResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFileResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFileResponse) GoString() string {
	return s.String()
}

func (s *ListFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFileResponse) GetBody() *ListFileResponseBody {
	return s.Body
}

func (s *ListFileResponse) SetHeaders(v map[string]*string) *ListFileResponse {
	s.Headers = v
	return s
}

func (s *ListFileResponse) SetStatusCode(v int32) *ListFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFileResponse) SetBody(v *ListFileResponseBody) *ListFileResponse {
	s.Body = v
	return s
}

func (s *ListFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
