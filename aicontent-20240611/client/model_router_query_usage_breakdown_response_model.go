// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryUsageBreakdownResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterQueryUsageBreakdownResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterQueryUsageBreakdownResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterQueryUsageBreakdownResponseBody) *ModelRouterQueryUsageBreakdownResponse
	GetBody() *ModelRouterQueryUsageBreakdownResponseBody
}

type ModelRouterQueryUsageBreakdownResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterQueryUsageBreakdownResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterQueryUsageBreakdownResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryUsageBreakdownResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryUsageBreakdownResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterQueryUsageBreakdownResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterQueryUsageBreakdownResponse) GetBody() *ModelRouterQueryUsageBreakdownResponseBody {
	return s.Body
}

func (s *ModelRouterQueryUsageBreakdownResponse) SetHeaders(v map[string]*string) *ModelRouterQueryUsageBreakdownResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterQueryUsageBreakdownResponse) SetStatusCode(v int32) *ModelRouterQueryUsageBreakdownResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterQueryUsageBreakdownResponse) SetBody(v *ModelRouterQueryUsageBreakdownResponseBody) *ModelRouterQueryUsageBreakdownResponse {
	s.Body = v
	return s
}

func (s *ModelRouterQueryUsageBreakdownResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
