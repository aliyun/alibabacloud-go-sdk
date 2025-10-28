// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertSwimmingLaneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *InsertSwimmingLaneResponseBody
	GetCode() *int32
	SetData(v *InsertSwimmingLaneResponseBodyData) *InsertSwimmingLaneResponseBody
	GetData() *InsertSwimmingLaneResponseBodyData
	SetMessage(v string) *InsertSwimmingLaneResponseBody
	GetMessage() *string
	SetRequestId(v string) *InsertSwimmingLaneResponseBody
	GetRequestId() *string
}

type InsertSwimmingLaneResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data that is returned.
	Data *InsertSwimmingLaneResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 75972A87-5682-5277-ADA7-DA2A01BE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InsertSwimmingLaneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsertSwimmingLaneResponseBody) GoString() string {
	return s.String()
}

func (s *InsertSwimmingLaneResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *InsertSwimmingLaneResponseBody) GetData() *InsertSwimmingLaneResponseBodyData {
	return s.Data
}

func (s *InsertSwimmingLaneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InsertSwimmingLaneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsertSwimmingLaneResponseBody) SetCode(v int32) *InsertSwimmingLaneResponseBody {
	s.Code = &v
	return s
}

func (s *InsertSwimmingLaneResponseBody) SetData(v *InsertSwimmingLaneResponseBodyData) *InsertSwimmingLaneResponseBody {
	s.Data = v
	return s
}

func (s *InsertSwimmingLaneResponseBody) SetMessage(v string) *InsertSwimmingLaneResponseBody {
	s.Message = &v
	return s
}

func (s *InsertSwimmingLaneResponseBody) SetRequestId(v string) *InsertSwimmingLaneResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertSwimmingLaneResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InsertSwimmingLaneResponseBodyData struct {
	// The applications that are associated with lanes.
	//
	// example:
	//
	// [{"appId":"9dcba109-ee9f-4e67-8916-41
	//
	// *******"}]
	AppInfos *string `json:"AppInfos,omitempty" xml:"AppInfos,omitempty"`
	// The throttling rule for the lane.
	//
	// priority: the priority of the throttling rule for the lane. Valid values: 1 to 100.
	//
	// path: the path that is matched by the throttling rule for the lane.
	//
	// restItems: conditions to be met.
	//
	// condition: the relationship among the conditions to be met.
	//
	// 	- AND: all conditions
	//
	// 	- OR: one of the conditions
	//
	// restItems.type: the type of the rule. Valid values:
	//
	// 	- header: matches by request header.
	//
	// 	- cookie: matches by request cookie.
	//
	// 	- param: matches by request parameters.
	//
	// restItems.name: the key that matches the rule.
	//
	// restItems.value: the value that matches the rule.
	//
	// restItems.cond: the condition that matches the rule. Valid values:
	//
	// 	- "==": The parameter value is equal to the value that you enter in the Value field.
	//
	// 	- "!=": The parameter value is not equal to the value that you enter in the Value field.
	//
	// 	- ">": The parameter value is greater than the value that you enter in the Value field.
	//
	// 	- "<": The parameter value is less than the value that you enter in the Value field.
	//
	// 	- ">=": The parameter value is greater than or equal to the value that you enter in the Value field.
	//
	// 	- "<=": The parameter value is less than or equal to the value that you enter in the Value field.
	//
	// 	- "in": The parameter value is within the values that you enter in the Value field.
	//
	// restItems.operator: the type of the value. Valid values:
	//
	// 	- rawvalue: the initial value
	//
	// 	- mod: the reminder obtained by performing modulo operation
	//
	// 	- list: the value from the list
	//
	// example:
	//
	// [{\\"condition\\":\\"AND\\",\\"enable\\":false,\\"path\\":\\"/traffic\\",\\"priority\\":1,\\"restItems\\":[{\\"cond\\":\\"==\\",\\"datum\\":\\"testvalue\\",\\"name\\":\\"testheader\\",\\"operator\\":\\"rawvalue\\",\\"type\\":\\"header\\",\\"value\\":\\"testvalue\\"}]}]
	EntryRule *string `json:"EntryRule,omitempty" xml:"EntryRule,omitempty"`
	// The ID of the lane group.
	//
	// example:
	//
	// 95
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the lane.
	//
	// example:
	//
	// 88
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the lane.
	//
	// example:
	//
	// test1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The association relationships between lanes and applications.
	SwimmingLaneAppRelationShipList []*InsertSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList `json:"SwimmingLaneAppRelationShipList,omitempty" xml:"SwimmingLaneAppRelationShipList,omitempty" type:"Repeated"`
	// The tag of the lane.
	//
	// example:
	//
	// 8202e09
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s InsertSwimmingLaneResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InsertSwimmingLaneResponseBodyData) GoString() string {
	return s.String()
}

func (s *InsertSwimmingLaneResponseBodyData) GetAppInfos() *string {
	return s.AppInfos
}

func (s *InsertSwimmingLaneResponseBodyData) GetEntryRule() *string {
	return s.EntryRule
}

func (s *InsertSwimmingLaneResponseBodyData) GetGroupId() *int64 {
	return s.GroupId
}

func (s *InsertSwimmingLaneResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *InsertSwimmingLaneResponseBodyData) GetName() *string {
	return s.Name
}

func (s *InsertSwimmingLaneResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *InsertSwimmingLaneResponseBodyData) GetSwimmingLaneAppRelationShipList() []*InsertSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList {
	return s.SwimmingLaneAppRelationShipList
}

func (s *InsertSwimmingLaneResponseBodyData) GetTag() *string {
	return s.Tag
}

func (s *InsertSwimmingLaneResponseBodyData) SetAppInfos(v string) *InsertSwimmingLaneResponseBodyData {
	s.AppInfos = &v
	return s
}

func (s *InsertSwimmingLaneResponseBodyData) SetEntryRule(v string) *InsertSwimmingLaneResponseBodyData {
	s.EntryRule = &v
	return s
}

func (s *InsertSwimmingLaneResponseBodyData) SetGroupId(v int64) *InsertSwimmingLaneResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *InsertSwimmingLaneResponseBodyData) SetId(v int64) *InsertSwimmingLaneResponseBodyData {
	s.Id = &v
	return s
}

func (s *InsertSwimmingLaneResponseBodyData) SetName(v string) *InsertSwimmingLaneResponseBodyData {
	s.Name = &v
	return s
}

func (s *InsertSwimmingLaneResponseBodyData) SetNamespaceId(v string) *InsertSwimmingLaneResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *InsertSwimmingLaneResponseBodyData) SetSwimmingLaneAppRelationShipList(v []*InsertSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) *InsertSwimmingLaneResponseBodyData {
	s.SwimmingLaneAppRelationShipList = v
	return s
}

func (s *InsertSwimmingLaneResponseBodyData) SetTag(v string) *InsertSwimmingLaneResponseBodyData {
	s.Tag = &v
	return s
}

func (s *InsertSwimmingLaneResponseBodyData) Validate() error {
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

type InsertSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList struct {
	// The ID of the application.
	//
	// example:
	//
	// bd170895-096c-4944-9007-d4582c77****
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
	// 88
	LaneId *int64 `json:"LaneId,omitempty" xml:"LaneId,omitempty"`
	// The association rule.
	//
	// example:
	//
	// dubbo
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s InsertSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) String() string {
	return dara.Prettify(s)
}

func (s InsertSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) GoString() string {
	return s.String()
}

func (s *InsertSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) GetAppId() *string {
	return s.AppId
}

func (s *InsertSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) GetAppName() *string {
	return s.AppName
}

func (s *InsertSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) GetLaneId() *int64 {
	return s.LaneId
}

func (s *InsertSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) GetRules() *string {
	return s.Rules
}

func (s *InsertSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) SetAppId(v string) *InsertSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList {
	s.AppId = &v
	return s
}

func (s *InsertSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) SetAppName(v string) *InsertSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList {
	s.AppName = &v
	return s
}

func (s *InsertSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) SetLaneId(v int64) *InsertSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList {
	s.LaneId = &v
	return s
}

func (s *InsertSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) SetRules(v string) *InsertSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList {
	s.Rules = &v
	return s
}

func (s *InsertSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) Validate() error {
	return dara.Validate(s)
}
