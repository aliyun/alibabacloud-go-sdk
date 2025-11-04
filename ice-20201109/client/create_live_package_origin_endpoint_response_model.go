// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLivePackageOriginEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLivePackageOriginEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLivePackageOriginEndpointResponse
	GetStatusCode() *int32
	SetBody(v *CreateLivePackageOriginEndpointResponseBody) *CreateLivePackageOriginEndpointResponse
	GetBody() *CreateLivePackageOriginEndpointResponseBody
}

type CreateLivePackageOriginEndpointResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLivePackageOriginEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLivePackageOriginEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLivePackageOriginEndpointResponse) GoString() string {
	return s.String()
}

func (s *CreateLivePackageOriginEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLivePackageOriginEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLivePackageOriginEndpointResponse) GetBody() *CreateLivePackageOriginEndpointResponseBody {
	return s.Body
}

func (s *CreateLivePackageOriginEndpointResponse) SetHeaders(v map[string]*string) *CreateLivePackageOriginEndpointResponse {
	s.Headers = v
	return s
}

func (s *CreateLivePackageOriginEndpointResponse) SetStatusCode(v int32) *CreateLivePackageOriginEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLivePackageOriginEndpointResponse) SetBody(v *CreateLivePackageOriginEndpointResponseBody) *CreateLivePackageOriginEndpointResponse {
	s.Body = v
	return s
}

func (s *CreateLivePackageOriginEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
