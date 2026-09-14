// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeResourceAuthUserMappingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComputeResourceId(v int64) *UpdateComputeResourceAuthUserMappingsRequest
	GetComputeResourceId() *int64
	SetProjectId(v int64) *UpdateComputeResourceAuthUserMappingsRequest
	GetProjectId() *int64
	SetRemoveUserIds(v []*string) *UpdateComputeResourceAuthUserMappingsRequest
	GetRemoveUserIds() []*string
	SetUpserts(v []*UpdateComputeResourceAuthUserMappingsRequestUpserts) *UpdateComputeResourceAuthUserMappingsRequest
	GetUpserts() []*UpdateComputeResourceAuthUserMappingsRequestUpserts
}

type UpdateComputeResourceAuthUserMappingsRequest struct {
	// The compute resource ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123455
	ComputeResourceId *int64 `json:"ComputeResourceId,omitempty" xml:"ComputeResourceId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The list of user mappings to remove.
	RemoveUserIds []*string `json:"RemoveUserIds,omitempty" xml:"RemoveUserIds,omitempty" type:"Repeated"`
	// The list of objects to update.
	Upserts []*UpdateComputeResourceAuthUserMappingsRequestUpserts `json:"Upserts,omitempty" xml:"Upserts,omitempty" type:"Repeated"`
}

func (s UpdateComputeResourceAuthUserMappingsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeResourceAuthUserMappingsRequest) GoString() string {
	return s.String()
}

func (s *UpdateComputeResourceAuthUserMappingsRequest) GetComputeResourceId() *int64 {
	return s.ComputeResourceId
}

func (s *UpdateComputeResourceAuthUserMappingsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateComputeResourceAuthUserMappingsRequest) GetRemoveUserIds() []*string {
	return s.RemoveUserIds
}

func (s *UpdateComputeResourceAuthUserMappingsRequest) GetUpserts() []*UpdateComputeResourceAuthUserMappingsRequestUpserts {
	return s.Upserts
}

func (s *UpdateComputeResourceAuthUserMappingsRequest) SetComputeResourceId(v int64) *UpdateComputeResourceAuthUserMappingsRequest {
	s.ComputeResourceId = &v
	return s
}

func (s *UpdateComputeResourceAuthUserMappingsRequest) SetProjectId(v int64) *UpdateComputeResourceAuthUserMappingsRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateComputeResourceAuthUserMappingsRequest) SetRemoveUserIds(v []*string) *UpdateComputeResourceAuthUserMappingsRequest {
	s.RemoveUserIds = v
	return s
}

func (s *UpdateComputeResourceAuthUserMappingsRequest) SetUpserts(v []*UpdateComputeResourceAuthUserMappingsRequestUpserts) *UpdateComputeResourceAuthUserMappingsRequest {
	s.Upserts = v
	return s
}

func (s *UpdateComputeResourceAuthUserMappingsRequest) Validate() error {
	if s.Upserts != nil {
		for _, item := range s.Upserts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateComputeResourceAuthUserMappingsRequestUpserts struct {
	// The password of the target system for the mapping, such as an LDAP password.
	//
	// example:
	//
	// 123xx
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The Alibaba Cloud UID.
	//
	// example:
	//
	// 12747300953xxx62
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username of the target system for the mapping, such as an LDAP username.
	//
	// example:
	//
	// lisa
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s UpdateComputeResourceAuthUserMappingsRequestUpserts) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeResourceAuthUserMappingsRequestUpserts) GoString() string {
	return s.String()
}

func (s *UpdateComputeResourceAuthUserMappingsRequestUpserts) GetPassword() *string {
	return s.Password
}

func (s *UpdateComputeResourceAuthUserMappingsRequestUpserts) GetUserId() *string {
	return s.UserId
}

func (s *UpdateComputeResourceAuthUserMappingsRequestUpserts) GetUsername() *string {
	return s.Username
}

func (s *UpdateComputeResourceAuthUserMappingsRequestUpserts) SetPassword(v string) *UpdateComputeResourceAuthUserMappingsRequestUpserts {
	s.Password = &v
	return s
}

func (s *UpdateComputeResourceAuthUserMappingsRequestUpserts) SetUserId(v string) *UpdateComputeResourceAuthUserMappingsRequestUpserts {
	s.UserId = &v
	return s
}

func (s *UpdateComputeResourceAuthUserMappingsRequestUpserts) SetUsername(v string) *UpdateComputeResourceAuthUserMappingsRequestUpserts {
	s.Username = &v
	return s
}

func (s *UpdateComputeResourceAuthUserMappingsRequestUpserts) Validate() error {
	return dara.Validate(s)
}
