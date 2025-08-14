// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRuleStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRuleStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRuleStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRuleStatusResponseBody) *ModifyRuleStatusResponse
	GetBody() *ModifyRuleStatusResponseBody
}

type ModifyRuleStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRuleStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRuleStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRuleStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyRuleStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRuleStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRuleStatusResponse) GetBody() *ModifyRuleStatusResponseBody {
	return s.Body
}

func (s *ModifyRuleStatusResponse) SetHeaders(v map[string]*string) *ModifyRuleStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyRuleStatusResponse) SetStatusCode(v int32) *ModifyRuleStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRuleStatusResponse) SetBody(v *ModifyRuleStatusResponseBody) *ModifyRuleStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyRuleStatusResponse) Validate() error {
	return dara.Validate(s)
}
