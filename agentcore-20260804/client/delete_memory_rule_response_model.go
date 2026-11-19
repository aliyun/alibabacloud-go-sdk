// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemoryRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMemoryRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMemoryRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMemoryRuleResponseBody) *DeleteMemoryRuleResponse
	GetBody() *DeleteMemoryRuleResponseBody
}

type DeleteMemoryRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMemoryRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMemoryRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoryRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteMemoryRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMemoryRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMemoryRuleResponse) GetBody() *DeleteMemoryRuleResponseBody {
	return s.Body
}

func (s *DeleteMemoryRuleResponse) SetHeaders(v map[string]*string) *DeleteMemoryRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteMemoryRuleResponse) SetStatusCode(v int32) *DeleteMemoryRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMemoryRuleResponse) SetBody(v *DeleteMemoryRuleResponseBody) *DeleteMemoryRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteMemoryRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
