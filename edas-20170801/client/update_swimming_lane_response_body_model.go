// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSwimmingLaneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateSwimmingLaneResponseBody
	GetCode() *int32
	SetData(v *UpdateSwimmingLaneResponseBodyData) *UpdateSwimmingLaneResponseBody
	GetData() *UpdateSwimmingLaneResponseBodyData
	SetMessage(v string) *UpdateSwimmingLaneResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSwimmingLaneResponseBody
	GetRequestId() *string
}

type UpdateSwimmingLaneResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data that is returned.
	Data *UpdateSwimmingLaneResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6CB46AEA-309C-5041-9EC7-FCF4478F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSwimmingLaneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSwimmingLaneResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSwimmingLaneResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateSwimmingLaneResponseBody) GetData() *UpdateSwimmingLaneResponseBodyData {
	return s.Data
}

func (s *UpdateSwimmingLaneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSwimmingLaneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSwimmingLaneResponseBody) SetCode(v int32) *UpdateSwimmingLaneResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSwimmingLaneResponseBody) SetData(v *UpdateSwimmingLaneResponseBodyData) *UpdateSwimmingLaneResponseBody {
	s.Data = v
	return s
}

func (s *UpdateSwimmingLaneResponseBody) SetMessage(v string) *UpdateSwimmingLaneResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSwimmingLaneResponseBody) SetRequestId(v string) *UpdateSwimmingLaneResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSwimmingLaneResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSwimmingLaneResponseBodyData struct {
	// The rule of the lane.
	//
	// example:
	//
	// [{\\"condition\\":\\"AND\\",\\"enable\\":true,\\"path\\":\\"/traffictest\\",\\"priority\\":1,\\"restItems\\":[{\\"cond\\":\\"==\\",\\"datum\\":\\"testheadervalue\\",\\"name\\":\\"testheader\\",\\"operator\\":\\"rawvalue\\",\\"type\\":\\"header\\",\\"value\\":\\"testheadervalue\\"}]}]"
	EntryRule *string `json:"EntryRule,omitempty" xml:"EntryRule,omitempty"`
	// The ID of the lane group.
	//
	// example:
	//
	// 171
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the lane.
	//
	// example:
	//
	// 321
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the lane.
	//
	// example:
	//
	// test-swimlane
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// cn-beijing:qa
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The list of associations between the lane and the related application.
	SwimmingLaneAppRelationShipList []*UpdateSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList `json:"SwimmingLaneAppRelationShipList,omitempty" xml:"SwimmingLaneAppRelationShipList,omitempty" type:"Repeated"`
	// The tag of the lane.
	//
	// example:
	//
	// 2cb6b8a
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s UpdateSwimmingLaneResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateSwimmingLaneResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateSwimmingLaneResponseBodyData) GetEntryRule() *string {
	return s.EntryRule
}

func (s *UpdateSwimmingLaneResponseBodyData) GetGroupId() *int64 {
	return s.GroupId
}

func (s *UpdateSwimmingLaneResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *UpdateSwimmingLaneResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateSwimmingLaneResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateSwimmingLaneResponseBodyData) GetSwimmingLaneAppRelationShipList() []*UpdateSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList {
	return s.SwimmingLaneAppRelationShipList
}

func (s *UpdateSwimmingLaneResponseBodyData) GetTag() *string {
	return s.Tag
}

func (s *UpdateSwimmingLaneResponseBodyData) SetEntryRule(v string) *UpdateSwimmingLaneResponseBodyData {
	s.EntryRule = &v
	return s
}

func (s *UpdateSwimmingLaneResponseBodyData) SetGroupId(v int64) *UpdateSwimmingLaneResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *UpdateSwimmingLaneResponseBodyData) SetId(v int64) *UpdateSwimmingLaneResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateSwimmingLaneResponseBodyData) SetName(v string) *UpdateSwimmingLaneResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateSwimmingLaneResponseBodyData) SetNamespaceId(v string) *UpdateSwimmingLaneResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *UpdateSwimmingLaneResponseBodyData) SetSwimmingLaneAppRelationShipList(v []*UpdateSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) *UpdateSwimmingLaneResponseBodyData {
	s.SwimmingLaneAppRelationShipList = v
	return s
}

func (s *UpdateSwimmingLaneResponseBodyData) SetTag(v string) *UpdateSwimmingLaneResponseBodyData {
	s.Tag = &v
	return s
}

func (s *UpdateSwimmingLaneResponseBodyData) Validate() error {
	if s.SwimmingLaneAppRelationShipList != nil {
		for _, item := range s.SwimmingLaneAppRelationShipList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList struct {
	// The ID of the application.
	//
	// example:
	//
	// 476d26d9-b54c-40b8-8af9-d898cdc2****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the lane.
	//
	// example:
	//
	// 321
	LaneId *int64 `json:"LaneId,omitempty" xml:"LaneId,omitempty"`
	// The association rule.
	//
	// example:
	//
	// dubbo
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s UpdateSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) String() string {
	return dara.Prettify(s)
}

func (s UpdateSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) GoString() string {
	return s.String()
}

func (s *UpdateSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) GetAppId() *string {
	return s.AppId
}

func (s *UpdateSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) GetAppName() *string {
	return s.AppName
}

func (s *UpdateSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) GetLaneId() *int64 {
	return s.LaneId
}

func (s *UpdateSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) GetRules() *string {
	return s.Rules
}

func (s *UpdateSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) SetAppId(v string) *UpdateSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList {
	s.AppId = &v
	return s
}

func (s *UpdateSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) SetAppName(v string) *UpdateSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList {
	s.AppName = &v
	return s
}

func (s *UpdateSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) SetLaneId(v int64) *UpdateSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList {
	s.LaneId = &v
	return s
}

func (s *UpdateSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) SetRules(v string) *UpdateSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList {
	s.Rules = &v
	return s
}

func (s *UpdateSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) Validate() error {
	return dara.Validate(s)
}
