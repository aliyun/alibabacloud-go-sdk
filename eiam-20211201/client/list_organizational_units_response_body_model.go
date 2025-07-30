// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationalUnitsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrganizationalUnits(v []*ListOrganizationalUnitsResponseBodyOrganizationalUnits) *ListOrganizationalUnitsResponseBody
	GetOrganizationalUnits() []*ListOrganizationalUnitsResponseBodyOrganizationalUnits
	SetRequestId(v string) *ListOrganizationalUnitsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListOrganizationalUnitsResponseBody
	GetTotalCount() *int64
}

type ListOrganizationalUnitsResponseBody struct {
	// The list of data objects of organizational units.
	OrganizationalUnits []*ListOrganizationalUnitsResponseBodyOrganizationalUnits `json:"OrganizationalUnits,omitempty" xml:"OrganizationalUnits,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries in the list.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOrganizationalUnitsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsResponseBody) GetOrganizationalUnits() []*ListOrganizationalUnitsResponseBodyOrganizationalUnits {
	return s.OrganizationalUnits
}

func (s *ListOrganizationalUnitsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOrganizationalUnitsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListOrganizationalUnitsResponseBody) SetOrganizationalUnits(v []*ListOrganizationalUnitsResponseBodyOrganizationalUnits) *ListOrganizationalUnitsResponseBody {
	s.OrganizationalUnits = v
	return s
}

func (s *ListOrganizationalUnitsResponseBody) SetRequestId(v string) *ListOrganizationalUnitsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBody) SetTotalCount(v int64) *ListOrganizationalUnitsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListOrganizationalUnitsResponseBodyOrganizationalUnits struct {
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
	// 组织名称。
	//
	// example:
	//
	// test_organizationalUnit_name
	OrganizationalUnitName *string `json:"OrganizationalUnitName,omitempty" xml:"OrganizationalUnitName,omitempty"`
	// The source ID of the organizational unit.
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

func (s ListOrganizationalUnitsResponseBodyOrganizationalUnits) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitsResponseBodyOrganizationalUnits) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) GetDescription() *string {
	return s.Description
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) GetLeaf() *bool {
	return s.Leaf
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) GetOrganizationalUnitExternalId() *string {
	return s.OrganizationalUnitExternalId
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) GetOrganizationalUnitName() *string {
	return s.OrganizationalUnitName
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) GetOrganizationalUnitSourceId() *string {
	return s.OrganizationalUnitSourceId
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) GetOrganizationalUnitSourceType() *string {
	return s.OrganizationalUnitSourceType
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) GetParentId() *string {
	return s.ParentId
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) SetCreateTime(v int64) *ListOrganizationalUnitsResponseBodyOrganizationalUnits {
	s.CreateTime = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) SetDescription(v string) *ListOrganizationalUnitsResponseBodyOrganizationalUnits {
	s.Description = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) SetInstanceId(v string) *ListOrganizationalUnitsResponseBodyOrganizationalUnits {
	s.InstanceId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) SetLeaf(v bool) *ListOrganizationalUnitsResponseBodyOrganizationalUnits {
	s.Leaf = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) SetOrganizationalUnitExternalId(v string) *ListOrganizationalUnitsResponseBodyOrganizationalUnits {
	s.OrganizationalUnitExternalId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) SetOrganizationalUnitId(v string) *ListOrganizationalUnitsResponseBodyOrganizationalUnits {
	s.OrganizationalUnitId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) SetOrganizationalUnitName(v string) *ListOrganizationalUnitsResponseBodyOrganizationalUnits {
	s.OrganizationalUnitName = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) SetOrganizationalUnitSourceId(v string) *ListOrganizationalUnitsResponseBodyOrganizationalUnits {
	s.OrganizationalUnitSourceId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) SetOrganizationalUnitSourceType(v string) *ListOrganizationalUnitsResponseBodyOrganizationalUnits {
	s.OrganizationalUnitSourceType = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) SetParentId(v string) *ListOrganizationalUnitsResponseBodyOrganizationalUnits {
	s.ParentId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) SetUpdateTime(v int64) *ListOrganizationalUnitsResponseBodyOrganizationalUnits {
	s.UpdateTime = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) Validate() error {
	return dara.Validate(s)
}
