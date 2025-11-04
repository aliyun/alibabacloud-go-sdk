// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveSnapshotTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *ListLiveSnapshotTemplatesResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListLiveSnapshotTemplatesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListLiveSnapshotTemplatesResponseBody
	GetRequestId() *string
	SetSortBy(v string) *ListLiveSnapshotTemplatesResponseBody
	GetSortBy() *string
	SetTemplateList(v []*ListLiveSnapshotTemplatesResponseBodyTemplateList) *ListLiveSnapshotTemplatesResponseBody
	GetTemplateList() []*ListLiveSnapshotTemplatesResponseBodyTemplateList
	SetTotalCount(v int64) *ListLiveSnapshotTemplatesResponseBody
	GetTotalCount() *int64
}

type ListLiveSnapshotTemplatesResponseBody struct {
	// The number of the returned page.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The sorting order of the results by creation time.
	//
	// example:
	//
	// desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The list of the templates.
	TemplateList []*ListLiveSnapshotTemplatesResponseBodyTemplateList `json:"TemplateList,omitempty" xml:"TemplateList,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLiveSnapshotTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLiveSnapshotTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveSnapshotTemplatesResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListLiveSnapshotTemplatesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLiveSnapshotTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLiveSnapshotTemplatesResponseBody) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLiveSnapshotTemplatesResponseBody) GetTemplateList() []*ListLiveSnapshotTemplatesResponseBodyTemplateList {
	return s.TemplateList
}

func (s *ListLiveSnapshotTemplatesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListLiveSnapshotTemplatesResponseBody) SetPageNo(v int32) *ListLiveSnapshotTemplatesResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListLiveSnapshotTemplatesResponseBody) SetPageSize(v int32) *ListLiveSnapshotTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLiveSnapshotTemplatesResponseBody) SetRequestId(v string) *ListLiveSnapshotTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLiveSnapshotTemplatesResponseBody) SetSortBy(v string) *ListLiveSnapshotTemplatesResponseBody {
	s.SortBy = &v
	return s
}

func (s *ListLiveSnapshotTemplatesResponseBody) SetTemplateList(v []*ListLiveSnapshotTemplatesResponseBodyTemplateList) *ListLiveSnapshotTemplatesResponseBody {
	s.TemplateList = v
	return s
}

func (s *ListLiveSnapshotTemplatesResponseBody) SetTotalCount(v int64) *ListLiveSnapshotTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLiveSnapshotTemplatesResponseBody) Validate() error {
	if s.TemplateList != nil {
		for _, item := range s.TemplateList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLiveSnapshotTemplatesResponseBodyTemplateList struct {
	// The time when the job was created.
	//
	// example:
	//
	// 2022-07-20T02:48:58Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The template ID.
	//
	// example:
	//
	// ****a046-263c-3560-978a-fb287782****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The template name.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The interval between two adjacent snapshots. Unit: seconds.
	//
	// example:
	//
	// 10
	TimeInterval *int32 `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty"`
	// The type of the template.
	//
	// Valid values:
	//
	// 	- system
	//
	// 	- custom
	//
	// example:
	//
	// custom
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListLiveSnapshotTemplatesResponseBodyTemplateList) String() string {
	return dara.Prettify(s)
}

func (s ListLiveSnapshotTemplatesResponseBodyTemplateList) GoString() string {
	return s.String()
}

func (s *ListLiveSnapshotTemplatesResponseBodyTemplateList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListLiveSnapshotTemplatesResponseBodyTemplateList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListLiveSnapshotTemplatesResponseBodyTemplateList) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListLiveSnapshotTemplatesResponseBodyTemplateList) GetTimeInterval() *int32 {
	return s.TimeInterval
}

func (s *ListLiveSnapshotTemplatesResponseBodyTemplateList) GetType() *string {
	return s.Type
}

func (s *ListLiveSnapshotTemplatesResponseBodyTemplateList) SetCreateTime(v string) *ListLiveSnapshotTemplatesResponseBodyTemplateList {
	s.CreateTime = &v
	return s
}

func (s *ListLiveSnapshotTemplatesResponseBodyTemplateList) SetTemplateId(v string) *ListLiveSnapshotTemplatesResponseBodyTemplateList {
	s.TemplateId = &v
	return s
}

func (s *ListLiveSnapshotTemplatesResponseBodyTemplateList) SetTemplateName(v string) *ListLiveSnapshotTemplatesResponseBodyTemplateList {
	s.TemplateName = &v
	return s
}

func (s *ListLiveSnapshotTemplatesResponseBodyTemplateList) SetTimeInterval(v int32) *ListLiveSnapshotTemplatesResponseBodyTemplateList {
	s.TimeInterval = &v
	return s
}

func (s *ListLiveSnapshotTemplatesResponseBodyTemplateList) SetType(v string) *ListLiveSnapshotTemplatesResponseBodyTemplateList {
	s.Type = &v
	return s
}

func (s *ListLiveSnapshotTemplatesResponseBodyTemplateList) Validate() error {
	return dara.Validate(s)
}
