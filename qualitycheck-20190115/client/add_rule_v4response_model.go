// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRuleV4Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddRuleV4Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddRuleV4Response
	GetStatusCode() *int32
	SetBody(v *AddRuleV4ResponseBody) *AddRuleV4Response
	GetBody() *AddRuleV4ResponseBody
}

type AddRuleV4Response struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRuleV4ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRuleV4Response) String() string {
	return dara.Prettify(s)
}

func (s AddRuleV4Response) GoString() string {
	return s.String()
}

func (s *AddRuleV4Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddRuleV4Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddRuleV4Response) GetBody() *AddRuleV4ResponseBody {
	return s.Body
}

func (s *AddRuleV4Response) SetHeaders(v map[string]*string) *AddRuleV4Response {
	s.Headers = v
	return s
}

func (s *AddRuleV4Response) SetStatusCode(v int32) *AddRuleV4Response {
	s.StatusCode = &v
	return s
}

func (s *AddRuleV4Response) SetBody(v *AddRuleV4ResponseBody) *AddRuleV4Response {
	s.Body = v
	return s
}

func (s *AddRuleV4Response) Validate() error {
	return dara.Validate(s)
}
