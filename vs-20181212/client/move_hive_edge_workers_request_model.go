// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveHiveEdgeWorkersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHiveId(v string) *MoveHiveEdgeWorkersRequest
	GetHiveId() *string
	SetInstanceIds(v []*string) *MoveHiveEdgeWorkersRequest
	GetInstanceIds() []*string
}

type MoveHiveEdgeWorkersRequest struct {
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
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s MoveHiveEdgeWorkersRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveHiveEdgeWorkersRequest) GoString() string {
	return s.String()
}

func (s *MoveHiveEdgeWorkersRequest) GetHiveId() *string {
	return s.HiveId
}

func (s *MoveHiveEdgeWorkersRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *MoveHiveEdgeWorkersRequest) SetHiveId(v string) *MoveHiveEdgeWorkersRequest {
	s.HiveId = &v
	return s
}

func (s *MoveHiveEdgeWorkersRequest) SetInstanceIds(v []*string) *MoveHiveEdgeWorkersRequest {
	s.InstanceIds = v
	return s
}

func (s *MoveHiveEdgeWorkersRequest) Validate() error {
	return dara.Validate(s)
}
