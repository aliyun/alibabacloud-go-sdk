// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServerIdeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteServerIdeInstanceRequest
	GetInstanceId() *string
}

type DeleteServerIdeInstanceRequest struct {
	// The ID of the personal development environment instance. You can call ListServerIdeInstances to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 699573
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteServerIdeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServerIdeInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteServerIdeInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteServerIdeInstanceRequest) SetInstanceId(v string) *DeleteServerIdeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteServerIdeInstanceRequest) Validate() error {
	return dara.Validate(s)
}
