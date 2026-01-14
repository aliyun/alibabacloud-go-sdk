// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DetailsResponseBody
	GetCode() *int64
	SetMessage(v string) *DetailsResponseBody
	GetMessage() *string
	SetModel(v *DetailsResponseBodyModel) *DetailsResponseBody
	GetModel() *DetailsResponseBodyModel
	SetRequestId(v string) *DetailsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DetailsResponseBody
	GetSuccess() *bool
	SetTimestamp(v int64) *DetailsResponseBody
	GetTimestamp() *int64
}

type DetailsResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值
	Message *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *DetailsResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 389b2d0a-37e2-406d-b576-1fc0827be46a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1686279332221
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DetailsResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DetailsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DetailsResponseBody) GetModel() *DetailsResponseBodyModel {
	return s.Model
}

func (s *DetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetailsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DetailsResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DetailsResponseBody) SetCode(v int64) *DetailsResponseBody {
	s.Code = &v
	return s
}

func (s *DetailsResponseBody) SetMessage(v string) *DetailsResponseBody {
	s.Message = &v
	return s
}

func (s *DetailsResponseBody) SetModel(v *DetailsResponseBodyModel) *DetailsResponseBody {
	s.Model = v
	return s
}

func (s *DetailsResponseBody) SetRequestId(v string) *DetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetailsResponseBody) SetSuccess(v bool) *DetailsResponseBody {
	s.Success = &v
	return s
}

func (s *DetailsResponseBody) SetTimestamp(v int64) *DetailsResponseBody {
	s.Timestamp = &v
	return s
}

func (s *DetailsResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetailsResponseBodyModel struct {
	List []*DetailsResponseBodyModelList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 94
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 79
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 71.8132
	TotalPage *float32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DetailsResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s DetailsResponseBodyModel) GoString() string {
	return s.String()
}

func (s *DetailsResponseBodyModel) GetList() []*DetailsResponseBodyModelList {
	return s.List
}

func (s *DetailsResponseBodyModel) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DetailsResponseBodyModel) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DetailsResponseBodyModel) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DetailsResponseBodyModel) GetTotalPage() *float32 {
	return s.TotalPage
}

func (s *DetailsResponseBodyModel) SetList(v []*DetailsResponseBodyModelList) *DetailsResponseBodyModel {
	s.List = v
	return s
}

func (s *DetailsResponseBodyModel) SetPageNo(v int64) *DetailsResponseBodyModel {
	s.PageNo = &v
	return s
}

func (s *DetailsResponseBodyModel) SetPageSize(v int64) *DetailsResponseBodyModel {
	s.PageSize = &v
	return s
}

func (s *DetailsResponseBodyModel) SetTotalCount(v int64) *DetailsResponseBodyModel {
	s.TotalCount = &v
	return s
}

func (s *DetailsResponseBodyModel) SetTotalPage(v float32) *DetailsResponseBodyModel {
	s.TotalPage = &v
	return s
}

func (s *DetailsResponseBodyModel) Validate() error {
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

type DetailsResponseBodyModelList struct {
	// 批次号
	//
	// example:
	//
	// 27
	BatchId *int64 `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// 呼叫状态描述
	//
	// example:
	//
	// 示例值示例值
	CallDesc *string `json:"CallDesc,omitempty" xml:"CallDesc,omitempty"`
	// 外呼ID
	//
	// example:
	//
	// 28dd74a4-30ec-43c0-9bec-272924c51eeb
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// 呼叫状态
	//
	// example:
	//
	// 1
	CallStatus *int64 `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	// 外呼类型
	//
	// example:
	//
	// 2001
	CallType *int64 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 导入时间
	//
	// example:
	//
	// 2023-04-25 15:19:02
	ImportTime *string `json:"ImportTime,omitempty" xml:"ImportTime,omitempty"`
	// 拦截原因
	//
	// example:
	//
	// 示例值示例值示例值
	InterceptReason *string `json:"InterceptReason,omitempty" xml:"InterceptReason,omitempty"`
	// 外呼号码
	//
	// example:
	//
	// 188******454
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// 号码状态描述
	//
	// example:
	//
	// 示例值
	NumberDesc *string `json:"NumberDesc,omitempty" xml:"NumberDesc,omitempty"`
	// 外呼号码MD5
	//
	// example:
	//
	// cbe1b40f76f2cca4735954886b706ffa
	NumberMD5 *string `json:"NumberMD5,omitempty" xml:"NumberMD5,omitempty"`
	// 号码状态
	//
	// example:
	//
	// 1
	NumberStatus *int64 `json:"NumberStatus,omitempty" xml:"NumberStatus,omitempty"`
	// 用户自定义标签
	//
	// example:
	//
	// A
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// 任务ID
	//
	// example:
	//
	// 28
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DetailsResponseBodyModelList) String() string {
	return dara.Prettify(s)
}

func (s DetailsResponseBodyModelList) GoString() string {
	return s.String()
}

func (s *DetailsResponseBodyModelList) GetBatchId() *int64 {
	return s.BatchId
}

func (s *DetailsResponseBodyModelList) GetCallDesc() *string {
	return s.CallDesc
}

func (s *DetailsResponseBodyModelList) GetCallId() *string {
	return s.CallId
}

func (s *DetailsResponseBodyModelList) GetCallStatus() *int64 {
	return s.CallStatus
}

func (s *DetailsResponseBodyModelList) GetCallType() *int64 {
	return s.CallType
}

func (s *DetailsResponseBodyModelList) GetImportTime() *string {
	return s.ImportTime
}

func (s *DetailsResponseBodyModelList) GetInterceptReason() *string {
	return s.InterceptReason
}

func (s *DetailsResponseBodyModelList) GetNumber() *string {
	return s.Number
}

func (s *DetailsResponseBodyModelList) GetNumberDesc() *string {
	return s.NumberDesc
}

func (s *DetailsResponseBodyModelList) GetNumberMD5() *string {
	return s.NumberMD5
}

func (s *DetailsResponseBodyModelList) GetNumberStatus() *int64 {
	return s.NumberStatus
}

func (s *DetailsResponseBodyModelList) GetTag() *string {
	return s.Tag
}

func (s *DetailsResponseBodyModelList) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DetailsResponseBodyModelList) SetBatchId(v int64) *DetailsResponseBodyModelList {
	s.BatchId = &v
	return s
}

func (s *DetailsResponseBodyModelList) SetCallDesc(v string) *DetailsResponseBodyModelList {
	s.CallDesc = &v
	return s
}

func (s *DetailsResponseBodyModelList) SetCallId(v string) *DetailsResponseBodyModelList {
	s.CallId = &v
	return s
}

func (s *DetailsResponseBodyModelList) SetCallStatus(v int64) *DetailsResponseBodyModelList {
	s.CallStatus = &v
	return s
}

func (s *DetailsResponseBodyModelList) SetCallType(v int64) *DetailsResponseBodyModelList {
	s.CallType = &v
	return s
}

func (s *DetailsResponseBodyModelList) SetImportTime(v string) *DetailsResponseBodyModelList {
	s.ImportTime = &v
	return s
}

func (s *DetailsResponseBodyModelList) SetInterceptReason(v string) *DetailsResponseBodyModelList {
	s.InterceptReason = &v
	return s
}

func (s *DetailsResponseBodyModelList) SetNumber(v string) *DetailsResponseBodyModelList {
	s.Number = &v
	return s
}

func (s *DetailsResponseBodyModelList) SetNumberDesc(v string) *DetailsResponseBodyModelList {
	s.NumberDesc = &v
	return s
}

func (s *DetailsResponseBodyModelList) SetNumberMD5(v string) *DetailsResponseBodyModelList {
	s.NumberMD5 = &v
	return s
}

func (s *DetailsResponseBodyModelList) SetNumberStatus(v int64) *DetailsResponseBodyModelList {
	s.NumberStatus = &v
	return s
}

func (s *DetailsResponseBodyModelList) SetTag(v string) *DetailsResponseBodyModelList {
	s.Tag = &v
	return s
}

func (s *DetailsResponseBodyModelList) SetTaskId(v int64) *DetailsResponseBodyModelList {
	s.TaskId = &v
	return s
}

func (s *DetailsResponseBodyModelList) Validate() error {
	return dara.Validate(s)
}
