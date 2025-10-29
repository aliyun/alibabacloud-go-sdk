// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExplorerRegistryModuleExamplesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExplorerRegistryModuleExamplesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExplorerRegistryModuleExamplesResponse
	GetStatusCode() *int32
	SetBody(v *ListExplorerRegistryModuleExamplesResponseBody) *ListExplorerRegistryModuleExamplesResponse
	GetBody() *ListExplorerRegistryModuleExamplesResponseBody
}

type ListExplorerRegistryModuleExamplesResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExplorerRegistryModuleExamplesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExplorerRegistryModuleExamplesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExplorerRegistryModuleExamplesResponse) GoString() string {
	return s.String()
}

func (s *ListExplorerRegistryModuleExamplesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExplorerRegistryModuleExamplesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExplorerRegistryModuleExamplesResponse) GetBody() *ListExplorerRegistryModuleExamplesResponseBody {
	return s.Body
}

func (s *ListExplorerRegistryModuleExamplesResponse) SetHeaders(v map[string]*string) *ListExplorerRegistryModuleExamplesResponse {
	s.Headers = v
	return s
}

func (s *ListExplorerRegistryModuleExamplesResponse) SetStatusCode(v int32) *ListExplorerRegistryModuleExamplesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExplorerRegistryModuleExamplesResponse) SetBody(v *ListExplorerRegistryModuleExamplesResponseBody) *ListExplorerRegistryModuleExamplesResponse {
	s.Body = v
	return s
}

func (s *ListExplorerRegistryModuleExamplesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
