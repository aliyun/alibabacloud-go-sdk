// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartServerIdeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StartServerIdeInstanceRequest
	GetInstanceId() *string
}

type StartServerIdeInstanceRequest struct {
	// The ID of the personal development environment instance. You can call ListServerIdeInstances to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 699573
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StartServerIdeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartServerIdeInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartServerIdeInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartServerIdeInstanceRequest) SetInstanceId(v string) *StartServerIdeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StartServerIdeInstanceRequest) Validate() error {
	return dara.Validate(s)
}
