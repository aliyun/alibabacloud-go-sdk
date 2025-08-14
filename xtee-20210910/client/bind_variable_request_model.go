// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *BindVariableRequest
	GetLang() *string
	SetApiRegionId(v string) *BindVariableRequest
	GetApiRegionId() *string
	SetApiType(v string) *BindVariableRequest
	GetApiType() *string
	SetCreateType(v string) *BindVariableRequest
	GetCreateType() *string
	SetDefineId(v string) *BindVariableRequest
	GetDefineId() *string
	SetDefineIds(v string) *BindVariableRequest
	GetDefineIds() *string
	SetDescription(v string) *BindVariableRequest
	GetDescription() *string
	SetEventCode(v string) *BindVariableRequest
	GetEventCode() *string
	SetExceptionValue(v string) *BindVariableRequest
	GetExceptionValue() *string
	SetId(v int64) *BindVariableRequest
	GetId() *int64
	SetOutputField(v string) *BindVariableRequest
	GetOutputField() *string
	SetOutputType(v string) *BindVariableRequest
	GetOutputType() *string
	SetParams(v string) *BindVariableRequest
	GetParams() *string
	SetParamsList(v string) *BindVariableRequest
	GetParamsList() *string
	SetRegId(v string) *BindVariableRequest
	GetRegId() *string
	SetSourceType(v string) *BindVariableRequest
	GetSourceType() *string
	SetTitle(v string) *BindVariableRequest
	GetTitle() *string
}

type BindVariableRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// API region ID
	//
	// example:
	//
	// cn-hangzhou
	ApiRegionId *string `json:"apiRegionId,omitempty" xml:"apiRegionId,omitempty"`
	// API type
	//
	// example:
	//
	// SELF
	ApiType *string `json:"apiType,omitempty" xml:"apiType,omitempty"`
	// Creation type
	//
	// example:
	//
	// NORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// Associated variable definition primary key ID
	//
	// example:
	//
	// 2438
	DefineId *string `json:"defineId,omitempty" xml:"defineId,omitempty"`
	// Variable definition IDs, can be multiple. If binding multiple IDs, separate them with commas
	//
	// example:
	//
	// 1546, 1547
	DefineIds *string `json:"defineIds,omitempty" xml:"defineIds,omitempty"`
	// Description information.
	//
	// example:
	//
	// 变量描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Event code
	//
	// This parameter is required.
	//
	// example:
	//
	// de_agbzfi5134
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Exception value
	//
	// example:
	//
	// SYS_ERROR
	ExceptionValue *string `json:"exceptionValue,omitempty" xml:"exceptionValue,omitempty"`
	// Variable primary key ID
	//
	// example:
	//
	// 106875
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Output field path
	//
	// example:
	//
	// BOOLEAN
	OutputField *string `json:"outputField,omitempty" xml:"outputField,omitempty"`
	// Output type
	//
	// example:
	//
	// STRING
	OutputType *string `json:"outputType,omitempty" xml:"outputType,omitempty"`
	// Binding input parameter information
	//
	// example:
	//
	// {"accountId":"accountId","mobile":""}
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
	// Event parameter mapping 2.0, either params or paramsList must not be empty. List, JSON structure
	//
	// example:
	//
	// [{"eventFieldName":"accountId","required":false}]
	ParamsList *string `json:"paramsList,omitempty" xml:"paramsList,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Variable source
	//
	// example:
	//
	// SAF
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// Title
	//
	// This parameter is required.
	//
	// example:
	//
	// 变量title
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s BindVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s BindVariableRequest) GoString() string {
	return s.String()
}

func (s *BindVariableRequest) GetLang() *string {
	return s.Lang
}

func (s *BindVariableRequest) GetApiRegionId() *string {
	return s.ApiRegionId
}

func (s *BindVariableRequest) GetApiType() *string {
	return s.ApiType
}

func (s *BindVariableRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *BindVariableRequest) GetDefineId() *string {
	return s.DefineId
}

func (s *BindVariableRequest) GetDefineIds() *string {
	return s.DefineIds
}

func (s *BindVariableRequest) GetDescription() *string {
	return s.Description
}

func (s *BindVariableRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *BindVariableRequest) GetExceptionValue() *string {
	return s.ExceptionValue
}

func (s *BindVariableRequest) GetId() *int64 {
	return s.Id
}

func (s *BindVariableRequest) GetOutputField() *string {
	return s.OutputField
}

func (s *BindVariableRequest) GetOutputType() *string {
	return s.OutputType
}

func (s *BindVariableRequest) GetParams() *string {
	return s.Params
}

func (s *BindVariableRequest) GetParamsList() *string {
	return s.ParamsList
}

func (s *BindVariableRequest) GetRegId() *string {
	return s.RegId
}

func (s *BindVariableRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *BindVariableRequest) GetTitle() *string {
	return s.Title
}

func (s *BindVariableRequest) SetLang(v string) *BindVariableRequest {
	s.Lang = &v
	return s
}

func (s *BindVariableRequest) SetApiRegionId(v string) *BindVariableRequest {
	s.ApiRegionId = &v
	return s
}

func (s *BindVariableRequest) SetApiType(v string) *BindVariableRequest {
	s.ApiType = &v
	return s
}

func (s *BindVariableRequest) SetCreateType(v string) *BindVariableRequest {
	s.CreateType = &v
	return s
}

func (s *BindVariableRequest) SetDefineId(v string) *BindVariableRequest {
	s.DefineId = &v
	return s
}

func (s *BindVariableRequest) SetDefineIds(v string) *BindVariableRequest {
	s.DefineIds = &v
	return s
}

func (s *BindVariableRequest) SetDescription(v string) *BindVariableRequest {
	s.Description = &v
	return s
}

func (s *BindVariableRequest) SetEventCode(v string) *BindVariableRequest {
	s.EventCode = &v
	return s
}

func (s *BindVariableRequest) SetExceptionValue(v string) *BindVariableRequest {
	s.ExceptionValue = &v
	return s
}

func (s *BindVariableRequest) SetId(v int64) *BindVariableRequest {
	s.Id = &v
	return s
}

func (s *BindVariableRequest) SetOutputField(v string) *BindVariableRequest {
	s.OutputField = &v
	return s
}

func (s *BindVariableRequest) SetOutputType(v string) *BindVariableRequest {
	s.OutputType = &v
	return s
}

func (s *BindVariableRequest) SetParams(v string) *BindVariableRequest {
	s.Params = &v
	return s
}

func (s *BindVariableRequest) SetParamsList(v string) *BindVariableRequest {
	s.ParamsList = &v
	return s
}

func (s *BindVariableRequest) SetRegId(v string) *BindVariableRequest {
	s.RegId = &v
	return s
}

func (s *BindVariableRequest) SetSourceType(v string) *BindVariableRequest {
	s.SourceType = &v
	return s
}

func (s *BindVariableRequest) SetTitle(v string) *BindVariableRequest {
	s.Title = &v
	return s
}

func (s *BindVariableRequest) Validate() error {
	return dara.Validate(s)
}
