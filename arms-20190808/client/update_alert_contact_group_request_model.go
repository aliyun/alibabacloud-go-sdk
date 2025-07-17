// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertContactGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroupId(v int64) *UpdateAlertContactGroupRequest
	GetContactGroupId() *int64
	SetContactGroupName(v string) *UpdateAlertContactGroupRequest
	GetContactGroupName() *string
	SetContactIds(v string) *UpdateAlertContactGroupRequest
	GetContactIds() *string
	SetRegionId(v string) *UpdateAlertContactGroupRequest
	GetRegionId() *string
}

type UpdateAlertContactGroupRequest struct {
	// The ID of the alert contact group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ContactGroupId *int64 `json:"ContactGroupId,omitempty" xml:"ContactGroupId,omitempty"`
	// The name of the alert contact group.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestGroup
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	// The ID of the alert contact.
	//
	// example:
	//
	// 123 234 345
	ContactIds *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateAlertContactGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertContactGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertContactGroupRequest) GetContactGroupId() *int64 {
	return s.ContactGroupId
}

func (s *UpdateAlertContactGroupRequest) GetContactGroupName() *string {
	return s.ContactGroupName
}

func (s *UpdateAlertContactGroupRequest) GetContactIds() *string {
	return s.ContactIds
}

func (s *UpdateAlertContactGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAlertContactGroupRequest) SetContactGroupId(v int64) *UpdateAlertContactGroupRequest {
	s.ContactGroupId = &v
	return s
}

func (s *UpdateAlertContactGroupRequest) SetContactGroupName(v string) *UpdateAlertContactGroupRequest {
	s.ContactGroupName = &v
	return s
}

func (s *UpdateAlertContactGroupRequest) SetContactIds(v string) *UpdateAlertContactGroupRequest {
	s.ContactIds = &v
	return s
}

func (s *UpdateAlertContactGroupRequest) SetRegionId(v string) *UpdateAlertContactGroupRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAlertContactGroupRequest) Validate() error {
	return dara.Validate(s)
}
