// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInstanceForDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CheckInstanceForDeleteRequest
	GetInstanceId() *string
}

type CheckInstanceForDeleteRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CheckInstanceForDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceForDeleteRequest) GoString() string {
	return s.String()
}

func (s *CheckInstanceForDeleteRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CheckInstanceForDeleteRequest) SetInstanceId(v string) *CheckInstanceForDeleteRequest {
	s.InstanceId = &v
	return s
}

func (s *CheckInstanceForDeleteRequest) Validate() error {
	return dara.Validate(s)
}
