// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRecognizeRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddRecognizeRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddRecognizeRuleResponse
	GetStatusCode() *int32
	SetBody(v *AddRecognizeRuleResponseBody) *AddRecognizeRuleResponse
	GetBody() *AddRecognizeRuleResponseBody
}

type AddRecognizeRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRecognizeRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRecognizeRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddRecognizeRuleResponse) GoString() string {
	return s.String()
}

func (s *AddRecognizeRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddRecognizeRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddRecognizeRuleResponse) GetBody() *AddRecognizeRuleResponseBody {
	return s.Body
}

func (s *AddRecognizeRuleResponse) SetHeaders(v map[string]*string) *AddRecognizeRuleResponse {
	s.Headers = v
	return s
}

func (s *AddRecognizeRuleResponse) SetStatusCode(v int32) *AddRecognizeRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRecognizeRuleResponse) SetBody(v *AddRecognizeRuleResponseBody) *AddRecognizeRuleResponse {
	s.Body = v
	return s
}

func (s *AddRecognizeRuleResponse) Validate() error {
	return dara.Validate(s)
}
