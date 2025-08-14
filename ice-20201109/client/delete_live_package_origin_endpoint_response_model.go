// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivePackageOriginEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLivePackageOriginEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLivePackageOriginEndpointResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLivePackageOriginEndpointResponseBody) *DeleteLivePackageOriginEndpointResponse
	GetBody() *DeleteLivePackageOriginEndpointResponseBody
}

type DeleteLivePackageOriginEndpointResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLivePackageOriginEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLivePackageOriginEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivePackageOriginEndpointResponse) GoString() string {
	return s.String()
}

func (s *DeleteLivePackageOriginEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLivePackageOriginEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLivePackageOriginEndpointResponse) GetBody() *DeleteLivePackageOriginEndpointResponseBody {
	return s.Body
}

func (s *DeleteLivePackageOriginEndpointResponse) SetHeaders(v map[string]*string) *DeleteLivePackageOriginEndpointResponse {
	s.Headers = v
	return s
}

func (s *DeleteLivePackageOriginEndpointResponse) SetStatusCode(v int32) *DeleteLivePackageOriginEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLivePackageOriginEndpointResponse) SetBody(v *DeleteLivePackageOriginEndpointResponseBody) *DeleteLivePackageOriginEndpointResponse {
	s.Body = v
	return s
}

func (s *DeleteLivePackageOriginEndpointResponse) Validate() error {
	return dara.Validate(s)
}
