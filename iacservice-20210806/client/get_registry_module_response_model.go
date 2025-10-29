// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegistryModuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRegistryModuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRegistryModuleResponse
	GetStatusCode() *int32
	SetBody(v *GetRegistryModuleResponseBody) *GetRegistryModuleResponse
	GetBody() *GetRegistryModuleResponseBody
}

type GetRegistryModuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRegistryModuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRegistryModuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRegistryModuleResponse) GoString() string {
	return s.String()
}

func (s *GetRegistryModuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRegistryModuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRegistryModuleResponse) GetBody() *GetRegistryModuleResponseBody {
	return s.Body
}

func (s *GetRegistryModuleResponse) SetHeaders(v map[string]*string) *GetRegistryModuleResponse {
	s.Headers = v
	return s
}

func (s *GetRegistryModuleResponse) SetStatusCode(v int32) *GetRegistryModuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRegistryModuleResponse) SetBody(v *GetRegistryModuleResponseBody) *GetRegistryModuleResponse {
	s.Body = v
	return s
}

func (s *GetRegistryModuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
