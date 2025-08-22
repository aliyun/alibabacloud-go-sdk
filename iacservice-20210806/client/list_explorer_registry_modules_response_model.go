// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExplorerRegistryModulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExplorerRegistryModulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExplorerRegistryModulesResponse
	GetStatusCode() *int32
	SetBody(v *ListExplorerRegistryModulesResponseBody) *ListExplorerRegistryModulesResponse
	GetBody() *ListExplorerRegistryModulesResponseBody
}

type ListExplorerRegistryModulesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExplorerRegistryModulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExplorerRegistryModulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExplorerRegistryModulesResponse) GoString() string {
	return s.String()
}

func (s *ListExplorerRegistryModulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExplorerRegistryModulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExplorerRegistryModulesResponse) GetBody() *ListExplorerRegistryModulesResponseBody {
	return s.Body
}

func (s *ListExplorerRegistryModulesResponse) SetHeaders(v map[string]*string) *ListExplorerRegistryModulesResponse {
	s.Headers = v
	return s
}

func (s *ListExplorerRegistryModulesResponse) SetStatusCode(v int32) *ListExplorerRegistryModulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExplorerRegistryModulesResponse) SetBody(v *ListExplorerRegistryModulesResponseBody) *ListExplorerRegistryModulesResponse {
	s.Body = v
	return s
}

func (s *ListExplorerRegistryModulesResponse) Validate() error {
	return dara.Validate(s)
}
