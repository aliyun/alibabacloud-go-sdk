// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEvaluatorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEvaluatorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEvaluatorResponse
	GetStatusCode() *int32
	SetBody(v *CreateEvaluatorResponseBody) *CreateEvaluatorResponse
	GetBody() *CreateEvaluatorResponseBody
}

type CreateEvaluatorResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEvaluatorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEvaluatorResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEvaluatorResponse) GoString() string {
	return s.String()
}

func (s *CreateEvaluatorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEvaluatorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEvaluatorResponse) GetBody() *CreateEvaluatorResponseBody {
	return s.Body
}

func (s *CreateEvaluatorResponse) SetHeaders(v map[string]*string) *CreateEvaluatorResponse {
	s.Headers = v
	return s
}

func (s *CreateEvaluatorResponse) SetStatusCode(v int32) *CreateEvaluatorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEvaluatorResponse) SetBody(v *CreateEvaluatorResponseBody) *CreateEvaluatorResponse {
	s.Body = v
	return s
}

func (s *CreateEvaluatorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
