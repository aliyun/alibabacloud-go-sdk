// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCloudPermissionTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableCloudPermissionTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableCloudPermissionTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DisableCloudPermissionTemplateResponseBody) *DisableCloudPermissionTemplateResponse
	GetBody() *DisableCloudPermissionTemplateResponseBody
}

type DisableCloudPermissionTemplateResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableCloudPermissionTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableCloudPermissionTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableCloudPermissionTemplateResponse) GoString() string {
	return s.String()
}

func (s *DisableCloudPermissionTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableCloudPermissionTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableCloudPermissionTemplateResponse) GetBody() *DisableCloudPermissionTemplateResponseBody {
	return s.Body
}

func (s *DisableCloudPermissionTemplateResponse) SetHeaders(v map[string]*string) *DisableCloudPermissionTemplateResponse {
	s.Headers = v
	return s
}

func (s *DisableCloudPermissionTemplateResponse) SetStatusCode(v int32) *DisableCloudPermissionTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableCloudPermissionTemplateResponse) SetBody(v *DisableCloudPermissionTemplateResponseBody) *DisableCloudPermissionTemplateResponse {
	s.Body = v
	return s
}

func (s *DisableCloudPermissionTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
