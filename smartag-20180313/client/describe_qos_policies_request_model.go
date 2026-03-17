// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQosPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DescribeQosPoliciesRequest
	GetDescription() *string
	SetOwnerAccount(v string) *DescribeQosPoliciesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeQosPoliciesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeQosPoliciesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeQosPoliciesRequest
	GetPageSize() *int32
	SetPriority(v int32) *DescribeQosPoliciesRequest
	GetPriority() *int32
	SetQosId(v string) *DescribeQosPoliciesRequest
	GetQosId() *string
	SetQosPolicyId(v string) *DescribeQosPoliciesRequest
	GetQosPolicyId() *string
	SetRegionId(v string) *DescribeQosPoliciesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeQosPoliciesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeQosPoliciesRequest
	GetResourceOwnerId() *int64
}

type DescribeQosPoliciesRequest struct {
	// The description of the 5-tuple.
	//
	// The description must be 1 to 512 characters in length, and can contain digits, underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// docdesc
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: **10**. A maximum of **50*	- entries can be returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The priority of the traffic throttling rule that is applied to the 5-tuple.
	//
	// Valid values: **1 to 3**. A smaller value indicates a higher priority.
	//
	// >  If you have submitted a ticket and created a QoS policy with the priority value 4 by calling the [CreateQosPolicy](https://help.aliyun.com/document_detail/131575.html) operation, you can set the value to 4.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the QoS policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-xitd8690ucu8ro****
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The ID of the 5-tuple.
	//
	// example:
	//
	// qospy-427m9fo6wkh****
	QosPolicyId *string `json:"QosPolicyId,omitempty" xml:"QosPolicyId,omitempty"`
	// The ID of the region to which the QoS policy belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeQosPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeQosPoliciesRequest) GoString() string {
	return s.String()
}

func (s *DescribeQosPoliciesRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeQosPoliciesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeQosPoliciesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeQosPoliciesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeQosPoliciesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeQosPoliciesRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeQosPoliciesRequest) GetQosId() *string {
	return s.QosId
}

func (s *DescribeQosPoliciesRequest) GetQosPolicyId() *string {
	return s.QosPolicyId
}

func (s *DescribeQosPoliciesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeQosPoliciesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeQosPoliciesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeQosPoliciesRequest) SetDescription(v string) *DescribeQosPoliciesRequest {
	s.Description = &v
	return s
}

func (s *DescribeQosPoliciesRequest) SetOwnerAccount(v string) *DescribeQosPoliciesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeQosPoliciesRequest) SetOwnerId(v int64) *DescribeQosPoliciesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeQosPoliciesRequest) SetPageNumber(v int32) *DescribeQosPoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeQosPoliciesRequest) SetPageSize(v int32) *DescribeQosPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeQosPoliciesRequest) SetPriority(v int32) *DescribeQosPoliciesRequest {
	s.Priority = &v
	return s
}

func (s *DescribeQosPoliciesRequest) SetQosId(v string) *DescribeQosPoliciesRequest {
	s.QosId = &v
	return s
}

func (s *DescribeQosPoliciesRequest) SetQosPolicyId(v string) *DescribeQosPoliciesRequest {
	s.QosPolicyId = &v
	return s
}

func (s *DescribeQosPoliciesRequest) SetRegionId(v string) *DescribeQosPoliciesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeQosPoliciesRequest) SetResourceOwnerAccount(v string) *DescribeQosPoliciesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeQosPoliciesRequest) SetResourceOwnerId(v int64) *DescribeQosPoliciesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeQosPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
