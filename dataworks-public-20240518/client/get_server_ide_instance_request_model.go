// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServerIdeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetServerIdeInstanceRequest
	GetInstanceId() *string
}

type GetServerIdeInstanceRequest struct {
	// The personal development environment instance ID. You can call ListServerIdeInstances to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 699573
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetServerIdeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServerIdeInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetServerIdeInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetServerIdeInstanceRequest) SetInstanceId(v string) *GetServerIdeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetServerIdeInstanceRequest) Validate() error {
	return dara.Validate(s)
}
