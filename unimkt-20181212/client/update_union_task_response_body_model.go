// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUnionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *UpdateUnionTaskResponseBody
	GetCode() *int64
	SetData(v int64) *UpdateUnionTaskResponseBody
	GetData() *int64
	SetErrorMsg(v string) *UpdateUnionTaskResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *UpdateUnionTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateUnionTaskResponseBody
	GetSuccess() *bool
}

type UpdateUnionTaskResponseBody struct {
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1687154040871094
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// error check permissions
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6CBD8D0F-D84A-5545-817C-16C7430FD612
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateUnionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUnionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUnionTaskResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *UpdateUnionTaskResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateUnionTaskResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *UpdateUnionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUnionTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateUnionTaskResponseBody) SetCode(v int64) *UpdateUnionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateUnionTaskResponseBody) SetData(v int64) *UpdateUnionTaskResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateUnionTaskResponseBody) SetErrorMsg(v string) *UpdateUnionTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateUnionTaskResponseBody) SetRequestId(v string) *UpdateUnionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUnionTaskResponseBody) SetSuccess(v bool) *UpdateUnionTaskResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateUnionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
