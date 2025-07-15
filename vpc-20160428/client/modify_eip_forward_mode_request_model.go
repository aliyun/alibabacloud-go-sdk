// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEipForwardModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyEipForwardModeRequest
	GetClientToken() *string
	SetInstanceId(v string) *ModifyEipForwardModeRequest
	GetInstanceId() *string
	SetMode(v string) *ModifyEipForwardModeRequest
	GetMode() *string
	SetOwnerId(v int64) *ModifyEipForwardModeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyEipForwardModeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyEipForwardModeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyEipForwardModeRequest
	GetResourceOwnerId() *int64
}

type ModifyEipForwardModeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **RequestId*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the EIP whose attributes you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// eip-j5ebhbw3br92fy****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The association mode. Valid values:
	//
	// 	- **NAT*	- (default): the standard NAT mode.
	//
	// 	- **MULTI_BINDED**: the multi-EIP-to-ENI mode.
	//
	// 	- **BINDED**: the cut-through mode.
	//
	// >  This parameter is required only if **InstanceType*	- is set to **NetworkInterface**.
	//
	// This parameter is required.
	//
	// example:
	//
	// BINDED
	Mode    *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the EIP belongs. You can call the DescribeRegions operation to query the most recent region list.
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

func (s ModifyEipForwardModeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyEipForwardModeRequest) GoString() string {
	return s.String()
}

func (s *ModifyEipForwardModeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyEipForwardModeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyEipForwardModeRequest) GetMode() *string {
	return s.Mode
}

func (s *ModifyEipForwardModeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyEipForwardModeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyEipForwardModeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyEipForwardModeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyEipForwardModeRequest) SetClientToken(v string) *ModifyEipForwardModeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyEipForwardModeRequest) SetInstanceId(v string) *ModifyEipForwardModeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyEipForwardModeRequest) SetMode(v string) *ModifyEipForwardModeRequest {
	s.Mode = &v
	return s
}

func (s *ModifyEipForwardModeRequest) SetOwnerId(v int64) *ModifyEipForwardModeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyEipForwardModeRequest) SetRegionId(v string) *ModifyEipForwardModeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyEipForwardModeRequest) SetResourceOwnerAccount(v string) *ModifyEipForwardModeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyEipForwardModeRequest) SetResourceOwnerId(v int64) *ModifyEipForwardModeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyEipForwardModeRequest) Validate() error {
	return dara.Validate(s)
}
