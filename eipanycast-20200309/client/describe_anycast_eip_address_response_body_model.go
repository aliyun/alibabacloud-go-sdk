// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAnycastEipAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v int64) *DescribeAnycastEipAddressResponseBody
	GetAliUid() *int64
	SetAnycastEipBindInfoList(v []*DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) *DescribeAnycastEipAddressResponseBody
	GetAnycastEipBindInfoList() []*DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList
	SetAnycastId(v string) *DescribeAnycastEipAddressResponseBody
	GetAnycastId() *string
	SetBandwidth(v int32) *DescribeAnycastEipAddressResponseBody
	GetBandwidth() *int32
	SetBid(v string) *DescribeAnycastEipAddressResponseBody
	GetBid() *string
	SetBusinessStatus(v string) *DescribeAnycastEipAddressResponseBody
	GetBusinessStatus() *string
	SetCreateTime(v string) *DescribeAnycastEipAddressResponseBody
	GetCreateTime() *string
	SetDescription(v string) *DescribeAnycastEipAddressResponseBody
	GetDescription() *string
	SetInstanceChargeType(v string) *DescribeAnycastEipAddressResponseBody
	GetInstanceChargeType() *string
	SetInternetChargeType(v string) *DescribeAnycastEipAddressResponseBody
	GetInternetChargeType() *string
	SetIpAddress(v string) *DescribeAnycastEipAddressResponseBody
	GetIpAddress() *string
	SetName(v string) *DescribeAnycastEipAddressResponseBody
	GetName() *string
	SetRequestId(v string) *DescribeAnycastEipAddressResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeAnycastEipAddressResponseBody
	GetResourceGroupId() *string
	SetServiceLocation(v string) *DescribeAnycastEipAddressResponseBody
	GetServiceLocation() *string
	SetServiceManaged(v int32) *DescribeAnycastEipAddressResponseBody
	GetServiceManaged() *int32
	SetStatus(v string) *DescribeAnycastEipAddressResponseBody
	GetStatus() *string
	SetTags(v []*DescribeAnycastEipAddressResponseBodyTags) *DescribeAnycastEipAddressResponseBody
	GetTags() []*DescribeAnycastEipAddressResponseBodyTags
}

type DescribeAnycastEipAddressResponseBody struct {
	// The ID of the account to which the Anycast EIP belongs.
	//
	// example:
	//
	// 25346073170691****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The information about the endpoint with which the Anycast EIP is associated.
	AnycastEipBindInfoList []*DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList `json:"AnycastEipBindInfoList,omitempty" xml:"AnycastEipBindInfoList,omitempty" type:"Repeated"`
	// The ID of the Anycast EIP.
	//
	// example:
	//
	// aeip-bp1ix34fralt4ykf3****
	AnycastId *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	// The maximum bandwidth of the Anycast EIP. Unit: Mbit/s.
	//
	// example:
	//
	// 200
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The BID of the account to which the Anycast EIP belongs.
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// The status of the Anycast EIP. Valid values:
	//
	// 	- **Normal**: running as expected
	//
	// 	- **FinancialLocked**: locked due to overdue payments
	//
	// 	- **RiskExpired**: locked due to security reasons.
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The point in time at which the Anycast EIP was created.
	//
	// The time follows the ISO8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-04-23T01:37:38Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the Anycast EIP.
	//
	// example:
	//
	// doctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The billing method of the Anycast EIP.
	//
	// Only **PostPaid*	- may be returned, which indicates the pay-as-you-go billing method.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The metering method of the Anycast EIP.
	//
	// Only **PayByTraffic*	- may be returned, which indicates the pay-by-data-transfer metering method.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The IP address of the Anycast EIP.
	//
	// example:
	//
	// 139.95.XX.XX
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The name of the Anycast EIP.
	//
	// example:
	//
	// docname
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfmzssisocarfy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The access area of the Anycast EIP.
	//
	// Only **international*	- may be returned, which indicates the areas outside the Chinese mainland.
	//
	// example:
	//
	// international
	ServiceLocation *string `json:"ServiceLocation,omitempty" xml:"ServiceLocation,omitempty"`
	// Indicates whether the instance is managed. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no.
	//
	// example:
	//
	// 1
	ServiceManaged *int32 `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The status of the Anycast EIP.
	//
	// 	- **Associating**
	//
	// 	- **Unassociating**
	//
	// 	- **Allocated**
	//
	// 	- **Associated**
	//
	// 	- **Modifying**
	//
	// 	- **Releasing**
	//
	// 	- **Released**
	//
	// example:
	//
	// Associated
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The information about the tags.
	Tags []*DescribeAnycastEipAddressResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeAnycastEipAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnycastEipAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAnycastEipAddressResponseBody) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeAnycastEipAddressResponseBody) GetAnycastEipBindInfoList() []*DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList {
	return s.AnycastEipBindInfoList
}

func (s *DescribeAnycastEipAddressResponseBody) GetAnycastId() *string {
	return s.AnycastId
}

func (s *DescribeAnycastEipAddressResponseBody) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeAnycastEipAddressResponseBody) GetBid() *string {
	return s.Bid
}

func (s *DescribeAnycastEipAddressResponseBody) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeAnycastEipAddressResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAnycastEipAddressResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeAnycastEipAddressResponseBody) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeAnycastEipAddressResponseBody) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeAnycastEipAddressResponseBody) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeAnycastEipAddressResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeAnycastEipAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAnycastEipAddressResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeAnycastEipAddressResponseBody) GetServiceLocation() *string {
	return s.ServiceLocation
}

func (s *DescribeAnycastEipAddressResponseBody) GetServiceManaged() *int32 {
	return s.ServiceManaged
}

func (s *DescribeAnycastEipAddressResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeAnycastEipAddressResponseBody) GetTags() []*DescribeAnycastEipAddressResponseBodyTags {
	return s.Tags
}

func (s *DescribeAnycastEipAddressResponseBody) SetAliUid(v int64) *DescribeAnycastEipAddressResponseBody {
	s.AliUid = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetAnycastEipBindInfoList(v []*DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) *DescribeAnycastEipAddressResponseBody {
	s.AnycastEipBindInfoList = v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetAnycastId(v string) *DescribeAnycastEipAddressResponseBody {
	s.AnycastId = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetBandwidth(v int32) *DescribeAnycastEipAddressResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetBid(v string) *DescribeAnycastEipAddressResponseBody {
	s.Bid = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetBusinessStatus(v string) *DescribeAnycastEipAddressResponseBody {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetCreateTime(v string) *DescribeAnycastEipAddressResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetDescription(v string) *DescribeAnycastEipAddressResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetInstanceChargeType(v string) *DescribeAnycastEipAddressResponseBody {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetInternetChargeType(v string) *DescribeAnycastEipAddressResponseBody {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetIpAddress(v string) *DescribeAnycastEipAddressResponseBody {
	s.IpAddress = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetName(v string) *DescribeAnycastEipAddressResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetRequestId(v string) *DescribeAnycastEipAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetResourceGroupId(v string) *DescribeAnycastEipAddressResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetServiceLocation(v string) *DescribeAnycastEipAddressResponseBody {
	s.ServiceLocation = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetServiceManaged(v int32) *DescribeAnycastEipAddressResponseBody {
	s.ServiceManaged = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetStatus(v string) *DescribeAnycastEipAddressResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) SetTags(v []*DescribeAnycastEipAddressResponseBodyTags) *DescribeAnycastEipAddressResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeAnycastEipAddressResponseBody) Validate() error {
	if s.AnycastEipBindInfoList != nil {
		for _, item := range s.AnycastEipBindInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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
	return nil
}

type DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList struct {
	// The association mode. Valid values:
	//
	// 	- **Default**: the default mode. In this mode, the associated endpoint serves as the default origin server.
	//
	// 	- **Normal**: the standard mode. In this mode, the associated endpoint serves as a standard origin server.
	//
	// example:
	//
	// Default
	AssociationMode *string `json:"AssociationMode,omitempty" xml:"AssociationMode,omitempty"`
	// The ID of the endpoint with which the Anycast EIP is associated.
	//
	// example:
	//
	// lb-2zebb08phyczzawe****
	BindInstanceId *string `json:"BindInstanceId,omitempty" xml:"BindInstanceId,omitempty"`
	// The ID of the region in which the endpoint is deployed.
	//
	// example:
	//
	// us-west-1
	BindInstanceRegionId *string `json:"BindInstanceRegionId,omitempty" xml:"BindInstanceRegionId,omitempty"`
	// The type of endpoint with which the Anycast EIP is associated. Valid values:
	//
	// 	- **SlbInstance**: a CLB instance in a VPC.
	//
	// 	- **NetworkInterface**: an elastic network interface (ENI).
	//
	// example:
	//
	// SlbInstance
	BindInstanceType *string `json:"BindInstanceType,omitempty" xml:"BindInstanceType,omitempty"`
	// The time when the Anycast EIP was associated.
	//
	// The time follows the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-04-23T02:37:38Z
	BindTime *string `json:"BindTime,omitempty" xml:"BindTime,omitempty"`
	// The information about the access points in associated access areas when you associate an Anycast EIP with a cloud resource.
	//
	// If this is your first time associating an Anycast EIP with an endpoint, the system returns information about access points in all access areas.
	PopLocations []*DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoListPopLocations `json:"PopLocations,omitempty" xml:"PopLocations,omitempty" type:"Repeated"`
	// The secondary private IP address of the associated ENI.
	//
	// This parameter is valid only when **BindInstanceType*	- is set to **NetworkInterface**.
	//
	// example:
	//
	// 192.168.XX.XX
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The status of the endpoint. Valid values:
	//
	// 	- **BINDING**
	//
	// 	- **BINDED**
	//
	// 	- **UNBINDING**
	//
	// 	- **DELETED**
	//
	// 	- **MODIFYING**
	//
	// example:
	//
	// BINDING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) GoString() string {
	return s.String()
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) GetAssociationMode() *string {
	return s.AssociationMode
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) GetBindInstanceId() *string {
	return s.BindInstanceId
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) GetBindInstanceRegionId() *string {
	return s.BindInstanceRegionId
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) GetBindInstanceType() *string {
	return s.BindInstanceType
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) GetBindTime() *string {
	return s.BindTime
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) GetPopLocations() []*DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoListPopLocations {
	return s.PopLocations
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) GetStatus() *string {
	return s.Status
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) SetAssociationMode(v string) *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList {
	s.AssociationMode = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) SetBindInstanceId(v string) *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList {
	s.BindInstanceId = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) SetBindInstanceRegionId(v string) *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList {
	s.BindInstanceRegionId = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) SetBindInstanceType(v string) *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList {
	s.BindInstanceType = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) SetBindTime(v string) *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList {
	s.BindTime = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) SetPopLocations(v []*DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoListPopLocations) *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList {
	s.PopLocations = v
	return s
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) SetPrivateIpAddress(v string) *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) SetStatus(v string) *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList {
	s.Status = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoList) Validate() error {
	if s.PopLocations != nil {
		for _, item := range s.PopLocations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoListPopLocations struct {
	// The information about the access points in associated access areas when you associate an Anycast EIP with a cloud resource.
	//
	// If this is your first time associating an Anycast EIP with an endpoint, the system returns information about access points in all access areas.
	//
	// example:
	//
	// us-west-1-pop
	PopLocation *string `json:"PopLocation,omitempty" xml:"PopLocation,omitempty"`
}

func (s DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoListPopLocations) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoListPopLocations) GoString() string {
	return s.String()
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoListPopLocations) GetPopLocation() *string {
	return s.PopLocation
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoListPopLocations) SetPopLocation(v string) *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoListPopLocations {
	s.PopLocation = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBodyAnycastEipBindInfoListPopLocations) Validate() error {
	return dara.Validate(s)
}

type DescribeAnycastEipAddressResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// 1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAnycastEipAddressResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnycastEipAddressResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeAnycastEipAddressResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *DescribeAnycastEipAddressResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *DescribeAnycastEipAddressResponseBodyTags) SetKey(v string) *DescribeAnycastEipAddressResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBodyTags) SetValue(v string) *DescribeAnycastEipAddressResponseBodyTags {
	s.Value = &v
	return s
}

func (s *DescribeAnycastEipAddressResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
