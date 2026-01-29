// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRayClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRayClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRayClusterResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRayClusterResponseBody) *UpdateRayClusterResponse
	GetBody() *UpdateRayClusterResponseBody
}

type UpdateRayClusterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRayClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRayClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRayClusterResponse) GoString() string {
	return s.String()
}

func (s *UpdateRayClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRayClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRayClusterResponse) GetBody() *UpdateRayClusterResponseBody {
	return s.Body
}

func (s *UpdateRayClusterResponse) SetHeaders(v map[string]*string) *UpdateRayClusterResponse {
	s.Headers = v
	return s
}

func (s *UpdateRayClusterResponse) SetStatusCode(v int32) *UpdateRayClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRayClusterResponse) SetBody(v *UpdateRayClusterResponseBody) *UpdateRayClusterResponse {
	s.Body = v
	return s
}

func (s *UpdateRayClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
