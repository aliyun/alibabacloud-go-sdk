// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuleV4Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRuleV4Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRuleV4Response
	GetStatusCode() *int32
	SetBody(v *GetRuleV4ResponseBody) *GetRuleV4Response
	GetBody() *GetRuleV4ResponseBody
}

type GetRuleV4Response struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRuleV4ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRuleV4Response) String() string {
	return dara.Prettify(s)
}

func (s GetRuleV4Response) GoString() string {
	return s.String()
}

func (s *GetRuleV4Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRuleV4Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRuleV4Response) GetBody() *GetRuleV4ResponseBody {
	return s.Body
}

func (s *GetRuleV4Response) SetHeaders(v map[string]*string) *GetRuleV4Response {
	s.Headers = v
	return s
}

func (s *GetRuleV4Response) SetStatusCode(v int32) *GetRuleV4Response {
	s.StatusCode = &v
	return s
}

func (s *GetRuleV4Response) SetBody(v *GetRuleV4ResponseBody) *GetRuleV4Response {
	s.Body = v
	return s
}

func (s *GetRuleV4Response) Validate() error {
	return dara.Validate(s)
}
