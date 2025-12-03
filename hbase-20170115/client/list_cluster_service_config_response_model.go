// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterServiceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterServiceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterServiceConfigResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterServiceConfigResponseBody) *ListClusterServiceConfigResponse
	GetBody() *ListClusterServiceConfigResponseBody
}

type ListClusterServiceConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterServiceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterServiceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterServiceConfigResponse) GoString() string {
	return s.String()
}

func (s *ListClusterServiceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterServiceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterServiceConfigResponse) GetBody() *ListClusterServiceConfigResponseBody {
	return s.Body
}

func (s *ListClusterServiceConfigResponse) SetHeaders(v map[string]*string) *ListClusterServiceConfigResponse {
	s.Headers = v
	return s
}

func (s *ListClusterServiceConfigResponse) SetStatusCode(v int32) *ListClusterServiceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterServiceConfigResponse) SetBody(v *ListClusterServiceConfigResponseBody) *ListClusterServiceConfigResponse {
	s.Body = v
	return s
}

func (s *ListClusterServiceConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
