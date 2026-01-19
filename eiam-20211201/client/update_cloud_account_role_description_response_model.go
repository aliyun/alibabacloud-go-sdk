// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudAccountRoleDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCloudAccountRoleDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCloudAccountRoleDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCloudAccountRoleDescriptionResponseBody) *UpdateCloudAccountRoleDescriptionResponse
	GetBody() *UpdateCloudAccountRoleDescriptionResponseBody
}

type UpdateCloudAccountRoleDescriptionResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloudAccountRoleDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloudAccountRoleDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudAccountRoleDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudAccountRoleDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCloudAccountRoleDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCloudAccountRoleDescriptionResponse) GetBody() *UpdateCloudAccountRoleDescriptionResponseBody {
	return s.Body
}

func (s *UpdateCloudAccountRoleDescriptionResponse) SetHeaders(v map[string]*string) *UpdateCloudAccountRoleDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudAccountRoleDescriptionResponse) SetStatusCode(v int32) *UpdateCloudAccountRoleDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloudAccountRoleDescriptionResponse) SetBody(v *UpdateCloudAccountRoleDescriptionResponseBody) *UpdateCloudAccountRoleDescriptionResponse {
	s.Body = v
	return s
}

func (s *UpdateCloudAccountRoleDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
