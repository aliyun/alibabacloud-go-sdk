// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAddRegionToExpressConnectRouterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckAddRegionToExpressConnectRouterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckAddRegionToExpressConnectRouterResponse
	GetStatusCode() *int32
	SetBody(v *CheckAddRegionToExpressConnectRouterResponseBody) *CheckAddRegionToExpressConnectRouterResponse
	GetBody() *CheckAddRegionToExpressConnectRouterResponseBody
}

type CheckAddRegionToExpressConnectRouterResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAddRegionToExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAddRegionToExpressConnectRouterResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckAddRegionToExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *CheckAddRegionToExpressConnectRouterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckAddRegionToExpressConnectRouterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckAddRegionToExpressConnectRouterResponse) GetBody() *CheckAddRegionToExpressConnectRouterResponseBody {
	return s.Body
}

func (s *CheckAddRegionToExpressConnectRouterResponse) SetHeaders(v map[string]*string) *CheckAddRegionToExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponse) SetStatusCode(v int32) *CheckAddRegionToExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponse) SetBody(v *CheckAddRegionToExpressConnectRouterResponseBody) *CheckAddRegionToExpressConnectRouterResponse {
	s.Body = v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponse) Validate() error {
	return dara.Validate(s)
}
