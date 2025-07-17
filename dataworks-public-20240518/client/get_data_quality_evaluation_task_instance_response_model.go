// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityEvaluationTaskInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataQualityEvaluationTaskInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataQualityEvaluationTaskInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetDataQualityEvaluationTaskInstanceResponseBody) *GetDataQualityEvaluationTaskInstanceResponse
	GetBody() *GetDataQualityEvaluationTaskInstanceResponseBody
}

type GetDataQualityEvaluationTaskInstanceResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataQualityEvaluationTaskInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataQualityEvaluationTaskInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataQualityEvaluationTaskInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataQualityEvaluationTaskInstanceResponse) GetBody() *GetDataQualityEvaluationTaskInstanceResponseBody {
	return s.Body
}

func (s *GetDataQualityEvaluationTaskInstanceResponse) SetHeaders(v map[string]*string) *GetDataQualityEvaluationTaskInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponse) SetStatusCode(v int32) *GetDataQualityEvaluationTaskInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponse) SetBody(v *GetDataQualityEvaluationTaskInstanceResponseBody) *GetDataQualityEvaluationTaskInstanceResponse {
	s.Body = v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceResponse) Validate() error {
	return dara.Validate(s)
}
