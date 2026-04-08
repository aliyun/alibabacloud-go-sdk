// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDIJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDIJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDIJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListDIJobsResponseBody) *ListDIJobsResponse
	GetBody() *ListDIJobsResponseBody
}

type ListDIJobsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDIJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDIJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobsResponse) GoString() string {
	return s.String()
}

func (s *ListDIJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDIJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDIJobsResponse) GetBody() *ListDIJobsResponseBody {
	return s.Body
}

func (s *ListDIJobsResponse) SetHeaders(v map[string]*string) *ListDIJobsResponse {
	s.Headers = v
	return s
}

func (s *ListDIJobsResponse) SetStatusCode(v int32) *ListDIJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDIJobsResponse) SetBody(v *ListDIJobsResponseBody) *ListDIJobsResponse {
	s.Body = v
	return s
}

func (s *ListDIJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
