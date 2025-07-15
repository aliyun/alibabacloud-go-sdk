// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociatePhysicalConnectionFromVirtualBorderRouterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnassociatePhysicalConnectionFromVirtualBorderRouterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnassociatePhysicalConnectionFromVirtualBorderRouterResponse
	GetStatusCode() *int32
	SetBody(v *UnassociatePhysicalConnectionFromVirtualBorderRouterResponseBody) *UnassociatePhysicalConnectionFromVirtualBorderRouterResponse
	GetBody() *UnassociatePhysicalConnectionFromVirtualBorderRouterResponseBody
}

type UnassociatePhysicalConnectionFromVirtualBorderRouterResponse struct {
	Headers    map[string]*string                                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnassociatePhysicalConnectionFromVirtualBorderRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnassociatePhysicalConnectionFromVirtualBorderRouterResponse) String() string {
	return dara.Prettify(s)
}

func (s UnassociatePhysicalConnectionFromVirtualBorderRouterResponse) GoString() string {
	return s.String()
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterResponse) GetBody() *UnassociatePhysicalConnectionFromVirtualBorderRouterResponseBody {
	return s.Body
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterResponse) SetHeaders(v map[string]*string) *UnassociatePhysicalConnectionFromVirtualBorderRouterResponse {
	s.Headers = v
	return s
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterResponse) SetStatusCode(v int32) *UnassociatePhysicalConnectionFromVirtualBorderRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterResponse) SetBody(v *UnassociatePhysicalConnectionFromVirtualBorderRouterResponseBody) *UnassociatePhysicalConnectionFromVirtualBorderRouterResponse {
	s.Body = v
	return s
}

func (s *UnassociatePhysicalConnectionFromVirtualBorderRouterResponse) Validate() error {
	return dara.Validate(s)
}
