// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNetworkInterfaceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionTrackingConfiguration(v *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration) *ModifyNetworkInterfaceAttributeRequest
	GetConnectionTrackingConfiguration() *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration
	SetDeleteOnRelease(v bool) *ModifyNetworkInterfaceAttributeRequest
	GetDeleteOnRelease() *bool
	SetDescription(v string) *ModifyNetworkInterfaceAttributeRequest
	GetDescription() *string
	SetEnablePrimaryIPv6(v bool) *ModifyNetworkInterfaceAttributeRequest
	GetEnablePrimaryIPv6() *bool
	SetEnhancedNetwork(v *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) *ModifyNetworkInterfaceAttributeRequest
	GetEnhancedNetwork() *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork
	SetNetworkInterfaceId(v string) *ModifyNetworkInterfaceAttributeRequest
	GetNetworkInterfaceId() *string
	SetNetworkInterfaceName(v string) *ModifyNetworkInterfaceAttributeRequest
	GetNetworkInterfaceName() *string
	SetNetworkInterfaceTrafficConfig(v *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) *ModifyNetworkInterfaceAttributeRequest
	GetNetworkInterfaceTrafficConfig() *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig
	SetOwnerAccount(v string) *ModifyNetworkInterfaceAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyNetworkInterfaceAttributeRequest
	GetOwnerId() *int64
	SetQueueNumber(v int32) *ModifyNetworkInterfaceAttributeRequest
	GetQueueNumber() *int32
	SetRegionId(v string) *ModifyNetworkInterfaceAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyNetworkInterfaceAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyNetworkInterfaceAttributeRequest
	GetResourceOwnerId() *int64
	SetRxQueueSize(v int32) *ModifyNetworkInterfaceAttributeRequest
	GetRxQueueSize() *int32
	SetSecurityGroupId(v []*string) *ModifyNetworkInterfaceAttributeRequest
	GetSecurityGroupId() []*string
	SetSourceDestCheck(v bool) *ModifyNetworkInterfaceAttributeRequest
	GetSourceDestCheck() *bool
	SetTxQueueSize(v int32) *ModifyNetworkInterfaceAttributeRequest
	GetTxQueueSize() *int32
}

type ModifyNetworkInterfaceAttributeRequest struct {
	// The network connectivity tracking configuration.
	ConnectionTrackingConfiguration *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration `json:"ConnectionTrackingConfiguration,omitempty" xml:"ConnectionTrackingConfiguration,omitempty" type:"Struct"`
	// Specifies whether to retain the ENI when the associated instance is released. Valid values:
	//
	// example:
	//
	// true
	DeleteOnRelease *bool `json:"DeleteOnRelease,omitempty" xml:"DeleteOnRelease,omitempty"`
	// The description of the network interface controller (NIC). The description must be 2 to 255 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// testDescription
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnablePrimaryIPv6 *bool   `json:"EnablePrimaryIPv6,omitempty" xml:"EnablePrimaryIPv6,omitempty"`
	// This parameter is not publicly available.
	EnhancedNetwork *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork `json:"EnhancedNetwork,omitempty" xml:"EnhancedNetwork,omitempty" type:"Struct"`
	// The ID of the network interface controller (NIC).
	//
	// This parameter is required.
	//
	// example:
	//
	// eni-bp67acfmxazb4p****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The name of the network interface controller (NIC). The name must be 2 to 128 characters in length and must start with a letter or a Chinese character. It cannot start with `http://` or `https://`. The name can contain characters under the letter categorization in Unicode, including English letters, Chinese characters, and digits. It can also contain colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// eniTestName
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	// The communication parameter of the network interface controller (NIC).
	NetworkInterfaceTrafficConfig *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig `json:"NetworkInterfaceTrafficConfig,omitempty" xml:"NetworkInterfaceTrafficConfig,omitempty" type:"Struct"`
	OwnerAccount                  *string                                                              `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                       *int64                                                               `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of queues for the ENI. Valid values: 1 to 2048.
	//
	// example:
	//
	// 8
	QueueNumber *int32 `json:"QueueNumber,omitempty" xml:"QueueNumber,omitempty"`
	// The region ID of the network interface controller (NIC). You can invoke [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The inbound queue depth of the network interface controller (NIC).
	//
	// example:
	//
	// 8192
	RxQueueSize *int32 `json:"RxQueueSize,omitempty" xml:"RxQueueSize,omitempty"`
	// The list of security group IDs. The secondary network interface controller (NIC) is added to the specified security groups and removed from the existing security groups.
	SecurityGroupId []*string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Repeated"`
	// Specifies whether to enable source/destination checking. We recommend that you enable this feature to improve network security. Valid values:
	//
	// example:
	//
	// false
	SourceDestCheck *bool `json:"SourceDestCheck,omitempty" xml:"SourceDestCheck,omitempty"`
	// The outbound queue depth of the network interface controller (NIC).
	//
	// example:
	//
	// 8192
	TxQueueSize *int32 `json:"TxQueueSize,omitempty" xml:"TxQueueSize,omitempty"`
}

func (s ModifyNetworkInterfaceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkInterfaceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetConnectionTrackingConfiguration() *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration {
	return s.ConnectionTrackingConfiguration
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetDeleteOnRelease() *bool {
	return s.DeleteOnRelease
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetEnablePrimaryIPv6() *bool {
	return s.EnablePrimaryIPv6
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetEnhancedNetwork() *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork {
	return s.EnhancedNetwork
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetNetworkInterfaceName() *string {
	return s.NetworkInterfaceName
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetNetworkInterfaceTrafficConfig() *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig {
	return s.NetworkInterfaceTrafficConfig
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetQueueNumber() *int32 {
	return s.QueueNumber
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetRxQueueSize() *int32 {
	return s.RxQueueSize
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetSecurityGroupId() []*string {
	return s.SecurityGroupId
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetSourceDestCheck() *bool {
	return s.SourceDestCheck
}

func (s *ModifyNetworkInterfaceAttributeRequest) GetTxQueueSize() *int32 {
	return s.TxQueueSize
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetConnectionTrackingConfiguration(v *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration) *ModifyNetworkInterfaceAttributeRequest {
	s.ConnectionTrackingConfiguration = v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetDeleteOnRelease(v bool) *ModifyNetworkInterfaceAttributeRequest {
	s.DeleteOnRelease = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetDescription(v string) *ModifyNetworkInterfaceAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetEnablePrimaryIPv6(v bool) *ModifyNetworkInterfaceAttributeRequest {
	s.EnablePrimaryIPv6 = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetEnhancedNetwork(v *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) *ModifyNetworkInterfaceAttributeRequest {
	s.EnhancedNetwork = v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetNetworkInterfaceId(v string) *ModifyNetworkInterfaceAttributeRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetNetworkInterfaceName(v string) *ModifyNetworkInterfaceAttributeRequest {
	s.NetworkInterfaceName = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetNetworkInterfaceTrafficConfig(v *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) *ModifyNetworkInterfaceAttributeRequest {
	s.NetworkInterfaceTrafficConfig = v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetOwnerAccount(v string) *ModifyNetworkInterfaceAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetOwnerId(v int64) *ModifyNetworkInterfaceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetQueueNumber(v int32) *ModifyNetworkInterfaceAttributeRequest {
	s.QueueNumber = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetRegionId(v string) *ModifyNetworkInterfaceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetResourceOwnerAccount(v string) *ModifyNetworkInterfaceAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetResourceOwnerId(v int64) *ModifyNetworkInterfaceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetRxQueueSize(v int32) *ModifyNetworkInterfaceAttributeRequest {
	s.RxQueueSize = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetSecurityGroupId(v []*string) *ModifyNetworkInterfaceAttributeRequest {
	s.SecurityGroupId = v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetSourceDestCheck(v bool) *ModifyNetworkInterfaceAttributeRequest {
	s.SourceDestCheck = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) SetTxQueueSize(v int32) *ModifyNetworkInterfaceAttributeRequest {
	s.TxQueueSize = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequest) Validate() error {
	if s.ConnectionTrackingConfiguration != nil {
		if err := s.ConnectionTrackingConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.EnhancedNetwork != nil {
		if err := s.EnhancedNetwork.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkInterfaceTrafficConfig != nil {
		if err := s.NetworkInterfaceTrafficConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration struct {
	// The timeout period for TCP connections in the TIME_WAIT and CLOSED states. Unit: seconds. Valid values: integers from 3 to 15.
	//
	// example:
	//
	// 3
	TcpClosedAndTimeWaitTimeout *int32 `json:"TcpClosedAndTimeWaitTimeout,omitempty" xml:"TcpClosedAndTimeWaitTimeout,omitempty"`
	// The timeout period for established TCP connections. Unit: seconds. Valid values: [30, 60, 80, 100, 200, 300, 500, 700, 910].
	//
	// example:
	//
	// 910
	TcpEstablishedTimeout *int32 `json:"TcpEstablishedTimeout,omitempty" xml:"TcpEstablishedTimeout,omitempty"`
	// The timeout period for UDP flows. Unit: seconds. Valid values: [10, 20, 30, 60, 80, 100].
	//
	// example:
	//
	// 30
	UdpTimeout *int32 `json:"UdpTimeout,omitempty" xml:"UdpTimeout,omitempty"`
}

func (s ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration) GoString() string {
	return s.String()
}

func (s *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration) GetTcpClosedAndTimeWaitTimeout() *int32 {
	return s.TcpClosedAndTimeWaitTimeout
}

func (s *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration) GetTcpEstablishedTimeout() *int32 {
	return s.TcpEstablishedTimeout
}

func (s *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration) GetUdpTimeout() *int32 {
	return s.UdpTimeout
}

func (s *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration) SetTcpClosedAndTimeWaitTimeout(v int32) *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration {
	s.TcpClosedAndTimeWaitTimeout = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration) SetTcpEstablishedTimeout(v int32) *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration {
	s.TcpEstablishedTimeout = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration) SetUdpTimeout(v int32) *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration {
	s.UdpTimeout = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequestConnectionTrackingConfiguration) Validate() error {
	return dara.Validate(s)
}

type ModifyNetworkInterfaceAttributeRequestEnhancedNetwork struct {
	EnableExpress *bool `json:"EnableExpress,omitempty" xml:"EnableExpress,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// false
	EnableRss *bool `json:"EnableRss,omitempty" xml:"EnableRss,omitempty"`
	// This parameter is not publicly available.
	//
	// example:
	//
	// true
	EnableSriov                     *bool  `json:"EnableSriov,omitempty" xml:"EnableSriov,omitempty"`
	VirtualFunctionQuantity         *int32 `json:"VirtualFunctionQuantity,omitempty" xml:"VirtualFunctionQuantity,omitempty"`
	VirtualFunctionTotalQueueNumber *int32 `json:"VirtualFunctionTotalQueueNumber,omitempty" xml:"VirtualFunctionTotalQueueNumber,omitempty"`
}

func (s ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) GoString() string {
	return s.String()
}

func (s *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) GetEnableExpress() *bool {
	return s.EnableExpress
}

func (s *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) GetEnableRss() *bool {
	return s.EnableRss
}

func (s *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) GetEnableSriov() *bool {
	return s.EnableSriov
}

func (s *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) GetVirtualFunctionQuantity() *int32 {
	return s.VirtualFunctionQuantity
}

func (s *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) GetVirtualFunctionTotalQueueNumber() *int32 {
	return s.VirtualFunctionTotalQueueNumber
}

func (s *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) SetEnableExpress(v bool) *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork {
	s.EnableExpress = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) SetEnableRss(v bool) *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork {
	s.EnableRss = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) SetEnableSriov(v bool) *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork {
	s.EnableSriov = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) SetVirtualFunctionQuantity(v int32) *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork {
	s.VirtualFunctionQuantity = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) SetVirtualFunctionTotalQueueNumber(v int32) *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork {
	s.VirtualFunctionTotalQueueNumber = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequestEnhancedNetwork) Validate() error {
	return dara.Validate(s)
}

type ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig struct {
	// The communication mode of the ENI. Valid values:
	//
	// example:
	//
	// HighPerformance
	NetworkInterfaceTrafficMode *string `json:"NetworkInterfaceTrafficMode,omitempty" xml:"NetworkInterfaceTrafficMode,omitempty"`
	// The number of queues for the network interface controller (NIC).
	//
	// example:
	//
	// 8
	QueueNumber *int32 `json:"QueueNumber,omitempty" xml:"QueueNumber,omitempty"`
	// The number of queues for the RDMA ENI.
	//
	// example:
	//
	// 8
	QueuePairNumber *int32 `json:"QueuePairNumber,omitempty" xml:"QueuePairNumber,omitempty"`
	// The inbound queue depth of the network interface controller (NIC).
	//
	// example:
	//
	// 8192
	RxQueueSize *int32 `json:"RxQueueSize,omitempty" xml:"RxQueueSize,omitempty"`
	// The outbound queue depth of the network interface controller (NIC).
	//
	// example:
	//
	// 8192
	TxQueueSize *int32 `json:"TxQueueSize,omitempty" xml:"TxQueueSize,omitempty"`
}

func (s ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) GoString() string {
	return s.String()
}

func (s *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) GetNetworkInterfaceTrafficMode() *string {
	return s.NetworkInterfaceTrafficMode
}

func (s *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) GetQueueNumber() *int32 {
	return s.QueueNumber
}

func (s *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) GetQueuePairNumber() *int32 {
	return s.QueuePairNumber
}

func (s *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) GetRxQueueSize() *int32 {
	return s.RxQueueSize
}

func (s *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) GetTxQueueSize() *int32 {
	return s.TxQueueSize
}

func (s *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) SetNetworkInterfaceTrafficMode(v string) *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig {
	s.NetworkInterfaceTrafficMode = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) SetQueueNumber(v int32) *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig {
	s.QueueNumber = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) SetQueuePairNumber(v int32) *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig {
	s.QueuePairNumber = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) SetRxQueueSize(v int32) *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig {
	s.RxQueueSize = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) SetTxQueueSize(v int32) *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig {
	s.TxQueueSize = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeRequestNetworkInterfaceTrafficConfig) Validate() error {
	return dara.Validate(s)
}
