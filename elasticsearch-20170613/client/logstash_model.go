// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogstash interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v map[string]*string) *Logstash
	GetConfig() map[string]*string
	SetCreatedAt(v string) *Logstash
	GetCreatedAt() *string
	SetDataNode(v bool) *Logstash
	GetDataNode() *bool
	SetDescription(v string) *Logstash
	GetDescription() *string
	SetEndTime(v int64) *Logstash
	GetEndTime() *int64
	SetEndpointList(v []*LogstashEndpointList) *Logstash
	GetEndpointList() []*LogstashEndpointList
	SetInstanceId(v string) *Logstash
	GetInstanceId() *string
	SetNetworkConfig(v *LogstashNetworkConfig) *Logstash
	GetNetworkConfig() *LogstashNetworkConfig
	SetNodeAmount(v int64) *Logstash
	GetNodeAmount() *int64
	SetNodeSpec(v *LogstashNodeSpec) *Logstash
	GetNodeSpec() *LogstashNodeSpec
	SetPaymentType(v string) *Logstash
	GetPaymentType() *string
	SetProtocol(v string) *Logstash
	GetProtocol() *string
	SetResourceGroupId(v string) *Logstash
	GetResourceGroupId() *string
	SetStatus(v string) *Logstash
	GetStatus() *string
	SetTags(v []*LogstashTags) *Logstash
	GetTags() []*LogstashTags
	SetUpdatedAt(v string) *Logstash
	GetUpdatedAt() *string
	SetVersion(v string) *Logstash
	GetVersion() *string
	SetZoneCount(v int64) *Logstash
	GetZoneCount() *int64
	SetZoneInfos(v []*LogstashZoneInfos) *Logstash
	GetZoneInfos() []*LogstashZoneInfos
}

type Logstash struct {
	Config map[string]*string `json:"config,omitempty" xml:"config,omitempty"`
	// example:
	//
	// 2018-07-13T03:58:07.253Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// true
	DataNode *bool `json:"dataNode,omitempty" xml:"dataNode,omitempty"`
	// example:
	//
	// ls-cn-abc
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 4749897600000
	EndTime      *int64                  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	EndpointList []*LogstashEndpointList `json:"endpointList,omitempty" xml:"endpointList,omitempty" type:"Repeated"`
	// example:
	//
	// ls-cn-n6w1o5jq****
	InstanceId    *string                `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	NetworkConfig *LogstashNetworkConfig `json:"networkConfig,omitempty" xml:"networkConfig,omitempty" type:"Struct"`
	// example:
	//
	// 2
	NodeAmount *int64            `json:"nodeAmount,omitempty" xml:"nodeAmount,omitempty"`
	NodeSpec   *LogstashNodeSpec `json:"nodeSpec,omitempty" xml:"nodeSpec,omitempty" type:"Struct"`
	// example:
	//
	// postpaid
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// rg-acfm2h5vbzd****
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// example:
	//
	// NORMAL
	Status *string         `json:"status,omitempty" xml:"status,omitempty"`
	Tags   []*LogstashTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// 2018-07-18T10:10:04.484Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// 6.7.0_with_X-Pack
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// example:
	//
	// 1
	ZoneCount *int64               `json:"zoneCount,omitempty" xml:"zoneCount,omitempty"`
	ZoneInfos []*LogstashZoneInfos `json:"zoneInfos,omitempty" xml:"zoneInfos,omitempty" type:"Repeated"`
}

func (s Logstash) String() string {
	return dara.Prettify(s)
}

func (s Logstash) GoString() string {
	return s.String()
}

func (s *Logstash) GetConfig() map[string]*string {
	return s.Config
}

func (s *Logstash) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *Logstash) GetDataNode() *bool {
	return s.DataNode
}

func (s *Logstash) GetDescription() *string {
	return s.Description
}

func (s *Logstash) GetEndTime() *int64 {
	return s.EndTime
}

func (s *Logstash) GetEndpointList() []*LogstashEndpointList {
	return s.EndpointList
}

func (s *Logstash) GetInstanceId() *string {
	return s.InstanceId
}

func (s *Logstash) GetNetworkConfig() *LogstashNetworkConfig {
	return s.NetworkConfig
}

func (s *Logstash) GetNodeAmount() *int64 {
	return s.NodeAmount
}

func (s *Logstash) GetNodeSpec() *LogstashNodeSpec {
	return s.NodeSpec
}

func (s *Logstash) GetPaymentType() *string {
	return s.PaymentType
}

func (s *Logstash) GetProtocol() *string {
	return s.Protocol
}

func (s *Logstash) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *Logstash) GetStatus() *string {
	return s.Status
}

func (s *Logstash) GetTags() []*LogstashTags {
	return s.Tags
}

func (s *Logstash) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *Logstash) GetVersion() *string {
	return s.Version
}

func (s *Logstash) GetZoneCount() *int64 {
	return s.ZoneCount
}

func (s *Logstash) GetZoneInfos() []*LogstashZoneInfos {
	return s.ZoneInfos
}

func (s *Logstash) SetConfig(v map[string]*string) *Logstash {
	s.Config = v
	return s
}

func (s *Logstash) SetCreatedAt(v string) *Logstash {
	s.CreatedAt = &v
	return s
}

func (s *Logstash) SetDataNode(v bool) *Logstash {
	s.DataNode = &v
	return s
}

func (s *Logstash) SetDescription(v string) *Logstash {
	s.Description = &v
	return s
}

func (s *Logstash) SetEndTime(v int64) *Logstash {
	s.EndTime = &v
	return s
}

func (s *Logstash) SetEndpointList(v []*LogstashEndpointList) *Logstash {
	s.EndpointList = v
	return s
}

func (s *Logstash) SetInstanceId(v string) *Logstash {
	s.InstanceId = &v
	return s
}

func (s *Logstash) SetNetworkConfig(v *LogstashNetworkConfig) *Logstash {
	s.NetworkConfig = v
	return s
}

func (s *Logstash) SetNodeAmount(v int64) *Logstash {
	s.NodeAmount = &v
	return s
}

func (s *Logstash) SetNodeSpec(v *LogstashNodeSpec) *Logstash {
	s.NodeSpec = v
	return s
}

func (s *Logstash) SetPaymentType(v string) *Logstash {
	s.PaymentType = &v
	return s
}

func (s *Logstash) SetProtocol(v string) *Logstash {
	s.Protocol = &v
	return s
}

func (s *Logstash) SetResourceGroupId(v string) *Logstash {
	s.ResourceGroupId = &v
	return s
}

func (s *Logstash) SetStatus(v string) *Logstash {
	s.Status = &v
	return s
}

func (s *Logstash) SetTags(v []*LogstashTags) *Logstash {
	s.Tags = v
	return s
}

func (s *Logstash) SetUpdatedAt(v string) *Logstash {
	s.UpdatedAt = &v
	return s
}

func (s *Logstash) SetVersion(v string) *Logstash {
	s.Version = &v
	return s
}

func (s *Logstash) SetZoneCount(v int64) *Logstash {
	s.ZoneCount = &v
	return s
}

func (s *Logstash) SetZoneInfos(v []*LogstashZoneInfos) *Logstash {
	s.ZoneInfos = v
	return s
}

func (s *Logstash) Validate() error {
	if s.EndpointList != nil {
		for _, item := range s.EndpointList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NetworkConfig != nil {
		if err := s.NetworkConfig.Validate(); err != nil {
			return err
		}
	}
	if s.NodeSpec != nil {
		if err := s.NodeSpec.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ZoneInfos != nil {
		for _, item := range s.ZoneInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type LogstashEndpointList struct {
	// example:
	//
	// 172.16.xx.xx
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// example:
	//
	// 9200
	Port *int64 `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s LogstashEndpointList) String() string {
	return dara.Prettify(s)
}

func (s LogstashEndpointList) GoString() string {
	return s.String()
}

func (s *LogstashEndpointList) GetHost() *string {
	return s.Host
}

func (s *LogstashEndpointList) GetPort() *int64 {
	return s.Port
}

func (s *LogstashEndpointList) GetZoneId() *string {
	return s.ZoneId
}

func (s *LogstashEndpointList) SetHost(v string) *LogstashEndpointList {
	s.Host = &v
	return s
}

func (s *LogstashEndpointList) SetPort(v int64) *LogstashEndpointList {
	s.Port = &v
	return s
}

func (s *LogstashEndpointList) SetZoneId(v string) *LogstashEndpointList {
	s.ZoneId = &v
	return s
}

func (s *LogstashEndpointList) Validate() error {
	return dara.Validate(s)
}

type LogstashNetworkConfig struct {
	// example:
	//
	// vpc
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// vpc-abc
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// example:
	//
	// cn-hangzhou-*
	VsArea *string `json:"vsArea,omitempty" xml:"vsArea,omitempty"`
	// example:
	//
	// vsw-def
	VswitchId *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
}

func (s LogstashNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s LogstashNetworkConfig) GoString() string {
	return s.String()
}

func (s *LogstashNetworkConfig) GetType() *string {
	return s.Type
}

func (s *LogstashNetworkConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *LogstashNetworkConfig) GetVsArea() *string {
	return s.VsArea
}

func (s *LogstashNetworkConfig) GetVswitchId() *string {
	return s.VswitchId
}

func (s *LogstashNetworkConfig) SetType(v string) *LogstashNetworkConfig {
	s.Type = &v
	return s
}

func (s *LogstashNetworkConfig) SetVpcId(v string) *LogstashNetworkConfig {
	s.VpcId = &v
	return s
}

func (s *LogstashNetworkConfig) SetVsArea(v string) *LogstashNetworkConfig {
	s.VsArea = &v
	return s
}

func (s *LogstashNetworkConfig) SetVswitchId(v string) *LogstashNetworkConfig {
	s.VswitchId = &v
	return s
}

func (s *LogstashNetworkConfig) Validate() error {
	return dara.Validate(s)
}

type LogstashNodeSpec struct {
	// example:
	//
	// 50
	Disk *int64 `json:"disk,omitempty" xml:"disk,omitempty"`
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// example:
	//
	// logstash.n4.small
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s LogstashNodeSpec) String() string {
	return dara.Prettify(s)
}

func (s LogstashNodeSpec) GoString() string {
	return s.String()
}

func (s *LogstashNodeSpec) GetDisk() *int64 {
	return s.Disk
}

func (s *LogstashNodeSpec) GetDiskType() *string {
	return s.DiskType
}

func (s *LogstashNodeSpec) GetSpec() *string {
	return s.Spec
}

func (s *LogstashNodeSpec) SetDisk(v int64) *LogstashNodeSpec {
	s.Disk = &v
	return s
}

func (s *LogstashNodeSpec) SetDiskType(v string) *LogstashNodeSpec {
	s.DiskType = &v
	return s
}

func (s *LogstashNodeSpec) SetSpec(v string) *LogstashNodeSpec {
	s.Spec = &v
	return s
}

func (s *LogstashNodeSpec) Validate() error {
	return dara.Validate(s)
}

type LogstashTags struct {
	// example:
	//
	// env
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// example:
	//
	// dev
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s LogstashTags) String() string {
	return dara.Prettify(s)
}

func (s LogstashTags) GoString() string {
	return s.String()
}

func (s *LogstashTags) GetTagKey() *string {
	return s.TagKey
}

func (s *LogstashTags) GetTagValue() *string {
	return s.TagValue
}

func (s *LogstashTags) SetTagKey(v string) *LogstashTags {
	s.TagKey = &v
	return s
}

func (s *LogstashTags) SetTagValue(v string) *LogstashTags {
	s.TagValue = &v
	return s
}

func (s *LogstashTags) Validate() error {
	return dara.Validate(s)
}

type LogstashZoneInfos struct {
	// example:
	//
	// NORMAL
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s LogstashZoneInfos) String() string {
	return dara.Prettify(s)
}

func (s LogstashZoneInfos) GoString() string {
	return s.String()
}

func (s *LogstashZoneInfos) GetStatus() *string {
	return s.Status
}

func (s *LogstashZoneInfos) GetZoneId() *string {
	return s.ZoneId
}

func (s *LogstashZoneInfos) SetStatus(v string) *LogstashZoneInfos {
	s.Status = &v
	return s
}

func (s *LogstashZoneInfos) SetZoneId(v string) *LogstashZoneInfos {
	s.ZoneId = &v
	return s
}

func (s *LogstashZoneInfos) Validate() error {
	return dara.Validate(s)
}
