// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEndPointEventListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodes(v []*DescribeEndPointEventListResponseBodyNodes) *DescribeEndPointEventListResponseBody
	GetNodes() []*DescribeEndPointEventListResponseBodyNodes
	SetRequestId(v string) *DescribeEndPointEventListResponseBody
	GetRequestId() *string
}

type DescribeEndPointEventListResponseBody struct {
	Nodes []*DescribeEndPointEventListResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEndPointEventListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndPointEventListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEndPointEventListResponseBody) GetNodes() []*DescribeEndPointEventListResponseBodyNodes {
	return s.Nodes
}

func (s *DescribeEndPointEventListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEndPointEventListResponseBody) SetNodes(v []*DescribeEndPointEventListResponseBodyNodes) *DescribeEndPointEventListResponseBody {
	s.Nodes = v
	return s
}

func (s *DescribeEndPointEventListResponseBody) SetRequestId(v string) *DescribeEndPointEventListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEndPointEventListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEndPointEventListResponseBodyNodes struct {
	EventDataItems []*DescribeEndPointEventListResponseBodyNodesEventDataItems `json:"EventDataItems,omitempty" xml:"EventDataItems,omitempty" type:"Repeated"`
	// example:
	//
	// testuserid
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeEndPointEventListResponseBodyNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndPointEventListResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *DescribeEndPointEventListResponseBodyNodes) GetEventDataItems() []*DescribeEndPointEventListResponseBodyNodesEventDataItems {
	return s.EventDataItems
}

func (s *DescribeEndPointEventListResponseBodyNodes) GetUserId() *string {
	return s.UserId
}

func (s *DescribeEndPointEventListResponseBodyNodes) SetEventDataItems(v []*DescribeEndPointEventListResponseBodyNodesEventDataItems) *DescribeEndPointEventListResponseBodyNodes {
	s.EventDataItems = v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodes) SetUserId(v string) *DescribeEndPointEventListResponseBodyNodes {
	s.UserId = &v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeEndPointEventListResponseBodyNodesEventDataItems struct {
	EventList []*DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
	// example:
	//
	// 1614936817
	Ts *int64 `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s DescribeEndPointEventListResponseBodyNodesEventDataItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndPointEventListResponseBodyNodesEventDataItems) GoString() string {
	return s.String()
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItems) GetEventList() []*DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList {
	return s.EventList
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItems) GetTs() *int64 {
	return s.Ts
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItems) SetEventList(v []*DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) *DescribeEndPointEventListResponseBodyNodesEventDataItems {
	s.EventList = v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItems) SetTs(v int64) *DescribeEndPointEventListResponseBodyNodesEventDataItems {
	s.Ts = &v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItems) Validate() error {
	return dara.Validate(s)
}

type DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList struct {
	// example:
	//
	// 开始发布
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// example:
	//
	// USER
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// example:
	//
	// 1614936817
	Ts *int64 `json:"Ts,omitempty" xml:"Ts,omitempty"`
	// example:
	//
	// 1614936817123
	TsInMs *string `json:"TsInMs,omitempty" xml:"TsInMs,omitempty"`
}

func (s DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) GoString() string {
	return s.String()
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) GetEventName() *string {
	return s.EventName
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) GetEventType() *string {
	return s.EventType
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) GetTs() *int64 {
	return s.Ts
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) GetTsInMs() *string {
	return s.TsInMs
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) SetEventName(v string) *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList {
	s.EventName = &v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) SetEventType(v string) *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList {
	s.EventType = &v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) SetTs(v int64) *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList {
	s.Ts = &v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) SetTsInMs(v string) *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList {
	s.TsInMs = &v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) Validate() error {
	return dara.Validate(s)
}
