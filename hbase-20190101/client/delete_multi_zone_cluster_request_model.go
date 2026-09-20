// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultiZoneClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteMultiZoneClusterRequest
	GetClusterId() *string
	SetImmediateDeleteFlag(v bool) *DeleteMultiZoneClusterRequest
	GetImmediateDeleteFlag() *bool
}

type DeleteMultiZoneClusterRequest struct {
	// The ID of the multi-zone cluster to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-t4nn71xa0yn56****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to immediately delete the instance. By default, the instance is moved to the recycle bin and permanently deleted after 7 days. Valid values:
	//
	// - true: Immediately deletes the instance without moving it to the recycle bin. Use this option with caution.
	//
	// - false: Moves the instance to the recycle bin. This is the default value.
	//
	// example:
	//
	// false
	ImmediateDeleteFlag *bool `json:"ImmediateDeleteFlag,omitempty" xml:"ImmediateDeleteFlag,omitempty"`
}

func (s DeleteMultiZoneClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultiZoneClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteMultiZoneClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteMultiZoneClusterRequest) GetImmediateDeleteFlag() *bool {
	return s.ImmediateDeleteFlag
}

func (s *DeleteMultiZoneClusterRequest) SetClusterId(v string) *DeleteMultiZoneClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteMultiZoneClusterRequest) SetImmediateDeleteFlag(v bool) *DeleteMultiZoneClusterRequest {
	s.ImmediateDeleteFlag = &v
	return s
}

func (s *DeleteMultiZoneClusterRequest) Validate() error {
	return dara.Validate(s)
}
