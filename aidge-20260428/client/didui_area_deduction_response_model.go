// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiduiAreaDeductionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DiduiAreaDeductionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DiduiAreaDeductionResponse
	GetStatusCode() *int32
	SetBody(v *DiduiAreaDeductionResponseBody) *DiduiAreaDeductionResponse
	GetBody() *DiduiAreaDeductionResponseBody
}

type DiduiAreaDeductionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DiduiAreaDeductionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DiduiAreaDeductionResponse) String() string {
	return dara.Prettify(s)
}

func (s DiduiAreaDeductionResponse) GoString() string {
	return s.String()
}

func (s *DiduiAreaDeductionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DiduiAreaDeductionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DiduiAreaDeductionResponse) GetBody() *DiduiAreaDeductionResponseBody {
	return s.Body
}

func (s *DiduiAreaDeductionResponse) SetHeaders(v map[string]*string) *DiduiAreaDeductionResponse {
	s.Headers = v
	return s
}

func (s *DiduiAreaDeductionResponse) SetStatusCode(v int32) *DiduiAreaDeductionResponse {
	s.StatusCode = &v
	return s
}

func (s *DiduiAreaDeductionResponse) SetBody(v *DiduiAreaDeductionResponseBody) *DiduiAreaDeductionResponse {
	s.Body = v
	return s
}

func (s *DiduiAreaDeductionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
