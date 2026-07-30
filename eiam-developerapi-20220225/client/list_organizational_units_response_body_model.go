// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationalUnitsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListOrganizationalUnitsResponseBodyData) *ListOrganizationalUnitsResponseBody
	GetData() []*ListOrganizationalUnitsResponseBodyData
	SetTotalCount(v int64) *ListOrganizationalUnitsResponseBody
	GetTotalCount() *int64
}

type ListOrganizationalUnitsResponseBody struct {
	// A list of data objects.
	Data []*ListOrganizationalUnitsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 1000
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListOrganizationalUnitsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsResponseBody) GetData() []*ListOrganizationalUnitsResponseBodyData {
	return s.Data
}

func (s *ListOrganizationalUnitsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListOrganizationalUnitsResponseBody) SetData(v []*ListOrganizationalUnitsResponseBodyData) *ListOrganizationalUnitsResponseBody {
	s.Data = v
	return s
}

func (s *ListOrganizationalUnitsResponseBody) SetTotalCount(v int64) *ListOrganizationalUnitsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOrganizationalUnitsResponseBodyData struct {
	// The time when the organization was created. This is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652083425923
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description of the organization.
	//
	// example:
	//
	// 测试组织
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The external ID of the organization. This ID is used to map external data to the organization\\"s data in IDaaS. The default value is the IDaaS organization ID.
	//
	// Note: The external ID must be unique for the same source type and source ID.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitExternalId *string `json:"organizationalUnitExternalId,omitempty" xml:"organizationalUnitExternalId,omitempty"`
	// The organization ID.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"organizationalUnitId,omitempty" xml:"organizationalUnitId,omitempty"`
	// The organization name.
	//
	// example:
	//
	// name001
	OrganizationalUnitName *string `json:"organizationalUnitName,omitempty" xml:"organizationalUnitName,omitempty"`
	// The source ID of the organization.
	//
	// For the \\`build_in\\` type, the default value is the instance ID. For other types, the value is the enterprise ID from the source. For example, if the source is DingTalk, the value is the \\`corpId\\` of the DingTalk enterprise.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	OrganizationalUnitSourceId *string `json:"organizationalUnitSourceId,omitempty" xml:"organizationalUnitSourceId,omitempty"`
	// The source type of the organization. Valid values:
	//
	// - \\`build_in\\`: The organization is created in IDaaS.
	//
	// - \\`ding_talk\\`: The organization is imported from DingTalk.
	//
	// - \\`ad\\`: The organization is imported from Active Directory (AD).
	//
	// - \\`ldap\\`: The organization is imported from LDAP.
	//
	// example:
	//
	// build_in
	OrganizationalUnitSourceType *string `json:"organizationalUnitSourceType,omitempty" xml:"organizationalUnitSourceType,omitempty"`
	// The parent organization ID.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The time when the organization was last updated. This is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652083425923
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListOrganizationalUnitsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListOrganizationalUnitsResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListOrganizationalUnitsResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOrganizationalUnitsResponseBodyData) GetOrganizationalUnitExternalId() *string {
	return s.OrganizationalUnitExternalId
}

func (s *ListOrganizationalUnitsResponseBodyData) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *ListOrganizationalUnitsResponseBodyData) GetOrganizationalUnitName() *string {
	return s.OrganizationalUnitName
}

func (s *ListOrganizationalUnitsResponseBodyData) GetOrganizationalUnitSourceId() *string {
	return s.OrganizationalUnitSourceId
}

func (s *ListOrganizationalUnitsResponseBodyData) GetOrganizationalUnitSourceType() *string {
	return s.OrganizationalUnitSourceType
}

func (s *ListOrganizationalUnitsResponseBodyData) GetParentId() *string {
	return s.ParentId
}

func (s *ListOrganizationalUnitsResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListOrganizationalUnitsResponseBodyData) SetCreateTime(v int64) *ListOrganizationalUnitsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetDescription(v string) *ListOrganizationalUnitsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetInstanceId(v string) *ListOrganizationalUnitsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetOrganizationalUnitExternalId(v string) *ListOrganizationalUnitsResponseBodyData {
	s.OrganizationalUnitExternalId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetOrganizationalUnitId(v string) *ListOrganizationalUnitsResponseBodyData {
	s.OrganizationalUnitId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetOrganizationalUnitName(v string) *ListOrganizationalUnitsResponseBodyData {
	s.OrganizationalUnitName = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetOrganizationalUnitSourceId(v string) *ListOrganizationalUnitsResponseBodyData {
	s.OrganizationalUnitSourceId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetOrganizationalUnitSourceType(v string) *ListOrganizationalUnitsResponseBodyData {
	s.OrganizationalUnitSourceType = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetParentId(v string) *ListOrganizationalUnitsResponseBodyData {
	s.ParentId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) SetUpdateTime(v int64) *ListOrganizationalUnitsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
