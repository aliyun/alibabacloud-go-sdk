// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateModelInstanceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateModelInstanceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateModelInstanceResponseBody
	GetMessage() *string
	SetModelInstanceId(v int64) *CreateModelInstanceResponseBody
	GetModelInstanceId() *int64
	SetRequestId(v string) *CreateModelInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateModelInstanceResponseBody
	GetSuccess() *bool
}

type CreateModelInstanceResponseBody struct {
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
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****F68692
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateModelInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateModelInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateModelInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateModelInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateModelInstanceResponseBody) GetModelInstanceId() *int64 {
	return s.ModelInstanceId
}

func (s *CreateModelInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateModelInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateModelInstanceResponseBody) SetCode(v string) *CreateModelInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateModelInstanceResponseBody) SetHttpStatusCode(v int32) *CreateModelInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateModelInstanceResponseBody) SetMessage(v string) *CreateModelInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateModelInstanceResponseBody) SetModelInstanceId(v int64) *CreateModelInstanceResponseBody {
	s.ModelInstanceId = &v
	return s
}

func (s *CreateModelInstanceResponseBody) SetRequestId(v string) *CreateModelInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelInstanceResponseBody) SetSuccess(v bool) *CreateModelInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateModelInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
