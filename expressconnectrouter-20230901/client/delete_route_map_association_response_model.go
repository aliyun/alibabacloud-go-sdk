// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteMapAssociationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRouteMapAssociationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRouteMapAssociationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRouteMapAssociationResponseBody) *DeleteRouteMapAssociationResponse
	GetBody() *DeleteRouteMapAssociationResponseBody
}

type DeleteRouteMapAssociationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRouteMapAssociationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRouteMapAssociationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteMapAssociationResponse) GoString() string {
	return s.String()
}

func (s *DeleteRouteMapAssociationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRouteMapAssociationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRouteMapAssociationResponse) GetBody() *DeleteRouteMapAssociationResponseBody {
	return s.Body
}

func (s *DeleteRouteMapAssociationResponse) SetHeaders(v map[string]*string) *DeleteRouteMapAssociationResponse {
	s.Headers = v
	return s
}

func (s *DeleteRouteMapAssociationResponse) SetStatusCode(v int32) *DeleteRouteMapAssociationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRouteMapAssociationResponse) SetBody(v *DeleteRouteMapAssociationResponseBody) *DeleteRouteMapAssociationResponse {
	s.Body = v
	return s
}

func (s *DeleteRouteMapAssociationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
