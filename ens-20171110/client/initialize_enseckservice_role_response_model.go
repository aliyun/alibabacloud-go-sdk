// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeENSECKServiceRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitializeENSECKServiceRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitializeENSECKServiceRoleResponse
	GetStatusCode() *int32
	SetBody(v *InitializeENSECKServiceRoleResponseBody) *InitializeENSECKServiceRoleResponse
	GetBody() *InitializeENSECKServiceRoleResponseBody
}

type InitializeENSECKServiceRoleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitializeENSECKServiceRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitializeENSECKServiceRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s InitializeENSECKServiceRoleResponse) GoString() string {
	return s.String()
}

func (s *InitializeENSECKServiceRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitializeENSECKServiceRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitializeENSECKServiceRoleResponse) GetBody() *InitializeENSECKServiceRoleResponseBody {
	return s.Body
}

func (s *InitializeENSECKServiceRoleResponse) SetHeaders(v map[string]*string) *InitializeENSECKServiceRoleResponse {
	s.Headers = v
	return s
}

func (s *InitializeENSECKServiceRoleResponse) SetStatusCode(v int32) *InitializeENSECKServiceRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *InitializeENSECKServiceRoleResponse) SetBody(v *InitializeENSECKServiceRoleResponseBody) *InitializeENSECKServiceRoleResponse {
	s.Body = v
	return s
}

func (s *InitializeENSECKServiceRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
