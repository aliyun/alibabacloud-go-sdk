// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNacosClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNacosClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNacosClusterResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNacosClusterResponseBody) *UpdateNacosClusterResponse
	GetBody() *UpdateNacosClusterResponseBody
}

type UpdateNacosClusterResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNacosClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNacosClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNacosClusterResponse) GoString() string {
	return s.String()
}

func (s *UpdateNacosClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNacosClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNacosClusterResponse) GetBody() *UpdateNacosClusterResponseBody {
	return s.Body
}

func (s *UpdateNacosClusterResponse) SetHeaders(v map[string]*string) *UpdateNacosClusterResponse {
	s.Headers = v
	return s
}

func (s *UpdateNacosClusterResponse) SetStatusCode(v int32) *UpdateNacosClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNacosClusterResponse) SetBody(v *UpdateNacosClusterResponseBody) *UpdateNacosClusterResponse {
	s.Body = v
	return s
}

func (s *UpdateNacosClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
