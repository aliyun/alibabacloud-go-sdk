// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveStreamWatermarkRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLiveStreamWatermarkRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLiveStreamWatermarkRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLiveStreamWatermarkRuleResponseBody) *UpdateLiveStreamWatermarkRuleResponse
	GetBody() *UpdateLiveStreamWatermarkRuleResponseBody
}

type UpdateLiveStreamWatermarkRuleResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveStreamWatermarkRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveStreamWatermarkRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveStreamWatermarkRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveStreamWatermarkRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLiveStreamWatermarkRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLiveStreamWatermarkRuleResponse) GetBody() *UpdateLiveStreamWatermarkRuleResponseBody {
	return s.Body
}

func (s *UpdateLiveStreamWatermarkRuleResponse) SetHeaders(v map[string]*string) *UpdateLiveStreamWatermarkRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveStreamWatermarkRuleResponse) SetStatusCode(v int32) *UpdateLiveStreamWatermarkRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveStreamWatermarkRuleResponse) SetBody(v *UpdateLiveStreamWatermarkRuleResponseBody) *UpdateLiveStreamWatermarkRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateLiveStreamWatermarkRuleResponse) Validate() error {
	return dara.Validate(s)
}
