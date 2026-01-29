// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRayClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRayClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRayClusterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRayClusterResponseBody) *DeleteRayClusterResponse
	GetBody() *DeleteRayClusterResponseBody
}

type DeleteRayClusterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRayClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRayClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRayClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteRayClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRayClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRayClusterResponse) GetBody() *DeleteRayClusterResponseBody {
	return s.Body
}

func (s *DeleteRayClusterResponse) SetHeaders(v map[string]*string) *DeleteRayClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteRayClusterResponse) SetStatusCode(v int32) *DeleteRayClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRayClusterResponse) SetBody(v *DeleteRayClusterResponseBody) *DeleteRayClusterResponse {
	s.Body = v
	return s
}

func (s *DeleteRayClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
