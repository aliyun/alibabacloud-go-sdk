// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *StartInstancesRequest
	GetInstanceIds() []*string
}

type StartInstancesRequest struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s StartInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s StartInstancesRequest) GoString() string {
	return s.String()
}

func (s *StartInstancesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *StartInstancesRequest) SetInstanceIds(v []*string) *StartInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *StartInstancesRequest) Validate() error {
	return dara.Validate(s)
}
