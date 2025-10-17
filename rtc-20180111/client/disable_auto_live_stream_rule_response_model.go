// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAutoLiveStreamRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableAutoLiveStreamRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableAutoLiveStreamRuleResponse
	GetStatusCode() *int32
	SetBody(v *DisableAutoLiveStreamRuleResponseBody) *DisableAutoLiveStreamRuleResponse
	GetBody() *DisableAutoLiveStreamRuleResponseBody
}

type DisableAutoLiveStreamRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableAutoLiveStreamRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableAutoLiveStreamRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableAutoLiveStreamRuleResponse) GoString() string {
	return s.String()
}

func (s *DisableAutoLiveStreamRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableAutoLiveStreamRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableAutoLiveStreamRuleResponse) GetBody() *DisableAutoLiveStreamRuleResponseBody {
	return s.Body
}

func (s *DisableAutoLiveStreamRuleResponse) SetHeaders(v map[string]*string) *DisableAutoLiveStreamRuleResponse {
	s.Headers = v
	return s
}

func (s *DisableAutoLiveStreamRuleResponse) SetStatusCode(v int32) *DisableAutoLiveStreamRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableAutoLiveStreamRuleResponse) SetBody(v *DisableAutoLiveStreamRuleResponseBody) *DisableAutoLiveStreamRuleResponse {
	s.Body = v
	return s
}

func (s *DisableAutoLiveStreamRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
