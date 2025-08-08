// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudPhoneNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeCloudPhoneNodesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeCloudPhoneNodesResponseBody
	GetNextToken() *string
	SetNodeModel(v []*DescribeCloudPhoneNodesResponseBodyNodeModel) *DescribeCloudPhoneNodesResponseBody
	GetNodeModel() []*DescribeCloudPhoneNodesResponseBodyNodeModel
	SetRequestId(v string) *DescribeCloudPhoneNodesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCloudPhoneNodesResponseBody
	GetTotalCount() *int32
}

type DescribeCloudPhoneNodesResponseBody struct {
	// The maximum number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- ****
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU+SQXzm0H9mu/FiSc****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The matrixes.
	NodeModel []*DescribeCloudPhoneNodesResponseBodyNodeModel `json:"NodeModel,omitempty" xml:"NodeModel,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F07A1DA1-E1EB-5CCA-8EED-12F85D32****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of cloud phone instances.
	//
	// example:
	//
	// 31
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCloudPhoneNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudPhoneNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeCloudPhoneNodesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCloudPhoneNodesResponseBody) GetNodeModel() []*DescribeCloudPhoneNodesResponseBodyNodeModel {
	return s.NodeModel
}

func (s *DescribeCloudPhoneNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudPhoneNodesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCloudPhoneNodesResponseBody) SetMaxResults(v int32) *DescribeCloudPhoneNodesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBody) SetNextToken(v string) *DescribeCloudPhoneNodesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBody) SetNodeModel(v []*DescribeCloudPhoneNodesResponseBodyNodeModel) *DescribeCloudPhoneNodesResponseBody {
	s.NodeModel = v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBody) SetRequestId(v string) *DescribeCloudPhoneNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBody) SetTotalCount(v int32) *DescribeCloudPhoneNodesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudPhoneNodesResponseBodyNodeModel struct {
	BandwidthPackageId     *string                                                `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	BandwidthPackageStatus *string                                                `json:"BandwidthPackageStatus,omitempty" xml:"BandwidthPackageStatus,omitempty"`
	BandwidthPackageType   *string                                                `json:"BandwidthPackageType,omitempty" xml:"BandwidthPackageType,omitempty"`
	BizTags                []*DescribeCloudPhoneNodesResponseBodyNodeModelBizTags `json:"BizTags,omitempty" xml:"BizTags,omitempty" type:"Repeated"`
	// The billing method.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The number of CPU cores.
	//
	// example:
	//
	// 2
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-11-13 02:03:14
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The expiration time of the subscription matrix.
	//
	// example:
	//
	// 2025-03-09 02:00:34
	GmtExpired *string `json:"GmtExpired,omitempty" xml:"GmtExpired,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 2025-02-13 02:03:14
	GmtModified  *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The memory size. Unit: GB.
	//
	// example:
	//
	// 32
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-5mwr9azebliva****
	NetworkId    *string                                                     `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	NetworkInfos []*DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos `json:"NetworkInfos,omitempty" xml:"NetworkInfos,omitempty" type:"Repeated"`
	NetworkType  *string                                                     `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The matrix ID.
	//
	// example:
	//
	// cpn-ehs0yoedq8ntm****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The matrix name.
	//
	// example:
	//
	// node_name
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The number of cloud phone instances per matrix.
	//
	// example:
	//
	// 25
	PhoneCount    *int32                                                     `json:"PhoneCount,omitempty" xml:"PhoneCount,omitempty"`
	PhoneDataInfo *DescribeCloudPhoneNodesResponseBodyNodeModelPhoneDataInfo `json:"PhoneDataInfo,omitempty" xml:"PhoneDataInfo,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The height of the resolution. Unit: pixel.
	//
	// example:
	//
	// 1280
	ResolutionHeight *int32 `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	// The width of the resolution. Unit: pixel.
	//
	// example:
	//
	// 720
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
	// The matrix specification.
	//
	// example:
	//
	// cpm.gn6.gx1
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// The size of the shared storage. Unit: GiB.
	//
	// example:
	//
	// 100
	ShareDataVolume *int32 `json:"ShareDataVolume,omitempty" xml:"ShareDataVolume,omitempty"`
	// The matrix status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-2zeekryyc1q3sm72l****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeCloudPhoneNodesResponseBodyNodeModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudPhoneNodesResponseBodyNodeModel) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) GetBandwidthPackageStatus() *string {
	return s.BandwidthPackageStatus
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) GetBandwidthPackageType() *string {
	return s.BandwidthPackageType
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) GetBizTags() []*DescribeCloudPhoneNodesResponseBodyNodeModelBizTags {
	return s.BizTags
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) GetCpu() *string {
	return s.Cpu
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) GetGmtExpired() *string {
	return s.GmtExpired
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) GetMemory() *int32 {
	return s.Memory
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) GetNetworkInfos() []*DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos {
	return s.NetworkInfos
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) GetPhoneCount() *int32 {
	return s.PhoneCount
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) GetPhoneDataInfo() *DescribeCloudPhoneNodesResponseBodyNodeModelPhoneDataInfo {
	return s.PhoneDataInfo
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) GetResolutionHeight() *int32 {
	return s.ResolutionHeight
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) GetResolutionWidth() *int32 {
	return s.ResolutionWidth
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) GetServerType() *string {
	return s.ServerType
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) GetShareDataVolume() *int32 {
	return s.ShareDataVolume
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) GetStatus() *string {
	return s.Status
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetBandwidthPackageId(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.BandwidthPackageId = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetBandwidthPackageStatus(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.BandwidthPackageStatus = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetBandwidthPackageType(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.BandwidthPackageType = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetBizTags(v []*DescribeCloudPhoneNodesResponseBodyNodeModelBizTags) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.BizTags = v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetChargeType(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.ChargeType = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetCpu(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.Cpu = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetGmtCreate(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.GmtCreate = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetGmtExpired(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.GmtExpired = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetGmtModified(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.GmtModified = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetInstanceType(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.InstanceType = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetMemory(v int32) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.Memory = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetNetworkId(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.NetworkId = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetNetworkInfos(v []*DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.NetworkInfos = v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetNetworkType(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.NetworkType = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetNodeId(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.NodeId = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetNodeName(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.NodeName = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetPhoneCount(v int32) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.PhoneCount = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetPhoneDataInfo(v *DescribeCloudPhoneNodesResponseBodyNodeModelPhoneDataInfo) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.PhoneDataInfo = v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetRegionId(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetResolutionHeight(v int32) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.ResolutionHeight = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetResolutionWidth(v int32) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.ResolutionWidth = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetServerType(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.ServerType = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetShareDataVolume(v int32) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.ShareDataVolume = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetStatus(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.Status = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) SetVSwitchId(v string) *DescribeCloudPhoneNodesResponseBodyNodeModel {
	s.VSwitchId = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModel) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudPhoneNodesResponseBodyNodeModelBizTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCloudPhoneNodesResponseBodyNodeModelBizTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudPhoneNodesResponseBodyNodeModelBizTags) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModelBizTags) GetKey() *string {
	return s.Key
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModelBizTags) GetValue() *string {
	return s.Value
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModelBizTags) SetKey(v string) *DescribeCloudPhoneNodesResponseBodyNodeModelBizTags {
	s.Key = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModelBizTags) SetValue(v string) *DescribeCloudPhoneNodesResponseBodyNodeModelBizTags {
	s.Value = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModelBizTags) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos struct {
	BandwidthPackageId   *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	BandwidthPackageType *string `json:"BandwidthPackageType,omitempty" xml:"BandwidthPackageType,omitempty"`
	NetworkId            *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	NetworkType          *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos) GetBandwidthPackageType() *string {
	return s.BandwidthPackageType
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos) SetBandwidthPackageId(v string) *DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos {
	s.BandwidthPackageId = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos) SetBandwidthPackageType(v string) *DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos {
	s.BandwidthPackageType = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos) SetNetworkId(v string) *DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos {
	s.NetworkId = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos) SetNetworkType(v string) *DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos {
	s.NetworkType = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos) SetVSwitchId(v string) *DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos {
	s.VSwitchId = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModelNetworkInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudPhoneNodesResponseBodyNodeModelPhoneDataInfo struct {
	PhoneDataId     *string `json:"PhoneDataId,omitempty" xml:"PhoneDataId,omitempty"`
	PhoneDataVolume *int32  `json:"PhoneDataVolume,omitempty" xml:"PhoneDataVolume,omitempty"`
}

func (s DescribeCloudPhoneNodesResponseBodyNodeModelPhoneDataInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudPhoneNodesResponseBodyNodeModelPhoneDataInfo) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModelPhoneDataInfo) GetPhoneDataId() *string {
	return s.PhoneDataId
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModelPhoneDataInfo) GetPhoneDataVolume() *int32 {
	return s.PhoneDataVolume
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModelPhoneDataInfo) SetPhoneDataId(v string) *DescribeCloudPhoneNodesResponseBodyNodeModelPhoneDataInfo {
	s.PhoneDataId = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModelPhoneDataInfo) SetPhoneDataVolume(v int32) *DescribeCloudPhoneNodesResponseBodyNodeModelPhoneDataInfo {
	s.PhoneDataVolume = &v
	return s
}

func (s *DescribeCloudPhoneNodesResponseBodyNodeModelPhoneDataInfo) Validate() error {
	return dara.Validate(s)
}
