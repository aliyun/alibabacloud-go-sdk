// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVcoRouteEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVcoRouteEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVcoRouteEntryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVcoRouteEntryResponseBody) *DeleteVcoRouteEntryResponse
	GetBody() *DeleteVcoRouteEntryResponseBody
}

type DeleteVcoRouteEntryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVcoRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVcoRouteEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVcoRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *DeleteVcoRouteEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVcoRouteEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVcoRouteEntryResponse) GetBody() *DeleteVcoRouteEntryResponseBody {
	return s.Body
}

func (s *DeleteVcoRouteEntryResponse) SetHeaders(v map[string]*string) *DeleteVcoRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *DeleteVcoRouteEntryResponse) SetStatusCode(v int32) *DeleteVcoRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVcoRouteEntryResponse) SetBody(v *DeleteVcoRouteEntryResponseBody) *DeleteVcoRouteEntryResponse {
	s.Body = v
	return s
}

func (s *DeleteVcoRouteEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
