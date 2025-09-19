// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyContainerPluginRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyContainerPluginRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyContainerPluginRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyContainerPluginRuleResponseBody) *ModifyContainerPluginRuleResponse
	GetBody() *ModifyContainerPluginRuleResponseBody
}

type ModifyContainerPluginRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyContainerPluginRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyContainerPluginRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyContainerPluginRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyContainerPluginRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyContainerPluginRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyContainerPluginRuleResponse) GetBody() *ModifyContainerPluginRuleResponseBody {
	return s.Body
}

func (s *ModifyContainerPluginRuleResponse) SetHeaders(v map[string]*string) *ModifyContainerPluginRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyContainerPluginRuleResponse) SetStatusCode(v int32) *ModifyContainerPluginRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyContainerPluginRuleResponse) SetBody(v *ModifyContainerPluginRuleResponseBody) *ModifyContainerPluginRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyContainerPluginRuleResponse) Validate() error {
	return dara.Validate(s)
}
