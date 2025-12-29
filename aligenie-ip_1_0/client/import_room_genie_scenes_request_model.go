// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportRoomGenieScenesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *ImportRoomGenieScenesRequest
	GetHotelId() *string
	SetRoomNo(v string) *ImportRoomGenieScenesRequest
	GetRoomNo() *string
	SetSceneList(v []*ImportRoomGenieScenesRequestSceneList) *ImportRoomGenieScenesRequest
	GetSceneList() []*ImportRoomGenieScenesRequestSceneList
}

type ImportRoomGenieScenesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a7a3***013
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1211
	RoomNo    *string                                  `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	SceneList []*ImportRoomGenieScenesRequestSceneList `json:"SceneList,omitempty" xml:"SceneList,omitempty" type:"Repeated"`
}

func (s ImportRoomGenieScenesRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportRoomGenieScenesRequest) GoString() string {
	return s.String()
}

func (s *ImportRoomGenieScenesRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ImportRoomGenieScenesRequest) GetRoomNo() *string {
	return s.RoomNo
}

func (s *ImportRoomGenieScenesRequest) GetSceneList() []*ImportRoomGenieScenesRequestSceneList {
	return s.SceneList
}

func (s *ImportRoomGenieScenesRequest) SetHotelId(v string) *ImportRoomGenieScenesRequest {
	s.HotelId = &v
	return s
}

func (s *ImportRoomGenieScenesRequest) SetRoomNo(v string) *ImportRoomGenieScenesRequest {
	s.RoomNo = &v
	return s
}

func (s *ImportRoomGenieScenesRequest) SetSceneList(v []*ImportRoomGenieScenesRequestSceneList) *ImportRoomGenieScenesRequest {
	s.SceneList = v
	return s
}

func (s *ImportRoomGenieScenesRequest) Validate() error {
	if s.SceneList != nil {
		for _, item := range s.SceneList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImportRoomGenieScenesRequestSceneList struct {
	// This parameter is required.
	Actions     []*ImportRoomGenieScenesRequestSceneListActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	Description *string                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Display *bool `json:"Display,omitempty" xml:"Display,omitempty"`
	// example:
	//
	// http://xxx.com/yyy.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// This parameter is required.
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TriggerLogical *int32 `json:"TriggerLogical,omitempty" xml:"TriggerLogical,omitempty"`
	// This parameter is required.
	Triggers []*ImportRoomGenieScenesRequestSceneListTriggers `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s ImportRoomGenieScenesRequestSceneList) String() string {
	return dara.Prettify(s)
}

func (s ImportRoomGenieScenesRequestSceneList) GoString() string {
	return s.String()
}

func (s *ImportRoomGenieScenesRequestSceneList) GetActions() []*ImportRoomGenieScenesRequestSceneListActions {
	return s.Actions
}

func (s *ImportRoomGenieScenesRequestSceneList) GetDescription() *string {
	return s.Description
}

func (s *ImportRoomGenieScenesRequestSceneList) GetDisplay() *bool {
	return s.Display
}

func (s *ImportRoomGenieScenesRequestSceneList) GetIcon() *string {
	return s.Icon
}

func (s *ImportRoomGenieScenesRequestSceneList) GetSceneName() *string {
	return s.SceneName
}

func (s *ImportRoomGenieScenesRequestSceneList) GetTriggerLogical() *int32 {
	return s.TriggerLogical
}

func (s *ImportRoomGenieScenesRequestSceneList) GetTriggers() []*ImportRoomGenieScenesRequestSceneListTriggers {
	return s.Triggers
}

func (s *ImportRoomGenieScenesRequestSceneList) SetActions(v []*ImportRoomGenieScenesRequestSceneListActions) *ImportRoomGenieScenesRequestSceneList {
	s.Actions = v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneList) SetDescription(v string) *ImportRoomGenieScenesRequestSceneList {
	s.Description = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneList) SetDisplay(v bool) *ImportRoomGenieScenesRequestSceneList {
	s.Display = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneList) SetIcon(v string) *ImportRoomGenieScenesRequestSceneList {
	s.Icon = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneList) SetSceneName(v string) *ImportRoomGenieScenesRequestSceneList {
	s.SceneName = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneList) SetTriggerLogical(v int32) *ImportRoomGenieScenesRequestSceneList {
	s.TriggerLogical = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneList) SetTriggers(v []*ImportRoomGenieScenesRequestSceneListTriggers) *ImportRoomGenieScenesRequestSceneList {
	s.Triggers = v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneList) Validate() error {
	if s.Actions != nil {
		for _, item := range s.Actions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Triggers != nil {
		for _, item := range s.Triggers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImportRoomGenieScenesRequestSceneListActions struct {
	AttributeValues []*ImportRoomGenieScenesRequestSceneListActionsAttributeValues `json:"AttributeValues,omitempty" xml:"AttributeValues,omitempty" type:"Repeated"`
	Device          *ImportRoomGenieScenesRequestSceneListActionsDevice            `json:"Device,omitempty" xml:"Device,omitempty" type:"Struct"`
	Reply           *string                                                        `json:"Reply,omitempty" xml:"Reply,omitempty"`
	Type            *int32                                                         `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ImportRoomGenieScenesRequestSceneListActions) String() string {
	return dara.Prettify(s)
}

func (s ImportRoomGenieScenesRequestSceneListActions) GoString() string {
	return s.String()
}

func (s *ImportRoomGenieScenesRequestSceneListActions) GetAttributeValues() []*ImportRoomGenieScenesRequestSceneListActionsAttributeValues {
	return s.AttributeValues
}

func (s *ImportRoomGenieScenesRequestSceneListActions) GetDevice() *ImportRoomGenieScenesRequestSceneListActionsDevice {
	return s.Device
}

func (s *ImportRoomGenieScenesRequestSceneListActions) GetReply() *string {
	return s.Reply
}

func (s *ImportRoomGenieScenesRequestSceneListActions) GetType() *int32 {
	return s.Type
}

func (s *ImportRoomGenieScenesRequestSceneListActions) SetAttributeValues(v []*ImportRoomGenieScenesRequestSceneListActionsAttributeValues) *ImportRoomGenieScenesRequestSceneListActions {
	s.AttributeValues = v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListActions) SetDevice(v *ImportRoomGenieScenesRequestSceneListActionsDevice) *ImportRoomGenieScenesRequestSceneListActions {
	s.Device = v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListActions) SetReply(v string) *ImportRoomGenieScenesRequestSceneListActions {
	s.Reply = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListActions) SetType(v int32) *ImportRoomGenieScenesRequestSceneListActions {
	s.Type = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListActions) Validate() error {
	if s.AttributeValues != nil {
		for _, item := range s.AttributeValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Device != nil {
		if err := s.Device.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImportRoomGenieScenesRequestSceneListActionsAttributeValues struct {
	// This parameter is required.
	//
	// example:
	//
	// powerstate
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	AttributeValue *string `json:"AttributeValue,omitempty" xml:"AttributeValue,omitempty"`
}

func (s ImportRoomGenieScenesRequestSceneListActionsAttributeValues) String() string {
	return dara.Prettify(s)
}

func (s ImportRoomGenieScenesRequestSceneListActionsAttributeValues) GoString() string {
	return s.String()
}

func (s *ImportRoomGenieScenesRequestSceneListActionsAttributeValues) GetAttributeName() *string {
	return s.AttributeName
}

func (s *ImportRoomGenieScenesRequestSceneListActionsAttributeValues) GetAttributeValue() *string {
	return s.AttributeValue
}

func (s *ImportRoomGenieScenesRequestSceneListActionsAttributeValues) SetAttributeName(v string) *ImportRoomGenieScenesRequestSceneListActionsAttributeValues {
	s.AttributeName = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListActionsAttributeValues) SetAttributeValue(v string) *ImportRoomGenieScenesRequestSceneListActionsAttributeValues {
	s.AttributeValue = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListActionsAttributeValues) Validate() error {
	return dara.Validate(s)
}

type ImportRoomGenieScenesRequestSceneListActionsDevice struct {
	// This parameter is required.
	//
	// example:
	//
	// light
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// 0
	DeviceIndex *int32 `json:"DeviceIndex,omitempty" xml:"DeviceIndex,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3c5d2ab8f9ec
	DeviceNumber *string `json:"DeviceNumber,omitempty" xml:"DeviceNumber,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ImportRoomGenieScenesRequestSceneListActionsDevice) String() string {
	return dara.Prettify(s)
}

func (s ImportRoomGenieScenesRequestSceneListActionsDevice) GoString() string {
	return s.String()
}

func (s *ImportRoomGenieScenesRequestSceneListActionsDevice) GetCategory() *string {
	return s.Category
}

func (s *ImportRoomGenieScenesRequestSceneListActionsDevice) GetDeviceIndex() *int32 {
	return s.DeviceIndex
}

func (s *ImportRoomGenieScenesRequestSceneListActionsDevice) GetDeviceNumber() *string {
	return s.DeviceNumber
}

func (s *ImportRoomGenieScenesRequestSceneListActionsDevice) GetType() *int32 {
	return s.Type
}

func (s *ImportRoomGenieScenesRequestSceneListActionsDevice) SetCategory(v string) *ImportRoomGenieScenesRequestSceneListActionsDevice {
	s.Category = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListActionsDevice) SetDeviceIndex(v int32) *ImportRoomGenieScenesRequestSceneListActionsDevice {
	s.DeviceIndex = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListActionsDevice) SetDeviceNumber(v string) *ImportRoomGenieScenesRequestSceneListActionsDevice {
	s.DeviceNumber = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListActionsDevice) SetType(v int32) *ImportRoomGenieScenesRequestSceneListActionsDevice {
	s.Type = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListActionsDevice) Validate() error {
	return dara.Validate(s)
}

type ImportRoomGenieScenesRequestSceneListTriggers struct {
	AttributeValues []*ImportRoomGenieScenesRequestSceneListTriggersAttributeValues `json:"AttributeValues,omitempty" xml:"AttributeValues,omitempty" type:"Repeated"`
	CorpusList      []*string                                                       `json:"CorpusList,omitempty" xml:"CorpusList,omitempty" type:"Repeated"`
	Device          *ImportRoomGenieScenesRequestSceneListTriggersDevice            `json:"Device,omitempty" xml:"Device,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ImportRoomGenieScenesRequestSceneListTriggers) String() string {
	return dara.Prettify(s)
}

func (s ImportRoomGenieScenesRequestSceneListTriggers) GoString() string {
	return s.String()
}

func (s *ImportRoomGenieScenesRequestSceneListTriggers) GetAttributeValues() []*ImportRoomGenieScenesRequestSceneListTriggersAttributeValues {
	return s.AttributeValues
}

func (s *ImportRoomGenieScenesRequestSceneListTriggers) GetCorpusList() []*string {
	return s.CorpusList
}

func (s *ImportRoomGenieScenesRequestSceneListTriggers) GetDevice() *ImportRoomGenieScenesRequestSceneListTriggersDevice {
	return s.Device
}

func (s *ImportRoomGenieScenesRequestSceneListTriggers) GetType() *int32 {
	return s.Type
}

func (s *ImportRoomGenieScenesRequestSceneListTriggers) SetAttributeValues(v []*ImportRoomGenieScenesRequestSceneListTriggersAttributeValues) *ImportRoomGenieScenesRequestSceneListTriggers {
	s.AttributeValues = v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListTriggers) SetCorpusList(v []*string) *ImportRoomGenieScenesRequestSceneListTriggers {
	s.CorpusList = v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListTriggers) SetDevice(v *ImportRoomGenieScenesRequestSceneListTriggersDevice) *ImportRoomGenieScenesRequestSceneListTriggers {
	s.Device = v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListTriggers) SetType(v int32) *ImportRoomGenieScenesRequestSceneListTriggers {
	s.Type = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListTriggers) Validate() error {
	if s.AttributeValues != nil {
		for _, item := range s.AttributeValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Device != nil {
		if err := s.Device.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImportRoomGenieScenesRequestSceneListTriggersAttributeValues struct {
	// This parameter is required.
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	// This parameter is required.
	AttributeValue *string `json:"AttributeValue,omitempty" xml:"AttributeValue,omitempty"`
}

func (s ImportRoomGenieScenesRequestSceneListTriggersAttributeValues) String() string {
	return dara.Prettify(s)
}

func (s ImportRoomGenieScenesRequestSceneListTriggersAttributeValues) GoString() string {
	return s.String()
}

func (s *ImportRoomGenieScenesRequestSceneListTriggersAttributeValues) GetAttributeName() *string {
	return s.AttributeName
}

func (s *ImportRoomGenieScenesRequestSceneListTriggersAttributeValues) GetAttributeValue() *string {
	return s.AttributeValue
}

func (s *ImportRoomGenieScenesRequestSceneListTriggersAttributeValues) SetAttributeName(v string) *ImportRoomGenieScenesRequestSceneListTriggersAttributeValues {
	s.AttributeName = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListTriggersAttributeValues) SetAttributeValue(v string) *ImportRoomGenieScenesRequestSceneListTriggersAttributeValues {
	s.AttributeValue = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListTriggersAttributeValues) Validate() error {
	return dara.Validate(s)
}

type ImportRoomGenieScenesRequestSceneListTriggersDevice struct {
	// This parameter is required.
	Category    *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DeviceIndex *string `json:"DeviceIndex,omitempty" xml:"DeviceIndex,omitempty"`
	// This parameter is required.
	DeviceNumber *string `json:"DeviceNumber,omitempty" xml:"DeviceNumber,omitempty"`
}

func (s ImportRoomGenieScenesRequestSceneListTriggersDevice) String() string {
	return dara.Prettify(s)
}

func (s ImportRoomGenieScenesRequestSceneListTriggersDevice) GoString() string {
	return s.String()
}

func (s *ImportRoomGenieScenesRequestSceneListTriggersDevice) GetCategory() *string {
	return s.Category
}

func (s *ImportRoomGenieScenesRequestSceneListTriggersDevice) GetDeviceIndex() *string {
	return s.DeviceIndex
}

func (s *ImportRoomGenieScenesRequestSceneListTriggersDevice) GetDeviceNumber() *string {
	return s.DeviceNumber
}

func (s *ImportRoomGenieScenesRequestSceneListTriggersDevice) SetCategory(v string) *ImportRoomGenieScenesRequestSceneListTriggersDevice {
	s.Category = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListTriggersDevice) SetDeviceIndex(v string) *ImportRoomGenieScenesRequestSceneListTriggersDevice {
	s.DeviceIndex = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListTriggersDevice) SetDeviceNumber(v string) *ImportRoomGenieScenesRequestSceneListTriggersDevice {
	s.DeviceNumber = &v
	return s
}

func (s *ImportRoomGenieScenesRequestSceneListTriggersDevice) Validate() error {
	return dara.Validate(s)
}
