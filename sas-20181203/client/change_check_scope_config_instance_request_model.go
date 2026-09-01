// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeCheckScopeConfigInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddAssetUuids(v []*string) *ChangeCheckScopeConfigInstanceRequest
	GetAddAssetUuids() []*string
	SetConfigId(v string) *ChangeCheckScopeConfigInstanceRequest
	GetConfigId() *string
	SetDeleteAssetUuids(v []*string) *ChangeCheckScopeConfigInstanceRequest
	GetDeleteAssetUuids() []*string
}

type ChangeCheckScopeConfigInstanceRequest struct {
	// The list of unique IDs of cloud assets to add.
	AddAssetUuids []*string `json:"AddAssetUuids,omitempty" xml:"AddAssetUuids,omitempty" type:"Repeated"`
	// The ID of the scan scope configuration.
	//
	// >Call the [GetCheckScopeConfig](~~GetCheckScopeConfig~~) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00cfa8161da093089e6804ba6a33****
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The list of unique IDs of cloud assets to delete.
	DeleteAssetUuids []*string `json:"DeleteAssetUuids,omitempty" xml:"DeleteAssetUuids,omitempty" type:"Repeated"`
}

func (s ChangeCheckScopeConfigInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeCheckScopeConfigInstanceRequest) GoString() string {
	return s.String()
}

func (s *ChangeCheckScopeConfigInstanceRequest) GetAddAssetUuids() []*string {
	return s.AddAssetUuids
}

func (s *ChangeCheckScopeConfigInstanceRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *ChangeCheckScopeConfigInstanceRequest) GetDeleteAssetUuids() []*string {
	return s.DeleteAssetUuids
}

func (s *ChangeCheckScopeConfigInstanceRequest) SetAddAssetUuids(v []*string) *ChangeCheckScopeConfigInstanceRequest {
	s.AddAssetUuids = v
	return s
}

func (s *ChangeCheckScopeConfigInstanceRequest) SetConfigId(v string) *ChangeCheckScopeConfigInstanceRequest {
	s.ConfigId = &v
	return s
}

func (s *ChangeCheckScopeConfigInstanceRequest) SetDeleteAssetUuids(v []*string) *ChangeCheckScopeConfigInstanceRequest {
	s.DeleteAssetUuids = v
	return s
}

func (s *ChangeCheckScopeConfigInstanceRequest) Validate() error {
	return dara.Validate(s)
}
