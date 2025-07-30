// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRuleV4Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRuleV4Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRuleV4Response
	GetStatusCode() *int32
	SetBody(v *UpdateRuleV4ResponseBody) *UpdateRuleV4Response
	GetBody() *UpdateRuleV4ResponseBody
}

type UpdateRuleV4Response struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRuleV4ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRuleV4Response) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleV4Response) GoString() string {
	return s.String()
}

func (s *UpdateRuleV4Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRuleV4Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRuleV4Response) GetBody() *UpdateRuleV4ResponseBody {
	return s.Body
}

func (s *UpdateRuleV4Response) SetHeaders(v map[string]*string) *UpdateRuleV4Response {
	s.Headers = v
	return s
}

func (s *UpdateRuleV4Response) SetStatusCode(v int32) *UpdateRuleV4Response {
	s.StatusCode = &v
	return s
}

func (s *UpdateRuleV4Response) SetBody(v *UpdateRuleV4ResponseBody) *UpdateRuleV4Response {
	s.Body = v
	return s
}

func (s *UpdateRuleV4Response) Validate() error {
	return dara.Validate(s)
}
