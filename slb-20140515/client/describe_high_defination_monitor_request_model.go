// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHighDefinationMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeHighDefinationMonitorRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeHighDefinationMonitorRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeHighDefinationMonitorRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeHighDefinationMonitorRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeHighDefinationMonitorRequest
	GetResourceOwnerId() *int64
	SetTags(v string) *DescribeHighDefinationMonitorRequest
	GetTags() *string
}

type DescribeHighDefinationMonitorRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where you want to query the configuration of fine-grained monitoring.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags of the logs. The tags must be key-value pairs that are contained in a JSON dictionary.
	//
	// example:
	//
	// [{"tagKey":"Key1","tagValue":"Value1"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s DescribeHighDefinationMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighDefinationMonitorRequest) GoString() string {
	return s.String()
}

func (s *DescribeHighDefinationMonitorRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeHighDefinationMonitorRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeHighDefinationMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHighDefinationMonitorRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeHighDefinationMonitorRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeHighDefinationMonitorRequest) GetTags() *string {
	return s.Tags
}

func (s *DescribeHighDefinationMonitorRequest) SetOwnerAccount(v string) *DescribeHighDefinationMonitorRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeHighDefinationMonitorRequest) SetOwnerId(v int64) *DescribeHighDefinationMonitorRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHighDefinationMonitorRequest) SetRegionId(v string) *DescribeHighDefinationMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHighDefinationMonitorRequest) SetResourceOwnerAccount(v string) *DescribeHighDefinationMonitorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHighDefinationMonitorRequest) SetResourceOwnerId(v int64) *DescribeHighDefinationMonitorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeHighDefinationMonitorRequest) SetTags(v string) *DescribeHighDefinationMonitorRequest {
	s.Tags = &v
	return s
}

func (s *DescribeHighDefinationMonitorRequest) Validate() error {
	return dara.Validate(s)
}
