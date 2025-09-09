// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkAvailableZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNetworkAvailableZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNetworkAvailableZonesResponse
	GetStatusCode() *int32
	SetBody(v *GetNetworkAvailableZonesResponseBody) *GetNetworkAvailableZonesResponse
	GetBody() *GetNetworkAvailableZonesResponseBody
}

type GetNetworkAvailableZonesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNetworkAvailableZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNetworkAvailableZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkAvailableZonesResponse) GoString() string {
	return s.String()
}

func (s *GetNetworkAvailableZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNetworkAvailableZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNetworkAvailableZonesResponse) GetBody() *GetNetworkAvailableZonesResponseBody {
	return s.Body
}

func (s *GetNetworkAvailableZonesResponse) SetHeaders(v map[string]*string) *GetNetworkAvailableZonesResponse {
	s.Headers = v
	return s
}

func (s *GetNetworkAvailableZonesResponse) SetStatusCode(v int32) *GetNetworkAvailableZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNetworkAvailableZonesResponse) SetBody(v *GetNetworkAvailableZonesResponseBody) *GetNetworkAvailableZonesResponse {
	s.Body = v
	return s
}

func (s *GetNetworkAvailableZonesResponse) Validate() error {
	return dara.Validate(s)
}
