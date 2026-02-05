// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListYunQiTaskByUidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListYunQiTaskByUidResponseBody
	GetCode() *string
	SetData(v *ListYunQiTaskByUidResponseBodyData) *ListYunQiTaskByUidResponseBody
	GetData() *ListYunQiTaskByUidResponseBodyData
	SetHttpStatusCode(v int32) *ListYunQiTaskByUidResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListYunQiTaskByUidResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListYunQiTaskByUidResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListYunQiTaskByUidResponseBody
	GetSuccess() *bool
}

type ListYunQiTaskByUidResponseBody struct {
	Code           *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data           *ListYunQiTaskByUidResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	HttpStatusCode *int32                              `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string                             `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool                               `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListYunQiTaskByUidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListYunQiTaskByUidResponseBody) GoString() string {
	return s.String()
}

func (s *ListYunQiTaskByUidResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListYunQiTaskByUidResponseBody) GetData() *ListYunQiTaskByUidResponseBodyData {
	return s.Data
}

func (s *ListYunQiTaskByUidResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListYunQiTaskByUidResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListYunQiTaskByUidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListYunQiTaskByUidResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListYunQiTaskByUidResponseBody) SetCode(v string) *ListYunQiTaskByUidResponseBody {
	s.Code = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBody) SetData(v *ListYunQiTaskByUidResponseBodyData) *ListYunQiTaskByUidResponseBody {
	s.Data = v
	return s
}

func (s *ListYunQiTaskByUidResponseBody) SetHttpStatusCode(v int32) *ListYunQiTaskByUidResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBody) SetMessage(v string) *ListYunQiTaskByUidResponseBody {
	s.Message = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBody) SetRequestId(v string) *ListYunQiTaskByUidResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBody) SetSuccess(v bool) *ListYunQiTaskByUidResponseBody {
	s.Success = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListYunQiTaskByUidResponseBodyData struct {
	Extend   interface{}                               `json:"extend,omitempty" xml:"extend,omitempty"`
	List     []*ListYunQiTaskByUidResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNum  *int32                                    `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32                                    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                                    `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListYunQiTaskByUidResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListYunQiTaskByUidResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListYunQiTaskByUidResponseBodyData) GetExtend() interface{} {
	return s.Extend
}

func (s *ListYunQiTaskByUidResponseBodyData) GetList() []*ListYunQiTaskByUidResponseBodyDataList {
	return s.List
}

func (s *ListYunQiTaskByUidResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListYunQiTaskByUidResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListYunQiTaskByUidResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListYunQiTaskByUidResponseBodyData) SetExtend(v interface{}) *ListYunQiTaskByUidResponseBodyData {
	s.Extend = v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyData) SetList(v []*ListYunQiTaskByUidResponseBodyDataList) *ListYunQiTaskByUidResponseBodyData {
	s.List = v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyData) SetPageNum(v int32) *ListYunQiTaskByUidResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyData) SetPageSize(v int32) *ListYunQiTaskByUidResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyData) SetTotal(v int32) *ListYunQiTaskByUidResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListYunQiTaskByUidResponseBodyDataList struct {
	ChatId               *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	CreateTime           *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorName          *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	EndTime              *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	EvaluationStar       *int32  `json:"evaluationStar,omitempty" xml:"evaluationStar,omitempty"`
	Important            *string `json:"important,omitempty" xml:"important,omitempty"`
	MainDingDepartmentId *string `json:"mainDingDepartmentId,omitempty" xml:"mainDingDepartmentId,omitempty"`
	MainDingGroupId      *string `json:"mainDingGroupId,omitempty" xml:"mainDingGroupId,omitempty"`
	MainDingGroupName    *string `json:"mainDingGroupName,omitempty" xml:"mainDingGroupName,omitempty"`
	ProductName          *string `json:"productName,omitempty" xml:"productName,omitempty"`
	RecordId             *string `json:"recordId,omitempty" xml:"recordId,omitempty"`
	Status               *string `json:"status,omitempty" xml:"status,omitempty"`
	SubDingDepartmentId  *string `json:"subDingDepartmentId,omitempty" xml:"subDingDepartmentId,omitempty"`
	SubDingGroupId       *string `json:"subDingGroupId,omitempty" xml:"subDingGroupId,omitempty"`
	SubDingGroupName     *string `json:"subDingGroupName,omitempty" xml:"subDingGroupName,omitempty"`
	Title                *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListYunQiTaskByUidResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListYunQiTaskByUidResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListYunQiTaskByUidResponseBodyDataList) GetChatId() *string {
	return s.ChatId
}

func (s *ListYunQiTaskByUidResponseBodyDataList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListYunQiTaskByUidResponseBodyDataList) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListYunQiTaskByUidResponseBodyDataList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListYunQiTaskByUidResponseBodyDataList) GetEvaluationStar() *int32 {
	return s.EvaluationStar
}

func (s *ListYunQiTaskByUidResponseBodyDataList) GetImportant() *string {
	return s.Important
}

func (s *ListYunQiTaskByUidResponseBodyDataList) GetMainDingDepartmentId() *string {
	return s.MainDingDepartmentId
}

func (s *ListYunQiTaskByUidResponseBodyDataList) GetMainDingGroupId() *string {
	return s.MainDingGroupId
}

func (s *ListYunQiTaskByUidResponseBodyDataList) GetMainDingGroupName() *string {
	return s.MainDingGroupName
}

func (s *ListYunQiTaskByUidResponseBodyDataList) GetProductName() *string {
	return s.ProductName
}

func (s *ListYunQiTaskByUidResponseBodyDataList) GetRecordId() *string {
	return s.RecordId
}

func (s *ListYunQiTaskByUidResponseBodyDataList) GetStatus() *string {
	return s.Status
}

func (s *ListYunQiTaskByUidResponseBodyDataList) GetSubDingDepartmentId() *string {
	return s.SubDingDepartmentId
}

func (s *ListYunQiTaskByUidResponseBodyDataList) GetSubDingGroupId() *string {
	return s.SubDingGroupId
}

func (s *ListYunQiTaskByUidResponseBodyDataList) GetSubDingGroupName() *string {
	return s.SubDingGroupName
}

func (s *ListYunQiTaskByUidResponseBodyDataList) GetTitle() *string {
	return s.Title
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetChatId(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.ChatId = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetCreateTime(v int64) *ListYunQiTaskByUidResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetCreatorName(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.CreatorName = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetEndTime(v int64) *ListYunQiTaskByUidResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetEvaluationStar(v int32) *ListYunQiTaskByUidResponseBodyDataList {
	s.EvaluationStar = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetImportant(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.Important = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetMainDingDepartmentId(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.MainDingDepartmentId = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetMainDingGroupId(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.MainDingGroupId = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetMainDingGroupName(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.MainDingGroupName = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetProductName(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.ProductName = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetRecordId(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.RecordId = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetStatus(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetSubDingDepartmentId(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.SubDingDepartmentId = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetSubDingGroupId(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.SubDingGroupId = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetSubDingGroupName(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.SubDingGroupName = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) SetTitle(v string) *ListYunQiTaskByUidResponseBodyDataList {
	s.Title = &v
	return s
}

func (s *ListYunQiTaskByUidResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
