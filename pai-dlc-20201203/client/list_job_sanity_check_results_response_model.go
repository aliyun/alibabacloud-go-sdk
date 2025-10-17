// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobSanityCheckResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListJobSanityCheckResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListJobSanityCheckResultsResponse
	GetStatusCode() *int32
	SetBody(v *ListJobSanityCheckResultsResponseBody) *ListJobSanityCheckResultsResponse
	GetBody() *ListJobSanityCheckResultsResponseBody
}

type ListJobSanityCheckResultsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobSanityCheckResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobSanityCheckResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListJobSanityCheckResultsResponse) GoString() string {
	return s.String()
}

func (s *ListJobSanityCheckResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListJobSanityCheckResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListJobSanityCheckResultsResponse) GetBody() *ListJobSanityCheckResultsResponseBody {
	return s.Body
}

func (s *ListJobSanityCheckResultsResponse) SetHeaders(v map[string]*string) *ListJobSanityCheckResultsResponse {
	s.Headers = v
	return s
}

func (s *ListJobSanityCheckResultsResponse) SetStatusCode(v int32) *ListJobSanityCheckResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobSanityCheckResultsResponse) SetBody(v *ListJobSanityCheckResultsResponseBody) *ListJobSanityCheckResultsResponse {
	s.Body = v
	return s
}

func (s *ListJobSanityCheckResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
