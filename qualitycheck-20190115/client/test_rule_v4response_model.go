// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestRuleV4Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TestRuleV4Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TestRuleV4Response
	GetStatusCode() *int32
	SetBody(v *TestRuleV4ResponseBody) *TestRuleV4Response
	GetBody() *TestRuleV4ResponseBody
}

type TestRuleV4Response struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TestRuleV4ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TestRuleV4Response) String() string {
	return dara.Prettify(s)
}

func (s TestRuleV4Response) GoString() string {
	return s.String()
}

func (s *TestRuleV4Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TestRuleV4Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TestRuleV4Response) GetBody() *TestRuleV4ResponseBody {
	return s.Body
}

func (s *TestRuleV4Response) SetHeaders(v map[string]*string) *TestRuleV4Response {
	s.Headers = v
	return s
}

func (s *TestRuleV4Response) SetStatusCode(v int32) *TestRuleV4Response {
	s.StatusCode = &v
	return s
}

func (s *TestRuleV4Response) SetBody(v *TestRuleV4ResponseBody) *TestRuleV4Response {
	s.Body = v
	return s
}

func (s *TestRuleV4Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
