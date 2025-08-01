// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListGatewayResponseBody
	GetCode() *int32
	SetData(v *ListGatewayResponseBodyData) *ListGatewayResponseBody
	GetData() *ListGatewayResponseBodyData
	SetHttpStatusCode(v int32) *ListGatewayResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListGatewayResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGatewayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListGatewayResponseBody
	GetSuccess() *bool
}

type ListGatewayResponseBody struct {
	// The return value.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ListGatewayResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// You are not authorized to perform this operation. Action: mse:ListGateway, Resource: acs:mse:cn-hangzhou:102123:*
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 34300B3-52EC-5049-8C96-914098CF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListGatewayResponseBody) GetData() *ListGatewayResponseBodyData {
	return s.Data
}

func (s *ListGatewayResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListGatewayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListGatewayResponseBody) SetCode(v int32) *ListGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *ListGatewayResponseBody) SetData(v *ListGatewayResponseBodyData) *ListGatewayResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayResponseBody) SetHttpStatusCode(v int32) *ListGatewayResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGatewayResponseBody) SetMessage(v string) *ListGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *ListGatewayResponseBody) SetRequestId(v string) *ListGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayResponseBody) SetSuccess(v bool) *ListGatewayResponseBody {
	s.Success = &v
	return s
}

func (s *ListGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListGatewayResponseBodyData struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The data returned.
	Result []*ListGatewayResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 9
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListGatewayResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayResponseBodyData) GetResult() []*ListGatewayResponseBodyDataResult {
	return s.Result
}

func (s *ListGatewayResponseBodyData) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *ListGatewayResponseBodyData) SetPageNumber(v int32) *ListGatewayResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayResponseBodyData) SetPageSize(v int32) *ListGatewayResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListGatewayResponseBodyData) SetResult(v []*ListGatewayResponseBodyDataResult) *ListGatewayResponseBodyData {
	s.Result = v
	return s
}

func (s *ListGatewayResponseBodyData) SetTotalSize(v int64) *ListGatewayResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListGatewayResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListGatewayResponseBodyDataResult struct {
	// Indicates whether Application High Availability Service (AHAS) is activated.
	//
	// example:
	//
	// false
	AhasOn *bool `json:"AhasOn,omitempty" xml:"AhasOn,omitempty"`
	// The version of the application.
	//
	// example:
	//
	// 1.0.1.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// Indicates whether Application Real-Time Monitoring Service (ARMS) is activated.
	//
	// example:
	//
	// false
	ArmsOn *bool `json:"ArmsOn,omitempty" xml:"ArmsOn,omitempty"`
	// The billing method.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The current version of the gateway.
	//
	// example:
	//
	// 0.1.0-mse-gw
	CurrentVersion *string `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// Indicates whether auto scale-out is enabled.
	//
	// example:
	//
	// true
	Elastic *bool `json:"Elastic,omitempty" xml:"Elastic,omitempty"`
	// The ID of the elastic gateway. This parameter is returned if auto scale-out is used.
	//
	// example:
	//
	// mse_ingresselastic_public_cn-uqm3d0*****
	ElasticInstanceId *string `json:"ElasticInstanceId,omitempty" xml:"ElasticInstanceId,omitempty"`
	// The auto scale-out policy.
	ElasticPolicy *ListGatewayResponseBodyDataResultElasticPolicy `json:"ElasticPolicy,omitempty" xml:"ElasticPolicy,omitempty" type:"Struct"`
	// The number of replicas that are automatically scaled out.
	//
	// example:
	//
	// 2
	ElasticReplica *int32 `json:"ElasticReplica,omitempty" xml:"ElasticReplica,omitempty"`
	// The type of auto scale-out. Valid value:
	//
	// 	- CronHPA: scale-out by time
	//
	// example:
	//
	// CronHPA
	ElasticType *string `json:"ElasticType,omitempty" xml:"ElasticType,omitempty"`
	// The time when the instance expires.
	//
	// example:
	//
	// 4792060800000
	EndDate      *string                                          `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	GatewayEntry []*ListGatewayResponseBodyDataResultGatewayEntry `json:"GatewayEntry,omitempty" xml:"GatewayEntry,omitempty" type:"Repeated"`
	// The gateway type.
	//
	// example:
	//
	// Ingress
	GatewayType *string `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-5017305290e14cebb8ce5cb6a4****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The version of the gateway.
	//
	// example:
	//
	// 1.2.9
	GatewayVersion *string `json:"GatewayVersion,omitempty" xml:"GatewayVersion,omitempty"`
	// The time when the gateway was created.
	//
	// example:
	//
	// 2021-09-13 19:24:23
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the gateway was modified.
	//
	// example:
	//
	// 2021-09-13 19:24:23
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The gateway ID.
	//
	// example:
	//
	// 153
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The configurations.
	InitConfig *ListGatewayResponseBodyDataResultInitConfig `json:"InitConfig,omitempty" xml:"InitConfig,omitempty" type:"Struct"`
	// The instance ID.
	//
	// example:
	//
	// mse_ingresspost-cn-0jbvrcex****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The details of the Internet-facing SLB instances.
	InternetSlb []*ListGatewayResponseBodyDataResultInternetSlb `json:"InternetSlb,omitempty" xml:"InternetSlb,omitempty" type:"Repeated"`
	// The latest version of the gateway.
	//
	// example:
	//
	// 0.1.0-mse-gw
	LatestVersion     *string                                             `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	MaintenancePeriod *ListGatewayResponseBodyDataResultMaintenancePeriod `json:"MaintenancePeriod,omitempty" xml:"MaintenancePeriod,omitempty" type:"Struct"`
	// The resource tag.
	//
	// example:
	//
	// {"tagKey":"tagValue"}
	MseTag     *string `json:"MseTag,omitempty" xml:"MseTag,omitempty"`
	MseVersion *string `json:"MseVersion,omitempty" xml:"MseVersion,omitempty"`
	// Indicates whether the instance was forcefully upgraded.
	//
	// example:
	//
	// false
	MustUpgrade *bool `json:"MustUpgrade,omitempty" xml:"MustUpgrade,omitempty"`
	// The gateway name.
	//
	// example:
	//
	// tesa-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The user information.
	//
	// example:
	//
	// 18278117654342
	PrimaryUser *string `json:"PrimaryUser,omitempty" xml:"PrimaryUser,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The number of replicas.
	//
	// example:
	//
	// 2
	Replica *int32 `json:"Replica,omitempty" xml:"Replica,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-97hg87vi****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether rollbacks are allowed.
	//
	// example:
	//
	// false
	RollBack *bool `json:"RollBack,omitempty" xml:"RollBack,omitempty"`
	// The details of Server Load Balancer (SLB) instances.
	Slb []*ListGatewayResponseBodyDataResultSlb `json:"Slb,omitempty" xml:"Slb,omitempty" type:"Repeated"`
	// The specifications of the gateway.
	//
	// example:
	//
	// MSE_GTW_1_2_200_c
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The gateway state. Valid values:
	//
	// 	- 0: The gateway is being created.
	//
	// 	- 1: The gateway failed to be created.
	//
	// 	- 2: The gateway is running.
	//
	// 	- 3: The gateway is being changed.
	//
	// 	- 4: The gateway is scaling in.
	//
	// 	- 6: The gateway is scaling out.
	//
	// 	- 8: The gateway is being deleted.
	//
	// 	- 9: The gateway is suspended and is to be released.
	//
	// 	- 10: The gateway is restarting.
	//
	// 	- 11: The gateway is being rebuilt.
	//
	// 	- 12: The gateway is being upgraded.
	//
	// 	- 13: The gateway failed to be upgraded.
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The description of the gateway state.
	//
	// example:
	//
	// Restarting
	StatusDesc *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	// Indicates whether WebAssembly (Wasm) is supported.
	//
	// example:
	//
	// true
	SupportWasm *bool `json:"SupportWasm,omitempty" xml:"SupportWasm,omitempty"`
	// The tag.
	//
	// example:
	//
	// test
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The total number of replicas, including the number of replicas that are automatically scaled out.
	//
	// example:
	//
	// 4
	TotalReplica *int32 `json:"TotalReplica,omitempty" xml:"TotalReplica,omitempty"`
	// Indicates whether the instance was upgraded.
	//
	// example:
	//
	// false
	Upgrade          *bool   `json:"Upgrade,omitempty" xml:"Upgrade,omitempty"`
	VersionLifecycle *string `json:"VersionLifecycle,omitempty" xml:"VersionLifecycle,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the gateway belongs.
	//
	// example:
	//
	// vpc-uf6heojei217tv14*****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the secondary vSwitch.
	//
	// example:
	//
	// vsw-bpbrveck45nf****
	Vswitch2 *string `json:"Vswitch2,omitempty" xml:"Vswitch2,omitempty"`
}

func (s ListGatewayResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListGatewayResponseBodyDataResult) GetAhasOn() *bool {
	return s.AhasOn
}

func (s *ListGatewayResponseBodyDataResult) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ListGatewayResponseBodyDataResult) GetArmsOn() *bool {
	return s.ArmsOn
}

func (s *ListGatewayResponseBodyDataResult) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListGatewayResponseBodyDataResult) GetCurrentVersion() *string {
	return s.CurrentVersion
}

func (s *ListGatewayResponseBodyDataResult) GetElastic() *bool {
	return s.Elastic
}

func (s *ListGatewayResponseBodyDataResult) GetElasticInstanceId() *string {
	return s.ElasticInstanceId
}

func (s *ListGatewayResponseBodyDataResult) GetElasticPolicy() *ListGatewayResponseBodyDataResultElasticPolicy {
	return s.ElasticPolicy
}

func (s *ListGatewayResponseBodyDataResult) GetElasticReplica() *int32 {
	return s.ElasticReplica
}

func (s *ListGatewayResponseBodyDataResult) GetElasticType() *string {
	return s.ElasticType
}

func (s *ListGatewayResponseBodyDataResult) GetEndDate() *string {
	return s.EndDate
}

func (s *ListGatewayResponseBodyDataResult) GetGatewayEntry() []*ListGatewayResponseBodyDataResultGatewayEntry {
	return s.GatewayEntry
}

func (s *ListGatewayResponseBodyDataResult) GetGatewayType() *string {
	return s.GatewayType
}

func (s *ListGatewayResponseBodyDataResult) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListGatewayResponseBodyDataResult) GetGatewayVersion() *string {
	return s.GatewayVersion
}

func (s *ListGatewayResponseBodyDataResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListGatewayResponseBodyDataResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListGatewayResponseBodyDataResult) GetId() *int64 {
	return s.Id
}

func (s *ListGatewayResponseBodyDataResult) GetInitConfig() *ListGatewayResponseBodyDataResultInitConfig {
	return s.InitConfig
}

func (s *ListGatewayResponseBodyDataResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListGatewayResponseBodyDataResult) GetInternetSlb() []*ListGatewayResponseBodyDataResultInternetSlb {
	return s.InternetSlb
}

func (s *ListGatewayResponseBodyDataResult) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *ListGatewayResponseBodyDataResult) GetMaintenancePeriod() *ListGatewayResponseBodyDataResultMaintenancePeriod {
	return s.MaintenancePeriod
}

func (s *ListGatewayResponseBodyDataResult) GetMseTag() *string {
	return s.MseTag
}

func (s *ListGatewayResponseBodyDataResult) GetMseVersion() *string {
	return s.MseVersion
}

func (s *ListGatewayResponseBodyDataResult) GetMustUpgrade() *bool {
	return s.MustUpgrade
}

func (s *ListGatewayResponseBodyDataResult) GetName() *string {
	return s.Name
}

func (s *ListGatewayResponseBodyDataResult) GetPrimaryUser() *string {
	return s.PrimaryUser
}

func (s *ListGatewayResponseBodyDataResult) GetRegion() *string {
	return s.Region
}

func (s *ListGatewayResponseBodyDataResult) GetReplica() *int32 {
	return s.Replica
}

func (s *ListGatewayResponseBodyDataResult) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListGatewayResponseBodyDataResult) GetRollBack() *bool {
	return s.RollBack
}

func (s *ListGatewayResponseBodyDataResult) GetSlb() []*ListGatewayResponseBodyDataResultSlb {
	return s.Slb
}

func (s *ListGatewayResponseBodyDataResult) GetSpec() *string {
	return s.Spec
}

func (s *ListGatewayResponseBodyDataResult) GetStatus() *int32 {
	return s.Status
}

func (s *ListGatewayResponseBodyDataResult) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *ListGatewayResponseBodyDataResult) GetSupportWasm() *bool {
	return s.SupportWasm
}

func (s *ListGatewayResponseBodyDataResult) GetTag() *string {
	return s.Tag
}

func (s *ListGatewayResponseBodyDataResult) GetTotalReplica() *int32 {
	return s.TotalReplica
}

func (s *ListGatewayResponseBodyDataResult) GetUpgrade() *bool {
	return s.Upgrade
}

func (s *ListGatewayResponseBodyDataResult) GetVersionLifecycle() *string {
	return s.VersionLifecycle
}

func (s *ListGatewayResponseBodyDataResult) GetVpcId() *string {
	return s.VpcId
}

func (s *ListGatewayResponseBodyDataResult) GetVswitch2() *string {
	return s.Vswitch2
}

func (s *ListGatewayResponseBodyDataResult) SetAhasOn(v bool) *ListGatewayResponseBodyDataResult {
	s.AhasOn = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetAppVersion(v string) *ListGatewayResponseBodyDataResult {
	s.AppVersion = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetArmsOn(v bool) *ListGatewayResponseBodyDataResult {
	s.ArmsOn = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetChargeType(v string) *ListGatewayResponseBodyDataResult {
	s.ChargeType = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetCurrentVersion(v string) *ListGatewayResponseBodyDataResult {
	s.CurrentVersion = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetElastic(v bool) *ListGatewayResponseBodyDataResult {
	s.Elastic = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetElasticInstanceId(v string) *ListGatewayResponseBodyDataResult {
	s.ElasticInstanceId = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetElasticPolicy(v *ListGatewayResponseBodyDataResultElasticPolicy) *ListGatewayResponseBodyDataResult {
	s.ElasticPolicy = v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetElasticReplica(v int32) *ListGatewayResponseBodyDataResult {
	s.ElasticReplica = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetElasticType(v string) *ListGatewayResponseBodyDataResult {
	s.ElasticType = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetEndDate(v string) *ListGatewayResponseBodyDataResult {
	s.EndDate = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetGatewayEntry(v []*ListGatewayResponseBodyDataResultGatewayEntry) *ListGatewayResponseBodyDataResult {
	s.GatewayEntry = v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetGatewayType(v string) *ListGatewayResponseBodyDataResult {
	s.GatewayType = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetGatewayUniqueId(v string) *ListGatewayResponseBodyDataResult {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetGatewayVersion(v string) *ListGatewayResponseBodyDataResult {
	s.GatewayVersion = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetGmtCreate(v string) *ListGatewayResponseBodyDataResult {
	s.GmtCreate = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetGmtModified(v string) *ListGatewayResponseBodyDataResult {
	s.GmtModified = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetId(v int64) *ListGatewayResponseBodyDataResult {
	s.Id = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetInitConfig(v *ListGatewayResponseBodyDataResultInitConfig) *ListGatewayResponseBodyDataResult {
	s.InitConfig = v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetInstanceId(v string) *ListGatewayResponseBodyDataResult {
	s.InstanceId = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetInternetSlb(v []*ListGatewayResponseBodyDataResultInternetSlb) *ListGatewayResponseBodyDataResult {
	s.InternetSlb = v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetLatestVersion(v string) *ListGatewayResponseBodyDataResult {
	s.LatestVersion = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetMaintenancePeriod(v *ListGatewayResponseBodyDataResultMaintenancePeriod) *ListGatewayResponseBodyDataResult {
	s.MaintenancePeriod = v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetMseTag(v string) *ListGatewayResponseBodyDataResult {
	s.MseTag = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetMseVersion(v string) *ListGatewayResponseBodyDataResult {
	s.MseVersion = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetMustUpgrade(v bool) *ListGatewayResponseBodyDataResult {
	s.MustUpgrade = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetName(v string) *ListGatewayResponseBodyDataResult {
	s.Name = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetPrimaryUser(v string) *ListGatewayResponseBodyDataResult {
	s.PrimaryUser = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetRegion(v string) *ListGatewayResponseBodyDataResult {
	s.Region = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetReplica(v int32) *ListGatewayResponseBodyDataResult {
	s.Replica = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetResourceGroupId(v string) *ListGatewayResponseBodyDataResult {
	s.ResourceGroupId = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetRollBack(v bool) *ListGatewayResponseBodyDataResult {
	s.RollBack = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetSlb(v []*ListGatewayResponseBodyDataResultSlb) *ListGatewayResponseBodyDataResult {
	s.Slb = v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetSpec(v string) *ListGatewayResponseBodyDataResult {
	s.Spec = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetStatus(v int32) *ListGatewayResponseBodyDataResult {
	s.Status = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetStatusDesc(v string) *ListGatewayResponseBodyDataResult {
	s.StatusDesc = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetSupportWasm(v bool) *ListGatewayResponseBodyDataResult {
	s.SupportWasm = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetTag(v string) *ListGatewayResponseBodyDataResult {
	s.Tag = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetTotalReplica(v int32) *ListGatewayResponseBodyDataResult {
	s.TotalReplica = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetUpgrade(v bool) *ListGatewayResponseBodyDataResult {
	s.Upgrade = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetVersionLifecycle(v string) *ListGatewayResponseBodyDataResult {
	s.VersionLifecycle = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetVpcId(v string) *ListGatewayResponseBodyDataResult {
	s.VpcId = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) SetVswitch2(v string) *ListGatewayResponseBodyDataResult {
	s.Vswitch2 = &v
	return s
}

func (s *ListGatewayResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}

type ListGatewayResponseBodyDataResultElasticPolicy struct {
	// Indicates whether auto scale-out is enabled.
	//
	// example:
	//
	// true
	Elastic *bool `json:"Elastic,omitempty" xml:"Elastic,omitempty"`
	// The type of auto scale-out. Valid value:
	//
	// 	- CronHPA: scale-out by time
	//
	// example:
	//
	// CronHPA
	ElasticType               *string                                                                    `json:"ElasticType,omitempty" xml:"ElasticType,omitempty"`
	EnableScaleTimePolicyList []*ListGatewayResponseBodyDataResultElasticPolicyEnableScaleTimePolicyList `json:"EnableScaleTimePolicyList,omitempty" xml:"EnableScaleTimePolicyList,omitempty" type:"Repeated"`
	LoadWarningThreshold      *bool                                                                      `json:"LoadWarningThreshold,omitempty" xml:"LoadWarningThreshold,omitempty"`
	// The maximum number of instances that are automatically scaled out. This parameter is used for horizontal scale-out.
	//
	// example:
	//
	// 10
	MaxReplica *int32 `json:"MaxReplica,omitempty" xml:"MaxReplica,omitempty"`
	// The time policy list for auto scale-out.
	TimePolicyList []*ListGatewayResponseBodyDataResultElasticPolicyTimePolicyList `json:"TimePolicyList,omitempty" xml:"TimePolicyList,omitempty" type:"Repeated"`
}

func (s ListGatewayResponseBodyDataResultElasticPolicy) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayResponseBodyDataResultElasticPolicy) GoString() string {
	return s.String()
}

func (s *ListGatewayResponseBodyDataResultElasticPolicy) GetElastic() *bool {
	return s.Elastic
}

func (s *ListGatewayResponseBodyDataResultElasticPolicy) GetElasticType() *string {
	return s.ElasticType
}

func (s *ListGatewayResponseBodyDataResultElasticPolicy) GetEnableScaleTimePolicyList() []*ListGatewayResponseBodyDataResultElasticPolicyEnableScaleTimePolicyList {
	return s.EnableScaleTimePolicyList
}

func (s *ListGatewayResponseBodyDataResultElasticPolicy) GetLoadWarningThreshold() *bool {
	return s.LoadWarningThreshold
}

func (s *ListGatewayResponseBodyDataResultElasticPolicy) GetMaxReplica() *int32 {
	return s.MaxReplica
}

func (s *ListGatewayResponseBodyDataResultElasticPolicy) GetTimePolicyList() []*ListGatewayResponseBodyDataResultElasticPolicyTimePolicyList {
	return s.TimePolicyList
}

func (s *ListGatewayResponseBodyDataResultElasticPolicy) SetElastic(v bool) *ListGatewayResponseBodyDataResultElasticPolicy {
	s.Elastic = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultElasticPolicy) SetElasticType(v string) *ListGatewayResponseBodyDataResultElasticPolicy {
	s.ElasticType = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultElasticPolicy) SetEnableScaleTimePolicyList(v []*ListGatewayResponseBodyDataResultElasticPolicyEnableScaleTimePolicyList) *ListGatewayResponseBodyDataResultElasticPolicy {
	s.EnableScaleTimePolicyList = v
	return s
}

func (s *ListGatewayResponseBodyDataResultElasticPolicy) SetLoadWarningThreshold(v bool) *ListGatewayResponseBodyDataResultElasticPolicy {
	s.LoadWarningThreshold = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultElasticPolicy) SetMaxReplica(v int32) *ListGatewayResponseBodyDataResultElasticPolicy {
	s.MaxReplica = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultElasticPolicy) SetTimePolicyList(v []*ListGatewayResponseBodyDataResultElasticPolicyTimePolicyList) *ListGatewayResponseBodyDataResultElasticPolicy {
	s.TimePolicyList = v
	return s
}

func (s *ListGatewayResponseBodyDataResultElasticPolicy) Validate() error {
	return dara.Validate(s)
}

type ListGatewayResponseBodyDataResultElasticPolicyEnableScaleTimePolicyList struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListGatewayResponseBodyDataResultElasticPolicyEnableScaleTimePolicyList) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayResponseBodyDataResultElasticPolicyEnableScaleTimePolicyList) GoString() string {
	return s.String()
}

func (s *ListGatewayResponseBodyDataResultElasticPolicyEnableScaleTimePolicyList) GetEndTime() *string {
	return s.EndTime
}

func (s *ListGatewayResponseBodyDataResultElasticPolicyEnableScaleTimePolicyList) GetStartTime() *string {
	return s.StartTime
}

func (s *ListGatewayResponseBodyDataResultElasticPolicyEnableScaleTimePolicyList) SetEndTime(v string) *ListGatewayResponseBodyDataResultElasticPolicyEnableScaleTimePolicyList {
	s.EndTime = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultElasticPolicyEnableScaleTimePolicyList) SetStartTime(v string) *ListGatewayResponseBodyDataResultElasticPolicyEnableScaleTimePolicyList {
	s.StartTime = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultElasticPolicyEnableScaleTimePolicyList) Validate() error {
	return dara.Validate(s)
}

type ListGatewayResponseBodyDataResultElasticPolicyTimePolicyList struct {
	// The expected number of replicas for auto scale-out.
	//
	// example:
	//
	// 4
	DesiredReplica *int32 `json:"DesiredReplica,omitempty" xml:"DesiredReplica,omitempty"`
	// The end time of auto scale-out.
	//
	// example:
	//
	// 09:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of auto scale-out.
	//
	// example:
	//
	// 07:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListGatewayResponseBodyDataResultElasticPolicyTimePolicyList) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayResponseBodyDataResultElasticPolicyTimePolicyList) GoString() string {
	return s.String()
}

func (s *ListGatewayResponseBodyDataResultElasticPolicyTimePolicyList) GetDesiredReplica() *int32 {
	return s.DesiredReplica
}

func (s *ListGatewayResponseBodyDataResultElasticPolicyTimePolicyList) GetEndTime() *string {
	return s.EndTime
}

func (s *ListGatewayResponseBodyDataResultElasticPolicyTimePolicyList) GetStartTime() *string {
	return s.StartTime
}

func (s *ListGatewayResponseBodyDataResultElasticPolicyTimePolicyList) SetDesiredReplica(v int32) *ListGatewayResponseBodyDataResultElasticPolicyTimePolicyList {
	s.DesiredReplica = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultElasticPolicyTimePolicyList) SetEndTime(v string) *ListGatewayResponseBodyDataResultElasticPolicyTimePolicyList {
	s.EndTime = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultElasticPolicyTimePolicyList) SetStartTime(v string) *ListGatewayResponseBodyDataResultElasticPolicyTimePolicyList {
	s.StartTime = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultElasticPolicyTimePolicyList) Validate() error {
	return dara.Validate(s)
}

type ListGatewayResponseBodyDataResultGatewayEntry struct {
	// example:
	//
	// gw-*****5c2cd6144f4bfa1c32289f45ea8.cn-hangzhou.alicloudapi.com
	EntryDomain *string   `json:"EntryDomain,omitempty" xml:"EntryDomain,omitempty"`
	HttpPorts   []*int32  `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty" type:"Repeated"`
	HttpsPorts  []*int32  `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty" type:"Repeated"`
	IpList      []*string `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	// example:
	//
	// PUB_NET
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s ListGatewayResponseBodyDataResultGatewayEntry) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayResponseBodyDataResultGatewayEntry) GoString() string {
	return s.String()
}

func (s *ListGatewayResponseBodyDataResultGatewayEntry) GetEntryDomain() *string {
	return s.EntryDomain
}

func (s *ListGatewayResponseBodyDataResultGatewayEntry) GetHttpPorts() []*int32 {
	return s.HttpPorts
}

func (s *ListGatewayResponseBodyDataResultGatewayEntry) GetHttpsPorts() []*int32 {
	return s.HttpsPorts
}

func (s *ListGatewayResponseBodyDataResultGatewayEntry) GetIpList() []*string {
	return s.IpList
}

func (s *ListGatewayResponseBodyDataResultGatewayEntry) GetNetType() *string {
	return s.NetType
}

func (s *ListGatewayResponseBodyDataResultGatewayEntry) SetEntryDomain(v string) *ListGatewayResponseBodyDataResultGatewayEntry {
	s.EntryDomain = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultGatewayEntry) SetHttpPorts(v []*int32) *ListGatewayResponseBodyDataResultGatewayEntry {
	s.HttpPorts = v
	return s
}

func (s *ListGatewayResponseBodyDataResultGatewayEntry) SetHttpsPorts(v []*int32) *ListGatewayResponseBodyDataResultGatewayEntry {
	s.HttpsPorts = v
	return s
}

func (s *ListGatewayResponseBodyDataResultGatewayEntry) SetIpList(v []*string) *ListGatewayResponseBodyDataResultGatewayEntry {
	s.IpList = v
	return s
}

func (s *ListGatewayResponseBodyDataResultGatewayEntry) SetNetType(v string) *ListGatewayResponseBodyDataResultGatewayEntry {
	s.NetType = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultGatewayEntry) Validate() error {
	return dara.Validate(s)
}

type ListGatewayResponseBodyDataResultInitConfig struct {
	// Indicates whether Web Application Firewall (WAF) is enabled.
	//
	// example:
	//
	// true
	EnableWaf *bool `json:"EnableWaf,omitempty" xml:"EnableWaf,omitempty"`
	// Indicates whether WAF is supported.
	//
	// example:
	//
	// true
	SupportWaf *bool `json:"SupportWaf,omitempty" xml:"SupportWaf,omitempty"`
}

func (s ListGatewayResponseBodyDataResultInitConfig) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayResponseBodyDataResultInitConfig) GoString() string {
	return s.String()
}

func (s *ListGatewayResponseBodyDataResultInitConfig) GetEnableWaf() *bool {
	return s.EnableWaf
}

func (s *ListGatewayResponseBodyDataResultInitConfig) GetSupportWaf() *bool {
	return s.SupportWaf
}

func (s *ListGatewayResponseBodyDataResultInitConfig) SetEnableWaf(v bool) *ListGatewayResponseBodyDataResultInitConfig {
	s.EnableWaf = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultInitConfig) SetSupportWaf(v bool) *ListGatewayResponseBodyDataResultInitConfig {
	s.SupportWaf = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultInitConfig) Validate() error {
	return dara.Validate(s)
}

type ListGatewayResponseBodyDataResultInternetSlb struct {
	// The mode of the SLB instance.
	//
	// example:
	//
	// UserHost
	GatewaySlbMode *string `json:"GatewaySlbMode,omitempty" xml:"GatewaySlbMode,omitempty"`
	// The state of the SLB instance.
	//
	// example:
	//
	// Ready
	GatewaySlbStatus *string `json:"GatewaySlbStatus,omitempty" xml:"GatewaySlbStatus,omitempty"`
	// The traffic of the gateway.
	//
	// example:
	//
	// 20
	InternetNetworkFlow *string `json:"InternetNetworkFlow,omitempty" xml:"InternetNetworkFlow,omitempty"`
	// The ID of the SLB instance.
	//
	// example:
	//
	// lb-bp1ut8asdfgucjk****
	SlbId *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	// The IP address of the SLB instance.
	//
	// example:
	//
	// 153.12.XX.XX
	SlbIp *string `json:"SlbIp,omitempty" xml:"SlbIp,omitempty"`
	// The port number of the SLB instance.
	//
	// example:
	//
	// slb.s2.small
	SlbPort *string `json:"SlbPort,omitempty" xml:"SlbPort,omitempty"`
	// The specifications of the SLB instance.
	//
	// example:
	//
	// slb.s2.small
	SlbSpec *string `json:"SlbSpec,omitempty" xml:"SlbSpec,omitempty"`
	// The description of the state.
	//
	// example:
	//
	// Creating
	StatusDesc *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	// The type of the SLB instance.
	//
	// example:
	//
	// PUB_NET
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListGatewayResponseBodyDataResultInternetSlb) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayResponseBodyDataResultInternetSlb) GoString() string {
	return s.String()
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) GetGatewaySlbMode() *string {
	return s.GatewaySlbMode
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) GetGatewaySlbStatus() *string {
	return s.GatewaySlbStatus
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) GetInternetNetworkFlow() *string {
	return s.InternetNetworkFlow
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) GetSlbId() *string {
	return s.SlbId
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) GetSlbIp() *string {
	return s.SlbIp
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) GetSlbPort() *string {
	return s.SlbPort
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) GetSlbSpec() *string {
	return s.SlbSpec
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) GetType() *string {
	return s.Type
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) SetGatewaySlbMode(v string) *ListGatewayResponseBodyDataResultInternetSlb {
	s.GatewaySlbMode = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) SetGatewaySlbStatus(v string) *ListGatewayResponseBodyDataResultInternetSlb {
	s.GatewaySlbStatus = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) SetInternetNetworkFlow(v string) *ListGatewayResponseBodyDataResultInternetSlb {
	s.InternetNetworkFlow = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) SetSlbId(v string) *ListGatewayResponseBodyDataResultInternetSlb {
	s.SlbId = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) SetSlbIp(v string) *ListGatewayResponseBodyDataResultInternetSlb {
	s.SlbIp = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) SetSlbPort(v string) *ListGatewayResponseBodyDataResultInternetSlb {
	s.SlbPort = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) SetSlbSpec(v string) *ListGatewayResponseBodyDataResultInternetSlb {
	s.SlbSpec = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) SetStatusDesc(v string) *ListGatewayResponseBodyDataResultInternetSlb {
	s.StatusDesc = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) SetType(v string) *ListGatewayResponseBodyDataResultInternetSlb {
	s.Type = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultInternetSlb) Validate() error {
	return dara.Validate(s)
}

type ListGatewayResponseBodyDataResultMaintenancePeriod struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TimeZone  *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s ListGatewayResponseBodyDataResultMaintenancePeriod) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayResponseBodyDataResultMaintenancePeriod) GoString() string {
	return s.String()
}

func (s *ListGatewayResponseBodyDataResultMaintenancePeriod) GetEndTime() *string {
	return s.EndTime
}

func (s *ListGatewayResponseBodyDataResultMaintenancePeriod) GetStartTime() *string {
	return s.StartTime
}

func (s *ListGatewayResponseBodyDataResultMaintenancePeriod) GetTimeZone() *string {
	return s.TimeZone
}

func (s *ListGatewayResponseBodyDataResultMaintenancePeriod) SetEndTime(v string) *ListGatewayResponseBodyDataResultMaintenancePeriod {
	s.EndTime = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultMaintenancePeriod) SetStartTime(v string) *ListGatewayResponseBodyDataResultMaintenancePeriod {
	s.StartTime = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultMaintenancePeriod) SetTimeZone(v string) *ListGatewayResponseBodyDataResultMaintenancePeriod {
	s.TimeZone = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultMaintenancePeriod) Validate() error {
	return dara.Validate(s)
}

type ListGatewayResponseBodyDataResultSlb struct {
	// The mode of the SLB instance.
	//
	// example:
	//
	// UserHost
	GatewaySlbMode *string `json:"GatewaySlbMode,omitempty" xml:"GatewaySlbMode,omitempty"`
	// The state of the SLB instance.
	//
	// example:
	//
	// Ready
	GatewaySlbStatus *string `json:"GatewaySlbStatus,omitempty" xml:"GatewaySlbStatus,omitempty"`
	// The ID of the SLB instance.
	//
	// example:
	//
	// lb-bp1ut8asdfgucjk****
	SlbId *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	// The IP address of the SLB instance.
	//
	// example:
	//
	// 153.12.XX.XX
	SlbIp *string `json:"SlbIp,omitempty" xml:"SlbIp,omitempty"`
	// The port number of the SLB instance.
	//
	// example:
	//
	// 80
	SlbPort *string `json:"SlbPort,omitempty" xml:"SlbPort,omitempty"`
	// The specifications of the SLB instance.
	//
	// example:
	//
	// slb.s2.small
	SlbSpec *string `json:"SlbSpec,omitempty" xml:"SlbSpec,omitempty"`
	// The description of the state.
	//
	// example:
	//
	// Creating
	StatusDesc *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	// The network type. Valid values:
	//
	// 	- PUB_NET
	//
	// 	- PRIVATE_NET
	//
	// example:
	//
	// PUB_NET
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListGatewayResponseBodyDataResultSlb) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayResponseBodyDataResultSlb) GoString() string {
	return s.String()
}

func (s *ListGatewayResponseBodyDataResultSlb) GetGatewaySlbMode() *string {
	return s.GatewaySlbMode
}

func (s *ListGatewayResponseBodyDataResultSlb) GetGatewaySlbStatus() *string {
	return s.GatewaySlbStatus
}

func (s *ListGatewayResponseBodyDataResultSlb) GetSlbId() *string {
	return s.SlbId
}

func (s *ListGatewayResponseBodyDataResultSlb) GetSlbIp() *string {
	return s.SlbIp
}

func (s *ListGatewayResponseBodyDataResultSlb) GetSlbPort() *string {
	return s.SlbPort
}

func (s *ListGatewayResponseBodyDataResultSlb) GetSlbSpec() *string {
	return s.SlbSpec
}

func (s *ListGatewayResponseBodyDataResultSlb) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *ListGatewayResponseBodyDataResultSlb) GetType() *string {
	return s.Type
}

func (s *ListGatewayResponseBodyDataResultSlb) SetGatewaySlbMode(v string) *ListGatewayResponseBodyDataResultSlb {
	s.GatewaySlbMode = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultSlb) SetGatewaySlbStatus(v string) *ListGatewayResponseBodyDataResultSlb {
	s.GatewaySlbStatus = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultSlb) SetSlbId(v string) *ListGatewayResponseBodyDataResultSlb {
	s.SlbId = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultSlb) SetSlbIp(v string) *ListGatewayResponseBodyDataResultSlb {
	s.SlbIp = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultSlb) SetSlbPort(v string) *ListGatewayResponseBodyDataResultSlb {
	s.SlbPort = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultSlb) SetSlbSpec(v string) *ListGatewayResponseBodyDataResultSlb {
	s.SlbSpec = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultSlb) SetStatusDesc(v string) *ListGatewayResponseBodyDataResultSlb {
	s.StatusDesc = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultSlb) SetType(v string) *ListGatewayResponseBodyDataResultSlb {
	s.Type = &v
	return s
}

func (s *ListGatewayResponseBodyDataResultSlb) Validate() error {
	return dara.Validate(s)
}
