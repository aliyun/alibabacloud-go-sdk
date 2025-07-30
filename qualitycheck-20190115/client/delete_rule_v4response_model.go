// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRuleV4Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRuleV4Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRuleV4Response
	GetStatusCode() *int32
	SetBody(v *DeleteRuleV4ResponseBody) *DeleteRuleV4Response
	GetBody() *DeleteRuleV4ResponseBody
}

type DeleteRuleV4Response struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRuleV4ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRuleV4Response) String() string {
	return dara.Prettify(s)
}

func (s DeleteRuleV4Response) GoString() string {
	return s.String()
}

func (s *DeleteRuleV4Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRuleV4Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRuleV4Response) GetBody() *DeleteRuleV4ResponseBody {
	return s.Body
}

func (s *DeleteRuleV4Response) SetHeaders(v map[string]*string) *DeleteRuleV4Response {
	s.Headers = v
	return s
}

func (s *DeleteRuleV4Response) SetStatusCode(v int32) *DeleteRuleV4Response {
	s.StatusCode = &v
	return s
}

func (s *DeleteRuleV4Response) SetBody(v *DeleteRuleV4ResponseBody) *DeleteRuleV4Response {
	s.Body = v
	return s
}

func (s *DeleteRuleV4Response) Validate() error {
	return dara.Validate(s)
}
