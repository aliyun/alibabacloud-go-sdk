// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAuthRuleDetailByRuleIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAuthRuleDetailByRuleIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAuthRuleDetailByRuleIdResponse
	GetStatusCode() *int32
	SetBody(v *QueryAuthRuleDetailByRuleIdResponseBody) *QueryAuthRuleDetailByRuleIdResponse
	GetBody() *QueryAuthRuleDetailByRuleIdResponseBody
}

type QueryAuthRuleDetailByRuleIdResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAuthRuleDetailByRuleIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAuthRuleDetailByRuleIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAuthRuleDetailByRuleIdResponse) GoString() string {
	return s.String()
}

func (s *QueryAuthRuleDetailByRuleIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAuthRuleDetailByRuleIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAuthRuleDetailByRuleIdResponse) GetBody() *QueryAuthRuleDetailByRuleIdResponseBody {
	return s.Body
}

func (s *QueryAuthRuleDetailByRuleIdResponse) SetHeaders(v map[string]*string) *QueryAuthRuleDetailByRuleIdResponse {
	s.Headers = v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponse) SetStatusCode(v int32) *QueryAuthRuleDetailByRuleIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponse) SetBody(v *QueryAuthRuleDetailByRuleIdResponseBody) *QueryAuthRuleDetailByRuleIdResponse {
	s.Body = v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdResponse) Validate() error {
	return dara.Validate(s)
}
