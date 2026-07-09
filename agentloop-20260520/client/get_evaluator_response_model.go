// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEvaluatorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEvaluatorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEvaluatorResponse
	GetStatusCode() *int32
	SetBody(v *GetEvaluatorResponseBody) *GetEvaluatorResponse
	GetBody() *GetEvaluatorResponseBody
}

type GetEvaluatorResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEvaluatorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEvaluatorResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEvaluatorResponse) GoString() string {
	return s.String()
}

func (s *GetEvaluatorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEvaluatorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEvaluatorResponse) GetBody() *GetEvaluatorResponseBody {
	return s.Body
}

func (s *GetEvaluatorResponse) SetHeaders(v map[string]*string) *GetEvaluatorResponse {
	s.Headers = v
	return s
}

func (s *GetEvaluatorResponse) SetStatusCode(v int32) *GetEvaluatorResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEvaluatorResponse) SetBody(v *GetEvaluatorResponseBody) *GetEvaluatorResponse {
	s.Body = v
	return s
}

func (s *GetEvaluatorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
