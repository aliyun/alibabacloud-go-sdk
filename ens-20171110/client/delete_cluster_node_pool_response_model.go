// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterNodePoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteClusterNodePoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteClusterNodePoolResponse
	GetStatusCode() *int32
	SetBody(v *DeleteClusterNodePoolResponseBody) *DeleteClusterNodePoolResponse
	GetBody() *DeleteClusterNodePoolResponseBody
}

type DeleteClusterNodePoolResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteClusterNodePoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteClusterNodePoolResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterNodePoolResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodePoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteClusterNodePoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteClusterNodePoolResponse) GetBody() *DeleteClusterNodePoolResponseBody {
	return s.Body
}

func (s *DeleteClusterNodePoolResponse) SetHeaders(v map[string]*string) *DeleteClusterNodePoolResponse {
	s.Headers = v
	return s
}

func (s *DeleteClusterNodePoolResponse) SetStatusCode(v int32) *DeleteClusterNodePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClusterNodePoolResponse) SetBody(v *DeleteClusterNodePoolResponseBody) *DeleteClusterNodePoolResponse {
	s.Body = v
	return s
}

func (s *DeleteClusterNodePoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
