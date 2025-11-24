// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveClusterFromServiceMeshResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveClusterFromServiceMeshResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveClusterFromServiceMeshResponse
	GetStatusCode() *int32
	SetBody(v *RemoveClusterFromServiceMeshResponseBody) *RemoveClusterFromServiceMeshResponse
	GetBody() *RemoveClusterFromServiceMeshResponseBody
}

type RemoveClusterFromServiceMeshResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveClusterFromServiceMeshResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveClusterFromServiceMeshResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveClusterFromServiceMeshResponse) GoString() string {
	return s.String()
}

func (s *RemoveClusterFromServiceMeshResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveClusterFromServiceMeshResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveClusterFromServiceMeshResponse) GetBody() *RemoveClusterFromServiceMeshResponseBody {
	return s.Body
}

func (s *RemoveClusterFromServiceMeshResponse) SetHeaders(v map[string]*string) *RemoveClusterFromServiceMeshResponse {
	s.Headers = v
	return s
}

func (s *RemoveClusterFromServiceMeshResponse) SetStatusCode(v int32) *RemoveClusterFromServiceMeshResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveClusterFromServiceMeshResponse) SetBody(v *RemoveClusterFromServiceMeshResponseBody) *RemoveClusterFromServiceMeshResponse {
	s.Body = v
	return s
}

func (s *RemoveClusterFromServiceMeshResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
