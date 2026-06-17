// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCfwInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyCfwInstanceRequest
	GetInstanceId() *string
	SetUpdateList(v []*ModifyCfwInstanceRequestUpdateList) *ModifyCfwInstanceRequest
	GetUpdateList() []*ModifyCfwInstanceRequestUpdateList
}

type ModifyCfwInstanceRequest struct {
	// The ID of the Cloud Firewall instance.
	//
	// example:
	//
	// cfw_elasticity_public_cn-zsk39m******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// A list of instance properties to update.
	UpdateList []*ModifyCfwInstanceRequestUpdateList `json:"UpdateList,omitempty" xml:"UpdateList,omitempty" type:"Repeated"`
}

func (s ModifyCfwInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCfwInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyCfwInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyCfwInstanceRequest) GetUpdateList() []*ModifyCfwInstanceRequestUpdateList {
	return s.UpdateList
}

func (s *ModifyCfwInstanceRequest) SetInstanceId(v string) *ModifyCfwInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyCfwInstanceRequest) SetUpdateList(v []*ModifyCfwInstanceRequestUpdateList) *ModifyCfwInstanceRequest {
	s.UpdateList = v
	return s
}

func (s *ModifyCfwInstanceRequest) Validate() error {
	if s.UpdateList != nil {
		for _, item := range s.UpdateList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyCfwInstanceRequestUpdateList struct {
	// The code of the instance property to update.
	//
	// The following codes are supported:
	//
	// - \\`Code\\`: \\`MajorVersion\\`. Set \\`Value\\` to \\`2\\`. This is available only for pay-as-you-go 1.0 users to upgrade their instances to pay-as-you-go 2.0.
	//
	//   	Warning:
	//
	//   Make sure you understand the billing methods and pricing of pay-as-you-go 2.0.
	//
	//
	//
	//   	Warning:
	//
	//   Note that if log delivery is enabled before the upgrade, it will remain enabled after the upgrade, and logs will be delivered to a new Logstore.
	//
	//
	//
	// - \\`Code\\`: \\`ThreatIntelligence\\`. This is available only for pay-as-you-go 2.0 users to enable or disable the threat intelligence feature. Set \\`Value\\` to \\`1\\` to enable the feature or \\`0\\` to disable it.
	//
	// - \\`Code\\`: \\`Sdl\\`. This is available only for pay-as-you-go 2.0 users to enable or disable the sensitive data leak detection feature. Set \\`Value\\` to \\`1\\` to enable the feature or \\`0\\` to disable it.
	//
	// example:
	//
	// Sdl
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The value for the specified \\`Code\\`. For valid values, see the description of the \\`Code\\` parameter.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyCfwInstanceRequestUpdateList) String() string {
	return dara.Prettify(s)
}

func (s ModifyCfwInstanceRequestUpdateList) GoString() string {
	return s.String()
}

func (s *ModifyCfwInstanceRequestUpdateList) GetCode() *string {
	return s.Code
}

func (s *ModifyCfwInstanceRequestUpdateList) GetValue() *string {
	return s.Value
}

func (s *ModifyCfwInstanceRequestUpdateList) SetCode(v string) *ModifyCfwInstanceRequestUpdateList {
	s.Code = &v
	return s
}

func (s *ModifyCfwInstanceRequestUpdateList) SetValue(v string) *ModifyCfwInstanceRequestUpdateList {
	s.Value = &v
	return s
}

func (s *ModifyCfwInstanceRequestUpdateList) Validate() error {
	return dara.Validate(s)
}
