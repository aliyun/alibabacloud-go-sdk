// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMappCenterWorkspacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMappCenterWorkspacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMappCenterWorkspacesResponse
	GetStatusCode() *int32
	SetBody(v *ListMappCenterWorkspacesResponseBody) *ListMappCenterWorkspacesResponse
	GetBody() *ListMappCenterWorkspacesResponseBody
}

type ListMappCenterWorkspacesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMappCenterWorkspacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMappCenterWorkspacesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMappCenterWorkspacesResponse) GoString() string {
	return s.String()
}

func (s *ListMappCenterWorkspacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMappCenterWorkspacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMappCenterWorkspacesResponse) GetBody() *ListMappCenterWorkspacesResponseBody {
	return s.Body
}

func (s *ListMappCenterWorkspacesResponse) SetHeaders(v map[string]*string) *ListMappCenterWorkspacesResponse {
	s.Headers = v
	return s
}

func (s *ListMappCenterWorkspacesResponse) SetStatusCode(v int32) *ListMappCenterWorkspacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMappCenterWorkspacesResponse) SetBody(v *ListMappCenterWorkspacesResponseBody) *ListMappCenterWorkspacesResponse {
	s.Body = v
	return s
}

func (s *ListMappCenterWorkspacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
