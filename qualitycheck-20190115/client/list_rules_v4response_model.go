// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRulesV4Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRulesV4Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRulesV4Response
	GetStatusCode() *int32
	SetBody(v *ListRulesV4ResponseBody) *ListRulesV4Response
	GetBody() *ListRulesV4ResponseBody
}

type ListRulesV4Response struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRulesV4ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRulesV4Response) String() string {
	return dara.Prettify(s)
}

func (s ListRulesV4Response) GoString() string {
	return s.String()
}

func (s *ListRulesV4Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRulesV4Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRulesV4Response) GetBody() *ListRulesV4ResponseBody {
	return s.Body
}

func (s *ListRulesV4Response) SetHeaders(v map[string]*string) *ListRulesV4Response {
	s.Headers = v
	return s
}

func (s *ListRulesV4Response) SetStatusCode(v int32) *ListRulesV4Response {
	s.StatusCode = &v
	return s
}

func (s *ListRulesV4Response) SetBody(v *ListRulesV4ResponseBody) *ListRulesV4Response {
	s.Body = v
	return s
}

func (s *ListRulesV4Response) Validate() error {
	return dara.Validate(s)
}
