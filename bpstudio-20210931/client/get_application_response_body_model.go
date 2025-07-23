// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetApplicationResponseBody
	GetCode() *string
	SetData(v *GetApplicationResponseBodyData) *GetApplicationResponseBody
	GetData() *GetApplicationResponseBodyData
	SetMessage(v string) *GetApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetApplicationResponseBody
	GetRequestId() *string
}

type GetApplicationResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the application.
	Data *GetApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Reason for the request failure
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetApplicationResponseBody) GetData() *GetApplicationResponseBodyData {
	return s.Data
}

func (s *GetApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationResponseBody) SetCode(v string) *GetApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *GetApplicationResponseBody) SetData(v *GetApplicationResponseBodyData) *GetApplicationResponseBody {
	s.Data = v
	return s
}

func (s *GetApplicationResponseBody) SetMessage(v string) *GetApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *GetApplicationResponseBody) SetRequestId(v string) *GetApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyData struct {
	// App ID
	//
	// example:
	//
	// VVK605ZH00OA4MRT
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The resource tag.
	Checklist      []*GetApplicationResponseBodyDataChecklist      `json:"Checklist,omitempty" xml:"Checklist,omitempty" type:"Repeated"`
	ComplianceList []*GetApplicationResponseBodyDataComplianceList `json:"ComplianceList,omitempty" xml:"ComplianceList,omitempty" type:"Repeated"`
	// The time when the app was created
	//
	// example:
	//
	// 2021-08-09 14:37:16
	CreateTime    *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeployPercent *float64 `json:"DeployPercent,omitempty" xml:"DeployPercent,omitempty"`
	// Application description
	//
	// example:
	//
	// remark
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The resource type.
	//
	// example:
	//
	// Success
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The URL of the image in the database.
	//
	// example:
	//
	// The details of the application.
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// App name
	//
	// example:
	//
	// 1411182597819805/sr-8DWU4RUS49NIDII0.png
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The billing results.
	PriceList []*GetApplicationResponseBodyDataPriceList `json:"PriceList,omitempty" xml:"PriceList,omitempty" type:"Repeated"`
	// The ID of the resource group to which the app belongs
	//
	// example:
	//
	// rg-aekzhfgmw4e6fwq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource specification.
	ResourceList []*GetApplicationResponseBodyDataResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	// Verification passed
	//
	// example:
	//
	// Deployed_Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the template associated with the application
	//
	// example:
	//
	// FYS9VZ535U20V7HT
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetApplicationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyData) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetApplicationResponseBodyData) GetChecklist() []*GetApplicationResponseBodyDataChecklist {
	return s.Checklist
}

func (s *GetApplicationResponseBodyData) GetComplianceList() []*GetApplicationResponseBodyDataComplianceList {
	return s.ComplianceList
}

func (s *GetApplicationResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetApplicationResponseBodyData) GetDeployPercent() *float64 {
	return s.DeployPercent
}

func (s *GetApplicationResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetApplicationResponseBodyData) GetError() *string {
	return s.Error
}

func (s *GetApplicationResponseBodyData) GetImageURL() *string {
	return s.ImageURL
}

func (s *GetApplicationResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetApplicationResponseBodyData) GetPriceList() []*GetApplicationResponseBodyDataPriceList {
	return s.PriceList
}

func (s *GetApplicationResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetApplicationResponseBodyData) GetResourceList() []*GetApplicationResponseBodyDataResourceList {
	return s.ResourceList
}

func (s *GetApplicationResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetApplicationResponseBodyData) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetApplicationResponseBodyData) SetApplicationId(v string) *GetApplicationResponseBodyData {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetChecklist(v []*GetApplicationResponseBodyDataChecklist) *GetApplicationResponseBodyData {
	s.Checklist = v
	return s
}

func (s *GetApplicationResponseBodyData) SetComplianceList(v []*GetApplicationResponseBodyDataComplianceList) *GetApplicationResponseBodyData {
	s.ComplianceList = v
	return s
}

func (s *GetApplicationResponseBodyData) SetCreateTime(v string) *GetApplicationResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetDeployPercent(v float64) *GetApplicationResponseBodyData {
	s.DeployPercent = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetDescription(v string) *GetApplicationResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetError(v string) *GetApplicationResponseBodyData {
	s.Error = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetImageURL(v string) *GetApplicationResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetName(v string) *GetApplicationResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetPriceList(v []*GetApplicationResponseBodyDataPriceList) *GetApplicationResponseBodyData {
	s.PriceList = v
	return s
}

func (s *GetApplicationResponseBodyData) SetResourceGroupId(v string) *GetApplicationResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetResourceList(v []*GetApplicationResponseBodyDataResourceList) *GetApplicationResponseBodyData {
	s.ResourceList = v
	return s
}

func (s *GetApplicationResponseBodyData) SetStatus(v string) *GetApplicationResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetTemplateId(v string) *GetApplicationResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *GetApplicationResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataChecklist struct {
	// The resource tag.
	//
	// example:
	//
	// Create
	Lifecycle *string `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	// The region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The message returned for verification.
	//
	// example:
	//
	// The ID of the region.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The service code.
	//
	// example:
	//
	// vpc
	ResourceCode *string `json:"ResourceCode,omitempty" xml:"ResourceCode,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// vpc
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The verification result.
	//
	// example:
	//
	// Finish
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The resource specifications.
	//
	// example:
	//
	// 192.168.0.0/16
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
}

func (s GetApplicationResponseBodyDataChecklist) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataChecklist) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataChecklist) GetLifecycle() *string {
	return s.Lifecycle
}

func (s *GetApplicationResponseBodyDataChecklist) GetRegion() *string {
	return s.Region
}

func (s *GetApplicationResponseBodyDataChecklist) GetRemark() *string {
	return s.Remark
}

func (s *GetApplicationResponseBodyDataChecklist) GetResourceCode() *string {
	return s.ResourceCode
}

func (s *GetApplicationResponseBodyDataChecklist) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetApplicationResponseBodyDataChecklist) GetResult() *string {
	return s.Result
}

func (s *GetApplicationResponseBodyDataChecklist) GetSpecification() *string {
	return s.Specification
}

func (s *GetApplicationResponseBodyDataChecklist) SetLifecycle(v string) *GetApplicationResponseBodyDataChecklist {
	s.Lifecycle = &v
	return s
}

func (s *GetApplicationResponseBodyDataChecklist) SetRegion(v string) *GetApplicationResponseBodyDataChecklist {
	s.Region = &v
	return s
}

func (s *GetApplicationResponseBodyDataChecklist) SetRemark(v string) *GetApplicationResponseBodyDataChecklist {
	s.Remark = &v
	return s
}

func (s *GetApplicationResponseBodyDataChecklist) SetResourceCode(v string) *GetApplicationResponseBodyDataChecklist {
	s.ResourceCode = &v
	return s
}

func (s *GetApplicationResponseBodyDataChecklist) SetResourceName(v string) *GetApplicationResponseBodyDataChecklist {
	s.ResourceName = &v
	return s
}

func (s *GetApplicationResponseBodyDataChecklist) SetResult(v string) *GetApplicationResponseBodyDataChecklist {
	s.Result = &v
	return s
}

func (s *GetApplicationResponseBodyDataChecklist) SetSpecification(v string) *GetApplicationResponseBodyDataChecklist {
	s.Specification = &v
	return s
}

func (s *GetApplicationResponseBodyDataChecklist) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataComplianceList struct {
	ResourceCode *string                                              `json:"ResourceCode,omitempty" xml:"ResourceCode,omitempty"`
	ResourceName *string                                              `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	Rules        []*GetApplicationResponseBodyDataComplianceListRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s GetApplicationResponseBodyDataComplianceList) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataComplianceList) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataComplianceList) GetResourceCode() *string {
	return s.ResourceCode
}

func (s *GetApplicationResponseBodyDataComplianceList) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetApplicationResponseBodyDataComplianceList) GetRules() []*GetApplicationResponseBodyDataComplianceListRules {
	return s.Rules
}

func (s *GetApplicationResponseBodyDataComplianceList) SetResourceCode(v string) *GetApplicationResponseBodyDataComplianceList {
	s.ResourceCode = &v
	return s
}

func (s *GetApplicationResponseBodyDataComplianceList) SetResourceName(v string) *GetApplicationResponseBodyDataComplianceList {
	s.ResourceName = &v
	return s
}

func (s *GetApplicationResponseBodyDataComplianceList) SetRules(v []*GetApplicationResponseBodyDataComplianceListRules) *GetApplicationResponseBodyDataComplianceList {
	s.Rules = v
	return s
}

func (s *GetApplicationResponseBodyDataComplianceList) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataComplianceListRules struct {
	RuleDetail *string `json:"ruleDetail,omitempty" xml:"ruleDetail,omitempty"`
	RuleId     *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
}

func (s GetApplicationResponseBodyDataComplianceListRules) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataComplianceListRules) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataComplianceListRules) GetRuleDetail() *string {
	return s.RuleDetail
}

func (s *GetApplicationResponseBodyDataComplianceListRules) GetRuleId() *string {
	return s.RuleId
}

func (s *GetApplicationResponseBodyDataComplianceListRules) SetRuleDetail(v string) *GetApplicationResponseBodyDataComplianceListRules {
	s.RuleDetail = &v
	return s
}

func (s *GetApplicationResponseBodyDataComplianceListRules) SetRuleId(v string) *GetApplicationResponseBodyDataComplianceListRules {
	s.RuleId = &v
	return s
}

func (s *GetApplicationResponseBodyDataComplianceListRules) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataPriceList struct {
	// The billing method.
	//
	// example:
	//
	// PayAsYouGo
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The quantity.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// ecs
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Resource Fill Labels.
	//
	// example:
	//
	// Create
	Lifecycle *string `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	// The unit price of the instance.
	//
	// example:
	//
	// 0.01
	OnePrice *float64 `json:"OnePrice,omitempty" xml:"OnePrice,omitempty"`
	// The original price of the instance.
	//
	// example:
	//
	// 3.570
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The service duration.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The total price.
	//
	// example:
	//
	// 0.01
	Price *float64 `json:"Price,omitempty" xml:"Price,omitempty"`
	// Unit: USD per hour
	//
	// example:
	//
	// The service duration.
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	// The region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The error message that is returned when a price query fails.
	//
	// example:
	//
	// ecs.e3.large
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Product code
	//
	// example:
	//
	// ecs
	ResourceCode *string `json:"ResourceCode,omitempty" xml:"ResourceCode,omitempty"`
	// The instance type. This parameter indicates the information about the instance type. For example, 192.168.0.0/16 may be returned for a Virtual Private Cloud (VPC) instance, ecs.g5.large may be returned for an Elastic Compute Service (ECS) instance, and slb.s1.small may be returned for a Server Load Balancer (SLB) instance. If the resource does not have a specific type, an empty value is returned.
	//
	// example:
	//
	// The billing method.
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// The creation mode. Valid values:\\
	//
	// 1: creates a new instance.\\
	//
	// 2: imports an instance.
	//
	// example:
	//
	// 1
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetApplicationResponseBodyDataPriceList) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataPriceList) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPriceList) GetChargeType() *string {
	return s.ChargeType
}

func (s *GetApplicationResponseBodyDataPriceList) GetCount() *int32 {
	return s.Count
}

func (s *GetApplicationResponseBodyDataPriceList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetApplicationResponseBodyDataPriceList) GetLifecycle() *string {
	return s.Lifecycle
}

func (s *GetApplicationResponseBodyDataPriceList) GetOnePrice() *float64 {
	return s.OnePrice
}

func (s *GetApplicationResponseBodyDataPriceList) GetOriginalPrice() *float64 {
	return s.OriginalPrice
}

func (s *GetApplicationResponseBodyDataPriceList) GetPeriod() *int32 {
	return s.Period
}

func (s *GetApplicationResponseBodyDataPriceList) GetPrice() *float64 {
	return s.Price
}

func (s *GetApplicationResponseBodyDataPriceList) GetPriceUnit() *string {
	return s.PriceUnit
}

func (s *GetApplicationResponseBodyDataPriceList) GetRegion() *string {
	return s.Region
}

func (s *GetApplicationResponseBodyDataPriceList) GetRemark() *string {
	return s.Remark
}

func (s *GetApplicationResponseBodyDataPriceList) GetResourceCode() *string {
	return s.ResourceCode
}

func (s *GetApplicationResponseBodyDataPriceList) GetSpecification() *string {
	return s.Specification
}

func (s *GetApplicationResponseBodyDataPriceList) GetType() *string {
	return s.Type
}

func (s *GetApplicationResponseBodyDataPriceList) SetChargeType(v string) *GetApplicationResponseBodyDataPriceList {
	s.ChargeType = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetCount(v int32) *GetApplicationResponseBodyDataPriceList {
	s.Count = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetInstanceName(v string) *GetApplicationResponseBodyDataPriceList {
	s.InstanceName = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetLifecycle(v string) *GetApplicationResponseBodyDataPriceList {
	s.Lifecycle = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetOnePrice(v float64) *GetApplicationResponseBodyDataPriceList {
	s.OnePrice = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetOriginalPrice(v float64) *GetApplicationResponseBodyDataPriceList {
	s.OriginalPrice = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetPeriod(v int32) *GetApplicationResponseBodyDataPriceList {
	s.Period = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetPrice(v float64) *GetApplicationResponseBodyDataPriceList {
	s.Price = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetPriceUnit(v string) *GetApplicationResponseBodyDataPriceList {
	s.PriceUnit = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetRegion(v string) *GetApplicationResponseBodyDataPriceList {
	s.Region = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetRemark(v string) *GetApplicationResponseBodyDataPriceList {
	s.Remark = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetResourceCode(v string) *GetApplicationResponseBodyDataPriceList {
	s.ResourceCode = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetSpecification(v string) *GetApplicationResponseBodyDataPriceList {
	s.Specification = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetType(v string) *GetApplicationResponseBodyDataPriceList {
	s.Type = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyDataResourceList struct {
	// The billing method.
	//
	// example:
	//
	// PayAsYouGo
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The resource tag.
	//
	// example:
	//
	// Create
	Lifecycle *string `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	NodeLabel *string `json:"NodeLabel,omitempty" xml:"NodeLabel,omitempty"`
	// The deployment result.
	//
	// example:
	//
	// {"hostName":"iZ2zehnzxqixu1pywsfbx1Z","memory":32768.0,"creationTime":"2021-09-28T11:23:46Z","instanceName":"ecs","internetMaxBandwidthOut":0.0,"description":"","clusterId":"","private_ip":"192.168.0.247","instanceId":"i-2zehnzxqixu1pywsfbx1","requestId":"F1C64344-3723-51A0-855B-5F08B5634323","zoneId":"cn-beijing-b","ioOptimized":"optimized","id":"i-2zehnzxqixu1pywsfbx1","instanceNetworkType":"vpc","instanceChargeType":"PostPaid","imageId":"centos_8_4_x64_20G_alibase_20210824.vhd","serialNumber":"cee246c4-38f3-4bf3-950b-c17e88ff6527","vlanId":"","instanceType":"ecs.e3.large","cpu":4.0,"creditSpecification":"","internetMaxBandwidthIn":-1.0,"expiredTime":"2099-12-31T15:59Z","internetChargeType":"PayByTraffic","regionId":"cn-beijing","refId":"79224644_0","stoppedMode":"Not-applicable","status":"Running"}
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The service code.
	//
	// example:
	//
	// ecs
	ResourceCode *string `json:"ResourceCode,omitempty" xml:"ResourceCode,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-2zehnzxqixu1pywsfbx1
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// ecs
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// ecs
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The resource deployment result.
	//
	// example:
	//
	// Finish
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetApplicationResponseBodyDataResourceList) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyDataResourceList) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataResourceList) GetChargeType() *string {
	return s.ChargeType
}

func (s *GetApplicationResponseBodyDataResourceList) GetLifecycle() *string {
	return s.Lifecycle
}

func (s *GetApplicationResponseBodyDataResourceList) GetNodeLabel() *string {
	return s.NodeLabel
}

func (s *GetApplicationResponseBodyDataResourceList) GetRemark() *string {
	return s.Remark
}

func (s *GetApplicationResponseBodyDataResourceList) GetResourceCode() *string {
	return s.ResourceCode
}

func (s *GetApplicationResponseBodyDataResourceList) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetApplicationResponseBodyDataResourceList) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetApplicationResponseBodyDataResourceList) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetApplicationResponseBodyDataResourceList) GetStatus() *string {
	return s.Status
}

func (s *GetApplicationResponseBodyDataResourceList) SetChargeType(v string) *GetApplicationResponseBodyDataResourceList {
	s.ChargeType = &v
	return s
}

func (s *GetApplicationResponseBodyDataResourceList) SetLifecycle(v string) *GetApplicationResponseBodyDataResourceList {
	s.Lifecycle = &v
	return s
}

func (s *GetApplicationResponseBodyDataResourceList) SetNodeLabel(v string) *GetApplicationResponseBodyDataResourceList {
	s.NodeLabel = &v
	return s
}

func (s *GetApplicationResponseBodyDataResourceList) SetRemark(v string) *GetApplicationResponseBodyDataResourceList {
	s.Remark = &v
	return s
}

func (s *GetApplicationResponseBodyDataResourceList) SetResourceCode(v string) *GetApplicationResponseBodyDataResourceList {
	s.ResourceCode = &v
	return s
}

func (s *GetApplicationResponseBodyDataResourceList) SetResourceId(v string) *GetApplicationResponseBodyDataResourceList {
	s.ResourceId = &v
	return s
}

func (s *GetApplicationResponseBodyDataResourceList) SetResourceName(v string) *GetApplicationResponseBodyDataResourceList {
	s.ResourceName = &v
	return s
}

func (s *GetApplicationResponseBodyDataResourceList) SetResourceType(v string) *GetApplicationResponseBodyDataResourceList {
	s.ResourceType = &v
	return s
}

func (s *GetApplicationResponseBodyDataResourceList) SetStatus(v string) *GetApplicationResponseBodyDataResourceList {
	s.Status = &v
	return s
}

func (s *GetApplicationResponseBodyDataResourceList) Validate() error {
	return dara.Validate(s)
}
