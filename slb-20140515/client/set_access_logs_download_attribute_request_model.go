// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAccessLogsDownloadAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerId(v string) *SetAccessLogsDownloadAttributeRequest
	GetLoadBalancerId() *string
	SetLogsDownloadAttributes(v string) *SetAccessLogsDownloadAttributeRequest
	GetLogsDownloadAttributes() *string
	SetOwnerAccount(v string) *SetAccessLogsDownloadAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SetAccessLogsDownloadAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetAccessLogsDownloadAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *SetAccessLogsDownloadAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetAccessLogsDownloadAttributeRequest
	GetResourceOwnerId() *int64
	SetTags(v string) *SetAccessLogsDownloadAttributeRequest
	GetTags() *string
}

type SetAccessLogsDownloadAttributeRequest struct {
	// The ID of the CLB instance.
	//
	// example:
	//
	// lb-bp1b6c719dfa08ex*****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The access log forwarding rule. Parameters:
	//
	// 	- **LogProject**: the name of the project of Simple Log Service.
	//
	// 	- **LogStore**: the name of the Logstore of Simple Log Service.
	//
	// 	- **LoadBalancerId**: the ID of the CLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"logProject":"my-project", "LogStore":"my-log-store", "LoadBalancerId":"lb-uf68ps3rekbljmdb0****"}]
	LogsDownloadAttributes *string `json:"LogsDownloadAttributes,omitempty" xml:"LogsDownloadAttributes,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the CLB instance.
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
	// The tags that are added to the CLB instance. The tags must be key-value pairs that are contained in a JSON dictionary.
	//
	// You can specify up to 10 tags in each call.
	//
	// example:
	//
	// [{"tagKey":"Key1","tagValue":"Value1"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s SetAccessLogsDownloadAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAccessLogsDownloadAttributeRequest) GoString() string {
	return s.String()
}

func (s *SetAccessLogsDownloadAttributeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *SetAccessLogsDownloadAttributeRequest) GetLogsDownloadAttributes() *string {
	return s.LogsDownloadAttributes
}

func (s *SetAccessLogsDownloadAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetAccessLogsDownloadAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetAccessLogsDownloadAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetAccessLogsDownloadAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetAccessLogsDownloadAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetAccessLogsDownloadAttributeRequest) GetTags() *string {
	return s.Tags
}

func (s *SetAccessLogsDownloadAttributeRequest) SetLoadBalancerId(v string) *SetAccessLogsDownloadAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetAccessLogsDownloadAttributeRequest) SetLogsDownloadAttributes(v string) *SetAccessLogsDownloadAttributeRequest {
	s.LogsDownloadAttributes = &v
	return s
}

func (s *SetAccessLogsDownloadAttributeRequest) SetOwnerAccount(v string) *SetAccessLogsDownloadAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetAccessLogsDownloadAttributeRequest) SetOwnerId(v int64) *SetAccessLogsDownloadAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *SetAccessLogsDownloadAttributeRequest) SetRegionId(v string) *SetAccessLogsDownloadAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *SetAccessLogsDownloadAttributeRequest) SetResourceOwnerAccount(v string) *SetAccessLogsDownloadAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetAccessLogsDownloadAttributeRequest) SetResourceOwnerId(v int64) *SetAccessLogsDownloadAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetAccessLogsDownloadAttributeRequest) SetTags(v string) *SetAccessLogsDownloadAttributeRequest {
	s.Tags = &v
	return s
}

func (s *SetAccessLogsDownloadAttributeRequest) Validate() error {
	return dara.Validate(s)
}
