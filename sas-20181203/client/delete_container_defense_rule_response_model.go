// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContainerDefenseRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteContainerDefenseRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteContainerDefenseRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteContainerDefenseRuleResponseBody) *DeleteContainerDefenseRuleResponse
	GetBody() *DeleteContainerDefenseRuleResponseBody
}

type DeleteContainerDefenseRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteContainerDefenseRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteContainerDefenseRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteContainerDefenseRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteContainerDefenseRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteContainerDefenseRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteContainerDefenseRuleResponse) GetBody() *DeleteContainerDefenseRuleResponseBody {
	return s.Body
}

func (s *DeleteContainerDefenseRuleResponse) SetHeaders(v map[string]*string) *DeleteContainerDefenseRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteContainerDefenseRuleResponse) SetStatusCode(v int32) *DeleteContainerDefenseRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContainerDefenseRuleResponse) SetBody(v *DeleteContainerDefenseRuleResponseBody) *DeleteContainerDefenseRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteContainerDefenseRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
