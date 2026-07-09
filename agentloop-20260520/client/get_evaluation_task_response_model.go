// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEvaluationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEvaluationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEvaluationTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetEvaluationTaskResponseBody) *GetEvaluationTaskResponse
	GetBody() *GetEvaluationTaskResponseBody
}

type GetEvaluationTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEvaluationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEvaluationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEvaluationTaskResponse) GoString() string {
	return s.String()
}

func (s *GetEvaluationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEvaluationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEvaluationTaskResponse) GetBody() *GetEvaluationTaskResponseBody {
	return s.Body
}

func (s *GetEvaluationTaskResponse) SetHeaders(v map[string]*string) *GetEvaluationTaskResponse {
	s.Headers = v
	return s
}

func (s *GetEvaluationTaskResponse) SetStatusCode(v int32) *GetEvaluationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEvaluationTaskResponse) SetBody(v *GetEvaluationTaskResponseBody) *GetEvaluationTaskResponse {
	s.Body = v
	return s
}

func (s *GetEvaluationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
