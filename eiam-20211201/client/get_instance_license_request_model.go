// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInstanceLicenseRequest
	GetInstanceId() *string
}

type GetInstanceLicenseRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk2676xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetInstanceLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceLicenseRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceLicenseRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceLicenseRequest) SetInstanceId(v string) *GetInstanceLicenseRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceLicenseRequest) Validate() error {
	return dara.Validate(s)
}
