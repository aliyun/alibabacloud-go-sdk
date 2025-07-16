// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSimpleListReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*SimpleListReportResponseBodyDataList) *SimpleListReportResponseBody
	GetDataList() []*SimpleListReportResponseBodyDataList
	SetHasMore(v bool) *SimpleListReportResponseBody
	GetHasMore() *bool
	SetNextCursor(v int64) *SimpleListReportResponseBody
	GetNextCursor() *int64
	SetRequestId(v string) *SimpleListReportResponseBody
	GetRequestId() *string
	SetSize(v int64) *SimpleListReportResponseBody
	GetSize() *int64
}

type SimpleListReportResponseBody struct {
	DataList []*SimpleListReportResponseBodyDataList `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// example:
	//
	// false
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 12312131231
	NextCursor *int64 `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 20
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s SimpleListReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SimpleListReportResponseBody) GoString() string {
	return s.String()
}

func (s *SimpleListReportResponseBody) GetDataList() []*SimpleListReportResponseBodyDataList {
	return s.DataList
}

func (s *SimpleListReportResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *SimpleListReportResponseBody) GetNextCursor() *int64 {
	return s.NextCursor
}

func (s *SimpleListReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SimpleListReportResponseBody) GetSize() *int64 {
	return s.Size
}

func (s *SimpleListReportResponseBody) SetDataList(v []*SimpleListReportResponseBodyDataList) *SimpleListReportResponseBody {
	s.DataList = v
	return s
}

func (s *SimpleListReportResponseBody) SetHasMore(v bool) *SimpleListReportResponseBody {
	s.HasMore = &v
	return s
}

func (s *SimpleListReportResponseBody) SetNextCursor(v int64) *SimpleListReportResponseBody {
	s.NextCursor = &v
	return s
}

func (s *SimpleListReportResponseBody) SetRequestId(v string) *SimpleListReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *SimpleListReportResponseBody) SetSize(v int64) *SimpleListReportResponseBody {
	s.Size = &v
	return s
}

func (s *SimpleListReportResponseBody) Validate() error {
	return dara.Validate(s)
}

type SimpleListReportResponseBodyDataList struct {
	// example:
	//
	// 1567034772000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1701038
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// example:
	//
	// xinmu
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	DeptName    *string `json:"DeptName,omitempty" xml:"DeptName,omitempty"`
	// example:
	//
	// ops
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// d64994d09916c76276dd9bfa23637644
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// example:
	//
	// WG-Template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s SimpleListReportResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s SimpleListReportResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *SimpleListReportResponseBodyDataList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *SimpleListReportResponseBodyDataList) GetCreatorId() *string {
	return s.CreatorId
}

func (s *SimpleListReportResponseBodyDataList) GetCreatorName() *string {
	return s.CreatorName
}

func (s *SimpleListReportResponseBodyDataList) GetDeptName() *string {
	return s.DeptName
}

func (s *SimpleListReportResponseBodyDataList) GetRemark() *string {
	return s.Remark
}

func (s *SimpleListReportResponseBodyDataList) GetReportId() *string {
	return s.ReportId
}

func (s *SimpleListReportResponseBodyDataList) GetTemplateName() *string {
	return s.TemplateName
}

func (s *SimpleListReportResponseBodyDataList) SetCreateTime(v int64) *SimpleListReportResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *SimpleListReportResponseBodyDataList) SetCreatorId(v string) *SimpleListReportResponseBodyDataList {
	s.CreatorId = &v
	return s
}

func (s *SimpleListReportResponseBodyDataList) SetCreatorName(v string) *SimpleListReportResponseBodyDataList {
	s.CreatorName = &v
	return s
}

func (s *SimpleListReportResponseBodyDataList) SetDeptName(v string) *SimpleListReportResponseBodyDataList {
	s.DeptName = &v
	return s
}

func (s *SimpleListReportResponseBodyDataList) SetRemark(v string) *SimpleListReportResponseBodyDataList {
	s.Remark = &v
	return s
}

func (s *SimpleListReportResponseBodyDataList) SetReportId(v string) *SimpleListReportResponseBodyDataList {
	s.ReportId = &v
	return s
}

func (s *SimpleListReportResponseBodyDataList) SetTemplateName(v string) *SimpleListReportResponseBodyDataList {
	s.TemplateName = &v
	return s
}

func (s *SimpleListReportResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
