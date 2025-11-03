// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociatePhysicalConnectionToVirtualBorderRouterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociatePhysicalConnectionToVirtualBorderRouterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociatePhysicalConnectionToVirtualBorderRouterResponse
	GetStatusCode() *int32
	SetBody(v *AssociatePhysicalConnectionToVirtualBorderRouterResponseBody) *AssociatePhysicalConnectionToVirtualBorderRouterResponse
	GetBody() *AssociatePhysicalConnectionToVirtualBorderRouterResponseBody
}

type AssociatePhysicalConnectionToVirtualBorderRouterResponse struct {
	Headers    map[string]*string                                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociatePhysicalConnectionToVirtualBorderRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociatePhysicalConnectionToVirtualBorderRouterResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociatePhysicalConnectionToVirtualBorderRouterResponse) GoString() string {
	return s.String()
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterResponse) GetBody() *AssociatePhysicalConnectionToVirtualBorderRouterResponseBody {
	return s.Body
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterResponse) SetHeaders(v map[string]*string) *AssociatePhysicalConnectionToVirtualBorderRouterResponse {
	s.Headers = v
	return s
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterResponse) SetStatusCode(v int32) *AssociatePhysicalConnectionToVirtualBorderRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterResponse) SetBody(v *AssociatePhysicalConnectionToVirtualBorderRouterResponseBody) *AssociatePhysicalConnectionToVirtualBorderRouterResponse {
	s.Body = v
	return s
}

func (s *AssociatePhysicalConnectionToVirtualBorderRouterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
