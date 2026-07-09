// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEvaluationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEvaluationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEvaluationTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateEvaluationTaskResponseBody) *CreateEvaluationTaskResponse
	GetBody() *CreateEvaluationTaskResponseBody
}

type CreateEvaluationTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEvaluationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEvaluationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEvaluationTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateEvaluationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEvaluationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEvaluationTaskResponse) GetBody() *CreateEvaluationTaskResponseBody {
	return s.Body
}

func (s *CreateEvaluationTaskResponse) SetHeaders(v map[string]*string) *CreateEvaluationTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateEvaluationTaskResponse) SetStatusCode(v int32) *CreateEvaluationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEvaluationTaskResponse) SetBody(v *CreateEvaluationTaskResponseBody) *CreateEvaluationTaskResponse {
	s.Body = v
	return s
}

func (s *CreateEvaluationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
