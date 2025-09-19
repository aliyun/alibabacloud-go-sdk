// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHybridProxyClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHybridProxyClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHybridProxyClusterResponse
	GetStatusCode() *int32
	SetBody(v *CreateHybridProxyClusterResponseBody) *CreateHybridProxyClusterResponse
	GetBody() *CreateHybridProxyClusterResponseBody
}

type CreateHybridProxyClusterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHybridProxyClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHybridProxyClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridProxyClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateHybridProxyClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHybridProxyClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHybridProxyClusterResponse) GetBody() *CreateHybridProxyClusterResponseBody {
	return s.Body
}

func (s *CreateHybridProxyClusterResponse) SetHeaders(v map[string]*string) *CreateHybridProxyClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateHybridProxyClusterResponse) SetStatusCode(v int32) *CreateHybridProxyClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHybridProxyClusterResponse) SetBody(v *CreateHybridProxyClusterResponseBody) *CreateHybridProxyClusterResponse {
	s.Body = v
	return s
}

func (s *CreateHybridProxyClusterResponse) Validate() error {
	return dara.Validate(s)
}
