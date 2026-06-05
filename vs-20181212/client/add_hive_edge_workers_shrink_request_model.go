// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddHiveEdgeWorkersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHiveId(v string) *AddHiveEdgeWorkersShrinkRequest
	GetHiveId() *string
	SetInstanceIdsShrink(v string) *AddHiveEdgeWorkersShrinkRequest
	GetInstanceIdsShrink() *string
}

type AddHiveEdgeWorkersShrinkRequest struct {
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
	// [\\"acp-c4b9pbj4fzkagfexv\\", \\"acp-c4b9pbj4fzkagfexw\\"]
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s AddHiveEdgeWorkersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddHiveEdgeWorkersShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddHiveEdgeWorkersShrinkRequest) GetHiveId() *string {
	return s.HiveId
}

func (s *AddHiveEdgeWorkersShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *AddHiveEdgeWorkersShrinkRequest) SetHiveId(v string) *AddHiveEdgeWorkersShrinkRequest {
	s.HiveId = &v
	return s
}

func (s *AddHiveEdgeWorkersShrinkRequest) SetInstanceIdsShrink(v string) *AddHiveEdgeWorkersShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *AddHiveEdgeWorkersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
