// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataQualityEvaluationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataQualityEvaluationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataQualityEvaluationTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataQualityEvaluationTaskResponseBody) *UpdateDataQualityEvaluationTaskResponse
	GetBody() *UpdateDataQualityEvaluationTaskResponseBody
}

type UpdateDataQualityEvaluationTaskResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataQualityEvaluationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataQualityEvaluationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityEvaluationTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityEvaluationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataQualityEvaluationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataQualityEvaluationTaskResponse) GetBody() *UpdateDataQualityEvaluationTaskResponseBody {
	return s.Body
}

func (s *UpdateDataQualityEvaluationTaskResponse) SetHeaders(v map[string]*string) *UpdateDataQualityEvaluationTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataQualityEvaluationTaskResponse) SetStatusCode(v int32) *UpdateDataQualityEvaluationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskResponse) SetBody(v *UpdateDataQualityEvaluationTaskResponseBody) *UpdateDataQualityEvaluationTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateDataQualityEvaluationTaskResponse) Validate() error {
	return dara.Validate(s)
}
