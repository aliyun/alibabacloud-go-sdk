// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddHiveEdgeWorkersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHiveId(v string) *AddHiveEdgeWorkersRequest
	GetHiveId() *string
	SetInstanceIds(v []*string) *AddHiveEdgeWorkersRequest
	GetInstanceIds() []*string
}

type AddHiveEdgeWorkersRequest struct {
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
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s AddHiveEdgeWorkersRequest) String() string {
	return dara.Prettify(s)
}

func (s AddHiveEdgeWorkersRequest) GoString() string {
	return s.String()
}

func (s *AddHiveEdgeWorkersRequest) GetHiveId() *string {
	return s.HiveId
}

func (s *AddHiveEdgeWorkersRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *AddHiveEdgeWorkersRequest) SetHiveId(v string) *AddHiveEdgeWorkersRequest {
	s.HiveId = &v
	return s
}

func (s *AddHiveEdgeWorkersRequest) SetInstanceIds(v []*string) *AddHiveEdgeWorkersRequest {
	s.InstanceIds = v
	return s
}

func (s *AddHiveEdgeWorkersRequest) Validate() error {
	return dara.Validate(s)
}
