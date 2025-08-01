// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchLosslessRuleListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FetchLosslessRuleListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FetchLosslessRuleListResponse
	GetStatusCode() *int32
	SetBody(v *FetchLosslessRuleListResponseBody) *FetchLosslessRuleListResponse
	GetBody() *FetchLosslessRuleListResponseBody
}

type FetchLosslessRuleListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FetchLosslessRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FetchLosslessRuleListResponse) String() string {
	return dara.Prettify(s)
}

func (s FetchLosslessRuleListResponse) GoString() string {
	return s.String()
}

func (s *FetchLosslessRuleListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FetchLosslessRuleListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FetchLosslessRuleListResponse) GetBody() *FetchLosslessRuleListResponseBody {
	return s.Body
}

func (s *FetchLosslessRuleListResponse) SetHeaders(v map[string]*string) *FetchLosslessRuleListResponse {
	s.Headers = v
	return s
}

func (s *FetchLosslessRuleListResponse) SetStatusCode(v int32) *FetchLosslessRuleListResponse {
	s.StatusCode = &v
	return s
}

func (s *FetchLosslessRuleListResponse) SetBody(v *FetchLosslessRuleListResponseBody) *FetchLosslessRuleListResponse {
	s.Body = v
	return s
}

func (s *FetchLosslessRuleListResponse) Validate() error {
	return dara.Validate(s)
}
