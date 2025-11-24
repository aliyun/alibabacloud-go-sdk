// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddClusterIntoServiceMeshResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddClusterIntoServiceMeshResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddClusterIntoServiceMeshResponse
	GetStatusCode() *int32
	SetBody(v *AddClusterIntoServiceMeshResponseBody) *AddClusterIntoServiceMeshResponse
	GetBody() *AddClusterIntoServiceMeshResponseBody
}

type AddClusterIntoServiceMeshResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddClusterIntoServiceMeshResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddClusterIntoServiceMeshResponse) String() string {
	return dara.Prettify(s)
}

func (s AddClusterIntoServiceMeshResponse) GoString() string {
	return s.String()
}

func (s *AddClusterIntoServiceMeshResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddClusterIntoServiceMeshResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddClusterIntoServiceMeshResponse) GetBody() *AddClusterIntoServiceMeshResponseBody {
	return s.Body
}

func (s *AddClusterIntoServiceMeshResponse) SetHeaders(v map[string]*string) *AddClusterIntoServiceMeshResponse {
	s.Headers = v
	return s
}

func (s *AddClusterIntoServiceMeshResponse) SetStatusCode(v int32) *AddClusterIntoServiceMeshResponse {
	s.StatusCode = &v
	return s
}

func (s *AddClusterIntoServiceMeshResponse) SetBody(v *AddClusterIntoServiceMeshResponseBody) *AddClusterIntoServiceMeshResponse {
	s.Body = v
	return s
}

func (s *AddClusterIntoServiceMeshResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
