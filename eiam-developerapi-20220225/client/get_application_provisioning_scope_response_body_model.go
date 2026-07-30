// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationProvisioningScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupIds(v []*string) *GetApplicationProvisioningScopeResponseBody
	GetGroupIds() []*string
	SetOrganizationalUnitIds(v []*string) *GetApplicationProvisioningScopeResponseBody
	GetOrganizationalUnitIds() []*string
}

type GetApplicationProvisioningScopeResponseBody struct {
	// The list of group IDs.
	GroupIds []*string `json:"groupIds,omitempty" xml:"groupIds,omitempty" type:"Repeated"`
	// The list of organization IDs.
	//
	// example:
	//
	// [ou_xxx001]
	OrganizationalUnitIds []*string `json:"organizationalUnitIds,omitempty" xml:"organizationalUnitIds,omitempty" type:"Repeated"`
}

func (s GetApplicationProvisioningScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationProvisioningScopeResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningScopeResponseBody) GetGroupIds() []*string {
	return s.GroupIds
}

func (s *GetApplicationProvisioningScopeResponseBody) GetOrganizationalUnitIds() []*string {
	return s.OrganizationalUnitIds
}

func (s *GetApplicationProvisioningScopeResponseBody) SetGroupIds(v []*string) *GetApplicationProvisioningScopeResponseBody {
	s.GroupIds = v
	return s
}

func (s *GetApplicationProvisioningScopeResponseBody) SetOrganizationalUnitIds(v []*string) *GetApplicationProvisioningScopeResponseBody {
	s.OrganizationalUnitIds = v
	return s
}

func (s *GetApplicationProvisioningScopeResponseBody) Validate() error {
	return dara.Validate(s)
}
