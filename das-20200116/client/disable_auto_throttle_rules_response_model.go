// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAutoThrottleRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableAutoThrottleRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableAutoThrottleRulesResponse
	GetStatusCode() *int32
	SetBody(v *DisableAutoThrottleRulesResponseBody) *DisableAutoThrottleRulesResponse
	GetBody() *DisableAutoThrottleRulesResponseBody
}

type DisableAutoThrottleRulesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableAutoThrottleRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableAutoThrottleRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableAutoThrottleRulesResponse) GoString() string {
	return s.String()
}

func (s *DisableAutoThrottleRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableAutoThrottleRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableAutoThrottleRulesResponse) GetBody() *DisableAutoThrottleRulesResponseBody {
	return s.Body
}

func (s *DisableAutoThrottleRulesResponse) SetHeaders(v map[string]*string) *DisableAutoThrottleRulesResponse {
	s.Headers = v
	return s
}

func (s *DisableAutoThrottleRulesResponse) SetStatusCode(v int32) *DisableAutoThrottleRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableAutoThrottleRulesResponse) SetBody(v *DisableAutoThrottleRulesResponseBody) *DisableAutoThrottleRulesResponse {
	s.Body = v
	return s
}

func (s *DisableAutoThrottleRulesResponse) Validate() error {
	return dara.Validate(s)
}
