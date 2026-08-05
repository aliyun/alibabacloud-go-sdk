// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRagEvaluatorTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRagEvaluatorTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRagEvaluatorTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetRagEvaluatorTaskResponseBody) *GetRagEvaluatorTaskResponse
	GetBody() *GetRagEvaluatorTaskResponseBody
}

type GetRagEvaluatorTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRagEvaluatorTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRagEvaluatorTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRagEvaluatorTaskResponse) GoString() string {
	return s.String()
}

func (s *GetRagEvaluatorTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRagEvaluatorTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRagEvaluatorTaskResponse) GetBody() *GetRagEvaluatorTaskResponseBody {
	return s.Body
}

func (s *GetRagEvaluatorTaskResponse) SetHeaders(v map[string]*string) *GetRagEvaluatorTaskResponse {
	s.Headers = v
	return s
}

func (s *GetRagEvaluatorTaskResponse) SetStatusCode(v int32) *GetRagEvaluatorTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRagEvaluatorTaskResponse) SetBody(v *GetRagEvaluatorTaskResponseBody) *GetRagEvaluatorTaskResponse {
	s.Body = v
	return s
}

func (s *GetRagEvaluatorTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
