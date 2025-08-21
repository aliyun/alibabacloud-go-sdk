// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWebRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWebRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWebRuleResponseBody) *ModifyWebRuleResponse
	GetBody() *ModifyWebRuleResponseBody
}

type ModifyWebRuleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWebRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWebRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWebRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWebRuleResponse) GetBody() *ModifyWebRuleResponseBody {
	return s.Body
}

func (s *ModifyWebRuleResponse) SetHeaders(v map[string]*string) *ModifyWebRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebRuleResponse) SetStatusCode(v int32) *ModifyWebRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWebRuleResponse) SetBody(v *ModifyWebRuleResponseBody) *ModifyWebRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyWebRuleResponse) Validate() error {
	return dara.Validate(s)
}
