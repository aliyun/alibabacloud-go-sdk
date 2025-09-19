// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterceptionRulePageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInterceptionRulePageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInterceptionRulePageResponse
	GetStatusCode() *int32
	SetBody(v *ListInterceptionRulePageResponseBody) *ListInterceptionRulePageResponse
	GetBody() *ListInterceptionRulePageResponseBody
}

type ListInterceptionRulePageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInterceptionRulePageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInterceptionRulePageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInterceptionRulePageResponse) GoString() string {
	return s.String()
}

func (s *ListInterceptionRulePageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInterceptionRulePageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInterceptionRulePageResponse) GetBody() *ListInterceptionRulePageResponseBody {
	return s.Body
}

func (s *ListInterceptionRulePageResponse) SetHeaders(v map[string]*string) *ListInterceptionRulePageResponse {
	s.Headers = v
	return s
}

func (s *ListInterceptionRulePageResponse) SetStatusCode(v int32) *ListInterceptionRulePageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInterceptionRulePageResponse) SetBody(v *ListInterceptionRulePageResponseBody) *ListInterceptionRulePageResponse {
	s.Body = v
	return s
}

func (s *ListInterceptionRulePageResponse) Validate() error {
	return dara.Validate(s)
}
