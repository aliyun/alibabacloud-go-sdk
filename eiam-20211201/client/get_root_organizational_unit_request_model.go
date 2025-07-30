// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRootOrganizationalUnitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetRootOrganizationalUnitRequest
	GetInstanceId() *string
}

type GetRootOrganizationalUnitRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetRootOrganizationalUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRootOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *GetRootOrganizationalUnitRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRootOrganizationalUnitRequest) SetInstanceId(v string) *GetRootOrganizationalUnitRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRootOrganizationalUnitRequest) Validate() error {
	return dara.Validate(s)
}
