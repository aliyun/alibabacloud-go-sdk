// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiduiAreaDeductionProResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DiduiAreaDeductionProResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DiduiAreaDeductionProResponse
	GetStatusCode() *int32
	SetBody(v *DiduiAreaDeductionProResponseBody) *DiduiAreaDeductionProResponse
	GetBody() *DiduiAreaDeductionProResponseBody
}

type DiduiAreaDeductionProResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DiduiAreaDeductionProResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DiduiAreaDeductionProResponse) String() string {
	return dara.Prettify(s)
}

func (s DiduiAreaDeductionProResponse) GoString() string {
	return s.String()
}

func (s *DiduiAreaDeductionProResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DiduiAreaDeductionProResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DiduiAreaDeductionProResponse) GetBody() *DiduiAreaDeductionProResponseBody {
	return s.Body
}

func (s *DiduiAreaDeductionProResponse) SetHeaders(v map[string]*string) *DiduiAreaDeductionProResponse {
	s.Headers = v
	return s
}

func (s *DiduiAreaDeductionProResponse) SetStatusCode(v int32) *DiduiAreaDeductionProResponse {
	s.StatusCode = &v
	return s
}

func (s *DiduiAreaDeductionProResponse) SetBody(v *DiduiAreaDeductionProResponseBody) *DiduiAreaDeductionProResponse {
	s.Body = v
	return s
}

func (s *DiduiAreaDeductionProResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
