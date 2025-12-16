// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEdgeContainerAppImageSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEdgeContainerAppImageSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEdgeContainerAppImageSecretResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEdgeContainerAppImageSecretResponseBody) *DeleteEdgeContainerAppImageSecretResponse
	GetBody() *DeleteEdgeContainerAppImageSecretResponseBody
}

type DeleteEdgeContainerAppImageSecretResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEdgeContainerAppImageSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEdgeContainerAppImageSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEdgeContainerAppImageSecretResponse) GoString() string {
	return s.String()
}

func (s *DeleteEdgeContainerAppImageSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEdgeContainerAppImageSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEdgeContainerAppImageSecretResponse) GetBody() *DeleteEdgeContainerAppImageSecretResponseBody {
	return s.Body
}

func (s *DeleteEdgeContainerAppImageSecretResponse) SetHeaders(v map[string]*string) *DeleteEdgeContainerAppImageSecretResponse {
	s.Headers = v
	return s
}

func (s *DeleteEdgeContainerAppImageSecretResponse) SetStatusCode(v int32) *DeleteEdgeContainerAppImageSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEdgeContainerAppImageSecretResponse) SetBody(v *DeleteEdgeContainerAppImageSecretResponseBody) *DeleteEdgeContainerAppImageSecretResponse {
	s.Body = v
	return s
}

func (s *DeleteEdgeContainerAppImageSecretResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
