// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudPermissionTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCloudPermissionTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCloudPermissionTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCloudPermissionTemplateResponseBody) *UpdateCloudPermissionTemplateResponse
	GetBody() *UpdateCloudPermissionTemplateResponseBody
}

type UpdateCloudPermissionTemplateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloudPermissionTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloudPermissionTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudPermissionTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudPermissionTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCloudPermissionTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCloudPermissionTemplateResponse) GetBody() *UpdateCloudPermissionTemplateResponseBody {
	return s.Body
}

func (s *UpdateCloudPermissionTemplateResponse) SetHeaders(v map[string]*string) *UpdateCloudPermissionTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudPermissionTemplateResponse) SetStatusCode(v int32) *UpdateCloudPermissionTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloudPermissionTemplateResponse) SetBody(v *UpdateCloudPermissionTemplateResponseBody) *UpdateCloudPermissionTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateCloudPermissionTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
