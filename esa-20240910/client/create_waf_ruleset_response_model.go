// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWafRulesetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWafRulesetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWafRulesetResponse
	GetStatusCode() *int32
	SetBody(v *CreateWafRulesetResponseBody) *CreateWafRulesetResponse
	GetBody() *CreateWafRulesetResponseBody
}

type CreateWafRulesetResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWafRulesetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWafRulesetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWafRulesetResponse) GoString() string {
	return s.String()
}

func (s *CreateWafRulesetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWafRulesetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWafRulesetResponse) GetBody() *CreateWafRulesetResponseBody {
	return s.Body
}

func (s *CreateWafRulesetResponse) SetHeaders(v map[string]*string) *CreateWafRulesetResponse {
	s.Headers = v
	return s
}

func (s *CreateWafRulesetResponse) SetStatusCode(v int32) *CreateWafRulesetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWafRulesetResponse) SetBody(v *CreateWafRulesetResponseBody) *CreateWafRulesetResponse {
	s.Body = v
	return s
}

func (s *CreateWafRulesetResponse) Validate() error {
	return dara.Validate(s)
}
