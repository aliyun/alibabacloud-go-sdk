// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateUnholdJobExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OperateUnholdJobExecutionResponseBody
	GetCode() *int32
	SetMessage(v string) *OperateUnholdJobExecutionResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateUnholdJobExecutionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateUnholdJobExecutionResponseBody
	GetSuccess() *bool
}

type OperateUnholdJobExecutionResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D0DE9C33-992A-580B-89C4-B609A292748D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateUnholdJobExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateUnholdJobExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *OperateUnholdJobExecutionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OperateUnholdJobExecutionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateUnholdJobExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateUnholdJobExecutionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateUnholdJobExecutionResponseBody) SetCode(v int32) *OperateUnholdJobExecutionResponseBody {
	s.Code = &v
	return s
}

func (s *OperateUnholdJobExecutionResponseBody) SetMessage(v string) *OperateUnholdJobExecutionResponseBody {
	s.Message = &v
	return s
}

func (s *OperateUnholdJobExecutionResponseBody) SetRequestId(v string) *OperateUnholdJobExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateUnholdJobExecutionResponseBody) SetSuccess(v bool) *OperateUnholdJobExecutionResponseBody {
	s.Success = &v
	return s
}

func (s *OperateUnholdJobExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
