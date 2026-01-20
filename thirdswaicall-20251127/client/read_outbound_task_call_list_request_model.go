// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadOutboundTaskCallListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrent(v int32) *ReadOutboundTaskCallListRequest
	GetCurrent() *int32
	SetCustomerNameOrPhone(v string) *ReadOutboundTaskCallListRequest
	GetCustomerNameOrPhone() *string
	SetDisplayStatusList(v []*string) *ReadOutboundTaskCallListRequest
	GetDisplayStatusList() []*string
	SetLabelTags(v []*string) *ReadOutboundTaskCallListRequest
	GetLabelTags() []*string
	SetMaxResults(v int32) *ReadOutboundTaskCallListRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ReadOutboundTaskCallListRequest
	GetNextToken() *string
	SetSize(v int32) *ReadOutboundTaskCallListRequest
	GetSize() *int32
	SetTaskId(v string) *ReadOutboundTaskCallListRequest
	GetTaskId() *string
	SetUserId(v string) *ReadOutboundTaskCallListRequest
	GetUserId() *string
}

type ReadOutboundTaskCallListRequest struct {
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
	DisplayStatusList []*string `json:"DisplayStatusList,omitempty" xml:"DisplayStatusList,omitempty" type:"Repeated"`
	// example:
	//
	// ["有意向", "高净值"]
	LabelTags []*string `json:"LabelTags,omitempty" xml:"LabelTags,omitempty" type:"Repeated"`
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

func (s ReadOutboundTaskCallListRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadOutboundTaskCallListRequest) GoString() string {
	return s.String()
}

func (s *ReadOutboundTaskCallListRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ReadOutboundTaskCallListRequest) GetCustomerNameOrPhone() *string {
	return s.CustomerNameOrPhone
}

func (s *ReadOutboundTaskCallListRequest) GetDisplayStatusList() []*string {
	return s.DisplayStatusList
}

func (s *ReadOutboundTaskCallListRequest) GetLabelTags() []*string {
	return s.LabelTags
}

func (s *ReadOutboundTaskCallListRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ReadOutboundTaskCallListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ReadOutboundTaskCallListRequest) GetSize() *int32 {
	return s.Size
}

func (s *ReadOutboundTaskCallListRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ReadOutboundTaskCallListRequest) GetUserId() *string {
	return s.UserId
}

func (s *ReadOutboundTaskCallListRequest) SetCurrent(v int32) *ReadOutboundTaskCallListRequest {
	s.Current = &v
	return s
}

func (s *ReadOutboundTaskCallListRequest) SetCustomerNameOrPhone(v string) *ReadOutboundTaskCallListRequest {
	s.CustomerNameOrPhone = &v
	return s
}

func (s *ReadOutboundTaskCallListRequest) SetDisplayStatusList(v []*string) *ReadOutboundTaskCallListRequest {
	s.DisplayStatusList = v
	return s
}

func (s *ReadOutboundTaskCallListRequest) SetLabelTags(v []*string) *ReadOutboundTaskCallListRequest {
	s.LabelTags = v
	return s
}

func (s *ReadOutboundTaskCallListRequest) SetMaxResults(v int32) *ReadOutboundTaskCallListRequest {
	s.MaxResults = &v
	return s
}

func (s *ReadOutboundTaskCallListRequest) SetNextToken(v string) *ReadOutboundTaskCallListRequest {
	s.NextToken = &v
	return s
}

func (s *ReadOutboundTaskCallListRequest) SetSize(v int32) *ReadOutboundTaskCallListRequest {
	s.Size = &v
	return s
}

func (s *ReadOutboundTaskCallListRequest) SetTaskId(v string) *ReadOutboundTaskCallListRequest {
	s.TaskId = &v
	return s
}

func (s *ReadOutboundTaskCallListRequest) SetUserId(v string) *ReadOutboundTaskCallListRequest {
	s.UserId = &v
	return s
}

func (s *ReadOutboundTaskCallListRequest) Validate() error {
	return dara.Validate(s)
}
