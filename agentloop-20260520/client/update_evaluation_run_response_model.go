// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEvaluationRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEvaluationRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEvaluationRunResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEvaluationRunResponseBody) *UpdateEvaluationRunResponse
	GetBody() *UpdateEvaluationRunResponseBody
}

type UpdateEvaluationRunResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEvaluationRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEvaluationRunResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEvaluationRunResponse) GoString() string {
	return s.String()
}

func (s *UpdateEvaluationRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEvaluationRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEvaluationRunResponse) GetBody() *UpdateEvaluationRunResponseBody {
	return s.Body
}

func (s *UpdateEvaluationRunResponse) SetHeaders(v map[string]*string) *UpdateEvaluationRunResponse {
	s.Headers = v
	return s
}

func (s *UpdateEvaluationRunResponse) SetStatusCode(v int32) *UpdateEvaluationRunResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEvaluationRunResponse) SetBody(v *UpdateEvaluationRunResponseBody) *UpdateEvaluationRunResponse {
	s.Body = v
	return s
}

func (s *UpdateEvaluationRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
