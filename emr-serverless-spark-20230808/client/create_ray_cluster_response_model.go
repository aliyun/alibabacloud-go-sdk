// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRayClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRayClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRayClusterResponse
	GetStatusCode() *int32
	SetBody(v *CreateRayClusterResponseBody) *CreateRayClusterResponse
	GetBody() *CreateRayClusterResponseBody
}

type CreateRayClusterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRayClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRayClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRayClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateRayClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRayClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRayClusterResponse) GetBody() *CreateRayClusterResponseBody {
	return s.Body
}

func (s *CreateRayClusterResponse) SetHeaders(v map[string]*string) *CreateRayClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateRayClusterResponse) SetStatusCode(v int32) *CreateRayClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRayClusterResponse) SetBody(v *CreateRayClusterResponseBody) *CreateRayClusterResponse {
	s.Body = v
	return s
}

func (s *CreateRayClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
