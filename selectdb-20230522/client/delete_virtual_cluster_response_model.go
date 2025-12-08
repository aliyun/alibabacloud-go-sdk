// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirtualClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVirtualClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVirtualClusterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVirtualClusterResponseBody) *DeleteVirtualClusterResponse
	GetBody() *DeleteVirtualClusterResponseBody
}

type DeleteVirtualClusterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVirtualClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVirtualClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirtualClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteVirtualClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVirtualClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVirtualClusterResponse) GetBody() *DeleteVirtualClusterResponseBody {
	return s.Body
}

func (s *DeleteVirtualClusterResponse) SetHeaders(v map[string]*string) *DeleteVirtualClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteVirtualClusterResponse) SetStatusCode(v int32) *DeleteVirtualClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVirtualClusterResponse) SetBody(v *DeleteVirtualClusterResponseBody) *DeleteVirtualClusterResponse {
	s.Body = v
	return s
}

func (s *DeleteVirtualClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
