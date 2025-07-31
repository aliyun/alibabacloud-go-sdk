// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckRecoveryConditionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckRecoveryConditionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckRecoveryConditionResponse
	GetStatusCode() *int32
	SetBody(v *CheckRecoveryConditionResponseBody) *CheckRecoveryConditionResponse
	GetBody() *CheckRecoveryConditionResponseBody
}

type CheckRecoveryConditionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckRecoveryConditionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckRecoveryConditionResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckRecoveryConditionResponse) GoString() string {
	return s.String()
}

func (s *CheckRecoveryConditionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckRecoveryConditionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckRecoveryConditionResponse) GetBody() *CheckRecoveryConditionResponseBody {
	return s.Body
}

func (s *CheckRecoveryConditionResponse) SetHeaders(v map[string]*string) *CheckRecoveryConditionResponse {
	s.Headers = v
	return s
}

func (s *CheckRecoveryConditionResponse) SetStatusCode(v int32) *CheckRecoveryConditionResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckRecoveryConditionResponse) SetBody(v *CheckRecoveryConditionResponseBody) *CheckRecoveryConditionResponse {
	s.Body = v
	return s
}

func (s *CheckRecoveryConditionResponse) Validate() error {
	return dara.Validate(s)
}
