// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateAppGroupResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateAppGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAppGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAppGroupResponseBody
	GetSuccess() *bool
}

type UpdateAppGroupResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// app is not existed, groupId=xxxx, namesapce=xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAppGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateAppGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAppGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAppGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAppGroupResponseBody) SetCode(v int32) *UpdateAppGroupResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAppGroupResponseBody) SetMessage(v string) *UpdateAppGroupResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAppGroupResponseBody) SetRequestId(v string) *UpdateAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppGroupResponseBody) SetSuccess(v bool) *UpdateAppGroupResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAppGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
