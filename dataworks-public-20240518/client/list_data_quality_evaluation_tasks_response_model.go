// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityEvaluationTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataQualityEvaluationTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataQualityEvaluationTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListDataQualityEvaluationTasksResponseBody) *ListDataQualityEvaluationTasksResponse
	GetBody() *ListDataQualityEvaluationTasksResponseBody
}

type ListDataQualityEvaluationTasksResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataQualityEvaluationTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataQualityEvaluationTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityEvaluationTasksResponse) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataQualityEvaluationTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataQualityEvaluationTasksResponse) GetBody() *ListDataQualityEvaluationTasksResponseBody {
	return s.Body
}

func (s *ListDataQualityEvaluationTasksResponse) SetHeaders(v map[string]*string) *ListDataQualityEvaluationTasksResponse {
	s.Headers = v
	return s
}

func (s *ListDataQualityEvaluationTasksResponse) SetStatusCode(v int32) *ListDataQualityEvaluationTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataQualityEvaluationTasksResponse) SetBody(v *ListDataQualityEvaluationTasksResponseBody) *ListDataQualityEvaluationTasksResponse {
	s.Body = v
	return s
}

func (s *ListDataQualityEvaluationTasksResponse) Validate() error {
	return dara.Validate(s)
}
