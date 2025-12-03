// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrganizationalUnitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *GetOrganizationalUnitResponseBody
	GetCreateTime() *int64
	SetDescription(v string) *GetOrganizationalUnitResponseBody
	GetDescription() *string
	SetInstanceId(v string) *GetOrganizationalUnitResponseBody
	GetInstanceId() *string
	SetOrganizationalUnitExternalId(v string) *GetOrganizationalUnitResponseBody
	GetOrganizationalUnitExternalId() *string
	SetOrganizationalUnitId(v string) *GetOrganizationalUnitResponseBody
	GetOrganizationalUnitId() *string
	SetOrganizationalUnitName(v string) *GetOrganizationalUnitResponseBody
	GetOrganizationalUnitName() *string
	SetOrganizationalUnitSourceId(v string) *GetOrganizationalUnitResponseBody
	GetOrganizationalUnitSourceId() *string
	SetOrganizationalUnitSourceType(v string) *GetOrganizationalUnitResponseBody
	GetOrganizationalUnitSourceType() *string
	SetParentId(v string) *GetOrganizationalUnitResponseBody
	GetParentId() *string
	SetUpdateTime(v int64) *GetOrganizationalUnitResponseBody
	GetUpdateTime() *int64
}

type GetOrganizationalUnitResponseBody struct {
	// The time when the organizational unit was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652083425923
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description of the organizational unit.
	//
	// example:
	//
	// xxxxx
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The external ID of the organizational unit.
	//
	// example:
	//
	// id_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitExternalId *string `json:"organizationalUnitExternalId,omitempty" xml:"organizationalUnitExternalId,omitempty"`
	// The ID of the organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"organizationalUnitId,omitempty" xml:"organizationalUnitId,omitempty"`
	// The name of the organizational unit.
	//
	// example:
	//
	// name001
	OrganizationalUnitName *string `json:"organizationalUnitName,omitempty" xml:"organizationalUnitName,omitempty"`
	// The source ID of the organizational unit.
	//
	// example:
	//
	// id_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitSourceId *string `json:"organizationalUnitSourceId,omitempty" xml:"organizationalUnitSourceId,omitempty"`
	// The source type of the organizational unit. Valid values:
	//
	// 	- build_in: The organizational unit was created in Identity as a Service (IDaaS).
	//
	// 	- ding_talk: The organizational unit was imported from DingTalk.
	//
	// 	- ad: The organizational unit was imported from Microsoft Active Directory (AD).
	//
	// 	- ldap: The organizational unit was imported from a Lightweight Directory Access Protocol (LDAP) service.
	//
	// example:
	//
	// build_in
	OrganizationalUnitSourceType *string `json:"organizationalUnitSourceType,omitempty" xml:"organizationalUnitSourceType,omitempty"`
	// The ID of the parent organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The time when the organizational unit was last updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652083425923
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetOrganizationalUnitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOrganizationalUnitResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetOrganizationalUnitResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetOrganizationalUnitResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetOrganizationalUnitResponseBody) GetOrganizationalUnitExternalId() *string {
	return s.OrganizationalUnitExternalId
}

func (s *GetOrganizationalUnitResponseBody) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *GetOrganizationalUnitResponseBody) GetOrganizationalUnitName() *string {
	return s.OrganizationalUnitName
}

func (s *GetOrganizationalUnitResponseBody) GetOrganizationalUnitSourceId() *string {
	return s.OrganizationalUnitSourceId
}

func (s *GetOrganizationalUnitResponseBody) GetOrganizationalUnitSourceType() *string {
	return s.OrganizationalUnitSourceType
}

func (s *GetOrganizationalUnitResponseBody) GetParentId() *string {
	return s.ParentId
}

func (s *GetOrganizationalUnitResponseBody) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetOrganizationalUnitResponseBody) SetCreateTime(v int64) *GetOrganizationalUnitResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetDescription(v string) *GetOrganizationalUnitResponseBody {
	s.Description = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetInstanceId(v string) *GetOrganizationalUnitResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetOrganizationalUnitExternalId(v string) *GetOrganizationalUnitResponseBody {
	s.OrganizationalUnitExternalId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetOrganizationalUnitId(v string) *GetOrganizationalUnitResponseBody {
	s.OrganizationalUnitId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetOrganizationalUnitName(v string) *GetOrganizationalUnitResponseBody {
	s.OrganizationalUnitName = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetOrganizationalUnitSourceId(v string) *GetOrganizationalUnitResponseBody {
	s.OrganizationalUnitSourceId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetOrganizationalUnitSourceType(v string) *GetOrganizationalUnitResponseBody {
	s.OrganizationalUnitSourceType = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetParentId(v string) *GetOrganizationalUnitResponseBody {
	s.ParentId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetUpdateTime(v int64) *GetOrganizationalUnitResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) Validate() error {
	return dara.Validate(s)
}
