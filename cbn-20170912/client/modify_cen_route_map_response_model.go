// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCenRouteMapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCenRouteMapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCenRouteMapResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCenRouteMapResponseBody) *ModifyCenRouteMapResponse
	GetBody() *ModifyCenRouteMapResponseBody
}

type ModifyCenRouteMapResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCenRouteMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCenRouteMapResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenRouteMapResponse) GoString() string {
	return s.String()
}

func (s *ModifyCenRouteMapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCenRouteMapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCenRouteMapResponse) GetBody() *ModifyCenRouteMapResponseBody {
	return s.Body
}

func (s *ModifyCenRouteMapResponse) SetHeaders(v map[string]*string) *ModifyCenRouteMapResponse {
	s.Headers = v
	return s
}

func (s *ModifyCenRouteMapResponse) SetStatusCode(v int32) *ModifyCenRouteMapResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCenRouteMapResponse) SetBody(v *ModifyCenRouteMapResponseBody) *ModifyCenRouteMapResponse {
	s.Body = v
	return s
}

func (s *ModifyCenRouteMapResponse) Validate() error {
	return dara.Validate(s)
}
