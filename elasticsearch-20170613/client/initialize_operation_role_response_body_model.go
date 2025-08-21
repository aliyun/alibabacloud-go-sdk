// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeOperationRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InitializeOperationRoleResponseBody
	GetRequestId() *string
	SetResult(v bool) *InitializeOperationRoleResponseBody
	GetResult() *bool
}

type InitializeOperationRoleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 29101430-4797-4D1D-96C3-9FCBCCA8F845
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the service-linked role is created. Valid values:
	//
	// 	- true: The service-linked role is created.
	//
	// 	- false: The service-linked role fails to be created.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s InitializeOperationRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitializeOperationRoleResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeOperationRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitializeOperationRoleResponseBody) GetResult() *bool {
	return s.Result
}

func (s *InitializeOperationRoleResponseBody) SetRequestId(v string) *InitializeOperationRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitializeOperationRoleResponseBody) SetResult(v bool) *InitializeOperationRoleResponseBody {
	s.Result = &v
	return s
}

func (s *InitializeOperationRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
