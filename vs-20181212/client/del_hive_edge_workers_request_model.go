// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelHiveEdgeWorkersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHiveId(v string) *DelHiveEdgeWorkersRequest
	GetHiveId() *string
	SetInstanceIds(v []*string) *DelHiveEdgeWorkersRequest
	GetInstanceIds() []*string
}

type DelHiveEdgeWorkersRequest struct {
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

func (s DelHiveEdgeWorkersRequest) String() string {
	return dara.Prettify(s)
}

func (s DelHiveEdgeWorkersRequest) GoString() string {
	return s.String()
}

func (s *DelHiveEdgeWorkersRequest) GetHiveId() *string {
	return s.HiveId
}

func (s *DelHiveEdgeWorkersRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DelHiveEdgeWorkersRequest) SetHiveId(v string) *DelHiveEdgeWorkersRequest {
	s.HiveId = &v
	return s
}

func (s *DelHiveEdgeWorkersRequest) SetInstanceIds(v []*string) *DelHiveEdgeWorkersRequest {
	s.InstanceIds = v
	return s
}

func (s *DelHiveEdgeWorkersRequest) Validate() error {
	return dara.Validate(s)
}
