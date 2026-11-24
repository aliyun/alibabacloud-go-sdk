// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRouteMapAssociationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRouteMapAssociationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRouteMapAssociationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRouteMapAssociationResponseBody) *UpdateRouteMapAssociationResponse
	GetBody() *UpdateRouteMapAssociationResponseBody
}

type UpdateRouteMapAssociationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRouteMapAssociationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRouteMapAssociationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRouteMapAssociationResponse) GoString() string {
	return s.String()
}

func (s *UpdateRouteMapAssociationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRouteMapAssociationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRouteMapAssociationResponse) GetBody() *UpdateRouteMapAssociationResponseBody {
	return s.Body
}

func (s *UpdateRouteMapAssociationResponse) SetHeaders(v map[string]*string) *UpdateRouteMapAssociationResponse {
	s.Headers = v
	return s
}

func (s *UpdateRouteMapAssociationResponse) SetStatusCode(v int32) *UpdateRouteMapAssociationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRouteMapAssociationResponse) SetBody(v *UpdateRouteMapAssociationResponseBody) *UpdateRouteMapAssociationResponse {
	s.Body = v
	return s
}

func (s *UpdateRouteMapAssociationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
