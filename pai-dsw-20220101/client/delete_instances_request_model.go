// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *DeleteInstancesRequest
	GetInstanceIds() []*string
}

type DeleteInstancesRequest struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DeleteInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstancesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DeleteInstancesRequest) SetInstanceIds(v []*string) *DeleteInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *DeleteInstancesRequest) Validate() error {
	return dara.Validate(s)
}
