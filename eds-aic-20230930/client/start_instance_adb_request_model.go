// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstanceAdbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *StartInstanceAdbRequest
	GetInstanceIds() []*string
}

type StartInstanceAdbRequest struct {
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s StartInstanceAdbRequest) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceAdbRequest) GoString() string {
	return s.String()
}

func (s *StartInstanceAdbRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *StartInstanceAdbRequest) SetInstanceIds(v []*string) *StartInstanceAdbRequest {
	s.InstanceIds = v
	return s
}

func (s *StartInstanceAdbRequest) Validate() error {
	return dara.Validate(s)
}
