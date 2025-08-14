// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCenRouteMapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCenRouteMapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCenRouteMapResponse
	GetStatusCode() *int32
	SetBody(v *CreateCenRouteMapResponseBody) *CreateCenRouteMapResponse
	GetBody() *CreateCenRouteMapResponseBody
}

type CreateCenRouteMapResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCenRouteMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCenRouteMapResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCenRouteMapResponse) GoString() string {
	return s.String()
}

func (s *CreateCenRouteMapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCenRouteMapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCenRouteMapResponse) GetBody() *CreateCenRouteMapResponseBody {
	return s.Body
}

func (s *CreateCenRouteMapResponse) SetHeaders(v map[string]*string) *CreateCenRouteMapResponse {
	s.Headers = v
	return s
}

func (s *CreateCenRouteMapResponse) SetStatusCode(v int32) *CreateCenRouteMapResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCenRouteMapResponse) SetBody(v *CreateCenRouteMapResponseBody) *CreateCenRouteMapResponse {
	s.Body = v
	return s
}

func (s *CreateCenRouteMapResponse) Validate() error {
	return dara.Validate(s)
}
