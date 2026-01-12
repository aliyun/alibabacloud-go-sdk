// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSourceFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSourceFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSourceFileResponse
	GetStatusCode() *int32
	SetBody(v *ListSourceFileResponseBody) *ListSourceFileResponse
	GetBody() *ListSourceFileResponseBody
}

type ListSourceFileResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSourceFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSourceFileResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSourceFileResponse) GoString() string {
	return s.String()
}

func (s *ListSourceFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSourceFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSourceFileResponse) GetBody() *ListSourceFileResponseBody {
	return s.Body
}

func (s *ListSourceFileResponse) SetHeaders(v map[string]*string) *ListSourceFileResponse {
	s.Headers = v
	return s
}

func (s *ListSourceFileResponse) SetStatusCode(v int32) *ListSourceFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSourceFileResponse) SetBody(v *ListSourceFileResponseBody) *ListSourceFileResponse {
	s.Body = v
	return s
}

func (s *ListSourceFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
