// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAttachmentAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivatePoolOptions(v *ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions) *ModifyInstanceAttachmentAttributesRequest
	GetPrivatePoolOptions() *ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions
	SetInstanceId(v string) *ModifyInstanceAttachmentAttributesRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ModifyInstanceAttachmentAttributesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceAttachmentAttributesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyInstanceAttachmentAttributesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyInstanceAttachmentAttributesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceAttachmentAttributesRequest
	GetResourceOwnerId() *int64
}

type ModifyInstanceAttachmentAttributesRequest struct {
	PrivatePoolOptions *ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	// The instance ID of the instance for which you want to modify the private pool matching property.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the private pool. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
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

func (s ModifyInstanceAttachmentAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAttachmentAttributesRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttachmentAttributesRequest) GetPrivatePoolOptions() *ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *ModifyInstanceAttachmentAttributesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceAttachmentAttributesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceAttachmentAttributesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceAttachmentAttributesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceAttachmentAttributesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceAttachmentAttributesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceAttachmentAttributesRequest) SetPrivatePoolOptions(v *ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions) *ModifyInstanceAttachmentAttributesRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *ModifyInstanceAttachmentAttributesRequest) SetInstanceId(v string) *ModifyInstanceAttachmentAttributesRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceAttachmentAttributesRequest) SetOwnerAccount(v string) *ModifyInstanceAttachmentAttributesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceAttachmentAttributesRequest) SetOwnerId(v int64) *ModifyInstanceAttachmentAttributesRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceAttachmentAttributesRequest) SetRegionId(v string) *ModifyInstanceAttachmentAttributesRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceAttachmentAttributesRequest) SetResourceOwnerAccount(v string) *ModifyInstanceAttachmentAttributesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceAttachmentAttributesRequest) SetResourceOwnerId(v int64) *ModifyInstanceAttachmentAttributesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceAttachmentAttributesRequest) Validate() error {
	if s.PrivatePoolOptions != nil {
		if err := s.PrivatePoolOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions struct {
	// The private pool ID, which is the elasticity assurance ID or capacity reservation ID.
	//
	// - This parameter is required when PrivatePoolOptions.MatchCriteria is set to `Target`.
	//
	// - Leave this parameter empty when PrivatePoolOptions.MatchCriteria is set to `Open` or `None`.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The private pool matching mode of the instance. Valid values:
	//
	// - Open: open mode. The system automatically matches the instance with an open private pool. If no matching private pool capacity is available, public pool resources are used to launch the instance.
	//
	// - Target: targeted mode. The instance is launched by using the capacity of the specified private pool. If the specified private pool capacity is unavailable, the instance fails to be launched. If you set this parameter to Target, you must also specify the PrivatePoolOptions.Id parameter to specify the private pool ID.
	//
	// - None: none. The instance is launched normally without using a private pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// Open
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
}

func (s ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions) GetId() *string {
	return s.Id
}

func (s *ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions) GetMatchCriteria() *string {
	return s.MatchCriteria
}

func (s *ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions) SetId(v string) *ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions) SetMatchCriteria(v string) *ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}
