// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*ListReportResponseBodyDataList) *ListReportResponseBody
	GetDataList() []*ListReportResponseBodyDataList
	SetHasMore(v bool) *ListReportResponseBody
	GetHasMore() *bool
	SetNextCursor(v int64) *ListReportResponseBody
	GetNextCursor() *int64
	SetRequestId(v string) *ListReportResponseBody
	GetRequestId() *string
	SetSize(v int64) *ListReportResponseBody
	GetSize() *int64
}

type ListReportResponseBody struct {
	DataList []*ListReportResponseBodyDataList `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// example:
	//
	// true
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

func (s ListReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListReportResponseBody) GetDataList() []*ListReportResponseBodyDataList {
	return s.DataList
}

func (s *ListReportResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListReportResponseBody) GetNextCursor() *int64 {
	return s.NextCursor
}

func (s *ListReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListReportResponseBody) GetSize() *int64 {
	return s.Size
}

func (s *ListReportResponseBody) SetDataList(v []*ListReportResponseBodyDataList) *ListReportResponseBody {
	s.DataList = v
	return s
}

func (s *ListReportResponseBody) SetHasMore(v bool) *ListReportResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListReportResponseBody) SetNextCursor(v int64) *ListReportResponseBody {
	s.NextCursor = &v
	return s
}

func (s *ListReportResponseBody) SetRequestId(v string) *ListReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListReportResponseBody) SetSize(v int64) *ListReportResponseBody {
	s.Size = &v
	return s
}

func (s *ListReportResponseBody) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListReportResponseBodyDataList struct {
	Contents []*ListReportResponseBodyDataListContents `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
	// example:
	//
	// 1678416166000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1127123
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// example:
	//
	// admin
	CreatorName *string   `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	DeptName    *string   `json:"DeptName,omitempty" xml:"DeptName,omitempty"`
	Images      []*string `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// 31.1264
	Latitude *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	// example:
	//
	// 108.938036
	Longitude *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	// example:
	//
	// 1653543540000
	ModifiedTime *int64  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// d64994d09916c76276dd9bfa23637644
	ReportId     *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListReportResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListReportResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListReportResponseBodyDataList) GetContents() []*ListReportResponseBodyDataListContents {
	return s.Contents
}

func (s *ListReportResponseBodyDataList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListReportResponseBodyDataList) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListReportResponseBodyDataList) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListReportResponseBodyDataList) GetDeptName() *string {
	return s.DeptName
}

func (s *ListReportResponseBodyDataList) GetImages() []*string {
	return s.Images
}

func (s *ListReportResponseBodyDataList) GetLatitude() *string {
	return s.Latitude
}

func (s *ListReportResponseBodyDataList) GetLongitude() *string {
	return s.Longitude
}

func (s *ListReportResponseBodyDataList) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *ListReportResponseBodyDataList) GetRemark() *string {
	return s.Remark
}

func (s *ListReportResponseBodyDataList) GetReportId() *string {
	return s.ReportId
}

func (s *ListReportResponseBodyDataList) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListReportResponseBodyDataList) SetContents(v []*ListReportResponseBodyDataListContents) *ListReportResponseBodyDataList {
	s.Contents = v
	return s
}

func (s *ListReportResponseBodyDataList) SetCreateTime(v int64) *ListReportResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *ListReportResponseBodyDataList) SetCreatorId(v string) *ListReportResponseBodyDataList {
	s.CreatorId = &v
	return s
}

func (s *ListReportResponseBodyDataList) SetCreatorName(v string) *ListReportResponseBodyDataList {
	s.CreatorName = &v
	return s
}

func (s *ListReportResponseBodyDataList) SetDeptName(v string) *ListReportResponseBodyDataList {
	s.DeptName = &v
	return s
}

func (s *ListReportResponseBodyDataList) SetImages(v []*string) *ListReportResponseBodyDataList {
	s.Images = v
	return s
}

func (s *ListReportResponseBodyDataList) SetLatitude(v string) *ListReportResponseBodyDataList {
	s.Latitude = &v
	return s
}

func (s *ListReportResponseBodyDataList) SetLongitude(v string) *ListReportResponseBodyDataList {
	s.Longitude = &v
	return s
}

func (s *ListReportResponseBodyDataList) SetModifiedTime(v int64) *ListReportResponseBodyDataList {
	s.ModifiedTime = &v
	return s
}

func (s *ListReportResponseBodyDataList) SetRemark(v string) *ListReportResponseBodyDataList {
	s.Remark = &v
	return s
}

func (s *ListReportResponseBodyDataList) SetReportId(v string) *ListReportResponseBodyDataList {
	s.ReportId = &v
	return s
}

func (s *ListReportResponseBodyDataList) SetTemplateName(v string) *ListReportResponseBodyDataList {
	s.TemplateName = &v
	return s
}

func (s *ListReportResponseBodyDataList) Validate() error {
	if s.Contents != nil {
		for _, item := range s.Contents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListReportResponseBodyDataListContents struct {
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 0
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 1
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// ce9ab5d4a80a9401f97c7077e6a9634bd
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListReportResponseBodyDataListContents) String() string {
	return dara.Prettify(s)
}

func (s ListReportResponseBodyDataListContents) GoString() string {
	return s.String()
}

func (s *ListReportResponseBodyDataListContents) GetKey() *string {
	return s.Key
}

func (s *ListReportResponseBodyDataListContents) GetSort() *string {
	return s.Sort
}

func (s *ListReportResponseBodyDataListContents) GetType() *string {
	return s.Type
}

func (s *ListReportResponseBodyDataListContents) GetValue() *string {
	return s.Value
}

func (s *ListReportResponseBodyDataListContents) SetKey(v string) *ListReportResponseBodyDataListContents {
	s.Key = &v
	return s
}

func (s *ListReportResponseBodyDataListContents) SetSort(v string) *ListReportResponseBodyDataListContents {
	s.Sort = &v
	return s
}

func (s *ListReportResponseBodyDataListContents) SetType(v string) *ListReportResponseBodyDataListContents {
	s.Type = &v
	return s
}

func (s *ListReportResponseBodyDataListContents) SetValue(v string) *ListReportResponseBodyDataListContents {
	s.Value = &v
	return s
}

func (s *ListReportResponseBodyDataListContents) Validate() error {
	return dara.Validate(s)
}
