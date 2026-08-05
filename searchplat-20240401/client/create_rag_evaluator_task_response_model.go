// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRagEvaluatorTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRagEvaluatorTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRagEvaluatorTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateRagEvaluatorTaskResponseBody) *CreateRagEvaluatorTaskResponse
	GetBody() *CreateRagEvaluatorTaskResponseBody
}

type CreateRagEvaluatorTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRagEvaluatorTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRagEvaluatorTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRagEvaluatorTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateRagEvaluatorTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRagEvaluatorTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRagEvaluatorTaskResponse) GetBody() *CreateRagEvaluatorTaskResponseBody {
	return s.Body
}

func (s *CreateRagEvaluatorTaskResponse) SetHeaders(v map[string]*string) *CreateRagEvaluatorTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateRagEvaluatorTaskResponse) SetStatusCode(v int32) *CreateRagEvaluatorTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRagEvaluatorTaskResponse) SetBody(v *CreateRagEvaluatorTaskResponseBody) *CreateRagEvaluatorTaskResponse {
	s.Body = v
	return s
}

func (s *CreateRagEvaluatorTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
