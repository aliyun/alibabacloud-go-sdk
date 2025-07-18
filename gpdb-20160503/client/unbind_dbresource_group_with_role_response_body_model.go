// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDBResourceGroupWithRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindDBResourceGroupWithRoleResponseBody
	GetRequestId() *string
}

type UnbindDBResourceGroupWithRoleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindDBResourceGroupWithRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindDBResourceGroupWithRoleResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindDBResourceGroupWithRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindDBResourceGroupWithRoleResponseBody) SetRequestId(v string) *UnbindDBResourceGroupWithRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindDBResourceGroupWithRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
