// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuickInitInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *QuickInitInstanceRequest
	GetInstanceId() *string
}

type QuickInitInstanceRequest struct {
	// The ID of the HSM.
	//
	// This parameter is required.
	//
	// example:
	//
	// hsm-cn-mp90fxef****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s QuickInitInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s QuickInitInstanceRequest) GoString() string {
	return s.String()
}

func (s *QuickInitInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QuickInitInstanceRequest) SetInstanceId(v string) *QuickInitInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *QuickInitInstanceRequest) Validate() error {
	return dara.Validate(s)
}
