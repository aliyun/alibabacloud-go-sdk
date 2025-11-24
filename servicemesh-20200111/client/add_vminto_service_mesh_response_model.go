// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVMIntoServiceMeshResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddVMIntoServiceMeshResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddVMIntoServiceMeshResponse
	GetStatusCode() *int32
	SetBody(v *AddVMIntoServiceMeshResponseBody) *AddVMIntoServiceMeshResponse
	GetBody() *AddVMIntoServiceMeshResponseBody
}

type AddVMIntoServiceMeshResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddVMIntoServiceMeshResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddVMIntoServiceMeshResponse) String() string {
	return dara.Prettify(s)
}

func (s AddVMIntoServiceMeshResponse) GoString() string {
	return s.String()
}

func (s *AddVMIntoServiceMeshResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddVMIntoServiceMeshResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddVMIntoServiceMeshResponse) GetBody() *AddVMIntoServiceMeshResponseBody {
	return s.Body
}

func (s *AddVMIntoServiceMeshResponse) SetHeaders(v map[string]*string) *AddVMIntoServiceMeshResponse {
	s.Headers = v
	return s
}

func (s *AddVMIntoServiceMeshResponse) SetStatusCode(v int32) *AddVMIntoServiceMeshResponse {
	s.StatusCode = &v
	return s
}

func (s *AddVMIntoServiceMeshResponse) SetBody(v *AddVMIntoServiceMeshResponseBody) *AddVMIntoServiceMeshResponse {
	s.Body = v
	return s
}

func (s *AddVMIntoServiceMeshResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
