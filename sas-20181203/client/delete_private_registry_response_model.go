// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateRegistryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePrivateRegistryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePrivateRegistryResponse
	GetStatusCode() *int32
	SetBody(v *DeletePrivateRegistryResponseBody) *DeletePrivateRegistryResponse
	GetBody() *DeletePrivateRegistryResponseBody
}

type DeletePrivateRegistryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePrivateRegistryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePrivateRegistryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateRegistryResponse) GoString() string {
	return s.String()
}

func (s *DeletePrivateRegistryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePrivateRegistryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePrivateRegistryResponse) GetBody() *DeletePrivateRegistryResponseBody {
	return s.Body
}

func (s *DeletePrivateRegistryResponse) SetHeaders(v map[string]*string) *DeletePrivateRegistryResponse {
	s.Headers = v
	return s
}

func (s *DeletePrivateRegistryResponse) SetStatusCode(v int32) *DeletePrivateRegistryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrivateRegistryResponse) SetBody(v *DeletePrivateRegistryResponseBody) *DeletePrivateRegistryResponse {
	s.Body = v
	return s
}

func (s *DeletePrivateRegistryResponse) Validate() error {
	return dara.Validate(s)
}
