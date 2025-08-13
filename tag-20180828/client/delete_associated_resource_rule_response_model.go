// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAssociatedResourceRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAssociatedResourceRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAssociatedResourceRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAssociatedResourceRuleResponseBody) *DeleteAssociatedResourceRuleResponse
	GetBody() *DeleteAssociatedResourceRuleResponseBody
}

type DeleteAssociatedResourceRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAssociatedResourceRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAssociatedResourceRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAssociatedResourceRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteAssociatedResourceRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAssociatedResourceRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAssociatedResourceRuleResponse) GetBody() *DeleteAssociatedResourceRuleResponseBody {
	return s.Body
}

func (s *DeleteAssociatedResourceRuleResponse) SetHeaders(v map[string]*string) *DeleteAssociatedResourceRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteAssociatedResourceRuleResponse) SetStatusCode(v int32) *DeleteAssociatedResourceRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAssociatedResourceRuleResponse) SetBody(v *DeleteAssociatedResourceRuleResponseBody) *DeleteAssociatedResourceRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteAssociatedResourceRuleResponse) Validate() error {
	return dara.Validate(s)
}
