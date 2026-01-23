// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityRuleTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQualityRuleTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetQualityRuleTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetQualityRuleTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQualityRuleTaskResponseBody
	GetRequestId() *string
	SetRuleTaskInfo(v *GetQualityRuleTaskResponseBodyRuleTaskInfo) *GetQualityRuleTaskResponseBody
	GetRuleTaskInfo() *GetQualityRuleTaskResponseBodyRuleTaskInfo
	SetSuccess(v bool) *GetQualityRuleTaskResponseBody
	GetSuccess() *bool
}

type GetQualityRuleTaskResponseBody struct {
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
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleTaskInfo *GetQualityRuleTaskResponseBodyRuleTaskInfo `json:"RuleTaskInfo,omitempty" xml:"RuleTaskInfo,omitempty" type:"Struct"`
	Success      *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQualityRuleTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityRuleTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQualityRuleTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetQualityRuleTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQualityRuleTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualityRuleTaskResponseBody) GetRuleTaskInfo() *GetQualityRuleTaskResponseBodyRuleTaskInfo {
	return s.RuleTaskInfo
}

func (s *GetQualityRuleTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQualityRuleTaskResponseBody) SetCode(v string) *GetQualityRuleTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityRuleTaskResponseBody) SetHttpStatusCode(v int32) *GetQualityRuleTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetQualityRuleTaskResponseBody) SetMessage(v string) *GetQualityRuleTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityRuleTaskResponseBody) SetRequestId(v string) *GetQualityRuleTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityRuleTaskResponseBody) SetRuleTaskInfo(v *GetQualityRuleTaskResponseBodyRuleTaskInfo) *GetQualityRuleTaskResponseBody {
	s.RuleTaskInfo = v
	return s
}

func (s *GetQualityRuleTaskResponseBody) SetSuccess(v bool) *GetQualityRuleTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityRuleTaskResponseBody) Validate() error {
	if s.RuleTaskInfo != nil {
		if err := s.RuleTaskInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQualityRuleTaskResponseBodyRuleTaskInfo struct {
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

func (s GetQualityRuleTaskResponseBodyRuleTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleTaskResponseBodyRuleTaskInfo) GoString() string {
	return s.String()
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) GetBizDate() *string {
	return s.BizDate
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) GetBizDateFormat() *string {
	return s.BizDateFormat
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) GetCreator() *string {
	return s.Creator
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) GetEndTime() *string {
	return s.EndTime
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) GetId() *int64 {
	return s.Id
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) GetModifier() *string {
	return s.Modifier
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) GetRuleId() *int64 {
	return s.RuleId
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) GetStatus() *string {
	return s.Status
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) GetValidateObjectName() *string {
	return s.ValidateObjectName
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) GetValidateObjectType() *string {
	return s.ValidateObjectType
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) GetValidatePartition() *string {
	return s.ValidatePartition
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) GetValidateSuccess() *bool {
	return s.ValidateSuccess
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) GetWatchId() *int64 {
	return s.WatchId
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) GetWatchTaskId() *int64 {
	return s.WatchTaskId
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) SetBizDate(v string) *GetQualityRuleTaskResponseBodyRuleTaskInfo {
	s.BizDate = &v
	return s
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) SetBizDateFormat(v string) *GetQualityRuleTaskResponseBodyRuleTaskInfo {
	s.BizDateFormat = &v
	return s
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) SetCreateTime(v string) *GetQualityRuleTaskResponseBodyRuleTaskInfo {
	s.CreateTime = &v
	return s
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) SetCreator(v string) *GetQualityRuleTaskResponseBodyRuleTaskInfo {
	s.Creator = &v
	return s
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) SetEndTime(v string) *GetQualityRuleTaskResponseBodyRuleTaskInfo {
	s.EndTime = &v
	return s
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) SetId(v int64) *GetQualityRuleTaskResponseBodyRuleTaskInfo {
	s.Id = &v
	return s
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) SetModifier(v string) *GetQualityRuleTaskResponseBodyRuleTaskInfo {
	s.Modifier = &v
	return s
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) SetModifyTime(v string) *GetQualityRuleTaskResponseBodyRuleTaskInfo {
	s.ModifyTime = &v
	return s
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) SetRuleId(v int64) *GetQualityRuleTaskResponseBodyRuleTaskInfo {
	s.RuleId = &v
	return s
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) SetStartTime(v string) *GetQualityRuleTaskResponseBodyRuleTaskInfo {
	s.StartTime = &v
	return s
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) SetStatus(v string) *GetQualityRuleTaskResponseBodyRuleTaskInfo {
	s.Status = &v
	return s
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) SetTemplateId(v int64) *GetQualityRuleTaskResponseBodyRuleTaskInfo {
	s.TemplateId = &v
	return s
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) SetValidateObjectName(v string) *GetQualityRuleTaskResponseBodyRuleTaskInfo {
	s.ValidateObjectName = &v
	return s
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) SetValidateObjectType(v string) *GetQualityRuleTaskResponseBodyRuleTaskInfo {
	s.ValidateObjectType = &v
	return s
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) SetValidatePartition(v string) *GetQualityRuleTaskResponseBodyRuleTaskInfo {
	s.ValidatePartition = &v
	return s
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) SetValidateSuccess(v bool) *GetQualityRuleTaskResponseBodyRuleTaskInfo {
	s.ValidateSuccess = &v
	return s
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) SetWatchId(v int64) *GetQualityRuleTaskResponseBodyRuleTaskInfo {
	s.WatchId = &v
	return s
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) SetWatchTaskId(v int64) *GetQualityRuleTaskResponseBodyRuleTaskInfo {
	s.WatchTaskId = &v
	return s
}

func (s *GetQualityRuleTaskResponseBodyRuleTaskInfo) Validate() error {
	return dara.Validate(s)
}
