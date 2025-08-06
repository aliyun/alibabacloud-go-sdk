// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTranscodeJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitTranscodeJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitTranscodeJobsResponse
	GetStatusCode() *int32
	SetBody(v *SubmitTranscodeJobsResponseBody) *SubmitTranscodeJobsResponse
	GetBody() *SubmitTranscodeJobsResponseBody
}

type SubmitTranscodeJobsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitTranscodeJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitTranscodeJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobsResponse) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitTranscodeJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitTranscodeJobsResponse) GetBody() *SubmitTranscodeJobsResponseBody {
	return s.Body
}

func (s *SubmitTranscodeJobsResponse) SetHeaders(v map[string]*string) *SubmitTranscodeJobsResponse {
	s.Headers = v
	return s
}

func (s *SubmitTranscodeJobsResponse) SetStatusCode(v int32) *SubmitTranscodeJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitTranscodeJobsResponse) SetBody(v *SubmitTranscodeJobsResponseBody) *SubmitTranscodeJobsResponse {
	s.Body = v
	return s
}

func (s *SubmitTranscodeJobsResponse) Validate() error {
	return dara.Validate(s)
}
