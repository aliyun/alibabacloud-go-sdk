// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoLiveStreamRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAutoLiveStreamRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAutoLiveStreamRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateAutoLiveStreamRuleResponseBody) *CreateAutoLiveStreamRuleResponse
	GetBody() *CreateAutoLiveStreamRuleResponseBody
}

type CreateAutoLiveStreamRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAutoLiveStreamRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAutoLiveStreamRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoLiveStreamRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateAutoLiveStreamRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAutoLiveStreamRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAutoLiveStreamRuleResponse) GetBody() *CreateAutoLiveStreamRuleResponseBody {
	return s.Body
}

func (s *CreateAutoLiveStreamRuleResponse) SetHeaders(v map[string]*string) *CreateAutoLiveStreamRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateAutoLiveStreamRuleResponse) SetStatusCode(v int32) *CreateAutoLiveStreamRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAutoLiveStreamRuleResponse) SetBody(v *CreateAutoLiveStreamRuleResponseBody) *CreateAutoLiveStreamRuleResponse {
	s.Body = v
	return s
}

func (s *CreateAutoLiveStreamRuleResponse) Validate() error {
	return dara.Validate(s)
}
