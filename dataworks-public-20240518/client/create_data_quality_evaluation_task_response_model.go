// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityEvaluationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataQualityEvaluationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataQualityEvaluationTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataQualityEvaluationTaskResponseBody) *CreateDataQualityEvaluationTaskResponse
	GetBody() *CreateDataQualityEvaluationTaskResponseBody
}

type CreateDataQualityEvaluationTaskResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataQualityEvaluationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataQualityEvaluationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityEvaluationTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDataQualityEvaluationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataQualityEvaluationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataQualityEvaluationTaskResponse) GetBody() *CreateDataQualityEvaluationTaskResponseBody {
	return s.Body
}

func (s *CreateDataQualityEvaluationTaskResponse) SetHeaders(v map[string]*string) *CreateDataQualityEvaluationTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDataQualityEvaluationTaskResponse) SetStatusCode(v int32) *CreateDataQualityEvaluationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskResponse) SetBody(v *CreateDataQualityEvaluationTaskResponseBody) *CreateDataQualityEvaluationTaskResponse {
	s.Body = v
	return s
}

func (s *CreateDataQualityEvaluationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
