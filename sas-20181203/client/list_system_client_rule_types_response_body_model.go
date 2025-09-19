// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemClientRuleTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSystemClientRuleTypesResponseBody
	GetRequestId() *string
	SetRuleTypes(v []*string) *ListSystemClientRuleTypesResponseBody
	GetRuleTypes() []*string
}

type ListSystemClientRuleTypesResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 79CFF74D-E967-5407-8A78-EE03B925****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the rule types.
	RuleTypes []*string `json:"RuleTypes,omitempty" xml:"RuleTypes,omitempty" type:"Repeated"`
}

func (s ListSystemClientRuleTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSystemClientRuleTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSystemClientRuleTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSystemClientRuleTypesResponseBody) GetRuleTypes() []*string {
	return s.RuleTypes
}

func (s *ListSystemClientRuleTypesResponseBody) SetRequestId(v string) *ListSystemClientRuleTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSystemClientRuleTypesResponseBody) SetRuleTypes(v []*string) *ListSystemClientRuleTypesResponseBody {
	s.RuleTypes = v
	return s
}

func (s *ListSystemClientRuleTypesResponseBody) Validate() error {
	return dara.Validate(s)
}
