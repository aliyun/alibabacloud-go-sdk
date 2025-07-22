// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExpressConnectRouterSupportedRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExpressConnectRouterSupportedRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExpressConnectRouterSupportedRegionResponse
	GetStatusCode() *int32
	SetBody(v *ListExpressConnectRouterSupportedRegionResponseBody) *ListExpressConnectRouterSupportedRegionResponse
	GetBody() *ListExpressConnectRouterSupportedRegionResponseBody
}

type ListExpressConnectRouterSupportedRegionResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExpressConnectRouterSupportedRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExpressConnectRouterSupportedRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExpressConnectRouterSupportedRegionResponse) GoString() string {
	return s.String()
}

func (s *ListExpressConnectRouterSupportedRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExpressConnectRouterSupportedRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExpressConnectRouterSupportedRegionResponse) GetBody() *ListExpressConnectRouterSupportedRegionResponseBody {
	return s.Body
}

func (s *ListExpressConnectRouterSupportedRegionResponse) SetHeaders(v map[string]*string) *ListExpressConnectRouterSupportedRegionResponse {
	s.Headers = v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionResponse) SetStatusCode(v int32) *ListExpressConnectRouterSupportedRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionResponse) SetBody(v *ListExpressConnectRouterSupportedRegionResponseBody) *ListExpressConnectRouterSupportedRegionResponse {
	s.Body = v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionResponse) Validate() error {
	return dara.Validate(s)
}
