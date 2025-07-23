// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogoTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateLogoTaskResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *CreateLogoTaskResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *CreateLogoTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateLogoTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *CreateLogoTaskResponseBody
	GetTaskId() *string
}

type CreateLogoTaskResponseBody struct {
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// aliuid:1998006665794443 assumeRole not exist,serviceName:realtimelogpush.dcdnservices.aliyuncs.com
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// c3r127e325at9yd
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateLogoTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLogoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLogoTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateLogoTaskResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CreateLogoTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLogoTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateLogoTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateLogoTaskResponseBody) SetErrorCode(v string) *CreateLogoTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateLogoTaskResponseBody) SetErrorMsg(v string) *CreateLogoTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateLogoTaskResponseBody) SetRequestId(v string) *CreateLogoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLogoTaskResponseBody) SetSuccess(v bool) *CreateLogoTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateLogoTaskResponseBody) SetTaskId(v string) *CreateLogoTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateLogoTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
