// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitorGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteMonitorGroupResponseBody
	GetCode() *int32
	SetGroup(v *DeleteMonitorGroupResponseBodyGroup) *DeleteMonitorGroupResponseBody
	GetGroup() *DeleteMonitorGroupResponseBodyGroup
	SetMessage(v string) *DeleteMonitorGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMonitorGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMonitorGroupResponseBody
	GetSuccess() *bool
}

type DeleteMonitorGroupResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the call was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The deleted application group.
	Group *DeleteMonitorGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CA35B3AE-4FFD-4A33-AE67-67EF68711EFA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. The value true indicates a success. The value false indicates a failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMonitorGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteMonitorGroupResponseBody) GetGroup() *DeleteMonitorGroupResponseBodyGroup {
	return s.Group
}

func (s *DeleteMonitorGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMonitorGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMonitorGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMonitorGroupResponseBody) SetCode(v int32) *DeleteMonitorGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMonitorGroupResponseBody) SetGroup(v *DeleteMonitorGroupResponseBodyGroup) *DeleteMonitorGroupResponseBody {
	s.Group = v
	return s
}

func (s *DeleteMonitorGroupResponseBody) SetMessage(v string) *DeleteMonitorGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMonitorGroupResponseBody) SetRequestId(v string) *DeleteMonitorGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMonitorGroupResponseBody) SetSuccess(v bool) *DeleteMonitorGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMonitorGroupResponseBody) Validate() error {
	if s.Group != nil {
		if err := s.Group.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMonitorGroupResponseBodyGroup struct {
	// The alert groups that receive alert notifications for the application group.
	ContactGroups *DeleteMonitorGroupResponseBodyGroupContactGroups `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Struct"`
	// The name of the application group.
	//
	// example:
	//
	// ECS_test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s DeleteMonitorGroupResponseBodyGroup) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorGroupResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupResponseBodyGroup) GetContactGroups() *DeleteMonitorGroupResponseBodyGroupContactGroups {
	return s.ContactGroups
}

func (s *DeleteMonitorGroupResponseBodyGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *DeleteMonitorGroupResponseBodyGroup) SetContactGroups(v *DeleteMonitorGroupResponseBodyGroupContactGroups) *DeleteMonitorGroupResponseBodyGroup {
	s.ContactGroups = v
	return s
}

func (s *DeleteMonitorGroupResponseBodyGroup) SetGroupName(v string) *DeleteMonitorGroupResponseBodyGroup {
	s.GroupName = &v
	return s
}

func (s *DeleteMonitorGroupResponseBodyGroup) Validate() error {
	if s.ContactGroups != nil {
		if err := s.ContactGroups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMonitorGroupResponseBodyGroupContactGroups struct {
	ContactGroup []*DeleteMonitorGroupResponseBodyGroupContactGroupsContactGroup `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty" type:"Repeated"`
}

func (s DeleteMonitorGroupResponseBodyGroupContactGroups) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorGroupResponseBodyGroupContactGroups) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupResponseBodyGroupContactGroups) GetContactGroup() []*DeleteMonitorGroupResponseBodyGroupContactGroupsContactGroup {
	return s.ContactGroup
}

func (s *DeleteMonitorGroupResponseBodyGroupContactGroups) SetContactGroup(v []*DeleteMonitorGroupResponseBodyGroupContactGroupsContactGroup) *DeleteMonitorGroupResponseBodyGroupContactGroups {
	s.ContactGroup = v
	return s
}

func (s *DeleteMonitorGroupResponseBodyGroupContactGroups) Validate() error {
	if s.ContactGroup != nil {
		for _, item := range s.ContactGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteMonitorGroupResponseBodyGroupContactGroupsContactGroup struct {
	// The name of the alert group.
	//
	// example:
	//
	// ECS_Group
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteMonitorGroupResponseBodyGroupContactGroupsContactGroup) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorGroupResponseBodyGroupContactGroupsContactGroup) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupResponseBodyGroupContactGroupsContactGroup) GetName() *string {
	return s.Name
}

func (s *DeleteMonitorGroupResponseBodyGroupContactGroupsContactGroup) SetName(v string) *DeleteMonitorGroupResponseBodyGroupContactGroupsContactGroup {
	s.Name = &v
	return s
}

func (s *DeleteMonitorGroupResponseBodyGroupContactGroupsContactGroup) Validate() error {
	return dara.Validate(s)
}
