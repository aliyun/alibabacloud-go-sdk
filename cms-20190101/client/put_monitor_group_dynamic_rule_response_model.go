// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutMonitorGroupDynamicRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutMonitorGroupDynamicRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutMonitorGroupDynamicRuleResponse
	GetStatusCode() *int32
	SetBody(v *PutMonitorGroupDynamicRuleResponseBody) *PutMonitorGroupDynamicRuleResponse
	GetBody() *PutMonitorGroupDynamicRuleResponseBody
}

type PutMonitorGroupDynamicRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutMonitorGroupDynamicRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutMonitorGroupDynamicRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s PutMonitorGroupDynamicRuleResponse) GoString() string {
	return s.String()
}

func (s *PutMonitorGroupDynamicRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutMonitorGroupDynamicRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutMonitorGroupDynamicRuleResponse) GetBody() *PutMonitorGroupDynamicRuleResponseBody {
	return s.Body
}

func (s *PutMonitorGroupDynamicRuleResponse) SetHeaders(v map[string]*string) *PutMonitorGroupDynamicRuleResponse {
	s.Headers = v
	return s
}

func (s *PutMonitorGroupDynamicRuleResponse) SetStatusCode(v int32) *PutMonitorGroupDynamicRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *PutMonitorGroupDynamicRuleResponse) SetBody(v *PutMonitorGroupDynamicRuleResponseBody) *PutMonitorGroupDynamicRuleResponse {
	s.Body = v
	return s
}

func (s *PutMonitorGroupDynamicRuleResponse) Validate() error {
	return dara.Validate(s)
}
