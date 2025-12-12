// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityCheckSchemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQualityCheckSchemeResponseBody
	GetCode() *string
	SetData(v *GetQualityCheckSchemeResponseBodyData) *GetQualityCheckSchemeResponseBody
	GetData() *GetQualityCheckSchemeResponseBodyData
	SetHttpStatusCode(v int32) *GetQualityCheckSchemeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetQualityCheckSchemeResponseBody
	GetMessage() *string
	SetMessages(v []*string) *GetQualityCheckSchemeResponseBody
	GetMessages() []*string
	SetRequestId(v string) *GetQualityCheckSchemeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQualityCheckSchemeResponseBody
	GetSuccess() *bool
}

type GetQualityCheckSchemeResponseBody struct {
	// example:
	//
	// 200
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetQualityCheckSchemeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages []*string `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// example:
	//
	// 96138D8D-8D26-4E41-BFF4-77AED1088BBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQualityCheckSchemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualityCheckSchemeResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityCheckSchemeResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQualityCheckSchemeResponseBody) GetData() *GetQualityCheckSchemeResponseBodyData {
	return s.Data
}

func (s *GetQualityCheckSchemeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetQualityCheckSchemeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQualityCheckSchemeResponseBody) GetMessages() []*string {
	return s.Messages
}

func (s *GetQualityCheckSchemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualityCheckSchemeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQualityCheckSchemeResponseBody) SetCode(v string) *GetQualityCheckSchemeResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBody) SetData(v *GetQualityCheckSchemeResponseBodyData) *GetQualityCheckSchemeResponseBody {
	s.Data = v
	return s
}

func (s *GetQualityCheckSchemeResponseBody) SetHttpStatusCode(v int32) *GetQualityCheckSchemeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBody) SetMessage(v string) *GetQualityCheckSchemeResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBody) SetMessages(v []*string) *GetQualityCheckSchemeResponseBody {
	s.Messages = v
	return s
}

func (s *GetQualityCheckSchemeResponseBody) SetRequestId(v string) *GetQualityCheckSchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBody) SetSuccess(v bool) *GetQualityCheckSchemeResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQualityCheckSchemeResponseBodyData struct {
	// example:
	//
	// 1616113198000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// xxx
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// 1
	DataType    *int32  `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InitScore   *string `json:"InitScore,omitempty" xml:"InitScore,omitempty"`
	// example:
	//
	// xxx
	Name                *string                                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleIds             []*string                                                   `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
	RuleList            []*RulesInfo                                                `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
	SchemeCheckTypeList []*GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList `json:"SchemeCheckTypeList,omitempty" xml:"SchemeCheckTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// 112**
	SchemeId *int64 `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
	// example:
	//
	// 1
	SchemeTemplateId *int64 `json:"SchemeTemplateId,omitempty" xml:"SchemeTemplateId,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	TemplateType *int32 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1616113198000
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// xxx
	UpdateUserName *string `json:"UpdateUserName,omitempty" xml:"UpdateUserName,omitempty"`
	// example:
	//
	// 1616113198000
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetQualityCheckSchemeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQualityCheckSchemeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQualityCheckSchemeResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetQualityCheckSchemeResponseBodyData) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *GetQualityCheckSchemeResponseBodyData) GetDataType() *int32 {
	return s.DataType
}

func (s *GetQualityCheckSchemeResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetQualityCheckSchemeResponseBodyData) GetInitScore() *string {
	return s.InitScore
}

func (s *GetQualityCheckSchemeResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetQualityCheckSchemeResponseBodyData) GetRuleIds() []*string {
	return s.RuleIds
}

func (s *GetQualityCheckSchemeResponseBodyData) GetRuleList() []*RulesInfo {
	return s.RuleList
}

func (s *GetQualityCheckSchemeResponseBodyData) GetSchemeCheckTypeList() []*GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	return s.SchemeCheckTypeList
}

func (s *GetQualityCheckSchemeResponseBodyData) GetSchemeId() *int64 {
	return s.SchemeId
}

func (s *GetQualityCheckSchemeResponseBodyData) GetSchemeTemplateId() *int64 {
	return s.SchemeTemplateId
}

func (s *GetQualityCheckSchemeResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetQualityCheckSchemeResponseBodyData) GetTemplateType() *int32 {
	return s.TemplateType
}

func (s *GetQualityCheckSchemeResponseBodyData) GetType() *int32 {
	return s.Type
}

func (s *GetQualityCheckSchemeResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetQualityCheckSchemeResponseBodyData) GetUpdateUserName() *string {
	return s.UpdateUserName
}

func (s *GetQualityCheckSchemeResponseBodyData) GetVersion() *int64 {
	return s.Version
}

func (s *GetQualityCheckSchemeResponseBodyData) SetCreateTime(v string) *GetQualityCheckSchemeResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetCreateUserName(v string) *GetQualityCheckSchemeResponseBodyData {
	s.CreateUserName = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetDataType(v int32) *GetQualityCheckSchemeResponseBodyData {
	s.DataType = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetDescription(v string) *GetQualityCheckSchemeResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetInitScore(v string) *GetQualityCheckSchemeResponseBodyData {
	s.InitScore = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetName(v string) *GetQualityCheckSchemeResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetRuleIds(v []*string) *GetQualityCheckSchemeResponseBodyData {
	s.RuleIds = v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetRuleList(v []*RulesInfo) *GetQualityCheckSchemeResponseBodyData {
	s.RuleList = v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetSchemeCheckTypeList(v []*GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) *GetQualityCheckSchemeResponseBodyData {
	s.SchemeCheckTypeList = v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetSchemeId(v int64) *GetQualityCheckSchemeResponseBodyData {
	s.SchemeId = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetSchemeTemplateId(v int64) *GetQualityCheckSchemeResponseBodyData {
	s.SchemeTemplateId = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetStatus(v int32) *GetQualityCheckSchemeResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetTemplateType(v int32) *GetQualityCheckSchemeResponseBodyData {
	s.TemplateType = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetType(v int32) *GetQualityCheckSchemeResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetUpdateTime(v string) *GetQualityCheckSchemeResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetUpdateUserName(v string) *GetQualityCheckSchemeResponseBodyData {
	s.UpdateUserName = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) SetVersion(v int64) *GetQualityCheckSchemeResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyData) Validate() error {
	if s.RuleList != nil {
		for _, item := range s.RuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SchemeCheckTypeList != nil {
		for _, item := range s.SchemeCheckTypeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList struct {
	CheckName *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	// example:
	//
	// 0
	CheckType *int64 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// 32
	SchemeId *int64 `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
	// example:
	//
	// 20
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// 10
	SourceScore *int32 `json:"SourceScore,omitempty" xml:"SourceScore,omitempty"`
}

func (s GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) String() string {
	return dara.Prettify(s)
}

func (s GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) GoString() string {
	return s.String()
}

func (s *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) GetCheckName() *string {
	return s.CheckName
}

func (s *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) GetCheckType() *int64 {
	return s.CheckType
}

func (s *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) GetEnable() *int32 {
	return s.Enable
}

func (s *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) GetSchemeId() *int64 {
	return s.SchemeId
}

func (s *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) GetScore() *int32 {
	return s.Score
}

func (s *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) GetSourceScore() *int32 {
	return s.SourceScore
}

func (s *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetCheckName(v string) *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.CheckName = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetCheckType(v int64) *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.CheckType = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetEnable(v int32) *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.Enable = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetSchemeId(v int64) *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.SchemeId = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetScore(v int32) *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.Score = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) SetSourceScore(v int32) *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList {
	s.SourceScore = &v
	return s
}

func (s *GetQualityCheckSchemeResponseBodyDataSchemeCheckTypeList) Validate() error {
	return dara.Validate(s)
}
