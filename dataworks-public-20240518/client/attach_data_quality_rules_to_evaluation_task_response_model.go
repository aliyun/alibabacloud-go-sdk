// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachDataQualityRulesToEvaluationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachDataQualityRulesToEvaluationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachDataQualityRulesToEvaluationTaskResponse
	GetStatusCode() *int32
	SetBody(v *AttachDataQualityRulesToEvaluationTaskResponseBody) *AttachDataQualityRulesToEvaluationTaskResponse
	GetBody() *AttachDataQualityRulesToEvaluationTaskResponseBody
}

type AttachDataQualityRulesToEvaluationTaskResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachDataQualityRulesToEvaluationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachDataQualityRulesToEvaluationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachDataQualityRulesToEvaluationTaskResponse) GoString() string {
	return s.String()
}

func (s *AttachDataQualityRulesToEvaluationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachDataQualityRulesToEvaluationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachDataQualityRulesToEvaluationTaskResponse) GetBody() *AttachDataQualityRulesToEvaluationTaskResponseBody {
	return s.Body
}

func (s *AttachDataQualityRulesToEvaluationTaskResponse) SetHeaders(v map[string]*string) *AttachDataQualityRulesToEvaluationTaskResponse {
	s.Headers = v
	return s
}

func (s *AttachDataQualityRulesToEvaluationTaskResponse) SetStatusCode(v int32) *AttachDataQualityRulesToEvaluationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachDataQualityRulesToEvaluationTaskResponse) SetBody(v *AttachDataQualityRulesToEvaluationTaskResponseBody) *AttachDataQualityRulesToEvaluationTaskResponse {
	s.Body = v
	return s
}

func (s *AttachDataQualityRulesToEvaluationTaskResponse) Validate() error {
	return dara.Validate(s)
}
