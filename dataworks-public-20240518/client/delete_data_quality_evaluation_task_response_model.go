// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataQualityEvaluationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataQualityEvaluationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataQualityEvaluationTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataQualityEvaluationTaskResponseBody) *DeleteDataQualityEvaluationTaskResponse
	GetBody() *DeleteDataQualityEvaluationTaskResponseBody
}

type DeleteDataQualityEvaluationTaskResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataQualityEvaluationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataQualityEvaluationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataQualityEvaluationTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataQualityEvaluationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataQualityEvaluationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataQualityEvaluationTaskResponse) GetBody() *DeleteDataQualityEvaluationTaskResponseBody {
	return s.Body
}

func (s *DeleteDataQualityEvaluationTaskResponse) SetHeaders(v map[string]*string) *DeleteDataQualityEvaluationTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataQualityEvaluationTaskResponse) SetStatusCode(v int32) *DeleteDataQualityEvaluationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataQualityEvaluationTaskResponse) SetBody(v *DeleteDataQualityEvaluationTaskResponseBody) *DeleteDataQualityEvaluationTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteDataQualityEvaluationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
