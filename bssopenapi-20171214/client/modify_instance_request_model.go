// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyInstanceRequest
	GetClientToken() *string
	SetInstanceId(v string) *ModifyInstanceRequest
	GetInstanceId() *string
	SetModifyType(v string) *ModifyInstanceRequest
	GetModifyType() *string
	SetOwnerId(v int64) *ModifyInstanceRequest
	GetOwnerId() *int64
	SetParameter(v []*ModifyInstanceRequestParameter) *ModifyInstanceRequest
	GetParameter() []*ModifyInstanceRequestParameter
	SetProductCode(v string) *ModifyInstanceRequest
	GetProductCode() *string
	SetProductType(v string) *ModifyInstanceRequest
	GetProductType() *string
	SetSubscriptionType(v string) *ModifyInstanceRequest
	GetSubscriptionType() *string
}

type ModifyInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests.
	//
	// example:
	//
	// JAKSJFHFAKJSF
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the instance for which you want to modify the configurations.
	//
	// example:
	//
	// rm-akjhkdsjhfskjfhd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of configuration modifications. Valid values:
	//
	// 	- Upgrade: upgrades the configurations of the instance.
	//
	// 	- Downgrade: downgrades the configurations of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// Upgrade
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The details about the parameters.
	Parameter []*ModifyInstanceRequestParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Repeated"`
	// The code of the service to which the instance belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The type of the service to which the instance belongs.
	//
	// example:
	//
	// rds
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The billing method. Valid values:
	//
	// 	- Subscription: subscription
	//
	// 	- PayAsYouGo: pay-as-you-go
	//
	// This parameter is required.
	//
	// example:
	//
	// Subscription
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s ModifyInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceRequest) GetModifyType() *string {
	return s.ModifyType
}

func (s *ModifyInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceRequest) GetParameter() []*ModifyInstanceRequestParameter {
	return s.Parameter
}

func (s *ModifyInstanceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ModifyInstanceRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ModifyInstanceRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *ModifyInstanceRequest) SetClientToken(v string) *ModifyInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyInstanceRequest) SetInstanceId(v string) *ModifyInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceRequest) SetModifyType(v string) *ModifyInstanceRequest {
	s.ModifyType = &v
	return s
}

func (s *ModifyInstanceRequest) SetOwnerId(v int64) *ModifyInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceRequest) SetParameter(v []*ModifyInstanceRequestParameter) *ModifyInstanceRequest {
	s.Parameter = v
	return s
}

func (s *ModifyInstanceRequest) SetProductCode(v string) *ModifyInstanceRequest {
	s.ProductCode = &v
	return s
}

func (s *ModifyInstanceRequest) SetProductType(v string) *ModifyInstanceRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyInstanceRequest) SetSubscriptionType(v string) *ModifyInstanceRequest {
	s.SubscriptionType = &v
	return s
}

func (s *ModifyInstanceRequest) Validate() error {
	if s.Parameter != nil {
		for _, item := range s.Parameter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyInstanceRequestParameter struct {
	// The code of the parameter n. Valid values of n: 1 to 100. Multiple parameters are concatenated in the order of n.
	//
	// >  Only the parameters of the attributes that you want to modify for the instance must be configured. For example, if the instance has Attribute A and Attribute B and only Attribute A must be modified, configure only the parameter of Attribute A.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The value of the parameter n. Valid values of n: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyInstanceRequestParameter) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceRequestParameter) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRequestParameter) GetCode() *string {
	return s.Code
}

func (s *ModifyInstanceRequestParameter) GetValue() *string {
	return s.Value
}

func (s *ModifyInstanceRequestParameter) SetCode(v string) *ModifyInstanceRequestParameter {
	s.Code = &v
	return s
}

func (s *ModifyInstanceRequestParameter) SetValue(v string) *ModifyInstanceRequestParameter {
	s.Value = &v
	return s
}

func (s *ModifyInstanceRequestParameter) Validate() error {
	return dara.Validate(s)
}
