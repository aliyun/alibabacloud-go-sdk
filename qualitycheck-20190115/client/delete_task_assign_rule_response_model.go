// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTaskAssignRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTaskAssignRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTaskAssignRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTaskAssignRuleResponseBody) *DeleteTaskAssignRuleResponse
	GetBody() *DeleteTaskAssignRuleResponseBody
}

type DeleteTaskAssignRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTaskAssignRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTaskAssignRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTaskAssignRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteTaskAssignRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTaskAssignRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTaskAssignRuleResponse) GetBody() *DeleteTaskAssignRuleResponseBody {
	return s.Body
}

func (s *DeleteTaskAssignRuleResponse) SetHeaders(v map[string]*string) *DeleteTaskAssignRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteTaskAssignRuleResponse) SetStatusCode(v int32) *DeleteTaskAssignRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTaskAssignRuleResponse) SetBody(v *DeleteTaskAssignRuleResponseBody) *DeleteTaskAssignRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteTaskAssignRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
