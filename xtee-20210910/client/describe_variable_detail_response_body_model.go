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
	// The request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned object.
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
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVariableDetailResponseBodyResultObject struct {
	// The basic properties.
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
	if s.BaseInfo != nil {
		if err := s.BaseInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVariableDetailResponseBodyResultObjectBaseInfo struct {
	// Specifies whether variable binding is allowed. Valid values:
	//
	// - **DISABLE**: unavailable
	//
	// - **ALL**: all
	//
	// - **ENABLE**: available
	//
	// - **PART_ENABLE**: partially available.
	//
	// example:
	//
	// ENABLE
	AllowBind *string `json:"allowBind,omitempty" xml:"allowBind,omitempty"`
	// The billing mode. Valid values:
	//
	// - **PAY_PER_VIEW**: paid
	//
	// - **FREE**: free.
	//
	// example:
	//
	// FREE
	ChargingMode *string `json:"chargingMode,omitempty" xml:"chargingMode,omitempty"`
	// The billing mode description.
	//
	// example:
	//
	// 免费
	ChargingModeDesc *string `json:"chargingModeDesc,omitempty" xml:"chargingModeDesc,omitempty"`
	// The creator.
	//
	// example:
	//
	// 176020
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The data distribution display in JSON format.
	//
	// example:
	//
	// {}
	DataDisplay *string `json:"dataDisplay,omitempty" xml:"dataDisplay,omitempty"`
	// The valid data range, inclusive on both ends.
	//
	// example:
	//
	// (0,10)
	DataThreshold *string `json:"dataThreshold,omitempty" xml:"dataThreshold,omitempty"`
	// The deduction coefficient.
	//
	// example:
	//
	// 10
	DeductionFactor *int32 `json:"deductionFactor,omitempty" xml:"deductionFactor,omitempty"`
	// The description.
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether front-end binding is allowed. Valid values:
	//
	// - **DISABLE**: not allowed
	//
	// - **ENABLE**: allowed.
	//
	// example:
	//
	// ENABLE
	FrontAllowBind *string `json:"frontAllowBind,omitempty" xml:"frontAllowBind,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1698143758000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 1698143758000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The primary key ID.
	//
	// example:
	//
	// 3144
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The required parameters.
	//
	//
	//
	// When inputRequired is set to __all__, all parameters are required.
	//
	// When inputRequired is set to __one__, only one input parameter is required.
	//
	// Required fields are separated by commas, such as mobile,ip,email.
	//
	// example:
	//
	// __one__
	InputRequired *string `json:"inputRequired,omitempty" xml:"inputRequired,omitempty"`
	// The input parameters.
	//
	// example:
	//
	// ip,age,mobile
	Inputs *string `json:"inputs,omitempty" xml:"inputs,omitempty"`
	// The input parameter description.
	//
	// example:
	//
	// ip,年龄,手机号
	InputsDesc *string `json:"inputsDesc,omitempty" xml:"inputsDesc,omitempty"`
	// The invocation key.
	//
	// example:
	//
	// onlineScamDetectionTags_v
	InvokeKey *string `json:"invokeKey,omitempty" xml:"invokeKey,omitempty"`
	// The invocation response time, in milliseconds.
	//
	// example:
	//
	// 10
	InvokeRt *int32 `json:"invokeRt,omitempty" xml:"invokeRt,omitempty"`
	// The invocation success rate.
	//
	// example:
	//
	// 100
	InvokeSuccessRate *string `json:"invokeSuccessRate,omitempty" xml:"invokeSuccessRate,omitempty"`
	// The number of invocations.
	//
	// example:
	//
	// 100000
	InvokeTimes *int64 `json:"invokeTimes,omitempty" xml:"invokeTimes,omitempty"`
	// The last modifier.
	//
	// example:
	//
	// root
	LastModifiedOperator *string `json:"lastModifiedOperator,omitempty" xml:"lastModifiedOperator,omitempty"`
	// The variable name.
	//
	// example:
	//
	// __onlineScamDetectionTags__
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The outputs.
	//
	// example:
	//
	// STRING
	Outputs *string `json:"outputs,omitempty" xml:"outputs,omitempty"`
	// The output description.
	//
	// example:
	//
	// 字符串
	OutputsDesc *string `json:"outputsDesc,omitempty" xml:"outputsDesc,omitempty"`
	// The applicable scenario code.
	Scene []*string `json:"scene,omitempty" xml:"scene,omitempty" type:"Repeated"`
	// The applicable scenario description.
	SceneDesc []*string `json:"sceneDesc,omitempty" xml:"sceneDesc,omitempty" type:"Repeated"`
	// The display order.
	//
	// example:
	//
	// 10
	ShowOrder *string `json:"showOrder,omitempty" xml:"showOrder,omitempty"`
	// The source.
	//
	// example:
	//
	// SAF
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The source description.
	//
	// example:
	//
	// 风险识别
	SourceDesc *string `json:"sourceDesc,omitempty" xml:"sourceDesc,omitempty"`
	// The status.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The list of supported regions.
	SupportRegions []*string `json:"supportRegions,omitempty" xml:"supportRegions,omitempty" type:"Repeated"`
	// The title.
	//
	// example:
	//
	// 诈骗引流识别_标签
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The type.
	//
	// example:
	//
	// NATIVE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The category description.
	//
	// example:
	//
	// 事件字段
	TypeDesc *string `json:"typeDesc,omitempty" xml:"typeDesc,omitempty"`
	// The X-axis label for the data distribution chart.
	//
	// example:
	//
	// 10
	XLabel *string `json:"xLabel,omitempty" xml:"xLabel,omitempty"`
	// The Y-axis label for the data distribution chart.
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
