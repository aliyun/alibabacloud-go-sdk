// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPublishedAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPublishedAgentResponseBody
	GetCode() *string
	SetData(v *GetPublishedAgentResponseBodyData) *GetPublishedAgentResponseBody
	GetData() *GetPublishedAgentResponseBodyData
	SetHttpStatusCode(v int32) *GetPublishedAgentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPublishedAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPublishedAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPublishedAgentResponseBody
	GetSuccess() *bool
}

type GetPublishedAgentResponseBody struct {
	Code           *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data           *GetPublishedAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	HttpStatusCode *int32                             `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string                            `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool                              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetPublishedAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPublishedAgentResponseBody) GoString() string {
	return s.String()
}

func (s *GetPublishedAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPublishedAgentResponseBody) GetData() *GetPublishedAgentResponseBodyData {
	return s.Data
}

func (s *GetPublishedAgentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPublishedAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPublishedAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPublishedAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPublishedAgentResponseBody) SetCode(v string) *GetPublishedAgentResponseBody {
	s.Code = &v
	return s
}

func (s *GetPublishedAgentResponseBody) SetData(v *GetPublishedAgentResponseBodyData) *GetPublishedAgentResponseBody {
	s.Data = v
	return s
}

func (s *GetPublishedAgentResponseBody) SetHttpStatusCode(v int32) *GetPublishedAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPublishedAgentResponseBody) SetMessage(v string) *GetPublishedAgentResponseBody {
	s.Message = &v
	return s
}

func (s *GetPublishedAgentResponseBody) SetRequestId(v string) *GetPublishedAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPublishedAgentResponseBody) SetSuccess(v bool) *GetPublishedAgentResponseBody {
	s.Success = &v
	return s
}

func (s *GetPublishedAgentResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPublishedAgentResponseBodyData struct {
	ApplicationConfig *GetPublishedAgentResponseBodyDataApplicationConfig `json:"applicationConfig,omitempty" xml:"applicationConfig,omitempty" type:"Struct"`
	Code              *string                                             `json:"code,omitempty" xml:"code,omitempty"`
	Instructions      *string                                             `json:"instructions,omitempty" xml:"instructions,omitempty"`
	ModelId           *string                                             `json:"modelId,omitempty" xml:"modelId,omitempty"`
	Name              *string                                             `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetPublishedAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPublishedAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPublishedAgentResponseBodyData) GetApplicationConfig() *GetPublishedAgentResponseBodyDataApplicationConfig {
	return s.ApplicationConfig
}

func (s *GetPublishedAgentResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *GetPublishedAgentResponseBodyData) GetInstructions() *string {
	return s.Instructions
}

func (s *GetPublishedAgentResponseBodyData) GetModelId() *string {
	return s.ModelId
}

func (s *GetPublishedAgentResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetPublishedAgentResponseBodyData) SetApplicationConfig(v *GetPublishedAgentResponseBodyDataApplicationConfig) *GetPublishedAgentResponseBodyData {
	s.ApplicationConfig = v
	return s
}

func (s *GetPublishedAgentResponseBodyData) SetCode(v string) *GetPublishedAgentResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetPublishedAgentResponseBodyData) SetInstructions(v string) *GetPublishedAgentResponseBodyData {
	s.Instructions = &v
	return s
}

func (s *GetPublishedAgentResponseBodyData) SetModelId(v string) *GetPublishedAgentResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *GetPublishedAgentResponseBodyData) SetName(v string) *GetPublishedAgentResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetPublishedAgentResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetPublishedAgentResponseBodyDataApplicationConfig struct {
	HistoryConfig  *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig  `json:"historyConfig,omitempty" xml:"historyConfig,omitempty" type:"Struct"`
	LongTermMemory *GetPublishedAgentResponseBodyDataApplicationConfigLongTermMemory `json:"longTermMemory,omitempty" xml:"longTermMemory,omitempty" type:"Struct"`
	Parameters     *GetPublishedAgentResponseBodyDataApplicationConfigParameters     `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
	RagConfig      *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig      `json:"ragConfig,omitempty" xml:"ragConfig,omitempty" type:"Struct"`
	Security       *GetPublishedAgentResponseBodyDataApplicationConfigSecurity       `json:"security,omitempty" xml:"security,omitempty" type:"Struct"`
	Tools          []*GetPublishedAgentResponseBodyDataApplicationConfigTools        `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
	WorkFlows      []*GetPublishedAgentResponseBodyDataApplicationConfigWorkFlows    `json:"workFlows,omitempty" xml:"workFlows,omitempty" type:"Repeated"`
}

func (s GetPublishedAgentResponseBodyDataApplicationConfig) String() string {
	return dara.Prettify(s)
}

func (s GetPublishedAgentResponseBodyDataApplicationConfig) GoString() string {
	return s.String()
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfig) GetHistoryConfig() *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig {
	return s.HistoryConfig
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfig) GetLongTermMemory() *GetPublishedAgentResponseBodyDataApplicationConfigLongTermMemory {
	return s.LongTermMemory
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfig) GetParameters() *GetPublishedAgentResponseBodyDataApplicationConfigParameters {
	return s.Parameters
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfig) GetRagConfig() *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig {
	return s.RagConfig
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfig) GetSecurity() *GetPublishedAgentResponseBodyDataApplicationConfigSecurity {
	return s.Security
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfig) GetTools() []*GetPublishedAgentResponseBodyDataApplicationConfigTools {
	return s.Tools
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfig) GetWorkFlows() []*GetPublishedAgentResponseBodyDataApplicationConfigWorkFlows {
	return s.WorkFlows
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfig) SetHistoryConfig(v *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig) *GetPublishedAgentResponseBodyDataApplicationConfig {
	s.HistoryConfig = v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfig) SetLongTermMemory(v *GetPublishedAgentResponseBodyDataApplicationConfigLongTermMemory) *GetPublishedAgentResponseBodyDataApplicationConfig {
	s.LongTermMemory = v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfig) SetParameters(v *GetPublishedAgentResponseBodyDataApplicationConfigParameters) *GetPublishedAgentResponseBodyDataApplicationConfig {
	s.Parameters = v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfig) SetRagConfig(v *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig) *GetPublishedAgentResponseBodyDataApplicationConfig {
	s.RagConfig = v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfig) SetSecurity(v *GetPublishedAgentResponseBodyDataApplicationConfigSecurity) *GetPublishedAgentResponseBodyDataApplicationConfig {
	s.Security = v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfig) SetTools(v []*GetPublishedAgentResponseBodyDataApplicationConfigTools) *GetPublishedAgentResponseBodyDataApplicationConfig {
	s.Tools = v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfig) SetWorkFlows(v []*GetPublishedAgentResponseBodyDataApplicationConfigWorkFlows) *GetPublishedAgentResponseBodyDataApplicationConfig {
	s.WorkFlows = v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfig) Validate() error {
	return dara.Validate(s)
}

type GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig struct {
	EnableAdbRecord *bool   `json:"enableAdbRecord,omitempty" xml:"enableAdbRecord,omitempty"`
	EnableRecord    *bool   `json:"enableRecord,omitempty" xml:"enableRecord,omitempty"`
	InstanceId      *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Region          *string `json:"region,omitempty" xml:"region,omitempty"`
	StoreCode       *string `json:"storeCode,omitempty" xml:"storeCode,omitempty"`
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig) String() string {
	return dara.Prettify(s)
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig) GoString() string {
	return s.String()
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig) GetEnableAdbRecord() *bool {
	return s.EnableAdbRecord
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig) GetEnableRecord() *bool {
	return s.EnableRecord
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig) GetRegion() *string {
	return s.Region
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig) GetStoreCode() *string {
	return s.StoreCode
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig) SetEnableAdbRecord(v bool) *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig {
	s.EnableAdbRecord = &v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig) SetEnableRecord(v bool) *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig {
	s.EnableRecord = &v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig) SetInstanceId(v string) *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig {
	s.InstanceId = &v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig) SetRegion(v string) *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig {
	s.Region = &v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig) SetStoreCode(v string) *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig {
	s.StoreCode = &v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig) Validate() error {
	return dara.Validate(s)
}

type GetPublishedAgentResponseBodyDataApplicationConfigLongTermMemory struct {
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigLongTermMemory) String() string {
	return dara.Prettify(s)
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigLongTermMemory) GoString() string {
	return s.String()
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigLongTermMemory) GetEnable() *bool {
	return s.Enable
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigLongTermMemory) SetEnable(v bool) *GetPublishedAgentResponseBodyDataApplicationConfigLongTermMemory {
	s.Enable = &v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigLongTermMemory) Validate() error {
	return dara.Validate(s)
}

type GetPublishedAgentResponseBodyDataApplicationConfigParameters struct {
	DialogRound *int32   `json:"dialogRound,omitempty" xml:"dialogRound,omitempty"`
	MaxTokens   *int32   `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	Temperature *float64 `json:"temperature,omitempty" xml:"temperature,omitempty"`
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigParameters) String() string {
	return dara.Prettify(s)
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigParameters) GoString() string {
	return s.String()
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigParameters) GetDialogRound() *int32 {
	return s.DialogRound
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigParameters) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigParameters) GetTemperature() *float64 {
	return s.Temperature
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigParameters) SetDialogRound(v int32) *GetPublishedAgentResponseBodyDataApplicationConfigParameters {
	s.DialogRound = &v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigParameters) SetMaxTokens(v int32) *GetPublishedAgentResponseBodyDataApplicationConfigParameters {
	s.MaxTokens = &v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigParameters) SetTemperature(v float64) *GetPublishedAgentResponseBodyDataApplicationConfigParameters {
	s.Temperature = &v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigParameters) Validate() error {
	return dara.Validate(s)
}

type GetPublishedAgentResponseBodyDataApplicationConfigRagConfig struct {
	EnableCitation        *bool     `json:"enableCitation,omitempty" xml:"enableCitation,omitempty"`
	EnableSearch          *bool     `json:"enableSearch,omitempty" xml:"enableSearch,omitempty"`
	KnowledgeBaseCodeList []*string `json:"knowledgeBaseCodeList,omitempty" xml:"knowledgeBaseCodeList,omitempty" type:"Repeated"`
	TopK                  *int32    `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigRagConfig) String() string {
	return dara.Prettify(s)
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigRagConfig) GoString() string {
	return s.String()
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig) GetEnableCitation() *bool {
	return s.EnableCitation
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig) GetEnableSearch() *bool {
	return s.EnableSearch
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig) GetKnowledgeBaseCodeList() []*string {
	return s.KnowledgeBaseCodeList
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig) GetTopK() *int32 {
	return s.TopK
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig) SetEnableCitation(v bool) *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig {
	s.EnableCitation = &v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig) SetEnableSearch(v bool) *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig {
	s.EnableSearch = &v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig) SetKnowledgeBaseCodeList(v []*string) *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig {
	s.KnowledgeBaseCodeList = v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig) SetTopK(v int32) *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig {
	s.TopK = &v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig) Validate() error {
	return dara.Validate(s)
}

type GetPublishedAgentResponseBodyDataApplicationConfigSecurity struct {
	ProcessingStrategy *string `json:"processingStrategy,omitempty" xml:"processingStrategy,omitempty"`
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigSecurity) String() string {
	return dara.Prettify(s)
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigSecurity) GoString() string {
	return s.String()
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigSecurity) GetProcessingStrategy() *string {
	return s.ProcessingStrategy
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigSecurity) SetProcessingStrategy(v string) *GetPublishedAgentResponseBodyDataApplicationConfigSecurity {
	s.ProcessingStrategy = &v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigSecurity) Validate() error {
	return dara.Validate(s)
}

type GetPublishedAgentResponseBodyDataApplicationConfigTools struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigTools) String() string {
	return dara.Prettify(s)
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigTools) GoString() string {
	return s.String()
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigTools) GetType() *string {
	return s.Type
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigTools) SetType(v string) *GetPublishedAgentResponseBodyDataApplicationConfigTools {
	s.Type = &v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigTools) Validate() error {
	return dara.Validate(s)
}

type GetPublishedAgentResponseBodyDataApplicationConfigWorkFlows struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigWorkFlows) String() string {
	return dara.Prettify(s)
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigWorkFlows) GoString() string {
	return s.String()
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigWorkFlows) GetType() *string {
	return s.Type
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigWorkFlows) SetType(v string) *GetPublishedAgentResponseBodyDataApplicationConfigWorkFlows {
	s.Type = &v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigWorkFlows) Validate() error {
	return dara.Validate(s)
}
