// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaConvertJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMediaConvertJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMediaConvertJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListMediaConvertJobsResponseBody) *ListMediaConvertJobsResponse
	GetBody() *ListMediaConvertJobsResponseBody
}

type ListMediaConvertJobsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMediaConvertJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMediaConvertJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMediaConvertJobsResponse) GoString() string {
	return s.String()
}

func (s *ListMediaConvertJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMediaConvertJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMediaConvertJobsResponse) GetBody() *ListMediaConvertJobsResponseBody {
	return s.Body
}

func (s *ListMediaConvertJobsResponse) SetHeaders(v map[string]*string) *ListMediaConvertJobsResponse {
	s.Headers = v
	return s
}

func (s *ListMediaConvertJobsResponse) SetStatusCode(v int32) *ListMediaConvertJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMediaConvertJobsResponse) SetBody(v *ListMediaConvertJobsResponseBody) *ListMediaConvertJobsResponse {
	s.Body = v
	return s
}

func (s *ListMediaConvertJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
