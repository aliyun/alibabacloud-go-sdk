// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAttackPathSensitiveAssetConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttackPathSensitiveAssetConfigId(v string) *DeleteAttackPathSensitiveAssetConfigRequest
	GetAttackPathSensitiveAssetConfigId() *string
}

type DeleteAttackPathSensitiveAssetConfigRequest struct {
	// ID of the attack path sensitive asset configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// apsac-123
	AttackPathSensitiveAssetConfigId *string `json:"AttackPathSensitiveAssetConfigId,omitempty" xml:"AttackPathSensitiveAssetConfigId,omitempty"`
}

func (s DeleteAttackPathSensitiveAssetConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAttackPathSensitiveAssetConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteAttackPathSensitiveAssetConfigRequest) GetAttackPathSensitiveAssetConfigId() *string {
	return s.AttackPathSensitiveAssetConfigId
}

func (s *DeleteAttackPathSensitiveAssetConfigRequest) SetAttackPathSensitiveAssetConfigId(v string) *DeleteAttackPathSensitiveAssetConfigRequest {
	s.AttackPathSensitiveAssetConfigId = &v
	return s
}

func (s *DeleteAttackPathSensitiveAssetConfigRequest) Validate() error {
	return dara.Validate(s)
}
