// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRootOrganizationalUnitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrganizationalUnit(v *GetRootOrganizationalUnitResponseBodyOrganizationalUnit) *GetRootOrganizationalUnitResponseBody
	GetOrganizationalUnit() *GetRootOrganizationalUnitResponseBodyOrganizationalUnit
	SetRequestId(v string) *GetRootOrganizationalUnitResponseBody
	GetRequestId() *string
}

type GetRootOrganizationalUnitResponseBody struct {
	// The data object of the organizational unit.
	OrganizationalUnit *GetRootOrganizationalUnitResponseBodyOrganizationalUnit `json:"OrganizationalUnit,omitempty" xml:"OrganizationalUnit,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRootOrganizationalUnitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRootOrganizationalUnitResponseBody) GoString() string {
	return s.String()
}

func (s *GetRootOrganizationalUnitResponseBody) GetOrganizationalUnit() *GetRootOrganizationalUnitResponseBodyOrganizationalUnit {
	return s.OrganizationalUnit
}

func (s *GetRootOrganizationalUnitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRootOrganizationalUnitResponseBody) SetOrganizationalUnit(v *GetRootOrganizationalUnitResponseBodyOrganizationalUnit) *GetRootOrganizationalUnitResponseBody {
	s.OrganizationalUnit = v
	return s
}

func (s *GetRootOrganizationalUnitResponseBody) SetRequestId(v string) *GetRootOrganizationalUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRootOrganizationalUnitResponseBody) Validate() error {
	if s.OrganizationalUnit != nil {
		if err := s.OrganizationalUnit.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRootOrganizationalUnitResponseBodyOrganizationalUnit struct {
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
	// The ID of the organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
	// The name of the organization.
	//
	// example:
	//
	// name001
	OrganizationalUnitName *string `json:"OrganizationalUnitName,omitempty" xml:"OrganizationalUnitName,omitempty"`
	// The time when the organizational unit was last updated. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetRootOrganizationalUnitResponseBodyOrganizationalUnit) String() string {
	return dara.Prettify(s)
}

func (s GetRootOrganizationalUnitResponseBodyOrganizationalUnit) GoString() string {
	return s.String()
}

func (s *GetRootOrganizationalUnitResponseBodyOrganizationalUnit) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetRootOrganizationalUnitResponseBodyOrganizationalUnit) GetDescription() *string {
	return s.Description
}

func (s *GetRootOrganizationalUnitResponseBodyOrganizationalUnit) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRootOrganizationalUnitResponseBodyOrganizationalUnit) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *GetRootOrganizationalUnitResponseBodyOrganizationalUnit) GetOrganizationalUnitName() *string {
	return s.OrganizationalUnitName
}

func (s *GetRootOrganizationalUnitResponseBodyOrganizationalUnit) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetRootOrganizationalUnitResponseBodyOrganizationalUnit) SetCreateTime(v int64) *GetRootOrganizationalUnitResponseBodyOrganizationalUnit {
	s.CreateTime = &v
	return s
}

func (s *GetRootOrganizationalUnitResponseBodyOrganizationalUnit) SetDescription(v string) *GetRootOrganizationalUnitResponseBodyOrganizationalUnit {
	s.Description = &v
	return s
}

func (s *GetRootOrganizationalUnitResponseBodyOrganizationalUnit) SetInstanceId(v string) *GetRootOrganizationalUnitResponseBodyOrganizationalUnit {
	s.InstanceId = &v
	return s
}

func (s *GetRootOrganizationalUnitResponseBodyOrganizationalUnit) SetOrganizationalUnitId(v string) *GetRootOrganizationalUnitResponseBodyOrganizationalUnit {
	s.OrganizationalUnitId = &v
	return s
}

func (s *GetRootOrganizationalUnitResponseBodyOrganizationalUnit) SetOrganizationalUnitName(v string) *GetRootOrganizationalUnitResponseBodyOrganizationalUnit {
	s.OrganizationalUnitName = &v
	return s
}

func (s *GetRootOrganizationalUnitResponseBodyOrganizationalUnit) SetUpdateTime(v int64) *GetRootOrganizationalUnitResponseBodyOrganizationalUnit {
	s.UpdateTime = &v
	return s
}

func (s *GetRootOrganizationalUnitResponseBodyOrganizationalUnit) Validate() error {
	return dara.Validate(s)
}
