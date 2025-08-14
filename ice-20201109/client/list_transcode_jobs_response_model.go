// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTranscodeJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTranscodeJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTranscodeJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListTranscodeJobsResponseBody) *ListTranscodeJobsResponse
	GetBody() *ListTranscodeJobsResponseBody
}

type ListTranscodeJobsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTranscodeJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTranscodeJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeJobsResponse) GoString() string {
	return s.String()
}

func (s *ListTranscodeJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTranscodeJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTranscodeJobsResponse) GetBody() *ListTranscodeJobsResponseBody {
	return s.Body
}

func (s *ListTranscodeJobsResponse) SetHeaders(v map[string]*string) *ListTranscodeJobsResponse {
	s.Headers = v
	return s
}

func (s *ListTranscodeJobsResponse) SetStatusCode(v int32) *ListTranscodeJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTranscodeJobsResponse) SetBody(v *ListTranscodeJobsResponseBody) *ListTranscodeJobsResponse {
	s.Body = v
	return s
}

func (s *ListTranscodeJobsResponse) Validate() error {
	return dara.Validate(s)
}
