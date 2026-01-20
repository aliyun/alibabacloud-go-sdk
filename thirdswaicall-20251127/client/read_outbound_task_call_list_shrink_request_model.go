// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadOutboundTaskCallListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrent(v int32) *ReadOutboundTaskCallListShrinkRequest
	GetCurrent() *int32
	SetCustomerNameOrPhone(v string) *ReadOutboundTaskCallListShrinkRequest
	GetCustomerNameOrPhone() *string
	SetDisplayStatusListShrink(v string) *ReadOutboundTaskCallListShrinkRequest
	GetDisplayStatusListShrink() *string
	SetLabelTagsShrink(v string) *ReadOutboundTaskCallListShrinkRequest
	GetLabelTagsShrink() *string
	SetMaxResults(v int32) *ReadOutboundTaskCallListShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ReadOutboundTaskCallListShrinkRequest
	GetNextToken() *string
	SetSize(v int32) *ReadOutboundTaskCallListShrinkRequest
	GetSize() *int32
	SetTaskId(v string) *ReadOutboundTaskCallListShrinkRequest
	GetTaskId() *string
	SetUserId(v string) *ReadOutboundTaskCallListShrinkRequest
	GetUserId() *string
}

type ReadOutboundTaskCallListShrinkRequest struct {
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// 张先生
	CustomerNameOrPhone *string `json:"CustomerNameOrPhone,omitempty" xml:"CustomerNameOrPhone,omitempty"`
	// example:
	//
	// ["1", "2"]
	DisplayStatusListShrink *string `json:"DisplayStatusList,omitempty" xml:"DisplayStatusList,omitempty"`
	// example:
	//
	// ["有意向", "高净值"]
	LabelTagsShrink *string `json:"LabelTags,omitempty" xml:"LabelTags,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// nextToken
	//
	// example:
	//
	// 51CC272E-D879-1B23-B98E-FCFB072D362B
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 1
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ReadOutboundTaskCallListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadOutboundTaskCallListShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReadOutboundTaskCallListShrinkRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ReadOutboundTaskCallListShrinkRequest) GetCustomerNameOrPhone() *string {
	return s.CustomerNameOrPhone
}

func (s *ReadOutboundTaskCallListShrinkRequest) GetDisplayStatusListShrink() *string {
	return s.DisplayStatusListShrink
}

func (s *ReadOutboundTaskCallListShrinkRequest) GetLabelTagsShrink() *string {
	return s.LabelTagsShrink
}

func (s *ReadOutboundTaskCallListShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ReadOutboundTaskCallListShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ReadOutboundTaskCallListShrinkRequest) GetSize() *int32 {
	return s.Size
}

func (s *ReadOutboundTaskCallListShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ReadOutboundTaskCallListShrinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *ReadOutboundTaskCallListShrinkRequest) SetCurrent(v int32) *ReadOutboundTaskCallListShrinkRequest {
	s.Current = &v
	return s
}

func (s *ReadOutboundTaskCallListShrinkRequest) SetCustomerNameOrPhone(v string) *ReadOutboundTaskCallListShrinkRequest {
	s.CustomerNameOrPhone = &v
	return s
}

func (s *ReadOutboundTaskCallListShrinkRequest) SetDisplayStatusListShrink(v string) *ReadOutboundTaskCallListShrinkRequest {
	s.DisplayStatusListShrink = &v
	return s
}

func (s *ReadOutboundTaskCallListShrinkRequest) SetLabelTagsShrink(v string) *ReadOutboundTaskCallListShrinkRequest {
	s.LabelTagsShrink = &v
	return s
}

func (s *ReadOutboundTaskCallListShrinkRequest) SetMaxResults(v int32) *ReadOutboundTaskCallListShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ReadOutboundTaskCallListShrinkRequest) SetNextToken(v string) *ReadOutboundTaskCallListShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ReadOutboundTaskCallListShrinkRequest) SetSize(v int32) *ReadOutboundTaskCallListShrinkRequest {
	s.Size = &v
	return s
}

func (s *ReadOutboundTaskCallListShrinkRequest) SetTaskId(v string) *ReadOutboundTaskCallListShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *ReadOutboundTaskCallListShrinkRequest) SetUserId(v string) *ReadOutboundTaskCallListShrinkRequest {
	s.UserId = &v
	return s
}

func (s *ReadOutboundTaskCallListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
