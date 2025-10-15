// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNetworkZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNetworkZoneResponse
	GetStatusCode() *int32
	SetBody(v *GetNetworkZoneResponseBody) *GetNetworkZoneResponse
	GetBody() *GetNetworkZoneResponseBody
}

type GetNetworkZoneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNetworkZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNetworkZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkZoneResponse) GoString() string {
	return s.String()
}

func (s *GetNetworkZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNetworkZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNetworkZoneResponse) GetBody() *GetNetworkZoneResponseBody {
	return s.Body
}

func (s *GetNetworkZoneResponse) SetHeaders(v map[string]*string) *GetNetworkZoneResponse {
	s.Headers = v
	return s
}

func (s *GetNetworkZoneResponse) SetStatusCode(v int32) *GetNetworkZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNetworkZoneResponse) SetBody(v *GetNetworkZoneResponseBody) *GetNetworkZoneResponse {
	s.Body = v
	return s
}

func (s *GetNetworkZoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
