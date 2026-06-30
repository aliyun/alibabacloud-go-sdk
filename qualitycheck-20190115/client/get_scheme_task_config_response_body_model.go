// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSchemeTaskConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSchemeTaskConfigResponseBody
	GetCode() *string
	SetData(v *GetSchemeTaskConfigResponseBodyData) *GetSchemeTaskConfigResponseBody
	GetData() *GetSchemeTaskConfigResponseBodyData
	SetHttpStatusCode(v string) *GetSchemeTaskConfigResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *GetSchemeTaskConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSchemeTaskConfigResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetSchemeTaskConfigResponseBody
	GetSuccess() *string
}

type GetSchemeTaskConfigResponseBody struct {
	// Result code. **200*	- means success.
	//
	// > Any other value means failure. The caller can use this field to identify the cause.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response data. See the additional notes below.
	Data *GetSchemeTaskConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Error details if the request failed. Returns successful if the request succeeded.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 3CEA0495-341B-4482-9AD9-8191EF4***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded. Use this field to check the result: true means success, false or null means failure.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSchemeTaskConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSchemeTaskConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetSchemeTaskConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSchemeTaskConfigResponseBody) GetData() *GetSchemeTaskConfigResponseBodyData {
	return s.Data
}

func (s *GetSchemeTaskConfigResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GetSchemeTaskConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSchemeTaskConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSchemeTaskConfigResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetSchemeTaskConfigResponseBody) SetCode(v string) *GetSchemeTaskConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetSchemeTaskConfigResponseBody) SetData(v *GetSchemeTaskConfigResponseBodyData) *GetSchemeTaskConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetSchemeTaskConfigResponseBody) SetHttpStatusCode(v string) *GetSchemeTaskConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSchemeTaskConfigResponseBody) SetMessage(v string) *GetSchemeTaskConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetSchemeTaskConfigResponseBody) SetRequestId(v string) *GetSchemeTaskConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSchemeTaskConfigResponseBody) SetSuccess(v string) *GetSchemeTaskConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetSchemeTaskConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSchemeTaskConfigResponseBodyData struct {
	// Task priority:
	//
	// - 0 (low)
	//
	// - 1 (medium)
	//
	// - 2 (high)
	//
	// example:
	//
	// 2
	AsrTaskPriority *int32 `json:"AsrTaskPriority,omitempty" xml:"AsrTaskPriority,omitempty"`
	// Assignment type
	//
	// example:
	//
	// 0
	AssignType *int32 `json:"AssignType,omitempty" xml:"AssignType,omitempty"`
	// Data configuration
	DataConfig *GetSchemeTaskConfigResponseBodyDataDataConfig `json:"DataConfig,omitempty" xml:"DataConfig,omitempty" type:"Struct"`
	// Quality inspection task ID
	//
	// example:
	//
	// 3
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Manual review
	//
	// example:
	//
	// 0
	ManualReview *int32 `json:"ManualReview,omitempty" xml:"ManualReview,omitempty"`
	// Language model ID
	//
	// example:
	//
	// cdae396590b*****ec40f3476e274fc
	ModeCustomizationId *string `json:"ModeCustomizationId,omitempty" xml:"ModeCustomizationId,omitempty"`
	// Language model name
	//
	// example:
	//
	// 自定义模型
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// Quality inspection task name
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Quality inspection scheme IDs
	SchemeIdList []*int64 `json:"SchemeIdList,omitempty" xml:"SchemeIdList,omitempty" type:"Repeated"`
	// Quality inspection schemes
	SchemeList []*GetSchemeTaskConfigResponseBodyDataSchemeList `json:"SchemeList,omitempty" xml:"SchemeList,omitempty" type:"Repeated"`
	// Quality inspection task ID
	//
	// example:
	//
	// 123
	SchemeTaskConfigId *int64 `json:"SchemeTaskConfigId,omitempty" xml:"SchemeTaskConfigId,omitempty"`
	// Quality inspection result type:
	//
	// - 1: offline voice
	//
	// - 2: offline text
	//
	// - 3: real-time voice
	//
	// - 4: real-time text
	//
	// - 5: contact center secondary quality inspection
	//
	// - 51: call center voice secondary quality inspection
	//
	// - 52: call center text secondary quality inspection
	//
	// - 11: dataset voice
	//
	// - 12: dataset text
	//
	// example:
	//
	// 1
	SourceDataType *string `json:"SourceDataType,omitempty" xml:"SourceDataType,omitempty"`
	// Enable status. Valid values: 0 (disabled) or 1 (enabled)
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetSchemeTaskConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSchemeTaskConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSchemeTaskConfigResponseBodyData) GetAsrTaskPriority() *int32 {
	return s.AsrTaskPriority
}

func (s *GetSchemeTaskConfigResponseBodyData) GetAssignType() *int32 {
	return s.AssignType
}

func (s *GetSchemeTaskConfigResponseBodyData) GetDataConfig() *GetSchemeTaskConfigResponseBodyDataDataConfig {
	return s.DataConfig
}

func (s *GetSchemeTaskConfigResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetSchemeTaskConfigResponseBodyData) GetManualReview() *int32 {
	return s.ManualReview
}

func (s *GetSchemeTaskConfigResponseBodyData) GetModeCustomizationId() *string {
	return s.ModeCustomizationId
}

func (s *GetSchemeTaskConfigResponseBodyData) GetModelName() *string {
	return s.ModelName
}

func (s *GetSchemeTaskConfigResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetSchemeTaskConfigResponseBodyData) GetSchemeIdList() []*int64 {
	return s.SchemeIdList
}

func (s *GetSchemeTaskConfigResponseBodyData) GetSchemeList() []*GetSchemeTaskConfigResponseBodyDataSchemeList {
	return s.SchemeList
}

func (s *GetSchemeTaskConfigResponseBodyData) GetSchemeTaskConfigId() *int64 {
	return s.SchemeTaskConfigId
}

func (s *GetSchemeTaskConfigResponseBodyData) GetSourceDataType() *string {
	return s.SourceDataType
}

func (s *GetSchemeTaskConfigResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetSchemeTaskConfigResponseBodyData) SetAsrTaskPriority(v int32) *GetSchemeTaskConfigResponseBodyData {
	s.AsrTaskPriority = &v
	return s
}

func (s *GetSchemeTaskConfigResponseBodyData) SetAssignType(v int32) *GetSchemeTaskConfigResponseBodyData {
	s.AssignType = &v
	return s
}

func (s *GetSchemeTaskConfigResponseBodyData) SetDataConfig(v *GetSchemeTaskConfigResponseBodyDataDataConfig) *GetSchemeTaskConfigResponseBodyData {
	s.DataConfig = v
	return s
}

func (s *GetSchemeTaskConfigResponseBodyData) SetId(v int64) *GetSchemeTaskConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetSchemeTaskConfigResponseBodyData) SetManualReview(v int32) *GetSchemeTaskConfigResponseBodyData {
	s.ManualReview = &v
	return s
}

func (s *GetSchemeTaskConfigResponseBodyData) SetModeCustomizationId(v string) *GetSchemeTaskConfigResponseBodyData {
	s.ModeCustomizationId = &v
	return s
}

func (s *GetSchemeTaskConfigResponseBodyData) SetModelName(v string) *GetSchemeTaskConfigResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *GetSchemeTaskConfigResponseBodyData) SetName(v string) *GetSchemeTaskConfigResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetSchemeTaskConfigResponseBodyData) SetSchemeIdList(v []*int64) *GetSchemeTaskConfigResponseBodyData {
	s.SchemeIdList = v
	return s
}

func (s *GetSchemeTaskConfigResponseBodyData) SetSchemeList(v []*GetSchemeTaskConfigResponseBodyDataSchemeList) *GetSchemeTaskConfigResponseBodyData {
	s.SchemeList = v
	return s
}

func (s *GetSchemeTaskConfigResponseBodyData) SetSchemeTaskConfigId(v int64) *GetSchemeTaskConfigResponseBodyData {
	s.SchemeTaskConfigId = &v
	return s
}

func (s *GetSchemeTaskConfigResponseBodyData) SetSourceDataType(v string) *GetSchemeTaskConfigResponseBodyData {
	s.SourceDataType = &v
	return s
}

func (s *GetSchemeTaskConfigResponseBodyData) SetStatus(v string) *GetSchemeTaskConfigResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetSchemeTaskConfigResponseBodyData) Validate() error {
	if s.DataConfig != nil {
		if err := s.DataConfig.Validate(); err != nil {
			return err
		}
	}
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

type GetSchemeTaskConfigResponseBodyDataDataConfig struct {
	// Data screening items for on-the-fly recording
	AssignConfigs []*GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigs `json:"AssignConfigs,omitempty" xml:"AssignConfigs,omitempty" type:"Repeated"`
	// Dataset task. Manage datasets.
	//
	// example:
	//
	// []
	DataSets *string `json:"DataSets,omitempty" xml:"DataSets,omitempty"`
	// Index number
	//
	// example:
	//
	// 0
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// JSON text for filtering conditions used in secondary quality inspection. For details, see the request parameters of the GetResult API.
	//
	// example:
	//
	// {}
	ResultParam *string `json:"ResultParam,omitempty" xml:"ResultParam,omitempty"`
}

func (s GetSchemeTaskConfigResponseBodyDataDataConfig) String() string {
	return dara.Prettify(s)
}

func (s GetSchemeTaskConfigResponseBodyDataDataConfig) GoString() string {
	return s.String()
}

func (s *GetSchemeTaskConfigResponseBodyDataDataConfig) GetAssignConfigs() []*GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigs {
	return s.AssignConfigs
}

func (s *GetSchemeTaskConfigResponseBodyDataDataConfig) GetDataSets() *string {
	return s.DataSets
}

func (s *GetSchemeTaskConfigResponseBodyDataDataConfig) GetIndex() *int64 {
	return s.Index
}

func (s *GetSchemeTaskConfigResponseBodyDataDataConfig) GetResultParam() *string {
	return s.ResultParam
}

func (s *GetSchemeTaskConfigResponseBodyDataDataConfig) SetAssignConfigs(v []*GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigs) *GetSchemeTaskConfigResponseBodyDataDataConfig {
	s.AssignConfigs = v
	return s
}

func (s *GetSchemeTaskConfigResponseBodyDataDataConfig) SetDataSets(v string) *GetSchemeTaskConfigResponseBodyDataDataConfig {
	s.DataSets = &v
	return s
}

func (s *GetSchemeTaskConfigResponseBodyDataDataConfig) SetIndex(v int64) *GetSchemeTaskConfigResponseBodyDataDataConfig {
	s.Index = &v
	return s
}

func (s *GetSchemeTaskConfigResponseBodyDataDataConfig) SetResultParam(v string) *GetSchemeTaskConfigResponseBodyDataDataConfig {
	s.ResultParam = &v
	return s
}

func (s *GetSchemeTaskConfigResponseBodyDataDataConfig) Validate() error {
	if s.AssignConfigs != nil {
		for _, item := range s.AssignConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigs struct {
	// Parameter matching configurations for on-the-fly recording
	AssignConfigContests []*GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigsAssignConfigContests `json:"AssignConfigContests,omitempty" xml:"AssignConfigContests,omitempty" type:"Repeated"`
}

func (s GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigs) GoString() string {
	return s.String()
}

func (s *GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigs) GetAssignConfigContests() []*GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigsAssignConfigContests {
	return s.AssignConfigContests
}

func (s *GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigs) SetAssignConfigContests(v []*GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigsAssignConfigContests) *GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigs {
	s.AssignConfigContests = v
	return s
}

func (s *GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigs) Validate() error {
	if s.AssignConfigContests != nil {
		for _, item := range s.AssignConfigContests {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigsAssignConfigContests struct {
	// Type of the value
	//
	// - 0: String
	//
	// - 1: Number
	//
	// - 2: List (use list type for all parameter values when using =)
	//
	// - 3: Date
	//
	// - 4: List_Json
	//
	// example:
	//
	// 3
	DataType *int32 `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// List of on-the-fly recording data
	ListObject []interface{} `json:"ListObject,omitempty" xml:"ListObject,omitempty" type:"Repeated"`
	// Check item name
	//
	// example:
	//
	// callStartTime
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Operator
	//
	// - 1: ==
	//
	// - 2: >
	//
	// - 3: <
	//
	// - 4: range
	//
	// - 5: >=
	//
	// - 6: <=
	//
	// - 7: !=
	//
	// - 8: null
	//
	// - 9: not null
	//
	// - 10: contains
	//
	// - 11: does not contain
	//
	// example:
	//
	// 4
	Symbol *int32 `json:"Symbol,omitempty" xml:"Symbol,omitempty"`
	// Matching value for on-the-fly recording data
	//
	// example:
	//
	// {\\"start\\":\\"2022-09-01 00:00:00\\",\\"end\\":\\"2022-09-30 00:00:00\\"}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigsAssignConfigContests) String() string {
	return dara.Prettify(s)
}

func (s GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigsAssignConfigContests) GoString() string {
	return s.String()
}

func (s *GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigsAssignConfigContests) GetDataType() *int32 {
	return s.DataType
}

func (s *GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigsAssignConfigContests) GetListObject() []interface{} {
	return s.ListObject
}

func (s *GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigsAssignConfigContests) GetName() *string {
	return s.Name
}

func (s *GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigsAssignConfigContests) GetSymbol() *int32 {
	return s.Symbol
}

func (s *GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigsAssignConfigContests) GetValue() *string {
	return s.Value
}

func (s *GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigsAssignConfigContests) SetDataType(v int32) *GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigsAssignConfigContests {
	s.DataType = &v
	return s
}

func (s *GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigsAssignConfigContests) SetListObject(v []interface{}) *GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigsAssignConfigContests {
	s.ListObject = v
	return s
}

func (s *GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigsAssignConfigContests) SetName(v string) *GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigsAssignConfigContests {
	s.Name = &v
	return s
}

func (s *GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigsAssignConfigContests) SetSymbol(v int32) *GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigsAssignConfigContests {
	s.Symbol = &v
	return s
}

func (s *GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigsAssignConfigContests) SetValue(v string) *GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigsAssignConfigContests {
	s.Value = &v
	return s
}

func (s *GetSchemeTaskConfigResponseBodyDataDataConfigAssignConfigsAssignConfigContests) Validate() error {
	return dara.Validate(s)
}

type GetSchemeTaskConfigResponseBodyDataSchemeList struct {
	// Quality inspection scheme name
	//
	// example:
	//
	// 质检方案B
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Quality inspection scheme ID
	//
	// example:
	//
	// 158
	SchemeId *int64 `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
}

func (s GetSchemeTaskConfigResponseBodyDataSchemeList) String() string {
	return dara.Prettify(s)
}

func (s GetSchemeTaskConfigResponseBodyDataSchemeList) GoString() string {
	return s.String()
}

func (s *GetSchemeTaskConfigResponseBodyDataSchemeList) GetName() *string {
	return s.Name
}

func (s *GetSchemeTaskConfigResponseBodyDataSchemeList) GetSchemeId() *int64 {
	return s.SchemeId
}

func (s *GetSchemeTaskConfigResponseBodyDataSchemeList) SetName(v string) *GetSchemeTaskConfigResponseBodyDataSchemeList {
	s.Name = &v
	return s
}

func (s *GetSchemeTaskConfigResponseBodyDataSchemeList) SetSchemeId(v int64) *GetSchemeTaskConfigResponseBodyDataSchemeList {
	s.SchemeId = &v
	return s
}

func (s *GetSchemeTaskConfigResponseBodyDataSchemeList) Validate() error {
	return dara.Validate(s)
}
