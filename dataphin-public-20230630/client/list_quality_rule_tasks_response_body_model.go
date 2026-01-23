// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityRuleTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListQualityRuleTasksResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListQualityRuleTasksResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListQualityRuleTasksResponseBody
	GetMessage() *string
	SetPageResult(v *ListQualityRuleTasksResponseBodyPageResult) *ListQualityRuleTasksResponseBody
	GetPageResult() *ListQualityRuleTasksResponseBodyPageResult
	SetRequestId(v string) *ListQualityRuleTasksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListQualityRuleTasksResponseBody
	GetSuccess() *bool
}

type ListQualityRuleTasksResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message    *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListQualityRuleTasksResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListQualityRuleTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRuleTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListQualityRuleTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListQualityRuleTasksResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListQualityRuleTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListQualityRuleTasksResponseBody) GetPageResult() *ListQualityRuleTasksResponseBodyPageResult {
	return s.PageResult
}

func (s *ListQualityRuleTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQualityRuleTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListQualityRuleTasksResponseBody) SetCode(v string) *ListQualityRuleTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListQualityRuleTasksResponseBody) SetHttpStatusCode(v int32) *ListQualityRuleTasksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListQualityRuleTasksResponseBody) SetMessage(v string) *ListQualityRuleTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListQualityRuleTasksResponseBody) SetPageResult(v *ListQualityRuleTasksResponseBodyPageResult) *ListQualityRuleTasksResponseBody {
	s.PageResult = v
	return s
}

func (s *ListQualityRuleTasksResponseBody) SetRequestId(v string) *ListQualityRuleTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQualityRuleTasksResponseBody) SetSuccess(v bool) *ListQualityRuleTasksResponseBody {
	s.Success = &v
	return s
}

func (s *ListQualityRuleTasksResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQualityRuleTasksResponseBodyPageResult struct {
	QualityRuleTaskList []*ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList `json:"QualityRuleTaskList,omitempty" xml:"QualityRuleTaskList,omitempty" type:"Repeated"`
	// example:
	//
	// 68
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQualityRuleTasksResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRuleTasksResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListQualityRuleTasksResponseBodyPageResult) GetQualityRuleTaskList() []*ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList {
	return s.QualityRuleTaskList
}

func (s *ListQualityRuleTasksResponseBodyPageResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListQualityRuleTasksResponseBodyPageResult) SetQualityRuleTaskList(v []*ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) *ListQualityRuleTasksResponseBodyPageResult {
	s.QualityRuleTaskList = v
	return s
}

func (s *ListQualityRuleTasksResponseBodyPageResult) SetTotalCount(v int64) *ListQualityRuleTasksResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListQualityRuleTasksResponseBodyPageResult) Validate() error {
	if s.QualityRuleTaskList != nil {
		for _, item := range s.QualityRuleTaskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList struct {
	// example:
	//
	// 2025-06-30
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// example:
	//
	// yyyy-MM-dd
	BizDateFormat *string `json:"BizDateFormat,omitempty" xml:"BizDateFormat,omitempty"`
	// example:
	//
	// 2025-06-30 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 30012011
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 2025-06-30 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 30012011
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// example:
	//
	// 2025-06-30 00:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 11
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// 2025-06-30 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 11
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// test
	ValidateObjectName *string `json:"ValidateObjectName,omitempty" xml:"ValidateObjectName,omitempty"`
	// example:
	//
	// column
	ValidateObjectType *string `json:"ValidateObjectType,omitempty" xml:"ValidateObjectType,omitempty"`
	// example:
	//
	// 20251011
	ValidatePartition *string `json:"ValidatePartition,omitempty" xml:"ValidatePartition,omitempty"`
	ValidateSuccess   *bool   `json:"ValidateSuccess,omitempty" xml:"ValidateSuccess,omitempty"`
	// example:
	//
	// 1
	WatchId *int64 `json:"WatchId,omitempty" xml:"WatchId,omitempty"`
	// example:
	//
	// 1
	WatchTaskId *int64 `json:"WatchTaskId,omitempty" xml:"WatchTaskId,omitempty"`
}

func (s ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) GoString() string {
	return s.String()
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) GetBizDate() *string {
	return s.BizDate
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) GetBizDateFormat() *string {
	return s.BizDateFormat
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) GetCreator() *string {
	return s.Creator
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) GetEndTime() *string {
	return s.EndTime
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) GetId() *int64 {
	return s.Id
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) GetModifier() *string {
	return s.Modifier
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) GetStartTime() *string {
	return s.StartTime
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) GetStatus() *string {
	return s.Status
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) GetValidateObjectName() *string {
	return s.ValidateObjectName
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) GetValidateObjectType() *string {
	return s.ValidateObjectType
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) GetValidatePartition() *string {
	return s.ValidatePartition
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) GetValidateSuccess() *bool {
	return s.ValidateSuccess
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) GetWatchId() *int64 {
	return s.WatchId
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) GetWatchTaskId() *int64 {
	return s.WatchTaskId
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) SetBizDate(v string) *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList {
	s.BizDate = &v
	return s
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) SetBizDateFormat(v string) *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList {
	s.BizDateFormat = &v
	return s
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) SetCreateTime(v string) *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList {
	s.CreateTime = &v
	return s
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) SetCreator(v string) *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList {
	s.Creator = &v
	return s
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) SetEndTime(v string) *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList {
	s.EndTime = &v
	return s
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) SetId(v int64) *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList {
	s.Id = &v
	return s
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) SetModifier(v string) *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList {
	s.Modifier = &v
	return s
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) SetModifyTime(v string) *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList {
	s.ModifyTime = &v
	return s
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) SetRuleId(v int64) *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList {
	s.RuleId = &v
	return s
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) SetStartTime(v string) *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList {
	s.StartTime = &v
	return s
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) SetStatus(v string) *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList {
	s.Status = &v
	return s
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) SetTemplateId(v int64) *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList {
	s.TemplateId = &v
	return s
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) SetValidateObjectName(v string) *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList {
	s.ValidateObjectName = &v
	return s
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) SetValidateObjectType(v string) *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList {
	s.ValidateObjectType = &v
	return s
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) SetValidatePartition(v string) *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList {
	s.ValidatePartition = &v
	return s
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) SetValidateSuccess(v bool) *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList {
	s.ValidateSuccess = &v
	return s
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) SetWatchId(v int64) *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList {
	s.WatchId = &v
	return s
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) SetWatchTaskId(v int64) *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList {
	s.WatchTaskId = &v
	return s
}

func (s *ListQualityRuleTasksResponseBodyPageResultQualityRuleTaskList) Validate() error {
	return dara.Validate(s)
}
