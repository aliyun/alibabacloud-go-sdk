// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInterceptionRuleSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInterceptionRuleSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInterceptionRuleSwitchResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInterceptionRuleSwitchResponseBody) *ModifyInterceptionRuleSwitchResponse
	GetBody() *ModifyInterceptionRuleSwitchResponseBody
}

type ModifyInterceptionRuleSwitchResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInterceptionRuleSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInterceptionRuleSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInterceptionRuleSwitchResponse) GoString() string {
	return s.String()
}

func (s *ModifyInterceptionRuleSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInterceptionRuleSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInterceptionRuleSwitchResponse) GetBody() *ModifyInterceptionRuleSwitchResponseBody {
	return s.Body
}

func (s *ModifyInterceptionRuleSwitchResponse) SetHeaders(v map[string]*string) *ModifyInterceptionRuleSwitchResponse {
	s.Headers = v
	return s
}

func (s *ModifyInterceptionRuleSwitchResponse) SetStatusCode(v int32) *ModifyInterceptionRuleSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInterceptionRuleSwitchResponse) SetBody(v *ModifyInterceptionRuleSwitchResponseBody) *ModifyInterceptionRuleSwitchResponse {
	s.Body = v
	return s
}

func (s *ModifyInterceptionRuleSwitchResponse) Validate() error {
	return dara.Validate(s)
}
