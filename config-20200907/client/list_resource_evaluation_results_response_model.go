// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceEvaluationResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceEvaluationResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceEvaluationResultsResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceEvaluationResultsResponseBody) *ListResourceEvaluationResultsResponse
	GetBody() *ListResourceEvaluationResultsResponseBody
}

type ListResourceEvaluationResultsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceEvaluationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceEvaluationResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceEvaluationResultsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceEvaluationResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceEvaluationResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceEvaluationResultsResponse) GetBody() *ListResourceEvaluationResultsResponseBody {
	return s.Body
}

func (s *ListResourceEvaluationResultsResponse) SetHeaders(v map[string]*string) *ListResourceEvaluationResultsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceEvaluationResultsResponse) SetStatusCode(v int32) *ListResourceEvaluationResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceEvaluationResultsResponse) SetBody(v *ListResourceEvaluationResultsResponseBody) *ListResourceEvaluationResultsResponse {
	s.Body = v
	return s
}

func (s *ListResourceEvaluationResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
