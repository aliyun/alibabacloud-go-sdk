// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudPermissionTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCloudPermissionTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCloudPermissionTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateCloudPermissionTemplateResponseBody) *CreateCloudPermissionTemplateResponse
	GetBody() *CreateCloudPermissionTemplateResponseBody
}

type CreateCloudPermissionTemplateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCloudPermissionTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCloudPermissionTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudPermissionTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudPermissionTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCloudPermissionTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCloudPermissionTemplateResponse) GetBody() *CreateCloudPermissionTemplateResponseBody {
	return s.Body
}

func (s *CreateCloudPermissionTemplateResponse) SetHeaders(v map[string]*string) *CreateCloudPermissionTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateCloudPermissionTemplateResponse) SetStatusCode(v int32) *CreateCloudPermissionTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCloudPermissionTemplateResponse) SetBody(v *CreateCloudPermissionTemplateResponseBody) *CreateCloudPermissionTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateCloudPermissionTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
