// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ResumeInstanceRequest
	GetInstanceId() *string
}

type ResumeInstanceRequest struct {
	// The ID of the HSM.
	//
	// This parameter is required.
	//
	// example:
	//
	// hsm-cn-vj30bil8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ResumeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeInstanceRequest) GoString() string {
	return s.String()
}

func (s *ResumeInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResumeInstanceRequest) SetInstanceId(v string) *ResumeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ResumeInstanceRequest) Validate() error {
	return dara.Validate(s)
}
