// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMappCenterAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMappCenterAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMappCenterAppsResponse
	GetStatusCode() *int32
	SetBody(v *ListMappCenterAppsResponseBody) *ListMappCenterAppsResponse
	GetBody() *ListMappCenterAppsResponseBody
}

type ListMappCenterAppsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMappCenterAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMappCenterAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMappCenterAppsResponse) GoString() string {
	return s.String()
}

func (s *ListMappCenterAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMappCenterAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMappCenterAppsResponse) GetBody() *ListMappCenterAppsResponseBody {
	return s.Body
}

func (s *ListMappCenterAppsResponse) SetHeaders(v map[string]*string) *ListMappCenterAppsResponse {
	s.Headers = v
	return s
}

func (s *ListMappCenterAppsResponse) SetStatusCode(v int32) *ListMappCenterAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMappCenterAppsResponse) SetBody(v *ListMappCenterAppsResponseBody) *ListMappCenterAppsResponse {
	s.Body = v
	return s
}

func (s *ListMappCenterAppsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
