// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelHiveEdgeWorkersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHiveId(v string) *DelHiveEdgeWorkersShrinkRequest
	GetHiveId() *string
	SetInstanceIdsShrink(v string) *DelHiveEdgeWorkersShrinkRequest
	GetInstanceIdsShrink() *string
}

type DelHiveEdgeWorkersShrinkRequest struct {
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

func (s DelHiveEdgeWorkersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DelHiveEdgeWorkersShrinkRequest) GoString() string {
	return s.String()
}

func (s *DelHiveEdgeWorkersShrinkRequest) GetHiveId() *string {
	return s.HiveId
}

func (s *DelHiveEdgeWorkersShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *DelHiveEdgeWorkersShrinkRequest) SetHiveId(v string) *DelHiveEdgeWorkersShrinkRequest {
	s.HiveId = &v
	return s
}

func (s *DelHiveEdgeWorkersShrinkRequest) SetInstanceIdsShrink(v string) *DelHiveEdgeWorkersShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *DelHiveEdgeWorkersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
