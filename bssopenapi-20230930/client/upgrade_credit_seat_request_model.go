// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeCreditSeatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpgradeCreditSeatRequest
	GetClientToken() *string
	SetConfigs(v []*UpgradeCreditSeatRequestConfigs) *UpgradeCreditSeatRequest
	GetConfigs() []*UpgradeCreditSeatRequestConfigs
	SetInstanceId(v string) *UpgradeCreditSeatRequest
	GetInstanceId() *string
	SetProductCode(v string) *UpgradeCreditSeatRequest
	GetProductCode() *string
	SetProductType(v string) *UpgradeCreditSeatRequest
	GetProductType() *string
	SetSubscriptionType(v string) *UpgradeCreditSeatRequest
	GetSubscriptionType() *string
}

type UpgradeCreditSeatRequest struct {
	ClientToken      *string                            `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Configs          []*UpgradeCreditSeatRequestConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	InstanceId       *string                            `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ProductCode      *string                            `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType      *string                            `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType *string                            `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s UpgradeCreditSeatRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeCreditSeatRequest) GoString() string {
	return s.String()
}

func (s *UpgradeCreditSeatRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpgradeCreditSeatRequest) GetConfigs() []*UpgradeCreditSeatRequestConfigs {
	return s.Configs
}

func (s *UpgradeCreditSeatRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpgradeCreditSeatRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *UpgradeCreditSeatRequest) GetProductType() *string {
	return s.ProductType
}

func (s *UpgradeCreditSeatRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *UpgradeCreditSeatRequest) SetClientToken(v string) *UpgradeCreditSeatRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeCreditSeatRequest) SetConfigs(v []*UpgradeCreditSeatRequestConfigs) *UpgradeCreditSeatRequest {
	s.Configs = v
	return s
}

func (s *UpgradeCreditSeatRequest) SetInstanceId(v string) *UpgradeCreditSeatRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeCreditSeatRequest) SetProductCode(v string) *UpgradeCreditSeatRequest {
	s.ProductCode = &v
	return s
}

func (s *UpgradeCreditSeatRequest) SetProductType(v string) *UpgradeCreditSeatRequest {
	s.ProductType = &v
	return s
}

func (s *UpgradeCreditSeatRequest) SetSubscriptionType(v string) *UpgradeCreditSeatRequest {
	s.SubscriptionType = &v
	return s
}

func (s *UpgradeCreditSeatRequest) Validate() error {
	if s.Configs != nil {
		for _, item := range s.Configs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpgradeCreditSeatRequestConfigs struct {
	Code  *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpgradeCreditSeatRequestConfigs) String() string {
	return dara.Prettify(s)
}

func (s UpgradeCreditSeatRequestConfigs) GoString() string {
	return s.String()
}

func (s *UpgradeCreditSeatRequestConfigs) GetCode() *string {
	return s.Code
}

func (s *UpgradeCreditSeatRequestConfigs) GetValue() *string {
	return s.Value
}

func (s *UpgradeCreditSeatRequestConfigs) SetCode(v string) *UpgradeCreditSeatRequestConfigs {
	s.Code = &v
	return s
}

func (s *UpgradeCreditSeatRequestConfigs) SetValue(v string) *UpgradeCreditSeatRequestConfigs {
	s.Value = &v
	return s
}

func (s *UpgradeCreditSeatRequestConfigs) Validate() error {
	return dara.Validate(s)
}
