// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVccRouteEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVccRouteEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVccRouteEntryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVccRouteEntryResponseBody) *DeleteVccRouteEntryResponse
	GetBody() *DeleteVccRouteEntryResponseBody
}

type DeleteVccRouteEntryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVccRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVccRouteEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVccRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *DeleteVccRouteEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVccRouteEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVccRouteEntryResponse) GetBody() *DeleteVccRouteEntryResponseBody {
	return s.Body
}

func (s *DeleteVccRouteEntryResponse) SetHeaders(v map[string]*string) *DeleteVccRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *DeleteVccRouteEntryResponse) SetStatusCode(v int32) *DeleteVccRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVccRouteEntryResponse) SetBody(v *DeleteVccRouteEntryResponseBody) *DeleteVccRouteEntryResponse {
	s.Body = v
	return s
}

func (s *DeleteVccRouteEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
