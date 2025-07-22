// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAutoResourceOptimizeRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableAutoResourceOptimizeRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableAutoResourceOptimizeRulesResponse
	GetStatusCode() *int32
	SetBody(v *DisableAutoResourceOptimizeRulesResponseBody) *DisableAutoResourceOptimizeRulesResponse
	GetBody() *DisableAutoResourceOptimizeRulesResponseBody
}

type DisableAutoResourceOptimizeRulesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableAutoResourceOptimizeRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableAutoResourceOptimizeRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableAutoResourceOptimizeRulesResponse) GoString() string {
	return s.String()
}

func (s *DisableAutoResourceOptimizeRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableAutoResourceOptimizeRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableAutoResourceOptimizeRulesResponse) GetBody() *DisableAutoResourceOptimizeRulesResponseBody {
	return s.Body
}

func (s *DisableAutoResourceOptimizeRulesResponse) SetHeaders(v map[string]*string) *DisableAutoResourceOptimizeRulesResponse {
	s.Headers = v
	return s
}

func (s *DisableAutoResourceOptimizeRulesResponse) SetStatusCode(v int32) *DisableAutoResourceOptimizeRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableAutoResourceOptimizeRulesResponse) SetBody(v *DisableAutoResourceOptimizeRulesResponseBody) *DisableAutoResourceOptimizeRulesResponse {
	s.Body = v
	return s
}

func (s *DisableAutoResourceOptimizeRulesResponse) Validate() error {
	return dara.Validate(s)
}
