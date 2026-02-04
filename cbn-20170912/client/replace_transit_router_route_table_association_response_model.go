// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceTransitRouterRouteTableAssociationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReplaceTransitRouterRouteTableAssociationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReplaceTransitRouterRouteTableAssociationResponse
	GetStatusCode() *int32
	SetBody(v *ReplaceTransitRouterRouteTableAssociationResponseBody) *ReplaceTransitRouterRouteTableAssociationResponse
	GetBody() *ReplaceTransitRouterRouteTableAssociationResponseBody
}

type ReplaceTransitRouterRouteTableAssociationResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReplaceTransitRouterRouteTableAssociationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReplaceTransitRouterRouteTableAssociationResponse) String() string {
	return dara.Prettify(s)
}

func (s ReplaceTransitRouterRouteTableAssociationResponse) GoString() string {
	return s.String()
}

func (s *ReplaceTransitRouterRouteTableAssociationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReplaceTransitRouterRouteTableAssociationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReplaceTransitRouterRouteTableAssociationResponse) GetBody() *ReplaceTransitRouterRouteTableAssociationResponseBody {
	return s.Body
}

func (s *ReplaceTransitRouterRouteTableAssociationResponse) SetHeaders(v map[string]*string) *ReplaceTransitRouterRouteTableAssociationResponse {
	s.Headers = v
	return s
}

func (s *ReplaceTransitRouterRouteTableAssociationResponse) SetStatusCode(v int32) *ReplaceTransitRouterRouteTableAssociationResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplaceTransitRouterRouteTableAssociationResponse) SetBody(v *ReplaceTransitRouterRouteTableAssociationResponseBody) *ReplaceTransitRouterRouteTableAssociationResponse {
	s.Body = v
	return s
}

func (s *ReplaceTransitRouterRouteTableAssociationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
