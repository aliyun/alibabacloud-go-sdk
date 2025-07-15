// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInterveneRuleDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInterveneRuleDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInterveneRuleDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetInterveneRuleDetailResponseBody) *GetInterveneRuleDetailResponse
	GetBody() *GetInterveneRuleDetailResponseBody
}

type GetInterveneRuleDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInterveneRuleDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInterveneRuleDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInterveneRuleDetailResponse) GoString() string {
	return s.String()
}

func (s *GetInterveneRuleDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInterveneRuleDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInterveneRuleDetailResponse) GetBody() *GetInterveneRuleDetailResponseBody {
	return s.Body
}

func (s *GetInterveneRuleDetailResponse) SetHeaders(v map[string]*string) *GetInterveneRuleDetailResponse {
	s.Headers = v
	return s
}

func (s *GetInterveneRuleDetailResponse) SetStatusCode(v int32) *GetInterveneRuleDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInterveneRuleDetailResponse) SetBody(v *GetInterveneRuleDetailResponseBody) *GetInterveneRuleDetailResponse {
	s.Body = v
	return s
}

func (s *GetInterveneRuleDetailResponse) Validate() error {
	return dara.Validate(s)
}
