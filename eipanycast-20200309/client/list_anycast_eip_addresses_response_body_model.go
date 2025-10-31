// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnycastEipAddressesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnycastList(v []*ListAnycastEipAddressesResponseBodyAnycastList) *ListAnycastEipAddressesResponseBody
	GetAnycastList() []*ListAnycastEipAddressesResponseBodyAnycastList
	SetNextToken(v string) *ListAnycastEipAddressesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAnycastEipAddressesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAnycastEipAddressesResponseBody
	GetTotalCount() *int32
}

type ListAnycastEipAddressesResponseBody struct {
	// The list of Anycast EIPs.
	AnycastList []*ListAnycastEipAddressesResponseBodyAnycastList `json:"AnycastList,omitempty" xml:"AnycastList,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If **NextToken*	- is not empty, the value of NextToken can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAnycastEipAddressesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAnycastEipAddressesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAnycastEipAddressesResponseBody) GetAnycastList() []*ListAnycastEipAddressesResponseBodyAnycastList {
	return s.AnycastList
}

func (s *ListAnycastEipAddressesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAnycastEipAddressesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAnycastEipAddressesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAnycastEipAddressesResponseBody) SetAnycastList(v []*ListAnycastEipAddressesResponseBodyAnycastList) *ListAnycastEipAddressesResponseBody {
	s.AnycastList = v
	return s
}

func (s *ListAnycastEipAddressesResponseBody) SetNextToken(v string) *ListAnycastEipAddressesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBody) SetRequestId(v string) *ListAnycastEipAddressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBody) SetTotalCount(v int32) *ListAnycastEipAddressesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBody) Validate() error {
	if s.AnycastList != nil {
		for _, item := range s.AnycastList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAnycastEipAddressesResponseBodyAnycastList struct {
	// The ID of the account to which the Anycast EIP belongs.
	//
	// example:
	//
	// 123440159596****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The list of cloud resources with which the Anycast EIPs are associated.
	AnycastEipBindInfoList []*ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList `json:"AnycastEipBindInfoList,omitempty" xml:"AnycastEipBindInfoList,omitempty" type:"Repeated"`
	// The ID of the Anycast EIP.
	//
	// example:
	//
	// aeip-2zeerraiwb7ujsxdc****
	AnycastId *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	// The maximum bandwidth of the Anycast EIP. Unit: Mbit/s.
	//
	// example:
	//
	// 200
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The service status of the Anycast EIP. Valid values:
	//
	// 	- **Normal**
	//
	// 	- **FinancialLocked**
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The time when the Anycast EIP was created.
	//
	// example:
	//
	// 2022-04-22T01:37:38Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the Anycast EIP.
	//
	// example:
	//
	// docdesc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The billing method of the Anycast EIP.
	//
	// **PostPaid**: pay-as-you-go
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The metering method of the Anycast EIP.
	//
	// **PayByTraffic**: pay-by-data-transfer
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
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-aekzthsmwsnfuni
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The access area of the Anycast EIP.
	//
	// **international**: regions outside the Chinese mainland
	//
	// example:
	//
	// international
	ServiceLocation *string `json:"ServiceLocation,omitempty" xml:"ServiceLocation,omitempty"`
	// Indicates whether the resource is created by the service account.
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 0
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
	// Associating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The information about the tags.
	Tags []*ListAnycastEipAddressesResponseBodyAnycastListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListAnycastEipAddressesResponseBodyAnycastList) String() string {
	return dara.Prettify(s)
}

func (s ListAnycastEipAddressesResponseBodyAnycastList) GoString() string {
	return s.String()
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) GetAliUid() *int64 {
	return s.AliUid
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) GetAnycastEipBindInfoList() []*ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList {
	return s.AnycastEipBindInfoList
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) GetAnycastId() *string {
	return s.AnycastId
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) GetDescription() *string {
	return s.Description
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) GetIpAddress() *string {
	return s.IpAddress
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) GetName() *string {
	return s.Name
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) GetServiceLocation() *string {
	return s.ServiceLocation
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) GetServiceManaged() *int32 {
	return s.ServiceManaged
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) GetStatus() *string {
	return s.Status
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) GetTags() []*ListAnycastEipAddressesResponseBodyAnycastListTags {
	return s.Tags
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetAliUid(v int64) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.AliUid = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetAnycastEipBindInfoList(v []*ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.AnycastEipBindInfoList = v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetAnycastId(v string) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.AnycastId = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetBandwidth(v int32) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.Bandwidth = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetBusinessStatus(v string) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.BusinessStatus = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetCreateTime(v string) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.CreateTime = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetDescription(v string) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.Description = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetInstanceChargeType(v string) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.InstanceChargeType = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetInternetChargeType(v string) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.InternetChargeType = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetIpAddress(v string) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.IpAddress = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetName(v string) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.Name = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetResourceGroupId(v string) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetServiceLocation(v string) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.ServiceLocation = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetServiceManaged(v int32) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.ServiceManaged = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetStatus(v string) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.Status = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) SetTags(v []*ListAnycastEipAddressesResponseBodyAnycastListTags) *ListAnycastEipAddressesResponseBodyAnycastList {
	s.Tags = v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastList) Validate() error {
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

type ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList struct {
	// The ID of the cloud resource with which the Anycast EIP is associated.
	//
	// example:
	//
	// lb-2zebb08phyczzawe****
	BindInstanceId *string `json:"BindInstanceId,omitempty" xml:"BindInstanceId,omitempty"`
	// The ID of the region where the cloud resource is deployed.
	//
	// example:
	//
	// us-west-1
	BindInstanceRegionId *string `json:"BindInstanceRegionId,omitempty" xml:"BindInstanceRegionId,omitempty"`
	// The type of cloud resource with which the Anycast EIP is associated.
	//
	// 	- **SlbInstance**: an internal-facing Classic Load Balancer (CLB) instance deployed in a virtual private cloud (VPC). CLB is formerly known as Server Load Balancer (SLB).
	//
	// 	- **NetworkInterface**: an elastic network interface (ENI).
	//
	// example:
	//
	// SlbInstance
	BindInstanceType *string `json:"BindInstanceType,omitempty" xml:"BindInstanceType,omitempty"`
	// The time when the Anycast EIP was associated.
	//
	// example:
	//
	// 2022-04-23T01:37:38Z
	BindTime *string `json:"BindTime,omitempty" xml:"BindTime,omitempty"`
}

func (s ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList) GoString() string {
	return s.String()
}

func (s *ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList) GetBindInstanceId() *string {
	return s.BindInstanceId
}

func (s *ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList) GetBindInstanceRegionId() *string {
	return s.BindInstanceRegionId
}

func (s *ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList) GetBindInstanceType() *string {
	return s.BindInstanceType
}

func (s *ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList) GetBindTime() *string {
	return s.BindTime
}

func (s *ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList) SetBindInstanceId(v string) *ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList {
	s.BindInstanceId = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList) SetBindInstanceRegionId(v string) *ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList {
	s.BindInstanceRegionId = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList) SetBindInstanceType(v string) *ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList {
	s.BindInstanceType = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList) SetBindTime(v string) *ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList {
	s.BindTime = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastListAnycastEipBindInfoList) Validate() error {
	return dara.Validate(s)
}

type ListAnycastEipAddressesResponseBodyAnycastListTags struct {
	// The tag key.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAnycastEipAddressesResponseBodyAnycastListTags) String() string {
	return dara.Prettify(s)
}

func (s ListAnycastEipAddressesResponseBodyAnycastListTags) GoString() string {
	return s.String()
}

func (s *ListAnycastEipAddressesResponseBodyAnycastListTags) GetKey() *string {
	return s.Key
}

func (s *ListAnycastEipAddressesResponseBodyAnycastListTags) GetValue() *string {
	return s.Value
}

func (s *ListAnycastEipAddressesResponseBodyAnycastListTags) SetKey(v string) *ListAnycastEipAddressesResponseBodyAnycastListTags {
	s.Key = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastListTags) SetValue(v string) *ListAnycastEipAddressesResponseBodyAnycastListTags {
	s.Value = &v
	return s
}

func (s *ListAnycastEipAddressesResponseBodyAnycastListTags) Validate() error {
	return dara.Validate(s)
}
