// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSchemeTaskConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSchemeTaskConfigResponseBody
	GetCode() *string
	SetCount(v int32) *ListSchemeTaskConfigResponseBody
	GetCount() *int32
	SetCurrentPage(v int32) *ListSchemeTaskConfigResponseBody
	GetCurrentPage() *int32
	SetData(v *ListSchemeTaskConfigResponseBodyData) *ListSchemeTaskConfigResponseBody
	GetData() *ListSchemeTaskConfigResponseBodyData
	SetHttpStatusCode(v int32) *ListSchemeTaskConfigResponseBody
	GetHttpStatusCode() *int32
	SetLastDataId(v string) *ListSchemeTaskConfigResponseBody
	GetLastDataId() *string
	SetMessage(v string) *ListSchemeTaskConfigResponseBody
	GetMessage() *string
	SetMessages(v *ListSchemeTaskConfigResponseBodyMessages) *ListSchemeTaskConfigResponseBody
	GetMessages() *ListSchemeTaskConfigResponseBodyMessages
	SetPageNumber(v int32) *ListSchemeTaskConfigResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSchemeTaskConfigResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListSchemeTaskConfigResponseBody
	GetRequestId() *string
	SetResultCountId(v string) *ListSchemeTaskConfigResponseBody
	GetResultCountId() *string
	SetSuccess(v bool) *ListSchemeTaskConfigResponseBody
	GetSuccess() *bool
}

type ListSchemeTaskConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 22
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32                                `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        *ListSchemeTaskConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// xxx
	LastDataId *string `json:"LastDataId,omitempty" xml:"LastDataId,omitempty"`
	// example:
	//
	// successful
	Message  *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *ListSchemeTaskConfigResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 4B0A8DCD-0DDF-5E80-8B9C-0A2F68B3403B
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCountId *string `json:"ResultCountId,omitempty" xml:"ResultCountId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSchemeTaskConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSchemeTaskConfigResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListSchemeTaskConfigResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListSchemeTaskConfigResponseBody) GetData() *ListSchemeTaskConfigResponseBodyData {
	return s.Data
}

func (s *ListSchemeTaskConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSchemeTaskConfigResponseBody) GetLastDataId() *string {
	return s.LastDataId
}

func (s *ListSchemeTaskConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSchemeTaskConfigResponseBody) GetMessages() *ListSchemeTaskConfigResponseBodyMessages {
	return s.Messages
}

func (s *ListSchemeTaskConfigResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSchemeTaskConfigResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSchemeTaskConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSchemeTaskConfigResponseBody) GetResultCountId() *string {
	return s.ResultCountId
}

func (s *ListSchemeTaskConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSchemeTaskConfigResponseBody) SetCode(v string) *ListSchemeTaskConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBody) SetCount(v int32) *ListSchemeTaskConfigResponseBody {
	s.Count = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBody) SetCurrentPage(v int32) *ListSchemeTaskConfigResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBody) SetData(v *ListSchemeTaskConfigResponseBodyData) *ListSchemeTaskConfigResponseBody {
	s.Data = v
	return s
}

func (s *ListSchemeTaskConfigResponseBody) SetHttpStatusCode(v int32) *ListSchemeTaskConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBody) SetLastDataId(v string) *ListSchemeTaskConfigResponseBody {
	s.LastDataId = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBody) SetMessage(v string) *ListSchemeTaskConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBody) SetMessages(v *ListSchemeTaskConfigResponseBodyMessages) *ListSchemeTaskConfigResponseBody {
	s.Messages = v
	return s
}

func (s *ListSchemeTaskConfigResponseBody) SetPageNumber(v int32) *ListSchemeTaskConfigResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBody) SetPageSize(v int32) *ListSchemeTaskConfigResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBody) SetRequestId(v string) *ListSchemeTaskConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBody) SetResultCountId(v string) *ListSchemeTaskConfigResponseBody {
	s.ResultCountId = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBody) SetSuccess(v bool) *ListSchemeTaskConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	if s.Messages != nil {
		if err := s.Messages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSchemeTaskConfigResponseBodyData struct {
	Data []*ListSchemeTaskConfigResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListSchemeTaskConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBodyData) GetData() []*ListSchemeTaskConfigResponseBodyDataData {
	return s.Data
}

func (s *ListSchemeTaskConfigResponseBodyData) SetData(v []*ListSchemeTaskConfigResponseBodyDataData) *ListSchemeTaskConfigResponseBodyData {
	s.Data = v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyData) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSchemeTaskConfigResponseBodyDataData struct {
	// example:
	//
	// 2
	AsrTaskPriority *int32 `json:"AsrTaskPriority,omitempty" xml:"AsrTaskPriority,omitempty"`
	AsrVersion      *int32 `json:"AsrVersion,omitempty" xml:"AsrVersion,omitempty"`
	// example:
	//
	// 0
	AssignType *int32 `json:"AssignType,omitempty" xml:"AssignType,omitempty"`
	// example:
	//
	// 1650418039000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1
	CreateUser *int64                                              `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	DataConfig *ListSchemeTaskConfigResponseBodyDataDataDataConfig `json:"DataConfig,omitempty" xml:"DataConfig,omitempty" type:"Struct"`
	// example:
	//
	// 100
	FinishRate *float64 `json:"FinishRate,omitempty" xml:"FinishRate,omitempty"`
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0
	ManualReview *int32 `json:"ManualReview,omitempty" xml:"ManualReview,omitempty"`
	// example:
	//
	// cdae396590b*****ec40f3476e274fc
	ModeCustomizationId *string `json:"ModeCustomizationId,omitempty" xml:"ModeCustomizationId,omitempty"`
	ModelName           *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0
	NumberExecuting *int32 `json:"NumberExecuting,omitempty" xml:"NumberExecuting,omitempty"`
	// example:
	//
	// 0
	NumberFail *int32 `json:"NumberFail,omitempty" xml:"NumberFail,omitempty"`
	// example:
	//
	// 1000
	NumberSuccess *int32 `json:"NumberSuccess,omitempty" xml:"NumberSuccess,omitempty"`
	// example:
	//
	// 1000
	NumberSum    *int32                                                `json:"NumberSum,omitempty" xml:"NumberSum,omitempty"`
	SchemeIdList *ListSchemeTaskConfigResponseBodyDataDataSchemeIdList `json:"SchemeIdList,omitempty" xml:"SchemeIdList,omitempty" type:"Struct"`
	SchemeList   *ListSchemeTaskConfigResponseBodyDataDataSchemeList   `json:"SchemeList,omitempty" xml:"SchemeList,omitempty" type:"Struct"`
	// example:
	//
	// 123
	SchemeTaskConfigId *int64 `json:"SchemeTaskConfigId,omitempty" xml:"SchemeTaskConfigId,omitempty"`
	// example:
	//
	// 2
	SourceDataType *int32 `json:"SourceDataType,omitempty" xml:"SourceDataType,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1650418039000
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1
	UpdateUser *int64 `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
	// example:
	//
	// 1
	UserGroup *string `json:"UserGroup,omitempty" xml:"UserGroup,omitempty"`
	// example:
	//
	// 9f90b3efa2****0a49acec226eafc17
	VocabId   *string `json:"VocabId,omitempty" xml:"VocabId,omitempty"`
	VocabName *string `json:"VocabName,omitempty" xml:"VocabName,omitempty"`
}

func (s ListSchemeTaskConfigResponseBodyDataData) String() string {
	return dara.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetAsrTaskPriority() *int32 {
	return s.AsrTaskPriority
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetAsrVersion() *int32 {
	return s.AsrVersion
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetAssignType() *int32 {
	return s.AssignType
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetCreateUser() *int64 {
	return s.CreateUser
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetDataConfig() *ListSchemeTaskConfigResponseBodyDataDataDataConfig {
	return s.DataConfig
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetFinishRate() *float64 {
	return s.FinishRate
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetId() *int64 {
	return s.Id
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetManualReview() *int32 {
	return s.ManualReview
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetModeCustomizationId() *string {
	return s.ModeCustomizationId
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetModelName() *string {
	return s.ModelName
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetName() *string {
	return s.Name
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetNumberExecuting() *int32 {
	return s.NumberExecuting
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetNumberFail() *int32 {
	return s.NumberFail
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetNumberSuccess() *int32 {
	return s.NumberSuccess
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetNumberSum() *int32 {
	return s.NumberSum
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetSchemeIdList() *ListSchemeTaskConfigResponseBodyDataDataSchemeIdList {
	return s.SchemeIdList
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetSchemeList() *ListSchemeTaskConfigResponseBodyDataDataSchemeList {
	return s.SchemeList
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetSchemeTaskConfigId() *int64 {
	return s.SchemeTaskConfigId
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetSourceDataType() *int32 {
	return s.SourceDataType
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetStatus() *int32 {
	return s.Status
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetType() *int32 {
	return s.Type
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetUpdateUser() *int64 {
	return s.UpdateUser
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetUserGroup() *string {
	return s.UserGroup
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetVocabId() *string {
	return s.VocabId
}

func (s *ListSchemeTaskConfigResponseBodyDataData) GetVocabName() *string {
	return s.VocabName
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetAsrTaskPriority(v int32) *ListSchemeTaskConfigResponseBodyDataData {
	s.AsrTaskPriority = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetAsrVersion(v int32) *ListSchemeTaskConfigResponseBodyDataData {
	s.AsrVersion = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetAssignType(v int32) *ListSchemeTaskConfigResponseBodyDataData {
	s.AssignType = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetCreateTime(v string) *ListSchemeTaskConfigResponseBodyDataData {
	s.CreateTime = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetCreateUser(v int64) *ListSchemeTaskConfigResponseBodyDataData {
	s.CreateUser = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetDataConfig(v *ListSchemeTaskConfigResponseBodyDataDataDataConfig) *ListSchemeTaskConfigResponseBodyDataData {
	s.DataConfig = v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetFinishRate(v float64) *ListSchemeTaskConfigResponseBodyDataData {
	s.FinishRate = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetId(v int64) *ListSchemeTaskConfigResponseBodyDataData {
	s.Id = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetManualReview(v int32) *ListSchemeTaskConfigResponseBodyDataData {
	s.ManualReview = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetModeCustomizationId(v string) *ListSchemeTaskConfigResponseBodyDataData {
	s.ModeCustomizationId = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetModelName(v string) *ListSchemeTaskConfigResponseBodyDataData {
	s.ModelName = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetName(v string) *ListSchemeTaskConfigResponseBodyDataData {
	s.Name = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetNumberExecuting(v int32) *ListSchemeTaskConfigResponseBodyDataData {
	s.NumberExecuting = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetNumberFail(v int32) *ListSchemeTaskConfigResponseBodyDataData {
	s.NumberFail = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetNumberSuccess(v int32) *ListSchemeTaskConfigResponseBodyDataData {
	s.NumberSuccess = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetNumberSum(v int32) *ListSchemeTaskConfigResponseBodyDataData {
	s.NumberSum = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetSchemeIdList(v *ListSchemeTaskConfigResponseBodyDataDataSchemeIdList) *ListSchemeTaskConfigResponseBodyDataData {
	s.SchemeIdList = v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetSchemeList(v *ListSchemeTaskConfigResponseBodyDataDataSchemeList) *ListSchemeTaskConfigResponseBodyDataData {
	s.SchemeList = v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetSchemeTaskConfigId(v int64) *ListSchemeTaskConfigResponseBodyDataData {
	s.SchemeTaskConfigId = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetSourceDataType(v int32) *ListSchemeTaskConfigResponseBodyDataData {
	s.SourceDataType = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetStatus(v int32) *ListSchemeTaskConfigResponseBodyDataData {
	s.Status = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetType(v int32) *ListSchemeTaskConfigResponseBodyDataData {
	s.Type = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetUpdateTime(v string) *ListSchemeTaskConfigResponseBodyDataData {
	s.UpdateTime = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetUpdateUser(v int64) *ListSchemeTaskConfigResponseBodyDataData {
	s.UpdateUser = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetUserGroup(v string) *ListSchemeTaskConfigResponseBodyDataData {
	s.UserGroup = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetVocabId(v string) *ListSchemeTaskConfigResponseBodyDataData {
	s.VocabId = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) SetVocabName(v string) *ListSchemeTaskConfigResponseBodyDataData {
	s.VocabName = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataData) Validate() error {
	if s.DataConfig != nil {
		if err := s.DataConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SchemeIdList != nil {
		if err := s.SchemeIdList.Validate(); err != nil {
			return err
		}
	}
	if s.SchemeList != nil {
		if err := s.SchemeList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSchemeTaskConfigResponseBodyDataDataDataConfig struct {
	AssignConfigs *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigs `json:"AssignConfigs,omitempty" xml:"AssignConfigs,omitempty" type:"Struct"`
	// example:
	//
	// []
	DataSets *string `json:"DataSets,omitempty" xml:"DataSets,omitempty"`
	// example:
	//
	// 0
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// {}
	ResultParam *string `json:"ResultParam,omitempty" xml:"ResultParam,omitempty"`
}

func (s ListSchemeTaskConfigResponseBodyDataDataDataConfig) String() string {
	return dara.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBodyDataDataDataConfig) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfig) GetAssignConfigs() *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigs {
	return s.AssignConfigs
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfig) GetDataSets() *string {
	return s.DataSets
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfig) GetIndex() *int32 {
	return s.Index
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfig) GetResultParam() *string {
	return s.ResultParam
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfig) SetAssignConfigs(v *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigs) *ListSchemeTaskConfigResponseBodyDataDataDataConfig {
	s.AssignConfigs = v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfig) SetDataSets(v string) *ListSchemeTaskConfigResponseBodyDataDataDataConfig {
	s.DataSets = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfig) SetIndex(v int32) *ListSchemeTaskConfigResponseBodyDataDataDataConfig {
	s.Index = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfig) SetResultParam(v string) *ListSchemeTaskConfigResponseBodyDataDataDataConfig {
	s.ResultParam = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfig) Validate() error {
	if s.AssignConfigs != nil {
		if err := s.AssignConfigs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigs struct {
	AssignConfig []*ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfig `json:"AssignConfig,omitempty" xml:"AssignConfig,omitempty" type:"Repeated"`
}

func (s ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigs) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigs) GetAssignConfig() []*ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfig {
	return s.AssignConfig
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigs) SetAssignConfig(v []*ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfig) *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigs {
	s.AssignConfig = v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigs) Validate() error {
	if s.AssignConfig != nil {
		for _, item := range s.AssignConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfig struct {
	AssignConfigContests *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContests `json:"AssignConfigContests,omitempty" xml:"AssignConfigContests,omitempty" type:"Struct"`
}

func (s ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfig) String() string {
	return dara.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfig) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfig) GetAssignConfigContests() *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContests {
	return s.AssignConfigContests
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfig) SetAssignConfigContests(v *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContests) *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfig {
	s.AssignConfigContests = v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfig) Validate() error {
	if s.AssignConfigContests != nil {
		if err := s.AssignConfigContests.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContests struct {
	AssignConfigContest []*ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest `json:"AssignConfigContest,omitempty" xml:"AssignConfigContest,omitempty" type:"Repeated"`
}

func (s ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContests) String() string {
	return dara.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContests) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContests) GetAssignConfigContest() []*ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest {
	return s.AssignConfigContest
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContests) SetAssignConfigContest(v []*ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest) *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContests {
	s.AssignConfigContest = v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContests) Validate() error {
	if s.AssignConfigContest != nil {
		for _, item := range s.AssignConfigContest {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest struct {
	// example:
	//
	// 3
	DataType   *int32                                                                                                                        `json:"DataType,omitempty" xml:"DataType,omitempty"`
	ListObject *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContestListObject `json:"ListObject,omitempty" xml:"ListObject,omitempty" type:"Struct"`
	// example:
	//
	// callStartTime
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 4
	Symbol *int32 `json:"Symbol,omitempty" xml:"Symbol,omitempty"`
	// example:
	//
	// {\\"start\\":\\"2022-09-01 00:00:00\\",\\"end\\":\\"2022-09-30 00:00:00\\"}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest) String() string {
	return dara.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest) GetDataType() *int32 {
	return s.DataType
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest) GetListObject() *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContestListObject {
	return s.ListObject
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest) GetName() *string {
	return s.Name
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest) GetSymbol() *int32 {
	return s.Symbol
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest) GetValue() *string {
	return s.Value
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest) SetDataType(v int32) *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest {
	s.DataType = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest) SetListObject(v *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContestListObject) *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest {
	s.ListObject = v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest) SetName(v string) *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest {
	s.Name = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest) SetSymbol(v int32) *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest {
	s.Symbol = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest) SetValue(v string) *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest {
	s.Value = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContest) Validate() error {
	if s.ListObject != nil {
		if err := s.ListObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContestListObject struct {
	ListObject []interface{} `json:"ListObject,omitempty" xml:"ListObject,omitempty" type:"Repeated"`
}

func (s ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContestListObject) String() string {
	return dara.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContestListObject) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContestListObject) GetListObject() []interface{} {
	return s.ListObject
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContestListObject) SetListObject(v []interface{}) *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContestListObject {
	s.ListObject = v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataDataDataConfigAssignConfigsAssignConfigAssignConfigContestsAssignConfigContestListObject) Validate() error {
	return dara.Validate(s)
}

type ListSchemeTaskConfigResponseBodyDataDataSchemeIdList struct {
	SchemeIdList []*int64 `json:"SchemeIdList,omitempty" xml:"SchemeIdList,omitempty" type:"Repeated"`
}

func (s ListSchemeTaskConfigResponseBodyDataDataSchemeIdList) String() string {
	return dara.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBodyDataDataSchemeIdList) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBodyDataDataSchemeIdList) GetSchemeIdList() []*int64 {
	return s.SchemeIdList
}

func (s *ListSchemeTaskConfigResponseBodyDataDataSchemeIdList) SetSchemeIdList(v []*int64) *ListSchemeTaskConfigResponseBodyDataDataSchemeIdList {
	s.SchemeIdList = v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataDataSchemeIdList) Validate() error {
	return dara.Validate(s)
}

type ListSchemeTaskConfigResponseBodyDataDataSchemeList struct {
	SchemeList []*ListSchemeTaskConfigResponseBodyDataDataSchemeListSchemeList `json:"SchemeList,omitempty" xml:"SchemeList,omitempty" type:"Repeated"`
}

func (s ListSchemeTaskConfigResponseBodyDataDataSchemeList) String() string {
	return dara.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBodyDataDataSchemeList) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBodyDataDataSchemeList) GetSchemeList() []*ListSchemeTaskConfigResponseBodyDataDataSchemeListSchemeList {
	return s.SchemeList
}

func (s *ListSchemeTaskConfigResponseBodyDataDataSchemeList) SetSchemeList(v []*ListSchemeTaskConfigResponseBodyDataDataSchemeListSchemeList) *ListSchemeTaskConfigResponseBodyDataDataSchemeList {
	s.SchemeList = v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataDataSchemeList) Validate() error {
	if s.SchemeList != nil {
		for _, item := range s.SchemeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSchemeTaskConfigResponseBodyDataDataSchemeListSchemeList struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 158
	SchemeId *int64 `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
}

func (s ListSchemeTaskConfigResponseBodyDataDataSchemeListSchemeList) String() string {
	return dara.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBodyDataDataSchemeListSchemeList) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBodyDataDataSchemeListSchemeList) GetName() *string {
	return s.Name
}

func (s *ListSchemeTaskConfigResponseBodyDataDataSchemeListSchemeList) GetSchemeId() *int64 {
	return s.SchemeId
}

func (s *ListSchemeTaskConfigResponseBodyDataDataSchemeListSchemeList) SetName(v string) *ListSchemeTaskConfigResponseBodyDataDataSchemeListSchemeList {
	s.Name = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataDataSchemeListSchemeList) SetSchemeId(v int64) *ListSchemeTaskConfigResponseBodyDataDataSchemeListSchemeList {
	s.SchemeId = &v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyDataDataSchemeListSchemeList) Validate() error {
	return dara.Validate(s)
}

type ListSchemeTaskConfigResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s ListSchemeTaskConfigResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s ListSchemeTaskConfigResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigResponseBodyMessages) GetMessage() []*string {
	return s.Message
}

func (s *ListSchemeTaskConfigResponseBodyMessages) SetMessage(v []*string) *ListSchemeTaskConfigResponseBodyMessages {
	s.Message = v
	return s
}

func (s *ListSchemeTaskConfigResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}
