// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteClusterNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteClusterNodesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteClusterNodesResponseBody) *DeleteClusterNodesResponse
	GetBody() *DeleteClusterNodesResponseBody
}

type DeleteClusterNodesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteClusterNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteClusterNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterNodesResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteClusterNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteClusterNodesResponse) GetBody() *DeleteClusterNodesResponseBody {
	return s.Body
}

func (s *DeleteClusterNodesResponse) SetHeaders(v map[string]*string) *DeleteClusterNodesResponse {
	s.Headers = v
	return s
}

func (s *DeleteClusterNodesResponse) SetStatusCode(v int32) *DeleteClusterNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClusterNodesResponse) SetBody(v *DeleteClusterNodesResponseBody) *DeleteClusterNodesResponse {
	s.Body = v
	return s
}

func (s *DeleteClusterNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
