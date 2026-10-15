// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppPluginFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppPluginFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppPluginFilesResponse
	GetStatusCode() *int32
	SetBody(v *ListAppPluginFilesResponseBody) *ListAppPluginFilesResponse
	GetBody() *ListAppPluginFilesResponseBody
}

type ListAppPluginFilesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppPluginFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppPluginFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppPluginFilesResponse) GoString() string {
	return s.String()
}

func (s *ListAppPluginFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppPluginFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppPluginFilesResponse) GetBody() *ListAppPluginFilesResponseBody {
	return s.Body
}

func (s *ListAppPluginFilesResponse) SetHeaders(v map[string]*string) *ListAppPluginFilesResponse {
	s.Headers = v
	return s
}

func (s *ListAppPluginFilesResponse) SetStatusCode(v int32) *ListAppPluginFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppPluginFilesResponse) SetBody(v *ListAppPluginFilesResponseBody) *ListAppPluginFilesResponse {
	s.Body = v
	return s
}

func (s *ListAppPluginFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
