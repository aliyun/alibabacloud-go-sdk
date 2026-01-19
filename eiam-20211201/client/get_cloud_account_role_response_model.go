// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudAccountRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCloudAccountRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCloudAccountRoleResponse
	GetStatusCode() *int32
	SetBody(v *GetCloudAccountRoleResponseBody) *GetCloudAccountRoleResponse
	GetBody() *GetCloudAccountRoleResponseBody
}

type GetCloudAccountRoleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCloudAccountRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCloudAccountRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAccountRoleResponse) GoString() string {
	return s.String()
}

func (s *GetCloudAccountRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCloudAccountRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCloudAccountRoleResponse) GetBody() *GetCloudAccountRoleResponseBody {
	return s.Body
}

func (s *GetCloudAccountRoleResponse) SetHeaders(v map[string]*string) *GetCloudAccountRoleResponse {
	s.Headers = v
	return s
}

func (s *GetCloudAccountRoleResponse) SetStatusCode(v int32) *GetCloudAccountRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCloudAccountRoleResponse) SetBody(v *GetCloudAccountRoleResponseBody) *GetCloudAccountRoleResponse {
	s.Body = v
	return s
}

func (s *GetCloudAccountRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
