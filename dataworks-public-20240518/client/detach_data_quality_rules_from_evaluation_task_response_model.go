// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachDataQualityRulesFromEvaluationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachDataQualityRulesFromEvaluationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachDataQualityRulesFromEvaluationTaskResponse
	GetStatusCode() *int32
	SetBody(v *DetachDataQualityRulesFromEvaluationTaskResponseBody) *DetachDataQualityRulesFromEvaluationTaskResponse
	GetBody() *DetachDataQualityRulesFromEvaluationTaskResponseBody
}

type DetachDataQualityRulesFromEvaluationTaskResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachDataQualityRulesFromEvaluationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachDataQualityRulesFromEvaluationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachDataQualityRulesFromEvaluationTaskResponse) GoString() string {
	return s.String()
}

func (s *DetachDataQualityRulesFromEvaluationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachDataQualityRulesFromEvaluationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachDataQualityRulesFromEvaluationTaskResponse) GetBody() *DetachDataQualityRulesFromEvaluationTaskResponseBody {
	return s.Body
}

func (s *DetachDataQualityRulesFromEvaluationTaskResponse) SetHeaders(v map[string]*string) *DetachDataQualityRulesFromEvaluationTaskResponse {
	s.Headers = v
	return s
}

func (s *DetachDataQualityRulesFromEvaluationTaskResponse) SetStatusCode(v int32) *DetachDataQualityRulesFromEvaluationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachDataQualityRulesFromEvaluationTaskResponse) SetBody(v *DetachDataQualityRulesFromEvaluationTaskResponseBody) *DetachDataQualityRulesFromEvaluationTaskResponse {
	s.Body = v
	return s
}

func (s *DetachDataQualityRulesFromEvaluationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
