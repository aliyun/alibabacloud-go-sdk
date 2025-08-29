// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrafficControlTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetTrafficControlTargetRequest
	GetInstanceId() *string
}

type GetTrafficControlTargetRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetTrafficControlTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTrafficControlTargetRequest) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTargetRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTrafficControlTargetRequest) SetInstanceId(v string) *GetTrafficControlTargetRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTrafficControlTargetRequest) Validate() error {
	return dara.Validate(s)
}
