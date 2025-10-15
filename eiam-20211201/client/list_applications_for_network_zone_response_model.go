// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForNetworkZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationsForNetworkZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationsForNetworkZoneResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationsForNetworkZoneResponseBody) *ListApplicationsForNetworkZoneResponse
	GetBody() *ListApplicationsForNetworkZoneResponseBody
}

type ListApplicationsForNetworkZoneResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationsForNetworkZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationsForNetworkZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForNetworkZoneResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationsForNetworkZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationsForNetworkZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationsForNetworkZoneResponse) GetBody() *ListApplicationsForNetworkZoneResponseBody {
	return s.Body
}

func (s *ListApplicationsForNetworkZoneResponse) SetHeaders(v map[string]*string) *ListApplicationsForNetworkZoneResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationsForNetworkZoneResponse) SetStatusCode(v int32) *ListApplicationsForNetworkZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationsForNetworkZoneResponse) SetBody(v *ListApplicationsForNetworkZoneResponseBody) *ListApplicationsForNetworkZoneResponse {
	s.Body = v
	return s
}

func (s *ListApplicationsForNetworkZoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
