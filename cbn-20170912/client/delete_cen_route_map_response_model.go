// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCenRouteMapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCenRouteMapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCenRouteMapResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCenRouteMapResponseBody) *DeleteCenRouteMapResponse
	GetBody() *DeleteCenRouteMapResponseBody
}

type DeleteCenRouteMapResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCenRouteMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCenRouteMapResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCenRouteMapResponse) GoString() string {
	return s.String()
}

func (s *DeleteCenRouteMapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCenRouteMapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCenRouteMapResponse) GetBody() *DeleteCenRouteMapResponseBody {
	return s.Body
}

func (s *DeleteCenRouteMapResponse) SetHeaders(v map[string]*string) *DeleteCenRouteMapResponse {
	s.Headers = v
	return s
}

func (s *DeleteCenRouteMapResponse) SetStatusCode(v int32) *DeleteCenRouteMapResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCenRouteMapResponse) SetBody(v *DeleteCenRouteMapResponseBody) *DeleteCenRouteMapResponse {
	s.Body = v
	return s
}

func (s *DeleteCenRouteMapResponse) Validate() error {
	return dara.Validate(s)
}
