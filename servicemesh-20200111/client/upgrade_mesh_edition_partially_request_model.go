// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeMeshEditionPartiallyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetASMGatewayContinue(v bool) *UpgradeMeshEditionPartiallyRequest
	GetASMGatewayContinue() *bool
	SetExpectedVersion(v string) *UpgradeMeshEditionPartiallyRequest
	GetExpectedVersion() *string
	SetPreCheck(v bool) *UpgradeMeshEditionPartiallyRequest
	GetPreCheck() *bool
	SetServiceMeshId(v string) *UpgradeMeshEditionPartiallyRequest
	GetServiceMeshId() *string
	SetSwitchToPro(v bool) *UpgradeMeshEditionPartiallyRequest
	GetSwitchToPro() *bool
	SetUpgradeGatewayRecords(v string) *UpgradeMeshEditionPartiallyRequest
	GetUpgradeGatewayRecords() *string
}

type UpgradeMeshEditionPartiallyRequest struct {
	// Specifies whether to upgrade the ASM gateways for the ASM instance. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	ASMGatewayContinue *bool `json:"ASMGatewayContinue,omitempty" xml:"ASMGatewayContinue,omitempty"`
	// The expected version that desired to be upgraded to.
	//
	// example:
	//
	// v1.15.3.118-g4712daf0-aliyun
	ExpectedVersion *string `json:"ExpectedVersion,omitempty" xml:"ExpectedVersion,omitempty"`
	// Specifies whether to perform an upgrade check. If the value of this parameter is set to true, only the upgrade check is performed and the ASM instance is not upgraded.
	//
	// example:
	//
	// true
	PreCheck *bool `json:"PreCheck,omitempty" xml:"PreCheck,omitempty"`
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca04bc38979214bf2882be79d39b4****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// Deprecated
	//
	// Specifies whether to upgrade the ASM instance to Professional Edition. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	SwitchToPro *bool `json:"SwitchToPro,omitempty" xml:"SwitchToPro,omitempty"`
	// Specifies the ASM gateways to be upgraded. Separate multiple ASM gateways with commas (,).
	//
	// example:
	//
	// ingressgateway1,ingressgateway2
	UpgradeGatewayRecords *string `json:"UpgradeGatewayRecords,omitempty" xml:"UpgradeGatewayRecords,omitempty"`
}

func (s UpgradeMeshEditionPartiallyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeMeshEditionPartiallyRequest) GoString() string {
	return s.String()
}

func (s *UpgradeMeshEditionPartiallyRequest) GetASMGatewayContinue() *bool {
	return s.ASMGatewayContinue
}

func (s *UpgradeMeshEditionPartiallyRequest) GetExpectedVersion() *string {
	return s.ExpectedVersion
}

func (s *UpgradeMeshEditionPartiallyRequest) GetPreCheck() *bool {
	return s.PreCheck
}

func (s *UpgradeMeshEditionPartiallyRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *UpgradeMeshEditionPartiallyRequest) GetSwitchToPro() *bool {
	return s.SwitchToPro
}

func (s *UpgradeMeshEditionPartiallyRequest) GetUpgradeGatewayRecords() *string {
	return s.UpgradeGatewayRecords
}

func (s *UpgradeMeshEditionPartiallyRequest) SetASMGatewayContinue(v bool) *UpgradeMeshEditionPartiallyRequest {
	s.ASMGatewayContinue = &v
	return s
}

func (s *UpgradeMeshEditionPartiallyRequest) SetExpectedVersion(v string) *UpgradeMeshEditionPartiallyRequest {
	s.ExpectedVersion = &v
	return s
}

func (s *UpgradeMeshEditionPartiallyRequest) SetPreCheck(v bool) *UpgradeMeshEditionPartiallyRequest {
	s.PreCheck = &v
	return s
}

func (s *UpgradeMeshEditionPartiallyRequest) SetServiceMeshId(v string) *UpgradeMeshEditionPartiallyRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpgradeMeshEditionPartiallyRequest) SetSwitchToPro(v bool) *UpgradeMeshEditionPartiallyRequest {
	s.SwitchToPro = &v
	return s
}

func (s *UpgradeMeshEditionPartiallyRequest) SetUpgradeGatewayRecords(v string) *UpgradeMeshEditionPartiallyRequest {
	s.UpgradeGatewayRecords = &v
	return s
}

func (s *UpgradeMeshEditionPartiallyRequest) Validate() error {
	return dara.Validate(s)
}
