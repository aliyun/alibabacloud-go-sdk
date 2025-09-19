// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSasContainerWebDefenseRuleCriteriaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCriteriaList(v []*GetSasContainerWebDefenseRuleCriteriaResponseBodyCriteriaList) *GetSasContainerWebDefenseRuleCriteriaResponseBody
	GetCriteriaList() []*GetSasContainerWebDefenseRuleCriteriaResponseBodyCriteriaList
	SetRequestId(v string) *GetSasContainerWebDefenseRuleCriteriaResponseBody
	GetRequestId() *string
}

type GetSasContainerWebDefenseRuleCriteriaResponseBody struct {
	// The search conditions.
	CriteriaList []*GetSasContainerWebDefenseRuleCriteriaResponseBodyCriteriaList `json:"CriteriaList,omitempty" xml:"CriteriaList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F8B6F758-BCD4-597A-8A2C-DA5A552C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSasContainerWebDefenseRuleCriteriaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSasContainerWebDefenseRuleCriteriaResponseBody) GoString() string {
	return s.String()
}

func (s *GetSasContainerWebDefenseRuleCriteriaResponseBody) GetCriteriaList() []*GetSasContainerWebDefenseRuleCriteriaResponseBodyCriteriaList {
	return s.CriteriaList
}

func (s *GetSasContainerWebDefenseRuleCriteriaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSasContainerWebDefenseRuleCriteriaResponseBody) SetCriteriaList(v []*GetSasContainerWebDefenseRuleCriteriaResponseBodyCriteriaList) *GetSasContainerWebDefenseRuleCriteriaResponseBody {
	s.CriteriaList = v
	return s
}

func (s *GetSasContainerWebDefenseRuleCriteriaResponseBody) SetRequestId(v string) *GetSasContainerWebDefenseRuleCriteriaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleCriteriaResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSasContainerWebDefenseRuleCriteriaResponseBodyCriteriaList struct {
	// The name of the search condition.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the search condition. Valid values:
	//
	// 	- **input**: You must manually enter the search condition.
	//
	// 	- **select**: You must select a search condition from the **Values*	- drop-down list.
	//
	// example:
	//
	// select
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The values of the search condition. This parameter is returned only if the value of **Type*	- is **select**.
	//
	// >  If the value of **Type*	- is **input**, the value of this parameter is empty.
	//
	// example:
	//
	// athor_bid
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s GetSasContainerWebDefenseRuleCriteriaResponseBodyCriteriaList) String() string {
	return dara.Prettify(s)
}

func (s GetSasContainerWebDefenseRuleCriteriaResponseBodyCriteriaList) GoString() string {
	return s.String()
}

func (s *GetSasContainerWebDefenseRuleCriteriaResponseBodyCriteriaList) GetName() *string {
	return s.Name
}

func (s *GetSasContainerWebDefenseRuleCriteriaResponseBodyCriteriaList) GetType() *string {
	return s.Type
}

func (s *GetSasContainerWebDefenseRuleCriteriaResponseBodyCriteriaList) GetValues() *string {
	return s.Values
}

func (s *GetSasContainerWebDefenseRuleCriteriaResponseBodyCriteriaList) SetName(v string) *GetSasContainerWebDefenseRuleCriteriaResponseBodyCriteriaList {
	s.Name = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleCriteriaResponseBodyCriteriaList) SetType(v string) *GetSasContainerWebDefenseRuleCriteriaResponseBodyCriteriaList {
	s.Type = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleCriteriaResponseBodyCriteriaList) SetValues(v string) *GetSasContainerWebDefenseRuleCriteriaResponseBodyCriteriaList {
	s.Values = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleCriteriaResponseBodyCriteriaList) Validate() error {
	return dara.Validate(s)
}
