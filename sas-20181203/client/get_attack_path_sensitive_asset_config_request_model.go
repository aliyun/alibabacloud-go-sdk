// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttackPathSensitiveAssetConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttackPathSensitiveAssetConfigId(v string) *GetAttackPathSensitiveAssetConfigRequest
	GetAttackPathSensitiveAssetConfigId() *string
	SetConfigType(v string) *GetAttackPathSensitiveAssetConfigRequest
	GetConfigType() *string
}

type GetAttackPathSensitiveAssetConfigRequest struct {
	// ID of the created attack path sensitive asset setting.
	//
	// example:
	//
	// apsac-123
	AttackPathSensitiveAssetConfigId *string `json:"AttackPathSensitiveAssetConfigId,omitempty" xml:"AttackPathSensitiveAssetConfigId,omitempty"`
	// Configuration type. Possible values:
	//
	// - asset_instance: Asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// asset_instance
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
}

func (s GetAttackPathSensitiveAssetConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAttackPathSensitiveAssetConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAttackPathSensitiveAssetConfigRequest) GetAttackPathSensitiveAssetConfigId() *string {
	return s.AttackPathSensitiveAssetConfigId
}

func (s *GetAttackPathSensitiveAssetConfigRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetAttackPathSensitiveAssetConfigRequest) SetAttackPathSensitiveAssetConfigId(v string) *GetAttackPathSensitiveAssetConfigRequest {
	s.AttackPathSensitiveAssetConfigId = &v
	return s
}

func (s *GetAttackPathSensitiveAssetConfigRequest) SetConfigType(v string) *GetAttackPathSensitiveAssetConfigRequest {
	s.ConfigType = &v
	return s
}

func (s *GetAttackPathSensitiveAssetConfigRequest) Validate() error {
	return dara.Validate(s)
}
