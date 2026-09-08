// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentifyModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListIdentifyModelsResponseBody
	GetRequestId() *string
	SetRuleList(v []*ListIdentifyModelsResponseBodyRuleList) *ListIdentifyModelsResponseBody
	GetRuleList() []*ListIdentifyModelsResponseBodyRuleList
}

type ListIdentifyModelsResponseBody struct {
	// example:
	//
	// 7C6D8E9F-1234-5678-ABCD-0123456789AB
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleList  []*ListIdentifyModelsResponseBodyRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
}

func (s ListIdentifyModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIdentifyModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIdentifyModelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIdentifyModelsResponseBody) GetRuleList() []*ListIdentifyModelsResponseBodyRuleList {
	return s.RuleList
}

func (s *ListIdentifyModelsResponseBody) SetRequestId(v string) *ListIdentifyModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIdentifyModelsResponseBody) SetRuleList(v []*ListIdentifyModelsResponseBodyRuleList) *ListIdentifyModelsResponseBody {
	s.RuleList = v
	return s
}

func (s *ListIdentifyModelsResponseBody) Validate() error {
	if s.RuleList != nil {
		for _, item := range s.RuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIdentifyModelsResponseBodyRuleList struct {
	// example:
	//
	// 1001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 手机号
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListIdentifyModelsResponseBodyRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListIdentifyModelsResponseBodyRuleList) GoString() string {
	return s.String()
}

func (s *ListIdentifyModelsResponseBodyRuleList) GetId() *int64 {
	return s.Id
}

func (s *ListIdentifyModelsResponseBodyRuleList) GetName() *string {
	return s.Name
}

func (s *ListIdentifyModelsResponseBodyRuleList) SetId(v int64) *ListIdentifyModelsResponseBodyRuleList {
	s.Id = &v
	return s
}

func (s *ListIdentifyModelsResponseBodyRuleList) SetName(v string) *ListIdentifyModelsResponseBodyRuleList {
	s.Name = &v
	return s
}

func (s *ListIdentifyModelsResponseBodyRuleList) Validate() error {
	return dara.Validate(s)
}
