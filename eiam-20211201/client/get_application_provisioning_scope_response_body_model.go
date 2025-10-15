// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationProvisioningScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationProvisioningScope(v *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope) *GetApplicationProvisioningScopeResponseBody
	GetApplicationProvisioningScope() *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope
	SetRequestId(v string) *GetApplicationProvisioningScopeResponseBody
	GetRequestId() *string
}

type GetApplicationProvisioningScopeResponseBody struct {
	// The scope of account synchronization.
	ApplicationProvisioningScope *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope `json:"ApplicationProvisioningScope,omitempty" xml:"ApplicationProvisioningScope,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationProvisioningScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationProvisioningScopeResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningScopeResponseBody) GetApplicationProvisioningScope() *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope {
	return s.ApplicationProvisioningScope
}

func (s *GetApplicationProvisioningScopeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationProvisioningScopeResponseBody) SetApplicationProvisioningScope(v *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope) *GetApplicationProvisioningScopeResponseBody {
	s.ApplicationProvisioningScope = v
	return s
}

func (s *GetApplicationProvisioningScopeResponseBody) SetRequestId(v string) *GetApplicationProvisioningScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationProvisioningScopeResponseBody) Validate() error {
	if s.ApplicationProvisioningScope != nil {
		if err := s.ApplicationProvisioningScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope struct {
	// Synchronize the list of authorized groups.
	GroupIds []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// Instance Indicates the maximum quota number of authorized agents.
	//
	// example:
	//
	// 20
	MaxQuota *int32 `json:"MaxQuota,omitempty" xml:"MaxQuota,omitempty"`
	// The list of organizational units that are authorized for account synchronization.
	OrganizationalUnitIds []*string `json:"OrganizationalUnitIds,omitempty" xml:"OrganizationalUnitIds,omitempty" type:"Repeated"`
	// Indicates the quota number of used authorized agents.
	//
	// example:
	//
	// 10
	UsedQuota *int32 `json:"UsedQuota,omitempty" xml:"UsedQuota,omitempty"`
}

func (s GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope) GetGroupIds() []*string {
	return s.GroupIds
}

func (s *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope) GetMaxQuota() *int32 {
	return s.MaxQuota
}

func (s *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope) GetOrganizationalUnitIds() []*string {
	return s.OrganizationalUnitIds
}

func (s *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope) GetUsedQuota() *int32 {
	return s.UsedQuota
}

func (s *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope) SetGroupIds(v []*string) *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope {
	s.GroupIds = v
	return s
}

func (s *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope) SetMaxQuota(v int32) *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope {
	s.MaxQuota = &v
	return s
}

func (s *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope) SetOrganizationalUnitIds(v []*string) *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope {
	s.OrganizationalUnitIds = v
	return s
}

func (s *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope) SetUsedQuota(v int32) *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope {
	s.UsedQuota = &v
	return s
}

func (s *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope) Validate() error {
	return dara.Validate(s)
}
