// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLivePackageOriginEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLivePackageOriginEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLivePackageOriginEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *ListLivePackageOriginEndpointsResponseBody) *ListLivePackageOriginEndpointsResponse
	GetBody() *ListLivePackageOriginEndpointsResponseBody
}

type ListLivePackageOriginEndpointsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLivePackageOriginEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLivePackageOriginEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLivePackageOriginEndpointsResponse) GoString() string {
	return s.String()
}

func (s *ListLivePackageOriginEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLivePackageOriginEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLivePackageOriginEndpointsResponse) GetBody() *ListLivePackageOriginEndpointsResponseBody {
	return s.Body
}

func (s *ListLivePackageOriginEndpointsResponse) SetHeaders(v map[string]*string) *ListLivePackageOriginEndpointsResponse {
	s.Headers = v
	return s
}

func (s *ListLivePackageOriginEndpointsResponse) SetStatusCode(v int32) *ListLivePackageOriginEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLivePackageOriginEndpointsResponse) SetBody(v *ListLivePackageOriginEndpointsResponseBody) *ListLivePackageOriginEndpointsResponse {
	s.Body = v
	return s
}

func (s *ListLivePackageOriginEndpointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
