// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEvaluationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEvaluationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEvaluationTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEvaluationTaskResponseBody) *UpdateEvaluationTaskResponse
	GetBody() *UpdateEvaluationTaskResponseBody
}

type UpdateEvaluationTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEvaluationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEvaluationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEvaluationTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateEvaluationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEvaluationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEvaluationTaskResponse) GetBody() *UpdateEvaluationTaskResponseBody {
	return s.Body
}

func (s *UpdateEvaluationTaskResponse) SetHeaders(v map[string]*string) *UpdateEvaluationTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateEvaluationTaskResponse) SetStatusCode(v int32) *UpdateEvaluationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEvaluationTaskResponse) SetBody(v *UpdateEvaluationTaskResponseBody) *UpdateEvaluationTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateEvaluationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
