// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v string) *ListEventRecordsRequest
	GetBeginTime() *string
	SetEndTime(v string) *ListEventRecordsRequest
	GetEndTime() *string
	SetPage(v int32) *ListEventRecordsRequest
	GetPage() *int32
	SetSize(v int32) *ListEventRecordsRequest
	GetSize() *int32
	SetTermContent(v string) *ListEventRecordsRequest
	GetTermContent() *string
	SetTermType(v string) *ListEventRecordsRequest
	GetTermType() *string
}

type ListEventRecordsRequest struct {
	// The start time for the event query. Specify a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1746516590000
	BeginTime *string `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// The end time for the event query. Specify a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1746775790000
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The page number in the paging Settings.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page in the paging Settings.
	//
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// The content of the object to filter. If this parameter is empty, event records for all instances are returned. Note: Specify the corresponding value based on the filter type.
	//
	// example:
	//
	// **
	TermContent *string `json:"termContent,omitempty" xml:"termContent,omitempty"`
	// The object type to filter. If this parameter is empty, event records of all types are returned.
	//
	// example:
	//
	// InstanceId
	TermType *string `json:"termType,omitempty" xml:"termType,omitempty"`
}

func (s ListEventRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEventRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListEventRecordsRequest) GetBeginTime() *string {
	return s.BeginTime
}

func (s *ListEventRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListEventRecordsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListEventRecordsRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListEventRecordsRequest) GetTermContent() *string {
	return s.TermContent
}

func (s *ListEventRecordsRequest) GetTermType() *string {
	return s.TermType
}

func (s *ListEventRecordsRequest) SetBeginTime(v string) *ListEventRecordsRequest {
	s.BeginTime = &v
	return s
}

func (s *ListEventRecordsRequest) SetEndTime(v string) *ListEventRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *ListEventRecordsRequest) SetPage(v int32) *ListEventRecordsRequest {
	s.Page = &v
	return s
}

func (s *ListEventRecordsRequest) SetSize(v int32) *ListEventRecordsRequest {
	s.Size = &v
	return s
}

func (s *ListEventRecordsRequest) SetTermContent(v string) *ListEventRecordsRequest {
	s.TermContent = &v
	return s
}

func (s *ListEventRecordsRequest) SetTermType(v string) *ListEventRecordsRequest {
	s.TermType = &v
	return s
}

func (s *ListEventRecordsRequest) Validate() error {
	return dara.Validate(s)
}
