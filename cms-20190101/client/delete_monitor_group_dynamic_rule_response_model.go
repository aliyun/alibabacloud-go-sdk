// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitorGroupDynamicRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMonitorGroupDynamicRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMonitorGroupDynamicRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMonitorGroupDynamicRuleResponseBody) *DeleteMonitorGroupDynamicRuleResponse
	GetBody() *DeleteMonitorGroupDynamicRuleResponseBody
}

type DeleteMonitorGroupDynamicRuleResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMonitorGroupDynamicRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMonitorGroupDynamicRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorGroupDynamicRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupDynamicRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMonitorGroupDynamicRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMonitorGroupDynamicRuleResponse) GetBody() *DeleteMonitorGroupDynamicRuleResponseBody {
	return s.Body
}

func (s *DeleteMonitorGroupDynamicRuleResponse) SetHeaders(v map[string]*string) *DeleteMonitorGroupDynamicRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteMonitorGroupDynamicRuleResponse) SetStatusCode(v int32) *DeleteMonitorGroupDynamicRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMonitorGroupDynamicRuleResponse) SetBody(v *DeleteMonitorGroupDynamicRuleResponseBody) *DeleteMonitorGroupDynamicRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteMonitorGroupDynamicRuleResponse) Validate() error {
	return dara.Validate(s)
}
