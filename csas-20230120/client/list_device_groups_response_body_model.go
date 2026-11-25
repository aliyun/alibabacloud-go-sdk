// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceGroups(v []*ListDeviceGroupsResponseBodyDeviceGroups) *ListDeviceGroupsResponseBody
	GetDeviceGroups() []*ListDeviceGroupsResponseBodyDeviceGroups
	SetRequestId(v string) *ListDeviceGroupsResponseBody
	GetRequestId() *string
	SetTotalNum(v int64) *ListDeviceGroupsResponseBody
	GetTotalNum() *int64
}

type ListDeviceGroupsResponseBody struct {
	// The list of device tags.
	DeviceGroups []*ListDeviceGroupsResponseBodyDeviceGroups `json:"DeviceGroups,omitempty" xml:"DeviceGroups,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// FD724DBC-CD76-5235-BF76-59C51B73296D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of device tags.
	//
	// example:
	//
	// 1
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListDeviceGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupsResponseBody) GetDeviceGroups() []*ListDeviceGroupsResponseBodyDeviceGroups {
	return s.DeviceGroups
}

func (s *ListDeviceGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeviceGroupsResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *ListDeviceGroupsResponseBody) SetDeviceGroups(v []*ListDeviceGroupsResponseBodyDeviceGroups) *ListDeviceGroupsResponseBody {
	s.DeviceGroups = v
	return s
}

func (s *ListDeviceGroupsResponseBody) SetRequestId(v string) *ListDeviceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceGroupsResponseBody) SetTotalNum(v int64) *ListDeviceGroupsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListDeviceGroupsResponseBody) Validate() error {
	if s.DeviceGroups != nil {
		for _, item := range s.DeviceGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDeviceGroupsResponseBodyDeviceGroups struct {
	// The description of the device tag.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the device tag.
	//
	// example:
	//
	// device-group-5191cf830a5e****
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	// Deprecated
	//
	// The operator for the dynamic device group rule.
	//
	// example:
	//
	// AND
	DynamicOperator *string `json:"DynamicOperator,omitempty" xml:"DynamicOperator,omitempty"`
	// The matching rule for the dynamic device tag.
	DynamicRule *Rule `json:"DynamicRule,omitempty" xml:"DynamicRule,omitempty"`
	// The type of the device tag. Valid values:
	//
	// - **static**: Static device tag. Members consist of manually added terminal devices.
	//
	// - **dynamic**: Dynamic device tag. Members are automatically calculated based on matching rules when terminal devices report heartbeats.
	//
	// example:
	//
	// static
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// Specifies whether the device tag is a built-in system device tag. Valid values:
	//
	// - **true**: Built-in system device tag.
	//
	// - **false**: Custom device tag.
	//
	// example:
	//
	// false
	IsDefault *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The IDs of the terminal devices that are associated with the device tag.
	MatchDevTags []*string `json:"MatchDevTags,omitempty" xml:"MatchDevTags,omitempty" type:"Repeated"`
	// The name of the device tag.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListDeviceGroupsResponseBodyDeviceGroups) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceGroupsResponseBodyDeviceGroups) GoString() string {
	return s.String()
}

func (s *ListDeviceGroupsResponseBodyDeviceGroups) GetDescription() *string {
	return s.Description
}

func (s *ListDeviceGroupsResponseBodyDeviceGroups) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *ListDeviceGroupsResponseBodyDeviceGroups) GetDynamicOperator() *string {
	return s.DynamicOperator
}

func (s *ListDeviceGroupsResponseBodyDeviceGroups) GetDynamicRule() *Rule {
	return s.DynamicRule
}

func (s *ListDeviceGroupsResponseBodyDeviceGroups) GetGroupType() *string {
	return s.GroupType
}

func (s *ListDeviceGroupsResponseBodyDeviceGroups) GetIsDefault() *string {
	return s.IsDefault
}

func (s *ListDeviceGroupsResponseBodyDeviceGroups) GetMatchDevTags() []*string {
	return s.MatchDevTags
}

func (s *ListDeviceGroupsResponseBodyDeviceGroups) GetName() *string {
	return s.Name
}

func (s *ListDeviceGroupsResponseBodyDeviceGroups) SetDescription(v string) *ListDeviceGroupsResponseBodyDeviceGroups {
	s.Description = &v
	return s
}

func (s *ListDeviceGroupsResponseBodyDeviceGroups) SetDeviceGroupId(v string) *ListDeviceGroupsResponseBodyDeviceGroups {
	s.DeviceGroupId = &v
	return s
}

func (s *ListDeviceGroupsResponseBodyDeviceGroups) SetDynamicOperator(v string) *ListDeviceGroupsResponseBodyDeviceGroups {
	s.DynamicOperator = &v
	return s
}

func (s *ListDeviceGroupsResponseBodyDeviceGroups) SetDynamicRule(v *Rule) *ListDeviceGroupsResponseBodyDeviceGroups {
	s.DynamicRule = v
	return s
}

func (s *ListDeviceGroupsResponseBodyDeviceGroups) SetGroupType(v string) *ListDeviceGroupsResponseBodyDeviceGroups {
	s.GroupType = &v
	return s
}

func (s *ListDeviceGroupsResponseBodyDeviceGroups) SetIsDefault(v string) *ListDeviceGroupsResponseBodyDeviceGroups {
	s.IsDefault = &v
	return s
}

func (s *ListDeviceGroupsResponseBodyDeviceGroups) SetMatchDevTags(v []*string) *ListDeviceGroupsResponseBodyDeviceGroups {
	s.MatchDevTags = v
	return s
}

func (s *ListDeviceGroupsResponseBodyDeviceGroups) SetName(v string) *ListDeviceGroupsResponseBodyDeviceGroups {
	s.Name = &v
	return s
}

func (s *ListDeviceGroupsResponseBodyDeviceGroups) Validate() error {
	if s.DynamicRule != nil {
		if err := s.DynamicRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}
