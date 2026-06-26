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
	// Asynchronous task ID. Enable asynchronous tasks beforehand.
	//
	// > When using the SDK for invocation, set a business-related ID. This helps with subsequent operations on the execution. For example, a video processing function can use the video filename as the invocation ID. Use this ID to check if the video processing is complete or to stop it. The ID naming convention must start with an English letter (uppercase or lowercase) or an underscore (_). It can contain English letters (uppercase or lowercase), digits (0-9), underscores (_), and hyphens (-). The ID cannot exceed 128 characters. If you do not set an ID for asynchronous invocation, the system automatically generates one.
	//
	// example:
	//
	// test-id
	XFcAsyncTaskId *string `json:"x-fc-async-task-id,omitempty" xml:"x-fc-async-task-id,omitempty"`
	// Function invocation type. Sync or Async.
	//
	// example:
	//
	// Sync
	XFcInvocationType *string `json:"x-fc-invocation-type,omitempty" xml:"x-fc-invocation-type,omitempty"`
	// Log type returned by function invocation. None or Tail.
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
