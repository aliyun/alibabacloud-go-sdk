// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloudAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloudAppsResponse
	GetStatusCode() *int32
	SetBody(v *ListCloudAppsResponseBody) *ListCloudAppsResponse
	GetBody() *ListCloudAppsResponseBody
}

type ListCloudAppsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAppsResponse) GoString() string {
	return s.String()
}

func (s *ListCloudAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloudAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloudAppsResponse) GetBody() *ListCloudAppsResponseBody {
	return s.Body
}

func (s *ListCloudAppsResponse) SetHeaders(v map[string]*string) *ListCloudAppsResponse {
	s.Headers = v
	return s
}

func (s *ListCloudAppsResponse) SetStatusCode(v int32) *ListCloudAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudAppsResponse) SetBody(v *ListCloudAppsResponseBody) *ListCloudAppsResponse {
	s.Body = v
	return s
}

func (s *ListCloudAppsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
