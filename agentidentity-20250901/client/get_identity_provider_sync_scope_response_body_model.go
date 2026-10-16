// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderSyncScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEstimateStatus(v string) *GetIdentityProviderSyncScopeResponseBody
	GetEstimateStatus() *string
	SetEstimatedUserCount(v int64) *GetIdentityProviderSyncScopeResponseBody
	GetEstimatedUserCount() *int64
	SetExpandedDepartmentCount(v int64) *GetIdentityProviderSyncScopeResponseBody
	GetExpandedDepartmentCount() *int64
	SetRequestId(v string) *GetIdentityProviderSyncScopeResponseBody
	GetRequestId() *string
	SetSelectedDepartments(v []*GetIdentityProviderSyncScopeResponseBodySelectedDepartments) *GetIdentityProviderSyncScopeResponseBody
	GetSelectedDepartments() []*GetIdentityProviderSyncScopeResponseBodySelectedDepartments
	SetSyncScope(v *GetIdentityProviderSyncScopeResponseBodySyncScope) *GetIdentityProviderSyncScopeResponseBody
	GetSyncScope() *GetIdentityProviderSyncScopeResponseBodySyncScope
}

type GetIdentityProviderSyncScopeResponseBody struct {
	// example:
	//
	// Ready
	EstimateStatus *string `json:"EstimateStatus,omitempty" xml:"EstimateStatus,omitempty"`
	// example:
	//
	// 24
	EstimatedUserCount *int64 `json:"EstimatedUserCount,omitempty" xml:"EstimatedUserCount,omitempty"`
	// example:
	//
	// 3
	ExpandedDepartmentCount *int64 `json:"ExpandedDepartmentCount,omitempty" xml:"ExpandedDepartmentCount,omitempty"`
	// example:
	//
	// EXAMPLE-REQUEST-ID
	RequestId           *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SelectedDepartments []*GetIdentityProviderSyncScopeResponseBodySelectedDepartments `json:"SelectedDepartments,omitempty" xml:"SelectedDepartments,omitempty" type:"Repeated"`
	SyncScope           *GetIdentityProviderSyncScopeResponseBodySyncScope             `json:"SyncScope,omitempty" xml:"SyncScope,omitempty" type:"Struct"`
}

func (s GetIdentityProviderSyncScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderSyncScopeResponseBody) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderSyncScopeResponseBody) GetEstimateStatus() *string {
	return s.EstimateStatus
}

func (s *GetIdentityProviderSyncScopeResponseBody) GetEstimatedUserCount() *int64 {
	return s.EstimatedUserCount
}

func (s *GetIdentityProviderSyncScopeResponseBody) GetExpandedDepartmentCount() *int64 {
	return s.ExpandedDepartmentCount
}

func (s *GetIdentityProviderSyncScopeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIdentityProviderSyncScopeResponseBody) GetSelectedDepartments() []*GetIdentityProviderSyncScopeResponseBodySelectedDepartments {
	return s.SelectedDepartments
}

func (s *GetIdentityProviderSyncScopeResponseBody) GetSyncScope() *GetIdentityProviderSyncScopeResponseBodySyncScope {
	return s.SyncScope
}

func (s *GetIdentityProviderSyncScopeResponseBody) SetEstimateStatus(v string) *GetIdentityProviderSyncScopeResponseBody {
	s.EstimateStatus = &v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBody) SetEstimatedUserCount(v int64) *GetIdentityProviderSyncScopeResponseBody {
	s.EstimatedUserCount = &v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBody) SetExpandedDepartmentCount(v int64) *GetIdentityProviderSyncScopeResponseBody {
	s.ExpandedDepartmentCount = &v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBody) SetRequestId(v string) *GetIdentityProviderSyncScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBody) SetSelectedDepartments(v []*GetIdentityProviderSyncScopeResponseBodySelectedDepartments) *GetIdentityProviderSyncScopeResponseBody {
	s.SelectedDepartments = v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBody) SetSyncScope(v *GetIdentityProviderSyncScopeResponseBodySyncScope) *GetIdentityProviderSyncScopeResponseBody {
	s.SyncScope = v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBody) Validate() error {
	if s.SelectedDepartments != nil {
		for _, item := range s.SelectedDepartments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SyncScope != nil {
		if err := s.SyncScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIdentityProviderSyncScopeResponseBodySelectedDepartments struct {
	// example:
	//
	// true
	Exists *bool `json:"Exists,omitempty" xml:"Exists,omitempty"`
	// example:
	//
	// 100
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 研发部
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetIdentityProviderSyncScopeResponseBodySelectedDepartments) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderSyncScopeResponseBodySelectedDepartments) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderSyncScopeResponseBodySelectedDepartments) GetExists() *bool {
	return s.Exists
}

func (s *GetIdentityProviderSyncScopeResponseBodySelectedDepartments) GetId() *string {
	return s.Id
}

func (s *GetIdentityProviderSyncScopeResponseBodySelectedDepartments) GetName() *string {
	return s.Name
}

func (s *GetIdentityProviderSyncScopeResponseBodySelectedDepartments) SetExists(v bool) *GetIdentityProviderSyncScopeResponseBodySelectedDepartments {
	s.Exists = &v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBodySelectedDepartments) SetId(v string) *GetIdentityProviderSyncScopeResponseBodySelectedDepartments {
	s.Id = &v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBodySelectedDepartments) SetName(v string) *GetIdentityProviderSyncScopeResponseBodySelectedDepartments {
	s.Name = &v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBodySelectedDepartments) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderSyncScopeResponseBodySyncScope struct {
	DepartmentIds []*string `json:"DepartmentIds,omitempty" xml:"DepartmentIds,omitempty" type:"Repeated"`
	// example:
	//
	// SelectedDepartments
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s GetIdentityProviderSyncScopeResponseBodySyncScope) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderSyncScopeResponseBodySyncScope) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderSyncScopeResponseBodySyncScope) GetDepartmentIds() []*string {
	return s.DepartmentIds
}

func (s *GetIdentityProviderSyncScopeResponseBodySyncScope) GetMode() *string {
	return s.Mode
}

func (s *GetIdentityProviderSyncScopeResponseBodySyncScope) SetDepartmentIds(v []*string) *GetIdentityProviderSyncScopeResponseBodySyncScope {
	s.DepartmentIds = v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBodySyncScope) SetMode(v string) *GetIdentityProviderSyncScopeResponseBodySyncScope {
	s.Mode = &v
	return s
}

func (s *GetIdentityProviderSyncScopeResponseBodySyncScope) Validate() error {
	return dara.Validate(s)
}
