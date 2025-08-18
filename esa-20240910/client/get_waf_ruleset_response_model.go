// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWafRulesetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWafRulesetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWafRulesetResponse
	GetStatusCode() *int32
	SetBody(v *GetWafRulesetResponseBody) *GetWafRulesetResponse
	GetBody() *GetWafRulesetResponseBody
}

type GetWafRulesetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWafRulesetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWafRulesetResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWafRulesetResponse) GoString() string {
	return s.String()
}

func (s *GetWafRulesetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWafRulesetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWafRulesetResponse) GetBody() *GetWafRulesetResponseBody {
	return s.Body
}

func (s *GetWafRulesetResponse) SetHeaders(v map[string]*string) *GetWafRulesetResponse {
	s.Headers = v
	return s
}

func (s *GetWafRulesetResponse) SetStatusCode(v int32) *GetWafRulesetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWafRulesetResponse) SetBody(v *GetWafRulesetResponseBody) *GetWafRulesetResponse {
	s.Body = v
	return s
}

func (s *GetWafRulesetResponse) Validate() error {
	return dara.Validate(s)
}
