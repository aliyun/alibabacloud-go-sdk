// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPullTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowRetry(v bool) *PullTaskResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *PullTaskResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *PullTaskResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *PullTaskResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *PullTaskResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *PullTaskResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *PullTaskResponseBody
	GetHttpStatusCode() *int32
	SetModule(v *PullTaskResponseBodyModule) *PullTaskResponseBody
	GetModule() *PullTaskResponseBodyModule
	SetRequestId(v string) *PullTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PullTaskResponseBody
	GetSuccess() *bool
}

type PullTaskResponseBody struct {
	// allowRetry
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// appName
	//
	// example:
	//
	// voldemort-aliyun-com
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamicCode
	//
	// example:
	//
	// 1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamicMessage
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// errorCode
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// errorMsg
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// httpStatusCode
	//
	// example:
	//
	// 200
	HttpStatusCode *int32                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Module         *PullTaskResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// 9831C9A6-3423-52C2-B0E5-5AE01D92C097
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PullTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PullTaskResponseBody) GoString() string {
	return s.String()
}

func (s *PullTaskResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *PullTaskResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *PullTaskResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *PullTaskResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *PullTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *PullTaskResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *PullTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PullTaskResponseBody) GetModule() *PullTaskResponseBodyModule {
	return s.Module
}

func (s *PullTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PullTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PullTaskResponseBody) SetAllowRetry(v bool) *PullTaskResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *PullTaskResponseBody) SetAppName(v string) *PullTaskResponseBody {
	s.AppName = &v
	return s
}

func (s *PullTaskResponseBody) SetDynamicCode(v string) *PullTaskResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *PullTaskResponseBody) SetDynamicMessage(v string) *PullTaskResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *PullTaskResponseBody) SetErrorCode(v string) *PullTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PullTaskResponseBody) SetErrorMsg(v string) *PullTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *PullTaskResponseBody) SetHttpStatusCode(v int32) *PullTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PullTaskResponseBody) SetModule(v *PullTaskResponseBodyModule) *PullTaskResponseBody {
	s.Module = v
	return s
}

func (s *PullTaskResponseBody) SetRequestId(v string) *PullTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *PullTaskResponseBody) SetSuccess(v bool) *PullTaskResponseBody {
	s.Success = &v
	return s
}

func (s *PullTaskResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PullTaskResponseBodyModule struct {
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// 1649470201045
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 1
	OutTaskId    *string `json:"OutTaskId,omitempty" xml:"OutTaskId,omitempty"`
	PrincipalKey *string `json:"PrincipalKey,omitempty" xml:"PrincipalKey,omitempty"`
	// example:
	//
	// {\\"result\\":\\"SUCCESS\\",\\"message\\":\\"null\\",\\"taskId\\":\\"d8800bab-88b6-4c60-9e4f-ed38dbbdd9b3\\"}
	TaskData *string `json:"TaskData,omitempty" xml:"TaskData,omitempty"`
	// example:
	//
	// 704614
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// PATENT_QUERY
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s PullTaskResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s PullTaskResponseBodyModule) GoString() string {
	return s.String()
}

func (s *PullTaskResponseBodyModule) GetBizCode() *string {
	return s.BizCode
}

func (s *PullTaskResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *PullTaskResponseBodyModule) GetOutTaskId() *string {
	return s.OutTaskId
}

func (s *PullTaskResponseBodyModule) GetPrincipalKey() *string {
	return s.PrincipalKey
}

func (s *PullTaskResponseBodyModule) GetTaskData() *string {
	return s.TaskData
}

func (s *PullTaskResponseBodyModule) GetTaskId() *string {
	return s.TaskId
}

func (s *PullTaskResponseBodyModule) GetTaskType() *string {
	return s.TaskType
}

func (s *PullTaskResponseBodyModule) SetBizCode(v string) *PullTaskResponseBodyModule {
	s.BizCode = &v
	return s
}

func (s *PullTaskResponseBodyModule) SetOrderId(v string) *PullTaskResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *PullTaskResponseBodyModule) SetOutTaskId(v string) *PullTaskResponseBodyModule {
	s.OutTaskId = &v
	return s
}

func (s *PullTaskResponseBodyModule) SetPrincipalKey(v string) *PullTaskResponseBodyModule {
	s.PrincipalKey = &v
	return s
}

func (s *PullTaskResponseBodyModule) SetTaskData(v string) *PullTaskResponseBodyModule {
	s.TaskData = &v
	return s
}

func (s *PullTaskResponseBodyModule) SetTaskId(v string) *PullTaskResponseBodyModule {
	s.TaskId = &v
	return s
}

func (s *PullTaskResponseBodyModule) SetTaskType(v string) *PullTaskResponseBodyModule {
	s.TaskType = &v
	return s
}

func (s *PullTaskResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
