// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyJVSInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyToAll(v bool) *ModifyJVSInstanceRequest
	GetApplyToAll() *bool
	SetCreditConfig(v []*ModifyJVSInstanceRequestCreditConfig) *ModifyJVSInstanceRequest
	GetCreditConfig() []*ModifyJVSInstanceRequestCreditConfig
	SetInstanceIds(v []*string) *ModifyJVSInstanceRequest
	GetInstanceIds() []*string
	SetInstanceName(v string) *ModifyJVSInstanceRequest
	GetInstanceName() *string
}

type ModifyJVSInstanceRequest struct {
	// example:
	//
	// true
	ApplyToAll   *bool                                   `json:"ApplyToAll,omitempty" xml:"ApplyToAll,omitempty"`
	CreditConfig []*ModifyJVSInstanceRequestCreditConfig `json:"CreditConfig,omitempty" xml:"CreditConfig,omitempty" type:"Repeated"`
	InstanceIds  []*string                               `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// defaultInstanceName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s ModifyJVSInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyJVSInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyJVSInstanceRequest) GetApplyToAll() *bool {
	return s.ApplyToAll
}

func (s *ModifyJVSInstanceRequest) GetCreditConfig() []*ModifyJVSInstanceRequestCreditConfig {
	return s.CreditConfig
}

func (s *ModifyJVSInstanceRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ModifyJVSInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyJVSInstanceRequest) SetApplyToAll(v bool) *ModifyJVSInstanceRequest {
	s.ApplyToAll = &v
	return s
}

func (s *ModifyJVSInstanceRequest) SetCreditConfig(v []*ModifyJVSInstanceRequestCreditConfig) *ModifyJVSInstanceRequest {
	s.CreditConfig = v
	return s
}

func (s *ModifyJVSInstanceRequest) SetInstanceIds(v []*string) *ModifyJVSInstanceRequest {
	s.InstanceIds = v
	return s
}

func (s *ModifyJVSInstanceRequest) SetInstanceName(v string) *ModifyJVSInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyJVSInstanceRequest) Validate() error {
	if s.CreditConfig != nil {
		for _, item := range s.CreditConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyJVSInstanceRequestCreditConfig struct {
	// example:
	//
	// 10
	CreditLimit *int64 `json:"CreditLimit,omitempty" xml:"CreditLimit,omitempty"`
	// example:
	//
	// day
	LimitPeriod *string `json:"LimitPeriod,omitempty" xml:"LimitPeriod,omitempty"`
}

func (s ModifyJVSInstanceRequestCreditConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyJVSInstanceRequestCreditConfig) GoString() string {
	return s.String()
}

func (s *ModifyJVSInstanceRequestCreditConfig) GetCreditLimit() *int64 {
	return s.CreditLimit
}

func (s *ModifyJVSInstanceRequestCreditConfig) GetLimitPeriod() *string {
	return s.LimitPeriod
}

func (s *ModifyJVSInstanceRequestCreditConfig) SetCreditLimit(v int64) *ModifyJVSInstanceRequestCreditConfig {
	s.CreditLimit = &v
	return s
}

func (s *ModifyJVSInstanceRequestCreditConfig) SetLimitPeriod(v string) *ModifyJVSInstanceRequestCreditConfig {
	s.LimitPeriod = &v
	return s
}

func (s *ModifyJVSInstanceRequestCreditConfig) Validate() error {
	return dara.Validate(s)
}
