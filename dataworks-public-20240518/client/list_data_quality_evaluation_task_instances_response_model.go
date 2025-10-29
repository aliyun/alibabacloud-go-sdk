// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityEvaluationTaskInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataQualityEvaluationTaskInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataQualityEvaluationTaskInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListDataQualityEvaluationTaskInstancesResponseBody) *ListDataQualityEvaluationTaskInstancesResponse
	GetBody() *ListDataQualityEvaluationTaskInstancesResponseBody
}

type ListDataQualityEvaluationTaskInstancesResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataQualityEvaluationTaskInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataQualityEvaluationTaskInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityEvaluationTaskInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTaskInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataQualityEvaluationTaskInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataQualityEvaluationTaskInstancesResponse) GetBody() *ListDataQualityEvaluationTaskInstancesResponseBody {
	return s.Body
}

func (s *ListDataQualityEvaluationTaskInstancesResponse) SetHeaders(v map[string]*string) *ListDataQualityEvaluationTaskInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponse) SetStatusCode(v int32) *ListDataQualityEvaluationTaskInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponse) SetBody(v *ListDataQualityEvaluationTaskInstancesResponseBody) *ListDataQualityEvaluationTaskInstancesResponse {
	s.Body = v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
