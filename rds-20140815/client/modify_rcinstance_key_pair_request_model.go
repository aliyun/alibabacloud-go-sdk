// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCInstanceKeyPairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyRCInstanceKeyPairRequest
	GetInstanceId() *string
	SetKeyPairName(v string) *ModifyRCInstanceKeyPairRequest
	GetKeyPairName() *string
	SetReboot(v bool) *ModifyRCInstanceKeyPairRequest
	GetReboot() *bool
	SetRegionId(v string) *ModifyRCInstanceKeyPairRequest
	GetRegionId() *string
}

type ModifyRCInstanceKeyPairRequest struct {
	// The instance ID.
	//
	// example:
	//
	// rc-m5sc1271fv344a1r****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// customer_keypairs
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// Specifies whether to restart the instance.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// ture
	Reboot *bool `json:"Reboot,omitempty" xml:"Reboot,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyRCInstanceKeyPairRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCInstanceKeyPairRequest) GoString() string {
	return s.String()
}

func (s *ModifyRCInstanceKeyPairRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyRCInstanceKeyPairRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *ModifyRCInstanceKeyPairRequest) GetReboot() *bool {
	return s.Reboot
}

func (s *ModifyRCInstanceKeyPairRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRCInstanceKeyPairRequest) SetInstanceId(v string) *ModifyRCInstanceKeyPairRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyRCInstanceKeyPairRequest) SetKeyPairName(v string) *ModifyRCInstanceKeyPairRequest {
	s.KeyPairName = &v
	return s
}

func (s *ModifyRCInstanceKeyPairRequest) SetReboot(v bool) *ModifyRCInstanceKeyPairRequest {
	s.Reboot = &v
	return s
}

func (s *ModifyRCInstanceKeyPairRequest) SetRegionId(v string) *ModifyRCInstanceKeyPairRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRCInstanceKeyPairRequest) Validate() error {
	return dara.Validate(s)
}
