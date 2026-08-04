// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExcessiveDeviceRegistrationApplicationsStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationIds(v []*string) *UpdateExcessiveDeviceRegistrationApplicationsStatusRequest
	GetApplicationIds() []*string
	SetStatus(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusRequest
	GetStatus() *string
}

type UpdateExcessiveDeviceRegistrationApplicationsStatusRequest struct {
	// List of IDs for device registration applications that exceed your quota.
	//
	// This parameter is required.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// Status of the device registration application. Valid values:
	//
	// - **Approved**: Approve the application. You can approve only applications with a Pending status.
	//
	// - **Rejected**: Reject the application. You can reject only applications with a Pending status.
	//
	// This parameter is required.
	//
	// example:
	//
	// Approved
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateExcessiveDeviceRegistrationApplicationsStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateExcessiveDeviceRegistrationApplicationsStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusRequest) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusRequest) SetApplicationIds(v []*string) *UpdateExcessiveDeviceRegistrationApplicationsStatusRequest {
	s.ApplicationIds = v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusRequest) SetStatus(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusRequest) Validate() error {
	return dara.Validate(s)
}
