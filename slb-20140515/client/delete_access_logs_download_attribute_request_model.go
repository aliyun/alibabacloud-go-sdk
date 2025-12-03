// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessLogsDownloadAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerId(v string) *DeleteAccessLogsDownloadAttributeRequest
	GetLoadBalancerId() *string
	SetLogsDownloadAttributes(v string) *DeleteAccessLogsDownloadAttributeRequest
	GetLogsDownloadAttributes() *string
	SetOwnerAccount(v string) *DeleteAccessLogsDownloadAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteAccessLogsDownloadAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteAccessLogsDownloadAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteAccessLogsDownloadAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteAccessLogsDownloadAttributeRequest
	GetResourceOwnerId() *int64
	SetTags(v string) *DeleteAccessLogsDownloadAttributeRequest
	GetTags() *string
}

type DeleteAccessLogsDownloadAttributeRequest struct {
	// The CLB instance ID.
	//
	// example:
	//
	// lb-uf68ps3rekbljmdb0****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The access log forwarding rule. Parameters:
	//
	// 	- **LogProject**: the name of the project of Log Service.
	//
	// 	- **LogStore**: the name of the Logstore of Log Service.
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
	// example:
	//
	// [{"tagKey":"Key1","tagValue":"Value1"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s DeleteAccessLogsDownloadAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessLogsDownloadAttributeRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessLogsDownloadAttributeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DeleteAccessLogsDownloadAttributeRequest) GetLogsDownloadAttributes() *string {
	return s.LogsDownloadAttributes
}

func (s *DeleteAccessLogsDownloadAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteAccessLogsDownloadAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteAccessLogsDownloadAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAccessLogsDownloadAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteAccessLogsDownloadAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteAccessLogsDownloadAttributeRequest) GetTags() *string {
	return s.Tags
}

func (s *DeleteAccessLogsDownloadAttributeRequest) SetLoadBalancerId(v string) *DeleteAccessLogsDownloadAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DeleteAccessLogsDownloadAttributeRequest) SetLogsDownloadAttributes(v string) *DeleteAccessLogsDownloadAttributeRequest {
	s.LogsDownloadAttributes = &v
	return s
}

func (s *DeleteAccessLogsDownloadAttributeRequest) SetOwnerAccount(v string) *DeleteAccessLogsDownloadAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteAccessLogsDownloadAttributeRequest) SetOwnerId(v int64) *DeleteAccessLogsDownloadAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAccessLogsDownloadAttributeRequest) SetRegionId(v string) *DeleteAccessLogsDownloadAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAccessLogsDownloadAttributeRequest) SetResourceOwnerAccount(v string) *DeleteAccessLogsDownloadAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAccessLogsDownloadAttributeRequest) SetResourceOwnerId(v int64) *DeleteAccessLogsDownloadAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAccessLogsDownloadAttributeRequest) SetTags(v string) *DeleteAccessLogsDownloadAttributeRequest {
	s.Tags = &v
	return s
}

func (s *DeleteAccessLogsDownloadAttributeRequest) Validate() error {
	return dara.Validate(s)
}
