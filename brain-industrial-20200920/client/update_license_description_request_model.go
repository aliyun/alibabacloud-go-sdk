// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLicenseDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateLicenseDescriptionRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateLicenseDescriptionRequest
	GetInstanceId() *string
}

type UpdateLicenseDescriptionRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// brainindustrial_aicsruntime_public_cn-mdu3ps3kw04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateLicenseDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLicenseDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateLicenseDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateLicenseDescriptionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateLicenseDescriptionRequest) SetDescription(v string) *UpdateLicenseDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateLicenseDescriptionRequest) SetInstanceId(v string) *UpdateLicenseDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateLicenseDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
