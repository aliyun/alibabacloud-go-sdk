// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAppResponseBody
	GetRequestId() *string
	SetResult(v *GetAppResponseBodyResult) *GetAppResponseBody
	GetResult() *GetAppResponseBodyResult
}

type GetAppResponseBody struct {
	// example:
	//
	// 2C5DAA30-****-5181-9B87-9D6181016197
	RequestId *string                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetAppResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppResponseBody) GetResult() *GetAppResponseBodyResult {
	return s.Result
}

func (s *GetAppResponseBody) SetRequestId(v string) *GetAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppResponseBody) SetResult(v *GetAppResponseBodyResult) *GetAppResponseBody {
	s.Result = v
	return s
}

func (s *GetAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAppResponseBodyResult struct {
	// example:
	//
	// test-app-abc
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// es-severless-test-app
	AppName    *string `json:"appName,omitempty" xml:"appName,omitempty"`
	AppType    *string `json:"appType,omitempty" xml:"appType,omitempty"`
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// example:
	//
	// 2022-08-15T11:20:52.370Z
	CreateTime  *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	InstanceId  *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 2022-08-15T11:21:50.000Z
	ModifiedTime *string                            `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	Network      []*GetAppResponseBodyResultNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Repeated"`
	// example:
	//
	// *******7595
	OwnerId        *string                                   `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	PrivateNetwork []*GetAppResponseBodyResultPrivateNetwork `json:"privateNetwork,omitempty" xml:"privateNetwork,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Scenario *string `json:"scenario,omitempty" xml:"scenario,omitempty"`
	// example:
	//
	// ACTIVE
	Status *string                         `json:"status,omitempty" xml:"status,omitempty"`
	Tags   []*GetAppResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// 7.10
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetAppResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyResult) GetAppId() *string {
	return s.AppId
}

func (s *GetAppResponseBodyResult) GetAppName() *string {
	return s.AppName
}

func (s *GetAppResponseBodyResult) GetAppType() *string {
	return s.AppType
}

func (s *GetAppResponseBodyResult) GetChargeType() *string {
	return s.ChargeType
}

func (s *GetAppResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetAppResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *GetAppResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAppResponseBodyResult) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetAppResponseBodyResult) GetNetwork() []*GetAppResponseBodyResultNetwork {
	return s.Network
}

func (s *GetAppResponseBodyResult) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetAppResponseBodyResult) GetPrivateNetwork() []*GetAppResponseBodyResultPrivateNetwork {
	return s.PrivateNetwork
}

func (s *GetAppResponseBodyResult) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAppResponseBodyResult) GetScenario() *string {
	return s.Scenario
}

func (s *GetAppResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetAppResponseBodyResult) GetTags() []*GetAppResponseBodyResultTags {
	return s.Tags
}

func (s *GetAppResponseBodyResult) GetVersion() *string {
	return s.Version
}

func (s *GetAppResponseBodyResult) SetAppId(v string) *GetAppResponseBodyResult {
	s.AppId = &v
	return s
}

func (s *GetAppResponseBodyResult) SetAppName(v string) *GetAppResponseBodyResult {
	s.AppName = &v
	return s
}

func (s *GetAppResponseBodyResult) SetAppType(v string) *GetAppResponseBodyResult {
	s.AppType = &v
	return s
}

func (s *GetAppResponseBodyResult) SetChargeType(v string) *GetAppResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *GetAppResponseBodyResult) SetCreateTime(v string) *GetAppResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetAppResponseBodyResult) SetDescription(v string) *GetAppResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetAppResponseBodyResult) SetInstanceId(v string) *GetAppResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *GetAppResponseBodyResult) SetModifiedTime(v string) *GetAppResponseBodyResult {
	s.ModifiedTime = &v
	return s
}

func (s *GetAppResponseBodyResult) SetNetwork(v []*GetAppResponseBodyResultNetwork) *GetAppResponseBodyResult {
	s.Network = v
	return s
}

func (s *GetAppResponseBodyResult) SetOwnerId(v string) *GetAppResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *GetAppResponseBodyResult) SetPrivateNetwork(v []*GetAppResponseBodyResultPrivateNetwork) *GetAppResponseBodyResult {
	s.PrivateNetwork = v
	return s
}

func (s *GetAppResponseBodyResult) SetRegionId(v string) *GetAppResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *GetAppResponseBodyResult) SetScenario(v string) *GetAppResponseBodyResult {
	s.Scenario = &v
	return s
}

func (s *GetAppResponseBodyResult) SetStatus(v string) *GetAppResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetAppResponseBodyResult) SetTags(v []*GetAppResponseBodyResultTags) *GetAppResponseBodyResult {
	s.Tags = v
	return s
}

func (s *GetAppResponseBodyResult) SetVersion(v string) *GetAppResponseBodyResult {
	s.Version = &v
	return s
}

func (s *GetAppResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetAppResponseBodyResultNetwork struct {
	Domain       *string                                        `json:"domain,omitempty" xml:"domain,omitempty"`
	Enabled      *bool                                          `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Port         *int32                                         `json:"port,omitempty" xml:"port,omitempty"`
	Type         *string                                        `json:"type,omitempty" xml:"type,omitempty"`
	WhiteIpGroup []*GetAppResponseBodyResultNetworkWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Repeated"`
}

func (s GetAppResponseBodyResultNetwork) String() string {
	return dara.Prettify(s)
}

func (s GetAppResponseBodyResultNetwork) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyResultNetwork) GetDomain() *string {
	return s.Domain
}

func (s *GetAppResponseBodyResultNetwork) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetAppResponseBodyResultNetwork) GetPort() *int32 {
	return s.Port
}

func (s *GetAppResponseBodyResultNetwork) GetType() *string {
	return s.Type
}

func (s *GetAppResponseBodyResultNetwork) GetWhiteIpGroup() []*GetAppResponseBodyResultNetworkWhiteIpGroup {
	return s.WhiteIpGroup
}

func (s *GetAppResponseBodyResultNetwork) SetDomain(v string) *GetAppResponseBodyResultNetwork {
	s.Domain = &v
	return s
}

func (s *GetAppResponseBodyResultNetwork) SetEnabled(v bool) *GetAppResponseBodyResultNetwork {
	s.Enabled = &v
	return s
}

func (s *GetAppResponseBodyResultNetwork) SetPort(v int32) *GetAppResponseBodyResultNetwork {
	s.Port = &v
	return s
}

func (s *GetAppResponseBodyResultNetwork) SetType(v string) *GetAppResponseBodyResultNetwork {
	s.Type = &v
	return s
}

func (s *GetAppResponseBodyResultNetwork) SetWhiteIpGroup(v []*GetAppResponseBodyResultNetworkWhiteIpGroup) *GetAppResponseBodyResultNetwork {
	s.WhiteIpGroup = v
	return s
}

func (s *GetAppResponseBodyResultNetwork) Validate() error {
	return dara.Validate(s)
}

type GetAppResponseBodyResultNetworkWhiteIpGroup struct {
	GroupName *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Ips       []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
}

func (s GetAppResponseBodyResultNetworkWhiteIpGroup) String() string {
	return dara.Prettify(s)
}

func (s GetAppResponseBodyResultNetworkWhiteIpGroup) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyResultNetworkWhiteIpGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *GetAppResponseBodyResultNetworkWhiteIpGroup) GetIps() []*string {
	return s.Ips
}

func (s *GetAppResponseBodyResultNetworkWhiteIpGroup) SetGroupName(v string) *GetAppResponseBodyResultNetworkWhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *GetAppResponseBodyResultNetworkWhiteIpGroup) SetIps(v []*string) *GetAppResponseBodyResultNetworkWhiteIpGroup {
	s.Ips = v
	return s
}

func (s *GetAppResponseBodyResultNetworkWhiteIpGroup) Validate() error {
	return dara.Validate(s)
}

type GetAppResponseBodyResultPrivateNetwork struct {
	Domain        *string                                               `json:"domain,omitempty" xml:"domain,omitempty"`
	Enabled       *bool                                                 `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Port          *int32                                                `json:"port,omitempty" xml:"port,omitempty"`
	PvlEndpointId *string                                               `json:"pvlEndpointId,omitempty" xml:"pvlEndpointId,omitempty"`
	Type          *string                                               `json:"type,omitempty" xml:"type,omitempty"`
	VpcId         *string                                               `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	WhiteIpGroup  []*GetAppResponseBodyResultPrivateNetworkWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Repeated"`
}

func (s GetAppResponseBodyResultPrivateNetwork) String() string {
	return dara.Prettify(s)
}

func (s GetAppResponseBodyResultPrivateNetwork) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyResultPrivateNetwork) GetDomain() *string {
	return s.Domain
}

func (s *GetAppResponseBodyResultPrivateNetwork) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetAppResponseBodyResultPrivateNetwork) GetPort() *int32 {
	return s.Port
}

func (s *GetAppResponseBodyResultPrivateNetwork) GetPvlEndpointId() *string {
	return s.PvlEndpointId
}

func (s *GetAppResponseBodyResultPrivateNetwork) GetType() *string {
	return s.Type
}

func (s *GetAppResponseBodyResultPrivateNetwork) GetVpcId() *string {
	return s.VpcId
}

func (s *GetAppResponseBodyResultPrivateNetwork) GetWhiteIpGroup() []*GetAppResponseBodyResultPrivateNetworkWhiteIpGroup {
	return s.WhiteIpGroup
}

func (s *GetAppResponseBodyResultPrivateNetwork) SetDomain(v string) *GetAppResponseBodyResultPrivateNetwork {
	s.Domain = &v
	return s
}

func (s *GetAppResponseBodyResultPrivateNetwork) SetEnabled(v bool) *GetAppResponseBodyResultPrivateNetwork {
	s.Enabled = &v
	return s
}

func (s *GetAppResponseBodyResultPrivateNetwork) SetPort(v int32) *GetAppResponseBodyResultPrivateNetwork {
	s.Port = &v
	return s
}

func (s *GetAppResponseBodyResultPrivateNetwork) SetPvlEndpointId(v string) *GetAppResponseBodyResultPrivateNetwork {
	s.PvlEndpointId = &v
	return s
}

func (s *GetAppResponseBodyResultPrivateNetwork) SetType(v string) *GetAppResponseBodyResultPrivateNetwork {
	s.Type = &v
	return s
}

func (s *GetAppResponseBodyResultPrivateNetwork) SetVpcId(v string) *GetAppResponseBodyResultPrivateNetwork {
	s.VpcId = &v
	return s
}

func (s *GetAppResponseBodyResultPrivateNetwork) SetWhiteIpGroup(v []*GetAppResponseBodyResultPrivateNetworkWhiteIpGroup) *GetAppResponseBodyResultPrivateNetwork {
	s.WhiteIpGroup = v
	return s
}

func (s *GetAppResponseBodyResultPrivateNetwork) Validate() error {
	return dara.Validate(s)
}

type GetAppResponseBodyResultPrivateNetworkWhiteIpGroup struct {
	GroupName *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Ips       []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
}

func (s GetAppResponseBodyResultPrivateNetworkWhiteIpGroup) String() string {
	return dara.Prettify(s)
}

func (s GetAppResponseBodyResultPrivateNetworkWhiteIpGroup) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyResultPrivateNetworkWhiteIpGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *GetAppResponseBodyResultPrivateNetworkWhiteIpGroup) GetIps() []*string {
	return s.Ips
}

func (s *GetAppResponseBodyResultPrivateNetworkWhiteIpGroup) SetGroupName(v string) *GetAppResponseBodyResultPrivateNetworkWhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *GetAppResponseBodyResultPrivateNetworkWhiteIpGroup) SetIps(v []*string) *GetAppResponseBodyResultPrivateNetworkWhiteIpGroup {
	s.Ips = v
	return s
}

func (s *GetAppResponseBodyResultPrivateNetworkWhiteIpGroup) Validate() error {
	return dara.Validate(s)
}

type GetAppResponseBodyResultTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetAppResponseBodyResultTags) String() string {
	return dara.Prettify(s)
}

func (s GetAppResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyResultTags) GetKey() *string {
	return s.Key
}

func (s *GetAppResponseBodyResultTags) GetValue() *string {
	return s.Value
}

func (s *GetAppResponseBodyResultTags) SetKey(v string) *GetAppResponseBodyResultTags {
	s.Key = &v
	return s
}

func (s *GetAppResponseBodyResultTags) SetValue(v string) *GetAppResponseBodyResultTags {
	s.Value = &v
	return s
}

func (s *GetAppResponseBodyResultTags) Validate() error {
	return dara.Validate(s)
}
