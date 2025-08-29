// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceRuleResponseBody) *GetResourceRuleResponse
	GetBody() *GetResourceRuleResponseBody
}

type GetResourceRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceRuleResponse) GoString() string {
	return s.String()
}

func (s *GetResourceRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceRuleResponse) GetBody() *GetResourceRuleResponseBody {
	return s.Body
}

func (s *GetResourceRuleResponse) SetHeaders(v map[string]*string) *GetResourceRuleResponse {
	s.Headers = v
	return s
}

func (s *GetResourceRuleResponse) SetStatusCode(v int32) *GetResourceRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceRuleResponse) SetBody(v *GetResourceRuleResponseBody) *GetResourceRuleResponse {
	s.Body = v
	return s
}

func (s *GetResourceRuleResponse) Validate() error {
	return dara.Validate(s)
}
