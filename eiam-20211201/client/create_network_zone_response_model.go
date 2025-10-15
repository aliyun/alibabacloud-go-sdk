// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNetworkZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNetworkZoneResponse
	GetStatusCode() *int32
	SetBody(v *CreateNetworkZoneResponseBody) *CreateNetworkZoneResponse
	GetBody() *CreateNetworkZoneResponseBody
}

type CreateNetworkZoneResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNetworkZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNetworkZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkZoneResponse) GoString() string {
	return s.String()
}

func (s *CreateNetworkZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNetworkZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNetworkZoneResponse) GetBody() *CreateNetworkZoneResponseBody {
	return s.Body
}

func (s *CreateNetworkZoneResponse) SetHeaders(v map[string]*string) *CreateNetworkZoneResponse {
	s.Headers = v
	return s
}

func (s *CreateNetworkZoneResponse) SetStatusCode(v int32) *CreateNetworkZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNetworkZoneResponse) SetBody(v *CreateNetworkZoneResponseBody) *CreateNetworkZoneResponse {
	s.Body = v
	return s
}

func (s *CreateNetworkZoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
