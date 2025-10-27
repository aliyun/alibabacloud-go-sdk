// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyContainerDefenseRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyContainerDefenseRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyContainerDefenseRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyContainerDefenseRuleResponseBody) *ModifyContainerDefenseRuleResponse
	GetBody() *ModifyContainerDefenseRuleResponseBody
}

type ModifyContainerDefenseRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyContainerDefenseRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyContainerDefenseRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyContainerDefenseRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyContainerDefenseRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyContainerDefenseRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyContainerDefenseRuleResponse) GetBody() *ModifyContainerDefenseRuleResponseBody {
	return s.Body
}

func (s *ModifyContainerDefenseRuleResponse) SetHeaders(v map[string]*string) *ModifyContainerDefenseRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyContainerDefenseRuleResponse) SetStatusCode(v int32) *ModifyContainerDefenseRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyContainerDefenseRuleResponse) SetBody(v *ModifyContainerDefenseRuleResponseBody) *ModifyContainerDefenseRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyContainerDefenseRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
