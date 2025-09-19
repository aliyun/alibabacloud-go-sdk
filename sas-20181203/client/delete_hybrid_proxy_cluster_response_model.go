// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHybridProxyClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHybridProxyClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHybridProxyClusterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHybridProxyClusterResponseBody) *DeleteHybridProxyClusterResponse
	GetBody() *DeleteHybridProxyClusterResponseBody
}

type DeleteHybridProxyClusterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHybridProxyClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHybridProxyClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHybridProxyClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteHybridProxyClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHybridProxyClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHybridProxyClusterResponse) GetBody() *DeleteHybridProxyClusterResponseBody {
	return s.Body
}

func (s *DeleteHybridProxyClusterResponse) SetHeaders(v map[string]*string) *DeleteHybridProxyClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteHybridProxyClusterResponse) SetStatusCode(v int32) *DeleteHybridProxyClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHybridProxyClusterResponse) SetBody(v *DeleteHybridProxyClusterResponseBody) *DeleteHybridProxyClusterResponse {
	s.Body = v
	return s
}

func (s *DeleteHybridProxyClusterResponse) Validate() error {
	return dara.Validate(s)
}
