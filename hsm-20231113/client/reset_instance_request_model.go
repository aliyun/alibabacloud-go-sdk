// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ResetInstanceRequest
	GetInstanceId() *string
}

type ResetInstanceRequest struct {
	// The ID of the HSM.
	//
	// This parameter is required.
	//
	// example:
	//
	// hsm-cn-vj30bil8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ResetInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetInstanceRequest) GoString() string {
	return s.String()
}

func (s *ResetInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResetInstanceRequest) SetInstanceId(v string) *ResetInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetInstanceRequest) Validate() error {
	return dara.Validate(s)
}
