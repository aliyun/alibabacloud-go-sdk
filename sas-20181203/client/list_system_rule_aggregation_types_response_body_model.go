// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemRuleAggregationTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAggregationTypeList(v []*ListSystemRuleAggregationTypesResponseBodyAggregationTypeList) *ListSystemRuleAggregationTypesResponseBody
	GetAggregationTypeList() []*ListSystemRuleAggregationTypesResponseBodyAggregationTypeList
	SetRequestId(v string) *ListSystemRuleAggregationTypesResponseBody
	GetRequestId() *string
}

type ListSystemRuleAggregationTypesResponseBody struct {
	// An array that consists of the aggregation types.
	AggregationTypeList []*ListSystemRuleAggregationTypesResponseBodyAggregationTypeList `json:"AggregationTypeList,omitempty" xml:"AggregationTypeList,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 1427F3BE-8A7E-57F9-BD4E-590B00D2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSystemRuleAggregationTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSystemRuleAggregationTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSystemRuleAggregationTypesResponseBody) GetAggregationTypeList() []*ListSystemRuleAggregationTypesResponseBodyAggregationTypeList {
	return s.AggregationTypeList
}

func (s *ListSystemRuleAggregationTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSystemRuleAggregationTypesResponseBody) SetAggregationTypeList(v []*ListSystemRuleAggregationTypesResponseBodyAggregationTypeList) *ListSystemRuleAggregationTypesResponseBody {
	s.AggregationTypeList = v
	return s
}

func (s *ListSystemRuleAggregationTypesResponseBody) SetRequestId(v string) *ListSystemRuleAggregationTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSystemRuleAggregationTypesResponseBody) Validate() error {
	if s.AggregationTypeList != nil {
		for _, item := range s.AggregationTypeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSystemRuleAggregationTypesResponseBodyAggregationTypeList struct {
	// The ID of the aggregation type.
	//
	// example:
	//
	// 1
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the aggregation type.
	//
	// example:
	//
	// Remote control\\*\\*\\*\\*
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListSystemRuleAggregationTypesResponseBodyAggregationTypeList) String() string {
	return dara.Prettify(s)
}

func (s ListSystemRuleAggregationTypesResponseBodyAggregationTypeList) GoString() string {
	return s.String()
}

func (s *ListSystemRuleAggregationTypesResponseBodyAggregationTypeList) GetId() *int32 {
	return s.Id
}

func (s *ListSystemRuleAggregationTypesResponseBodyAggregationTypeList) GetName() *string {
	return s.Name
}

func (s *ListSystemRuleAggregationTypesResponseBodyAggregationTypeList) SetId(v int32) *ListSystemRuleAggregationTypesResponseBodyAggregationTypeList {
	s.Id = &v
	return s
}

func (s *ListSystemRuleAggregationTypesResponseBodyAggregationTypeList) SetName(v string) *ListSystemRuleAggregationTypesResponseBodyAggregationTypeList {
	s.Name = &v
	return s
}

func (s *ListSystemRuleAggregationTypesResponseBodyAggregationTypeList) Validate() error {
	return dara.Validate(s)
}
