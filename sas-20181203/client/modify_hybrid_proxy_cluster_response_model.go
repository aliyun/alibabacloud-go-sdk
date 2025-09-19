// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridProxyClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHybridProxyClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHybridProxyClusterResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHybridProxyClusterResponseBody) *ModifyHybridProxyClusterResponse
	GetBody() *ModifyHybridProxyClusterResponseBody
}

type ModifyHybridProxyClusterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHybridProxyClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHybridProxyClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridProxyClusterResponse) GoString() string {
	return s.String()
}

func (s *ModifyHybridProxyClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHybridProxyClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHybridProxyClusterResponse) GetBody() *ModifyHybridProxyClusterResponseBody {
	return s.Body
}

func (s *ModifyHybridProxyClusterResponse) SetHeaders(v map[string]*string) *ModifyHybridProxyClusterResponse {
	s.Headers = v
	return s
}

func (s *ModifyHybridProxyClusterResponse) SetStatusCode(v int32) *ModifyHybridProxyClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHybridProxyClusterResponse) SetBody(v *ModifyHybridProxyClusterResponseBody) *ModifyHybridProxyClusterResponse {
	s.Body = v
	return s
}

func (s *ModifyHybridProxyClusterResponse) Validate() error {
	return dara.Validate(s)
}
