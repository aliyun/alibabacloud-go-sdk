// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSmartJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSmartJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSmartJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListSmartJobsResponseBody) *ListSmartJobsResponse
	GetBody() *ListSmartJobsResponseBody
}

type ListSmartJobsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSmartJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSmartJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSmartJobsResponse) GoString() string {
	return s.String()
}

func (s *ListSmartJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSmartJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSmartJobsResponse) GetBody() *ListSmartJobsResponseBody {
	return s.Body
}

func (s *ListSmartJobsResponse) SetHeaders(v map[string]*string) *ListSmartJobsResponse {
	s.Headers = v
	return s
}

func (s *ListSmartJobsResponse) SetStatusCode(v int32) *ListSmartJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSmartJobsResponse) SetBody(v *ListSmartJobsResponseBody) *ListSmartJobsResponse {
	s.Body = v
	return s
}

func (s *ListSmartJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
