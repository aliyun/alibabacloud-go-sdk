// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSwimmingLaneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListSwimmingLaneResponseBody
	GetCode() *int32
	SetData(v []*ListSwimmingLaneResponseBodyData) *ListSwimmingLaneResponseBody
	GetData() []*ListSwimmingLaneResponseBodyData
	SetMessage(v string) *ListSwimmingLaneResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSwimmingLaneResponseBody
	GetRequestId() *string
}

type ListSwimmingLaneResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data that is returned.
	Data []*ListSwimmingLaneResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// C2CDBBF9-9C9A-5AA1-9F39-094ADEB3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSwimmingLaneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSwimmingLaneResponseBody) GoString() string {
	return s.String()
}

func (s *ListSwimmingLaneResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListSwimmingLaneResponseBody) GetData() []*ListSwimmingLaneResponseBodyData {
	return s.Data
}

func (s *ListSwimmingLaneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSwimmingLaneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSwimmingLaneResponseBody) SetCode(v int32) *ListSwimmingLaneResponseBody {
	s.Code = &v
	return s
}

func (s *ListSwimmingLaneResponseBody) SetData(v []*ListSwimmingLaneResponseBodyData) *ListSwimmingLaneResponseBody {
	s.Data = v
	return s
}

func (s *ListSwimmingLaneResponseBody) SetMessage(v string) *ListSwimmingLaneResponseBody {
	s.Message = &v
	return s
}

func (s *ListSwimmingLaneResponseBody) SetRequestId(v string) *ListSwimmingLaneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSwimmingLaneResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSwimmingLaneResponseBodyData struct {
	// Indicates whether the throttling rule is enabled.
	//
	// example:
	//
	// true
	EnableRules *bool `json:"EnableRules,omitempty" xml:"EnableRules,omitempty"`
	// The conditions.
	//
	// example:
	//
	// [{\\"condition\\":\\"AND\\",\\"enable\\":true,\\"path\\":\\"/\\",\\"priority\\":1,\\"restItems\\":[{\\"cond\\":\\"==\\",\\"datum\\":\\"value\\",\\"name\\":\\"tags\\",\\"operator\\":\\"rawvalue\\",\\"type\\":\\"header\\",\\"value\\":\\"value\\"}]}]
	EntryRule *string `json:"EntryRule,omitempty" xml:"EntryRule,omitempty"`
	// The ID of the lane group.
	//
	// example:
	//
	// 156
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the lane.
	//
	// example:
	//
	// 348
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the lane.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the microservices namespace.
	//
	// example:
	//
	// cn-hangzhou:pre2
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The expected tag.
	//
	// example:
	//
	// d0ad1052
	ScenarioSign *string `json:"ScenarioSign,omitempty" xml:"ScenarioSign,omitempty"`
	// The applications that are related to the lane.
	SwimmingLaneAppRelationShipList []*ListSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList `json:"SwimmingLaneAppRelationShipList,omitempty" xml:"SwimmingLaneAppRelationShipList,omitempty" type:"Repeated"`
	// The tag.
	//
	// example:
	//
	// 2cb6b8a
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListSwimmingLaneResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSwimmingLaneResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSwimmingLaneResponseBodyData) GetEnableRules() *bool {
	return s.EnableRules
}

func (s *ListSwimmingLaneResponseBodyData) GetEntryRule() *string {
	return s.EntryRule
}

func (s *ListSwimmingLaneResponseBodyData) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ListSwimmingLaneResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListSwimmingLaneResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListSwimmingLaneResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListSwimmingLaneResponseBodyData) GetScenarioSign() *string {
	return s.ScenarioSign
}

func (s *ListSwimmingLaneResponseBodyData) GetSwimmingLaneAppRelationShipList() []*ListSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList {
	return s.SwimmingLaneAppRelationShipList
}

func (s *ListSwimmingLaneResponseBodyData) GetTag() *string {
	return s.Tag
}

func (s *ListSwimmingLaneResponseBodyData) SetEnableRules(v bool) *ListSwimmingLaneResponseBodyData {
	s.EnableRules = &v
	return s
}

func (s *ListSwimmingLaneResponseBodyData) SetEntryRule(v string) *ListSwimmingLaneResponseBodyData {
	s.EntryRule = &v
	return s
}

func (s *ListSwimmingLaneResponseBodyData) SetGroupId(v int64) *ListSwimmingLaneResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *ListSwimmingLaneResponseBodyData) SetId(v int64) *ListSwimmingLaneResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListSwimmingLaneResponseBodyData) SetName(v string) *ListSwimmingLaneResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListSwimmingLaneResponseBodyData) SetNamespaceId(v string) *ListSwimmingLaneResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *ListSwimmingLaneResponseBodyData) SetScenarioSign(v string) *ListSwimmingLaneResponseBodyData {
	s.ScenarioSign = &v
	return s
}

func (s *ListSwimmingLaneResponseBodyData) SetSwimmingLaneAppRelationShipList(v []*ListSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) *ListSwimmingLaneResponseBodyData {
	s.SwimmingLaneAppRelationShipList = v
	return s
}

func (s *ListSwimmingLaneResponseBodyData) SetTag(v string) *ListSwimmingLaneResponseBodyData {
	s.Tag = &v
	return s
}

func (s *ListSwimmingLaneResponseBodyData) Validate() error {
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

type ListSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList struct {
	// The ID of the application.
	//
	// example:
	//
	// 3b615783-01f1-4569-baa3-cb71bdb6****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// prod-app-58846
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Additional information.
	//
	// example:
	//
	// edas-canary
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The ID of the lane.
	//
	// example:
	//
	// 348
	LaneId *int64 `json:"LaneId,omitempty" xml:"LaneId,omitempty"`
	// The association rule.
	//
	// example:
	//
	// ""
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s ListSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) String() string {
	return dara.Prettify(s)
}

func (s ListSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) GoString() string {
	return s.String()
}

func (s *ListSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) GetAppId() *string {
	return s.AppId
}

func (s *ListSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) GetAppName() *string {
	return s.AppName
}

func (s *ListSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) GetExtra() *string {
	return s.Extra
}

func (s *ListSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) GetLaneId() *int64 {
	return s.LaneId
}

func (s *ListSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) GetRules() *string {
	return s.Rules
}

func (s *ListSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) SetAppId(v string) *ListSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList {
	s.AppId = &v
	return s
}

func (s *ListSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) SetAppName(v string) *ListSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList {
	s.AppName = &v
	return s
}

func (s *ListSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) SetExtra(v string) *ListSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList {
	s.Extra = &v
	return s
}

func (s *ListSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) SetLaneId(v int64) *ListSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList {
	s.LaneId = &v
	return s
}

func (s *ListSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) SetRules(v string) *ListSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList {
	s.Rules = &v
	return s
}

func (s *ListSwimmingLaneResponseBodyDataSwimmingLaneAppRelationShipList) Validate() error {
	return dara.Validate(s)
}
