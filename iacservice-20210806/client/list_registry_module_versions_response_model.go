// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegistryModuleVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRegistryModuleVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRegistryModuleVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListRegistryModuleVersionsResponseBody) *ListRegistryModuleVersionsResponse
	GetBody() *ListRegistryModuleVersionsResponseBody
}

type ListRegistryModuleVersionsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRegistryModuleVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRegistryModuleVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRegistryModuleVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListRegistryModuleVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRegistryModuleVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRegistryModuleVersionsResponse) GetBody() *ListRegistryModuleVersionsResponseBody {
	return s.Body
}

func (s *ListRegistryModuleVersionsResponse) SetHeaders(v map[string]*string) *ListRegistryModuleVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListRegistryModuleVersionsResponse) SetStatusCode(v int32) *ListRegistryModuleVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegistryModuleVersionsResponse) SetBody(v *ListRegistryModuleVersionsResponseBody) *ListRegistryModuleVersionsResponse {
	s.Body = v
	return s
}

func (s *ListRegistryModuleVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
