// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEvaluationRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEvaluationRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEvaluationRunResponse
	GetStatusCode() *int32
	SetBody(v *GetEvaluationRunResponseBody) *GetEvaluationRunResponse
	GetBody() *GetEvaluationRunResponseBody
}

type GetEvaluationRunResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEvaluationRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEvaluationRunResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEvaluationRunResponse) GoString() string {
	return s.String()
}

func (s *GetEvaluationRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEvaluationRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEvaluationRunResponse) GetBody() *GetEvaluationRunResponseBody {
	return s.Body
}

func (s *GetEvaluationRunResponse) SetHeaders(v map[string]*string) *GetEvaluationRunResponse {
	s.Headers = v
	return s
}

func (s *GetEvaluationRunResponse) SetStatusCode(v int32) *GetEvaluationRunResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEvaluationRunResponse) SetBody(v *GetEvaluationRunResponseBody) *GetEvaluationRunResponse {
	s.Body = v
	return s
}

func (s *GetEvaluationRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
