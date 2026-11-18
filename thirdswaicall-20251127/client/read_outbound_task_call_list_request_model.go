// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadOutboundTaskCallListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallEndTimeBegin(v int64) *ReadOutboundTaskCallListRequest
	GetCallEndTimeBegin() *int64
	SetCallEndTimeEnd(v int64) *ReadOutboundTaskCallListRequest
	GetCallEndTimeEnd() *int64
	SetCallStartTimeBegin(v int64) *ReadOutboundTaskCallListRequest
	GetCallStartTimeBegin() *int64
	SetCallStartTimeEnd(v int64) *ReadOutboundTaskCallListRequest
	GetCallStartTimeEnd() *int64
	SetCallerUacAccountId(v string) *ReadOutboundTaskCallListRequest
	GetCallerUacAccountId() *string
	SetCurrent(v int32) *ReadOutboundTaskCallListRequest
	GetCurrent() *int32
	SetCurrentWorkspaceCode(v string) *ReadOutboundTaskCallListRequest
	GetCurrentWorkspaceCode() *string
	SetCurrentWorkspaceId(v string) *ReadOutboundTaskCallListRequest
	GetCurrentWorkspaceId() *string
	SetCustomerNameOrPhone(v string) *ReadOutboundTaskCallListRequest
	GetCustomerNameOrPhone() *string
	SetDisplayStatusList(v []*string) *ReadOutboundTaskCallListRequest
	GetDisplayStatusList() []*string
	SetDurationRangeList(v []*string) *ReadOutboundTaskCallListRequest
	GetDurationRangeList() []*string
	SetLabelTags(v []*string) *ReadOutboundTaskCallListRequest
	GetLabelTags() []*string
	SetMaxResults(v int32) *ReadOutboundTaskCallListRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ReadOutboundTaskCallListRequest
	GetNextToken() *string
	SetSize(v int32) *ReadOutboundTaskCallListRequest
	GetSize() *int32
	SetStructuredLabelTags(v []*ReadOutboundTaskCallListRequestStructuredLabelTags) *ReadOutboundTaskCallListRequest
	GetStructuredLabelTags() []*ReadOutboundTaskCallListRequestStructuredLabelTags
	SetTaskId(v string) *ReadOutboundTaskCallListRequest
	GetTaskId() *string
	SetUserId(v string) *ReadOutboundTaskCallListRequest
	GetUserId() *string
}

type ReadOutboundTaskCallListRequest struct {
	// example:
	//
	// 1735689600000
	CallEndTimeBegin *int64 `json:"CallEndTimeBegin,omitempty" xml:"CallEndTimeBegin,omitempty"`
	// example:
	//
	// 1767225600000
	CallEndTimeEnd *int64 `json:"CallEndTimeEnd,omitempty" xml:"CallEndTimeEnd,omitempty"`
	// example:
	//
	// 1735689600000
	CallStartTimeBegin *int64 `json:"CallStartTimeBegin,omitempty" xml:"CallStartTimeBegin,omitempty"`
	// example:
	//
	// 1767225600000
	CallStartTimeEnd *int64 `json:"CallStartTimeEnd,omitempty" xml:"CallStartTimeEnd,omitempty"`
	// example:
	//
	// abc123***
	CallerUacAccountId *string `json:"CallerUacAccountId,omitempty" xml:"CallerUacAccountId,omitempty"`
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// 1ab92c
	CurrentWorkspaceCode *string `json:"CurrentWorkspaceCode,omitempty" xml:"CurrentWorkspaceCode,omitempty"`
	// example:
	//
	// abc123***
	CurrentWorkspaceId *string `json:"CurrentWorkspaceId,omitempty" xml:"CurrentWorkspaceId,omitempty"`
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
	// ["WITHIN_15_SECONDS", "FROM_1_TO_3_MINUTES"]
	DurationRangeList []*string `json:"DurationRangeList,omitempty" xml:"DurationRangeList,omitempty" type:"Repeated"`
	// example:
	//
	// ["有意向", "高净值"]
	LabelTags []*string `json:"LabelTags,omitempty" xml:"LabelTags,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 51CC272E-D879-1B23-B98E-FCFB072D362B
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 10
	Size                *int32                                                `json:"Size,omitempty" xml:"Size,omitempty"`
	StructuredLabelTags []*ReadOutboundTaskCallListRequestStructuredLabelTags `json:"StructuredLabelTags,omitempty" xml:"StructuredLabelTags,omitempty" type:"Repeated"`
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

func (s *ReadOutboundTaskCallListRequest) GetCallEndTimeBegin() *int64 {
	return s.CallEndTimeBegin
}

func (s *ReadOutboundTaskCallListRequest) GetCallEndTimeEnd() *int64 {
	return s.CallEndTimeEnd
}

func (s *ReadOutboundTaskCallListRequest) GetCallStartTimeBegin() *int64 {
	return s.CallStartTimeBegin
}

func (s *ReadOutboundTaskCallListRequest) GetCallStartTimeEnd() *int64 {
	return s.CallStartTimeEnd
}

func (s *ReadOutboundTaskCallListRequest) GetCallerUacAccountId() *string {
	return s.CallerUacAccountId
}

func (s *ReadOutboundTaskCallListRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ReadOutboundTaskCallListRequest) GetCurrentWorkspaceCode() *string {
	return s.CurrentWorkspaceCode
}

func (s *ReadOutboundTaskCallListRequest) GetCurrentWorkspaceId() *string {
	return s.CurrentWorkspaceId
}

func (s *ReadOutboundTaskCallListRequest) GetCustomerNameOrPhone() *string {
	return s.CustomerNameOrPhone
}

func (s *ReadOutboundTaskCallListRequest) GetDisplayStatusList() []*string {
	return s.DisplayStatusList
}

func (s *ReadOutboundTaskCallListRequest) GetDurationRangeList() []*string {
	return s.DurationRangeList
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

func (s *ReadOutboundTaskCallListRequest) GetStructuredLabelTags() []*ReadOutboundTaskCallListRequestStructuredLabelTags {
	return s.StructuredLabelTags
}

func (s *ReadOutboundTaskCallListRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ReadOutboundTaskCallListRequest) GetUserId() *string {
	return s.UserId
}

func (s *ReadOutboundTaskCallListRequest) SetCallEndTimeBegin(v int64) *ReadOutboundTaskCallListRequest {
	s.CallEndTimeBegin = &v
	return s
}

func (s *ReadOutboundTaskCallListRequest) SetCallEndTimeEnd(v int64) *ReadOutboundTaskCallListRequest {
	s.CallEndTimeEnd = &v
	return s
}

func (s *ReadOutboundTaskCallListRequest) SetCallStartTimeBegin(v int64) *ReadOutboundTaskCallListRequest {
	s.CallStartTimeBegin = &v
	return s
}

func (s *ReadOutboundTaskCallListRequest) SetCallStartTimeEnd(v int64) *ReadOutboundTaskCallListRequest {
	s.CallStartTimeEnd = &v
	return s
}

func (s *ReadOutboundTaskCallListRequest) SetCallerUacAccountId(v string) *ReadOutboundTaskCallListRequest {
	s.CallerUacAccountId = &v
	return s
}

func (s *ReadOutboundTaskCallListRequest) SetCurrent(v int32) *ReadOutboundTaskCallListRequest {
	s.Current = &v
	return s
}

func (s *ReadOutboundTaskCallListRequest) SetCurrentWorkspaceCode(v string) *ReadOutboundTaskCallListRequest {
	s.CurrentWorkspaceCode = &v
	return s
}

func (s *ReadOutboundTaskCallListRequest) SetCurrentWorkspaceId(v string) *ReadOutboundTaskCallListRequest {
	s.CurrentWorkspaceId = &v
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

func (s *ReadOutboundTaskCallListRequest) SetDurationRangeList(v []*string) *ReadOutboundTaskCallListRequest {
	s.DurationRangeList = v
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

func (s *ReadOutboundTaskCallListRequest) SetStructuredLabelTags(v []*ReadOutboundTaskCallListRequestStructuredLabelTags) *ReadOutboundTaskCallListRequest {
	s.StructuredLabelTags = v
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
	if s.StructuredLabelTags != nil {
		for _, item := range s.StructuredLabelTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReadOutboundTaskCallListRequestStructuredLabelTags struct {
	// example:
	//
	// 是否有意向
	Key   *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ReadOutboundTaskCallListRequestStructuredLabelTags) String() string {
	return dara.Prettify(s)
}

func (s ReadOutboundTaskCallListRequestStructuredLabelTags) GoString() string {
	return s.String()
}

func (s *ReadOutboundTaskCallListRequestStructuredLabelTags) GetKey() *string {
	return s.Key
}

func (s *ReadOutboundTaskCallListRequestStructuredLabelTags) GetValue() []*string {
	return s.Value
}

func (s *ReadOutboundTaskCallListRequestStructuredLabelTags) SetKey(v string) *ReadOutboundTaskCallListRequestStructuredLabelTags {
	s.Key = &v
	return s
}

func (s *ReadOutboundTaskCallListRequestStructuredLabelTags) SetValue(v []*string) *ReadOutboundTaskCallListRequestStructuredLabelTags {
	s.Value = v
	return s
}

func (s *ReadOutboundTaskCallListRequestStructuredLabelTags) Validate() error {
	return dara.Validate(s)
}
