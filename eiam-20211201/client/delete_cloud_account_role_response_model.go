// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudAccountRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCloudAccountRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCloudAccountRoleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCloudAccountRoleResponseBody) *DeleteCloudAccountRoleResponse
	GetBody() *DeleteCloudAccountRoleResponseBody
}

type DeleteCloudAccountRoleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCloudAccountRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCloudAccountRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudAccountRoleResponse) GoString() string {
	return s.String()
}

func (s *DeleteCloudAccountRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCloudAccountRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCloudAccountRoleResponse) GetBody() *DeleteCloudAccountRoleResponseBody {
	return s.Body
}

func (s *DeleteCloudAccountRoleResponse) SetHeaders(v map[string]*string) *DeleteCloudAccountRoleResponse {
	s.Headers = v
	return s
}

func (s *DeleteCloudAccountRoleResponse) SetStatusCode(v int32) *DeleteCloudAccountRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCloudAccountRoleResponse) SetBody(v *DeleteCloudAccountRoleResponseBody) *DeleteCloudAccountRoleResponse {
	s.Body = v
	return s
}

func (s *DeleteCloudAccountRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
