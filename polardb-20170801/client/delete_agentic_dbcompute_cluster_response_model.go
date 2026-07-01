// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticDBComputeClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAgenticDBComputeClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAgenticDBComputeClusterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAgenticDBComputeClusterResponseBody) *DeleteAgenticDBComputeClusterResponse
	GetBody() *DeleteAgenticDBComputeClusterResponseBody
}

type DeleteAgenticDBComputeClusterResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAgenticDBComputeClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAgenticDBComputeClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticDBComputeClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgenticDBComputeClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAgenticDBComputeClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAgenticDBComputeClusterResponse) GetBody() *DeleteAgenticDBComputeClusterResponseBody {
	return s.Body
}

func (s *DeleteAgenticDBComputeClusterResponse) SetHeaders(v map[string]*string) *DeleteAgenticDBComputeClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgenticDBComputeClusterResponse) SetStatusCode(v int32) *DeleteAgenticDBComputeClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgenticDBComputeClusterResponse) SetBody(v *DeleteAgenticDBComputeClusterResponseBody) *DeleteAgenticDBComputeClusterResponse {
	s.Body = v
	return s
}

func (s *DeleteAgenticDBComputeClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
