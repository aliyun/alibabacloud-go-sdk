// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventPageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeEventPageListRequest
	GetLang() *string
	SetCreateType(v string) *DescribeEventPageListRequest
	GetCreateType() *string
	SetCurrentPage(v int32) *DescribeEventPageListRequest
	GetCurrentPage() *int32
	SetEventCode(v string) *DescribeEventPageListRequest
	GetEventCode() *string
	SetEventName(v string) *DescribeEventPageListRequest
	GetEventName() *string
	SetEventStatus(v string) *DescribeEventPageListRequest
	GetEventStatus() *string
	SetPageSize(v int32) *DescribeEventPageListRequest
	GetPageSize() *int32
	SetRegId(v string) *DescribeEventPageListRequest
	GetRegId() *string
}

type DescribeEventPageListRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Creation type
	//
	// example:
	//
	// NORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Event code
	//
	// example:
	//
	// de_aikqxy3122
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Event name.
	//
	// example:
	//
	// 注册风险旁路
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Event status.
	//
	// example:
	//
	// ONLINE
	EventStatus *string `json:"eventStatus,omitempty" xml:"eventStatus,omitempty"`
	// Page size, with a default value of 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeEventPageListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventPageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventPageListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeEventPageListRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *DescribeEventPageListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeEventPageListRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeEventPageListRequest) GetEventName() *string {
	return s.EventName
}

func (s *DescribeEventPageListRequest) GetEventStatus() *string {
	return s.EventStatus
}

func (s *DescribeEventPageListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEventPageListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeEventPageListRequest) SetLang(v string) *DescribeEventPageListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventPageListRequest) SetCreateType(v string) *DescribeEventPageListRequest {
	s.CreateType = &v
	return s
}

func (s *DescribeEventPageListRequest) SetCurrentPage(v int32) *DescribeEventPageListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeEventPageListRequest) SetEventCode(v string) *DescribeEventPageListRequest {
	s.EventCode = &v
	return s
}

func (s *DescribeEventPageListRequest) SetEventName(v string) *DescribeEventPageListRequest {
	s.EventName = &v
	return s
}

func (s *DescribeEventPageListRequest) SetEventStatus(v string) *DescribeEventPageListRequest {
	s.EventStatus = &v
	return s
}

func (s *DescribeEventPageListRequest) SetPageSize(v int32) *DescribeEventPageListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEventPageListRequest) SetRegId(v string) *DescribeEventPageListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeEventPageListRequest) Validate() error {
	return dara.Validate(s)
}
