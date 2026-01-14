// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAcceleratorAutoRenewAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *UpdateAcceleratorAutoRenewAttributeRequest
	GetAcceleratorId() *string
	SetAutoRenew(v bool) *UpdateAcceleratorAutoRenewAttributeRequest
	GetAutoRenew() *bool
	SetAutoRenewDuration(v int32) *UpdateAcceleratorAutoRenewAttributeRequest
	GetAutoRenewDuration() *int32
	SetClientToken(v string) *UpdateAcceleratorAutoRenewAttributeRequest
	GetClientToken() *string
	SetName(v string) *UpdateAcceleratorAutoRenewAttributeRequest
	GetName() *string
	SetRegionId(v string) *UpdateAcceleratorAutoRenewAttributeRequest
	GetRegionId() *string
	SetRenewalStatus(v string) *UpdateAcceleratorAutoRenewAttributeRequest
	GetRenewalStatus() *string
}

type UpdateAcceleratorAutoRenewAttributeRequest struct {
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// Specifies whether to enable auto-renewal for the GA instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// >  **AutoRenew*	- and **RenewalStatus*	- cannot be left empty at the same time.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal duration. Unit: month.
	//
	// Valid values: **1*	- to **12**.
	//
	// >  This parameter takes effect only if you set **AutoRenew*	- to **true**.
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
	// The name of the GA instance.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies how to renew the GA instance. Valid values:
	//
	// 	- **AutoRenewal**: The system automatically renews the GA instance.
	//
	// 	- **Normal**: You must manually renew the GA instance.
	//
	// 	- **NotRenewal**: The GA instance is not renewed after the instance expires. The system sends only a non-renewal reminder three days before the expiration date. The system no longer reminds you to renew the GA instance. To renew a GA instance whose RenewalStatus is set to NotRenewal, change the value of RenewalStatus from NotRenewal to **Normal**, and then manually renew the instance. You can also set RenewalStatus to **AutoRenewal**.
	//
	// >
	//
	// 	- **AutoRenew*	- and **RenewalStatus*	- cannot be left empty at the same time.
	//
	// 	- **RenewalStatus*	- takes precedence over **AutoRenew**. By default, if you do not specify **RenewalStatus**, **AutoRenew*	- is used.
	//
	// example:
	//
	// Normal
	RenewalStatus *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
}

func (s UpdateAcceleratorAutoRenewAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAcceleratorAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateAcceleratorAutoRenewAttributeRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *UpdateAcceleratorAutoRenewAttributeRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *UpdateAcceleratorAutoRenewAttributeRequest) GetAutoRenewDuration() *int32 {
	return s.AutoRenewDuration
}

func (s *UpdateAcceleratorAutoRenewAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAcceleratorAutoRenewAttributeRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAcceleratorAutoRenewAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAcceleratorAutoRenewAttributeRequest) GetRenewalStatus() *string {
	return s.RenewalStatus
}

func (s *UpdateAcceleratorAutoRenewAttributeRequest) SetAcceleratorId(v string) *UpdateAcceleratorAutoRenewAttributeRequest {
	s.AcceleratorId = &v
	return s
}

func (s *UpdateAcceleratorAutoRenewAttributeRequest) SetAutoRenew(v bool) *UpdateAcceleratorAutoRenewAttributeRequest {
	s.AutoRenew = &v
	return s
}

func (s *UpdateAcceleratorAutoRenewAttributeRequest) SetAutoRenewDuration(v int32) *UpdateAcceleratorAutoRenewAttributeRequest {
	s.AutoRenewDuration = &v
	return s
}

func (s *UpdateAcceleratorAutoRenewAttributeRequest) SetClientToken(v string) *UpdateAcceleratorAutoRenewAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAcceleratorAutoRenewAttributeRequest) SetName(v string) *UpdateAcceleratorAutoRenewAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateAcceleratorAutoRenewAttributeRequest) SetRegionId(v string) *UpdateAcceleratorAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAcceleratorAutoRenewAttributeRequest) SetRenewalStatus(v string) *UpdateAcceleratorAutoRenewAttributeRequest {
	s.RenewalStatus = &v
	return s
}

func (s *UpdateAcceleratorAutoRenewAttributeRequest) Validate() error {
	return dara.Validate(s)
}
