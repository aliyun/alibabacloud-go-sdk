// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventVariableListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeEventVariableListResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeEventVariableListResponseBodyResultObject) *DescribeEventVariableListResponseBody
	GetResultObject() *DescribeEventVariableListResponseBodyResultObject
}

type DescribeEventVariableListResponseBody struct {
	RequestId    *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *DescribeEventVariableListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeEventVariableListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventVariableListResponseBody) GetResultObject() *DescribeEventVariableListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeEventVariableListResponseBody) SetRequestId(v string) *DescribeEventVariableListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventVariableListResponseBody) SetResultObject(v *DescribeEventVariableListResponseBodyResultObject) *DescribeEventVariableListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeEventVariableListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObject struct {
	Actions                 []*DescribeEventVariableListResponseBodyResultObjectActions                 `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	DeviceVariables         []*DescribeEventVariableListResponseBodyResultObjectDeviceVariables         `json:"deviceVariables,omitempty" xml:"deviceVariables,omitempty" type:"Repeated"`
	ExpressionVariables     []*DescribeEventVariableListResponseBodyResultObjectExpressionVariables     `json:"expressionVariables,omitempty" xml:"expressionVariables,omitempty" type:"Repeated"`
	FavoriteVariables       []*DescribeEventVariableListResponseBodyResultObjectFavoriteVariables       `json:"favoriteVariables,omitempty" xml:"favoriteVariables,omitempty" type:"Repeated"`
	MiddleVariables         []*DescribeEventVariableListResponseBodyResultObjectMiddleVariables         `json:"middleVariables,omitempty" xml:"middleVariables,omitempty" type:"Repeated"`
	ModelVariables          []*DescribeEventVariableListResponseBodyResultObjectModelVariables          `json:"modelVariables,omitempty" xml:"modelVariables,omitempty" type:"Repeated"`
	NameList                []*DescribeEventVariableListResponseBodyResultObjectNameList                `json:"nameList,omitempty" xml:"nameList,omitempty" type:"Repeated"`
	NativeVariableFunctions []*DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions `json:"nativeVariableFunctions,omitempty" xml:"nativeVariableFunctions,omitempty" type:"Repeated"`
	NativeVariables         []*DescribeEventVariableListResponseBodyResultObjectNativeVariables         `json:"nativeVariables,omitempty" xml:"nativeVariables,omitempty" type:"Repeated"`
	QueryVariables          []*DescribeEventVariableListResponseBodyResultObjectQueryVariables          `json:"queryVariables,omitempty" xml:"queryVariables,omitempty" type:"Repeated"`
	SelfVariables           []*DescribeEventVariableListResponseBodyResultObjectSelfVariables           `json:"selfVariables,omitempty" xml:"selfVariables,omitempty" type:"Repeated"`
	SysVariables            []*DescribeEventVariableListResponseBodyResultObjectSysVariables            `json:"sysVariables,omitempty" xml:"sysVariables,omitempty" type:"Repeated"`
	ThirdVariables          map[string]interface{}                                                      `json:"thirdVariables,omitempty" xml:"thirdVariables,omitempty"`
	VelocityVariables       []*DescribeEventVariableListResponseBodyResultObjectVelocityVariables       `json:"velocityVariables,omitempty" xml:"velocityVariables,omitempty" type:"Repeated"`
}

func (s DescribeEventVariableListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObject) GetActions() []*DescribeEventVariableListResponseBodyResultObjectActions {
	return s.Actions
}

func (s *DescribeEventVariableListResponseBodyResultObject) GetDeviceVariables() []*DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	return s.DeviceVariables
}

func (s *DescribeEventVariableListResponseBodyResultObject) GetExpressionVariables() []*DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	return s.ExpressionVariables
}

func (s *DescribeEventVariableListResponseBodyResultObject) GetFavoriteVariables() []*DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	return s.FavoriteVariables
}

func (s *DescribeEventVariableListResponseBodyResultObject) GetMiddleVariables() []*DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	return s.MiddleVariables
}

func (s *DescribeEventVariableListResponseBodyResultObject) GetModelVariables() []*DescribeEventVariableListResponseBodyResultObjectModelVariables {
	return s.ModelVariables
}

func (s *DescribeEventVariableListResponseBodyResultObject) GetNameList() []*DescribeEventVariableListResponseBodyResultObjectNameList {
	return s.NameList
}

func (s *DescribeEventVariableListResponseBodyResultObject) GetNativeVariableFunctions() []*DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	return s.NativeVariableFunctions
}

func (s *DescribeEventVariableListResponseBodyResultObject) GetNativeVariables() []*DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	return s.NativeVariables
}

func (s *DescribeEventVariableListResponseBodyResultObject) GetQueryVariables() []*DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	return s.QueryVariables
}

func (s *DescribeEventVariableListResponseBodyResultObject) GetSelfVariables() []*DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	return s.SelfVariables
}

func (s *DescribeEventVariableListResponseBodyResultObject) GetSysVariables() []*DescribeEventVariableListResponseBodyResultObjectSysVariables {
	return s.SysVariables
}

func (s *DescribeEventVariableListResponseBodyResultObject) GetThirdVariables() map[string]interface{} {
	return s.ThirdVariables
}

func (s *DescribeEventVariableListResponseBodyResultObject) GetVelocityVariables() []*DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	return s.VelocityVariables
}

func (s *DescribeEventVariableListResponseBodyResultObject) SetActions(v []*DescribeEventVariableListResponseBodyResultObjectActions) *DescribeEventVariableListResponseBodyResultObject {
	s.Actions = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObject) SetDeviceVariables(v []*DescribeEventVariableListResponseBodyResultObjectDeviceVariables) *DescribeEventVariableListResponseBodyResultObject {
	s.DeviceVariables = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObject) SetExpressionVariables(v []*DescribeEventVariableListResponseBodyResultObjectExpressionVariables) *DescribeEventVariableListResponseBodyResultObject {
	s.ExpressionVariables = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObject) SetFavoriteVariables(v []*DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) *DescribeEventVariableListResponseBodyResultObject {
	s.FavoriteVariables = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObject) SetMiddleVariables(v []*DescribeEventVariableListResponseBodyResultObjectMiddleVariables) *DescribeEventVariableListResponseBodyResultObject {
	s.MiddleVariables = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObject) SetModelVariables(v []*DescribeEventVariableListResponseBodyResultObjectModelVariables) *DescribeEventVariableListResponseBodyResultObject {
	s.ModelVariables = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObject) SetNameList(v []*DescribeEventVariableListResponseBodyResultObjectNameList) *DescribeEventVariableListResponseBodyResultObject {
	s.NameList = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObject) SetNativeVariableFunctions(v []*DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) *DescribeEventVariableListResponseBodyResultObject {
	s.NativeVariableFunctions = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObject) SetNativeVariables(v []*DescribeEventVariableListResponseBodyResultObjectNativeVariables) *DescribeEventVariableListResponseBodyResultObject {
	s.NativeVariables = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObject) SetQueryVariables(v []*DescribeEventVariableListResponseBodyResultObjectQueryVariables) *DescribeEventVariableListResponseBodyResultObject {
	s.QueryVariables = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObject) SetSelfVariables(v []*DescribeEventVariableListResponseBodyResultObjectSelfVariables) *DescribeEventVariableListResponseBodyResultObject {
	s.SelfVariables = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObject) SetSysVariables(v []*DescribeEventVariableListResponseBodyResultObjectSysVariables) *DescribeEventVariableListResponseBodyResultObject {
	s.SysVariables = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObject) SetThirdVariables(v map[string]interface{}) *DescribeEventVariableListResponseBodyResultObject {
	s.ThirdVariables = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObject) SetVelocityVariables(v []*DescribeEventVariableListResponseBodyResultObjectVelocityVariables) *DescribeEventVariableListResponseBodyResultObject {
	s.VelocityVariables = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectActions struct {
	Code             *string                                                                   `json:"code,omitempty" xml:"code,omitempty"`
	DataDisplay      *string                                                                   `json:"dataDisplay,omitempty" xml:"dataDisplay,omitempty"`
	DefineId         *string                                                                   `json:"defineId,omitempty" xml:"defineId,omitempty"`
	Description      *string                                                                   `json:"description,omitempty" xml:"description,omitempty"`
	DisplayType      *string                                                                   `json:"displayType,omitempty" xml:"displayType,omitempty"`
	ExpressionTitle  *string                                                                   `json:"expressionTitle,omitempty" xml:"expressionTitle,omitempty"`
	FavoriteFlag     *bool                                                                     `json:"favoriteFlag,omitempty" xml:"favoriteFlag,omitempty"`
	FieldDetail      *string                                                                   `json:"fieldDetail,omitempty" xml:"fieldDetail,omitempty"`
	FieldRank        *int32                                                                    `json:"fieldRank,omitempty" xml:"fieldRank,omitempty"`
	FieldSource      *string                                                                   `json:"fieldSource,omitempty" xml:"fieldSource,omitempty"`
	FieldType        *string                                                                   `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	Id               *int64                                                                    `json:"id,omitempty" xml:"id,omitempty"`
	InputFieldType   *string                                                                   `json:"inputFieldType,omitempty" xml:"inputFieldType,omitempty"`
	InputRequired    *string                                                                   `json:"inputRequired,omitempty" xml:"inputRequired,omitempty"`
	Inputs           *string                                                                   `json:"inputs,omitempty" xml:"inputs,omitempty"`
	Name             *string                                                                   `json:"name,omitempty" xml:"name,omitempty"`
	Outlier          *string                                                                   `json:"outlier,omitempty" xml:"outlier,omitempty"`
	OutputThreshold  *DescribeEventVariableListResponseBodyResultObjectActionsOutputThreshold  `json:"outputThreshold,omitempty" xml:"outputThreshold,omitempty" type:"Struct"`
	ParentName       *string                                                                   `json:"parentName,omitempty" xml:"parentName,omitempty"`
	SourceType       *string                                                                   `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Title            *string                                                                   `json:"title,omitempty" xml:"title,omitempty"`
	Type             *string                                                                   `json:"type,omitempty" xml:"type,omitempty"`
	VariableVelocity *DescribeEventVariableListResponseBodyResultObjectActionsVariableVelocity `json:"variableVelocity,omitempty" xml:"variableVelocity,omitempty" type:"Struct"`
	XLabel           *string                                                                   `json:"xLabel,omitempty" xml:"xLabel,omitempty"`
	YLabel           *string                                                                   `json:"yLabel,omitempty" xml:"yLabel,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectActions) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectActions) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) GetCode() *string {
	return s.Code
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) GetDataDisplay() *string {
	return s.DataDisplay
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) GetDefineId() *string {
	return s.DefineId
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) GetDescription() *string {
	return s.Description
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) GetDisplayType() *string {
	return s.DisplayType
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) GetExpressionTitle() *string {
	return s.ExpressionTitle
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) GetFavoriteFlag() *bool {
	return s.FavoriteFlag
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) GetFieldDetail() *string {
	return s.FieldDetail
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) GetFieldRank() *int32 {
	return s.FieldRank
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) GetFieldSource() *string {
	return s.FieldSource
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) GetInputFieldType() *string {
	return s.InputFieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) GetInputRequired() *string {
	return s.InputRequired
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) GetInputs() *string {
	return s.Inputs
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) GetName() *string {
	return s.Name
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) GetOutlier() *string {
	return s.Outlier
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) GetOutputThreshold() *DescribeEventVariableListResponseBodyResultObjectActionsOutputThreshold {
	return s.OutputThreshold
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) GetParentName() *string {
	return s.ParentName
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) GetTitle() *string {
	return s.Title
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) GetType() *string {
	return s.Type
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) GetVariableVelocity() *DescribeEventVariableListResponseBodyResultObjectActionsVariableVelocity {
	return s.VariableVelocity
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) GetXLabel() *string {
	return s.XLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) GetYLabel() *string {
	return s.YLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) SetCode(v string) *DescribeEventVariableListResponseBodyResultObjectActions {
	s.Code = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) SetDataDisplay(v string) *DescribeEventVariableListResponseBodyResultObjectActions {
	s.DataDisplay = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) SetDefineId(v string) *DescribeEventVariableListResponseBodyResultObjectActions {
	s.DefineId = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) SetDescription(v string) *DescribeEventVariableListResponseBodyResultObjectActions {
	s.Description = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) SetDisplayType(v string) *DescribeEventVariableListResponseBodyResultObjectActions {
	s.DisplayType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) SetExpressionTitle(v string) *DescribeEventVariableListResponseBodyResultObjectActions {
	s.ExpressionTitle = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) SetFavoriteFlag(v bool) *DescribeEventVariableListResponseBodyResultObjectActions {
	s.FavoriteFlag = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) SetFieldDetail(v string) *DescribeEventVariableListResponseBodyResultObjectActions {
	s.FieldDetail = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) SetFieldRank(v int32) *DescribeEventVariableListResponseBodyResultObjectActions {
	s.FieldRank = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) SetFieldSource(v string) *DescribeEventVariableListResponseBodyResultObjectActions {
	s.FieldSource = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) SetFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectActions {
	s.FieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) SetId(v int64) *DescribeEventVariableListResponseBodyResultObjectActions {
	s.Id = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) SetInputFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectActions {
	s.InputFieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) SetInputRequired(v string) *DescribeEventVariableListResponseBodyResultObjectActions {
	s.InputRequired = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) SetInputs(v string) *DescribeEventVariableListResponseBodyResultObjectActions {
	s.Inputs = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) SetName(v string) *DescribeEventVariableListResponseBodyResultObjectActions {
	s.Name = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) SetOutlier(v string) *DescribeEventVariableListResponseBodyResultObjectActions {
	s.Outlier = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) SetOutputThreshold(v *DescribeEventVariableListResponseBodyResultObjectActionsOutputThreshold) *DescribeEventVariableListResponseBodyResultObjectActions {
	s.OutputThreshold = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) SetParentName(v string) *DescribeEventVariableListResponseBodyResultObjectActions {
	s.ParentName = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) SetSourceType(v string) *DescribeEventVariableListResponseBodyResultObjectActions {
	s.SourceType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) SetTitle(v string) *DescribeEventVariableListResponseBodyResultObjectActions {
	s.Title = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) SetType(v string) *DescribeEventVariableListResponseBodyResultObjectActions {
	s.Type = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) SetVariableVelocity(v *DescribeEventVariableListResponseBodyResultObjectActionsVariableVelocity) *DescribeEventVariableListResponseBodyResultObjectActions {
	s.VariableVelocity = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) SetXLabel(v string) *DescribeEventVariableListResponseBodyResultObjectActions {
	s.XLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) SetYLabel(v string) *DescribeEventVariableListResponseBodyResultObjectActions {
	s.YLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActions) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectActionsOutputThreshold struct {
	MaxValue *float64 `json:"maxValue,omitempty" xml:"maxValue,omitempty"`
	MinValue *float64 `json:"minValue,omitempty" xml:"minValue,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectActionsOutputThreshold) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectActionsOutputThreshold) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectActionsOutputThreshold) GetMaxValue() *float64 {
	return s.MaxValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectActionsOutputThreshold) GetMinValue() *float64 {
	return s.MinValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectActionsOutputThreshold) SetMaxValue(v float64) *DescribeEventVariableListResponseBodyResultObjectActionsOutputThreshold {
	s.MaxValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActionsOutputThreshold) SetMinValue(v float64) *DescribeEventVariableListResponseBodyResultObjectActionsOutputThreshold {
	s.MinValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActionsOutputThreshold) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectActionsVariableVelocity struct {
	Iv *string `json:"iv,omitempty" xml:"iv,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectActionsVariableVelocity) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectActionsVariableVelocity) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectActionsVariableVelocity) GetIv() *string {
	return s.Iv
}

func (s *DescribeEventVariableListResponseBodyResultObjectActionsVariableVelocity) SetIv(v string) *DescribeEventVariableListResponseBodyResultObjectActionsVariableVelocity {
	s.Iv = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectActionsVariableVelocity) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectDeviceVariables struct {
	Code             *string                                                                           `json:"code,omitempty" xml:"code,omitempty"`
	DataDisplay      *string                                                                           `json:"dataDisplay,omitempty" xml:"dataDisplay,omitempty"`
	DefineId         *string                                                                           `json:"defineId,omitempty" xml:"defineId,omitempty"`
	Description      *string                                                                           `json:"description,omitempty" xml:"description,omitempty"`
	DisplayType      *string                                                                           `json:"displayType,omitempty" xml:"displayType,omitempty"`
	ExpressionTitle  *string                                                                           `json:"expressionTitle,omitempty" xml:"expressionTitle,omitempty"`
	FavoriteFlag     *bool                                                                             `json:"favoriteFlag,omitempty" xml:"favoriteFlag,omitempty"`
	FieldDetail      *string                                                                           `json:"fieldDetail,omitempty" xml:"fieldDetail,omitempty"`
	FieldRank        *int32                                                                            `json:"fieldRank,omitempty" xml:"fieldRank,omitempty"`
	FieldSource      *string                                                                           `json:"fieldSource,omitempty" xml:"fieldSource,omitempty"`
	FieldType        *string                                                                           `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	Id               *int64                                                                            `json:"id,omitempty" xml:"id,omitempty"`
	InputFieldType   *string                                                                           `json:"inputFieldType,omitempty" xml:"inputFieldType,omitempty"`
	InputRequired    *string                                                                           `json:"inputRequired,omitempty" xml:"inputRequired,omitempty"`
	Inputs           *string                                                                           `json:"inputs,omitempty" xml:"inputs,omitempty"`
	Name             *string                                                                           `json:"name,omitempty" xml:"name,omitempty"`
	Outlier          *string                                                                           `json:"outlier,omitempty" xml:"outlier,omitempty"`
	OutputThreshold  *DescribeEventVariableListResponseBodyResultObjectDeviceVariablesOutputThreshold  `json:"outputThreshold,omitempty" xml:"outputThreshold,omitempty" type:"Struct"`
	ParentName       *string                                                                           `json:"parentName,omitempty" xml:"parentName,omitempty"`
	SourceType       *string                                                                           `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Title            *string                                                                           `json:"title,omitempty" xml:"title,omitempty"`
	Type             *string                                                                           `json:"type,omitempty" xml:"type,omitempty"`
	VariableVelocity *DescribeEventVariableListResponseBodyResultObjectDeviceVariablesVariableVelocity `json:"variableVelocity,omitempty" xml:"variableVelocity,omitempty" type:"Struct"`
	XLabel           *string                                                                           `json:"xLabel,omitempty" xml:"xLabel,omitempty"`
	YLabel           *string                                                                           `json:"yLabel,omitempty" xml:"yLabel,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectDeviceVariables) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GetCode() *string {
	return s.Code
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GetDataDisplay() *string {
	return s.DataDisplay
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GetDefineId() *string {
	return s.DefineId
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GetDescription() *string {
	return s.Description
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GetDisplayType() *string {
	return s.DisplayType
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GetExpressionTitle() *string {
	return s.ExpressionTitle
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GetFavoriteFlag() *bool {
	return s.FavoriteFlag
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GetFieldDetail() *string {
	return s.FieldDetail
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GetFieldRank() *int32 {
	return s.FieldRank
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GetFieldSource() *string {
	return s.FieldSource
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GetInputFieldType() *string {
	return s.InputFieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GetInputRequired() *string {
	return s.InputRequired
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GetInputs() *string {
	return s.Inputs
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GetName() *string {
	return s.Name
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GetOutlier() *string {
	return s.Outlier
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GetOutputThreshold() *DescribeEventVariableListResponseBodyResultObjectDeviceVariablesOutputThreshold {
	return s.OutputThreshold
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GetParentName() *string {
	return s.ParentName
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GetTitle() *string {
	return s.Title
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GetType() *string {
	return s.Type
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GetVariableVelocity() *DescribeEventVariableListResponseBodyResultObjectDeviceVariablesVariableVelocity {
	return s.VariableVelocity
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GetXLabel() *string {
	return s.XLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) GetYLabel() *string {
	return s.YLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) SetCode(v string) *DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	s.Code = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) SetDataDisplay(v string) *DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	s.DataDisplay = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) SetDefineId(v string) *DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	s.DefineId = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) SetDescription(v string) *DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	s.Description = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) SetDisplayType(v string) *DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	s.DisplayType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) SetExpressionTitle(v string) *DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	s.ExpressionTitle = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) SetFavoriteFlag(v bool) *DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	s.FavoriteFlag = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) SetFieldDetail(v string) *DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	s.FieldDetail = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) SetFieldRank(v int32) *DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	s.FieldRank = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) SetFieldSource(v string) *DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	s.FieldSource = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) SetFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	s.FieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) SetId(v int64) *DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	s.Id = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) SetInputFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	s.InputFieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) SetInputRequired(v string) *DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	s.InputRequired = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) SetInputs(v string) *DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	s.Inputs = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) SetName(v string) *DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	s.Name = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) SetOutlier(v string) *DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	s.Outlier = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) SetOutputThreshold(v *DescribeEventVariableListResponseBodyResultObjectDeviceVariablesOutputThreshold) *DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	s.OutputThreshold = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) SetParentName(v string) *DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	s.ParentName = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) SetSourceType(v string) *DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	s.SourceType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) SetTitle(v string) *DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	s.Title = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) SetType(v string) *DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	s.Type = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) SetVariableVelocity(v *DescribeEventVariableListResponseBodyResultObjectDeviceVariablesVariableVelocity) *DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	s.VariableVelocity = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) SetXLabel(v string) *DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	s.XLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) SetYLabel(v string) *DescribeEventVariableListResponseBodyResultObjectDeviceVariables {
	s.YLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariables) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectDeviceVariablesOutputThreshold struct {
	MaxValue *float64 `json:"maxValue,omitempty" xml:"maxValue,omitempty"`
	MinValue *float64 `json:"minValue,omitempty" xml:"minValue,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectDeviceVariablesOutputThreshold) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectDeviceVariablesOutputThreshold) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariablesOutputThreshold) GetMaxValue() *float64 {
	return s.MaxValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariablesOutputThreshold) GetMinValue() *float64 {
	return s.MinValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariablesOutputThreshold) SetMaxValue(v float64) *DescribeEventVariableListResponseBodyResultObjectDeviceVariablesOutputThreshold {
	s.MaxValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariablesOutputThreshold) SetMinValue(v float64) *DescribeEventVariableListResponseBodyResultObjectDeviceVariablesOutputThreshold {
	s.MinValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariablesOutputThreshold) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectDeviceVariablesVariableVelocity struct {
	Iv *string `json:"iv,omitempty" xml:"iv,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectDeviceVariablesVariableVelocity) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectDeviceVariablesVariableVelocity) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariablesVariableVelocity) GetIv() *string {
	return s.Iv
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariablesVariableVelocity) SetIv(v string) *DescribeEventVariableListResponseBodyResultObjectDeviceVariablesVariableVelocity {
	s.Iv = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectDeviceVariablesVariableVelocity) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectExpressionVariables struct {
	Code             *string                                                                               `json:"code,omitempty" xml:"code,omitempty"`
	DataDisplay      *string                                                                               `json:"dataDisplay,omitempty" xml:"dataDisplay,omitempty"`
	DefineId         *string                                                                               `json:"defineId,omitempty" xml:"defineId,omitempty"`
	Description      *string                                                                               `json:"description,omitempty" xml:"description,omitempty"`
	DisplayType      *string                                                                               `json:"displayType,omitempty" xml:"displayType,omitempty"`
	ExpressionTitle  *string                                                                               `json:"expressionTitle,omitempty" xml:"expressionTitle,omitempty"`
	FavoriteFlag     *bool                                                                                 `json:"favoriteFlag,omitempty" xml:"favoriteFlag,omitempty"`
	FieldDetail      *string                                                                               `json:"fieldDetail,omitempty" xml:"fieldDetail,omitempty"`
	FieldRank        *int32                                                                                `json:"fieldRank,omitempty" xml:"fieldRank,omitempty"`
	FieldSource      *string                                                                               `json:"fieldSource,omitempty" xml:"fieldSource,omitempty"`
	FieldType        *string                                                                               `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	Id               *int64                                                                                `json:"id,omitempty" xml:"id,omitempty"`
	InputFieldType   *string                                                                               `json:"inputFieldType,omitempty" xml:"inputFieldType,omitempty"`
	InputRequired    *string                                                                               `json:"inputRequired,omitempty" xml:"inputRequired,omitempty"`
	Inputs           *string                                                                               `json:"inputs,omitempty" xml:"inputs,omitempty"`
	Name             *string                                                                               `json:"name,omitempty" xml:"name,omitempty"`
	Outlier          *string                                                                               `json:"outlier,omitempty" xml:"outlier,omitempty"`
	OutputThreshold  *DescribeEventVariableListResponseBodyResultObjectExpressionVariablesOutputThreshold  `json:"outputThreshold,omitempty" xml:"outputThreshold,omitempty" type:"Struct"`
	ParentName       *string                                                                               `json:"parentName,omitempty" xml:"parentName,omitempty"`
	SourceType       *string                                                                               `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Title            *string                                                                               `json:"title,omitempty" xml:"title,omitempty"`
	Type             *string                                                                               `json:"type,omitempty" xml:"type,omitempty"`
	VariableVelocity *DescribeEventVariableListResponseBodyResultObjectExpressionVariablesVariableVelocity `json:"variableVelocity,omitempty" xml:"variableVelocity,omitempty" type:"Struct"`
	XLabel           *string                                                                               `json:"xLabel,omitempty" xml:"xLabel,omitempty"`
	YLabel           *string                                                                               `json:"yLabel,omitempty" xml:"yLabel,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectExpressionVariables) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GetCode() *string {
	return s.Code
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GetDataDisplay() *string {
	return s.DataDisplay
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GetDefineId() *string {
	return s.DefineId
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GetDescription() *string {
	return s.Description
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GetDisplayType() *string {
	return s.DisplayType
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GetExpressionTitle() *string {
	return s.ExpressionTitle
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GetFavoriteFlag() *bool {
	return s.FavoriteFlag
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GetFieldDetail() *string {
	return s.FieldDetail
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GetFieldRank() *int32 {
	return s.FieldRank
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GetFieldSource() *string {
	return s.FieldSource
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GetInputFieldType() *string {
	return s.InputFieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GetInputRequired() *string {
	return s.InputRequired
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GetInputs() *string {
	return s.Inputs
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GetName() *string {
	return s.Name
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GetOutlier() *string {
	return s.Outlier
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GetOutputThreshold() *DescribeEventVariableListResponseBodyResultObjectExpressionVariablesOutputThreshold {
	return s.OutputThreshold
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GetParentName() *string {
	return s.ParentName
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GetTitle() *string {
	return s.Title
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GetType() *string {
	return s.Type
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GetVariableVelocity() *DescribeEventVariableListResponseBodyResultObjectExpressionVariablesVariableVelocity {
	return s.VariableVelocity
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GetXLabel() *string {
	return s.XLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) GetYLabel() *string {
	return s.YLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) SetCode(v string) *DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	s.Code = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) SetDataDisplay(v string) *DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	s.DataDisplay = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) SetDefineId(v string) *DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	s.DefineId = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) SetDescription(v string) *DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	s.Description = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) SetDisplayType(v string) *DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	s.DisplayType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) SetExpressionTitle(v string) *DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	s.ExpressionTitle = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) SetFavoriteFlag(v bool) *DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	s.FavoriteFlag = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) SetFieldDetail(v string) *DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	s.FieldDetail = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) SetFieldRank(v int32) *DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	s.FieldRank = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) SetFieldSource(v string) *DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	s.FieldSource = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) SetFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	s.FieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) SetId(v int64) *DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	s.Id = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) SetInputFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	s.InputFieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) SetInputRequired(v string) *DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	s.InputRequired = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) SetInputs(v string) *DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	s.Inputs = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) SetName(v string) *DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	s.Name = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) SetOutlier(v string) *DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	s.Outlier = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) SetOutputThreshold(v *DescribeEventVariableListResponseBodyResultObjectExpressionVariablesOutputThreshold) *DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	s.OutputThreshold = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) SetParentName(v string) *DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	s.ParentName = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) SetSourceType(v string) *DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	s.SourceType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) SetTitle(v string) *DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	s.Title = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) SetType(v string) *DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	s.Type = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) SetVariableVelocity(v *DescribeEventVariableListResponseBodyResultObjectExpressionVariablesVariableVelocity) *DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	s.VariableVelocity = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) SetXLabel(v string) *DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	s.XLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) SetYLabel(v string) *DescribeEventVariableListResponseBodyResultObjectExpressionVariables {
	s.YLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariables) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectExpressionVariablesOutputThreshold struct {
	MaxValue *float64 `json:"maxValue,omitempty" xml:"maxValue,omitempty"`
	MinValue *float64 `json:"minValue,omitempty" xml:"minValue,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectExpressionVariablesOutputThreshold) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectExpressionVariablesOutputThreshold) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariablesOutputThreshold) GetMaxValue() *float64 {
	return s.MaxValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariablesOutputThreshold) GetMinValue() *float64 {
	return s.MinValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariablesOutputThreshold) SetMaxValue(v float64) *DescribeEventVariableListResponseBodyResultObjectExpressionVariablesOutputThreshold {
	s.MaxValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariablesOutputThreshold) SetMinValue(v float64) *DescribeEventVariableListResponseBodyResultObjectExpressionVariablesOutputThreshold {
	s.MinValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariablesOutputThreshold) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectExpressionVariablesVariableVelocity struct {
	Iv *string `json:"iv,omitempty" xml:"iv,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectExpressionVariablesVariableVelocity) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectExpressionVariablesVariableVelocity) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariablesVariableVelocity) GetIv() *string {
	return s.Iv
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariablesVariableVelocity) SetIv(v string) *DescribeEventVariableListResponseBodyResultObjectExpressionVariablesVariableVelocity {
	s.Iv = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectExpressionVariablesVariableVelocity) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectFavoriteVariables struct {
	Code             *string                                                                             `json:"code,omitempty" xml:"code,omitempty"`
	DataDisplay      *string                                                                             `json:"dataDisplay,omitempty" xml:"dataDisplay,omitempty"`
	DefineId         *string                                                                             `json:"defineId,omitempty" xml:"defineId,omitempty"`
	Description      *string                                                                             `json:"description,omitempty" xml:"description,omitempty"`
	DisplayType      *string                                                                             `json:"displayType,omitempty" xml:"displayType,omitempty"`
	ExpressionTitle  *string                                                                             `json:"expressionTitle,omitempty" xml:"expressionTitle,omitempty"`
	FavoriteFlag     *bool                                                                               `json:"favoriteFlag,omitempty" xml:"favoriteFlag,omitempty"`
	FieldDetail      *string                                                                             `json:"fieldDetail,omitempty" xml:"fieldDetail,omitempty"`
	FieldRank        *int32                                                                              `json:"fieldRank,omitempty" xml:"fieldRank,omitempty"`
	FieldSource      *string                                                                             `json:"fieldSource,omitempty" xml:"fieldSource,omitempty"`
	FieldType        *string                                                                             `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	Id               *int64                                                                              `json:"id,omitempty" xml:"id,omitempty"`
	InputFieldType   *string                                                                             `json:"inputFieldType,omitempty" xml:"inputFieldType,omitempty"`
	InputRequired    *string                                                                             `json:"inputRequired,omitempty" xml:"inputRequired,omitempty"`
	Inputs           *string                                                                             `json:"inputs,omitempty" xml:"inputs,omitempty"`
	Name             *string                                                                             `json:"name,omitempty" xml:"name,omitempty"`
	Outlier          *string                                                                             `json:"outlier,omitempty" xml:"outlier,omitempty"`
	OutputThreshold  *DescribeEventVariableListResponseBodyResultObjectFavoriteVariablesOutputThreshold  `json:"outputThreshold,omitempty" xml:"outputThreshold,omitempty" type:"Struct"`
	ParentName       *string                                                                             `json:"parentName,omitempty" xml:"parentName,omitempty"`
	SourceType       *string                                                                             `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Title            *string                                                                             `json:"title,omitempty" xml:"title,omitempty"`
	Type             *string                                                                             `json:"type,omitempty" xml:"type,omitempty"`
	VariableVelocity *DescribeEventVariableListResponseBodyResultObjectFavoriteVariablesVariableVelocity `json:"variableVelocity,omitempty" xml:"variableVelocity,omitempty" type:"Struct"`
	XLabel           *string                                                                             `json:"xLabel,omitempty" xml:"xLabel,omitempty"`
	YLabel           *string                                                                             `json:"yLabel,omitempty" xml:"yLabel,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GetCode() *string {
	return s.Code
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GetDataDisplay() *string {
	return s.DataDisplay
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GetDefineId() *string {
	return s.DefineId
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GetDescription() *string {
	return s.Description
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GetDisplayType() *string {
	return s.DisplayType
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GetExpressionTitle() *string {
	return s.ExpressionTitle
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GetFavoriteFlag() *bool {
	return s.FavoriteFlag
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GetFieldDetail() *string {
	return s.FieldDetail
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GetFieldRank() *int32 {
	return s.FieldRank
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GetFieldSource() *string {
	return s.FieldSource
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GetInputFieldType() *string {
	return s.InputFieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GetInputRequired() *string {
	return s.InputRequired
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GetInputs() *string {
	return s.Inputs
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GetName() *string {
	return s.Name
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GetOutlier() *string {
	return s.Outlier
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GetOutputThreshold() *DescribeEventVariableListResponseBodyResultObjectFavoriteVariablesOutputThreshold {
	return s.OutputThreshold
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GetParentName() *string {
	return s.ParentName
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GetTitle() *string {
	return s.Title
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GetType() *string {
	return s.Type
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GetVariableVelocity() *DescribeEventVariableListResponseBodyResultObjectFavoriteVariablesVariableVelocity {
	return s.VariableVelocity
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GetXLabel() *string {
	return s.XLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) GetYLabel() *string {
	return s.YLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) SetCode(v string) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	s.Code = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) SetDataDisplay(v string) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	s.DataDisplay = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) SetDefineId(v string) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	s.DefineId = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) SetDescription(v string) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	s.Description = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) SetDisplayType(v string) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	s.DisplayType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) SetExpressionTitle(v string) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	s.ExpressionTitle = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) SetFavoriteFlag(v bool) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	s.FavoriteFlag = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) SetFieldDetail(v string) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	s.FieldDetail = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) SetFieldRank(v int32) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	s.FieldRank = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) SetFieldSource(v string) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	s.FieldSource = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) SetFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	s.FieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) SetId(v int64) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	s.Id = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) SetInputFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	s.InputFieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) SetInputRequired(v string) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	s.InputRequired = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) SetInputs(v string) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	s.Inputs = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) SetName(v string) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	s.Name = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) SetOutlier(v string) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	s.Outlier = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) SetOutputThreshold(v *DescribeEventVariableListResponseBodyResultObjectFavoriteVariablesOutputThreshold) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	s.OutputThreshold = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) SetParentName(v string) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	s.ParentName = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) SetSourceType(v string) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	s.SourceType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) SetTitle(v string) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	s.Title = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) SetType(v string) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	s.Type = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) SetVariableVelocity(v *DescribeEventVariableListResponseBodyResultObjectFavoriteVariablesVariableVelocity) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	s.VariableVelocity = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) SetXLabel(v string) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	s.XLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) SetYLabel(v string) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables {
	s.YLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariables) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectFavoriteVariablesOutputThreshold struct {
	MaxValue *float64 `json:"maxValue,omitempty" xml:"maxValue,omitempty"`
	MinValue *float64 `json:"minValue,omitempty" xml:"minValue,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectFavoriteVariablesOutputThreshold) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectFavoriteVariablesOutputThreshold) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariablesOutputThreshold) GetMaxValue() *float64 {
	return s.MaxValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariablesOutputThreshold) GetMinValue() *float64 {
	return s.MinValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariablesOutputThreshold) SetMaxValue(v float64) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariablesOutputThreshold {
	s.MaxValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariablesOutputThreshold) SetMinValue(v float64) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariablesOutputThreshold {
	s.MinValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariablesOutputThreshold) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectFavoriteVariablesVariableVelocity struct {
	Iv *string `json:"iv,omitempty" xml:"iv,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectFavoriteVariablesVariableVelocity) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectFavoriteVariablesVariableVelocity) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariablesVariableVelocity) GetIv() *string {
	return s.Iv
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariablesVariableVelocity) SetIv(v string) *DescribeEventVariableListResponseBodyResultObjectFavoriteVariablesVariableVelocity {
	s.Iv = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectFavoriteVariablesVariableVelocity) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectMiddleVariables struct {
	Code             *string                                                                           `json:"code,omitempty" xml:"code,omitempty"`
	DataDisplay      *string                                                                           `json:"dataDisplay,omitempty" xml:"dataDisplay,omitempty"`
	DefineId         *string                                                                           `json:"defineId,omitempty" xml:"defineId,omitempty"`
	Description      *string                                                                           `json:"description,omitempty" xml:"description,omitempty"`
	DisplayType      *string                                                                           `json:"displayType,omitempty" xml:"displayType,omitempty"`
	ExpressionTitle  *string                                                                           `json:"expressionTitle,omitempty" xml:"expressionTitle,omitempty"`
	FavoriteFlag     *bool                                                                             `json:"favoriteFlag,omitempty" xml:"favoriteFlag,omitempty"`
	FieldDetail      *string                                                                           `json:"fieldDetail,omitempty" xml:"fieldDetail,omitempty"`
	FieldRank        *int32                                                                            `json:"fieldRank,omitempty" xml:"fieldRank,omitempty"`
	FieldSource      *string                                                                           `json:"fieldSource,omitempty" xml:"fieldSource,omitempty"`
	FieldType        *string                                                                           `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	Id               *int64                                                                            `json:"id,omitempty" xml:"id,omitempty"`
	InputFieldType   *string                                                                           `json:"inputFieldType,omitempty" xml:"inputFieldType,omitempty"`
	InputRequired    *string                                                                           `json:"inputRequired,omitempty" xml:"inputRequired,omitempty"`
	Inputs           *string                                                                           `json:"inputs,omitempty" xml:"inputs,omitempty"`
	Name             *string                                                                           `json:"name,omitempty" xml:"name,omitempty"`
	Outlier          *string                                                                           `json:"outlier,omitempty" xml:"outlier,omitempty"`
	OutputThreshold  *DescribeEventVariableListResponseBodyResultObjectMiddleVariablesOutputThreshold  `json:"outputThreshold,omitempty" xml:"outputThreshold,omitempty" type:"Struct"`
	ParentName       *string                                                                           `json:"parentName,omitempty" xml:"parentName,omitempty"`
	SourceType       *string                                                                           `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Title            *string                                                                           `json:"title,omitempty" xml:"title,omitempty"`
	Type             *string                                                                           `json:"type,omitempty" xml:"type,omitempty"`
	VariableVelocity *DescribeEventVariableListResponseBodyResultObjectMiddleVariablesVariableVelocity `json:"variableVelocity,omitempty" xml:"variableVelocity,omitempty" type:"Struct"`
	XLabel           *string                                                                           `json:"xLabel,omitempty" xml:"xLabel,omitempty"`
	YLabel           *string                                                                           `json:"yLabel,omitempty" xml:"yLabel,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectMiddleVariables) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GetCode() *string {
	return s.Code
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GetDataDisplay() *string {
	return s.DataDisplay
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GetDefineId() *string {
	return s.DefineId
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GetDescription() *string {
	return s.Description
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GetDisplayType() *string {
	return s.DisplayType
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GetExpressionTitle() *string {
	return s.ExpressionTitle
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GetFavoriteFlag() *bool {
	return s.FavoriteFlag
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GetFieldDetail() *string {
	return s.FieldDetail
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GetFieldRank() *int32 {
	return s.FieldRank
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GetFieldSource() *string {
	return s.FieldSource
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GetInputFieldType() *string {
	return s.InputFieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GetInputRequired() *string {
	return s.InputRequired
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GetInputs() *string {
	return s.Inputs
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GetName() *string {
	return s.Name
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GetOutlier() *string {
	return s.Outlier
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GetOutputThreshold() *DescribeEventVariableListResponseBodyResultObjectMiddleVariablesOutputThreshold {
	return s.OutputThreshold
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GetParentName() *string {
	return s.ParentName
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GetTitle() *string {
	return s.Title
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GetType() *string {
	return s.Type
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GetVariableVelocity() *DescribeEventVariableListResponseBodyResultObjectMiddleVariablesVariableVelocity {
	return s.VariableVelocity
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GetXLabel() *string {
	return s.XLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) GetYLabel() *string {
	return s.YLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) SetCode(v string) *DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	s.Code = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) SetDataDisplay(v string) *DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	s.DataDisplay = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) SetDefineId(v string) *DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	s.DefineId = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) SetDescription(v string) *DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	s.Description = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) SetDisplayType(v string) *DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	s.DisplayType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) SetExpressionTitle(v string) *DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	s.ExpressionTitle = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) SetFavoriteFlag(v bool) *DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	s.FavoriteFlag = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) SetFieldDetail(v string) *DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	s.FieldDetail = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) SetFieldRank(v int32) *DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	s.FieldRank = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) SetFieldSource(v string) *DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	s.FieldSource = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) SetFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	s.FieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) SetId(v int64) *DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	s.Id = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) SetInputFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	s.InputFieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) SetInputRequired(v string) *DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	s.InputRequired = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) SetInputs(v string) *DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	s.Inputs = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) SetName(v string) *DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	s.Name = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) SetOutlier(v string) *DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	s.Outlier = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) SetOutputThreshold(v *DescribeEventVariableListResponseBodyResultObjectMiddleVariablesOutputThreshold) *DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	s.OutputThreshold = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) SetParentName(v string) *DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	s.ParentName = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) SetSourceType(v string) *DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	s.SourceType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) SetTitle(v string) *DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	s.Title = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) SetType(v string) *DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	s.Type = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) SetVariableVelocity(v *DescribeEventVariableListResponseBodyResultObjectMiddleVariablesVariableVelocity) *DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	s.VariableVelocity = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) SetXLabel(v string) *DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	s.XLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) SetYLabel(v string) *DescribeEventVariableListResponseBodyResultObjectMiddleVariables {
	s.YLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariables) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectMiddleVariablesOutputThreshold struct {
	MaxValue *float64 `json:"maxValue,omitempty" xml:"maxValue,omitempty"`
	MinValue *float64 `json:"minValue,omitempty" xml:"minValue,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectMiddleVariablesOutputThreshold) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectMiddleVariablesOutputThreshold) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariablesOutputThreshold) GetMaxValue() *float64 {
	return s.MaxValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariablesOutputThreshold) GetMinValue() *float64 {
	return s.MinValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariablesOutputThreshold) SetMaxValue(v float64) *DescribeEventVariableListResponseBodyResultObjectMiddleVariablesOutputThreshold {
	s.MaxValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariablesOutputThreshold) SetMinValue(v float64) *DescribeEventVariableListResponseBodyResultObjectMiddleVariablesOutputThreshold {
	s.MinValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariablesOutputThreshold) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectMiddleVariablesVariableVelocity struct {
	Iv *string `json:"iv,omitempty" xml:"iv,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectMiddleVariablesVariableVelocity) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectMiddleVariablesVariableVelocity) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariablesVariableVelocity) GetIv() *string {
	return s.Iv
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariablesVariableVelocity) SetIv(v string) *DescribeEventVariableListResponseBodyResultObjectMiddleVariablesVariableVelocity {
	s.Iv = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectMiddleVariablesVariableVelocity) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectModelVariables struct {
	Code             *string                                                                          `json:"code,omitempty" xml:"code,omitempty"`
	DataDisplay      *string                                                                          `json:"dataDisplay,omitempty" xml:"dataDisplay,omitempty"`
	DefineId         *string                                                                          `json:"defineId,omitempty" xml:"defineId,omitempty"`
	Description      *string                                                                          `json:"description,omitempty" xml:"description,omitempty"`
	DisplayType      *string                                                                          `json:"displayType,omitempty" xml:"displayType,omitempty"`
	ExpressionTitle  *string                                                                          `json:"expressionTitle,omitempty" xml:"expressionTitle,omitempty"`
	FavoriteFlag     *bool                                                                            `json:"favoriteFlag,omitempty" xml:"favoriteFlag,omitempty"`
	FieldDetail      *string                                                                          `json:"fieldDetail,omitempty" xml:"fieldDetail,omitempty"`
	FieldRank        *int32                                                                           `json:"fieldRank,omitempty" xml:"fieldRank,omitempty"`
	FieldSource      *string                                                                          `json:"fieldSource,omitempty" xml:"fieldSource,omitempty"`
	FieldType        *string                                                                          `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	Id               *int64                                                                           `json:"id,omitempty" xml:"id,omitempty"`
	InputFieldType   *string                                                                          `json:"inputFieldType,omitempty" xml:"inputFieldType,omitempty"`
	InputRequired    *string                                                                          `json:"inputRequired,omitempty" xml:"inputRequired,omitempty"`
	Inputs           *string                                                                          `json:"inputs,omitempty" xml:"inputs,omitempty"`
	Name             *string                                                                          `json:"name,omitempty" xml:"name,omitempty"`
	Outlier          *string                                                                          `json:"outlier,omitempty" xml:"outlier,omitempty"`
	OutputThreshold  *DescribeEventVariableListResponseBodyResultObjectModelVariablesOutputThreshold  `json:"outputThreshold,omitempty" xml:"outputThreshold,omitempty" type:"Struct"`
	ParentName       *string                                                                          `json:"parentName,omitempty" xml:"parentName,omitempty"`
	SourceType       *string                                                                          `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Title            *string                                                                          `json:"title,omitempty" xml:"title,omitempty"`
	Type             *string                                                                          `json:"type,omitempty" xml:"type,omitempty"`
	VariableVelocity *DescribeEventVariableListResponseBodyResultObjectModelVariablesVariableVelocity `json:"variableVelocity,omitempty" xml:"variableVelocity,omitempty" type:"Struct"`
	XLabel           *string                                                                          `json:"xLabel,omitempty" xml:"xLabel,omitempty"`
	YLabel           *string                                                                          `json:"yLabel,omitempty" xml:"yLabel,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectModelVariables) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectModelVariables) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) GetCode() *string {
	return s.Code
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) GetDataDisplay() *string {
	return s.DataDisplay
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) GetDefineId() *string {
	return s.DefineId
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) GetDescription() *string {
	return s.Description
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) GetDisplayType() *string {
	return s.DisplayType
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) GetExpressionTitle() *string {
	return s.ExpressionTitle
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) GetFavoriteFlag() *bool {
	return s.FavoriteFlag
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) GetFieldDetail() *string {
	return s.FieldDetail
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) GetFieldRank() *int32 {
	return s.FieldRank
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) GetFieldSource() *string {
	return s.FieldSource
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) GetInputFieldType() *string {
	return s.InputFieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) GetInputRequired() *string {
	return s.InputRequired
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) GetInputs() *string {
	return s.Inputs
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) GetName() *string {
	return s.Name
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) GetOutlier() *string {
	return s.Outlier
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) GetOutputThreshold() *DescribeEventVariableListResponseBodyResultObjectModelVariablesOutputThreshold {
	return s.OutputThreshold
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) GetParentName() *string {
	return s.ParentName
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) GetTitle() *string {
	return s.Title
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) GetType() *string {
	return s.Type
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) GetVariableVelocity() *DescribeEventVariableListResponseBodyResultObjectModelVariablesVariableVelocity {
	return s.VariableVelocity
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) GetXLabel() *string {
	return s.XLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) GetYLabel() *string {
	return s.YLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) SetCode(v string) *DescribeEventVariableListResponseBodyResultObjectModelVariables {
	s.Code = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) SetDataDisplay(v string) *DescribeEventVariableListResponseBodyResultObjectModelVariables {
	s.DataDisplay = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) SetDefineId(v string) *DescribeEventVariableListResponseBodyResultObjectModelVariables {
	s.DefineId = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) SetDescription(v string) *DescribeEventVariableListResponseBodyResultObjectModelVariables {
	s.Description = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) SetDisplayType(v string) *DescribeEventVariableListResponseBodyResultObjectModelVariables {
	s.DisplayType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) SetExpressionTitle(v string) *DescribeEventVariableListResponseBodyResultObjectModelVariables {
	s.ExpressionTitle = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) SetFavoriteFlag(v bool) *DescribeEventVariableListResponseBodyResultObjectModelVariables {
	s.FavoriteFlag = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) SetFieldDetail(v string) *DescribeEventVariableListResponseBodyResultObjectModelVariables {
	s.FieldDetail = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) SetFieldRank(v int32) *DescribeEventVariableListResponseBodyResultObjectModelVariables {
	s.FieldRank = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) SetFieldSource(v string) *DescribeEventVariableListResponseBodyResultObjectModelVariables {
	s.FieldSource = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) SetFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectModelVariables {
	s.FieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) SetId(v int64) *DescribeEventVariableListResponseBodyResultObjectModelVariables {
	s.Id = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) SetInputFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectModelVariables {
	s.InputFieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) SetInputRequired(v string) *DescribeEventVariableListResponseBodyResultObjectModelVariables {
	s.InputRequired = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) SetInputs(v string) *DescribeEventVariableListResponseBodyResultObjectModelVariables {
	s.Inputs = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) SetName(v string) *DescribeEventVariableListResponseBodyResultObjectModelVariables {
	s.Name = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) SetOutlier(v string) *DescribeEventVariableListResponseBodyResultObjectModelVariables {
	s.Outlier = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) SetOutputThreshold(v *DescribeEventVariableListResponseBodyResultObjectModelVariablesOutputThreshold) *DescribeEventVariableListResponseBodyResultObjectModelVariables {
	s.OutputThreshold = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) SetParentName(v string) *DescribeEventVariableListResponseBodyResultObjectModelVariables {
	s.ParentName = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) SetSourceType(v string) *DescribeEventVariableListResponseBodyResultObjectModelVariables {
	s.SourceType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) SetTitle(v string) *DescribeEventVariableListResponseBodyResultObjectModelVariables {
	s.Title = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) SetType(v string) *DescribeEventVariableListResponseBodyResultObjectModelVariables {
	s.Type = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) SetVariableVelocity(v *DescribeEventVariableListResponseBodyResultObjectModelVariablesVariableVelocity) *DescribeEventVariableListResponseBodyResultObjectModelVariables {
	s.VariableVelocity = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) SetXLabel(v string) *DescribeEventVariableListResponseBodyResultObjectModelVariables {
	s.XLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) SetYLabel(v string) *DescribeEventVariableListResponseBodyResultObjectModelVariables {
	s.YLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariables) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectModelVariablesOutputThreshold struct {
	MaxValue *float64 `json:"maxValue,omitempty" xml:"maxValue,omitempty"`
	MinValue *float64 `json:"minValue,omitempty" xml:"minValue,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectModelVariablesOutputThreshold) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectModelVariablesOutputThreshold) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariablesOutputThreshold) GetMaxValue() *float64 {
	return s.MaxValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariablesOutputThreshold) GetMinValue() *float64 {
	return s.MinValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariablesOutputThreshold) SetMaxValue(v float64) *DescribeEventVariableListResponseBodyResultObjectModelVariablesOutputThreshold {
	s.MaxValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariablesOutputThreshold) SetMinValue(v float64) *DescribeEventVariableListResponseBodyResultObjectModelVariablesOutputThreshold {
	s.MinValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariablesOutputThreshold) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectModelVariablesVariableVelocity struct {
	Iv *string `json:"iv,omitempty" xml:"iv,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectModelVariablesVariableVelocity) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectModelVariablesVariableVelocity) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariablesVariableVelocity) GetIv() *string {
	return s.Iv
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariablesVariableVelocity) SetIv(v string) *DescribeEventVariableListResponseBodyResultObjectModelVariablesVariableVelocity {
	s.Iv = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectModelVariablesVariableVelocity) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectNameList struct {
	Code             *string                                                                    `json:"code,omitempty" xml:"code,omitempty"`
	DataDisplay      *string                                                                    `json:"dataDisplay,omitempty" xml:"dataDisplay,omitempty"`
	DefineId         *string                                                                    `json:"defineId,omitempty" xml:"defineId,omitempty"`
	Description      *string                                                                    `json:"description,omitempty" xml:"description,omitempty"`
	DisplayType      *string                                                                    `json:"displayType,omitempty" xml:"displayType,omitempty"`
	ExpressionTitle  *string                                                                    `json:"expressionTitle,omitempty" xml:"expressionTitle,omitempty"`
	FavoriteFlag     *bool                                                                      `json:"favoriteFlag,omitempty" xml:"favoriteFlag,omitempty"`
	FieldDetail      *string                                                                    `json:"fieldDetail,omitempty" xml:"fieldDetail,omitempty"`
	FieldRank        *int32                                                                     `json:"fieldRank,omitempty" xml:"fieldRank,omitempty"`
	FieldSource      *string                                                                    `json:"fieldSource,omitempty" xml:"fieldSource,omitempty"`
	FieldType        *string                                                                    `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	Id               *int64                                                                     `json:"id,omitempty" xml:"id,omitempty"`
	InputFieldType   *string                                                                    `json:"inputFieldType,omitempty" xml:"inputFieldType,omitempty"`
	InputRequired    *string                                                                    `json:"inputRequired,omitempty" xml:"inputRequired,omitempty"`
	Inputs           *string                                                                    `json:"inputs,omitempty" xml:"inputs,omitempty"`
	Name             *string                                                                    `json:"name,omitempty" xml:"name,omitempty"`
	Outlier          *string                                                                    `json:"outlier,omitempty" xml:"outlier,omitempty"`
	OutputThreshold  *DescribeEventVariableListResponseBodyResultObjectNameListOutputThreshold  `json:"outputThreshold,omitempty" xml:"outputThreshold,omitempty" type:"Struct"`
	ParentName       *string                                                                    `json:"parentName,omitempty" xml:"parentName,omitempty"`
	SourceType       *string                                                                    `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Title            *string                                                                    `json:"title,omitempty" xml:"title,omitempty"`
	Type             *string                                                                    `json:"type,omitempty" xml:"type,omitempty"`
	VariableVelocity *DescribeEventVariableListResponseBodyResultObjectNameListVariableVelocity `json:"variableVelocity,omitempty" xml:"variableVelocity,omitempty" type:"Struct"`
	XLabel           *string                                                                    `json:"xLabel,omitempty" xml:"xLabel,omitempty"`
	YLabel           *string                                                                    `json:"yLabel,omitempty" xml:"yLabel,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectNameList) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectNameList) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) GetCode() *string {
	return s.Code
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) GetDataDisplay() *string {
	return s.DataDisplay
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) GetDefineId() *string {
	return s.DefineId
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) GetDescription() *string {
	return s.Description
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) GetDisplayType() *string {
	return s.DisplayType
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) GetExpressionTitle() *string {
	return s.ExpressionTitle
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) GetFavoriteFlag() *bool {
	return s.FavoriteFlag
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) GetFieldDetail() *string {
	return s.FieldDetail
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) GetFieldRank() *int32 {
	return s.FieldRank
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) GetFieldSource() *string {
	return s.FieldSource
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) GetInputFieldType() *string {
	return s.InputFieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) GetInputRequired() *string {
	return s.InputRequired
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) GetInputs() *string {
	return s.Inputs
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) GetName() *string {
	return s.Name
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) GetOutlier() *string {
	return s.Outlier
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) GetOutputThreshold() *DescribeEventVariableListResponseBodyResultObjectNameListOutputThreshold {
	return s.OutputThreshold
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) GetParentName() *string {
	return s.ParentName
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) GetTitle() *string {
	return s.Title
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) GetType() *string {
	return s.Type
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) GetVariableVelocity() *DescribeEventVariableListResponseBodyResultObjectNameListVariableVelocity {
	return s.VariableVelocity
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) GetXLabel() *string {
	return s.XLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) GetYLabel() *string {
	return s.YLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) SetCode(v string) *DescribeEventVariableListResponseBodyResultObjectNameList {
	s.Code = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) SetDataDisplay(v string) *DescribeEventVariableListResponseBodyResultObjectNameList {
	s.DataDisplay = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) SetDefineId(v string) *DescribeEventVariableListResponseBodyResultObjectNameList {
	s.DefineId = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) SetDescription(v string) *DescribeEventVariableListResponseBodyResultObjectNameList {
	s.Description = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) SetDisplayType(v string) *DescribeEventVariableListResponseBodyResultObjectNameList {
	s.DisplayType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) SetExpressionTitle(v string) *DescribeEventVariableListResponseBodyResultObjectNameList {
	s.ExpressionTitle = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) SetFavoriteFlag(v bool) *DescribeEventVariableListResponseBodyResultObjectNameList {
	s.FavoriteFlag = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) SetFieldDetail(v string) *DescribeEventVariableListResponseBodyResultObjectNameList {
	s.FieldDetail = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) SetFieldRank(v int32) *DescribeEventVariableListResponseBodyResultObjectNameList {
	s.FieldRank = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) SetFieldSource(v string) *DescribeEventVariableListResponseBodyResultObjectNameList {
	s.FieldSource = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) SetFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectNameList {
	s.FieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) SetId(v int64) *DescribeEventVariableListResponseBodyResultObjectNameList {
	s.Id = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) SetInputFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectNameList {
	s.InputFieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) SetInputRequired(v string) *DescribeEventVariableListResponseBodyResultObjectNameList {
	s.InputRequired = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) SetInputs(v string) *DescribeEventVariableListResponseBodyResultObjectNameList {
	s.Inputs = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) SetName(v string) *DescribeEventVariableListResponseBodyResultObjectNameList {
	s.Name = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) SetOutlier(v string) *DescribeEventVariableListResponseBodyResultObjectNameList {
	s.Outlier = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) SetOutputThreshold(v *DescribeEventVariableListResponseBodyResultObjectNameListOutputThreshold) *DescribeEventVariableListResponseBodyResultObjectNameList {
	s.OutputThreshold = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) SetParentName(v string) *DescribeEventVariableListResponseBodyResultObjectNameList {
	s.ParentName = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) SetSourceType(v string) *DescribeEventVariableListResponseBodyResultObjectNameList {
	s.SourceType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) SetTitle(v string) *DescribeEventVariableListResponseBodyResultObjectNameList {
	s.Title = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) SetType(v string) *DescribeEventVariableListResponseBodyResultObjectNameList {
	s.Type = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) SetVariableVelocity(v *DescribeEventVariableListResponseBodyResultObjectNameListVariableVelocity) *DescribeEventVariableListResponseBodyResultObjectNameList {
	s.VariableVelocity = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) SetXLabel(v string) *DescribeEventVariableListResponseBodyResultObjectNameList {
	s.XLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) SetYLabel(v string) *DescribeEventVariableListResponseBodyResultObjectNameList {
	s.YLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameList) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectNameListOutputThreshold struct {
	MaxValue *float64 `json:"maxValue,omitempty" xml:"maxValue,omitempty"`
	MinValue *float64 `json:"minValue,omitempty" xml:"minValue,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectNameListOutputThreshold) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectNameListOutputThreshold) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameListOutputThreshold) GetMaxValue() *float64 {
	return s.MaxValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameListOutputThreshold) GetMinValue() *float64 {
	return s.MinValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameListOutputThreshold) SetMaxValue(v float64) *DescribeEventVariableListResponseBodyResultObjectNameListOutputThreshold {
	s.MaxValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameListOutputThreshold) SetMinValue(v float64) *DescribeEventVariableListResponseBodyResultObjectNameListOutputThreshold {
	s.MinValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameListOutputThreshold) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectNameListVariableVelocity struct {
	Iv *string `json:"iv,omitempty" xml:"iv,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectNameListVariableVelocity) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectNameListVariableVelocity) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameListVariableVelocity) GetIv() *string {
	return s.Iv
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameListVariableVelocity) SetIv(v string) *DescribeEventVariableListResponseBodyResultObjectNameListVariableVelocity {
	s.Iv = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNameListVariableVelocity) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions struct {
	Code             *string                                                                                   `json:"code,omitempty" xml:"code,omitempty"`
	DataDisplay      *string                                                                                   `json:"dataDisplay,omitempty" xml:"dataDisplay,omitempty"`
	DefineId         *string                                                                                   `json:"defineId,omitempty" xml:"defineId,omitempty"`
	Description      *string                                                                                   `json:"description,omitempty" xml:"description,omitempty"`
	DisplayType      *string                                                                                   `json:"displayType,omitempty" xml:"displayType,omitempty"`
	ExpressionTitle  *string                                                                                   `json:"expressionTitle,omitempty" xml:"expressionTitle,omitempty"`
	FavoriteFlag     *bool                                                                                     `json:"favoriteFlag,omitempty" xml:"favoriteFlag,omitempty"`
	FieldDetail      *string                                                                                   `json:"fieldDetail,omitempty" xml:"fieldDetail,omitempty"`
	FieldRank        *int32                                                                                    `json:"fieldRank,omitempty" xml:"fieldRank,omitempty"`
	FieldSource      *string                                                                                   `json:"fieldSource,omitempty" xml:"fieldSource,omitempty"`
	FieldType        *string                                                                                   `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	Id               *int64                                                                                    `json:"id,omitempty" xml:"id,omitempty"`
	InputFieldType   *string                                                                                   `json:"inputFieldType,omitempty" xml:"inputFieldType,omitempty"`
	InputRequired    *string                                                                                   `json:"inputRequired,omitempty" xml:"inputRequired,omitempty"`
	Inputs           *string                                                                                   `json:"inputs,omitempty" xml:"inputs,omitempty"`
	Name             *string                                                                                   `json:"name,omitempty" xml:"name,omitempty"`
	Outlier          *string                                                                                   `json:"outlier,omitempty" xml:"outlier,omitempty"`
	OutputThreshold  *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctionsOutputThreshold  `json:"outputThreshold,omitempty" xml:"outputThreshold,omitempty" type:"Struct"`
	ParentName       *string                                                                                   `json:"parentName,omitempty" xml:"parentName,omitempty"`
	SourceType       *string                                                                                   `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Title            *string                                                                                   `json:"title,omitempty" xml:"title,omitempty"`
	Type             *string                                                                                   `json:"type,omitempty" xml:"type,omitempty"`
	VariableVelocity *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctionsVariableVelocity `json:"variableVelocity,omitempty" xml:"variableVelocity,omitempty" type:"Struct"`
	XLabel           *string                                                                                   `json:"xLabel,omitempty" xml:"xLabel,omitempty"`
	YLabel           *string                                                                                   `json:"yLabel,omitempty" xml:"yLabel,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GetCode() *string {
	return s.Code
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GetDataDisplay() *string {
	return s.DataDisplay
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GetDefineId() *string {
	return s.DefineId
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GetDescription() *string {
	return s.Description
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GetDisplayType() *string {
	return s.DisplayType
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GetExpressionTitle() *string {
	return s.ExpressionTitle
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GetFavoriteFlag() *bool {
	return s.FavoriteFlag
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GetFieldDetail() *string {
	return s.FieldDetail
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GetFieldRank() *int32 {
	return s.FieldRank
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GetFieldSource() *string {
	return s.FieldSource
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GetInputFieldType() *string {
	return s.InputFieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GetInputRequired() *string {
	return s.InputRequired
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GetInputs() *string {
	return s.Inputs
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GetName() *string {
	return s.Name
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GetOutlier() *string {
	return s.Outlier
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GetOutputThreshold() *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctionsOutputThreshold {
	return s.OutputThreshold
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GetParentName() *string {
	return s.ParentName
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GetTitle() *string {
	return s.Title
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GetType() *string {
	return s.Type
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GetVariableVelocity() *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctionsVariableVelocity {
	return s.VariableVelocity
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GetXLabel() *string {
	return s.XLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) GetYLabel() *string {
	return s.YLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) SetCode(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	s.Code = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) SetDataDisplay(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	s.DataDisplay = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) SetDefineId(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	s.DefineId = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) SetDescription(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	s.Description = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) SetDisplayType(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	s.DisplayType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) SetExpressionTitle(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	s.ExpressionTitle = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) SetFavoriteFlag(v bool) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	s.FavoriteFlag = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) SetFieldDetail(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	s.FieldDetail = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) SetFieldRank(v int32) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	s.FieldRank = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) SetFieldSource(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	s.FieldSource = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) SetFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	s.FieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) SetId(v int64) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	s.Id = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) SetInputFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	s.InputFieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) SetInputRequired(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	s.InputRequired = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) SetInputs(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	s.Inputs = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) SetName(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	s.Name = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) SetOutlier(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	s.Outlier = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) SetOutputThreshold(v *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctionsOutputThreshold) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	s.OutputThreshold = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) SetParentName(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	s.ParentName = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) SetSourceType(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	s.SourceType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) SetTitle(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	s.Title = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) SetType(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	s.Type = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) SetVariableVelocity(v *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctionsVariableVelocity) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	s.VariableVelocity = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) SetXLabel(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	s.XLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) SetYLabel(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions {
	s.YLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctions) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctionsOutputThreshold struct {
	MaxValue *float64 `json:"maxValue,omitempty" xml:"maxValue,omitempty"`
	MinValue *float64 `json:"minValue,omitempty" xml:"minValue,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctionsOutputThreshold) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctionsOutputThreshold) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctionsOutputThreshold) GetMaxValue() *float64 {
	return s.MaxValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctionsOutputThreshold) GetMinValue() *float64 {
	return s.MinValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctionsOutputThreshold) SetMaxValue(v float64) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctionsOutputThreshold {
	s.MaxValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctionsOutputThreshold) SetMinValue(v float64) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctionsOutputThreshold {
	s.MinValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctionsOutputThreshold) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctionsVariableVelocity struct {
	Iv *string `json:"iv,omitempty" xml:"iv,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctionsVariableVelocity) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctionsVariableVelocity) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctionsVariableVelocity) GetIv() *string {
	return s.Iv
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctionsVariableVelocity) SetIv(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctionsVariableVelocity {
	s.Iv = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariableFunctionsVariableVelocity) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectNativeVariables struct {
	Code             *string                                                                           `json:"code,omitempty" xml:"code,omitempty"`
	DataDisplay      *string                                                                           `json:"dataDisplay,omitempty" xml:"dataDisplay,omitempty"`
	DefineId         *string                                                                           `json:"defineId,omitempty" xml:"defineId,omitempty"`
	Description      *string                                                                           `json:"description,omitempty" xml:"description,omitempty"`
	DisplayType      *string                                                                           `json:"displayType,omitempty" xml:"displayType,omitempty"`
	ExpressionTitle  *string                                                                           `json:"expressionTitle,omitempty" xml:"expressionTitle,omitempty"`
	FavoriteFlag     *bool                                                                             `json:"favoriteFlag,omitempty" xml:"favoriteFlag,omitempty"`
	FieldDetail      *string                                                                           `json:"fieldDetail,omitempty" xml:"fieldDetail,omitempty"`
	FieldRank        *int32                                                                            `json:"fieldRank,omitempty" xml:"fieldRank,omitempty"`
	FieldSource      *string                                                                           `json:"fieldSource,omitempty" xml:"fieldSource,omitempty"`
	FieldType        *string                                                                           `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	Id               *int64                                                                            `json:"id,omitempty" xml:"id,omitempty"`
	InputFieldType   *string                                                                           `json:"inputFieldType,omitempty" xml:"inputFieldType,omitempty"`
	InputRequired    *string                                                                           `json:"inputRequired,omitempty" xml:"inputRequired,omitempty"`
	Inputs           *string                                                                           `json:"inputs,omitempty" xml:"inputs,omitempty"`
	Name             *string                                                                           `json:"name,omitempty" xml:"name,omitempty"`
	Outlier          *string                                                                           `json:"outlier,omitempty" xml:"outlier,omitempty"`
	OutputThreshold  *DescribeEventVariableListResponseBodyResultObjectNativeVariablesOutputThreshold  `json:"outputThreshold,omitempty" xml:"outputThreshold,omitempty" type:"Struct"`
	ParentName       *string                                                                           `json:"parentName,omitempty" xml:"parentName,omitempty"`
	SourceType       *string                                                                           `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Title            *string                                                                           `json:"title,omitempty" xml:"title,omitempty"`
	Type             *string                                                                           `json:"type,omitempty" xml:"type,omitempty"`
	VariableVelocity *DescribeEventVariableListResponseBodyResultObjectNativeVariablesVariableVelocity `json:"variableVelocity,omitempty" xml:"variableVelocity,omitempty" type:"Struct"`
	XLabel           *string                                                                           `json:"xLabel,omitempty" xml:"xLabel,omitempty"`
	YLabel           *string                                                                           `json:"yLabel,omitempty" xml:"yLabel,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectNativeVariables) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectNativeVariables) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) GetCode() *string {
	return s.Code
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) GetDataDisplay() *string {
	return s.DataDisplay
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) GetDefineId() *string {
	return s.DefineId
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) GetDescription() *string {
	return s.Description
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) GetDisplayType() *string {
	return s.DisplayType
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) GetExpressionTitle() *string {
	return s.ExpressionTitle
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) GetFavoriteFlag() *bool {
	return s.FavoriteFlag
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) GetFieldDetail() *string {
	return s.FieldDetail
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) GetFieldRank() *int32 {
	return s.FieldRank
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) GetFieldSource() *string {
	return s.FieldSource
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) GetInputFieldType() *string {
	return s.InputFieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) GetInputRequired() *string {
	return s.InputRequired
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) GetInputs() *string {
	return s.Inputs
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) GetName() *string {
	return s.Name
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) GetOutlier() *string {
	return s.Outlier
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) GetOutputThreshold() *DescribeEventVariableListResponseBodyResultObjectNativeVariablesOutputThreshold {
	return s.OutputThreshold
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) GetParentName() *string {
	return s.ParentName
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) GetTitle() *string {
	return s.Title
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) GetType() *string {
	return s.Type
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) GetVariableVelocity() *DescribeEventVariableListResponseBodyResultObjectNativeVariablesVariableVelocity {
	return s.VariableVelocity
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) GetXLabel() *string {
	return s.XLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) GetYLabel() *string {
	return s.YLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) SetCode(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	s.Code = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) SetDataDisplay(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	s.DataDisplay = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) SetDefineId(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	s.DefineId = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) SetDescription(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	s.Description = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) SetDisplayType(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	s.DisplayType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) SetExpressionTitle(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	s.ExpressionTitle = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) SetFavoriteFlag(v bool) *DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	s.FavoriteFlag = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) SetFieldDetail(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	s.FieldDetail = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) SetFieldRank(v int32) *DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	s.FieldRank = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) SetFieldSource(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	s.FieldSource = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) SetFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	s.FieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) SetId(v int64) *DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	s.Id = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) SetInputFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	s.InputFieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) SetInputRequired(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	s.InputRequired = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) SetInputs(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	s.Inputs = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) SetName(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	s.Name = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) SetOutlier(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	s.Outlier = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) SetOutputThreshold(v *DescribeEventVariableListResponseBodyResultObjectNativeVariablesOutputThreshold) *DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	s.OutputThreshold = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) SetParentName(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	s.ParentName = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) SetSourceType(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	s.SourceType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) SetTitle(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	s.Title = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) SetType(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	s.Type = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) SetVariableVelocity(v *DescribeEventVariableListResponseBodyResultObjectNativeVariablesVariableVelocity) *DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	s.VariableVelocity = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) SetXLabel(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	s.XLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) SetYLabel(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariables {
	s.YLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariables) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectNativeVariablesOutputThreshold struct {
	MaxValue *float64 `json:"maxValue,omitempty" xml:"maxValue,omitempty"`
	MinValue *float64 `json:"minValue,omitempty" xml:"minValue,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectNativeVariablesOutputThreshold) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectNativeVariablesOutputThreshold) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariablesOutputThreshold) GetMaxValue() *float64 {
	return s.MaxValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariablesOutputThreshold) GetMinValue() *float64 {
	return s.MinValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariablesOutputThreshold) SetMaxValue(v float64) *DescribeEventVariableListResponseBodyResultObjectNativeVariablesOutputThreshold {
	s.MaxValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariablesOutputThreshold) SetMinValue(v float64) *DescribeEventVariableListResponseBodyResultObjectNativeVariablesOutputThreshold {
	s.MinValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariablesOutputThreshold) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectNativeVariablesVariableVelocity struct {
	Iv *string `json:"iv,omitempty" xml:"iv,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectNativeVariablesVariableVelocity) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectNativeVariablesVariableVelocity) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariablesVariableVelocity) GetIv() *string {
	return s.Iv
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariablesVariableVelocity) SetIv(v string) *DescribeEventVariableListResponseBodyResultObjectNativeVariablesVariableVelocity {
	s.Iv = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectNativeVariablesVariableVelocity) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectQueryVariables struct {
	Code             *string                                                                          `json:"code,omitempty" xml:"code,omitempty"`
	DataDisplay      *string                                                                          `json:"dataDisplay,omitempty" xml:"dataDisplay,omitempty"`
	DefineId         *string                                                                          `json:"defineId,omitempty" xml:"defineId,omitempty"`
	Description      *string                                                                          `json:"description,omitempty" xml:"description,omitempty"`
	DisplayType      *string                                                                          `json:"displayType,omitempty" xml:"displayType,omitempty"`
	ExpressionTitle  *string                                                                          `json:"expressionTitle,omitempty" xml:"expressionTitle,omitempty"`
	FavoriteFlag     *bool                                                                            `json:"favoriteFlag,omitempty" xml:"favoriteFlag,omitempty"`
	FieldDetail      *string                                                                          `json:"fieldDetail,omitempty" xml:"fieldDetail,omitempty"`
	FieldRank        *int32                                                                           `json:"fieldRank,omitempty" xml:"fieldRank,omitempty"`
	FieldSource      *string                                                                          `json:"fieldSource,omitempty" xml:"fieldSource,omitempty"`
	FieldType        *string                                                                          `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	Id               *int64                                                                           `json:"id,omitempty" xml:"id,omitempty"`
	InputFieldType   *string                                                                          `json:"inputFieldType,omitempty" xml:"inputFieldType,omitempty"`
	InputRequired    *string                                                                          `json:"inputRequired,omitempty" xml:"inputRequired,omitempty"`
	Inputs           *string                                                                          `json:"inputs,omitempty" xml:"inputs,omitempty"`
	Name             *string                                                                          `json:"name,omitempty" xml:"name,omitempty"`
	Outlier          *string                                                                          `json:"outlier,omitempty" xml:"outlier,omitempty"`
	OutputThreshold  *DescribeEventVariableListResponseBodyResultObjectQueryVariablesOutputThreshold  `json:"outputThreshold,omitempty" xml:"outputThreshold,omitempty" type:"Struct"`
	ParentName       *string                                                                          `json:"parentName,omitempty" xml:"parentName,omitempty"`
	SourceType       *string                                                                          `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Title            *string                                                                          `json:"title,omitempty" xml:"title,omitempty"`
	Type             *string                                                                          `json:"type,omitempty" xml:"type,omitempty"`
	VariableVelocity *DescribeEventVariableListResponseBodyResultObjectQueryVariablesVariableVelocity `json:"variableVelocity,omitempty" xml:"variableVelocity,omitempty" type:"Struct"`
	XLabel           *string                                                                          `json:"xLabel,omitempty" xml:"xLabel,omitempty"`
	YLabel           *string                                                                          `json:"yLabel,omitempty" xml:"yLabel,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectQueryVariables) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectQueryVariables) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) GetCode() *string {
	return s.Code
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) GetDataDisplay() *string {
	return s.DataDisplay
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) GetDefineId() *string {
	return s.DefineId
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) GetDescription() *string {
	return s.Description
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) GetDisplayType() *string {
	return s.DisplayType
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) GetExpressionTitle() *string {
	return s.ExpressionTitle
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) GetFavoriteFlag() *bool {
	return s.FavoriteFlag
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) GetFieldDetail() *string {
	return s.FieldDetail
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) GetFieldRank() *int32 {
	return s.FieldRank
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) GetFieldSource() *string {
	return s.FieldSource
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) GetInputFieldType() *string {
	return s.InputFieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) GetInputRequired() *string {
	return s.InputRequired
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) GetInputs() *string {
	return s.Inputs
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) GetName() *string {
	return s.Name
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) GetOutlier() *string {
	return s.Outlier
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) GetOutputThreshold() *DescribeEventVariableListResponseBodyResultObjectQueryVariablesOutputThreshold {
	return s.OutputThreshold
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) GetParentName() *string {
	return s.ParentName
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) GetTitle() *string {
	return s.Title
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) GetType() *string {
	return s.Type
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) GetVariableVelocity() *DescribeEventVariableListResponseBodyResultObjectQueryVariablesVariableVelocity {
	return s.VariableVelocity
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) GetXLabel() *string {
	return s.XLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) GetYLabel() *string {
	return s.YLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) SetCode(v string) *DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	s.Code = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) SetDataDisplay(v string) *DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	s.DataDisplay = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) SetDefineId(v string) *DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	s.DefineId = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) SetDescription(v string) *DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	s.Description = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) SetDisplayType(v string) *DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	s.DisplayType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) SetExpressionTitle(v string) *DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	s.ExpressionTitle = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) SetFavoriteFlag(v bool) *DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	s.FavoriteFlag = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) SetFieldDetail(v string) *DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	s.FieldDetail = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) SetFieldRank(v int32) *DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	s.FieldRank = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) SetFieldSource(v string) *DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	s.FieldSource = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) SetFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	s.FieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) SetId(v int64) *DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	s.Id = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) SetInputFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	s.InputFieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) SetInputRequired(v string) *DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	s.InputRequired = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) SetInputs(v string) *DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	s.Inputs = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) SetName(v string) *DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	s.Name = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) SetOutlier(v string) *DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	s.Outlier = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) SetOutputThreshold(v *DescribeEventVariableListResponseBodyResultObjectQueryVariablesOutputThreshold) *DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	s.OutputThreshold = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) SetParentName(v string) *DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	s.ParentName = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) SetSourceType(v string) *DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	s.SourceType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) SetTitle(v string) *DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	s.Title = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) SetType(v string) *DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	s.Type = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) SetVariableVelocity(v *DescribeEventVariableListResponseBodyResultObjectQueryVariablesVariableVelocity) *DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	s.VariableVelocity = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) SetXLabel(v string) *DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	s.XLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) SetYLabel(v string) *DescribeEventVariableListResponseBodyResultObjectQueryVariables {
	s.YLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariables) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectQueryVariablesOutputThreshold struct {
	MaxValue *float64 `json:"maxValue,omitempty" xml:"maxValue,omitempty"`
	MinValue *float64 `json:"minValue,omitempty" xml:"minValue,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectQueryVariablesOutputThreshold) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectQueryVariablesOutputThreshold) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariablesOutputThreshold) GetMaxValue() *float64 {
	return s.MaxValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariablesOutputThreshold) GetMinValue() *float64 {
	return s.MinValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariablesOutputThreshold) SetMaxValue(v float64) *DescribeEventVariableListResponseBodyResultObjectQueryVariablesOutputThreshold {
	s.MaxValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariablesOutputThreshold) SetMinValue(v float64) *DescribeEventVariableListResponseBodyResultObjectQueryVariablesOutputThreshold {
	s.MinValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariablesOutputThreshold) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectQueryVariablesVariableVelocity struct {
	Iv *string `json:"iv,omitempty" xml:"iv,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectQueryVariablesVariableVelocity) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectQueryVariablesVariableVelocity) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariablesVariableVelocity) GetIv() *string {
	return s.Iv
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariablesVariableVelocity) SetIv(v string) *DescribeEventVariableListResponseBodyResultObjectQueryVariablesVariableVelocity {
	s.Iv = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectQueryVariablesVariableVelocity) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectSelfVariables struct {
	Code             *string                                                                         `json:"code,omitempty" xml:"code,omitempty"`
	DataDisplay      *string                                                                         `json:"dataDisplay,omitempty" xml:"dataDisplay,omitempty"`
	DefineId         *string                                                                         `json:"defineId,omitempty" xml:"defineId,omitempty"`
	Description      *string                                                                         `json:"description,omitempty" xml:"description,omitempty"`
	DisplayType      *string                                                                         `json:"displayType,omitempty" xml:"displayType,omitempty"`
	ExpressionTitle  *string                                                                         `json:"expressionTitle,omitempty" xml:"expressionTitle,omitempty"`
	FavoriteFlag     *bool                                                                           `json:"favoriteFlag,omitempty" xml:"favoriteFlag,omitempty"`
	FieldDetail      *string                                                                         `json:"fieldDetail,omitempty" xml:"fieldDetail,omitempty"`
	FieldRank        *int32                                                                          `json:"fieldRank,omitempty" xml:"fieldRank,omitempty"`
	FieldSource      *string                                                                         `json:"fieldSource,omitempty" xml:"fieldSource,omitempty"`
	FieldType        *string                                                                         `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	Id               *int64                                                                          `json:"id,omitempty" xml:"id,omitempty"`
	InputFieldType   *string                                                                         `json:"inputFieldType,omitempty" xml:"inputFieldType,omitempty"`
	InputRequired    *string                                                                         `json:"inputRequired,omitempty" xml:"inputRequired,omitempty"`
	Inputs           *string                                                                         `json:"inputs,omitempty" xml:"inputs,omitempty"`
	Name             *string                                                                         `json:"name,omitempty" xml:"name,omitempty"`
	Outlier          *string                                                                         `json:"outlier,omitempty" xml:"outlier,omitempty"`
	OutputThreshold  *DescribeEventVariableListResponseBodyResultObjectSelfVariablesOutputThreshold  `json:"outputThreshold,omitempty" xml:"outputThreshold,omitempty" type:"Struct"`
	ParentName       *string                                                                         `json:"parentName,omitempty" xml:"parentName,omitempty"`
	SourceType       *string                                                                         `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Title            *string                                                                         `json:"title,omitempty" xml:"title,omitempty"`
	Type             *string                                                                         `json:"type,omitempty" xml:"type,omitempty"`
	VariableVelocity *DescribeEventVariableListResponseBodyResultObjectSelfVariablesVariableVelocity `json:"variableVelocity,omitempty" xml:"variableVelocity,omitempty" type:"Struct"`
	XLabel           *string                                                                         `json:"xLabel,omitempty" xml:"xLabel,omitempty"`
	YLabel           *string                                                                         `json:"yLabel,omitempty" xml:"yLabel,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectSelfVariables) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectSelfVariables) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) GetCode() *string {
	return s.Code
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) GetDataDisplay() *string {
	return s.DataDisplay
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) GetDefineId() *string {
	return s.DefineId
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) GetDescription() *string {
	return s.Description
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) GetDisplayType() *string {
	return s.DisplayType
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) GetExpressionTitle() *string {
	return s.ExpressionTitle
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) GetFavoriteFlag() *bool {
	return s.FavoriteFlag
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) GetFieldDetail() *string {
	return s.FieldDetail
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) GetFieldRank() *int32 {
	return s.FieldRank
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) GetFieldSource() *string {
	return s.FieldSource
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) GetInputFieldType() *string {
	return s.InputFieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) GetInputRequired() *string {
	return s.InputRequired
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) GetInputs() *string {
	return s.Inputs
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) GetName() *string {
	return s.Name
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) GetOutlier() *string {
	return s.Outlier
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) GetOutputThreshold() *DescribeEventVariableListResponseBodyResultObjectSelfVariablesOutputThreshold {
	return s.OutputThreshold
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) GetParentName() *string {
	return s.ParentName
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) GetTitle() *string {
	return s.Title
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) GetType() *string {
	return s.Type
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) GetVariableVelocity() *DescribeEventVariableListResponseBodyResultObjectSelfVariablesVariableVelocity {
	return s.VariableVelocity
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) GetXLabel() *string {
	return s.XLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) GetYLabel() *string {
	return s.YLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) SetCode(v string) *DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	s.Code = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) SetDataDisplay(v string) *DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	s.DataDisplay = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) SetDefineId(v string) *DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	s.DefineId = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) SetDescription(v string) *DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	s.Description = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) SetDisplayType(v string) *DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	s.DisplayType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) SetExpressionTitle(v string) *DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	s.ExpressionTitle = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) SetFavoriteFlag(v bool) *DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	s.FavoriteFlag = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) SetFieldDetail(v string) *DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	s.FieldDetail = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) SetFieldRank(v int32) *DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	s.FieldRank = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) SetFieldSource(v string) *DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	s.FieldSource = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) SetFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	s.FieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) SetId(v int64) *DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	s.Id = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) SetInputFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	s.InputFieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) SetInputRequired(v string) *DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	s.InputRequired = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) SetInputs(v string) *DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	s.Inputs = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) SetName(v string) *DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	s.Name = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) SetOutlier(v string) *DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	s.Outlier = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) SetOutputThreshold(v *DescribeEventVariableListResponseBodyResultObjectSelfVariablesOutputThreshold) *DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	s.OutputThreshold = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) SetParentName(v string) *DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	s.ParentName = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) SetSourceType(v string) *DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	s.SourceType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) SetTitle(v string) *DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	s.Title = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) SetType(v string) *DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	s.Type = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) SetVariableVelocity(v *DescribeEventVariableListResponseBodyResultObjectSelfVariablesVariableVelocity) *DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	s.VariableVelocity = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) SetXLabel(v string) *DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	s.XLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) SetYLabel(v string) *DescribeEventVariableListResponseBodyResultObjectSelfVariables {
	s.YLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariables) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectSelfVariablesOutputThreshold struct {
	MaxValue *float64 `json:"maxValue,omitempty" xml:"maxValue,omitempty"`
	MinValue *float64 `json:"minValue,omitempty" xml:"minValue,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectSelfVariablesOutputThreshold) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectSelfVariablesOutputThreshold) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariablesOutputThreshold) GetMaxValue() *float64 {
	return s.MaxValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariablesOutputThreshold) GetMinValue() *float64 {
	return s.MinValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariablesOutputThreshold) SetMaxValue(v float64) *DescribeEventVariableListResponseBodyResultObjectSelfVariablesOutputThreshold {
	s.MaxValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariablesOutputThreshold) SetMinValue(v float64) *DescribeEventVariableListResponseBodyResultObjectSelfVariablesOutputThreshold {
	s.MinValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariablesOutputThreshold) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectSelfVariablesVariableVelocity struct {
	Iv *string `json:"iv,omitempty" xml:"iv,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectSelfVariablesVariableVelocity) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectSelfVariablesVariableVelocity) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariablesVariableVelocity) GetIv() *string {
	return s.Iv
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariablesVariableVelocity) SetIv(v string) *DescribeEventVariableListResponseBodyResultObjectSelfVariablesVariableVelocity {
	s.Iv = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSelfVariablesVariableVelocity) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectSysVariables struct {
	Code             *string                                                                        `json:"code,omitempty" xml:"code,omitempty"`
	DataDisplay      *string                                                                        `json:"dataDisplay,omitempty" xml:"dataDisplay,omitempty"`
	DefineId         *string                                                                        `json:"defineId,omitempty" xml:"defineId,omitempty"`
	Description      *string                                                                        `json:"description,omitempty" xml:"description,omitempty"`
	DisplayType      *string                                                                        `json:"displayType,omitempty" xml:"displayType,omitempty"`
	ExpressionTitle  *string                                                                        `json:"expressionTitle,omitempty" xml:"expressionTitle,omitempty"`
	FavoriteFlag     *bool                                                                          `json:"favoriteFlag,omitempty" xml:"favoriteFlag,omitempty"`
	FieldDetail      *string                                                                        `json:"fieldDetail,omitempty" xml:"fieldDetail,omitempty"`
	FieldRank        *int32                                                                         `json:"fieldRank,omitempty" xml:"fieldRank,omitempty"`
	FieldSource      *string                                                                        `json:"fieldSource,omitempty" xml:"fieldSource,omitempty"`
	FieldType        *string                                                                        `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	Id               *int64                                                                         `json:"id,omitempty" xml:"id,omitempty"`
	InputFieldType   *string                                                                        `json:"inputFieldType,omitempty" xml:"inputFieldType,omitempty"`
	InputRequired    *string                                                                        `json:"inputRequired,omitempty" xml:"inputRequired,omitempty"`
	Inputs           *string                                                                        `json:"inputs,omitempty" xml:"inputs,omitempty"`
	Name             *string                                                                        `json:"name,omitempty" xml:"name,omitempty"`
	Outlier          *string                                                                        `json:"outlier,omitempty" xml:"outlier,omitempty"`
	OutputThreshold  *DescribeEventVariableListResponseBodyResultObjectSysVariablesOutputThreshold  `json:"outputThreshold,omitempty" xml:"outputThreshold,omitempty" type:"Struct"`
	ParentName       *string                                                                        `json:"parentName,omitempty" xml:"parentName,omitempty"`
	SourceType       *string                                                                        `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Title            *string                                                                        `json:"title,omitempty" xml:"title,omitempty"`
	Type             *string                                                                        `json:"type,omitempty" xml:"type,omitempty"`
	VariableVelocity *DescribeEventVariableListResponseBodyResultObjectSysVariablesVariableVelocity `json:"variableVelocity,omitempty" xml:"variableVelocity,omitempty" type:"Struct"`
	XLabel           *string                                                                        `json:"xLabel,omitempty" xml:"xLabel,omitempty"`
	YLabel           *string                                                                        `json:"yLabel,omitempty" xml:"yLabel,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectSysVariables) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectSysVariables) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) GetCode() *string {
	return s.Code
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) GetDataDisplay() *string {
	return s.DataDisplay
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) GetDefineId() *string {
	return s.DefineId
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) GetDescription() *string {
	return s.Description
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) GetDisplayType() *string {
	return s.DisplayType
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) GetExpressionTitle() *string {
	return s.ExpressionTitle
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) GetFavoriteFlag() *bool {
	return s.FavoriteFlag
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) GetFieldDetail() *string {
	return s.FieldDetail
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) GetFieldRank() *int32 {
	return s.FieldRank
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) GetFieldSource() *string {
	return s.FieldSource
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) GetInputFieldType() *string {
	return s.InputFieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) GetInputRequired() *string {
	return s.InputRequired
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) GetInputs() *string {
	return s.Inputs
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) GetName() *string {
	return s.Name
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) GetOutlier() *string {
	return s.Outlier
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) GetOutputThreshold() *DescribeEventVariableListResponseBodyResultObjectSysVariablesOutputThreshold {
	return s.OutputThreshold
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) GetParentName() *string {
	return s.ParentName
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) GetTitle() *string {
	return s.Title
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) GetType() *string {
	return s.Type
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) GetVariableVelocity() *DescribeEventVariableListResponseBodyResultObjectSysVariablesVariableVelocity {
	return s.VariableVelocity
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) GetXLabel() *string {
	return s.XLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) GetYLabel() *string {
	return s.YLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) SetCode(v string) *DescribeEventVariableListResponseBodyResultObjectSysVariables {
	s.Code = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) SetDataDisplay(v string) *DescribeEventVariableListResponseBodyResultObjectSysVariables {
	s.DataDisplay = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) SetDefineId(v string) *DescribeEventVariableListResponseBodyResultObjectSysVariables {
	s.DefineId = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) SetDescription(v string) *DescribeEventVariableListResponseBodyResultObjectSysVariables {
	s.Description = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) SetDisplayType(v string) *DescribeEventVariableListResponseBodyResultObjectSysVariables {
	s.DisplayType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) SetExpressionTitle(v string) *DescribeEventVariableListResponseBodyResultObjectSysVariables {
	s.ExpressionTitle = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) SetFavoriteFlag(v bool) *DescribeEventVariableListResponseBodyResultObjectSysVariables {
	s.FavoriteFlag = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) SetFieldDetail(v string) *DescribeEventVariableListResponseBodyResultObjectSysVariables {
	s.FieldDetail = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) SetFieldRank(v int32) *DescribeEventVariableListResponseBodyResultObjectSysVariables {
	s.FieldRank = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) SetFieldSource(v string) *DescribeEventVariableListResponseBodyResultObjectSysVariables {
	s.FieldSource = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) SetFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectSysVariables {
	s.FieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) SetId(v int64) *DescribeEventVariableListResponseBodyResultObjectSysVariables {
	s.Id = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) SetInputFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectSysVariables {
	s.InputFieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) SetInputRequired(v string) *DescribeEventVariableListResponseBodyResultObjectSysVariables {
	s.InputRequired = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) SetInputs(v string) *DescribeEventVariableListResponseBodyResultObjectSysVariables {
	s.Inputs = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) SetName(v string) *DescribeEventVariableListResponseBodyResultObjectSysVariables {
	s.Name = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) SetOutlier(v string) *DescribeEventVariableListResponseBodyResultObjectSysVariables {
	s.Outlier = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) SetOutputThreshold(v *DescribeEventVariableListResponseBodyResultObjectSysVariablesOutputThreshold) *DescribeEventVariableListResponseBodyResultObjectSysVariables {
	s.OutputThreshold = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) SetParentName(v string) *DescribeEventVariableListResponseBodyResultObjectSysVariables {
	s.ParentName = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) SetSourceType(v string) *DescribeEventVariableListResponseBodyResultObjectSysVariables {
	s.SourceType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) SetTitle(v string) *DescribeEventVariableListResponseBodyResultObjectSysVariables {
	s.Title = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) SetType(v string) *DescribeEventVariableListResponseBodyResultObjectSysVariables {
	s.Type = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) SetVariableVelocity(v *DescribeEventVariableListResponseBodyResultObjectSysVariablesVariableVelocity) *DescribeEventVariableListResponseBodyResultObjectSysVariables {
	s.VariableVelocity = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) SetXLabel(v string) *DescribeEventVariableListResponseBodyResultObjectSysVariables {
	s.XLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) SetYLabel(v string) *DescribeEventVariableListResponseBodyResultObjectSysVariables {
	s.YLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariables) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectSysVariablesOutputThreshold struct {
	MaxValue *float64 `json:"maxValue,omitempty" xml:"maxValue,omitempty"`
	MinValue *float64 `json:"minValue,omitempty" xml:"minValue,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectSysVariablesOutputThreshold) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectSysVariablesOutputThreshold) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariablesOutputThreshold) GetMaxValue() *float64 {
	return s.MaxValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariablesOutputThreshold) GetMinValue() *float64 {
	return s.MinValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariablesOutputThreshold) SetMaxValue(v float64) *DescribeEventVariableListResponseBodyResultObjectSysVariablesOutputThreshold {
	s.MaxValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariablesOutputThreshold) SetMinValue(v float64) *DescribeEventVariableListResponseBodyResultObjectSysVariablesOutputThreshold {
	s.MinValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariablesOutputThreshold) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectSysVariablesVariableVelocity struct {
	Iv *string `json:"iv,omitempty" xml:"iv,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectSysVariablesVariableVelocity) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectSysVariablesVariableVelocity) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariablesVariableVelocity) GetIv() *string {
	return s.Iv
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariablesVariableVelocity) SetIv(v string) *DescribeEventVariableListResponseBodyResultObjectSysVariablesVariableVelocity {
	s.Iv = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectSysVariablesVariableVelocity) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectVelocityVariables struct {
	Code             *string                                                                             `json:"code,omitempty" xml:"code,omitempty"`
	DataDisplay      *string                                                                             `json:"dataDisplay,omitempty" xml:"dataDisplay,omitempty"`
	DefineId         *string                                                                             `json:"defineId,omitempty" xml:"defineId,omitempty"`
	Description      *string                                                                             `json:"description,omitempty" xml:"description,omitempty"`
	DisplayType      *string                                                                             `json:"displayType,omitempty" xml:"displayType,omitempty"`
	ExpressionTitle  *string                                                                             `json:"expressionTitle,omitempty" xml:"expressionTitle,omitempty"`
	FavoriteFlag     *bool                                                                               `json:"favoriteFlag,omitempty" xml:"favoriteFlag,omitempty"`
	FieldDetail      *string                                                                             `json:"fieldDetail,omitempty" xml:"fieldDetail,omitempty"`
	FieldRank        *int32                                                                              `json:"fieldRank,omitempty" xml:"fieldRank,omitempty"`
	FieldSource      *string                                                                             `json:"fieldSource,omitempty" xml:"fieldSource,omitempty"`
	FieldType        *string                                                                             `json:"fieldType,omitempty" xml:"fieldType,omitempty"`
	Id               *int64                                                                              `json:"id,omitempty" xml:"id,omitempty"`
	InputFieldType   *string                                                                             `json:"inputFieldType,omitempty" xml:"inputFieldType,omitempty"`
	InputRequired    *string                                                                             `json:"inputRequired,omitempty" xml:"inputRequired,omitempty"`
	Inputs           *string                                                                             `json:"inputs,omitempty" xml:"inputs,omitempty"`
	Name             *string                                                                             `json:"name,omitempty" xml:"name,omitempty"`
	Outlier          *string                                                                             `json:"outlier,omitempty" xml:"outlier,omitempty"`
	OutputThreshold  *DescribeEventVariableListResponseBodyResultObjectVelocityVariablesOutputThreshold  `json:"outputThreshold,omitempty" xml:"outputThreshold,omitempty" type:"Struct"`
	ParentName       *string                                                                             `json:"parentName,omitempty" xml:"parentName,omitempty"`
	SourceType       *string                                                                             `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Title            *string                                                                             `json:"title,omitempty" xml:"title,omitempty"`
	Type             *string                                                                             `json:"type,omitempty" xml:"type,omitempty"`
	VariableVelocity *DescribeEventVariableListResponseBodyResultObjectVelocityVariablesVariableVelocity `json:"variableVelocity,omitempty" xml:"variableVelocity,omitempty" type:"Struct"`
	XLabel           *string                                                                             `json:"xLabel,omitempty" xml:"xLabel,omitempty"`
	YLabel           *string                                                                             `json:"yLabel,omitempty" xml:"yLabel,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectVelocityVariables) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GetCode() *string {
	return s.Code
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GetDataDisplay() *string {
	return s.DataDisplay
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GetDefineId() *string {
	return s.DefineId
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GetDescription() *string {
	return s.Description
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GetDisplayType() *string {
	return s.DisplayType
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GetExpressionTitle() *string {
	return s.ExpressionTitle
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GetFavoriteFlag() *bool {
	return s.FavoriteFlag
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GetFieldDetail() *string {
	return s.FieldDetail
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GetFieldRank() *int32 {
	return s.FieldRank
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GetFieldSource() *string {
	return s.FieldSource
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GetFieldType() *string {
	return s.FieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GetInputFieldType() *string {
	return s.InputFieldType
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GetInputRequired() *string {
	return s.InputRequired
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GetInputs() *string {
	return s.Inputs
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GetName() *string {
	return s.Name
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GetOutlier() *string {
	return s.Outlier
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GetOutputThreshold() *DescribeEventVariableListResponseBodyResultObjectVelocityVariablesOutputThreshold {
	return s.OutputThreshold
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GetParentName() *string {
	return s.ParentName
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GetTitle() *string {
	return s.Title
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GetType() *string {
	return s.Type
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GetVariableVelocity() *DescribeEventVariableListResponseBodyResultObjectVelocityVariablesVariableVelocity {
	return s.VariableVelocity
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GetXLabel() *string {
	return s.XLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) GetYLabel() *string {
	return s.YLabel
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) SetCode(v string) *DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	s.Code = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) SetDataDisplay(v string) *DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	s.DataDisplay = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) SetDefineId(v string) *DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	s.DefineId = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) SetDescription(v string) *DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	s.Description = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) SetDisplayType(v string) *DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	s.DisplayType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) SetExpressionTitle(v string) *DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	s.ExpressionTitle = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) SetFavoriteFlag(v bool) *DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	s.FavoriteFlag = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) SetFieldDetail(v string) *DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	s.FieldDetail = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) SetFieldRank(v int32) *DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	s.FieldRank = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) SetFieldSource(v string) *DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	s.FieldSource = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) SetFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	s.FieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) SetId(v int64) *DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	s.Id = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) SetInputFieldType(v string) *DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	s.InputFieldType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) SetInputRequired(v string) *DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	s.InputRequired = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) SetInputs(v string) *DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	s.Inputs = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) SetName(v string) *DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	s.Name = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) SetOutlier(v string) *DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	s.Outlier = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) SetOutputThreshold(v *DescribeEventVariableListResponseBodyResultObjectVelocityVariablesOutputThreshold) *DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	s.OutputThreshold = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) SetParentName(v string) *DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	s.ParentName = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) SetSourceType(v string) *DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	s.SourceType = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) SetTitle(v string) *DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	s.Title = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) SetType(v string) *DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	s.Type = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) SetVariableVelocity(v *DescribeEventVariableListResponseBodyResultObjectVelocityVariablesVariableVelocity) *DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	s.VariableVelocity = v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) SetXLabel(v string) *DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	s.XLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) SetYLabel(v string) *DescribeEventVariableListResponseBodyResultObjectVelocityVariables {
	s.YLabel = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariables) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectVelocityVariablesOutputThreshold struct {
	MaxValue *float64 `json:"maxValue,omitempty" xml:"maxValue,omitempty"`
	MinValue *float64 `json:"minValue,omitempty" xml:"minValue,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectVelocityVariablesOutputThreshold) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectVelocityVariablesOutputThreshold) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariablesOutputThreshold) GetMaxValue() *float64 {
	return s.MaxValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariablesOutputThreshold) GetMinValue() *float64 {
	return s.MinValue
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariablesOutputThreshold) SetMaxValue(v float64) *DescribeEventVariableListResponseBodyResultObjectVelocityVariablesOutputThreshold {
	s.MaxValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariablesOutputThreshold) SetMinValue(v float64) *DescribeEventVariableListResponseBodyResultObjectVelocityVariablesOutputThreshold {
	s.MinValue = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariablesOutputThreshold) Validate() error {
	return dara.Validate(s)
}

type DescribeEventVariableListResponseBodyResultObjectVelocityVariablesVariableVelocity struct {
	Iv *string `json:"iv,omitempty" xml:"iv,omitempty"`
}

func (s DescribeEventVariableListResponseBodyResultObjectVelocityVariablesVariableVelocity) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponseBodyResultObjectVelocityVariablesVariableVelocity) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariablesVariableVelocity) GetIv() *string {
	return s.Iv
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariablesVariableVelocity) SetIv(v string) *DescribeEventVariableListResponseBodyResultObjectVelocityVariablesVariableVelocity {
	s.Iv = &v
	return s
}

func (s *DescribeEventVariableListResponseBodyResultObjectVelocityVariablesVariableVelocity) Validate() error {
	return dara.Validate(s)
}
