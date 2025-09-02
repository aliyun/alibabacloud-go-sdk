// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQualityRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateQualityRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateQualityRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateQualityRuleResponseBody) *CreateQualityRuleResponse
	GetBody() *CreateQualityRuleResponseBody
}

type CreateQualityRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQualityRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQualityRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateQualityRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateQualityRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateQualityRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateQualityRuleResponse) GetBody() *CreateQualityRuleResponseBody {
	return s.Body
}

func (s *CreateQualityRuleResponse) SetHeaders(v map[string]*string) *CreateQualityRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateQualityRuleResponse) SetStatusCode(v int32) *CreateQualityRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQualityRuleResponse) SetBody(v *CreateQualityRuleResponseBody) *CreateQualityRuleResponse {
	s.Body = v
	return s
}

func (s *CreateQualityRuleResponse) Validate() error {
	return dara.Validate(s)
}
