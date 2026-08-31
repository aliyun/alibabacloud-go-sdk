// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComputeClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteComputeClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteComputeClusterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteComputeClusterResponseBody) *DeleteComputeClusterResponse
	GetBody() *DeleteComputeClusterResponseBody
}

type DeleteComputeClusterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteComputeClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteComputeClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteComputeClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteComputeClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteComputeClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteComputeClusterResponse) GetBody() *DeleteComputeClusterResponseBody {
	return s.Body
}

func (s *DeleteComputeClusterResponse) SetHeaders(v map[string]*string) *DeleteComputeClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteComputeClusterResponse) SetStatusCode(v int32) *DeleteComputeClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteComputeClusterResponse) SetBody(v *DeleteComputeClusterResponseBody) *DeleteComputeClusterResponse {
	s.Body = v
	return s
}

func (s *DeleteComputeClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
