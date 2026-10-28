// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeFunctionHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InvokeFunctionHeaders
	GetCommonHeaders() map[string]*string
	SetXFcAsyncTaskId(v string) *InvokeFunctionHeaders
	GetXFcAsyncTaskId() *string
	SetXFcInvocationType(v string) *InvokeFunctionHeaders
	GetXFcInvocationType() *string
	SetXFcLogType(v string) *InvokeFunctionHeaders
	GetXFcLogType() *string
}

type InvokeFunctionHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders" xml:"commonHeaders"`
	// The ID of the asynchronous task. You must enable asynchronous tasks in advance.
	//
	// > We recommend that you set a business-related ID when you use the SDK to invoke the function, which facilitates subsequent operations on the execution. For example, a video processing function can use the video file name as the invoke ID, through which you can check whether the video processing is complete or stop the video processing. The ID must start with an uppercase or lowercase letter or an underscore (_), and can contain uppercase and lowercase letters, digits (0-9), underscores (_), and hyphens (-). It can be up to 128 characters in length. If you do not specify an ID for the asynchronous invoke, the system automatically generates one.
	//
	// example:
	//
	// test-id
	XFcAsyncTaskId *string `json:"x-fc-async-task-id,omitempty" xml:"x-fc-async-task-id,omitempty"`
	// The function invoke type. Valid values: Sync and Async.
	//
	// example:
	//
	// Sync
	XFcInvocationType *string `json:"x-fc-invocation-type,omitempty" xml:"x-fc-invocation-type,omitempty"`
	// The log type returned for function invoke. Valid values: None and Tail.
	//
	// example:
	//
	// Tail
	XFcLogType *string `json:"x-fc-log-type,omitempty" xml:"x-fc-log-type,omitempty"`
}

func (s InvokeFunctionHeaders) String() string {
	return dara.Prettify(s)
}

func (s InvokeFunctionHeaders) GoString() string {
	return s.String()
}

func (s *InvokeFunctionHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InvokeFunctionHeaders) GetXFcAsyncTaskId() *string {
	return s.XFcAsyncTaskId
}

func (s *InvokeFunctionHeaders) GetXFcInvocationType() *string {
	return s.XFcInvocationType
}

func (s *InvokeFunctionHeaders) GetXFcLogType() *string {
	return s.XFcLogType
}

func (s *InvokeFunctionHeaders) SetCommonHeaders(v map[string]*string) *InvokeFunctionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvokeFunctionHeaders) SetXFcAsyncTaskId(v string) *InvokeFunctionHeaders {
	s.XFcAsyncTaskId = &v
	return s
}

func (s *InvokeFunctionHeaders) SetXFcInvocationType(v string) *InvokeFunctionHeaders {
	s.XFcInvocationType = &v
	return s
}

func (s *InvokeFunctionHeaders) SetXFcLogType(v string) *InvokeFunctionHeaders {
	s.XFcLogType = &v
	return s
}

func (s *InvokeFunctionHeaders) Validate() error {
	return dara.Validate(s)
}
