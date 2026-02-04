// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceTrialLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateInstanceTrialLicenseRequest
	GetInstanceId() *string
}

type CreateInstanceTrialLicenseRequest struct {
	// IDaaS EIAM的实例id
	//
	// This parameter is required.
	//
	// example:
	//
	// eiam-111ccc1111
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateInstanceTrialLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceTrialLicenseRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceTrialLicenseRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstanceTrialLicenseRequest) SetInstanceId(v string) *CreateInstanceTrialLicenseRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceTrialLicenseRequest) Validate() error {
	return dara.Validate(s)
}
