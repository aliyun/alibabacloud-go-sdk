// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateContactGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertContactGroup(v *CreateOrUpdateContactGroupResponseBodyAlertContactGroup) *CreateOrUpdateContactGroupResponseBody
	GetAlertContactGroup() *CreateOrUpdateContactGroupResponseBodyAlertContactGroup
	SetRequestId(v string) *CreateOrUpdateContactGroupResponseBody
	GetRequestId() *string
}

type CreateOrUpdateContactGroupResponseBody struct {
	// The information about the alert contact group.
	AlertContactGroup *CreateOrUpdateContactGroupResponseBodyAlertContactGroup `json:"AlertContactGroup,omitempty" xml:"AlertContactGroup,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9319A57D-2D9E-472A-B69B-CF3CD16D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrUpdateContactGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateContactGroupResponseBody) GetAlertContactGroup() *CreateOrUpdateContactGroupResponseBodyAlertContactGroup {
	return s.AlertContactGroup
}

func (s *CreateOrUpdateContactGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrUpdateContactGroupResponseBody) SetAlertContactGroup(v *CreateOrUpdateContactGroupResponseBodyAlertContactGroup) *CreateOrUpdateContactGroupResponseBody {
	s.AlertContactGroup = v
	return s
}

func (s *CreateOrUpdateContactGroupResponseBody) SetRequestId(v string) *CreateOrUpdateContactGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrUpdateContactGroupResponseBody) Validate() error {
	if s.AlertContactGroup != nil {
		if err := s.AlertContactGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOrUpdateContactGroupResponseBodyAlertContactGroup struct {
	// The ID of the alert contact group.
	//
	// example:
	//
	// 123
	ContactGroupId *float32 `json:"ContactGroupId,omitempty" xml:"ContactGroupId,omitempty"`
	// The name of the alert contact group.
	//
	// example:
	//
	// TestGroup
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	// The IDs of the contacts that are included in the alert contact group.
	//
	// example:
	//
	// [1,2,3]
	ContactIds *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
}

func (s CreateOrUpdateContactGroupResponseBodyAlertContactGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateContactGroupResponseBodyAlertContactGroup) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateContactGroupResponseBodyAlertContactGroup) GetContactGroupId() *float32 {
	return s.ContactGroupId
}

func (s *CreateOrUpdateContactGroupResponseBodyAlertContactGroup) GetContactGroupName() *string {
	return s.ContactGroupName
}

func (s *CreateOrUpdateContactGroupResponseBodyAlertContactGroup) GetContactIds() *string {
	return s.ContactIds
}

func (s *CreateOrUpdateContactGroupResponseBodyAlertContactGroup) SetContactGroupId(v float32) *CreateOrUpdateContactGroupResponseBodyAlertContactGroup {
	s.ContactGroupId = &v
	return s
}

func (s *CreateOrUpdateContactGroupResponseBodyAlertContactGroup) SetContactGroupName(v string) *CreateOrUpdateContactGroupResponseBodyAlertContactGroup {
	s.ContactGroupName = &v
	return s
}

func (s *CreateOrUpdateContactGroupResponseBodyAlertContactGroup) SetContactIds(v string) *CreateOrUpdateContactGroupResponseBodyAlertContactGroup {
	s.ContactIds = &v
	return s
}

func (s *CreateOrUpdateContactGroupResponseBodyAlertContactGroup) Validate() error {
	return dara.Validate(s)
}
