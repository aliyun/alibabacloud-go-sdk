// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExperimentRunsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExperimentRunsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExperimentRunsResponse
	GetStatusCode() *int32
	SetBody(v *ListExperimentRunsResponseBody) *ListExperimentRunsResponse
	GetBody() *ListExperimentRunsResponseBody
}

type ListExperimentRunsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExperimentRunsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExperimentRunsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExperimentRunsResponse) GoString() string {
	return s.String()
}

func (s *ListExperimentRunsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExperimentRunsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExperimentRunsResponse) GetBody() *ListExperimentRunsResponseBody {
	return s.Body
}

func (s *ListExperimentRunsResponse) SetHeaders(v map[string]*string) *ListExperimentRunsResponse {
	s.Headers = v
	return s
}

func (s *ListExperimentRunsResponse) SetStatusCode(v int32) *ListExperimentRunsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExperimentRunsResponse) SetBody(v *ListExperimentRunsResponseBody) *ListExperimentRunsResponse {
	s.Body = v
	return s
}

func (s *ListExperimentRunsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
