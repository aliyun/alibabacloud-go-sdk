// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityEvaluationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataQualityEvaluationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataQualityEvaluationTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetDataQualityEvaluationTaskResponseBody) *GetDataQualityEvaluationTaskResponse
	GetBody() *GetDataQualityEvaluationTaskResponseBody
}

type GetDataQualityEvaluationTaskResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataQualityEvaluationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataQualityEvaluationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskResponse) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataQualityEvaluationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataQualityEvaluationTaskResponse) GetBody() *GetDataQualityEvaluationTaskResponseBody {
	return s.Body
}

func (s *GetDataQualityEvaluationTaskResponse) SetHeaders(v map[string]*string) *GetDataQualityEvaluationTaskResponse {
	s.Headers = v
	return s
}

func (s *GetDataQualityEvaluationTaskResponse) SetStatusCode(v int32) *GetDataQualityEvaluationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataQualityEvaluationTaskResponse) SetBody(v *GetDataQualityEvaluationTaskResponseBody) *GetDataQualityEvaluationTaskResponse {
	s.Body = v
	return s
}

func (s *GetDataQualityEvaluationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
