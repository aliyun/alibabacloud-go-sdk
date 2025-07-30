// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrganizationalUnitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrganizationalUnit(v *GetOrganizationalUnitResponseBodyOrganizationalUnit) *GetOrganizationalUnitResponseBody
	GetOrganizationalUnit() *GetOrganizationalUnitResponseBodyOrganizationalUnit
	SetRequestId(v string) *GetOrganizationalUnitResponseBody
	GetRequestId() *string
}

type GetOrganizationalUnitResponseBody struct {
	// The data object of the organizational unit.
	OrganizationalUnit *GetOrganizationalUnitResponseBodyOrganizationalUnit `json:"OrganizationalUnit,omitempty" xml:"OrganizationalUnit,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOrganizationalUnitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOrganizationalUnitResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitResponseBody) GetOrganizationalUnit() *GetOrganizationalUnitResponseBodyOrganizationalUnit {
	return s.OrganizationalUnit
}

func (s *GetOrganizationalUnitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOrganizationalUnitResponseBody) SetOrganizationalUnit(v *GetOrganizationalUnitResponseBodyOrganizationalUnit) *GetOrganizationalUnitResponseBody {
	s.OrganizationalUnit = v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetRequestId(v string) *GetOrganizationalUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetOrganizationalUnitResponseBodyOrganizationalUnit struct {
	// The time when the organizational unit was created. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the organizational unit.
	//
	// example:
	//
	// Test organizational unit
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the node is a leaf node.
	//
	// example:
	//
	// false
	Leaf *bool `json:"Leaf,omitempty" xml:"Leaf,omitempty"`
	// The external ID of the organizational unit. The external ID can be used by external data to map the data of the organizational unit in IDaaS EIAM. By default, the external ID is the organizational unit ID.
	//
	// For organizational units with the same source type and source ID, each organizational unit has a unique external ID.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitExternalId *string `json:"OrganizationalUnitExternalId,omitempty" xml:"OrganizationalUnitExternalId,omitempty"`
	// The ID of the organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
	// The Name of the organizational unit.
	//
	// example:
	//
	// test_organizationalUnit_name
	OrganizationalUnitName *string `json:"OrganizationalUnitName,omitempty" xml:"OrganizationalUnitName,omitempty"`
	// The source ID of the organizational unit.
	//
	// If the organizational unit was created in IDaaS, its source ID is the ID of the IDaaS instance. If the organizational unit was imported, its source ID is the enterprise ID in the source. For example, if the organizational unit was imported from DingTalk, its source ID is the corpId value of the enterprise in DingTalk.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	OrganizationalUnitSourceId *string `json:"OrganizationalUnitSourceId,omitempty" xml:"OrganizationalUnitSourceId,omitempty"`
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
	// example:
	//
	// build_in
	OrganizationalUnitSourceType *string `json:"OrganizationalUnitSourceType,omitempty" xml:"OrganizationalUnitSourceType,omitempty"`
	// The ID of the parent organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The time when the organizational unit was last updated. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetOrganizationalUnitResponseBodyOrganizationalUnit) String() string {
	return dara.Prettify(s)
}

func (s GetOrganizationalUnitResponseBodyOrganizationalUnit) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) GetDescription() *string {
	return s.Description
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) GetLeaf() *bool {
	return s.Leaf
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) GetOrganizationalUnitExternalId() *string {
	return s.OrganizationalUnitExternalId
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) GetOrganizationalUnitName() *string {
	return s.OrganizationalUnitName
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) GetOrganizationalUnitSourceId() *string {
	return s.OrganizationalUnitSourceId
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) GetOrganizationalUnitSourceType() *string {
	return s.OrganizationalUnitSourceType
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) GetParentId() *string {
	return s.ParentId
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) SetCreateTime(v int64) *GetOrganizationalUnitResponseBodyOrganizationalUnit {
	s.CreateTime = &v
	return s
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) SetDescription(v string) *GetOrganizationalUnitResponseBodyOrganizationalUnit {
	s.Description = &v
	return s
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) SetInstanceId(v string) *GetOrganizationalUnitResponseBodyOrganizationalUnit {
	s.InstanceId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) SetLeaf(v bool) *GetOrganizationalUnitResponseBodyOrganizationalUnit {
	s.Leaf = &v
	return s
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) SetOrganizationalUnitExternalId(v string) *GetOrganizationalUnitResponseBodyOrganizationalUnit {
	s.OrganizationalUnitExternalId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) SetOrganizationalUnitId(v string) *GetOrganizationalUnitResponseBodyOrganizationalUnit {
	s.OrganizationalUnitId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) SetOrganizationalUnitName(v string) *GetOrganizationalUnitResponseBodyOrganizationalUnit {
	s.OrganizationalUnitName = &v
	return s
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) SetOrganizationalUnitSourceId(v string) *GetOrganizationalUnitResponseBodyOrganizationalUnit {
	s.OrganizationalUnitSourceId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) SetOrganizationalUnitSourceType(v string) *GetOrganizationalUnitResponseBodyOrganizationalUnit {
	s.OrganizationalUnitSourceType = &v
	return s
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) SetParentId(v string) *GetOrganizationalUnitResponseBodyOrganizationalUnit {
	s.ParentId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) SetUpdateTime(v int64) *GetOrganizationalUnitResponseBodyOrganizationalUnit {
	s.UpdateTime = &v
	return s
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) Validate() error {
	return dara.Validate(s)
}
