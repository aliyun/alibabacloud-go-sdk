// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetStackResponseBody
	GetRequestId() *string
	SetStackInfo(v []*GetStackResponseBodyStackInfo) *GetStackResponseBody
	GetStackInfo() []*GetStackResponseBodyStackInfo
}

type GetStackResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B34C3A43-A901-5F94-9DAD-758CE4******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The method stack details.
	StackInfo []*GetStackResponseBodyStackInfo `json:"StackInfo,omitempty" xml:"StackInfo,omitempty" type:"Repeated"`
}

func (s GetStackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStackResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStackResponseBody) GetStackInfo() []*GetStackResponseBodyStackInfo {
	return s.StackInfo
}

func (s *GetStackResponseBody) SetRequestId(v string) *GetStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackResponseBody) SetStackInfo(v []*GetStackResponseBodyStackInfo) *GetStackResponseBody {
	s.StackInfo = v
	return s
}

func (s *GetStackResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetStackResponseBodyStackInfo struct {
	// The name of the operation.
	//
	// example:
	//
	// Tomcat Servlet Process
	Api *string `json:"Api,omitempty" xml:"Api,omitempty"`
	// The number of times the method is called.
	//
	// example:
	//
	// 1
	CallCount *string `json:"CallCount,omitempty" xml:"CallCount,omitempty"`
	// The duration. Unit: milliseconds.
	//
	// example:
	//
	// 32
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The error message.
	//
	// example:
	//
	// java.lang.NullPointerException
	Exception *string `json:"Exception,omitempty" xml:"Exception,omitempty"`
	// The information about the array object.
	ExtInfo *GetStackResponseBodyStackInfoExtInfo `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty" type:"Struct"`
	// The number of rows in the method stack information.
	//
	// example:
	//
	// 34
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// The ID of the RPC mode.
	//
	// example:
	//
	// 0.1
	RpcId *string `json:"RpcId,omitempty" xml:"RpcId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// /com/test
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The start time of the call method.
	//
	// example:
	//
	// 1653555396
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetStackResponseBodyStackInfo) String() string {
	return dara.Prettify(s)
}

func (s GetStackResponseBodyStackInfo) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyStackInfo) GetApi() *string {
	return s.Api
}

func (s *GetStackResponseBodyStackInfo) GetCallCount() *string {
	return s.CallCount
}

func (s *GetStackResponseBodyStackInfo) GetDuration() *int64 {
	return s.Duration
}

func (s *GetStackResponseBodyStackInfo) GetException() *string {
	return s.Exception
}

func (s *GetStackResponseBodyStackInfo) GetExtInfo() *GetStackResponseBodyStackInfoExtInfo {
	return s.ExtInfo
}

func (s *GetStackResponseBodyStackInfo) GetLine() *string {
	return s.Line
}

func (s *GetStackResponseBodyStackInfo) GetRpcId() *string {
	return s.RpcId
}

func (s *GetStackResponseBodyStackInfo) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetStackResponseBodyStackInfo) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetStackResponseBodyStackInfo) SetApi(v string) *GetStackResponseBodyStackInfo {
	s.Api = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetCallCount(v string) *GetStackResponseBodyStackInfo {
	s.CallCount = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetDuration(v int64) *GetStackResponseBodyStackInfo {
	s.Duration = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetException(v string) *GetStackResponseBodyStackInfo {
	s.Exception = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetExtInfo(v *GetStackResponseBodyStackInfoExtInfo) *GetStackResponseBodyStackInfo {
	s.ExtInfo = v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetLine(v string) *GetStackResponseBodyStackInfo {
	s.Line = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetRpcId(v string) *GetStackResponseBodyStackInfo {
	s.RpcId = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetServiceName(v string) *GetStackResponseBodyStackInfo {
	s.ServiceName = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetStartTime(v int64) *GetStackResponseBodyStackInfo {
	s.StartTime = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) Validate() error {
	return dara.Validate(s)
}

type GetStackResponseBodyStackInfoExtInfo struct {
	// The content of the custom parameter.
	//
	// example:
	//
	// input=254275&amp;
	Info *string `json:"Info,omitempty" xml:"Info,omitempty"`
	// The type of the custom parameter.
	//
	// example:
	//
	// 41
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetStackResponseBodyStackInfoExtInfo) String() string {
	return dara.Prettify(s)
}

func (s GetStackResponseBodyStackInfoExtInfo) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyStackInfoExtInfo) GetInfo() *string {
	return s.Info
}

func (s *GetStackResponseBodyStackInfoExtInfo) GetType() *string {
	return s.Type
}

func (s *GetStackResponseBodyStackInfoExtInfo) SetInfo(v string) *GetStackResponseBodyStackInfoExtInfo {
	s.Info = &v
	return s
}

func (s *GetStackResponseBodyStackInfoExtInfo) SetType(v string) *GetStackResponseBodyStackInfoExtInfo {
	s.Type = &v
	return s
}

func (s *GetStackResponseBodyStackInfoExtInfo) Validate() error {
	return dara.Validate(s)
}
