// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudAccountRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCloudAccountRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCloudAccountRoleResponse
	GetStatusCode() *int32
	SetBody(v *CreateCloudAccountRoleResponseBody) *CreateCloudAccountRoleResponse
	GetBody() *CreateCloudAccountRoleResponseBody
}

type CreateCloudAccountRoleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCloudAccountRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCloudAccountRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudAccountRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudAccountRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCloudAccountRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCloudAccountRoleResponse) GetBody() *CreateCloudAccountRoleResponseBody {
	return s.Body
}

func (s *CreateCloudAccountRoleResponse) SetHeaders(v map[string]*string) *CreateCloudAccountRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateCloudAccountRoleResponse) SetStatusCode(v int32) *CreateCloudAccountRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCloudAccountRoleResponse) SetBody(v *CreateCloudAccountRoleResponseBody) *CreateCloudAccountRoleResponse {
	s.Body = v
	return s
}

func (s *CreateCloudAccountRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
