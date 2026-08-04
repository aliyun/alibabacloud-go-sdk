// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterGetDeptBalanceSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterGetDeptBalanceSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterGetDeptBalanceSummaryResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterGetDeptBalanceSummaryResponseBody) *ModelRouterGetDeptBalanceSummaryResponse
	GetBody() *ModelRouterGetDeptBalanceSummaryResponseBody
}

type ModelRouterGetDeptBalanceSummaryResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterGetDeptBalanceSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterGetDeptBalanceSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetDeptBalanceSummaryResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterGetDeptBalanceSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterGetDeptBalanceSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterGetDeptBalanceSummaryResponse) GetBody() *ModelRouterGetDeptBalanceSummaryResponseBody {
	return s.Body
}

func (s *ModelRouterGetDeptBalanceSummaryResponse) SetHeaders(v map[string]*string) *ModelRouterGetDeptBalanceSummaryResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterGetDeptBalanceSummaryResponse) SetStatusCode(v int32) *ModelRouterGetDeptBalanceSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterGetDeptBalanceSummaryResponse) SetBody(v *ModelRouterGetDeptBalanceSummaryResponseBody) *ModelRouterGetDeptBalanceSummaryResponse {
	s.Body = v
	return s
}

func (s *ModelRouterGetDeptBalanceSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
