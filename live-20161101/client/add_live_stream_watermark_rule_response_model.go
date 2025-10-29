// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveStreamWatermarkRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLiveStreamWatermarkRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLiveStreamWatermarkRuleResponse
	GetStatusCode() *int32
	SetBody(v *AddLiveStreamWatermarkRuleResponseBody) *AddLiveStreamWatermarkRuleResponse
	GetBody() *AddLiveStreamWatermarkRuleResponseBody
}

type AddLiveStreamWatermarkRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLiveStreamWatermarkRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLiveStreamWatermarkRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLiveStreamWatermarkRuleResponse) GoString() string {
	return s.String()
}

func (s *AddLiveStreamWatermarkRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLiveStreamWatermarkRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLiveStreamWatermarkRuleResponse) GetBody() *AddLiveStreamWatermarkRuleResponseBody {
	return s.Body
}

func (s *AddLiveStreamWatermarkRuleResponse) SetHeaders(v map[string]*string) *AddLiveStreamWatermarkRuleResponse {
	s.Headers = v
	return s
}

func (s *AddLiveStreamWatermarkRuleResponse) SetStatusCode(v int32) *AddLiveStreamWatermarkRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLiveStreamWatermarkRuleResponse) SetBody(v *AddLiveStreamWatermarkRuleResponseBody) *AddLiveStreamWatermarkRuleResponse {
	s.Body = v
	return s
}

func (s *AddLiveStreamWatermarkRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
