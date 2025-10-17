// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoLiveStreamRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAutoLiveStreamRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAutoLiveStreamRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAutoLiveStreamRuleResponseBody) *DeleteAutoLiveStreamRuleResponse
	GetBody() *DeleteAutoLiveStreamRuleResponseBody
}

type DeleteAutoLiveStreamRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAutoLiveStreamRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAutoLiveStreamRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoLiveStreamRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteAutoLiveStreamRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAutoLiveStreamRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAutoLiveStreamRuleResponse) GetBody() *DeleteAutoLiveStreamRuleResponseBody {
	return s.Body
}

func (s *DeleteAutoLiveStreamRuleResponse) SetHeaders(v map[string]*string) *DeleteAutoLiveStreamRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteAutoLiveStreamRuleResponse) SetStatusCode(v int32) *DeleteAutoLiveStreamRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAutoLiveStreamRuleResponse) SetBody(v *DeleteAutoLiveStreamRuleResponseBody) *DeleteAutoLiveStreamRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteAutoLiveStreamRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
