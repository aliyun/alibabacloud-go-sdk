// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumSymbolicStackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetColumn(v int64) *GetRumSymbolicStackResponseBody
	GetColumn() *int64
	SetLine(v int64) *GetRumSymbolicStackResponseBody
	GetLine() *int64
	SetLines(v []*string) *GetRumSymbolicStackResponseBody
	GetLines() []*string
	SetName(v string) *GetRumSymbolicStackResponseBody
	GetName() *string
	SetRequestId(v string) *GetRumSymbolicStackResponseBody
	GetRequestId() *string
	SetSource(v string) *GetRumSymbolicStackResponseBody
	GetSource() *string
	SetThreadId(v string) *GetRumSymbolicStackResponseBody
	GetThreadId() *string
	SetTraceInfo(v []*string) *GetRumSymbolicStackResponseBody
	GetTraceInfo() []*string
	SetTraceMarkLine(v int64) *GetRumSymbolicStackResponseBody
	GetTraceMarkLine() *int64
}

type GetRumSymbolicStackResponseBody struct {
	// The column number restored by the JS SourceMap.
	//
	// example:
	//
	// 10
	Column *int64 `json:"column,omitempty" xml:"column,omitempty"`
	// The line number restored by the JS SourceMap.
	//
	// example:
	//
	// 42
	Line *int64 `json:"line,omitempty" xml:"line,omitempty"`
	// The list of symbolicated stack lines.
	//
	// example:
	//
	// ["\\tat com.example.UserService.getUser(UserService.java:25)"]
	Lines []*string `json:"lines,omitempty" xml:"lines,omitempty" type:"Repeated"`
	// The symbol name restored by the JS SourceMap.
	//
	// example:
	//
	// handleClick
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The source file path restored by the JS SourceMap.
	//
	// example:
	//
	// src/app.js
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The ID of the exception thread.
	//
	// example:
	//
	// main-thread
	ThreadId *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
	// The list of source code context lines.
	//
	// example:
	//
	// ["line1"]
	TraceInfo []*string `json:"traceInfo,omitempty" xml:"traceInfo,omitempty" type:"Repeated"`
	// The highlighted line number of the error line in traceInfo.
	//
	// example:
	//
	// 1
	TraceMarkLine *int64 `json:"traceMarkLine,omitempty" xml:"traceMarkLine,omitempty"`
}

func (s GetRumSymbolicStackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRumSymbolicStackResponseBody) GoString() string {
	return s.String()
}

func (s *GetRumSymbolicStackResponseBody) GetColumn() *int64 {
	return s.Column
}

func (s *GetRumSymbolicStackResponseBody) GetLine() *int64 {
	return s.Line
}

func (s *GetRumSymbolicStackResponseBody) GetLines() []*string {
	return s.Lines
}

func (s *GetRumSymbolicStackResponseBody) GetName() *string {
	return s.Name
}

func (s *GetRumSymbolicStackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRumSymbolicStackResponseBody) GetSource() *string {
	return s.Source
}

func (s *GetRumSymbolicStackResponseBody) GetThreadId() *string {
	return s.ThreadId
}

func (s *GetRumSymbolicStackResponseBody) GetTraceInfo() []*string {
	return s.TraceInfo
}

func (s *GetRumSymbolicStackResponseBody) GetTraceMarkLine() *int64 {
	return s.TraceMarkLine
}

func (s *GetRumSymbolicStackResponseBody) SetColumn(v int64) *GetRumSymbolicStackResponseBody {
	s.Column = &v
	return s
}

func (s *GetRumSymbolicStackResponseBody) SetLine(v int64) *GetRumSymbolicStackResponseBody {
	s.Line = &v
	return s
}

func (s *GetRumSymbolicStackResponseBody) SetLines(v []*string) *GetRumSymbolicStackResponseBody {
	s.Lines = v
	return s
}

func (s *GetRumSymbolicStackResponseBody) SetName(v string) *GetRumSymbolicStackResponseBody {
	s.Name = &v
	return s
}

func (s *GetRumSymbolicStackResponseBody) SetRequestId(v string) *GetRumSymbolicStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRumSymbolicStackResponseBody) SetSource(v string) *GetRumSymbolicStackResponseBody {
	s.Source = &v
	return s
}

func (s *GetRumSymbolicStackResponseBody) SetThreadId(v string) *GetRumSymbolicStackResponseBody {
	s.ThreadId = &v
	return s
}

func (s *GetRumSymbolicStackResponseBody) SetTraceInfo(v []*string) *GetRumSymbolicStackResponseBody {
	s.TraceInfo = v
	return s
}

func (s *GetRumSymbolicStackResponseBody) SetTraceMarkLine(v int64) *GetRumSymbolicStackResponseBody {
	s.TraceMarkLine = &v
	return s
}

func (s *GetRumSymbolicStackResponseBody) Validate() error {
	return dara.Validate(s)
}
