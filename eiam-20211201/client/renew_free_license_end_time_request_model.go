// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewFreeLicenseEndTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RenewFreeLicenseEndTimeRequest
	GetInstanceId() *string
}

type RenewFreeLicenseEndTimeRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RenewFreeLicenseEndTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewFreeLicenseEndTimeRequest) GoString() string {
	return s.String()
}

func (s *RenewFreeLicenseEndTimeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RenewFreeLicenseEndTimeRequest) SetInstanceId(v string) *RenewFreeLicenseEndTimeRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewFreeLicenseEndTimeRequest) Validate() error {
	return dara.Validate(s)
}
