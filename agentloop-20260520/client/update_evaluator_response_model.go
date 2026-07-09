// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEvaluatorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEvaluatorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEvaluatorResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEvaluatorResponseBody) *UpdateEvaluatorResponse
	GetBody() *UpdateEvaluatorResponseBody
}

type UpdateEvaluatorResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEvaluatorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEvaluatorResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEvaluatorResponse) GoString() string {
	return s.String()
}

func (s *UpdateEvaluatorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEvaluatorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEvaluatorResponse) GetBody() *UpdateEvaluatorResponseBody {
	return s.Body
}

func (s *UpdateEvaluatorResponse) SetHeaders(v map[string]*string) *UpdateEvaluatorResponse {
	s.Headers = v
	return s
}

func (s *UpdateEvaluatorResponse) SetStatusCode(v int32) *UpdateEvaluatorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEvaluatorResponse) SetBody(v *UpdateEvaluatorResponseBody) *UpdateEvaluatorResponse {
	s.Body = v
	return s
}

func (s *UpdateEvaluatorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
