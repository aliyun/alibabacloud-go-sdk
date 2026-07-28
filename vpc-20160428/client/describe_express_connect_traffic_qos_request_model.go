// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectTrafficQosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeExpressConnectTrafficQosRequest
	GetClientToken() *string
	SetMaxResults(v int32) *DescribeExpressConnectTrafficQosRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeExpressConnectTrafficQosRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeExpressConnectTrafficQosRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeExpressConnectTrafficQosRequest
	GetOwnerId() *int64
	SetQosIdList(v []*string) *DescribeExpressConnectTrafficQosRequest
	GetQosIdList() []*string
	SetQosNameList(v []*string) *DescribeExpressConnectTrafficQosRequest
	GetQosNameList() []*string
	SetRegionId(v string) *DescribeExpressConnectTrafficQosRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeExpressConnectTrafficQosRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeExpressConnectTrafficQosRequest
	GetResourceOwnerAccount() *string
	SetTags(v []*DescribeExpressConnectTrafficQosRequestTags) *DescribeExpressConnectTrafficQosRequest
	GetTags() []*DescribeExpressConnectTrafficQosRequestTags
}

type DescribeExpressConnectTrafficQosRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may differ for each API request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The number of entries per page for paginated queries. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Valid values:
	//
	// 	- Leave this parameter empty for the first query or if no subsequent query is required.
	//
	// 	- If a next query is to be sent, set the value to the **NextToken*	- value returned in the previous API call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The list of QoS policy IDs.
	QosIdList []*string `json:"QosIdList,omitempty" xml:"QosIdList,omitempty" type:"Repeated"`
	// The list of QoS policy names.
	QosNameList []*string `json:"QosNameList,omitempty" xml:"QosNameList,omitempty" type:"Repeated"`
	// The region ID of the QoS policy.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxazfdgdg****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The tags of the resource.
	Tags []*DescribeExpressConnectTrafficQosRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeExpressConnectTrafficQosRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectTrafficQosRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectTrafficQosRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeExpressConnectTrafficQosRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeExpressConnectTrafficQosRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeExpressConnectTrafficQosRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeExpressConnectTrafficQosRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeExpressConnectTrafficQosRequest) GetQosIdList() []*string {
	return s.QosIdList
}

func (s *DescribeExpressConnectTrafficQosRequest) GetQosNameList() []*string {
	return s.QosNameList
}

func (s *DescribeExpressConnectTrafficQosRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeExpressConnectTrafficQosRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeExpressConnectTrafficQosRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeExpressConnectTrafficQosRequest) GetTags() []*DescribeExpressConnectTrafficQosRequestTags {
	return s.Tags
}

func (s *DescribeExpressConnectTrafficQosRequest) SetClientToken(v string) *DescribeExpressConnectTrafficQosRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRequest) SetMaxResults(v int32) *DescribeExpressConnectTrafficQosRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRequest) SetNextToken(v string) *DescribeExpressConnectTrafficQosRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRequest) SetOwnerAccount(v string) *DescribeExpressConnectTrafficQosRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRequest) SetOwnerId(v int64) *DescribeExpressConnectTrafficQosRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRequest) SetQosIdList(v []*string) *DescribeExpressConnectTrafficQosRequest {
	s.QosIdList = v
	return s
}

func (s *DescribeExpressConnectTrafficQosRequest) SetQosNameList(v []*string) *DescribeExpressConnectTrafficQosRequest {
	s.QosNameList = v
	return s
}

func (s *DescribeExpressConnectTrafficQosRequest) SetRegionId(v string) *DescribeExpressConnectTrafficQosRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRequest) SetResourceGroupId(v string) *DescribeExpressConnectTrafficQosRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRequest) SetResourceOwnerAccount(v string) *DescribeExpressConnectTrafficQosRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRequest) SetTags(v []*DescribeExpressConnectTrafficQosRequestTags) *DescribeExpressConnectTrafficQosRequest {
	s.Tags = v
	return s
}

func (s *DescribeExpressConnectTrafficQosRequest) Validate() error {
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

type DescribeExpressConnectTrafficQosRequestTags struct {
	// The tag key of the resource. You must specify at least 1 tag key and can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// A tag key can be up to 128 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeExpressConnectTrafficQosRequestTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectTrafficQosRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectTrafficQosRequestTags) GetKey() *string {
	return s.Key
}

func (s *DescribeExpressConnectTrafficQosRequestTags) GetValue() *string {
	return s.Value
}

func (s *DescribeExpressConnectTrafficQosRequestTags) SetKey(v string) *DescribeExpressConnectTrafficQosRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRequestTags) SetValue(v string) *DescribeExpressConnectTrafficQosRequestTags {
	s.Value = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRequestTags) Validate() error {
	return dara.Validate(s)
}
