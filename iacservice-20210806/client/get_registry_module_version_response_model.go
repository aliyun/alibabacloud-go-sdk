// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegistryModuleVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRegistryModuleVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRegistryModuleVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetRegistryModuleVersionResponseBody) *GetRegistryModuleVersionResponse
	GetBody() *GetRegistryModuleVersionResponseBody
}

type GetRegistryModuleVersionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRegistryModuleVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRegistryModuleVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRegistryModuleVersionResponse) GoString() string {
	return s.String()
}

func (s *GetRegistryModuleVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRegistryModuleVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRegistryModuleVersionResponse) GetBody() *GetRegistryModuleVersionResponseBody {
	return s.Body
}

func (s *GetRegistryModuleVersionResponse) SetHeaders(v map[string]*string) *GetRegistryModuleVersionResponse {
	s.Headers = v
	return s
}

func (s *GetRegistryModuleVersionResponse) SetStatusCode(v int32) *GetRegistryModuleVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRegistryModuleVersionResponse) SetBody(v *GetRegistryModuleVersionResponseBody) *GetRegistryModuleVersionResponse {
	s.Body = v
	return s
}

func (s *GetRegistryModuleVersionResponse) Validate() error {
	return dara.Validate(s)
}
