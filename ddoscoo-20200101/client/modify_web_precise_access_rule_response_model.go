// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebPreciseAccessRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWebPreciseAccessRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWebPreciseAccessRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWebPreciseAccessRuleResponseBody) *ModifyWebPreciseAccessRuleResponse
	GetBody() *ModifyWebPreciseAccessRuleResponseBody
}

type ModifyWebPreciseAccessRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWebPreciseAccessRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWebPreciseAccessRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebPreciseAccessRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebPreciseAccessRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWebPreciseAccessRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWebPreciseAccessRuleResponse) GetBody() *ModifyWebPreciseAccessRuleResponseBody {
	return s.Body
}

func (s *ModifyWebPreciseAccessRuleResponse) SetHeaders(v map[string]*string) *ModifyWebPreciseAccessRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebPreciseAccessRuleResponse) SetStatusCode(v int32) *ModifyWebPreciseAccessRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWebPreciseAccessRuleResponse) SetBody(v *ModifyWebPreciseAccessRuleResponseBody) *ModifyWebPreciseAccessRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyWebPreciseAccessRuleResponse) Validate() error {
	return dara.Validate(s)
}
