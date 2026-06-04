// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCustomAttributeResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateCustomAttributeResponseBody
	GetResult() *bool
	SetSuccess(v bool) *UpdateCustomAttributeResponseBody
	GetSuccess() *bool
}

type UpdateCustomAttributeResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 6B56C4A2-C7F3-52AF-8417-6DFF7447011B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCustomAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomAttributeResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateCustomAttributeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCustomAttributeResponseBody) SetRequestId(v string) *UpdateCustomAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomAttributeResponseBody) SetResult(v bool) *UpdateCustomAttributeResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateCustomAttributeResponseBody) SetSuccess(v bool) *UpdateCustomAttributeResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCustomAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
