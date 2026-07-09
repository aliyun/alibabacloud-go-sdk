// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvaluationRunsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEvaluationRunsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEvaluationRunsResponse
	GetStatusCode() *int32
	SetBody(v *ListEvaluationRunsResponseBody) *ListEvaluationRunsResponse
	GetBody() *ListEvaluationRunsResponseBody
}

type ListEvaluationRunsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEvaluationRunsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEvaluationRunsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationRunsResponse) GoString() string {
	return s.String()
}

func (s *ListEvaluationRunsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEvaluationRunsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEvaluationRunsResponse) GetBody() *ListEvaluationRunsResponseBody {
	return s.Body
}

func (s *ListEvaluationRunsResponse) SetHeaders(v map[string]*string) *ListEvaluationRunsResponse {
	s.Headers = v
	return s
}

func (s *ListEvaluationRunsResponse) SetStatusCode(v int32) *ListEvaluationRunsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEvaluationRunsResponse) SetBody(v *ListEvaluationRunsResponseBody) *ListEvaluationRunsResponse {
	s.Body = v
	return s
}

func (s *ListEvaluationRunsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
