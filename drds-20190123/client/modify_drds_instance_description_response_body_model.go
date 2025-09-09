// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDrdsInstanceDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDrdsInstanceDescriptionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyDrdsInstanceDescriptionResponseBody
	GetSuccess() *bool
}

type ModifyDrdsInstanceDescriptionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C44CA24C-C7C4-4C0F-8AC9-1343F2******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDrdsInstanceDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDrdsInstanceDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDrdsInstanceDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDrdsInstanceDescriptionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyDrdsInstanceDescriptionResponseBody) SetRequestId(v string) *ModifyDrdsInstanceDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDrdsInstanceDescriptionResponseBody) SetSuccess(v bool) *ModifyDrdsInstanceDescriptionResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDrdsInstanceDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
