// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateContactGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroupId(v int64) *CreateOrUpdateContactGroupRequest
	GetContactGroupId() *int64
	SetContactGroupName(v string) *CreateOrUpdateContactGroupRequest
	GetContactGroupName() *string
	SetContactIds(v string) *CreateOrUpdateContactGroupRequest
	GetContactIds() *string
}

type CreateOrUpdateContactGroupRequest struct {
	// The ID of the alert contact group.
	//
	// 	- If you do not specify this parameter, an alert contact group is created.
	//
	// 	- If you specify this parameter, the specified alert contact group is modified.
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
	// The ID of the contact that you want to add to the contact group. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// [1,2,3]
	ContactIds *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
}

func (s CreateOrUpdateContactGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateContactGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateContactGroupRequest) GetContactGroupId() *int64 {
	return s.ContactGroupId
}

func (s *CreateOrUpdateContactGroupRequest) GetContactGroupName() *string {
	return s.ContactGroupName
}

func (s *CreateOrUpdateContactGroupRequest) GetContactIds() *string {
	return s.ContactIds
}

func (s *CreateOrUpdateContactGroupRequest) SetContactGroupId(v int64) *CreateOrUpdateContactGroupRequest {
	s.ContactGroupId = &v
	return s
}

func (s *CreateOrUpdateContactGroupRequest) SetContactGroupName(v string) *CreateOrUpdateContactGroupRequest {
	s.ContactGroupName = &v
	return s
}

func (s *CreateOrUpdateContactGroupRequest) SetContactIds(v string) *CreateOrUpdateContactGroupRequest {
	s.ContactIds = &v
	return s
}

func (s *CreateOrUpdateContactGroupRequest) Validate() error {
	return dara.Validate(s)
}
