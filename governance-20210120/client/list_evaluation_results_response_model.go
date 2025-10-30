// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvaluationResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEvaluationResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEvaluationResultsResponse
	GetStatusCode() *int32
	SetBody(v *ListEvaluationResultsResponseBody) *ListEvaluationResultsResponse
	GetBody() *ListEvaluationResultsResponseBody
}

type ListEvaluationResultsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEvaluationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEvaluationResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationResultsResponse) GoString() string {
	return s.String()
}

func (s *ListEvaluationResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEvaluationResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEvaluationResultsResponse) GetBody() *ListEvaluationResultsResponseBody {
	return s.Body
}

func (s *ListEvaluationResultsResponse) SetHeaders(v map[string]*string) *ListEvaluationResultsResponse {
	s.Headers = v
	return s
}

func (s *ListEvaluationResultsResponse) SetStatusCode(v int32) *ListEvaluationResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEvaluationResultsResponse) SetBody(v *ListEvaluationResultsResponseBody) *ListEvaluationResultsResponse {
	s.Body = v
	return s
}

func (s *ListEvaluationResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
