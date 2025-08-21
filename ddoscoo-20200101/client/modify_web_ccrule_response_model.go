// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebCCRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWebCCRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWebCCRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWebCCRuleResponseBody) *ModifyWebCCRuleResponse
	GetBody() *ModifyWebCCRuleResponseBody
}

type ModifyWebCCRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWebCCRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWebCCRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebCCRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebCCRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWebCCRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWebCCRuleResponse) GetBody() *ModifyWebCCRuleResponseBody {
	return s.Body
}

func (s *ModifyWebCCRuleResponse) SetHeaders(v map[string]*string) *ModifyWebCCRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebCCRuleResponse) SetStatusCode(v int32) *ModifyWebCCRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWebCCRuleResponse) SetBody(v *ModifyWebCCRuleResponseBody) *ModifyWebCCRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyWebCCRuleResponse) Validate() error {
	return dara.Validate(s)
}
