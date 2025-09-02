// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQualityRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateQualityRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateQualityRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateQualityRuleResponseBody) *UpdateQualityRuleResponse
	GetBody() *UpdateQualityRuleResponseBody
}

type UpdateQualityRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateQualityRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateQualityRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateQualityRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateQualityRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateQualityRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateQualityRuleResponse) GetBody() *UpdateQualityRuleResponseBody {
	return s.Body
}

func (s *UpdateQualityRuleResponse) SetHeaders(v map[string]*string) *UpdateQualityRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateQualityRuleResponse) SetStatusCode(v int32) *UpdateQualityRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQualityRuleResponse) SetBody(v *UpdateQualityRuleResponseBody) *UpdateQualityRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateQualityRuleResponse) Validate() error {
	return dara.Validate(s)
}
