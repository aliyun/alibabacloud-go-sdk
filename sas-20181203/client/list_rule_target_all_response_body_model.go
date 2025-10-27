// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRuleTargetAllResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListRuleTargetAllResponseBody
	GetRequestId() *string
	SetRuleTargetList(v []*ListRuleTargetAllResponseBodyRuleTargetList) *ListRuleTargetAllResponseBody
	GetRuleTargetList() []*ListRuleTargetAllResponseBodyRuleTargetList
}

type ListRuleTargetAllResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 30CBF632-109F-596F-97F2-451C8B2A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the network objects.
	RuleTargetList []*ListRuleTargetAllResponseBodyRuleTargetList `json:"RuleTargetList,omitempty" xml:"RuleTargetList,omitempty" type:"Repeated"`
}

func (s ListRuleTargetAllResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRuleTargetAllResponseBody) GoString() string {
	return s.String()
}

func (s *ListRuleTargetAllResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRuleTargetAllResponseBody) GetRuleTargetList() []*ListRuleTargetAllResponseBodyRuleTargetList {
	return s.RuleTargetList
}

func (s *ListRuleTargetAllResponseBody) SetRequestId(v string) *ListRuleTargetAllResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRuleTargetAllResponseBody) SetRuleTargetList(v []*ListRuleTargetAllResponseBodyRuleTargetList) *ListRuleTargetAllResponseBody {
	s.RuleTargetList = v
	return s
}

func (s *ListRuleTargetAllResponseBody) Validate() error {
	if s.RuleTargetList != nil {
		for _, item := range s.RuleTargetList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRuleTargetAllResponseBodyRuleTargetList struct {
	// The ID of the network object.
	//
	// example:
	//
	// 301944
	TargetId *int64 `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the network object.
	//
	// example:
	//
	// source-test-obj-xFKcx
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The type of the object. Valid values:
	//
	// 	- IMAGE
	//
	// example:
	//
	// IMAGE
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListRuleTargetAllResponseBodyRuleTargetList) String() string {
	return dara.Prettify(s)
}

func (s ListRuleTargetAllResponseBodyRuleTargetList) GoString() string {
	return s.String()
}

func (s *ListRuleTargetAllResponseBodyRuleTargetList) GetTargetId() *int64 {
	return s.TargetId
}

func (s *ListRuleTargetAllResponseBodyRuleTargetList) GetTargetName() *string {
	return s.TargetName
}

func (s *ListRuleTargetAllResponseBodyRuleTargetList) GetTargetType() *string {
	return s.TargetType
}

func (s *ListRuleTargetAllResponseBodyRuleTargetList) SetTargetId(v int64) *ListRuleTargetAllResponseBodyRuleTargetList {
	s.TargetId = &v
	return s
}

func (s *ListRuleTargetAllResponseBodyRuleTargetList) SetTargetName(v string) *ListRuleTargetAllResponseBodyRuleTargetList {
	s.TargetName = &v
	return s
}

func (s *ListRuleTargetAllResponseBodyRuleTargetList) SetTargetType(v string) *ListRuleTargetAllResponseBodyRuleTargetList {
	s.TargetType = &v
	return s
}

func (s *ListRuleTargetAllResponseBodyRuleTargetList) Validate() error {
	return dara.Validate(s)
}
