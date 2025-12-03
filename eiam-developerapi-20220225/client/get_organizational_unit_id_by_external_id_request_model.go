// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrganizationalUnitIdByExternalIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrganizationalUnitExternalId(v string) *GetOrganizationalUnitIdByExternalIdRequest
	GetOrganizationalUnitExternalId() *string
	SetOrganizationalUnitSourceId(v string) *GetOrganizationalUnitIdByExternalIdRequest
	GetOrganizationalUnitSourceId() *string
	SetOrganizationalUnitSourceType(v string) *GetOrganizationalUnitIdByExternalIdRequest
	GetOrganizationalUnitSourceType() *string
}

type GetOrganizationalUnitIdByExternalIdRequest struct {
	// The external ID of the organizational unit. The external ID can be used to map external data to the data of the organizational unit in Employee Identity and Access Management (EIAM) of Identity as a Service (IDaaS). By default, the external ID is the organizational unit ID.
	//
	// Note: For organizational units with the same source type and source ID, each organizational unit has a unique external ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitExternalId *string `json:"organizationalUnitExternalId,omitempty" xml:"organizationalUnitExternalId,omitempty"`
	// The source ID of the organizational unit.
	//
	// If the organizational unit was created in IDaaS, its source ID is the ID of the IDaaS instance. If the organizational unit was imported, its source ID is the enterprise ID in the source. For example, if the organizational unit was imported from DingTalk, its source ID is the corpId value of the enterprise in DingTalk.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	OrganizationalUnitSourceId *string `json:"organizationalUnitSourceId,omitempty" xml:"organizationalUnitSourceId,omitempty"`
	// The source type of the organizational unit. Valid values:
	//
	// 	- build_in: The organizational unit was created in IDaaS.
	//
	// 	- ding_talk: The organizational unit was imported from DingTalk.
	//
	// 	- ad: The organizational unit was imported from Microsoft Active Directory (AD).
	//
	// 	- ldap: The organizational unit was imported from a Lightweight Directory Access Protocol (LDAP) service.
	//
	// This parameter is required.
	//
	// example:
	//
	// build_in
	OrganizationalUnitSourceType *string `json:"organizationalUnitSourceType,omitempty" xml:"organizationalUnitSourceType,omitempty"`
}

func (s GetOrganizationalUnitIdByExternalIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOrganizationalUnitIdByExternalIdRequest) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitIdByExternalIdRequest) GetOrganizationalUnitExternalId() *string {
	return s.OrganizationalUnitExternalId
}

func (s *GetOrganizationalUnitIdByExternalIdRequest) GetOrganizationalUnitSourceId() *string {
	return s.OrganizationalUnitSourceId
}

func (s *GetOrganizationalUnitIdByExternalIdRequest) GetOrganizationalUnitSourceType() *string {
	return s.OrganizationalUnitSourceType
}

func (s *GetOrganizationalUnitIdByExternalIdRequest) SetOrganizationalUnitExternalId(v string) *GetOrganizationalUnitIdByExternalIdRequest {
	s.OrganizationalUnitExternalId = &v
	return s
}

func (s *GetOrganizationalUnitIdByExternalIdRequest) SetOrganizationalUnitSourceId(v string) *GetOrganizationalUnitIdByExternalIdRequest {
	s.OrganizationalUnitSourceId = &v
	return s
}

func (s *GetOrganizationalUnitIdByExternalIdRequest) SetOrganizationalUnitSourceType(v string) *GetOrganizationalUnitIdByExternalIdRequest {
	s.OrganizationalUnitSourceType = &v
	return s
}

func (s *GetOrganizationalUnitIdByExternalIdRequest) Validate() error {
	return dara.Validate(s)
}
