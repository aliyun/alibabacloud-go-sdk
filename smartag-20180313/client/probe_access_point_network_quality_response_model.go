// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProbeAccessPointNetworkQualityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ProbeAccessPointNetworkQualityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ProbeAccessPointNetworkQualityResponse
	GetStatusCode() *int32
	SetBody(v *ProbeAccessPointNetworkQualityResponseBody) *ProbeAccessPointNetworkQualityResponse
	GetBody() *ProbeAccessPointNetworkQualityResponseBody
}

type ProbeAccessPointNetworkQualityResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProbeAccessPointNetworkQualityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ProbeAccessPointNetworkQualityResponse) String() string {
	return dara.Prettify(s)
}

func (s ProbeAccessPointNetworkQualityResponse) GoString() string {
	return s.String()
}

func (s *ProbeAccessPointNetworkQualityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ProbeAccessPointNetworkQualityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ProbeAccessPointNetworkQualityResponse) GetBody() *ProbeAccessPointNetworkQualityResponseBody {
	return s.Body
}

func (s *ProbeAccessPointNetworkQualityResponse) SetHeaders(v map[string]*string) *ProbeAccessPointNetworkQualityResponse {
	s.Headers = v
	return s
}

func (s *ProbeAccessPointNetworkQualityResponse) SetStatusCode(v int32) *ProbeAccessPointNetworkQualityResponse {
	s.StatusCode = &v
	return s
}

func (s *ProbeAccessPointNetworkQualityResponse) SetBody(v *ProbeAccessPointNetworkQualityResponseBody) *ProbeAccessPointNetworkQualityResponse {
	s.Body = v
	return s
}

func (s *ProbeAccessPointNetworkQualityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
