// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvaluationTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEvaluationTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEvaluationTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListEvaluationTasksResponseBody) *ListEvaluationTasksResponse
	GetBody() *ListEvaluationTasksResponseBody
}

type ListEvaluationTasksResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEvaluationTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEvaluationTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationTasksResponse) GoString() string {
	return s.String()
}

func (s *ListEvaluationTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEvaluationTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEvaluationTasksResponse) GetBody() *ListEvaluationTasksResponseBody {
	return s.Body
}

func (s *ListEvaluationTasksResponse) SetHeaders(v map[string]*string) *ListEvaluationTasksResponse {
	s.Headers = v
	return s
}

func (s *ListEvaluationTasksResponse) SetStatusCode(v int32) *ListEvaluationTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEvaluationTasksResponse) SetBody(v *ListEvaluationTasksResponseBody) *ListEvaluationTasksResponse {
	s.Body = v
	return s
}

func (s *ListEvaluationTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
