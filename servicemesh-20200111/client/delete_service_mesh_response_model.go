// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceMeshResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteServiceMeshResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteServiceMeshResponse
	GetStatusCode() *int32
	SetBody(v *DeleteServiceMeshResponseBody) *DeleteServiceMeshResponse
	GetBody() *DeleteServiceMeshResponseBody
}

type DeleteServiceMeshResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceMeshResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceMeshResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceMeshResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceMeshResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteServiceMeshResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteServiceMeshResponse) GetBody() *DeleteServiceMeshResponseBody {
	return s.Body
}

func (s *DeleteServiceMeshResponse) SetHeaders(v map[string]*string) *DeleteServiceMeshResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceMeshResponse) SetStatusCode(v int32) *DeleteServiceMeshResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceMeshResponse) SetBody(v *DeleteServiceMeshResponseBody) *DeleteServiceMeshResponse {
	s.Body = v
	return s
}

func (s *DeleteServiceMeshResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
