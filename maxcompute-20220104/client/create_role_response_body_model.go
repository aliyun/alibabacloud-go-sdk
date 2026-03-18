// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CreateRoleResponseBody
	GetData() *string
	SetRequestId(v string) *CreateRoleResponseBody
	GetRequestId() *string
}

type CreateRoleResponseBody struct {
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRoleResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRoleResponseBody) SetData(v string) *CreateRoleResponseBody {
	s.Data = &v
	return s
}

func (s *CreateRoleResponseBody) SetRequestId(v string) *CreateRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
