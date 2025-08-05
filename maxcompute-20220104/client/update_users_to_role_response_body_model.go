// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUsersToRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UpdateUsersToRoleResponseBody
	GetData() *string
	SetRequestId(v string) *UpdateUsersToRoleResponseBody
	GetRequestId() *string
}

type UpdateUsersToRoleResponseBody struct {
	// The data returned.
	//
	// example:
	//
	// success
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0a032a1317254153012687347ef4ee
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateUsersToRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUsersToRoleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUsersToRoleResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateUsersToRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUsersToRoleResponseBody) SetData(v string) *UpdateUsersToRoleResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateUsersToRoleResponseBody) SetRequestId(v string) *UpdateUsersToRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUsersToRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
