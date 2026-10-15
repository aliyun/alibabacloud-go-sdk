// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorizationAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAuthorizationAttributesResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateAuthorizationAttributesResponseBody
	GetResult() *bool
	SetSuccess(v bool) *UpdateAuthorizationAttributesResponseBody
	GetSuccess() *bool
}

type UpdateAuthorizationAttributesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0c52dbfa-85de-41e8-9dab-43af8e1952a3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the operation. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAuthorizationAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAuthorizationAttributesResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateAuthorizationAttributesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAuthorizationAttributesResponseBody) SetRequestId(v string) *UpdateAuthorizationAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAuthorizationAttributesResponseBody) SetResult(v bool) *UpdateAuthorizationAttributesResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateAuthorizationAttributesResponseBody) SetSuccess(v bool) *UpdateAuthorizationAttributesResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAuthorizationAttributesResponseBody) Validate() error {
	return dara.Validate(s)
}
