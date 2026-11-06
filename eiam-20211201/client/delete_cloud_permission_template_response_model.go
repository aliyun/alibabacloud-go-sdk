// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudPermissionTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCloudPermissionTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCloudPermissionTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCloudPermissionTemplateResponseBody) *DeleteCloudPermissionTemplateResponse
	GetBody() *DeleteCloudPermissionTemplateResponseBody
}

type DeleteCloudPermissionTemplateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCloudPermissionTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCloudPermissionTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudPermissionTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteCloudPermissionTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCloudPermissionTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCloudPermissionTemplateResponse) GetBody() *DeleteCloudPermissionTemplateResponseBody {
	return s.Body
}

func (s *DeleteCloudPermissionTemplateResponse) SetHeaders(v map[string]*string) *DeleteCloudPermissionTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteCloudPermissionTemplateResponse) SetStatusCode(v int32) *DeleteCloudPermissionTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCloudPermissionTemplateResponse) SetBody(v *DeleteCloudPermissionTemplateResponseBody) *DeleteCloudPermissionTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteCloudPermissionTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
