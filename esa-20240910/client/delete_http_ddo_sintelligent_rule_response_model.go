// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpDDoSIntelligentRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHttpDDoSIntelligentRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHttpDDoSIntelligentRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHttpDDoSIntelligentRuleResponseBody) *DeleteHttpDDoSIntelligentRuleResponse
	GetBody() *DeleteHttpDDoSIntelligentRuleResponseBody
}

type DeleteHttpDDoSIntelligentRuleResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHttpDDoSIntelligentRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHttpDDoSIntelligentRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpDDoSIntelligentRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteHttpDDoSIntelligentRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHttpDDoSIntelligentRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHttpDDoSIntelligentRuleResponse) GetBody() *DeleteHttpDDoSIntelligentRuleResponseBody {
	return s.Body
}

func (s *DeleteHttpDDoSIntelligentRuleResponse) SetHeaders(v map[string]*string) *DeleteHttpDDoSIntelligentRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteHttpDDoSIntelligentRuleResponse) SetStatusCode(v int32) *DeleteHttpDDoSIntelligentRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHttpDDoSIntelligentRuleResponse) SetBody(v *DeleteHttpDDoSIntelligentRuleResponseBody) *DeleteHttpDDoSIntelligentRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteHttpDDoSIntelligentRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
