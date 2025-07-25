// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudGtmAvailableAlertGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertGroups(v *ListCloudGtmAvailableAlertGroupsResponseBodyAlertGroups) *ListCloudGtmAvailableAlertGroupsResponseBody
	GetAlertGroups() *ListCloudGtmAvailableAlertGroupsResponseBodyAlertGroups
	SetRequestId(v string) *ListCloudGtmAvailableAlertGroupsResponseBody
	GetRequestId() *string
}

type ListCloudGtmAvailableAlertGroupsResponseBody struct {
	// The alert contact groups.
	AlertGroups *ListCloudGtmAvailableAlertGroupsResponseBodyAlertGroups `json:"AlertGroups,omitempty" xml:"AlertGroups,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCloudGtmAvailableAlertGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAvailableAlertGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAvailableAlertGroupsResponseBody) GetAlertGroups() *ListCloudGtmAvailableAlertGroupsResponseBodyAlertGroups {
	return s.AlertGroups
}

func (s *ListCloudGtmAvailableAlertGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloudGtmAvailableAlertGroupsResponseBody) SetAlertGroups(v *ListCloudGtmAvailableAlertGroupsResponseBodyAlertGroups) *ListCloudGtmAvailableAlertGroupsResponseBody {
	s.AlertGroups = v
	return s
}

func (s *ListCloudGtmAvailableAlertGroupsResponseBody) SetRequestId(v string) *ListCloudGtmAvailableAlertGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudGtmAvailableAlertGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCloudGtmAvailableAlertGroupsResponseBodyAlertGroups struct {
	AlertGroup []*ListCloudGtmAvailableAlertGroupsResponseBodyAlertGroupsAlertGroup `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty" type:"Repeated"`
}

func (s ListCloudGtmAvailableAlertGroupsResponseBodyAlertGroups) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAvailableAlertGroupsResponseBodyAlertGroups) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAvailableAlertGroupsResponseBodyAlertGroups) GetAlertGroup() []*ListCloudGtmAvailableAlertGroupsResponseBodyAlertGroupsAlertGroup {
	return s.AlertGroup
}

func (s *ListCloudGtmAvailableAlertGroupsResponseBodyAlertGroups) SetAlertGroup(v []*ListCloudGtmAvailableAlertGroupsResponseBodyAlertGroupsAlertGroup) *ListCloudGtmAvailableAlertGroupsResponseBodyAlertGroups {
	s.AlertGroup = v
	return s
}

func (s *ListCloudGtmAvailableAlertGroupsResponseBodyAlertGroups) Validate() error {
	return dara.Validate(s)
}

type ListCloudGtmAvailableAlertGroupsResponseBodyAlertGroupsAlertGroup struct {
	// The name of the alert contact group.
	//
	// example:
	//
	// [\\"Default\\"]
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s ListCloudGtmAvailableAlertGroupsResponseBodyAlertGroupsAlertGroup) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAvailableAlertGroupsResponseBodyAlertGroupsAlertGroup) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAvailableAlertGroupsResponseBodyAlertGroupsAlertGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *ListCloudGtmAvailableAlertGroupsResponseBodyAlertGroupsAlertGroup) SetGroupName(v string) *ListCloudGtmAvailableAlertGroupsResponseBodyAlertGroupsAlertGroup {
	s.GroupName = &v
	return s
}

func (s *ListCloudGtmAvailableAlertGroupsResponseBodyAlertGroupsAlertGroup) Validate() error {
	return dara.Validate(s)
}
