// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutEvaluationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutEvaluationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutEvaluationsResponse
	GetStatusCode() *int32
	SetBody(v *PutEvaluationsResponseBody) *PutEvaluationsResponse
	GetBody() *PutEvaluationsResponseBody
}

type PutEvaluationsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutEvaluationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutEvaluationsResponse) String() string {
	return dara.Prettify(s)
}

func (s PutEvaluationsResponse) GoString() string {
	return s.String()
}

func (s *PutEvaluationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutEvaluationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutEvaluationsResponse) GetBody() *PutEvaluationsResponseBody {
	return s.Body
}

func (s *PutEvaluationsResponse) SetHeaders(v map[string]*string) *PutEvaluationsResponse {
	s.Headers = v
	return s
}

func (s *PutEvaluationsResponse) SetStatusCode(v int32) *PutEvaluationsResponse {
	s.StatusCode = &v
	return s
}

func (s *PutEvaluationsResponse) SetBody(v *PutEvaluationsResponseBody) *PutEvaluationsResponse {
	s.Body = v
	return s
}

func (s *PutEvaluationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
