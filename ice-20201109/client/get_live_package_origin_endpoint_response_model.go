// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLivePackageOriginEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLivePackageOriginEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLivePackageOriginEndpointResponse
	GetStatusCode() *int32
	SetBody(v *GetLivePackageOriginEndpointResponseBody) *GetLivePackageOriginEndpointResponse
	GetBody() *GetLivePackageOriginEndpointResponseBody
}

type GetLivePackageOriginEndpointResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLivePackageOriginEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLivePackageOriginEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLivePackageOriginEndpointResponse) GoString() string {
	return s.String()
}

func (s *GetLivePackageOriginEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLivePackageOriginEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLivePackageOriginEndpointResponse) GetBody() *GetLivePackageOriginEndpointResponseBody {
	return s.Body
}

func (s *GetLivePackageOriginEndpointResponse) SetHeaders(v map[string]*string) *GetLivePackageOriginEndpointResponse {
	s.Headers = v
	return s
}

func (s *GetLivePackageOriginEndpointResponse) SetStatusCode(v int32) *GetLivePackageOriginEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLivePackageOriginEndpointResponse) SetBody(v *GetLivePackageOriginEndpointResponseBody) *GetLivePackageOriginEndpointResponse {
	s.Body = v
	return s
}

func (s *GetLivePackageOriginEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
