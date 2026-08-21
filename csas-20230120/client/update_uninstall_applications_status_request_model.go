// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUninstallApplicationsStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationIds(v []*string) *UpdateUninstallApplicationsStatusRequest
	GetApplicationIds() []*string
	SetStatus(v string) *UpdateUninstallApplicationsStatusRequest
	GetStatus() *string
}

type UpdateUninstallApplicationsStatusRequest struct {
	// The list of uninstall application IDs.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// The status of the uninstall application. Valid values:
	//
	// - **Approved**: Approved.
	//
	// - **Rejected**: Rejected.
	//
	// example:
	//
	// Approved
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateUninstallApplicationsStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUninstallApplicationsStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateUninstallApplicationsStatusRequest) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *UpdateUninstallApplicationsStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateUninstallApplicationsStatusRequest) SetApplicationIds(v []*string) *UpdateUninstallApplicationsStatusRequest {
	s.ApplicationIds = v
	return s
}

func (s *UpdateUninstallApplicationsStatusRequest) SetStatus(v string) *UpdateUninstallApplicationsStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateUninstallApplicationsStatusRequest) Validate() error {
	return dara.Validate(s)
}
