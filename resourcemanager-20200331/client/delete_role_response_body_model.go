// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRoleResponseBody
	GetRequestId() *string
}

type DeleteRoleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 898FAB24-7509-43EE-A287-086FE4C44394
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRoleResponseBody) SetRequestId(v string) *DeleteRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
