// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveHiveEdgeWorkersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHiveId(v string) *MoveHiveEdgeWorkersShrinkRequest
	GetHiveId() *string
	SetInstanceIdsShrink(v string) *MoveHiveEdgeWorkersShrinkRequest
	GetInstanceIdsShrink() *string
}

type MoveHiveEdgeWorkersShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hive-xxxx-xxx-xxx
	HiveId *string `json:"HiveId,omitempty" xml:"HiveId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ["ew-xxxx"]
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s MoveHiveEdgeWorkersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveHiveEdgeWorkersShrinkRequest) GoString() string {
	return s.String()
}

func (s *MoveHiveEdgeWorkersShrinkRequest) GetHiveId() *string {
	return s.HiveId
}

func (s *MoveHiveEdgeWorkersShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *MoveHiveEdgeWorkersShrinkRequest) SetHiveId(v string) *MoveHiveEdgeWorkersShrinkRequest {
	s.HiveId = &v
	return s
}

func (s *MoveHiveEdgeWorkersShrinkRequest) SetInstanceIdsShrink(v string) *MoveHiveEdgeWorkersShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *MoveHiveEdgeWorkersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
