// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBaseGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v int64) *BaseGroupResponse
	GetCreatedAt() *int64
	SetCreator(v string) *BaseGroupResponse
	GetCreator() *string
	SetDescription(v string) *BaseGroupResponse
	GetDescription() *string
	SetDomainId(v string) *BaseGroupResponse
	GetDomainId() *string
	SetGroupId(v string) *BaseGroupResponse
	GetGroupId() *string
	SetGroupName(v string) *BaseGroupResponse
	GetGroupName() *string
	SetIsSync(v bool) *BaseGroupResponse
	GetIsSync() *bool
	SetPermission(v map[string]*IDPermission) *BaseGroupResponse
	GetPermission() map[string]*IDPermission
	SetUpdatedAt(v string) *BaseGroupResponse
	GetUpdatedAt() *string
}

type BaseGroupResponse struct {
	// example:
	//
	// 111111
	CreatedAt *int64 `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// example:
	//
	// system
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// desc-111
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// bj123
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// example:
	//
	// b38b5681bd964950ad8bc0f8ea504793
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// example:
	//
	// name-111
	GroupName  *string                  `json:"group_name,omitempty" xml:"group_name,omitempty"`
	IsSync     *bool                    `json:"is_sync,omitempty" xml:"is_sync,omitempty"`
	Permission map[string]*IDPermission `json:"permission,omitempty" xml:"permission,omitempty"`
	// example:
	//
	// 111111
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s BaseGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s BaseGroupResponse) GoString() string {
	return s.String()
}

func (s *BaseGroupResponse) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *BaseGroupResponse) GetCreator() *string {
	return s.Creator
}

func (s *BaseGroupResponse) GetDescription() *string {
	return s.Description
}

func (s *BaseGroupResponse) GetDomainId() *string {
	return s.DomainId
}

func (s *BaseGroupResponse) GetGroupId() *string {
	return s.GroupId
}

func (s *BaseGroupResponse) GetGroupName() *string {
	return s.GroupName
}

func (s *BaseGroupResponse) GetIsSync() *bool {
	return s.IsSync
}

func (s *BaseGroupResponse) GetPermission() map[string]*IDPermission {
	return s.Permission
}

func (s *BaseGroupResponse) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *BaseGroupResponse) SetCreatedAt(v int64) *BaseGroupResponse {
	s.CreatedAt = &v
	return s
}

func (s *BaseGroupResponse) SetCreator(v string) *BaseGroupResponse {
	s.Creator = &v
	return s
}

func (s *BaseGroupResponse) SetDescription(v string) *BaseGroupResponse {
	s.Description = &v
	return s
}

func (s *BaseGroupResponse) SetDomainId(v string) *BaseGroupResponse {
	s.DomainId = &v
	return s
}

func (s *BaseGroupResponse) SetGroupId(v string) *BaseGroupResponse {
	s.GroupId = &v
	return s
}

func (s *BaseGroupResponse) SetGroupName(v string) *BaseGroupResponse {
	s.GroupName = &v
	return s
}

func (s *BaseGroupResponse) SetIsSync(v bool) *BaseGroupResponse {
	s.IsSync = &v
	return s
}

func (s *BaseGroupResponse) SetPermission(v map[string]*IDPermission) *BaseGroupResponse {
	s.Permission = v
	return s
}

func (s *BaseGroupResponse) SetUpdatedAt(v string) *BaseGroupResponse {
	s.UpdatedAt = &v
	return s
}

func (s *BaseGroupResponse) Validate() error {
	return dara.Validate(s)
}
