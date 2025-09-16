// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetInstanceResponseBody
	GetRequestId() *string
	SetResult(v *GetInstanceResponseBodyResult) *GetInstanceResponseBody
	GetResult() *GetInstanceResponseBodyResult
}

type GetInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Response parameters
	Result *GetInstanceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceResponseBody) GetResult() *GetInstanceResponseBodyResult {
	return s.Result
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetResult(v *GetInstanceResponseBodyResult) *GetInstanceResponseBody {
	s.Result = v
	return s
}

func (s *GetInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyResult struct {
	BsVersion *string `json:"bsVersion,omitempty" xml:"bsVersion,omitempty"`
	// The billing method.
	//
	// example:
	//
	// POSYPAY
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The commodity code of the instance.
	//
	// example:
	//
	// commodityCode
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2022-06-17T02:01:26Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// ha3_test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The edition of the instance. Valid values: vector and engine.
	//
	// example:
	//
	// vector
	Edition *string `json:"edition,omitempty" xml:"edition,omitempty"`
	// The time when the instance expires.
	//
	// example:
	//
	// 1634609702
	ExpiredTime *string `json:"expiredTime,omitempty" xml:"expiredTime,omitempty"`
	// Indicates whether an overdue payment is involved.
	//
	// example:
	//
	// false
	InDebt *bool `json:"inDebt,omitempty" xml:"inDebt,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// ha-cn-7mz2qsgq301
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The lock status.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	// The network information of the instance.
	Network *GetInstanceResponseBodyResultNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
	// Specifies whether the instance is of the new version.
	//
	// example:
	//
	// true
	NewMode *bool `json:"newMode,omitempty" xml:"newMode,omitempty"`
	// Specifies whether the instance has only one node.
	//
	// example:
	//
	// false
	NoQrs *bool `json:"noQrs,omitempty" xml:"noQrs,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzjvw24el5lma
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The node specifications.
	Spec *GetInstanceResponseBodyResultSpec `json:"spec,omitempty" xml:"spec,omitempty" type:"Struct"`
	// The status of the instance. Valid values:
	//
	// 	- INIT: being initialized
	//
	// 	- WAIT_CONFIG: to be configured
	//
	// 	- CONFIG_UPDATING: configuration taking effect
	//
	// 	- READY: normal
	//
	// example:
	//
	// INIT
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tags of the instance.
	Tags []*GetInstanceResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The time when the instance was updated.
	//
	// example:
	//
	// 1634609702
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The username.
	//
	// example:
	//
	// admin
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// The version of the engine.
	//
	// example:
	//
	// ha3_3.10.0
	Version   *string `json:"version,omitempty" xml:"version,omitempty"`
	ZoneCount *int32  `json:"zoneCount,omitempty" xml:"zoneCount,omitempty"`
}

func (s GetInstanceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyResult) GetBsVersion() *string {
	return s.BsVersion
}

func (s *GetInstanceResponseBodyResult) GetChargeType() *string {
	return s.ChargeType
}

func (s *GetInstanceResponseBodyResult) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetInstanceResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetInstanceResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *GetInstanceResponseBodyResult) GetEdition() *string {
	return s.Edition
}

func (s *GetInstanceResponseBodyResult) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *GetInstanceResponseBodyResult) GetInDebt() *bool {
	return s.InDebt
}

func (s *GetInstanceResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceResponseBodyResult) GetLockMode() *string {
	return s.LockMode
}

func (s *GetInstanceResponseBodyResult) GetNetwork() *GetInstanceResponseBodyResultNetwork {
	return s.Network
}

func (s *GetInstanceResponseBodyResult) GetNewMode() *bool {
	return s.NewMode
}

func (s *GetInstanceResponseBodyResult) GetNoQrs() *bool {
	return s.NoQrs
}

func (s *GetInstanceResponseBodyResult) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetInstanceResponseBodyResult) GetSpec() *GetInstanceResponseBodyResultSpec {
	return s.Spec
}

func (s *GetInstanceResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceResponseBodyResult) GetTags() []*GetInstanceResponseBodyResultTags {
	return s.Tags
}

func (s *GetInstanceResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetInstanceResponseBodyResult) GetUserName() *string {
	return s.UserName
}

func (s *GetInstanceResponseBodyResult) GetVersion() *string {
	return s.Version
}

func (s *GetInstanceResponseBodyResult) GetZoneCount() *int32 {
	return s.ZoneCount
}

func (s *GetInstanceResponseBodyResult) SetBsVersion(v string) *GetInstanceResponseBodyResult {
	s.BsVersion = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetChargeType(v string) *GetInstanceResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetCommodityCode(v string) *GetInstanceResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetCreateTime(v string) *GetInstanceResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetDescription(v string) *GetInstanceResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetEdition(v string) *GetInstanceResponseBodyResult {
	s.Edition = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetExpiredTime(v string) *GetInstanceResponseBodyResult {
	s.ExpiredTime = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetInDebt(v bool) *GetInstanceResponseBodyResult {
	s.InDebt = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetInstanceId(v string) *GetInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetLockMode(v string) *GetInstanceResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetNetwork(v *GetInstanceResponseBodyResultNetwork) *GetInstanceResponseBodyResult {
	s.Network = v
	return s
}

func (s *GetInstanceResponseBodyResult) SetNewMode(v bool) *GetInstanceResponseBodyResult {
	s.NewMode = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetNoQrs(v bool) *GetInstanceResponseBodyResult {
	s.NoQrs = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetResourceGroupId(v string) *GetInstanceResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetSpec(v *GetInstanceResponseBodyResultSpec) *GetInstanceResponseBodyResult {
	s.Spec = v
	return s
}

func (s *GetInstanceResponseBodyResult) SetStatus(v string) *GetInstanceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetTags(v []*GetInstanceResponseBodyResultTags) *GetInstanceResponseBodyResult {
	s.Tags = v
	return s
}

func (s *GetInstanceResponseBodyResult) SetUpdateTime(v string) *GetInstanceResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetUserName(v string) *GetInstanceResponseBodyResult {
	s.UserName = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetVersion(v string) *GetInstanceResponseBodyResult {
	s.Version = &v
	return s
}

func (s *GetInstanceResponseBodyResult) SetZoneCount(v int32) *GetInstanceResponseBodyResult {
	s.ZoneCount = &v
	return s
}

func (s *GetInstanceResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyResultNetwork struct {
	// The public domain name whitelist.
	//
	// example:
	//
	// 127.0.0.1
	Allow *string `json:"allow,omitempty" xml:"allow,omitempty"`
	// The instance endpoint.
	//
	// example:
	//
	// ha-cn-35t3r****.ha.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The public endpoint.
	//
	// example:
	//
	// ha-cn-35t3ni****.public.ha.aliyuncs.com
	PublicEndpoint *string `json:"publicEndpoint,omitempty" xml:"publicEndpoint,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp11ldcf59q2n****
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-wz9axk41d9vff****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s GetInstanceResponseBodyResultNetwork) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyResultNetwork) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyResultNetwork) GetAllow() *string {
	return s.Allow
}

func (s *GetInstanceResponseBodyResultNetwork) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetInstanceResponseBodyResultNetwork) GetPublicEndpoint() *string {
	return s.PublicEndpoint
}

func (s *GetInstanceResponseBodyResultNetwork) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetInstanceResponseBodyResultNetwork) GetVpcId() *string {
	return s.VpcId
}

func (s *GetInstanceResponseBodyResultNetwork) SetAllow(v string) *GetInstanceResponseBodyResultNetwork {
	s.Allow = &v
	return s
}

func (s *GetInstanceResponseBodyResultNetwork) SetEndpoint(v string) *GetInstanceResponseBodyResultNetwork {
	s.Endpoint = &v
	return s
}

func (s *GetInstanceResponseBodyResultNetwork) SetPublicEndpoint(v string) *GetInstanceResponseBodyResultNetwork {
	s.PublicEndpoint = &v
	return s
}

func (s *GetInstanceResponseBodyResultNetwork) SetVSwitchId(v string) *GetInstanceResponseBodyResultNetwork {
	s.VSwitchId = &v
	return s
}

func (s *GetInstanceResponseBodyResultNetwork) SetVpcId(v string) *GetInstanceResponseBodyResultNetwork {
	s.VpcId = &v
	return s
}

func (s *GetInstanceResponseBodyResultNetwork) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyResultSpec struct {
	// The QRS worker specifications.
	QrsResource *GetInstanceResponseBodyResultSpecQrsResource `json:"qrsResource,omitempty" xml:"qrsResource,omitempty" type:"Struct"`
	// The searcher worker specifications.
	SearchResource *GetInstanceResponseBodyResultSpecSearchResource `json:"searchResource,omitempty" xml:"searchResource,omitempty" type:"Struct"`
}

func (s GetInstanceResponseBodyResultSpec) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyResultSpec) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyResultSpec) GetQrsResource() *GetInstanceResponseBodyResultSpecQrsResource {
	return s.QrsResource
}

func (s *GetInstanceResponseBodyResultSpec) GetSearchResource() *GetInstanceResponseBodyResultSpecSearchResource {
	return s.SearchResource
}

func (s *GetInstanceResponseBodyResultSpec) SetQrsResource(v *GetInstanceResponseBodyResultSpecQrsResource) *GetInstanceResponseBodyResultSpec {
	s.QrsResource = v
	return s
}

func (s *GetInstanceResponseBodyResultSpec) SetSearchResource(v *GetInstanceResponseBodyResultSpecSearchResource) *GetInstanceResponseBodyResultSpec {
	s.SearchResource = v
	return s
}

func (s *GetInstanceResponseBodyResultSpec) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyResultSpecQrsResource struct {
	// The category. Valid values: local_ssd, SSD, and cloud.
	//
	// example:
	//
	// local_ssd
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 2
	Cpu *int32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The storage capacity. Unit: GB.
	//
	// example:
	//
	// 100
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// The memory of the instance. Unit: GB.
	//
	// example:
	//
	// 10
	Mem *int32 `json:"mem,omitempty" xml:"mem,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 1
	NodeCount *int32 `json:"nodeCount,omitempty" xml:"nodeCount,omitempty"`
}

func (s GetInstanceResponseBodyResultSpecQrsResource) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyResultSpecQrsResource) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyResultSpecQrsResource) GetCategory() *string {
	return s.Category
}

func (s *GetInstanceResponseBodyResultSpecQrsResource) GetCpu() *int32 {
	return s.Cpu
}

func (s *GetInstanceResponseBodyResultSpecQrsResource) GetDisk() *int32 {
	return s.Disk
}

func (s *GetInstanceResponseBodyResultSpecQrsResource) GetMem() *int32 {
	return s.Mem
}

func (s *GetInstanceResponseBodyResultSpecQrsResource) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *GetInstanceResponseBodyResultSpecQrsResource) SetCategory(v string) *GetInstanceResponseBodyResultSpecQrsResource {
	s.Category = &v
	return s
}

func (s *GetInstanceResponseBodyResultSpecQrsResource) SetCpu(v int32) *GetInstanceResponseBodyResultSpecQrsResource {
	s.Cpu = &v
	return s
}

func (s *GetInstanceResponseBodyResultSpecQrsResource) SetDisk(v int32) *GetInstanceResponseBodyResultSpecQrsResource {
	s.Disk = &v
	return s
}

func (s *GetInstanceResponseBodyResultSpecQrsResource) SetMem(v int32) *GetInstanceResponseBodyResultSpecQrsResource {
	s.Mem = &v
	return s
}

func (s *GetInstanceResponseBodyResultSpecQrsResource) SetNodeCount(v int32) *GetInstanceResponseBodyResultSpecQrsResource {
	s.NodeCount = &v
	return s
}

func (s *GetInstanceResponseBodyResultSpecQrsResource) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyResultSpecSearchResource struct {
	// The category. Valid values: local_ssd, SSD, and cloud.
	//
	// example:
	//
	// local_ssd
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 2
	Cpu *int32 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The storage capacity. Unit: GB.
	//
	// example:
	//
	// 100
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// The memory of the instance. Unit: GB.
	//
	// example:
	//
	// 10
	Mem *int32 `json:"mem,omitempty" xml:"mem,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 1
	NodeCount *int32 `json:"nodeCount,omitempty" xml:"nodeCount,omitempty"`
}

func (s GetInstanceResponseBodyResultSpecSearchResource) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyResultSpecSearchResource) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyResultSpecSearchResource) GetCategory() *string {
	return s.Category
}

func (s *GetInstanceResponseBodyResultSpecSearchResource) GetCpu() *int32 {
	return s.Cpu
}

func (s *GetInstanceResponseBodyResultSpecSearchResource) GetDisk() *int32 {
	return s.Disk
}

func (s *GetInstanceResponseBodyResultSpecSearchResource) GetMem() *int32 {
	return s.Mem
}

func (s *GetInstanceResponseBodyResultSpecSearchResource) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *GetInstanceResponseBodyResultSpecSearchResource) SetCategory(v string) *GetInstanceResponseBodyResultSpecSearchResource {
	s.Category = &v
	return s
}

func (s *GetInstanceResponseBodyResultSpecSearchResource) SetCpu(v int32) *GetInstanceResponseBodyResultSpecSearchResource {
	s.Cpu = &v
	return s
}

func (s *GetInstanceResponseBodyResultSpecSearchResource) SetDisk(v int32) *GetInstanceResponseBodyResultSpecSearchResource {
	s.Disk = &v
	return s
}

func (s *GetInstanceResponseBodyResultSpecSearchResource) SetMem(v int32) *GetInstanceResponseBodyResultSpecSearchResource {
	s.Mem = &v
	return s
}

func (s *GetInstanceResponseBodyResultSpecSearchResource) SetNodeCount(v int32) *GetInstanceResponseBodyResultSpecSearchResource {
	s.NodeCount = &v
	return s
}

func (s *GetInstanceResponseBodyResultSpecSearchResource) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyResultTags struct {
	// The tag key.
	//
	// example:
	//
	// env
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// prod
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetInstanceResponseBodyResultTags) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyResultTags) GetKey() *string {
	return s.Key
}

func (s *GetInstanceResponseBodyResultTags) GetValue() *string {
	return s.Value
}

func (s *GetInstanceResponseBodyResultTags) SetKey(v string) *GetInstanceResponseBodyResultTags {
	s.Key = &v
	return s
}

func (s *GetInstanceResponseBodyResultTags) SetValue(v string) *GetInstanceResponseBodyResultTags {
	s.Value = &v
	return s
}

func (s *GetInstanceResponseBodyResultTags) Validate() error {
	return dara.Validate(s)
}
