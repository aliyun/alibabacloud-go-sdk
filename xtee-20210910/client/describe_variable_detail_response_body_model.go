// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVariableDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVariableDetailResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeVariableDetailResponseBodyResultObject) *DescribeVariableDetailResponseBody
	GetResultObject() *DescribeVariableDetailResponseBodyResultObject
}

type DescribeVariableDetailResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject *DescribeVariableDetailResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeVariableDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVariableDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVariableDetailResponseBody) GetResultObject() *DescribeVariableDetailResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeVariableDetailResponseBody) SetRequestId(v string) *DescribeVariableDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVariableDetailResponseBody) SetResultObject(v *DescribeVariableDetailResponseBodyResultObject) *DescribeVariableDetailResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeVariableDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVariableDetailResponseBodyResultObject struct {
	// Basic attributes.
	BaseInfo *DescribeVariableDetailResponseBodyResultObjectBaseInfo `json:"baseInfo,omitempty" xml:"baseInfo,omitempty" type:"Struct"`
}

func (s DescribeVariableDetailResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableDetailResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeVariableDetailResponseBodyResultObject) GetBaseInfo() *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	return s.BaseInfo
}

func (s *DescribeVariableDetailResponseBodyResultObject) SetBaseInfo(v *DescribeVariableDetailResponseBodyResultObjectBaseInfo) *DescribeVariableDetailResponseBodyResultObject {
	s.BaseInfo = v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}

type DescribeVariableDetailResponseBodyResultObjectBaseInfo struct {
	// Whether variable binding is allowed
	//
	// example:
	//
	// ENABLE
	AllowBind *string `json:"allowBind,omitempty" xml:"allowBind,omitempty"`
	// Charging mode
	//
	// example:
	//
	// FREE
	ChargingMode *string `json:"chargingMode,omitempty" xml:"chargingMode,omitempty"`
	// Charging mode description
	//
	// example:
	//
	// 免费
	ChargingModeDesc *string `json:"chargingModeDesc,omitempty" xml:"chargingModeDesc,omitempty"`
	// Creator.
	//
	// example:
	//
	// 176020
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// Data distribution display, in JSON format
	//
	// example:
	//
	// {}
	DataDisplay *string `json:"dataDisplay,omitempty" xml:"dataDisplay,omitempty"`
	// Data valid range, left-closed and right-closed
	//
	// example:
	//
	// (0,10)
	DataThreshold *string `json:"dataThreshold,omitempty" xml:"dataThreshold,omitempty"`
	// Deduction factor
	//
	// example:
	//
	// 10
	DeductionFactor *int32 `json:"deductionFactor,omitempty" xml:"deductionFactor,omitempty"`
	// Description.
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Front-end binding allowed
	//
	// example:
	//
	// ENABLE
	FrontAllowBind *string `json:"frontAllowBind,omitempty" xml:"frontAllowBind,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1698143758000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// Modification time.
	//
	// example:
	//
	// 1698143758000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Primary key ID
	//
	// example:
	//
	// 3144
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Required parameters
	//
	//
	//
	//      When inputRequired=__all__, it means all parameters are required
	//
	//      When inputRequired=__one__, it means only one input is needed
	//
	//      Required fields are separated by commas, e.g., mobile,ip,email
	//
	// example:
	//
	// __one__
	InputRequired *string `json:"inputRequired,omitempty" xml:"inputRequired,omitempty"`
	// Input parameters.
	//
	// example:
	//
	// ip,age,mobile
	Inputs *string `json:"inputs,omitempty" xml:"inputs,omitempty"`
	// Input parameter description.
	//
	// example:
	//
	// ip,年龄,手机号
	InputsDesc *string `json:"inputsDesc,omitempty" xml:"inputsDesc,omitempty"`
	// Invoke key
	//
	// example:
	//
	// onlineScamDetectionTags_v
	InvokeKey *string `json:"invokeKey,omitempty" xml:"invokeKey,omitempty"`
	// Invoke RT, unit: milliseconds
	//
	// example:
	//
	// 10
	InvokeRt *int32 `json:"invokeRt,omitempty" xml:"invokeRt,omitempty"`
	// Invocation success rate
	//
	// example:
	//
	// 100
	InvokeSuccessRate *string `json:"invokeSuccessRate,omitempty" xml:"invokeSuccessRate,omitempty"`
	// Number of invocations
	//
	// example:
	//
	// 100000
	InvokeTimes *int64 `json:"invokeTimes,omitempty" xml:"invokeTimes,omitempty"`
	// Last modifier.
	//
	// example:
	//
	// root
	LastModifiedOperator *string `json:"lastModifiedOperator,omitempty" xml:"lastModifiedOperator,omitempty"`
	// Variable name
	//
	// example:
	//
	// __onlineScamDetectionTags__
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Output
	//
	// example:
	//
	// STRING
	Outputs *string `json:"outputs,omitempty" xml:"outputs,omitempty"`
	// Output description
	//
	// example:
	//
	// 字符串
	OutputsDesc *string `json:"outputsDesc,omitempty" xml:"outputsDesc,omitempty"`
	// Code of applicable scenarios
	Scene []*string `json:"scene,omitempty" xml:"scene,omitempty" type:"Repeated"`
	// Applicable scenario description
	SceneDesc []*string `json:"sceneDesc,omitempty" xml:"sceneDesc,omitempty" type:"Repeated"`
	// Display order
	//
	// example:
	//
	// 10
	ShowOrder *string `json:"showOrder,omitempty" xml:"showOrder,omitempty"`
	// Source
	//
	// example:
	//
	// SAF
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// Source description
	//
	// example:
	//
	// 风险识别
	SourceDesc *string `json:"sourceDesc,omitempty" xml:"sourceDesc,omitempty"`
	// Status.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// List of supported regions.
	SupportRegions []*string `json:"supportRegions,omitempty" xml:"supportRegions,omitempty" type:"Repeated"`
	// Title.
	//
	// example:
	//
	// 诈骗引流识别_标签
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// Type
	//
	// example:
	//
	// NATIVE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Category description
	//
	// example:
	//
	// 事件字段
	TypeDesc *string `json:"typeDesc,omitempty" xml:"typeDesc,omitempty"`
	// X-axis label for data distribution display
	//
	// example:
	//
	// 10
	XLabel *string `json:"xLabel,omitempty" xml:"xLabel,omitempty"`
	// Data distribution display y-axis label
	//
	// example:
	//
	// 10
	YLabel *string `json:"yLabel,omitempty" xml:"yLabel,omitempty"`
}

func (s DescribeVariableDetailResponseBodyResultObjectBaseInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableDetailResponseBodyResultObjectBaseInfo) GoString() string {
	return s.String()
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetAllowBind() *string {
	return s.AllowBind
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetChargingMode() *string {
	return s.ChargingMode
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetChargingModeDesc() *string {
	return s.ChargingModeDesc
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetCreator() *string {
	return s.Creator
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetDataDisplay() *string {
	return s.DataDisplay
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetDataThreshold() *string {
	return s.DataThreshold
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetDeductionFactor() *int32 {
	return s.DeductionFactor
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetDescription() *string {
	return s.Description
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetFrontAllowBind() *string {
	return s.FrontAllowBind
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetId() *int64 {
	return s.Id
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetInputRequired() *string {
	return s.InputRequired
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetInputs() *string {
	return s.Inputs
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetInputsDesc() *string {
	return s.InputsDesc
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetInvokeKey() *string {
	return s.InvokeKey
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetInvokeRt() *int32 {
	return s.InvokeRt
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetInvokeSuccessRate() *string {
	return s.InvokeSuccessRate
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetInvokeTimes() *int64 {
	return s.InvokeTimes
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetLastModifiedOperator() *string {
	return s.LastModifiedOperator
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetName() *string {
	return s.Name
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetOutputs() *string {
	return s.Outputs
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetOutputsDesc() *string {
	return s.OutputsDesc
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetScene() []*string {
	return s.Scene
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetSceneDesc() []*string {
	return s.SceneDesc
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetShowOrder() *string {
	return s.ShowOrder
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetSource() *string {
	return s.Source
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetSourceDesc() *string {
	return s.SourceDesc
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetSupportRegions() []*string {
	return s.SupportRegions
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetTitle() *string {
	return s.Title
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetType() *string {
	return s.Type
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetTypeDesc() *string {
	return s.TypeDesc
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetXLabel() *string {
	return s.XLabel
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) GetYLabel() *string {
	return s.YLabel
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetAllowBind(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.AllowBind = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetChargingMode(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.ChargingMode = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetChargingModeDesc(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.ChargingModeDesc = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetCreator(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.Creator = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetDataDisplay(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.DataDisplay = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetDataThreshold(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.DataThreshold = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetDeductionFactor(v int32) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.DeductionFactor = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetDescription(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.Description = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetFrontAllowBind(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.FrontAllowBind = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetGmtCreate(v int64) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.GmtCreate = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetGmtModified(v int64) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.GmtModified = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetId(v int64) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.Id = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetInputRequired(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.InputRequired = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetInputs(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.Inputs = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetInputsDesc(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.InputsDesc = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetInvokeKey(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.InvokeKey = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetInvokeRt(v int32) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.InvokeRt = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetInvokeSuccessRate(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.InvokeSuccessRate = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetInvokeTimes(v int64) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.InvokeTimes = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetLastModifiedOperator(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.LastModifiedOperator = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetName(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.Name = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetOutputs(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.Outputs = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetOutputsDesc(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.OutputsDesc = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetScene(v []*string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.Scene = v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetSceneDesc(v []*string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.SceneDesc = v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetShowOrder(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.ShowOrder = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetSource(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.Source = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetSourceDesc(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.SourceDesc = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetStatus(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.Status = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetSupportRegions(v []*string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.SupportRegions = v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetTitle(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.Title = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetType(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.Type = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetTypeDesc(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.TypeDesc = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetXLabel(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.XLabel = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) SetYLabel(v string) *DescribeVariableDetailResponseBodyResultObjectBaseInfo {
	s.YLabel = &v
	return s
}

func (s *DescribeVariableDetailResponseBodyResultObjectBaseInfo) Validate() error {
	return dara.Validate(s)
}
