// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConnectedClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteConnectedClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteConnectedClusterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteConnectedClusterResponseBody) *DeleteConnectedClusterResponse
	GetBody() *DeleteConnectedClusterResponseBody
}

type DeleteConnectedClusterResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConnectedClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConnectedClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteConnectedClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteConnectedClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteConnectedClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteConnectedClusterResponse) GetBody() *DeleteConnectedClusterResponseBody {
	return s.Body
}

func (s *DeleteConnectedClusterResponse) SetHeaders(v map[string]*string) *DeleteConnectedClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteConnectedClusterResponse) SetStatusCode(v int32) *DeleteConnectedClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConnectedClusterResponse) SetBody(v *DeleteConnectedClusterResponseBody) *DeleteConnectedClusterResponse {
	s.Body = v
	return s
}

func (s *DeleteConnectedClusterResponse) Validate() error {
	return dara.Validate(s)
}
