// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteMapAssociationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRouteMapAssociationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRouteMapAssociationResponse
	GetStatusCode() *int32
	SetBody(v *CreateRouteMapAssociationResponseBody) *CreateRouteMapAssociationResponse
	GetBody() *CreateRouteMapAssociationResponseBody
}

type CreateRouteMapAssociationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRouteMapAssociationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRouteMapAssociationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteMapAssociationResponse) GoString() string {
	return s.String()
}

func (s *CreateRouteMapAssociationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRouteMapAssociationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRouteMapAssociationResponse) GetBody() *CreateRouteMapAssociationResponseBody {
	return s.Body
}

func (s *CreateRouteMapAssociationResponse) SetHeaders(v map[string]*string) *CreateRouteMapAssociationResponse {
	s.Headers = v
	return s
}

func (s *CreateRouteMapAssociationResponse) SetStatusCode(v int32) *CreateRouteMapAssociationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRouteMapAssociationResponse) SetBody(v *CreateRouteMapAssociationResponseBody) *CreateRouteMapAssociationResponse {
	s.Body = v
	return s
}

func (s *CreateRouteMapAssociationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
