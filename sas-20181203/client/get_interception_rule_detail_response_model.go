// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInterceptionRuleDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInterceptionRuleDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInterceptionRuleDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetInterceptionRuleDetailResponseBody) *GetInterceptionRuleDetailResponse
	GetBody() *GetInterceptionRuleDetailResponseBody
}

type GetInterceptionRuleDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInterceptionRuleDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInterceptionRuleDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInterceptionRuleDetailResponse) GoString() string {
	return s.String()
}

func (s *GetInterceptionRuleDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInterceptionRuleDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInterceptionRuleDetailResponse) GetBody() *GetInterceptionRuleDetailResponseBody {
	return s.Body
}

func (s *GetInterceptionRuleDetailResponse) SetHeaders(v map[string]*string) *GetInterceptionRuleDetailResponse {
	s.Headers = v
	return s
}

func (s *GetInterceptionRuleDetailResponse) SetStatusCode(v int32) *GetInterceptionRuleDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInterceptionRuleDetailResponse) SetBody(v *GetInterceptionRuleDetailResponseBody) *GetInterceptionRuleDetailResponse {
	s.Body = v
	return s
}

func (s *GetInterceptionRuleDetailResponse) Validate() error {
	return dara.Validate(s)
}
