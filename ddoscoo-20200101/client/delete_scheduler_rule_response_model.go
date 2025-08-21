// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSchedulerRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSchedulerRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSchedulerRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSchedulerRuleResponseBody) *DeleteSchedulerRuleResponse
	GetBody() *DeleteSchedulerRuleResponseBody
}

type DeleteSchedulerRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSchedulerRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSchedulerRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSchedulerRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteSchedulerRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSchedulerRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSchedulerRuleResponse) GetBody() *DeleteSchedulerRuleResponseBody {
	return s.Body
}

func (s *DeleteSchedulerRuleResponse) SetHeaders(v map[string]*string) *DeleteSchedulerRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteSchedulerRuleResponse) SetStatusCode(v int32) *DeleteSchedulerRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSchedulerRuleResponse) SetBody(v *DeleteSchedulerRuleResponseBody) *DeleteSchedulerRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteSchedulerRuleResponse) Validate() error {
	return dara.Validate(s)
}
