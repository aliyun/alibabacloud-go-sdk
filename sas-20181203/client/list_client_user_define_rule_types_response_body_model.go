// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientUserDefineRuleTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListClientUserDefineRuleTypesResponseBody
	GetRequestId() *string
	SetUserDefineRuleTypes(v []*string) *ListClientUserDefineRuleTypesResponseBody
	GetUserDefineRuleTypes() []*string
}

type ListClientUserDefineRuleTypesResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 340D7FC4-D575-1661-8ACD-CFA7BE57****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array consisting of the rule types that are supported.
	UserDefineRuleTypes []*string `json:"UserDefineRuleTypes,omitempty" xml:"UserDefineRuleTypes,omitempty" type:"Repeated"`
}

func (s ListClientUserDefineRuleTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClientUserDefineRuleTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClientUserDefineRuleTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClientUserDefineRuleTypesResponseBody) GetUserDefineRuleTypes() []*string {
	return s.UserDefineRuleTypes
}

func (s *ListClientUserDefineRuleTypesResponseBody) SetRequestId(v string) *ListClientUserDefineRuleTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClientUserDefineRuleTypesResponseBody) SetUserDefineRuleTypes(v []*string) *ListClientUserDefineRuleTypesResponseBody {
	s.UserDefineRuleTypes = v
	return s
}

func (s *ListClientUserDefineRuleTypesResponseBody) Validate() error {
	return dara.Validate(s)
}
