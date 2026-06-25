// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetClusterResponseBody
	GetCode() *int32
	SetData(v *GetClusterResponseBodyData) *GetClusterResponseBody
	GetData() *GetClusterResponseBodyData
	SetMessage(v string) *GetClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetClusterResponseBody
	GetSuccess() *bool
}

type GetClusterResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID. The value is a unique identifier that Alibaba Cloud generates for the request and can be used to troubleshoot issues.
	//
	// example:
	//
	// D0DE9C33-992A-580B-89C4-B609A292748D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetClusterResponseBody) GetData() *GetClusterResponseBodyData {
	return s.Data
}

func (s *GetClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetClusterResponseBody) SetCode(v int32) *GetClusterResponseBody {
	s.Code = &v
	return s
}

func (s *GetClusterResponseBody) SetData(v *GetClusterResponseBodyData) *GetClusterResponseBody {
	s.Data = v
	return s
}

func (s *GetClusterResponseBody) SetMessage(v string) *GetClusterResponseBody {
	s.Message = &v
	return s
}

func (s *GetClusterResponseBody) SetRequestId(v string) *GetClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterResponseBody) SetSuccess(v bool) *GetClusterResponseBody {
	s.Success = &v
	return s
}

func (s *GetClusterResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetClusterResponseBodyData struct {
	// The billing method. Valid values:
	//
	// - PREPAY: subscription.
	//
	// - POSTPAY: pay-as-you-go.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// xxljob-e0d018c6df8
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// xxl-job-test-1730427575152
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The cluster specification. Valid values:
	//
	// - scx.dev.x1.
	//
	// - scx.small.x1.
	//
	// - scx.small.x2.
	//
	// - scx.medium.x1.
	//
	// - scx.medium.x2.
	//
	// example:
	//
	// scx.small.x2
	ClusterSpec *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// example:
	//
	// 1
	ClusterType *int32 `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The time when the cluster was created.
	//
	// example:
	//
	// 2024-10-29 15:56:36
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2024-10-29 15:56:36
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The engine type. Valid value: xxljob.
	//
	// example:
	//
	// xxljob
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The engine version.
	//
	// example:
	//
	// 2.0.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The public endpoint.
	//
	// > Currently not supported.
	//
	// example:
	//
	// 暂无
	InternetDomain *string `json:"InternetDomain,omitempty" xml:"InternetDomain,omitempty"`
	// The internal endpoint.
	//
	// example:
	//
	// http://xxljob-xxxxxx.schedulerx.mse.aliyuncs.com
	IntranetDomain *string `json:"IntranetDomain,omitempty" xml:"IntranetDomain,omitempty"`
	// example:
	//
	// 192.168.1.0/24
	IpWhitelist *string `json:"IpWhitelist,omitempty" xml:"IpWhitelist,omitempty"`
	// The maximum number of jobs for the current specification.
	//
	// example:
	//
	// 100
	JobNum *int32 `json:"JobNum,omitempty" xml:"JobNum,omitempty"`
	// The configuration of the Kubernetes server.
	//
	// example:
	//
	// {
	//
	//   "vSwitchIdList": [
	//
	//     "xxx",
	//
	//     "xxx"
	//
	//   ],
	//
	//   "cpu": xxx,
	//
	//   "cpuUnit": "xxx",
	//
	//   "diskCapacity": xxx,
	//
	//   "memoryCapacity": xxx,
	//
	//   "zoneIds": [
	//
	//     "xxx",
	//
	//     "xxx"
	//
	//   ],
	//
	//   "securityGroupList": [
	//
	//     "xxx"
	//
	//   ],
	//
	//   "eniCrossZone": "xxx",
	//
	//   "regionId": "xxx",
	//
	//   "instanceCount": xxx,
	//
	//   "vpcId": "xxx",
	//
	//   "memoryUnit": "xxx",
	//
	//   "diskType": "xxx",
	//
	//   "appClusterId": "xxx"
	//
	// }
	KubeConfig *string `json:"KubeConfig,omitempty" xml:"KubeConfig,omitempty"`
	// The maximum number of jobs for the current specification.
	//
	// example:
	//
	// 1000
	MaxJobNum *int32 `json:"MaxJobNum,omitempty" xml:"MaxJobNum,omitempty"`
	// The maximum number of workflows supported.
	//
	// example:
	//
	// 100
	MaxWorkflowNum *int32 `json:"MaxWorkflowNum,omitempty" xml:"MaxWorkflowNum,omitempty"`
	// The product edition.
	//
	// - 1: Developer Edition.
	//
	// - 2: Professional Edition.
	//
	// - 3: Enterprise Edition.
	//
	// example:
	//
	// 2
	ProductType *int32  `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	Source      *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The peak number of scheduling operations per minute.
	//
	// example:
	//
	// 10
	Spm *int32 `json:"Spm,omitempty" xml:"Spm,omitempty"`
	// The cluster status.
	//
	// - 1: Being created.
	//
	// - 2: Running.
	//
	// - 3: Restarting.
	//
	// - 4: Being released.
	//
	// - 5: Creation failed.
	//
	// - 6: Stopped.
	//
	// - 99: Deleted.
	//
	// example:
	//
	// 2
	Status *int32                 `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The vSwitch information of zones.
	VSwitches        []*GetClusterResponseBodyDataVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	VersionLifecycle *string                                `json:"VersionLifecycle,omitempty" xml:"VersionLifecycle,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp1fiz967u39lt8yuxcs0
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The number of workers.
	//
	// example:
	//
	// 10
	WorkerNum *int32 `json:"WorkerNum,omitempty" xml:"WorkerNum,omitempty"`
	// The current number of workflows.
	//
	// example:
	//
	// 20
	WorkflowNum *int32 `json:"WorkflowNum,omitempty" xml:"WorkflowNum,omitempty"`
	// The list of zones.
	Zones []*string `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s GetClusterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyData) GetChargeType() *string {
	return s.ChargeType
}

func (s *GetClusterResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetClusterResponseBodyData) GetClusterName() *string {
	return s.ClusterName
}

func (s *GetClusterResponseBodyData) GetClusterSpec() *string {
	return s.ClusterSpec
}

func (s *GetClusterResponseBodyData) GetClusterType() *int32 {
	return s.ClusterType
}

func (s *GetClusterResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetClusterResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *GetClusterResponseBodyData) GetEngineType() *string {
	return s.EngineType
}

func (s *GetClusterResponseBodyData) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *GetClusterResponseBodyData) GetInternetDomain() *string {
	return s.InternetDomain
}

func (s *GetClusterResponseBodyData) GetIntranetDomain() *string {
	return s.IntranetDomain
}

func (s *GetClusterResponseBodyData) GetIpWhitelist() *string {
	return s.IpWhitelist
}

func (s *GetClusterResponseBodyData) GetJobNum() *int32 {
	return s.JobNum
}

func (s *GetClusterResponseBodyData) GetKubeConfig() *string {
	return s.KubeConfig
}

func (s *GetClusterResponseBodyData) GetMaxJobNum() *int32 {
	return s.MaxJobNum
}

func (s *GetClusterResponseBodyData) GetMaxWorkflowNum() *int32 {
	return s.MaxWorkflowNum
}

func (s *GetClusterResponseBodyData) GetProductType() *int32 {
	return s.ProductType
}

func (s *GetClusterResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *GetClusterResponseBodyData) GetSpm() *int32 {
	return s.Spm
}

func (s *GetClusterResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetClusterResponseBodyData) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *GetClusterResponseBodyData) GetVSwitches() []*GetClusterResponseBodyDataVSwitches {
	return s.VSwitches
}

func (s *GetClusterResponseBodyData) GetVersionLifecycle() *string {
	return s.VersionLifecycle
}

func (s *GetClusterResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *GetClusterResponseBodyData) GetWorkerNum() *int32 {
	return s.WorkerNum
}

func (s *GetClusterResponseBodyData) GetWorkflowNum() *int32 {
	return s.WorkflowNum
}

func (s *GetClusterResponseBodyData) GetZones() []*string {
	return s.Zones
}

func (s *GetClusterResponseBodyData) SetChargeType(v string) *GetClusterResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *GetClusterResponseBodyData) SetClusterId(v string) *GetClusterResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *GetClusterResponseBodyData) SetClusterName(v string) *GetClusterResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *GetClusterResponseBodyData) SetClusterSpec(v string) *GetClusterResponseBodyData {
	s.ClusterSpec = &v
	return s
}

func (s *GetClusterResponseBodyData) SetClusterType(v int32) *GetClusterResponseBodyData {
	s.ClusterType = &v
	return s
}

func (s *GetClusterResponseBodyData) SetCreateTime(v string) *GetClusterResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetClusterResponseBodyData) SetEndTime(v string) *GetClusterResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetClusterResponseBodyData) SetEngineType(v string) *GetClusterResponseBodyData {
	s.EngineType = &v
	return s
}

func (s *GetClusterResponseBodyData) SetEngineVersion(v string) *GetClusterResponseBodyData {
	s.EngineVersion = &v
	return s
}

func (s *GetClusterResponseBodyData) SetInternetDomain(v string) *GetClusterResponseBodyData {
	s.InternetDomain = &v
	return s
}

func (s *GetClusterResponseBodyData) SetIntranetDomain(v string) *GetClusterResponseBodyData {
	s.IntranetDomain = &v
	return s
}

func (s *GetClusterResponseBodyData) SetIpWhitelist(v string) *GetClusterResponseBodyData {
	s.IpWhitelist = &v
	return s
}

func (s *GetClusterResponseBodyData) SetJobNum(v int32) *GetClusterResponseBodyData {
	s.JobNum = &v
	return s
}

func (s *GetClusterResponseBodyData) SetKubeConfig(v string) *GetClusterResponseBodyData {
	s.KubeConfig = &v
	return s
}

func (s *GetClusterResponseBodyData) SetMaxJobNum(v int32) *GetClusterResponseBodyData {
	s.MaxJobNum = &v
	return s
}

func (s *GetClusterResponseBodyData) SetMaxWorkflowNum(v int32) *GetClusterResponseBodyData {
	s.MaxWorkflowNum = &v
	return s
}

func (s *GetClusterResponseBodyData) SetProductType(v int32) *GetClusterResponseBodyData {
	s.ProductType = &v
	return s
}

func (s *GetClusterResponseBodyData) SetSource(v string) *GetClusterResponseBodyData {
	s.Source = &v
	return s
}

func (s *GetClusterResponseBodyData) SetSpm(v int32) *GetClusterResponseBodyData {
	s.Spm = &v
	return s
}

func (s *GetClusterResponseBodyData) SetStatus(v int32) *GetClusterResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetClusterResponseBodyData) SetTags(v map[string]interface{}) *GetClusterResponseBodyData {
	s.Tags = v
	return s
}

func (s *GetClusterResponseBodyData) SetVSwitches(v []*GetClusterResponseBodyDataVSwitches) *GetClusterResponseBodyData {
	s.VSwitches = v
	return s
}

func (s *GetClusterResponseBodyData) SetVersionLifecycle(v string) *GetClusterResponseBodyData {
	s.VersionLifecycle = &v
	return s
}

func (s *GetClusterResponseBodyData) SetVpcId(v string) *GetClusterResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *GetClusterResponseBodyData) SetWorkerNum(v int32) *GetClusterResponseBodyData {
	s.WorkerNum = &v
	return s
}

func (s *GetClusterResponseBodyData) SetWorkflowNum(v int32) *GetClusterResponseBodyData {
	s.WorkflowNum = &v
	return s
}

func (s *GetClusterResponseBodyData) SetZones(v []*string) *GetClusterResponseBodyData {
	s.Zones = v
	return s
}

func (s *GetClusterResponseBodyData) Validate() error {
	if s.VSwitches != nil {
		for _, item := range s.VSwitches {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetClusterResponseBodyDataVSwitches struct {
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-8vbf1n216nshvfjdyff8a
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetClusterResponseBodyDataVSwitches) String() string {
	return dara.Prettify(s)
}

func (s GetClusterResponseBodyDataVSwitches) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyDataVSwitches) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetClusterResponseBodyDataVSwitches) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetClusterResponseBodyDataVSwitches) SetVSwitchId(v string) *GetClusterResponseBodyDataVSwitches {
	s.VSwitchId = &v
	return s
}

func (s *GetClusterResponseBodyDataVSwitches) SetZoneId(v string) *GetClusterResponseBodyDataVSwitches {
	s.ZoneId = &v
	return s
}

func (s *GetClusterResponseBodyDataVSwitches) Validate() error {
	return dara.Validate(s)
}
