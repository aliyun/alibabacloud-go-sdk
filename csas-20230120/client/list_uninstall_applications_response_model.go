// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUninstallApplicationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUninstallApplicationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUninstallApplicationsResponse
	GetStatusCode() *int32
	SetBody(v *ListUninstallApplicationsResponseBody) *ListUninstallApplicationsResponse
	GetBody() *ListUninstallApplicationsResponseBody
}

type ListUninstallApplicationsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUninstallApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUninstallApplicationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUninstallApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListUninstallApplicationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUninstallApplicationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUninstallApplicationsResponse) GetBody() *ListUninstallApplicationsResponseBody {
	return s.Body
}

func (s *ListUninstallApplicationsResponse) SetHeaders(v map[string]*string) *ListUninstallApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListUninstallApplicationsResponse) SetStatusCode(v int32) *ListUninstallApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUninstallApplicationsResponse) SetBody(v *ListUninstallApplicationsResponseBody) *ListUninstallApplicationsResponse {
	s.Body = v
	return s
}

func (s *ListUninstallApplicationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
