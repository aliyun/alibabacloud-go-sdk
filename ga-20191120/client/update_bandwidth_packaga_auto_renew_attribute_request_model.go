// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBandwidthPackagaAutoRenewAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *UpdateBandwidthPackagaAutoRenewAttributeRequest
	GetAutoRenew() *bool
	SetAutoRenewDuration(v int32) *UpdateBandwidthPackagaAutoRenewAttributeRequest
	GetAutoRenewDuration() *int32
	SetClientToken(v string) *UpdateBandwidthPackagaAutoRenewAttributeRequest
	GetClientToken() *string
	SetInstanceId(v string) *UpdateBandwidthPackagaAutoRenewAttributeRequest
	GetInstanceId() *string
	SetName(v string) *UpdateBandwidthPackagaAutoRenewAttributeRequest
	GetName() *string
	SetRegionId(v string) *UpdateBandwidthPackagaAutoRenewAttributeRequest
	GetRegionId() *string
	SetRenewalStatus(v string) *UpdateBandwidthPackagaAutoRenewAttributeRequest
	GetRenewalStatus() *string
}

type UpdateBandwidthPackagaAutoRenewAttributeRequest struct {
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// >  You must specify **AutoRenew*	- or **RenewalStatus**.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal duration. Unit: months. Valid values: **1*	- to **12**.
	//
	// > This parameter takes effect only if **AutoRenew*	- is set to **true**.
	//
	// example:
	//
	// 1
	AutoRenewDuration *int32 `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the bandwidth plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// gbwp-bp1sgzldyj6b4q7cx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the bandwidth plan.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The auto-renewal status of the bandwidth plan. Valid values:
	//
	// 	- **AutoRenewal**: The bandwidth plan is automatically renewed.
	//
	// 	- **Normal**: You must manually renew the bandwidth plan.
	//
	// 	- **NotRenewal**: The bandwidth plan is not renewed after it expires. The system sends only a non-renewal reminder three days before the expiration date. To renew a bandwidth plan for which you set RenewalStatus to NotRenewal, you can change the value of RenewalStatus from NotRenewal to Normal, and then manually renew the bandwidth plan. You can also set RenewalStatus to **AutoRenewal**.
	//
	// > 	- You must specify **AutoRenew*	- or **RenewalStatus**.
	//
	// > 	- **RenewalStatus*	- takes precedence over **AutoRenew**. If you do not specify **RenewalStatus**, **AutoRenew*	- is used.
	//
	// example:
	//
	// Normal
	RenewalStatus *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
}

func (s UpdateBandwidthPackagaAutoRenewAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBandwidthPackagaAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeRequest) GetAutoRenewDuration() *int32 {
	return s.AutoRenewDuration
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeRequest) GetName() *string {
	return s.Name
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeRequest) GetRenewalStatus() *string {
	return s.RenewalStatus
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeRequest) SetAutoRenew(v bool) *UpdateBandwidthPackagaAutoRenewAttributeRequest {
	s.AutoRenew = &v
	return s
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeRequest) SetAutoRenewDuration(v int32) *UpdateBandwidthPackagaAutoRenewAttributeRequest {
	s.AutoRenewDuration = &v
	return s
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeRequest) SetClientToken(v string) *UpdateBandwidthPackagaAutoRenewAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeRequest) SetInstanceId(v string) *UpdateBandwidthPackagaAutoRenewAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeRequest) SetName(v string) *UpdateBandwidthPackagaAutoRenewAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeRequest) SetRegionId(v string) *UpdateBandwidthPackagaAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeRequest) SetRenewalStatus(v string) *UpdateBandwidthPackagaAutoRenewAttributeRequest {
	s.RenewalStatus = &v
	return s
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeRequest) Validate() error {
	return dara.Validate(s)
}
