// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTerminalCommandsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListTerminalCommandsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTerminalCommandsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTerminalCommandsResponseBody
	GetRequestId() *string
	SetTerminalCommandList(v []*ListTerminalCommandsResponseBodyTerminalCommandList) *ListTerminalCommandsResponseBody
	GetTerminalCommandList() []*ListTerminalCommandsResponseBodyTerminalCommandList
	SetTotalCount(v int32) *ListTerminalCommandsResponseBody
	GetTotalCount() *int32
}

type ListTerminalCommandsResponseBody struct {
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId           *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TerminalCommandList []*ListTerminalCommandsResponseBodyTerminalCommandList `json:"TerminalCommandList,omitempty" xml:"TerminalCommandList,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTerminalCommandsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTerminalCommandsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTerminalCommandsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTerminalCommandsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTerminalCommandsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTerminalCommandsResponseBody) GetTerminalCommandList() []*ListTerminalCommandsResponseBodyTerminalCommandList {
	return s.TerminalCommandList
}

func (s *ListTerminalCommandsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTerminalCommandsResponseBody) SetPageNumber(v int32) *ListTerminalCommandsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTerminalCommandsResponseBody) SetPageSize(v int32) *ListTerminalCommandsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTerminalCommandsResponseBody) SetRequestId(v string) *ListTerminalCommandsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTerminalCommandsResponseBody) SetTerminalCommandList(v []*ListTerminalCommandsResponseBodyTerminalCommandList) *ListTerminalCommandsResponseBody {
	s.TerminalCommandList = v
	return s
}

func (s *ListTerminalCommandsResponseBody) SetTotalCount(v int32) *ListTerminalCommandsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTerminalCommandsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTerminalCommandsResponseBodyTerminalCommandList struct {
	// example:
	//
	// ls
	CommandLine *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	// example:
	//
	// 2024-04-16T03:53:18Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// /root
	ExecutePath *string `json:"ExecutePath,omitempty" xml:"ExecutePath,omitempty"`
	// example:
	//
	// root
	LoginUser *string `json:"LoginUser,omitempty" xml:"LoginUser,omitempty"`
}

func (s ListTerminalCommandsResponseBodyTerminalCommandList) String() string {
	return dara.Prettify(s)
}

func (s ListTerminalCommandsResponseBodyTerminalCommandList) GoString() string {
	return s.String()
}

func (s *ListTerminalCommandsResponseBodyTerminalCommandList) GetCommandLine() *string {
	return s.CommandLine
}

func (s *ListTerminalCommandsResponseBodyTerminalCommandList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListTerminalCommandsResponseBodyTerminalCommandList) GetExecutePath() *string {
	return s.ExecutePath
}

func (s *ListTerminalCommandsResponseBodyTerminalCommandList) GetLoginUser() *string {
	return s.LoginUser
}

func (s *ListTerminalCommandsResponseBodyTerminalCommandList) SetCommandLine(v string) *ListTerminalCommandsResponseBodyTerminalCommandList {
	s.CommandLine = &v
	return s
}

func (s *ListTerminalCommandsResponseBodyTerminalCommandList) SetCreateTime(v string) *ListTerminalCommandsResponseBodyTerminalCommandList {
	s.CreateTime = &v
	return s
}

func (s *ListTerminalCommandsResponseBodyTerminalCommandList) SetExecutePath(v string) *ListTerminalCommandsResponseBodyTerminalCommandList {
	s.ExecutePath = &v
	return s
}

func (s *ListTerminalCommandsResponseBodyTerminalCommandList) SetLoginUser(v string) *ListTerminalCommandsResponseBodyTerminalCommandList {
	s.LoginUser = &v
	return s
}

func (s *ListTerminalCommandsResponseBodyTerminalCommandList) Validate() error {
	return dara.Validate(s)
}
