// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveTranscodeJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLiveTranscodeJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLiveTranscodeJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListLiveTranscodeJobsResponseBody) *ListLiveTranscodeJobsResponse
	GetBody() *ListLiveTranscodeJobsResponseBody
}

type ListLiveTranscodeJobsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLiveTranscodeJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLiveTranscodeJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLiveTranscodeJobsResponse) GoString() string {
	return s.String()
}

func (s *ListLiveTranscodeJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLiveTranscodeJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLiveTranscodeJobsResponse) GetBody() *ListLiveTranscodeJobsResponseBody {
	return s.Body
}

func (s *ListLiveTranscodeJobsResponse) SetHeaders(v map[string]*string) *ListLiveTranscodeJobsResponse {
	s.Headers = v
	return s
}

func (s *ListLiveTranscodeJobsResponse) SetStatusCode(v int32) *ListLiveTranscodeJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLiveTranscodeJobsResponse) SetBody(v *ListLiveTranscodeJobsResponseBody) *ListLiveTranscodeJobsResponse {
	s.Body = v
	return s
}

func (s *ListLiveTranscodeJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
