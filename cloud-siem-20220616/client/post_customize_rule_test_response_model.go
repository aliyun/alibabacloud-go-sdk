// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostCustomizeRuleTestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PostCustomizeRuleTestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PostCustomizeRuleTestResponse
	GetStatusCode() *int32
	SetBody(v *PostCustomizeRuleTestResponseBody) *PostCustomizeRuleTestResponse
	GetBody() *PostCustomizeRuleTestResponseBody
}

type PostCustomizeRuleTestResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostCustomizeRuleTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostCustomizeRuleTestResponse) String() string {
	return dara.Prettify(s)
}

func (s PostCustomizeRuleTestResponse) GoString() string {
	return s.String()
}

func (s *PostCustomizeRuleTestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PostCustomizeRuleTestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PostCustomizeRuleTestResponse) GetBody() *PostCustomizeRuleTestResponseBody {
	return s.Body
}

func (s *PostCustomizeRuleTestResponse) SetHeaders(v map[string]*string) *PostCustomizeRuleTestResponse {
	s.Headers = v
	return s
}

func (s *PostCustomizeRuleTestResponse) SetStatusCode(v int32) *PostCustomizeRuleTestResponse {
	s.StatusCode = &v
	return s
}

func (s *PostCustomizeRuleTestResponse) SetBody(v *PostCustomizeRuleTestResponseBody) *PostCustomizeRuleTestResponse {
	s.Body = v
	return s
}

func (s *PostCustomizeRuleTestResponse) Validate() error {
	return dara.Validate(s)
}
