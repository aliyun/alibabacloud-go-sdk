// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNetworkZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNetworkZoneResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNetworkZoneResponseBody) *DeleteNetworkZoneResponse
	GetBody() *DeleteNetworkZoneResponseBody
}

type DeleteNetworkZoneResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNetworkZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNetworkZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkZoneResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNetworkZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNetworkZoneResponse) GetBody() *DeleteNetworkZoneResponseBody {
	return s.Body
}

func (s *DeleteNetworkZoneResponse) SetHeaders(v map[string]*string) *DeleteNetworkZoneResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetworkZoneResponse) SetStatusCode(v int32) *DeleteNetworkZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNetworkZoneResponse) SetBody(v *DeleteNetworkZoneResponseBody) *DeleteNetworkZoneResponse {
	s.Body = v
	return s
}

func (s *DeleteNetworkZoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
