// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrQuitLaneCmd interface {
	dara.Model
	String() string
	GoString() string
	SetLaneIds(v string) *AddOrQuitLaneCmd
	GetLaneIds() *string
	SetOperateType(v string) *AddOrQuitLaneCmd
	GetOperateType() *string
	SetServiceGroupId(v int64) *AddOrQuitLaneCmd
	GetServiceGroupId() *int64
}

type AddOrQuitLaneCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	LaneIds *string `json:"laneIds,omitempty" xml:"laneIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// add
	OperateType *string `json:"operateType,omitempty" xml:"operateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ServiceGroupId *int64 `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
}

func (s AddOrQuitLaneCmd) String() string {
	return dara.Prettify(s)
}

func (s AddOrQuitLaneCmd) GoString() string {
	return s.String()
}

func (s *AddOrQuitLaneCmd) GetLaneIds() *string {
	return s.LaneIds
}

func (s *AddOrQuitLaneCmd) GetOperateType() *string {
	return s.OperateType
}

func (s *AddOrQuitLaneCmd) GetServiceGroupId() *int64 {
	return s.ServiceGroupId
}

func (s *AddOrQuitLaneCmd) SetLaneIds(v string) *AddOrQuitLaneCmd {
	s.LaneIds = &v
	return s
}

func (s *AddOrQuitLaneCmd) SetOperateType(v string) *AddOrQuitLaneCmd {
	s.OperateType = &v
	return s
}

func (s *AddOrQuitLaneCmd) SetServiceGroupId(v int64) *AddOrQuitLaneCmd {
	s.ServiceGroupId = &v
	return s
}

func (s *AddOrQuitLaneCmd) Validate() error {
	return dara.Validate(s)
}
