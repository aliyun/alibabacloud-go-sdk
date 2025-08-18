// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWafRulesetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWafRulesetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWafRulesetResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWafRulesetResponseBody) *UpdateWafRulesetResponse
	GetBody() *UpdateWafRulesetResponseBody
}

type UpdateWafRulesetResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWafRulesetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWafRulesetResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWafRulesetResponse) GoString() string {
	return s.String()
}

func (s *UpdateWafRulesetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWafRulesetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWafRulesetResponse) GetBody() *UpdateWafRulesetResponseBody {
	return s.Body
}

func (s *UpdateWafRulesetResponse) SetHeaders(v map[string]*string) *UpdateWafRulesetResponse {
	s.Headers = v
	return s
}

func (s *UpdateWafRulesetResponse) SetStatusCode(v int32) *UpdateWafRulesetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWafRulesetResponse) SetBody(v *UpdateWafRulesetResponseBody) *UpdateWafRulesetResponse {
	s.Body = v
	return s
}

func (s *UpdateWafRulesetResponse) Validate() error {
	return dara.Validate(s)
}
