// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityEvaluationTaskInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataQualityEvaluationTaskInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataQualityEvaluationTaskInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataQualityEvaluationTaskInstanceResponseBody) *CreateDataQualityEvaluationTaskInstanceResponse
	GetBody() *CreateDataQualityEvaluationTaskInstanceResponseBody
}

type CreateDataQualityEvaluationTaskInstanceResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataQualityEvaluationTaskInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataQualityEvaluationTaskInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityEvaluationTaskInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateDataQualityEvaluationTaskInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataQualityEvaluationTaskInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataQualityEvaluationTaskInstanceResponse) GetBody() *CreateDataQualityEvaluationTaskInstanceResponseBody {
	return s.Body
}

func (s *CreateDataQualityEvaluationTaskInstanceResponse) SetHeaders(v map[string]*string) *CreateDataQualityEvaluationTaskInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateDataQualityEvaluationTaskInstanceResponse) SetStatusCode(v int32) *CreateDataQualityEvaluationTaskInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskInstanceResponse) SetBody(v *CreateDataQualityEvaluationTaskInstanceResponseBody) *CreateDataQualityEvaluationTaskInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateDataQualityEvaluationTaskInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
