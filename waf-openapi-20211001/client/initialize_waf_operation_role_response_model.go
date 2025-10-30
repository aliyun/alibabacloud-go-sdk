// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeWafOperationRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitializeWafOperationRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitializeWafOperationRoleResponse
	GetStatusCode() *int32
	SetBody(v *InitializeWafOperationRoleResponseBody) *InitializeWafOperationRoleResponse
	GetBody() *InitializeWafOperationRoleResponseBody
}

type InitializeWafOperationRoleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitializeWafOperationRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitializeWafOperationRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s InitializeWafOperationRoleResponse) GoString() string {
	return s.String()
}

func (s *InitializeWafOperationRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitializeWafOperationRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitializeWafOperationRoleResponse) GetBody() *InitializeWafOperationRoleResponseBody {
	return s.Body
}

func (s *InitializeWafOperationRoleResponse) SetHeaders(v map[string]*string) *InitializeWafOperationRoleResponse {
	s.Headers = v
	return s
}

func (s *InitializeWafOperationRoleResponse) SetStatusCode(v int32) *InitializeWafOperationRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *InitializeWafOperationRoleResponse) SetBody(v *InitializeWafOperationRoleResponseBody) *InitializeWafOperationRoleResponse {
	s.Body = v
	return s
}

func (s *InitializeWafOperationRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
