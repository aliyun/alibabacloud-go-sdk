// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceGroup(v *GetDeviceGroupResponseBodyDeviceGroup) *GetDeviceGroupResponseBody
	GetDeviceGroup() *GetDeviceGroupResponseBodyDeviceGroup
	SetRequestId(v string) *GetDeviceGroupResponseBody
	GetRequestId() *string
}

type GetDeviceGroupResponseBody struct {
	// The device label details.
	DeviceGroup *GetDeviceGroupResponseBodyDeviceGroup `json:"DeviceGroup,omitempty" xml:"DeviceGroup,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// BE4FB974-11BC-5453-9BE1-1606A73EACA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeviceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceGroupResponseBody) GetDeviceGroup() *GetDeviceGroupResponseBodyDeviceGroup {
	return s.DeviceGroup
}

func (s *GetDeviceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeviceGroupResponseBody) SetDeviceGroup(v *GetDeviceGroupResponseBodyDeviceGroup) *GetDeviceGroupResponseBody {
	s.DeviceGroup = v
	return s
}

func (s *GetDeviceGroupResponseBody) SetRequestId(v string) *GetDeviceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceGroupResponseBody) Validate() error {
	if s.DeviceGroup != nil {
		if err := s.DeviceGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDeviceGroupResponseBodyDeviceGroup struct {
	// The device label description.
	//
	// example:
	//
	// Test device group description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The device label ID.
	//
	// example:
	//
	// device-group-5191cf830a5e****
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	// Deprecated
	//
	// The dynamic device group rule operator.
	//
	// example:
	//
	// AND
	DynamicOperator *string `json:"DynamicOperator,omitempty" xml:"DynamicOperator,omitempty"`
	// The matching rule of the dynamic device label.
	DynamicRule *Rule `json:"DynamicRule,omitempty" xml:"DynamicRule,omitempty"`
	// The device label type. Valid values:
	//
	// - **static**: Static device label. Members consist of manually added terminal devices.
	//
	// - **dynamic**: Dynamic device label.
	//
	// example:
	//
	// static
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// Indicates whether the device label is a system built-in device label. Valid values:
	//
	// - **true**: System built-in device label.
	//
	// - **false**: User-defined device label.
	//
	// example:
	//
	// true
	IsDefault *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The collection of terminal device IDs associated with the device label.
	MatchDevTags []*string `json:"MatchDevTags,omitempty" xml:"MatchDevTags,omitempty" type:"Repeated"`
	// The device label name.
	//
	// example:
	//
	// autotest_a1b2bfd0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetDeviceGroupResponseBodyDeviceGroup) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceGroupResponseBodyDeviceGroup) GoString() string {
	return s.String()
}

func (s *GetDeviceGroupResponseBodyDeviceGroup) GetDescription() *string {
	return s.Description
}

func (s *GetDeviceGroupResponseBodyDeviceGroup) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *GetDeviceGroupResponseBodyDeviceGroup) GetDynamicOperator() *string {
	return s.DynamicOperator
}

func (s *GetDeviceGroupResponseBodyDeviceGroup) GetDynamicRule() *Rule {
	return s.DynamicRule
}

func (s *GetDeviceGroupResponseBodyDeviceGroup) GetGroupType() *string {
	return s.GroupType
}

func (s *GetDeviceGroupResponseBodyDeviceGroup) GetIsDefault() *string {
	return s.IsDefault
}

func (s *GetDeviceGroupResponseBodyDeviceGroup) GetMatchDevTags() []*string {
	return s.MatchDevTags
}

func (s *GetDeviceGroupResponseBodyDeviceGroup) GetName() *string {
	return s.Name
}

func (s *GetDeviceGroupResponseBodyDeviceGroup) SetDescription(v string) *GetDeviceGroupResponseBodyDeviceGroup {
	s.Description = &v
	return s
}

func (s *GetDeviceGroupResponseBodyDeviceGroup) SetDeviceGroupId(v string) *GetDeviceGroupResponseBodyDeviceGroup {
	s.DeviceGroupId = &v
	return s
}

func (s *GetDeviceGroupResponseBodyDeviceGroup) SetDynamicOperator(v string) *GetDeviceGroupResponseBodyDeviceGroup {
	s.DynamicOperator = &v
	return s
}

func (s *GetDeviceGroupResponseBodyDeviceGroup) SetDynamicRule(v *Rule) *GetDeviceGroupResponseBodyDeviceGroup {
	s.DynamicRule = v
	return s
}

func (s *GetDeviceGroupResponseBodyDeviceGroup) SetGroupType(v string) *GetDeviceGroupResponseBodyDeviceGroup {
	s.GroupType = &v
	return s
}

func (s *GetDeviceGroupResponseBodyDeviceGroup) SetIsDefault(v string) *GetDeviceGroupResponseBodyDeviceGroup {
	s.IsDefault = &v
	return s
}

func (s *GetDeviceGroupResponseBodyDeviceGroup) SetMatchDevTags(v []*string) *GetDeviceGroupResponseBodyDeviceGroup {
	s.MatchDevTags = v
	return s
}

func (s *GetDeviceGroupResponseBodyDeviceGroup) SetName(v string) *GetDeviceGroupResponseBodyDeviceGroup {
	s.Name = &v
	return s
}

func (s *GetDeviceGroupResponseBodyDeviceGroup) Validate() error {
	if s.DynamicRule != nil {
		if err := s.DynamicRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}
