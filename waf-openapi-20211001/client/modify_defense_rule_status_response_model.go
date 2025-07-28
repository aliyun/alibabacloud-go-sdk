// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseRuleStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDefenseRuleStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDefenseRuleStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDefenseRuleStatusResponseBody) *ModifyDefenseRuleStatusResponse
	GetBody() *ModifyDefenseRuleStatusResponseBody
}

type ModifyDefenseRuleStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDefenseRuleStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDefenseRuleStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseRuleStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefenseRuleStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDefenseRuleStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDefenseRuleStatusResponse) GetBody() *ModifyDefenseRuleStatusResponseBody {
	return s.Body
}

func (s *ModifyDefenseRuleStatusResponse) SetHeaders(v map[string]*string) *ModifyDefenseRuleStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefenseRuleStatusResponse) SetStatusCode(v int32) *ModifyDefenseRuleStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefenseRuleStatusResponse) SetBody(v *ModifyDefenseRuleStatusResponseBody) *ModifyDefenseRuleStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyDefenseRuleStatusResponse) Validate() error {
	return dara.Validate(s)
}
