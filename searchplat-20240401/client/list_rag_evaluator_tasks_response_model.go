// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRagEvaluatorTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRagEvaluatorTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRagEvaluatorTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListRagEvaluatorTasksResponseBody) *ListRagEvaluatorTasksResponse
	GetBody() *ListRagEvaluatorTasksResponseBody
}

type ListRagEvaluatorTasksResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRagEvaluatorTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRagEvaluatorTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRagEvaluatorTasksResponse) GoString() string {
	return s.String()
}

func (s *ListRagEvaluatorTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRagEvaluatorTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRagEvaluatorTasksResponse) GetBody() *ListRagEvaluatorTasksResponseBody {
	return s.Body
}

func (s *ListRagEvaluatorTasksResponse) SetHeaders(v map[string]*string) *ListRagEvaluatorTasksResponse {
	s.Headers = v
	return s
}

func (s *ListRagEvaluatorTasksResponse) SetStatusCode(v int32) *ListRagEvaluatorTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRagEvaluatorTasksResponse) SetBody(v *ListRagEvaluatorTasksResponseBody) *ListRagEvaluatorTasksResponse {
	s.Body = v
	return s
}

func (s *ListRagEvaluatorTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
