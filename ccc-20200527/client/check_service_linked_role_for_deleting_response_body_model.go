// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckServiceLinkedRoleForDeletingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeleTable(v bool) *CheckServiceLinkedRoleForDeletingResponseBody
	GetDeleTable() *bool
	SetRequestId(v string) *CheckServiceLinkedRoleForDeletingResponseBody
	GetRequestId() *string
	SetRoleUsages(v []*CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages) *CheckServiceLinkedRoleForDeletingResponseBody
	GetRoleUsages() []*CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages
}

type CheckServiceLinkedRoleForDeletingResponseBody struct {
	DeleTable  *bool                                                      `json:"DeleTable,omitempty" xml:"DeleTable,omitempty"`
	RequestId  *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RoleUsages []*CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages `json:"RoleUsages,omitempty" xml:"RoleUsages,omitempty" type:"Repeated"`
}

func (s CheckServiceLinkedRoleForDeletingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceLinkedRoleForDeletingResponseBody) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleForDeletingResponseBody) GetDeleTable() *bool {
	return s.DeleTable
}

func (s *CheckServiceLinkedRoleForDeletingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckServiceLinkedRoleForDeletingResponseBody) GetRoleUsages() []*CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages {
	return s.RoleUsages
}

func (s *CheckServiceLinkedRoleForDeletingResponseBody) SetDeleTable(v bool) *CheckServiceLinkedRoleForDeletingResponseBody {
	s.DeleTable = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingResponseBody) SetRequestId(v string) *CheckServiceLinkedRoleForDeletingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingResponseBody) SetRoleUsages(v []*CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages) *CheckServiceLinkedRoleForDeletingResponseBody {
	s.RoleUsages = v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingResponseBody) Validate() error {
	if s.RoleUsages != nil {
		for _, item := range s.RoleUsages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages struct {
	Region    *string   `json:"Region,omitempty" xml:"Region,omitempty"`
	Resources []*string `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages) GetRegion() *string {
	return s.Region
}

func (s *CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages) GetResources() []*string {
	return s.Resources
}

func (s *CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages) SetRegion(v string) *CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages {
	s.Region = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages) SetResources(v []*string) *CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages {
	s.Resources = v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingResponseBodyRoleUsages) Validate() error {
	return dara.Validate(s)
}
