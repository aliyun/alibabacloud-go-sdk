// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalDataNetworkListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeGlobalDataNetworkListResponseBodyItems) *DescribeGlobalDataNetworkListResponseBody
	GetItems() *DescribeGlobalDataNetworkListResponseBodyItems
	SetPageNumber(v string) *DescribeGlobalDataNetworkListResponseBody
	GetPageNumber() *string
	SetPageRecordCount(v string) *DescribeGlobalDataNetworkListResponseBody
	GetPageRecordCount() *string
	SetRequestId(v string) *DescribeGlobalDataNetworkListResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v string) *DescribeGlobalDataNetworkListResponseBody
	GetTotalRecordCount() *string
}

type DescribeGlobalDataNetworkListResponseBody struct {
	Items *DescribeGlobalDataNetworkListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageRecordCount *string `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// CD35F3-F3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalRecordCount *string `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeGlobalDataNetworkListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDataNetworkListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDataNetworkListResponseBody) GetItems() *DescribeGlobalDataNetworkListResponseBodyItems {
	return s.Items
}

func (s *DescribeGlobalDataNetworkListResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeGlobalDataNetworkListResponseBody) GetPageRecordCount() *string {
	return s.PageRecordCount
}

func (s *DescribeGlobalDataNetworkListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGlobalDataNetworkListResponseBody) GetTotalRecordCount() *string {
	return s.TotalRecordCount
}

func (s *DescribeGlobalDataNetworkListResponseBody) SetItems(v *DescribeGlobalDataNetworkListResponseBodyItems) *DescribeGlobalDataNetworkListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBody) SetPageNumber(v string) *DescribeGlobalDataNetworkListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBody) SetPageRecordCount(v string) *DescribeGlobalDataNetworkListResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBody) SetRequestId(v string) *DescribeGlobalDataNetworkListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBody) SetTotalRecordCount(v string) *DescribeGlobalDataNetworkListResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGlobalDataNetworkListResponseBodyItems struct {
	Networks []*DescribeGlobalDataNetworkListResponseBodyItemsNetworks `json:"Networks,omitempty" xml:"Networks,omitempty" type:"Repeated"`
}

func (s DescribeGlobalDataNetworkListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDataNetworkListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDataNetworkListResponseBodyItems) GetNetworks() []*DescribeGlobalDataNetworkListResponseBodyItemsNetworks {
	return s.Networks
}

func (s *DescribeGlobalDataNetworkListResponseBodyItems) SetNetworks(v []*DescribeGlobalDataNetworkListResponseBodyItemsNetworks) *DescribeGlobalDataNetworkListResponseBodyItems {
	s.Networks = v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBodyItems) Validate() error {
	if s.Networks != nil {
		for _, item := range s.Networks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGlobalDataNetworkListResponseBodyItemsNetworks struct {
	Channels []*DescribeGlobalDataNetworkListResponseBodyItemsNetworksChannels `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-03-25T09:37:10Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// mygdn
	NetworkDescription *string `json:"NetworkDescription,omitempty" xml:"NetworkDescription,omitempty"`
	// GDN ID
	//
	// example:
	//
	// gdn-xxx
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// example:
	//
	// Running
	NetworkStatus   *string                                                                `json:"NetworkStatus,omitempty" xml:"NetworkStatus,omitempty"`
	NetworkTopology *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopology `json:"NetworkTopology,omitempty" xml:"NetworkTopology,omitempty" type:"Struct"`
}

func (s DescribeGlobalDataNetworkListResponseBodyItemsNetworks) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDataNetworkListResponseBodyItemsNetworks) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworks) GetChannels() []*DescribeGlobalDataNetworkListResponseBodyItemsNetworksChannels {
	return s.Channels
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworks) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworks) GetNetworkDescription() *string {
	return s.NetworkDescription
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworks) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworks) GetNetworkStatus() *string {
	return s.NetworkStatus
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworks) GetNetworkTopology() *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopology {
	return s.NetworkTopology
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworks) SetChannels(v []*DescribeGlobalDataNetworkListResponseBodyItemsNetworksChannels) *DescribeGlobalDataNetworkListResponseBodyItemsNetworks {
	s.Channels = v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworks) SetCreateTime(v string) *DescribeGlobalDataNetworkListResponseBodyItemsNetworks {
	s.CreateTime = &v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworks) SetNetworkDescription(v string) *DescribeGlobalDataNetworkListResponseBodyItemsNetworks {
	s.NetworkDescription = &v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworks) SetNetworkId(v string) *DescribeGlobalDataNetworkListResponseBodyItemsNetworks {
	s.NetworkId = &v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworks) SetNetworkStatus(v string) *DescribeGlobalDataNetworkListResponseBodyItemsNetworks {
	s.NetworkStatus = &v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworks) SetNetworkTopology(v *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopology) *DescribeGlobalDataNetworkListResponseBodyItemsNetworks {
	s.NetworkTopology = v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworks) Validate() error {
	if s.Channels != nil {
		for _, item := range s.Channels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NetworkTopology != nil {
		if err := s.NetworkTopology.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGlobalDataNetworkListResponseBodyItemsNetworksChannels struct {
	// example:
	//
	// gdc-xxx
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// Syncing
	ChannelStatus *string `json:"ChannelStatus,omitempty" xml:"ChannelStatus,omitempty"`
	// example:
	//
	// true
	FreezeSourceDuringSync *bool `json:"FreezeSourceDuringSync,omitempty" xml:"FreezeSourceDuringSync,omitempty"`
	// example:
	//
	// 11.45%
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s DescribeGlobalDataNetworkListResponseBodyItemsNetworksChannels) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDataNetworkListResponseBodyItemsNetworksChannels) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksChannels) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksChannels) GetChannelStatus() *string {
	return s.ChannelStatus
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksChannels) GetFreezeSourceDuringSync() *bool {
	return s.FreezeSourceDuringSync
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksChannels) GetProgress() *string {
	return s.Progress
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksChannels) SetChannelId(v string) *DescribeGlobalDataNetworkListResponseBodyItemsNetworksChannels {
	s.ChannelId = &v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksChannels) SetChannelStatus(v string) *DescribeGlobalDataNetworkListResponseBodyItemsNetworksChannels {
	s.ChannelStatus = &v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksChannels) SetFreezeSourceDuringSync(v bool) *DescribeGlobalDataNetworkListResponseBodyItemsNetworksChannels {
	s.FreezeSourceDuringSync = &v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksChannels) SetProgress(v string) *DescribeGlobalDataNetworkListResponseBodyItemsNetworksChannels {
	s.Progress = &v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksChannels) Validate() error {
	return dara.Validate(s)
}

type DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopology struct {
	Destinations []*DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologyDestinations `json:"Destinations,omitempty" xml:"Destinations,omitempty" type:"Repeated"`
	Sources      []*DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologySources      `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
}

func (s DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopology) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopology) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopology) GetDestinations() []*DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologyDestinations {
	return s.Destinations
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopology) GetSources() []*DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologySources {
	return s.Sources
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopology) SetDestinations(v []*DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologyDestinations) *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopology {
	s.Destinations = v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopology) SetSources(v []*DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologySources) *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopology {
	s.Sources = v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopology) Validate() error {
	if s.Destinations != nil {
		for _, item := range s.Destinations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Sources != nil {
		for _, item := range s.Sources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologyDestinations struct {
	// example:
	//
	// /
	DestinationFileSystemPath *string `json:"DestinationFileSystemPath,omitempty" xml:"DestinationFileSystemPath,omitempty"`
	// example:
	//
	// pfs-xxx
	DestinationId *string `json:"DestinationId,omitempty" xml:"DestinationId,omitempty"`
	// example:
	//
	// cn-beijing
	DestinationRegion *string `json:"DestinationRegion,omitempty" xml:"DestinationRegion,omitempty"`
	// example:
	//
	// pfs
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
}

func (s DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologyDestinations) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologyDestinations) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologyDestinations) GetDestinationFileSystemPath() *string {
	return s.DestinationFileSystemPath
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologyDestinations) GetDestinationId() *string {
	return s.DestinationId
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologyDestinations) GetDestinationRegion() *string {
	return s.DestinationRegion
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologyDestinations) GetDestinationType() *string {
	return s.DestinationType
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologyDestinations) SetDestinationFileSystemPath(v string) *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologyDestinations {
	s.DestinationFileSystemPath = &v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologyDestinations) SetDestinationId(v string) *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologyDestinations {
	s.DestinationId = &v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologyDestinations) SetDestinationRegion(v string) *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologyDestinations {
	s.DestinationRegion = &v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologyDestinations) SetDestinationType(v string) *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologyDestinations {
	s.DestinationType = &v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologyDestinations) Validate() error {
	return dara.Validate(s)
}

type DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologySources struct {
	// example:
	//
	// /
	SourceFileSystemPath *string `json:"SourceFileSystemPath,omitempty" xml:"SourceFileSystemPath,omitempty"`
	// example:
	//
	// oss-xxx
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// example:
	//
	// cn-wulanchabu
	SourceRegion *string `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
	// example:
	//
	// oss
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologySources) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologySources) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologySources) GetSourceFileSystemPath() *string {
	return s.SourceFileSystemPath
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologySources) GetSourceId() *string {
	return s.SourceId
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologySources) GetSourceRegion() *string {
	return s.SourceRegion
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologySources) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologySources) SetSourceFileSystemPath(v string) *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologySources {
	s.SourceFileSystemPath = &v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologySources) SetSourceId(v string) *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologySources {
	s.SourceId = &v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologySources) SetSourceRegion(v string) *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologySources {
	s.SourceRegion = &v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologySources) SetSourceType(v string) *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologySources {
	s.SourceType = &v
	return s
}

func (s *DescribeGlobalDataNetworkListResponseBodyItemsNetworksNetworkTopologySources) Validate() error {
	return dara.Validate(s)
}
