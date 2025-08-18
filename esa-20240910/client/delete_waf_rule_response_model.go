// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWafRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWafRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWafRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWafRuleResponseBody) *DeleteWafRuleResponse
	GetBody() *DeleteWafRuleResponseBody
}

type DeleteWafRuleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWafRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWafRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWafRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteWafRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWafRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWafRuleResponse) GetBody() *DeleteWafRuleResponseBody {
	return s.Body
}

func (s *DeleteWafRuleResponse) SetHeaders(v map[string]*string) *DeleteWafRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteWafRuleResponse) SetStatusCode(v int32) *DeleteWafRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWafRuleResponse) SetBody(v *DeleteWafRuleResponseBody) *DeleteWafRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteWafRuleResponse) Validate() error {
	return dara.Validate(s)
}
