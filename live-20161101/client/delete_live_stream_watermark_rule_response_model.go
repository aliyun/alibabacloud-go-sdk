// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveStreamWatermarkRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveStreamWatermarkRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveStreamWatermarkRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveStreamWatermarkRuleResponseBody) *DeleteLiveStreamWatermarkRuleResponse
	GetBody() *DeleteLiveStreamWatermarkRuleResponseBody
}

type DeleteLiveStreamWatermarkRuleResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveStreamWatermarkRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveStreamWatermarkRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamWatermarkRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamWatermarkRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveStreamWatermarkRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveStreamWatermarkRuleResponse) GetBody() *DeleteLiveStreamWatermarkRuleResponseBody {
	return s.Body
}

func (s *DeleteLiveStreamWatermarkRuleResponse) SetHeaders(v map[string]*string) *DeleteLiveStreamWatermarkRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveStreamWatermarkRuleResponse) SetStatusCode(v int32) *DeleteLiveStreamWatermarkRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveStreamWatermarkRuleResponse) SetBody(v *DeleteLiveStreamWatermarkRuleResponseBody) *DeleteLiveStreamWatermarkRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveStreamWatermarkRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
