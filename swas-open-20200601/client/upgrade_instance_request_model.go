// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpgradeInstanceRequest
	GetClientToken() *string
	SetInstanceId(v string) *UpgradeInstanceRequest
	GetInstanceId() *string
	SetPlanId(v string) *UpgradeInstanceRequest
	GetPlanId() *string
	SetRegionId(v string) *UpgradeInstanceRequest
	GetRegionId() *string
}

type UpgradeInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of ClientToken can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID of the specified simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the plan to which you want to upgrade. You can call [ListPlans](https://help.aliyun.com/document_detail/189314.html) to query plan information.
	//
	// This parameter is required.
	//
	// example:
	//
	// swas.s2.c2m2s50b4t08
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The region ID of the specified simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpgradeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpgradeInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpgradeInstanceRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *UpgradeInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradeInstanceRequest) SetClientToken(v string) *UpgradeInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeInstanceRequest) SetInstanceId(v string) *UpgradeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeInstanceRequest) SetPlanId(v string) *UpgradeInstanceRequest {
	s.PlanId = &v
	return s
}

func (s *UpgradeInstanceRequest) SetRegionId(v string) *UpgradeInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeInstanceRequest) Validate() error {
	return dara.Validate(s)
}
