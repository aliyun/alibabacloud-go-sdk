// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnycastEipAddressesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnycastEipAddress(v string) *ListAnycastEipAddressesRequest
	GetAnycastEipAddress() *string
	SetAnycastId(v string) *ListAnycastEipAddressesRequest
	GetAnycastId() *string
	SetAnycastIds(v []*string) *ListAnycastEipAddressesRequest
	GetAnycastIds() []*string
	SetBindInstanceIds(v []*string) *ListAnycastEipAddressesRequest
	GetBindInstanceIds() []*string
	SetBusinessStatus(v string) *ListAnycastEipAddressesRequest
	GetBusinessStatus() *string
	SetInstanceChargeType(v string) *ListAnycastEipAddressesRequest
	GetInstanceChargeType() *string
	SetInternetChargeType(v string) *ListAnycastEipAddressesRequest
	GetInternetChargeType() *string
	SetMaxResults(v int32) *ListAnycastEipAddressesRequest
	GetMaxResults() *int32
	SetName(v string) *ListAnycastEipAddressesRequest
	GetName() *string
	SetNextToken(v string) *ListAnycastEipAddressesRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListAnycastEipAddressesRequest
	GetResourceGroupId() *string
	SetServiceLocation(v string) *ListAnycastEipAddressesRequest
	GetServiceLocation() *string
	SetStatus(v string) *ListAnycastEipAddressesRequest
	GetStatus() *string
	SetTags(v []*ListAnycastEipAddressesRequestTags) *ListAnycastEipAddressesRequest
	GetTags() []*ListAnycastEipAddressesRequestTags
}

type ListAnycastEipAddressesRequest struct {
	// The IP address that is allocated to the Anycast EIP.
	//
	// example:
	//
	// 139.95.XX.XX
	AnycastEipAddress *string `json:"AnycastEipAddress,omitempty" xml:"AnycastEipAddress,omitempty"`
	// The ID of the Anycast EIP.
	//
	// >  To ensure the accuracy of the query result, we do not recommend that you specify **AnycastIds*	- and **AnycastId*	- at the same time.
	//
	// example:
	//
	// aeip-2zeerraiwb7ujsxdc****
	AnycastId *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	// The IDs of Anycast EIPs.
	//
	// You can enter at most 50 Anycast EIP IDs.
	//
	// >  To ensure the accuracy of the query result, we do not recommend that you specify **AnycastIds*	- and **AnycastId*	- at the same time.
	AnycastIds []*string `json:"AnycastIds,omitempty" xml:"AnycastIds,omitempty" type:"Repeated"`
	// The IDs of the cloud resources with which the Anycast EIPs are associated.
	//
	// You can enter at most 100 cloud resource IDs.
	//
	// example:
	//
	// lb-2zebb08phyczzawe****
	BindInstanceIds []*string `json:"BindInstanceIds,omitempty" xml:"BindInstanceIds,omitempty" type:"Repeated"`
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
	// The billing method of the Anycast EIP.
	//
	// Set the value to **PostPaid**, which specifies the pay-as-you-go billing method.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The metering method of the Anycast EIP.
	//
	// Set the value to **PayByTraffic**, which specifies the pay-by-data-transfer metering method.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The number of entries to return on each page. Valid values: **20 to 100**. Default value: **50**.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the Anycast EIP.
	//
	// The name must be 0 to 128 characters in length, and can contain digits, hyphens (-), and underscores (_). The name must start with a letter.
	//
	// example:
	//
	// doctest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-aekzthsmwsnfuni
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The access area of the Anycast EIP.
	//
	// Set the value to **international**, which specifies the regions outside the Chinese mainland.
	//
	// example:
	//
	// international
	ServiceLocation *string `json:"ServiceLocation,omitempty" xml:"ServiceLocation,omitempty"`
	// The status of the Anycast EIP. Valid values:
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
	// The tags.
	Tags []*ListAnycastEipAddressesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListAnycastEipAddressesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAnycastEipAddressesRequest) GoString() string {
	return s.String()
}

func (s *ListAnycastEipAddressesRequest) GetAnycastEipAddress() *string {
	return s.AnycastEipAddress
}

func (s *ListAnycastEipAddressesRequest) GetAnycastId() *string {
	return s.AnycastId
}

func (s *ListAnycastEipAddressesRequest) GetAnycastIds() []*string {
	return s.AnycastIds
}

func (s *ListAnycastEipAddressesRequest) GetBindInstanceIds() []*string {
	return s.BindInstanceIds
}

func (s *ListAnycastEipAddressesRequest) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *ListAnycastEipAddressesRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *ListAnycastEipAddressesRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *ListAnycastEipAddressesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAnycastEipAddressesRequest) GetName() *string {
	return s.Name
}

func (s *ListAnycastEipAddressesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAnycastEipAddressesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListAnycastEipAddressesRequest) GetServiceLocation() *string {
	return s.ServiceLocation
}

func (s *ListAnycastEipAddressesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAnycastEipAddressesRequest) GetTags() []*ListAnycastEipAddressesRequestTags {
	return s.Tags
}

func (s *ListAnycastEipAddressesRequest) SetAnycastEipAddress(v string) *ListAnycastEipAddressesRequest {
	s.AnycastEipAddress = &v
	return s
}

func (s *ListAnycastEipAddressesRequest) SetAnycastId(v string) *ListAnycastEipAddressesRequest {
	s.AnycastId = &v
	return s
}

func (s *ListAnycastEipAddressesRequest) SetAnycastIds(v []*string) *ListAnycastEipAddressesRequest {
	s.AnycastIds = v
	return s
}

func (s *ListAnycastEipAddressesRequest) SetBindInstanceIds(v []*string) *ListAnycastEipAddressesRequest {
	s.BindInstanceIds = v
	return s
}

func (s *ListAnycastEipAddressesRequest) SetBusinessStatus(v string) *ListAnycastEipAddressesRequest {
	s.BusinessStatus = &v
	return s
}

func (s *ListAnycastEipAddressesRequest) SetInstanceChargeType(v string) *ListAnycastEipAddressesRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *ListAnycastEipAddressesRequest) SetInternetChargeType(v string) *ListAnycastEipAddressesRequest {
	s.InternetChargeType = &v
	return s
}

func (s *ListAnycastEipAddressesRequest) SetMaxResults(v int32) *ListAnycastEipAddressesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAnycastEipAddressesRequest) SetName(v string) *ListAnycastEipAddressesRequest {
	s.Name = &v
	return s
}

func (s *ListAnycastEipAddressesRequest) SetNextToken(v string) *ListAnycastEipAddressesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAnycastEipAddressesRequest) SetResourceGroupId(v string) *ListAnycastEipAddressesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAnycastEipAddressesRequest) SetServiceLocation(v string) *ListAnycastEipAddressesRequest {
	s.ServiceLocation = &v
	return s
}

func (s *ListAnycastEipAddressesRequest) SetStatus(v string) *ListAnycastEipAddressesRequest {
	s.Status = &v
	return s
}

func (s *ListAnycastEipAddressesRequest) SetTags(v []*ListAnycastEipAddressesRequestTags) *ListAnycastEipAddressesRequest {
	s.Tags = v
	return s
}

func (s *ListAnycastEipAddressesRequest) Validate() error {
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

type ListAnycastEipAddressesRequestTags struct {
	// The tag key of the resource. You can specify up to 20 tag keys. You cannot specify empty strings as tag keys.
	//
	// The key can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The key must start with a letter but cannot start with `aliyun` or `acs:`. The key cannot contain `http://` or `https://`.
	//
	// >  You must specify at least one of **Tag.N*	- (**Tag.N.Key*	- and **Tag.N.Value**).
	//
	// example:
	//
	// 1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. It can be an empty string.
	//
	// The value cannot exceed 128 characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). The value must start with a letter but cannot start with `aliyun` or `acs:`. The value cannot contain `http://` or `https://`.
	//
	// >  You must specify at least one of **Tag.N*	- (**Tag.N.Key*	- and **Tag.N.Value**).
	//
	// example:
	//
	// tag1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAnycastEipAddressesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListAnycastEipAddressesRequestTags) GoString() string {
	return s.String()
}

func (s *ListAnycastEipAddressesRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListAnycastEipAddressesRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListAnycastEipAddressesRequestTags) SetKey(v string) *ListAnycastEipAddressesRequestTags {
	s.Key = &v
	return s
}

func (s *ListAnycastEipAddressesRequestTags) SetValue(v string) *ListAnycastEipAddressesRequestTags {
	s.Value = &v
	return s
}

func (s *ListAnycastEipAddressesRequestTags) Validate() error {
	return dara.Validate(s)
}
