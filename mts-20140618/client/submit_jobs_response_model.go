// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitJobsResponse
	GetStatusCode() *int32
	SetBody(v *SubmitJobsResponseBody) *SubmitJobsResponse
	GetBody() *SubmitJobsResponseBody
}

type SubmitJobsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsResponse) GoString() string {
	return s.String()
}

func (s *SubmitJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitJobsResponse) GetBody() *SubmitJobsResponseBody {
	return s.Body
}

func (s *SubmitJobsResponse) SetHeaders(v map[string]*string) *SubmitJobsResponse {
	s.Headers = v
	return s
}

func (s *SubmitJobsResponse) SetStatusCode(v int32) *SubmitJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitJobsResponse) SetBody(v *SubmitJobsResponseBody) *SubmitJobsResponse {
	s.Body = v
	return s
}

func (s *SubmitJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
