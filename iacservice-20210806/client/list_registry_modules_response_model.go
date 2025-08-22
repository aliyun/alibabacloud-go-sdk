// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegistryModulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRegistryModulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRegistryModulesResponse
	GetStatusCode() *int32
	SetBody(v *ListRegistryModulesResponseBody) *ListRegistryModulesResponse
	GetBody() *ListRegistryModulesResponseBody
}

type ListRegistryModulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRegistryModulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRegistryModulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRegistryModulesResponse) GoString() string {
	return s.String()
}

func (s *ListRegistryModulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRegistryModulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRegistryModulesResponse) GetBody() *ListRegistryModulesResponseBody {
	return s.Body
}

func (s *ListRegistryModulesResponse) SetHeaders(v map[string]*string) *ListRegistryModulesResponse {
	s.Headers = v
	return s
}

func (s *ListRegistryModulesResponse) SetStatusCode(v int32) *ListRegistryModulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegistryModulesResponse) SetBody(v *ListRegistryModulesResponseBody) *ListRegistryModulesResponse {
	s.Body = v
	return s
}

func (s *ListRegistryModulesResponse) Validate() error {
	return dara.Validate(s)
}
