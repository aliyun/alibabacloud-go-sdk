// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeHotpatchResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMcubeHotpatchResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMcubeHotpatchResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListMcubeHotpatchResourcesResponseBody) *ListMcubeHotpatchResourcesResponse
	GetBody() *ListMcubeHotpatchResourcesResponseBody
}

type ListMcubeHotpatchResourcesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMcubeHotpatchResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMcubeHotpatchResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeHotpatchResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListMcubeHotpatchResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMcubeHotpatchResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMcubeHotpatchResourcesResponse) GetBody() *ListMcubeHotpatchResourcesResponseBody {
	return s.Body
}

func (s *ListMcubeHotpatchResourcesResponse) SetHeaders(v map[string]*string) *ListMcubeHotpatchResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListMcubeHotpatchResourcesResponse) SetStatusCode(v int32) *ListMcubeHotpatchResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMcubeHotpatchResourcesResponse) SetBody(v *ListMcubeHotpatchResourcesResponseBody) *ListMcubeHotpatchResourcesResponse {
	s.Body = v
	return s
}

func (s *ListMcubeHotpatchResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
