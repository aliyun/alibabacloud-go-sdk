// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRegistryModuleVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRegistryModuleVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRegistryModuleVersionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRegistryModuleVersionResponseBody) *DeleteRegistryModuleVersionResponse
	GetBody() *DeleteRegistryModuleVersionResponseBody
}

type DeleteRegistryModuleVersionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRegistryModuleVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRegistryModuleVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRegistryModuleVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteRegistryModuleVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRegistryModuleVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRegistryModuleVersionResponse) GetBody() *DeleteRegistryModuleVersionResponseBody {
	return s.Body
}

func (s *DeleteRegistryModuleVersionResponse) SetHeaders(v map[string]*string) *DeleteRegistryModuleVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteRegistryModuleVersionResponse) SetStatusCode(v int32) *DeleteRegistryModuleVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRegistryModuleVersionResponse) SetBody(v *DeleteRegistryModuleVersionResponseBody) *DeleteRegistryModuleVersionResponse {
	s.Body = v
	return s
}

func (s *DeleteRegistryModuleVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
