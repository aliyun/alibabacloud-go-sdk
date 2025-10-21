// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateModelInstanceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateModelInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateModelInstanceResponseBody
	GetMessage() *string
	SetModelInstanceId(v int64) *UpdateModelInstanceResponseBody
	GetModelInstanceId() *int64
	SetRequestId(v string) *UpdateModelInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateModelInstanceResponseBody
	GetSuccess() *bool
}

type UpdateModelInstanceResponseBody struct {
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	ModelInstanceId *int64 `json:"ModelInstanceId,omitempty" xml:"ModelInstanceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateModelInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateModelInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateModelInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateModelInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateModelInstanceResponseBody) GetModelInstanceId() *int64 {
	return s.ModelInstanceId
}

func (s *UpdateModelInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateModelInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateModelInstanceResponseBody) SetCode(v string) *UpdateModelInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateModelInstanceResponseBody) SetHttpStatusCode(v int32) *UpdateModelInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateModelInstanceResponseBody) SetMessage(v string) *UpdateModelInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateModelInstanceResponseBody) SetModelInstanceId(v int64) *UpdateModelInstanceResponseBody {
	s.ModelInstanceId = &v
	return s
}

func (s *UpdateModelInstanceResponseBody) SetRequestId(v string) *UpdateModelInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateModelInstanceResponseBody) SetSuccess(v bool) *UpdateModelInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateModelInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
