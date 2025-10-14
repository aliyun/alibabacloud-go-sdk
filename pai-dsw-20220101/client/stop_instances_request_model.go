// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *StopInstancesRequest
	GetInstanceIds() []*string
}

type StopInstancesRequest struct {
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s StopInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s StopInstancesRequest) GoString() string {
	return s.String()
}

func (s *StopInstancesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *StopInstancesRequest) SetInstanceIds(v []*string) *StopInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *StopInstancesRequest) Validate() error {
	return dara.Validate(s)
}
