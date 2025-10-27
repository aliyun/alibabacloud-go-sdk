// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAssetCleanConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetCleanConfigs(v []*ModifyAssetCleanConfigRequestAssetCleanConfigs) *ModifyAssetCleanConfigRequest
	GetAssetCleanConfigs() []*ModifyAssetCleanConfigRequestAssetCleanConfigs
}

type ModifyAssetCleanConfigRequest struct {
	// The asset cleanup configurations.
	AssetCleanConfigs []*ModifyAssetCleanConfigRequestAssetCleanConfigs `json:"AssetCleanConfigs,omitempty" xml:"AssetCleanConfigs,omitempty" type:"Repeated"`
}

func (s ModifyAssetCleanConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAssetCleanConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyAssetCleanConfigRequest) GetAssetCleanConfigs() []*ModifyAssetCleanConfigRequestAssetCleanConfigs {
	return s.AssetCleanConfigs
}

func (s *ModifyAssetCleanConfigRequest) SetAssetCleanConfigs(v []*ModifyAssetCleanConfigRequestAssetCleanConfigs) *ModifyAssetCleanConfigRequest {
	s.AssetCleanConfigs = v
	return s
}

func (s *ModifyAssetCleanConfigRequest) Validate() error {
	if s.AssetCleanConfigs != nil {
		for _, item := range s.AssetCleanConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyAssetCleanConfigRequestAssetCleanConfigs struct {
	// The number of days before hosts whose provider cannot be identified are automatically cleaned after they enter the offline state. Valid value: an integer that ranges from 1 to 30.
	//
	// example:
	//
	// 7
	CleanDays *int32 `json:"CleanDays,omitempty" xml:"CleanDays,omitempty"`
	// Specifies whether to enable the feature of cleaning the offline hosts whose provider cannot be identified. Valid values:
	//
	// 	- **0**: disables the feature.
	//
	// 	- **1**: enables the feature.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of hosts that you want to clean.
	//
	// Set the value to **1**, which indicates hosts whose provider cannot be identified.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyAssetCleanConfigRequestAssetCleanConfigs) String() string {
	return dara.Prettify(s)
}

func (s ModifyAssetCleanConfigRequestAssetCleanConfigs) GoString() string {
	return s.String()
}

func (s *ModifyAssetCleanConfigRequestAssetCleanConfigs) GetCleanDays() *int32 {
	return s.CleanDays
}

func (s *ModifyAssetCleanConfigRequestAssetCleanConfigs) GetStatus() *int32 {
	return s.Status
}

func (s *ModifyAssetCleanConfigRequestAssetCleanConfigs) GetType() *int32 {
	return s.Type
}

func (s *ModifyAssetCleanConfigRequestAssetCleanConfigs) SetCleanDays(v int32) *ModifyAssetCleanConfigRequestAssetCleanConfigs {
	s.CleanDays = &v
	return s
}

func (s *ModifyAssetCleanConfigRequestAssetCleanConfigs) SetStatus(v int32) *ModifyAssetCleanConfigRequestAssetCleanConfigs {
	s.Status = &v
	return s
}

func (s *ModifyAssetCleanConfigRequestAssetCleanConfigs) SetType(v int32) *ModifyAssetCleanConfigRequestAssetCleanConfigs {
	s.Type = &v
	return s
}

func (s *ModifyAssetCleanConfigRequestAssetCleanConfigs) Validate() error {
	return dara.Validate(s)
}
