// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoLiveStreamRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAutoLiveStreamRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAutoLiveStreamRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAutoLiveStreamRuleResponseBody) *UpdateAutoLiveStreamRuleResponse
	GetBody() *UpdateAutoLiveStreamRuleResponseBody
}

type UpdateAutoLiveStreamRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAutoLiveStreamRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAutoLiveStreamRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoLiveStreamRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateAutoLiveStreamRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAutoLiveStreamRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAutoLiveStreamRuleResponse) GetBody() *UpdateAutoLiveStreamRuleResponseBody {
	return s.Body
}

func (s *UpdateAutoLiveStreamRuleResponse) SetHeaders(v map[string]*string) *UpdateAutoLiveStreamRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateAutoLiveStreamRuleResponse) SetStatusCode(v int32) *UpdateAutoLiveStreamRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAutoLiveStreamRuleResponse) SetBody(v *UpdateAutoLiveStreamRuleResponseBody) *UpdateAutoLiveStreamRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateAutoLiveStreamRuleResponse) Validate() error {
	return dara.Validate(s)
}
