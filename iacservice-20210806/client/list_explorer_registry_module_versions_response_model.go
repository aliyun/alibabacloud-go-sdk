// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExplorerRegistryModuleVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExplorerRegistryModuleVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExplorerRegistryModuleVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListExplorerRegistryModuleVersionsResponseBody) *ListExplorerRegistryModuleVersionsResponse
	GetBody() *ListExplorerRegistryModuleVersionsResponseBody
}

type ListExplorerRegistryModuleVersionsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExplorerRegistryModuleVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExplorerRegistryModuleVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExplorerRegistryModuleVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListExplorerRegistryModuleVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExplorerRegistryModuleVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExplorerRegistryModuleVersionsResponse) GetBody() *ListExplorerRegistryModuleVersionsResponseBody {
	return s.Body
}

func (s *ListExplorerRegistryModuleVersionsResponse) SetHeaders(v map[string]*string) *ListExplorerRegistryModuleVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListExplorerRegistryModuleVersionsResponse) SetStatusCode(v int32) *ListExplorerRegistryModuleVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExplorerRegistryModuleVersionsResponse) SetBody(v *ListExplorerRegistryModuleVersionsResponseBody) *ListExplorerRegistryModuleVersionsResponse {
	s.Body = v
	return s
}

func (s *ListExplorerRegistryModuleVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
