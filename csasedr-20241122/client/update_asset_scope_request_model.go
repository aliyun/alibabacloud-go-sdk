// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAssetScopeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *UpdateAssetScopeRequest
	GetId() *int64
	SetMatchMode(v string) *UpdateAssetScopeRequest
	GetMatchMode() *string
	SetMatchTargetIds(v []*string) *UpdateAssetScopeRequest
	GetMatchTargetIds() []*string
}

type UpdateAssetScopeRequest struct {
	// The asset scope ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1932
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The effective scope mode. Valid values:
	//
	// - **UserGroupAll**: All user groups.
	//
	// - **UserGroupNormal**: Specified user groups.
	//
	// - **DeviceGroupNormal**: Specified device groups.
	//
	// - **DevTagNormal**: Specified device tags.
	//
	// This field determines the meaning of MatchTargetIds.
	//
	// This parameter is required.
	//
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The IDs of the user groups, device tags, or tag groups.
	MatchTargetIds []*string `json:"MatchTargetIds,omitempty" xml:"MatchTargetIds,omitempty" type:"Repeated"`
}

func (s UpdateAssetScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAssetScopeRequest) GoString() string {
	return s.String()
}

func (s *UpdateAssetScopeRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateAssetScopeRequest) GetMatchMode() *string {
	return s.MatchMode
}

func (s *UpdateAssetScopeRequest) GetMatchTargetIds() []*string {
	return s.MatchTargetIds
}

func (s *UpdateAssetScopeRequest) SetId(v int64) *UpdateAssetScopeRequest {
	s.Id = &v
	return s
}

func (s *UpdateAssetScopeRequest) SetMatchMode(v string) *UpdateAssetScopeRequest {
	s.MatchMode = &v
	return s
}

func (s *UpdateAssetScopeRequest) SetMatchTargetIds(v []*string) *UpdateAssetScopeRequest {
	s.MatchTargetIds = v
	return s
}

func (s *UpdateAssetScopeRequest) Validate() error {
	return dara.Validate(s)
}
