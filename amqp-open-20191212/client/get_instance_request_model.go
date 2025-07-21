// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInstanceRequest
	GetInstanceId() *string
}

type GetInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-v0h1kb9nu***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceRequest) SetInstanceId(v string) *GetInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceRequest) Validate() error {
	return dara.Validate(s)
}
