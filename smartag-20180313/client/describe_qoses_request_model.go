// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQosesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeQosesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeQosesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeQosesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeQosesRequest
	GetPageSize() *int32
	SetQosIds(v string) *DescribeQosesRequest
	GetQosIds() *string
	SetQosName(v string) *DescribeQosesRequest
	GetQosName() *string
	SetRegionId(v string) *DescribeQosesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeQosesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeQosesRequest
	GetResourceOwnerId() *int64
}

type DescribeQosesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**. Maximum value: **50**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the QoS policy.
	//
	// 	- Separate multiple QoS polices with commas (,).
	//
	// 	- If you do not specify this parameter, all QoS policies in the region are queried.
	//
	// example:
	//
	// qos-oek3r2cmvk7m8q****
	QosIds *string `json:"QosIds,omitempty" xml:"QosIds,omitempty"`
	// The name of the QoS policy.
	//
	// The name must be 2 to 100 characters in length and can contain letters, digits, periods (.), underscores (_),and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// zxtest
	QosName *string `json:"QosName,omitempty" xml:"QosName,omitempty"`
	// The ID of the region where the QoS policy is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeQosesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeQosesRequest) GoString() string {
	return s.String()
}

func (s *DescribeQosesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeQosesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeQosesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeQosesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeQosesRequest) GetQosIds() *string {
	return s.QosIds
}

func (s *DescribeQosesRequest) GetQosName() *string {
	return s.QosName
}

func (s *DescribeQosesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeQosesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeQosesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeQosesRequest) SetOwnerAccount(v string) *DescribeQosesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeQosesRequest) SetOwnerId(v int64) *DescribeQosesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeQosesRequest) SetPageNumber(v int32) *DescribeQosesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeQosesRequest) SetPageSize(v int32) *DescribeQosesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeQosesRequest) SetQosIds(v string) *DescribeQosesRequest {
	s.QosIds = &v
	return s
}

func (s *DescribeQosesRequest) SetQosName(v string) *DescribeQosesRequest {
	s.QosName = &v
	return s
}

func (s *DescribeQosesRequest) SetRegionId(v string) *DescribeQosesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeQosesRequest) SetResourceOwnerAccount(v string) *DescribeQosesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeQosesRequest) SetResourceOwnerId(v int64) *DescribeQosesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeQosesRequest) Validate() error {
	return dara.Validate(s)
}
