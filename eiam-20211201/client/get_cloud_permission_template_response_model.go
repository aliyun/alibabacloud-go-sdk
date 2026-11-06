// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudPermissionTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCloudPermissionTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCloudPermissionTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetCloudPermissionTemplateResponseBody) *GetCloudPermissionTemplateResponse
	GetBody() *GetCloudPermissionTemplateResponseBody
}

type GetCloudPermissionTemplateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCloudPermissionTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCloudPermissionTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCloudPermissionTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetCloudPermissionTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCloudPermissionTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCloudPermissionTemplateResponse) GetBody() *GetCloudPermissionTemplateResponseBody {
	return s.Body
}

func (s *GetCloudPermissionTemplateResponse) SetHeaders(v map[string]*string) *GetCloudPermissionTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetCloudPermissionTemplateResponse) SetStatusCode(v int32) *GetCloudPermissionTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCloudPermissionTemplateResponse) SetBody(v *GetCloudPermissionTemplateResponseBody) *GetCloudPermissionTemplateResponse {
	s.Body = v
	return s
}

func (s *GetCloudPermissionTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
