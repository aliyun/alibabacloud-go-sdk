// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuleDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRuleDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRuleDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetRuleDetailResponseBody) *GetRuleDetailResponse
	GetBody() *GetRuleDetailResponseBody
}

type GetRuleDetailResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRuleDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRuleDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRuleDetailResponse) GoString() string {
	return s.String()
}

func (s *GetRuleDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRuleDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRuleDetailResponse) GetBody() *GetRuleDetailResponseBody {
	return s.Body
}

func (s *GetRuleDetailResponse) SetHeaders(v map[string]*string) *GetRuleDetailResponse {
	s.Headers = v
	return s
}

func (s *GetRuleDetailResponse) SetStatusCode(v int32) *GetRuleDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRuleDetailResponse) SetBody(v *GetRuleDetailResponseBody) *GetRuleDetailResponse {
	s.Body = v
	return s
}

func (s *GetRuleDetailResponse) Validate() error {
	return dara.Validate(s)
}
