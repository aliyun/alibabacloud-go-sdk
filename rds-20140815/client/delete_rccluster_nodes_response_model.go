// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRCClusterNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRCClusterNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRCClusterNodesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRCClusterNodesResponseBody) *DeleteRCClusterNodesResponse
	GetBody() *DeleteRCClusterNodesResponseBody
}

type DeleteRCClusterNodesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRCClusterNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRCClusterNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRCClusterNodesResponse) GoString() string {
	return s.String()
}

func (s *DeleteRCClusterNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRCClusterNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRCClusterNodesResponse) GetBody() *DeleteRCClusterNodesResponseBody {
	return s.Body
}

func (s *DeleteRCClusterNodesResponse) SetHeaders(v map[string]*string) *DeleteRCClusterNodesResponse {
	s.Headers = v
	return s
}

func (s *DeleteRCClusterNodesResponse) SetStatusCode(v int32) *DeleteRCClusterNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRCClusterNodesResponse) SetBody(v *DeleteRCClusterNodesResponseBody) *DeleteRCClusterNodesResponse {
	s.Body = v
	return s
}

func (s *DeleteRCClusterNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
