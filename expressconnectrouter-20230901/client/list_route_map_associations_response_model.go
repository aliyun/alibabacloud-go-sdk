// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRouteMapAssociationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRouteMapAssociationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRouteMapAssociationsResponse
	GetStatusCode() *int32
	SetBody(v *ListRouteMapAssociationsResponseBody) *ListRouteMapAssociationsResponse
	GetBody() *ListRouteMapAssociationsResponseBody
}

type ListRouteMapAssociationsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRouteMapAssociationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRouteMapAssociationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRouteMapAssociationsResponse) GoString() string {
	return s.String()
}

func (s *ListRouteMapAssociationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRouteMapAssociationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRouteMapAssociationsResponse) GetBody() *ListRouteMapAssociationsResponseBody {
	return s.Body
}

func (s *ListRouteMapAssociationsResponse) SetHeaders(v map[string]*string) *ListRouteMapAssociationsResponse {
	s.Headers = v
	return s
}

func (s *ListRouteMapAssociationsResponse) SetStatusCode(v int32) *ListRouteMapAssociationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRouteMapAssociationsResponse) SetBody(v *ListRouteMapAssociationsResponseBody) *ListRouteMapAssociationsResponse {
	s.Body = v
	return s
}

func (s *ListRouteMapAssociationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
