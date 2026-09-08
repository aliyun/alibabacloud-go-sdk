// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPickOutboundNumbersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalledNumber(v string) *PickOutboundNumbersRequest
	GetCalledNumber() *string
	SetCount(v int32) *PickOutboundNumbersRequest
	GetCount() *int32
	SetInstanceId(v string) *PickOutboundNumbersRequest
	GetInstanceId() *string
	SetSkillGroupIdList(v string) *PickOutboundNumbersRequest
	GetSkillGroupIdList() *string
}

type PickOutboundNumbersRequest struct {
	// Called number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1388888****
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The desired quantity of selectable numbers to return. Default is 1.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// A collection of skill group IDs, formatted as a JSON array string. Each array element is a skill group ID. Numbers are associated with skill groups, and this parameter specifies from which skill groups to select numbers.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["skillgroup@ccc-test"]
	SkillGroupIdList *string `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty"`
}

func (s PickOutboundNumbersRequest) String() string {
	return dara.Prettify(s)
}

func (s PickOutboundNumbersRequest) GoString() string {
	return s.String()
}

func (s *PickOutboundNumbersRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *PickOutboundNumbersRequest) GetCount() *int32 {
	return s.Count
}

func (s *PickOutboundNumbersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PickOutboundNumbersRequest) GetSkillGroupIdList() *string {
	return s.SkillGroupIdList
}

func (s *PickOutboundNumbersRequest) SetCalledNumber(v string) *PickOutboundNumbersRequest {
	s.CalledNumber = &v
	return s
}

func (s *PickOutboundNumbersRequest) SetCount(v int32) *PickOutboundNumbersRequest {
	s.Count = &v
	return s
}

func (s *PickOutboundNumbersRequest) SetInstanceId(v string) *PickOutboundNumbersRequest {
	s.InstanceId = &v
	return s
}

func (s *PickOutboundNumbersRequest) SetSkillGroupIdList(v string) *PickOutboundNumbersRequest {
	s.SkillGroupIdList = &v
	return s
}

func (s *PickOutboundNumbersRequest) Validate() error {
	return dara.Validate(s)
}
