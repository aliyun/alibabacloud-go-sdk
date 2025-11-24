// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceMeshResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceMeshResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceMeshResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceMeshResponseBody) *CreateServiceMeshResponse
	GetBody() *CreateServiceMeshResponseBody
}

type CreateServiceMeshResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceMeshResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceMeshResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceMeshResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceMeshResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceMeshResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceMeshResponse) GetBody() *CreateServiceMeshResponseBody {
	return s.Body
}

func (s *CreateServiceMeshResponse) SetHeaders(v map[string]*string) *CreateServiceMeshResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceMeshResponse) SetStatusCode(v int32) *CreateServiceMeshResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceMeshResponse) SetBody(v *CreateServiceMeshResponseBody) *CreateServiceMeshResponse {
	s.Body = v
	return s
}

func (s *CreateServiceMeshResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
