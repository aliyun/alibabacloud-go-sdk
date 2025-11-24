// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveVMFromServiceMeshResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveVMFromServiceMeshResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveVMFromServiceMeshResponse
	GetStatusCode() *int32
	SetBody(v *RemoveVMFromServiceMeshResponseBody) *RemoveVMFromServiceMeshResponse
	GetBody() *RemoveVMFromServiceMeshResponseBody
}

type RemoveVMFromServiceMeshResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveVMFromServiceMeshResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveVMFromServiceMeshResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveVMFromServiceMeshResponse) GoString() string {
	return s.String()
}

func (s *RemoveVMFromServiceMeshResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveVMFromServiceMeshResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveVMFromServiceMeshResponse) GetBody() *RemoveVMFromServiceMeshResponseBody {
	return s.Body
}

func (s *RemoveVMFromServiceMeshResponse) SetHeaders(v map[string]*string) *RemoveVMFromServiceMeshResponse {
	s.Headers = v
	return s
}

func (s *RemoveVMFromServiceMeshResponse) SetStatusCode(v int32) *RemoveVMFromServiceMeshResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveVMFromServiceMeshResponse) SetBody(v *RemoveVMFromServiceMeshResponseBody) *RemoveVMFromServiceMeshResponse {
	s.Body = v
	return s
}

func (s *RemoveVMFromServiceMeshResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
