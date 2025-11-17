// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstanceAdbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *StopInstanceAdbRequest
	GetInstanceIds() []*string
}

type StopInstanceAdbRequest struct {
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s StopInstanceAdbRequest) String() string {
	return dara.Prettify(s)
}

func (s StopInstanceAdbRequest) GoString() string {
	return s.String()
}

func (s *StopInstanceAdbRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *StopInstanceAdbRequest) SetInstanceIds(v []*string) *StopInstanceAdbRequest {
	s.InstanceIds = v
	return s
}

func (s *StopInstanceAdbRequest) Validate() error {
	return dara.Validate(s)
}
