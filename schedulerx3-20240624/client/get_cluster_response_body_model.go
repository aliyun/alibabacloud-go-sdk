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
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *GetClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D0DE9C33-992A-580B-89C4-B609A292748D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// xxljob-e0d018c6df8
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// xxl-job-test-1730427575152
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// scx.small.x2
	ClusterSpec *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// example:
	//
	// 2024-10-29 15:56:36
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2024-10-29 15:56:36
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// xxljob
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// example:
	//
	// 2.0.0
	EngineVersion  *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	InternetDomain *string `json:"InternetDomain,omitempty" xml:"InternetDomain,omitempty"`
	// example:
	//
	// http://xxljob-xxxxxx.schedulerx.mse.aliyuncs.com
	IntranetDomain *string `json:"IntranetDomain,omitempty" xml:"IntranetDomain,omitempty"`
	// example:
	//
	// 100
	JobNum *int32 `json:"JobNum,omitempty" xml:"JobNum,omitempty"`
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
	// example:
	//
	// 1000
	MaxJobNum *int32 `json:"MaxJobNum,omitempty" xml:"MaxJobNum,omitempty"`
	// example:
	//
	// 2
	ProductType *int32 `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// 10
	Spm *int32 `json:"Spm,omitempty" xml:"Spm,omitempty"`
	// example:
	//
	// 2
	Status           *int32                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags             map[string]interface{}                 `json:"Tags,omitempty" xml:"Tags,omitempty"`
	VSwitches        []*GetClusterResponseBodyDataVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	VersionLifecycle *string                                `json:"VersionLifecycle,omitempty" xml:"VersionLifecycle,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-bp1fiz967u39lt8yuxcs0
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// 10
	WorkerNum *int32    `json:"WorkerNum,omitempty" xml:"WorkerNum,omitempty"`
	Zones     []*string `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
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

func (s *GetClusterResponseBodyData) GetJobNum() *int32 {
	return s.JobNum
}

func (s *GetClusterResponseBodyData) GetKubeConfig() *string {
	return s.KubeConfig
}

func (s *GetClusterResponseBodyData) GetMaxJobNum() *int32 {
	return s.MaxJobNum
}

func (s *GetClusterResponseBodyData) GetProductType() *int32 {
	return s.ProductType
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

func (s *GetClusterResponseBodyData) SetProductType(v int32) *GetClusterResponseBodyData {
	s.ProductType = &v
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
	// example:
	//
	// vsw-8vbf1n216nshvfjdyff8a
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
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
