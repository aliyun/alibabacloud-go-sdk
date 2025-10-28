// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListServiceGroupsResponseBody
	GetCode() *int32
	SetMessage(v string) *ListServiceGroupsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListServiceGroupsResponseBody
	GetRequestId() *string
	SetServiceGroupsList(v *ListServiceGroupsResponseBodyServiceGroupsList) *ListServiceGroupsResponseBody
	GetServiceGroupsList() *ListServiceGroupsResponseBodyServiceGroupsList
}

type ListServiceGroupsResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// a5281053-08e4-47a5-b2ab-5c0323de7b5a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about service groups.
	ServiceGroupsList *ListServiceGroupsResponseBodyServiceGroupsList `json:"ServiceGroupsList,omitempty" xml:"ServiceGroupsList,omitempty" type:"Struct"`
}

func (s ListServiceGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceGroupsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListServiceGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListServiceGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceGroupsResponseBody) GetServiceGroupsList() *ListServiceGroupsResponseBodyServiceGroupsList {
	return s.ServiceGroupsList
}

func (s *ListServiceGroupsResponseBody) SetCode(v int32) *ListServiceGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListServiceGroupsResponseBody) SetMessage(v string) *ListServiceGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListServiceGroupsResponseBody) SetRequestId(v string) *ListServiceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceGroupsResponseBody) SetServiceGroupsList(v *ListServiceGroupsResponseBodyServiceGroupsList) *ListServiceGroupsResponseBody {
	s.ServiceGroupsList = v
	return s
}

func (s *ListServiceGroupsResponseBody) Validate() error {
	if s.ServiceGroupsList != nil {
		if err := s.ServiceGroupsList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListServiceGroupsResponseBodyServiceGroupsList struct {
	ListServiceGroups []*ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups `json:"ListServiceGroups,omitempty" xml:"ListServiceGroups,omitempty" type:"Repeated"`
}

func (s ListServiceGroupsResponseBodyServiceGroupsList) String() string {
	return dara.Prettify(s)
}

func (s ListServiceGroupsResponseBodyServiceGroupsList) GoString() string {
	return s.String()
}

func (s *ListServiceGroupsResponseBodyServiceGroupsList) GetListServiceGroups() []*ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups {
	return s.ListServiceGroups
}

func (s *ListServiceGroupsResponseBodyServiceGroupsList) SetListServiceGroups(v []*ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups) *ListServiceGroupsResponseBodyServiceGroupsList {
	s.ListServiceGroups = v
	return s
}

func (s *ListServiceGroupsResponseBodyServiceGroupsList) Validate() error {
	if s.ListServiceGroups != nil {
		for _, item := range s.ListServiceGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups struct {
	// The time when the service group was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1575357165770
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the service group.
	//
	// example:
	//
	// 789d9cda-74b1-****-****-05e21a0a7661
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the service group.
	//
	// example:
	//
	// edas-test-group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups) String() string {
	return dara.Prettify(s)
}

func (s ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups) GoString() string {
	return s.String()
}

func (s *ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups) SetCreateTime(v string) *ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups {
	s.CreateTime = &v
	return s
}

func (s *ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups) SetGroupId(v string) *ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups {
	s.GroupId = &v
	return s
}

func (s *ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups) SetGroupName(v string) *ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups {
	s.GroupName = &v
	return s
}

func (s *ListServiceGroupsResponseBodyServiceGroupsListListServiceGroups) Validate() error {
	return dara.Validate(s)
}
