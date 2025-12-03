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
	// This parameter is required.
	//
	// example:
	//
	// d-t4nn71xa0yn56****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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
