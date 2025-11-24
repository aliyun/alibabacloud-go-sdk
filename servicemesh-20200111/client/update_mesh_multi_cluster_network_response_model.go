// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMeshMultiClusterNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMeshMultiClusterNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMeshMultiClusterNetworkResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMeshMultiClusterNetworkResponseBody) *UpdateMeshMultiClusterNetworkResponse
	GetBody() *UpdateMeshMultiClusterNetworkResponseBody
}

type UpdateMeshMultiClusterNetworkResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMeshMultiClusterNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMeshMultiClusterNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeshMultiClusterNetworkResponse) GoString() string {
	return s.String()
}

func (s *UpdateMeshMultiClusterNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMeshMultiClusterNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMeshMultiClusterNetworkResponse) GetBody() *UpdateMeshMultiClusterNetworkResponseBody {
	return s.Body
}

func (s *UpdateMeshMultiClusterNetworkResponse) SetHeaders(v map[string]*string) *UpdateMeshMultiClusterNetworkResponse {
	s.Headers = v
	return s
}

func (s *UpdateMeshMultiClusterNetworkResponse) SetStatusCode(v int32) *UpdateMeshMultiClusterNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMeshMultiClusterNetworkResponse) SetBody(v *UpdateMeshMultiClusterNetworkResponseBody) *UpdateMeshMultiClusterNetworkResponse {
	s.Body = v
	return s
}

func (s *UpdateMeshMultiClusterNetworkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
