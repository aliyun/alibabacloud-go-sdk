// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeARMServerInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeARMServerInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeARMServerInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeARMServerInstancesResponseBody
	GetRequestId() *string
	SetServers(v []*DescribeARMServerInstancesResponseBodyServers) *DescribeARMServerInstancesResponseBody
	GetServers() []*DescribeARMServerInstancesResponseBodyServers
	SetTotalCount(v int32) *DescribeARMServerInstancesResponseBody
	GetTotalCount() *int32
}

type DescribeARMServerInstancesResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the servers and the AIC instances.
	Servers []*DescribeARMServerInstancesResponseBodyServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeARMServerInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeARMServerInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeARMServerInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeARMServerInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeARMServerInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeARMServerInstancesResponseBody) GetServers() []*DescribeARMServerInstancesResponseBodyServers {
	return s.Servers
}

func (s *DescribeARMServerInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeARMServerInstancesResponseBody) SetPageNumber(v int32) *DescribeARMServerInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBody) SetPageSize(v int32) *DescribeARMServerInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBody) SetRequestId(v string) *DescribeARMServerInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBody) SetServers(v []*DescribeARMServerInstancesResponseBodyServers) *DescribeARMServerInstancesResponseBody {
	s.Servers = v
	return s
}

func (s *DescribeARMServerInstancesResponseBody) SetTotalCount(v int32) *DescribeARMServerInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeARMServerInstancesResponseBodyServers struct {
	// The information about the AIC instances.
	AICInstances []*DescribeARMServerInstancesResponseBodyServersAICInstances `json:"AICInstances,omitempty" xml:"AICInstances,omitempty" type:"Repeated"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2022-05-07 11:59:09
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the ENS node.
	//
	// example:
	//
	// cn-hanghzou-27
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The time when the instance expires.
	//
	// example:
	//
	// 2027-05-07 11:59:09
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The operation that was most recently performed.
	//
	// example:
	//
	// ServerCreate
	LatestAction *string `json:"LatestAction,omitempty" xml:"LatestAction,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// Server-Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace of the cluster to which the server belongs.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The billing method.
	//
	// example:
	//
	// prepay
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The ID of the server.
	//
	// example:
	//
	// cas-******
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The server specification.
	//
	// example:
	//
	// cas.cf53r
	SpecName *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	// The operation status of the server. Valid values:
	//
	// 	- **success**
	//
	// 	- **failed**
	//
	// 	- **creating**
	//
	// 	- **releasing**
	//
	// 	- **rebooting**
	//
	// 	- **upgrading**
	//
	// example:
	//
	// success
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The running status of the server. Valid values:
	//
	// 	- **running**
	//
	// 	- **stopping**
	//
	// 	- **down**
	//
	// 	- **starting**
	//
	// example:
	//
	// running
	Status *string                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*DescribeARMServerInstancesResponseBodyServersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeARMServerInstancesResponseBodyServers) String() string {
	return dara.Prettify(s)
}

func (s DescribeARMServerInstancesResponseBodyServers) GoString() string {
	return s.String()
}

func (s *DescribeARMServerInstancesResponseBodyServers) GetAICInstances() []*DescribeARMServerInstancesResponseBodyServersAICInstances {
	return s.AICInstances
}

func (s *DescribeARMServerInstancesResponseBodyServers) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeARMServerInstancesResponseBodyServers) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeARMServerInstancesResponseBodyServers) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeARMServerInstancesResponseBodyServers) GetLatestAction() *string {
	return s.LatestAction
}

func (s *DescribeARMServerInstancesResponseBodyServers) GetName() *string {
	return s.Name
}

func (s *DescribeARMServerInstancesResponseBodyServers) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeARMServerInstancesResponseBodyServers) GetPayType() *string {
	return s.PayType
}

func (s *DescribeARMServerInstancesResponseBodyServers) GetServerId() *string {
	return s.ServerId
}

func (s *DescribeARMServerInstancesResponseBodyServers) GetSpecName() *string {
	return s.SpecName
}

func (s *DescribeARMServerInstancesResponseBodyServers) GetState() *string {
	return s.State
}

func (s *DescribeARMServerInstancesResponseBodyServers) GetStatus() *string {
	return s.Status
}

func (s *DescribeARMServerInstancesResponseBodyServers) GetTags() []*DescribeARMServerInstancesResponseBodyServersTags {
	return s.Tags
}

func (s *DescribeARMServerInstancesResponseBodyServers) SetAICInstances(v []*DescribeARMServerInstancesResponseBodyServersAICInstances) *DescribeARMServerInstancesResponseBodyServers {
	s.AICInstances = v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServers) SetCreationTime(v string) *DescribeARMServerInstancesResponseBodyServers {
	s.CreationTime = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServers) SetEnsRegionId(v string) *DescribeARMServerInstancesResponseBodyServers {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServers) SetExpiredTime(v string) *DescribeARMServerInstancesResponseBodyServers {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServers) SetLatestAction(v string) *DescribeARMServerInstancesResponseBodyServers {
	s.LatestAction = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServers) SetName(v string) *DescribeARMServerInstancesResponseBodyServers {
	s.Name = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServers) SetNamespace(v string) *DescribeARMServerInstancesResponseBodyServers {
	s.Namespace = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServers) SetPayType(v string) *DescribeARMServerInstancesResponseBodyServers {
	s.PayType = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServers) SetServerId(v string) *DescribeARMServerInstancesResponseBodyServers {
	s.ServerId = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServers) SetSpecName(v string) *DescribeARMServerInstancesResponseBodyServers {
	s.SpecName = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServers) SetState(v string) *DescribeARMServerInstancesResponseBodyServers {
	s.State = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServers) SetStatus(v string) *DescribeARMServerInstancesResponseBodyServers {
	s.Status = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServers) SetTags(v []*DescribeARMServerInstancesResponseBodyServersTags) *DescribeARMServerInstancesResponseBodyServers {
	s.Tags = v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServers) Validate() error {
	return dara.Validate(s)
}

type DescribeARMServerInstancesResponseBodyServersAICInstances struct {
	// The refresh rate of the AIC instance. Unit: Hz.
	//
	// example:
	//
	// 120
	Frequency *int64 `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The ID of the AIC image.
	//
	// example:
	//
	// m-****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the AIC instance.
	//
	// example:
	//
	// aic-instance****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The operation that was most recently performed.
	//
	// example:
	//
	// ServerCreate
	LatestAction *string `json:"LatestAction,omitempty" xml:"LatestAction,omitempty"`
	// The name of the container.
	//
	// example:
	//
	// AIC-Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network attributes of the AIC instance.
	NetworkAttributes *DescribeARMServerInstancesResponseBodyServersAICInstancesNetworkAttributes `json:"NetworkAttributes,omitempty" xml:"NetworkAttributes,omitempty" type:"Struct"`
	// The resolution of the AIC instance.
	//
	// example:
	//
	// 1920*1080
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// The information about the shared data group (SDG) that is deployed on the AIC instance.
	SdgDeployInfo *DescribeARMServerInstancesResponseBodyServersAICInstancesSdgDeployInfo `json:"SdgDeployInfo,omitempty" xml:"SdgDeployInfo,omitempty" type:"Struct"`
	// The specification of the AIC instance.
	//
	// example:
	//
	// aic.cf53r.c2.np
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The operation status of the AIC instance. Valid values:
	//
	// 	- **success**
	//
	// 	- **failed**
	//
	// 	- **creating**
	//
	// 	- **releasing**
	//
	// 	- **rebooting**
	//
	// 	- **reseting**
	//
	// example:
	//
	// success
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The running status of the AIC instance. Valid values:
	//
	// 	- **running**
	//
	// 	- **pending**
	//
	// 	- **terminating**
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeARMServerInstancesResponseBodyServersAICInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeARMServerInstancesResponseBodyServersAICInstances) GoString() string {
	return s.String()
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstances) GetFrequency() *int64 {
	return s.Frequency
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstances) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstances) GetLatestAction() *string {
	return s.LatestAction
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstances) GetName() *string {
	return s.Name
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstances) GetNetworkAttributes() *DescribeARMServerInstancesResponseBodyServersAICInstancesNetworkAttributes {
	return s.NetworkAttributes
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstances) GetResolution() *string {
	return s.Resolution
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstances) GetSdgDeployInfo() *DescribeARMServerInstancesResponseBodyServersAICInstancesSdgDeployInfo {
	return s.SdgDeployInfo
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstances) GetSpec() *string {
	return s.Spec
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstances) GetState() *string {
	return s.State
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstances) GetStatus() *string {
	return s.Status
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstances) SetFrequency(v int64) *DescribeARMServerInstancesResponseBodyServersAICInstances {
	s.Frequency = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstances) SetImageId(v string) *DescribeARMServerInstancesResponseBodyServersAICInstances {
	s.ImageId = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstances) SetInstanceId(v string) *DescribeARMServerInstancesResponseBodyServersAICInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstances) SetLatestAction(v string) *DescribeARMServerInstancesResponseBodyServersAICInstances {
	s.LatestAction = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstances) SetName(v string) *DescribeARMServerInstancesResponseBodyServersAICInstances {
	s.Name = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstances) SetNetworkAttributes(v *DescribeARMServerInstancesResponseBodyServersAICInstancesNetworkAttributes) *DescribeARMServerInstancesResponseBodyServersAICInstances {
	s.NetworkAttributes = v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstances) SetResolution(v string) *DescribeARMServerInstancesResponseBodyServersAICInstances {
	s.Resolution = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstances) SetSdgDeployInfo(v *DescribeARMServerInstancesResponseBodyServersAICInstancesSdgDeployInfo) *DescribeARMServerInstancesResponseBodyServersAICInstances {
	s.SdgDeployInfo = v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstances) SetSpec(v string) *DescribeARMServerInstancesResponseBodyServersAICInstances {
	s.Spec = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstances) SetState(v string) *DescribeARMServerInstancesResponseBodyServersAICInstances {
	s.State = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstances) SetStatus(v string) *DescribeARMServerInstancesResponseBodyServersAICInstances {
	s.Status = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeARMServerInstancesResponseBodyServersAICInstancesNetworkAttributes struct {
	// The IP address of the AIC instance.
	//
	// example:
	//
	// 192.168.XX.XX
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The network ID of the AIC instance.
	//
	// example:
	//
	// n-*****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The vSwitch ID of the AIC instance.
	//
	// example:
	//
	// vsw-****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeARMServerInstancesResponseBodyServersAICInstancesNetworkAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeARMServerInstancesResponseBodyServersAICInstancesNetworkAttributes) GoString() string {
	return s.String()
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstancesNetworkAttributes) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstancesNetworkAttributes) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstancesNetworkAttributes) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstancesNetworkAttributes) SetIpAddress(v string) *DescribeARMServerInstancesResponseBodyServersAICInstancesNetworkAttributes {
	s.IpAddress = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstancesNetworkAttributes) SetNetworkId(v string) *DescribeARMServerInstancesResponseBodyServersAICInstancesNetworkAttributes {
	s.NetworkId = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstancesNetworkAttributes) SetVSwitchId(v string) *DescribeARMServerInstancesResponseBodyServersAICInstancesNetworkAttributes {
	s.VSwitchId = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstancesNetworkAttributes) Validate() error {
	return dara.Validate(s)
}

type DescribeARMServerInstancesResponseBodyServersAICInstancesSdgDeployInfo struct {
	// The ID of the SDG.
	//
	// example:
	//
	// sdg-xxxxx
	SDGId *string `json:"SDGId,omitempty" xml:"SDGId,omitempty"`
	// The deployment status of the SDG. Valid values:
	//
	// 	- **sdg_deploying**
	//
	// 	- **failed**
	//
	// 	- **success**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeARMServerInstancesResponseBodyServersAICInstancesSdgDeployInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeARMServerInstancesResponseBodyServersAICInstancesSdgDeployInfo) GoString() string {
	return s.String()
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstancesSdgDeployInfo) GetSDGId() *string {
	return s.SDGId
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstancesSdgDeployInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstancesSdgDeployInfo) SetSDGId(v string) *DescribeARMServerInstancesResponseBodyServersAICInstancesSdgDeployInfo {
	s.SDGId = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstancesSdgDeployInfo) SetStatus(v string) *DescribeARMServerInstancesResponseBodyServersAICInstancesSdgDeployInfo {
	s.Status = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServersAICInstancesSdgDeployInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeARMServerInstancesResponseBodyServersTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeARMServerInstancesResponseBodyServersTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeARMServerInstancesResponseBodyServersTags) GoString() string {
	return s.String()
}

func (s *DescribeARMServerInstancesResponseBodyServersTags) GetKey() *string {
	return s.Key
}

func (s *DescribeARMServerInstancesResponseBodyServersTags) GetValue() *string {
	return s.Value
}

func (s *DescribeARMServerInstancesResponseBodyServersTags) SetKey(v string) *DescribeARMServerInstancesResponseBodyServersTags {
	s.Key = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServersTags) SetValue(v string) *DescribeARMServerInstancesResponseBodyServersTags {
	s.Value = &v
	return s
}

func (s *DescribeARMServerInstancesResponseBodyServersTags) Validate() error {
	return dara.Validate(s)
}
