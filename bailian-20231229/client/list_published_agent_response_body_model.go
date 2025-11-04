// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishedAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPublishedAgentResponseBody
	GetCode() *string
	SetData(v *ListPublishedAgentResponseBodyData) *ListPublishedAgentResponseBody
	GetData() *ListPublishedAgentResponseBodyData
	SetHttpStatusCode(v int32) *ListPublishedAgentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListPublishedAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPublishedAgentResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ListPublishedAgentResponseBody
	GetSuccess() *string
}

type ListPublishedAgentResponseBody struct {
	Code           *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data           *ListPublishedAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	HttpStatusCode *int32                              `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string                             `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *string                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListPublishedAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAgentResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPublishedAgentResponseBody) GetData() *ListPublishedAgentResponseBodyData {
	return s.Data
}

func (s *ListPublishedAgentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListPublishedAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPublishedAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPublishedAgentResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ListPublishedAgentResponseBody) SetCode(v string) *ListPublishedAgentResponseBody {
	s.Code = &v
	return s
}

func (s *ListPublishedAgentResponseBody) SetData(v *ListPublishedAgentResponseBodyData) *ListPublishedAgentResponseBody {
	s.Data = v
	return s
}

func (s *ListPublishedAgentResponseBody) SetHttpStatusCode(v int32) *ListPublishedAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPublishedAgentResponseBody) SetMessage(v string) *ListPublishedAgentResponseBody {
	s.Message = &v
	return s
}

func (s *ListPublishedAgentResponseBody) SetRequestId(v string) *ListPublishedAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublishedAgentResponseBody) SetSuccess(v string) *ListPublishedAgentResponseBody {
	s.Success = &v
	return s
}

func (s *ListPublishedAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPublishedAgentResponseBodyData struct {
	List     []*ListPublishedAgentResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNo   *int32                                    `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
	PageSize *int32                                    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                                    `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListPublishedAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentResponseBodyData) GetList() []*ListPublishedAgentResponseBodyDataList {
	return s.List
}

func (s *ListPublishedAgentResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListPublishedAgentResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPublishedAgentResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListPublishedAgentResponseBodyData) SetList(v []*ListPublishedAgentResponseBodyDataList) *ListPublishedAgentResponseBodyData {
	s.List = v
	return s
}

func (s *ListPublishedAgentResponseBodyData) SetPageNo(v int32) *ListPublishedAgentResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *ListPublishedAgentResponseBodyData) SetPageSize(v int32) *ListPublishedAgentResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListPublishedAgentResponseBodyData) SetTotal(v int32) *ListPublishedAgentResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListPublishedAgentResponseBodyData) Validate() error {
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

type ListPublishedAgentResponseBodyDataList struct {
	ApplicationConfig *ListPublishedAgentResponseBodyDataListApplicationConfig `json:"applicationConfig,omitempty" xml:"applicationConfig,omitempty" type:"Struct"`
	Code              *string                                                  `json:"code,omitempty" xml:"code,omitempty"`
	Instructions      *string                                                  `json:"instructions,omitempty" xml:"instructions,omitempty"`
	ModelId           *string                                                  `json:"modelId,omitempty" xml:"modelId,omitempty"`
	Name              *string                                                  `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListPublishedAgentResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAgentResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentResponseBodyDataList) GetApplicationConfig() *ListPublishedAgentResponseBodyDataListApplicationConfig {
	return s.ApplicationConfig
}

func (s *ListPublishedAgentResponseBodyDataList) GetCode() *string {
	return s.Code
}

func (s *ListPublishedAgentResponseBodyDataList) GetInstructions() *string {
	return s.Instructions
}

func (s *ListPublishedAgentResponseBodyDataList) GetModelId() *string {
	return s.ModelId
}

func (s *ListPublishedAgentResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *ListPublishedAgentResponseBodyDataList) SetApplicationConfig(v *ListPublishedAgentResponseBodyDataListApplicationConfig) *ListPublishedAgentResponseBodyDataList {
	s.ApplicationConfig = v
	return s
}

func (s *ListPublishedAgentResponseBodyDataList) SetCode(v string) *ListPublishedAgentResponseBodyDataList {
	s.Code = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataList) SetInstructions(v string) *ListPublishedAgentResponseBodyDataList {
	s.Instructions = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataList) SetModelId(v string) *ListPublishedAgentResponseBodyDataList {
	s.ModelId = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataList) SetName(v string) *ListPublishedAgentResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataList) Validate() error {
	if s.ApplicationConfig != nil {
		if err := s.ApplicationConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPublishedAgentResponseBodyDataListApplicationConfig struct {
	HistoryConfig  *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig  `json:"historyConfig,omitempty" xml:"historyConfig,omitempty" type:"Struct"`
	LongTermMemory *ListPublishedAgentResponseBodyDataListApplicationConfigLongTermMemory `json:"longTermMemory,omitempty" xml:"longTermMemory,omitempty" type:"Struct"`
	Parameters     *ListPublishedAgentResponseBodyDataListApplicationConfigParameters     `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
	RagConfig      *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig      `json:"ragConfig,omitempty" xml:"ragConfig,omitempty" type:"Struct"`
	Security       *ListPublishedAgentResponseBodyDataListApplicationConfigSecurity       `json:"security,omitempty" xml:"security,omitempty" type:"Struct"`
	Tools          []*ListPublishedAgentResponseBodyDataListApplicationConfigTools        `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
	WorkFlows      []*ListPublishedAgentResponseBodyDataListApplicationConfigWorkFlows    `json:"workFlows,omitempty" xml:"workFlows,omitempty" type:"Repeated"`
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfig) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfig) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfig) GetHistoryConfig() *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig {
	return s.HistoryConfig
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfig) GetLongTermMemory() *ListPublishedAgentResponseBodyDataListApplicationConfigLongTermMemory {
	return s.LongTermMemory
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfig) GetParameters() *ListPublishedAgentResponseBodyDataListApplicationConfigParameters {
	return s.Parameters
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfig) GetRagConfig() *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig {
	return s.RagConfig
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfig) GetSecurity() *ListPublishedAgentResponseBodyDataListApplicationConfigSecurity {
	return s.Security
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfig) GetTools() []*ListPublishedAgentResponseBodyDataListApplicationConfigTools {
	return s.Tools
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfig) GetWorkFlows() []*ListPublishedAgentResponseBodyDataListApplicationConfigWorkFlows {
	return s.WorkFlows
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfig) SetHistoryConfig(v *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig) *ListPublishedAgentResponseBodyDataListApplicationConfig {
	s.HistoryConfig = v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfig) SetLongTermMemory(v *ListPublishedAgentResponseBodyDataListApplicationConfigLongTermMemory) *ListPublishedAgentResponseBodyDataListApplicationConfig {
	s.LongTermMemory = v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfig) SetParameters(v *ListPublishedAgentResponseBodyDataListApplicationConfigParameters) *ListPublishedAgentResponseBodyDataListApplicationConfig {
	s.Parameters = v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfig) SetRagConfig(v *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig) *ListPublishedAgentResponseBodyDataListApplicationConfig {
	s.RagConfig = v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfig) SetSecurity(v *ListPublishedAgentResponseBodyDataListApplicationConfigSecurity) *ListPublishedAgentResponseBodyDataListApplicationConfig {
	s.Security = v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfig) SetTools(v []*ListPublishedAgentResponseBodyDataListApplicationConfigTools) *ListPublishedAgentResponseBodyDataListApplicationConfig {
	s.Tools = v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfig) SetWorkFlows(v []*ListPublishedAgentResponseBodyDataListApplicationConfigWorkFlows) *ListPublishedAgentResponseBodyDataListApplicationConfig {
	s.WorkFlows = v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfig) Validate() error {
	if s.HistoryConfig != nil {
		if err := s.HistoryConfig.Validate(); err != nil {
			return err
		}
	}
	if s.LongTermMemory != nil {
		if err := s.LongTermMemory.Validate(); err != nil {
			return err
		}
	}
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	if s.RagConfig != nil {
		if err := s.RagConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Security != nil {
		if err := s.Security.Validate(); err != nil {
			return err
		}
	}
	if s.Tools != nil {
		for _, item := range s.Tools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WorkFlows != nil {
		for _, item := range s.WorkFlows {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig struct {
	EnableAdbRecord *bool   `json:"enableAdbRecord,omitempty" xml:"enableAdbRecord,omitempty"`
	EnableRecord    *bool   `json:"enableRecord,omitempty" xml:"enableRecord,omitempty"`
	InstanceId      *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Region          *string `json:"region,omitempty" xml:"region,omitempty"`
	StoreCode       *string `json:"storeCode,omitempty" xml:"storeCode,omitempty"`
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig) GetEnableAdbRecord() *bool {
	return s.EnableAdbRecord
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig) GetEnableRecord() *bool {
	return s.EnableRecord
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig) GetRegion() *string {
	return s.Region
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig) GetStoreCode() *string {
	return s.StoreCode
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig) SetEnableAdbRecord(v bool) *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig {
	s.EnableAdbRecord = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig) SetEnableRecord(v bool) *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig {
	s.EnableRecord = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig) SetInstanceId(v string) *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig {
	s.InstanceId = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig) SetRegion(v string) *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig {
	s.Region = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig) SetStoreCode(v string) *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig {
	s.StoreCode = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig) Validate() error {
	return dara.Validate(s)
}

type ListPublishedAgentResponseBodyDataListApplicationConfigLongTermMemory struct {
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigLongTermMemory) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigLongTermMemory) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigLongTermMemory) GetEnable() *bool {
	return s.Enable
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigLongTermMemory) SetEnable(v bool) *ListPublishedAgentResponseBodyDataListApplicationConfigLongTermMemory {
	s.Enable = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigLongTermMemory) Validate() error {
	return dara.Validate(s)
}

type ListPublishedAgentResponseBodyDataListApplicationConfigParameters struct {
	DialogRound *int32   `json:"dialogRound,omitempty" xml:"dialogRound,omitempty"`
	MaxTokens   *int32   `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	Temperature *float64 `json:"temperature,omitempty" xml:"temperature,omitempty"`
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigParameters) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigParameters) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigParameters) GetDialogRound() *int32 {
	return s.DialogRound
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigParameters) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigParameters) GetTemperature() *float64 {
	return s.Temperature
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigParameters) SetDialogRound(v int32) *ListPublishedAgentResponseBodyDataListApplicationConfigParameters {
	s.DialogRound = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigParameters) SetMaxTokens(v int32) *ListPublishedAgentResponseBodyDataListApplicationConfigParameters {
	s.MaxTokens = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigParameters) SetTemperature(v float64) *ListPublishedAgentResponseBodyDataListApplicationConfigParameters {
	s.Temperature = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigParameters) Validate() error {
	return dara.Validate(s)
}

type ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig struct {
	EnableCitation        *bool     `json:"enableCitation,omitempty" xml:"enableCitation,omitempty"`
	EnableSearch          *bool     `json:"enableSearch,omitempty" xml:"enableSearch,omitempty"`
	KnowledgeBaseCodeList []*string `json:"knowledgeBaseCodeList,omitempty" xml:"knowledgeBaseCodeList,omitempty" type:"Repeated"`
	TopK                  *int32    `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig) GetEnableCitation() *bool {
	return s.EnableCitation
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig) GetEnableSearch() *bool {
	return s.EnableSearch
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig) GetKnowledgeBaseCodeList() []*string {
	return s.KnowledgeBaseCodeList
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig) GetTopK() *int32 {
	return s.TopK
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig) SetEnableCitation(v bool) *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig {
	s.EnableCitation = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig) SetEnableSearch(v bool) *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig {
	s.EnableSearch = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig) SetKnowledgeBaseCodeList(v []*string) *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig {
	s.KnowledgeBaseCodeList = v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig) SetTopK(v int32) *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig {
	s.TopK = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig) Validate() error {
	return dara.Validate(s)
}

type ListPublishedAgentResponseBodyDataListApplicationConfigSecurity struct {
	ProcessingStrategy *string `json:"processingStrategy,omitempty" xml:"processingStrategy,omitempty"`
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigSecurity) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigSecurity) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigSecurity) GetProcessingStrategy() *string {
	return s.ProcessingStrategy
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigSecurity) SetProcessingStrategy(v string) *ListPublishedAgentResponseBodyDataListApplicationConfigSecurity {
	s.ProcessingStrategy = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigSecurity) Validate() error {
	return dara.Validate(s)
}

type ListPublishedAgentResponseBodyDataListApplicationConfigTools struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigTools) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigTools) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigTools) GetType() *string {
	return s.Type
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigTools) SetType(v string) *ListPublishedAgentResponseBodyDataListApplicationConfigTools {
	s.Type = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigTools) Validate() error {
	return dara.Validate(s)
}

type ListPublishedAgentResponseBodyDataListApplicationConfigWorkFlows struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigWorkFlows) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigWorkFlows) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigWorkFlows) GetType() *string {
	return s.Type
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigWorkFlows) SetType(v string) *ListPublishedAgentResponseBodyDataListApplicationConfigWorkFlows {
	s.Type = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigWorkFlows) Validate() error {
	return dara.Validate(s)
}
