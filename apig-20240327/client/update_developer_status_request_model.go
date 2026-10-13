// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeveloperStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPortalId(v string) *UpdateDeveloperStatusRequest
	GetPortalId() *string
	SetStatus(v string) *UpdateDeveloperStatusRequest
	GetStatus() *string
}

type UpdateDeveloperStatusRequest struct {
	// The ID of the portal to which the developer belongs, which is the primary ID of the portal resource. You can obtain this ID from the portalId field in the CreatePortal response or from the portalId field of an item in the ListPortals response.
	//
	// example:
	//
	// portal-xxxxxxxx
	PortalId *string `json:"portalId,omitempty" xml:"portalId,omitempty"`
	// The target status of the developer. Valid values: PENDING, APPROVED, and REJECTED.
	//
	// example:
	//
	// APPROVED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateDeveloperStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeveloperStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeveloperStatusRequest) GetPortalId() *string {
	return s.PortalId
}

func (s *UpdateDeveloperStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateDeveloperStatusRequest) SetPortalId(v string) *UpdateDeveloperStatusRequest {
	s.PortalId = &v
	return s
}

func (s *UpdateDeveloperStatusRequest) SetStatus(v string) *UpdateDeveloperStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateDeveloperStatusRequest) Validate() error {
	return dara.Validate(s)
}
