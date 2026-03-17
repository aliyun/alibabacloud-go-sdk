// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudConnectNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCloudConnectNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCloudConnectNetworkResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCloudConnectNetworkResponseBody) *DeleteCloudConnectNetworkResponse
	GetBody() *DeleteCloudConnectNetworkResponseBody
}

type DeleteCloudConnectNetworkResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCloudConnectNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCloudConnectNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudConnectNetworkResponse) GoString() string {
	return s.String()
}

func (s *DeleteCloudConnectNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCloudConnectNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCloudConnectNetworkResponse) GetBody() *DeleteCloudConnectNetworkResponseBody {
	return s.Body
}

func (s *DeleteCloudConnectNetworkResponse) SetHeaders(v map[string]*string) *DeleteCloudConnectNetworkResponse {
	s.Headers = v
	return s
}

func (s *DeleteCloudConnectNetworkResponse) SetStatusCode(v int32) *DeleteCloudConnectNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCloudConnectNetworkResponse) SetBody(v *DeleteCloudConnectNetworkResponseBody) *DeleteCloudConnectNetworkResponse {
	s.Body = v
	return s
}

func (s *DeleteCloudConnectNetworkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
