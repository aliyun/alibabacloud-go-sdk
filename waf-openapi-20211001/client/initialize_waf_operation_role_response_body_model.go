// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeWafOperationRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InitializeWafOperationRoleResponseBody
	GetRequestId() *string
}

type InitializeWafOperationRoleResponseBody struct {
	// example:
	//
	// 4EC9EA6C-F80A-5D25-A8F7-***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InitializeWafOperationRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitializeWafOperationRoleResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeWafOperationRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitializeWafOperationRoleResponseBody) SetRequestId(v string) *InitializeWafOperationRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitializeWafOperationRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
