// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopServerIdeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StopServerIdeInstanceRequest
	GetInstanceId() *string
}

type StopServerIdeInstanceRequest struct {
	// The ID of the personal development environment instance. You can call ListServerIdeInstances to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 699573
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StopServerIdeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StopServerIdeInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopServerIdeInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StopServerIdeInstanceRequest) SetInstanceId(v string) *StopServerIdeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StopServerIdeInstanceRequest) Validate() error {
	return dara.Validate(s)
}
