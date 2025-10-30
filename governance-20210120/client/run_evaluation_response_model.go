// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunEvaluationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunEvaluationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunEvaluationResponse
	GetStatusCode() *int32
	SetBody(v *RunEvaluationResponseBody) *RunEvaluationResponse
	GetBody() *RunEvaluationResponseBody
}

type RunEvaluationResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunEvaluationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunEvaluationResponse) String() string {
	return dara.Prettify(s)
}

func (s RunEvaluationResponse) GoString() string {
	return s.String()
}

func (s *RunEvaluationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunEvaluationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunEvaluationResponse) GetBody() *RunEvaluationResponseBody {
	return s.Body
}

func (s *RunEvaluationResponse) SetHeaders(v map[string]*string) *RunEvaluationResponse {
	s.Headers = v
	return s
}

func (s *RunEvaluationResponse) SetStatusCode(v int32) *RunEvaluationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunEvaluationResponse) SetBody(v *RunEvaluationResponseBody) *RunEvaluationResponse {
	s.Body = v
	return s
}

func (s *RunEvaluationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
