// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetworkZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNetworkZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNetworkZoneResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNetworkZoneResponseBody) *UpdateNetworkZoneResponse
	GetBody() *UpdateNetworkZoneResponseBody
}

type UpdateNetworkZoneResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNetworkZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNetworkZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkZoneResponse) GoString() string {
	return s.String()
}

func (s *UpdateNetworkZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNetworkZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNetworkZoneResponse) GetBody() *UpdateNetworkZoneResponseBody {
	return s.Body
}

func (s *UpdateNetworkZoneResponse) SetHeaders(v map[string]*string) *UpdateNetworkZoneResponse {
	s.Headers = v
	return s
}

func (s *UpdateNetworkZoneResponse) SetStatusCode(v int32) *UpdateNetworkZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNetworkZoneResponse) SetBody(v *UpdateNetworkZoneResponseBody) *UpdateNetworkZoneResponse {
	s.Body = v
	return s
}

func (s *UpdateNetworkZoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
