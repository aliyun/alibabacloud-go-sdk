// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DisableInstanceRequest
	GetInstanceId() *string
}

type DisableInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// da37319b-6c83-4268-9f19-814aed62e401
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableInstanceRequest) GoString() string {
	return s.String()
}

func (s *DisableInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableInstanceRequest) SetInstanceId(v string) *DisableInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableInstanceRequest) Validate() error {
	return dara.Validate(s)
}
