// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDedicatedBlockStorageClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDedicatedBlockStorageClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDedicatedBlockStorageClusterResponse
	GetStatusCode() *int32
	SetBody(v *CreateDedicatedBlockStorageClusterResponseBody) *CreateDedicatedBlockStorageClusterResponse
	GetBody() *CreateDedicatedBlockStorageClusterResponseBody
}

type CreateDedicatedBlockStorageClusterResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDedicatedBlockStorageClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDedicatedBlockStorageClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDedicatedBlockStorageClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateDedicatedBlockStorageClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDedicatedBlockStorageClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDedicatedBlockStorageClusterResponse) GetBody() *CreateDedicatedBlockStorageClusterResponseBody {
	return s.Body
}

func (s *CreateDedicatedBlockStorageClusterResponse) SetHeaders(v map[string]*string) *CreateDedicatedBlockStorageClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateDedicatedBlockStorageClusterResponse) SetStatusCode(v int32) *CreateDedicatedBlockStorageClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDedicatedBlockStorageClusterResponse) SetBody(v *CreateDedicatedBlockStorageClusterResponseBody) *CreateDedicatedBlockStorageClusterResponse {
	s.Body = v
	return s
}

func (s *CreateDedicatedBlockStorageClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
