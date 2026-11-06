// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudPermissionTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloudPermissionTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloudPermissionTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListCloudPermissionTemplatesResponseBody) *ListCloudPermissionTemplatesResponse
	GetBody() *ListCloudPermissionTemplatesResponseBody
}

type ListCloudPermissionTemplatesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudPermissionTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudPermissionTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloudPermissionTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListCloudPermissionTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloudPermissionTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloudPermissionTemplatesResponse) GetBody() *ListCloudPermissionTemplatesResponseBody {
	return s.Body
}

func (s *ListCloudPermissionTemplatesResponse) SetHeaders(v map[string]*string) *ListCloudPermissionTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListCloudPermissionTemplatesResponse) SetStatusCode(v int32) *ListCloudPermissionTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudPermissionTemplatesResponse) SetBody(v *ListCloudPermissionTemplatesResponseBody) *ListCloudPermissionTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListCloudPermissionTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
