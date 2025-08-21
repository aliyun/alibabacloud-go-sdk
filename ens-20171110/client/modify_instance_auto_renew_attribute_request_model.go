// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAutoRenewAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v string) *ModifyInstanceAutoRenewAttributeRequest
	GetAutoRenew() *string
	SetDuration(v string) *ModifyInstanceAutoRenewAttributeRequest
	GetDuration() *string
	SetInstanceIds(v string) *ModifyInstanceAutoRenewAttributeRequest
	GetInstanceIds() *string
	SetOwnerId(v string) *ModifyInstanceAutoRenewAttributeRequest
	GetOwnerId() *string
	SetRenewalStatus(v string) *ModifyInstanceAutoRenewAttributeRequest
	GetRenewalStatus() *string
}

type ModifyInstanceAutoRenewAttributeRequest struct {
	// Specifies whether to enable the auto-renewal feature. Valid values: **True and False**. Default value: False.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period of the instance. Unit: months. Valid values: 1 to 9 and 12. This parameter is required if the AutoRenew parameter is set to true.
	//
	// example:
	//
	// 12
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The IDs of the instances. Separate IDs with semicolons (;).
	//
	// This parameter is required.
	//
	// example:
	//
	// instance-test
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to renew the instance. The **RenewalStatus*	- parameter has a higher priority than the **AutoRenew*	- parameter. If you do not specify **RenewalStatus**, the **AutoRenew*	- parameter is used by default.
	//
	// 	- AutoRenewal: Auto-renewal is enabled for the instance.
	//
	// 	- Normal: Auto-renewal is disabled for the instance.
	//
	// 	- NotRenewal: The instance is not renewed.
	//
	// The system no longer sends an expiration notification but sends only a renewal notification three days before the instance expires. To renew the instance, you can change the value of this parameter from NotRenewal to Normal and then manually renew the instance, or change the value of this parameter from NotRenewal to AutoRenewal.
	//
	// example:
	//
	// Normal
	RenewalStatus *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
}

func (s ModifyInstanceAutoRenewAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAutoRenewAttributeRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *ModifyInstanceAutoRenewAttributeRequest) GetDuration() *string {
	return s.Duration
}

func (s *ModifyInstanceAutoRenewAttributeRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *ModifyInstanceAutoRenewAttributeRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ModifyInstanceAutoRenewAttributeRequest) GetRenewalStatus() *string {
	return s.RenewalStatus
}

func (s *ModifyInstanceAutoRenewAttributeRequest) SetAutoRenew(v string) *ModifyInstanceAutoRenewAttributeRequest {
	s.AutoRenew = &v
	return s
}

func (s *ModifyInstanceAutoRenewAttributeRequest) SetDuration(v string) *ModifyInstanceAutoRenewAttributeRequest {
	s.Duration = &v
	return s
}

func (s *ModifyInstanceAutoRenewAttributeRequest) SetInstanceIds(v string) *ModifyInstanceAutoRenewAttributeRequest {
	s.InstanceIds = &v
	return s
}

func (s *ModifyInstanceAutoRenewAttributeRequest) SetOwnerId(v string) *ModifyInstanceAutoRenewAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceAutoRenewAttributeRequest) SetRenewalStatus(v string) *ModifyInstanceAutoRenewAttributeRequest {
	s.RenewalStatus = &v
	return s
}

func (s *ModifyInstanceAutoRenewAttributeRequest) Validate() error {
	return dara.Validate(s)
}
