// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRegistryModuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRegistryModuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRegistryModuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRegistryModuleResponseBody) *DeleteRegistryModuleResponse
	GetBody() *DeleteRegistryModuleResponseBody
}

type DeleteRegistryModuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRegistryModuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRegistryModuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRegistryModuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRegistryModuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRegistryModuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRegistryModuleResponse) GetBody() *DeleteRegistryModuleResponseBody {
	return s.Body
}

func (s *DeleteRegistryModuleResponse) SetHeaders(v map[string]*string) *DeleteRegistryModuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRegistryModuleResponse) SetStatusCode(v int32) *DeleteRegistryModuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRegistryModuleResponse) SetBody(v *DeleteRegistryModuleResponseBody) *DeleteRegistryModuleResponse {
	s.Body = v
	return s
}

func (s *DeleteRegistryModuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
