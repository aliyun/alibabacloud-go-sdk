// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeRuleListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQueryOptimizeRuleListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQueryOptimizeRuleListResponse
	GetStatusCode() *int32
	SetBody(v *GetQueryOptimizeRuleListResponseBody) *GetQueryOptimizeRuleListResponse
	GetBody() *GetQueryOptimizeRuleListResponseBody
}

type GetQueryOptimizeRuleListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQueryOptimizeRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQueryOptimizeRuleListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeRuleListResponse) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeRuleListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQueryOptimizeRuleListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQueryOptimizeRuleListResponse) GetBody() *GetQueryOptimizeRuleListResponseBody {
	return s.Body
}

func (s *GetQueryOptimizeRuleListResponse) SetHeaders(v map[string]*string) *GetQueryOptimizeRuleListResponse {
	s.Headers = v
	return s
}

func (s *GetQueryOptimizeRuleListResponse) SetStatusCode(v int32) *GetQueryOptimizeRuleListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQueryOptimizeRuleListResponse) SetBody(v *GetQueryOptimizeRuleListResponseBody) *GetQueryOptimizeRuleListResponse {
	s.Body = v
	return s
}

func (s *GetQueryOptimizeRuleListResponse) Validate() error {
	return dara.Validate(s)
}
