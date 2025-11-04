// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivePackageOriginEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLivePackageOriginEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLivePackageOriginEndpointResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLivePackageOriginEndpointResponseBody) *UpdateLivePackageOriginEndpointResponse
	GetBody() *UpdateLivePackageOriginEndpointResponseBody
}

type UpdateLivePackageOriginEndpointResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLivePackageOriginEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLivePackageOriginEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePackageOriginEndpointResponse) GoString() string {
	return s.String()
}

func (s *UpdateLivePackageOriginEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLivePackageOriginEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLivePackageOriginEndpointResponse) GetBody() *UpdateLivePackageOriginEndpointResponseBody {
	return s.Body
}

func (s *UpdateLivePackageOriginEndpointResponse) SetHeaders(v map[string]*string) *UpdateLivePackageOriginEndpointResponse {
	s.Headers = v
	return s
}

func (s *UpdateLivePackageOriginEndpointResponse) SetStatusCode(v int32) *UpdateLivePackageOriginEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointResponse) SetBody(v *UpdateLivePackageOriginEndpointResponseBody) *UpdateLivePackageOriginEndpointResponse {
	s.Body = v
	return s
}

func (s *UpdateLivePackageOriginEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
