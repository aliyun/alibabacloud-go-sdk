// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindDBResourceGroupWithRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindDBResourceGroupWithRoleResponseBody
	GetRequestId() *string
}

type BindDBResourceGroupWithRoleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindDBResourceGroupWithRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindDBResourceGroupWithRoleResponseBody) GoString() string {
	return s.String()
}

func (s *BindDBResourceGroupWithRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindDBResourceGroupWithRoleResponseBody) SetRequestId(v string) *BindDBResourceGroupWithRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindDBResourceGroupWithRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
