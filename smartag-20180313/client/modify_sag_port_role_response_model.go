// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagPortRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySagPortRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySagPortRoleResponse
	GetStatusCode() *int32
	SetBody(v *ModifySagPortRoleResponseBody) *ModifySagPortRoleResponse
	GetBody() *ModifySagPortRoleResponseBody
}

type ModifySagPortRoleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySagPortRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySagPortRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySagPortRoleResponse) GoString() string {
	return s.String()
}

func (s *ModifySagPortRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySagPortRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySagPortRoleResponse) GetBody() *ModifySagPortRoleResponseBody {
	return s.Body
}

func (s *ModifySagPortRoleResponse) SetHeaders(v map[string]*string) *ModifySagPortRoleResponse {
	s.Headers = v
	return s
}

func (s *ModifySagPortRoleResponse) SetStatusCode(v int32) *ModifySagPortRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySagPortRoleResponse) SetBody(v *ModifySagPortRoleResponseBody) *ModifySagPortRoleResponse {
	s.Body = v
	return s
}

func (s *ModifySagPortRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
