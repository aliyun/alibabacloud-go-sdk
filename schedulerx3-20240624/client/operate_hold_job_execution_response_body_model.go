// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateHoldJobExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OperateHoldJobExecutionResponseBody
	GetCode() *int32
	SetMessage(v string) *OperateHoldJobExecutionResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateHoldJobExecutionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateHoldJobExecutionResponseBody
	GetSuccess() *bool
}

type OperateHoldJobExecutionResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// not support query script history, please upgrade engine version to 2.2.2+
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2ECA6FC9-7557-5576-AF5F-FC3E7BCC9C21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateHoldJobExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateHoldJobExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *OperateHoldJobExecutionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OperateHoldJobExecutionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateHoldJobExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateHoldJobExecutionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateHoldJobExecutionResponseBody) SetCode(v int32) *OperateHoldJobExecutionResponseBody {
	s.Code = &v
	return s
}

func (s *OperateHoldJobExecutionResponseBody) SetMessage(v string) *OperateHoldJobExecutionResponseBody {
	s.Message = &v
	return s
}

func (s *OperateHoldJobExecutionResponseBody) SetRequestId(v string) *OperateHoldJobExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateHoldJobExecutionResponseBody) SetSuccess(v bool) *OperateHoldJobExecutionResponseBody {
	s.Success = &v
	return s
}

func (s *OperateHoldJobExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
