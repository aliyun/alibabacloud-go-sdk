// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExpressConnectTrafficQosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateExpressConnectTrafficQosRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *CreateExpressConnectTrafficQosRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateExpressConnectTrafficQosRequest
	GetOwnerId() *int64
	SetQosDescription(v string) *CreateExpressConnectTrafficQosRequest
	GetQosDescription() *string
	SetQosName(v string) *CreateExpressConnectTrafficQosRequest
	GetQosName() *string
	SetRegionId(v string) *CreateExpressConnectTrafficQosRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateExpressConnectTrafficQosRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateExpressConnectTrafficQosRequest
	GetResourceOwnerAccount() *string
	SetTags(v []*CreateExpressConnectTrafficQosRequestTags) *CreateExpressConnectTrafficQosRequest
	GetTags() []*CreateExpressConnectTrafficQosRequestTags
}

type CreateExpressConnectTrafficQosRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- in each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The description of the QoS policy.
	//
	// The description must be 0 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// qos-test
	QosDescription *string `json:"QosDescription,omitempty" xml:"QosDescription,omitempty"`
	// The name of the QoS policy.
	//
	// The name must be 0 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// qos-test
	QosName *string `json:"QosName,omitempty" xml:"QosName,omitempty"`
	// The region ID of the QoS policy.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxazfdgdg****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The tag to add to the resource.
	Tags []*CreateExpressConnectTrafficQosRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateExpressConnectTrafficQosRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExpressConnectTrafficQosRequest) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectTrafficQosRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateExpressConnectTrafficQosRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateExpressConnectTrafficQosRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateExpressConnectTrafficQosRequest) GetQosDescription() *string {
	return s.QosDescription
}

func (s *CreateExpressConnectTrafficQosRequest) GetQosName() *string {
	return s.QosName
}

func (s *CreateExpressConnectTrafficQosRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateExpressConnectTrafficQosRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateExpressConnectTrafficQosRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateExpressConnectTrafficQosRequest) GetTags() []*CreateExpressConnectTrafficQosRequestTags {
	return s.Tags
}

func (s *CreateExpressConnectTrafficQosRequest) SetClientToken(v string) *CreateExpressConnectTrafficQosRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRequest) SetOwnerAccount(v string) *CreateExpressConnectTrafficQosRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRequest) SetOwnerId(v int64) *CreateExpressConnectTrafficQosRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRequest) SetQosDescription(v string) *CreateExpressConnectTrafficQosRequest {
	s.QosDescription = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRequest) SetQosName(v string) *CreateExpressConnectTrafficQosRequest {
	s.QosName = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRequest) SetRegionId(v string) *CreateExpressConnectTrafficQosRequest {
	s.RegionId = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRequest) SetResourceGroupId(v string) *CreateExpressConnectTrafficQosRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRequest) SetResourceOwnerAccount(v string) *CreateExpressConnectTrafficQosRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRequest) SetTags(v []*CreateExpressConnectTrafficQosRequestTags) *CreateExpressConnectTrafficQosRequest {
	s.Tags = v
	return s
}

func (s *CreateExpressConnectTrafficQosRequest) Validate() error {
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

type CreateExpressConnectTrafficQosRequestTags struct {
	// The tag key to add to the resource. You must enter at least one tag key. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// A tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value to add to the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateExpressConnectTrafficQosRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateExpressConnectTrafficQosRequestTags) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectTrafficQosRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateExpressConnectTrafficQosRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateExpressConnectTrafficQosRequestTags) SetKey(v string) *CreateExpressConnectTrafficQosRequestTags {
	s.Key = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRequestTags) SetValue(v string) *CreateExpressConnectTrafficQosRequestTags {
	s.Value = &v
	return s
}

func (s *CreateExpressConnectTrafficQosRequestTags) Validate() error {
	return dara.Validate(s)
}
