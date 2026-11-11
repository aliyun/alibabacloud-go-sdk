// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAccountResourcesForCloudPermissionTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloudAccountResourcesForCloudPermissionTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloudAccountResourcesForCloudPermissionTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody) *ListCloudAccountResourcesForCloudPermissionTemplateResponse
	GetBody() *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody
}

type ListCloudAccountResourcesForCloudPermissionTemplateResponse struct {
	Headers    map[string]*string                                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudAccountResourcesForCloudPermissionTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccountResourcesForCloudPermissionTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponse) GetBody() *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody {
	return s.Body
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponse) SetHeaders(v map[string]*string) *ListCloudAccountResourcesForCloudPermissionTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponse) SetStatusCode(v int32) *ListCloudAccountResourcesForCloudPermissionTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponse) SetBody(v *ListCloudAccountResourcesForCloudPermissionTemplateResponseBody) *ListCloudAccountResourcesForCloudPermissionTemplateResponse {
	s.Body = v
	return s
}

func (s *ListCloudAccountResourcesForCloudPermissionTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
