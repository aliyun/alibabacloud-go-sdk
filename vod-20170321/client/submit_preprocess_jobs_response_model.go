// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPreprocessJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitPreprocessJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitPreprocessJobsResponse
	GetStatusCode() *int32
	SetBody(v *SubmitPreprocessJobsResponseBody) *SubmitPreprocessJobsResponse
	GetBody() *SubmitPreprocessJobsResponseBody
}

type SubmitPreprocessJobsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitPreprocessJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitPreprocessJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitPreprocessJobsResponse) GoString() string {
	return s.String()
}

func (s *SubmitPreprocessJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitPreprocessJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitPreprocessJobsResponse) GetBody() *SubmitPreprocessJobsResponseBody {
	return s.Body
}

func (s *SubmitPreprocessJobsResponse) SetHeaders(v map[string]*string) *SubmitPreprocessJobsResponse {
	s.Headers = v
	return s
}

func (s *SubmitPreprocessJobsResponse) SetStatusCode(v int32) *SubmitPreprocessJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitPreprocessJobsResponse) SetBody(v *SubmitPreprocessJobsResponseBody) *SubmitPreprocessJobsResponse {
	s.Body = v
	return s
}

func (s *SubmitPreprocessJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
