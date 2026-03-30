// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBackgroundMusicsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBackgroundMusicsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBackgroundMusicsResponse
	GetStatusCode() *int32
	SetBody(v *ListBackgroundMusicsResponseBody) *ListBackgroundMusicsResponse
	GetBody() *ListBackgroundMusicsResponseBody
}

type ListBackgroundMusicsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBackgroundMusicsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBackgroundMusicsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBackgroundMusicsResponse) GoString() string {
	return s.String()
}

func (s *ListBackgroundMusicsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBackgroundMusicsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBackgroundMusicsResponse) GetBody() *ListBackgroundMusicsResponseBody {
	return s.Body
}

func (s *ListBackgroundMusicsResponse) SetHeaders(v map[string]*string) *ListBackgroundMusicsResponse {
	s.Headers = v
	return s
}

func (s *ListBackgroundMusicsResponse) SetStatusCode(v int32) *ListBackgroundMusicsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBackgroundMusicsResponse) SetBody(v *ListBackgroundMusicsResponseBody) *ListBackgroundMusicsResponse {
	s.Body = v
	return s
}

func (s *ListBackgroundMusicsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
