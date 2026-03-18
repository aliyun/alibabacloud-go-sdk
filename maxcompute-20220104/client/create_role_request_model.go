// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *CreateRoleRequest
	GetBody() *string
}

type CreateRoleRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateRoleRequest) GetBody() *string {
	return s.Body
}

func (s *CreateRoleRequest) SetBody(v string) *CreateRoleRequest {
	s.Body = &v
	return s
}

func (s *CreateRoleRequest) Validate() error {
	return dara.Validate(s)
}
