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
	// The ID of the asynchronous task. You must enable the asynchronous task feature in advance.
	//
	// >  If you use an SDK to invoke a function, we recommend that you specify a business-related ID to facilitate subsequent operations. For example, a video processing function can use video file names as invocation IDs. This way, you can easily check whether a video is successfully processed or terminated before it is processed. The ID can start only with letters or underscores. An ID can contain *letters, digits (0 - 9), underscores*, and hyphens (-). It can be up to 128 characters in length. If you do not specify the ID of the asynchronous invocation, the system automatically generates an ID.
	//
	// example:
	//
	// test-id
	XFcAsyncTaskId *string `json:"x-fc-async-task-id,omitempty" xml:"x-fc-async-task-id,omitempty"`
	// The type of function invocation. Valid values: Sync and Async.
	//
	// example:
	//
	// Sync
	XFcInvocationType *string `json:"x-fc-invocation-type,omitempty" xml:"x-fc-invocation-type,omitempty"`
	// The log type of function invocation. Valid values: None and Tail.
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
