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
	// The list of query criteria.
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
	if s.CriteriaList != nil {
		for _, item := range s.CriteriaList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSasContainerWebDefenseRuleCriteriaResponseBodyCriteriaList struct {
	// The name of the query criterion.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the query criterion. Valid values:
	//
	// - **input**: Input type. You must manually enter the query content when using this query criterion.
	//
	// - **select**: Selection type. You must select a value from the options list (the **Values*	- field) when using this query criterion.
	//
	// example:
	//
	// select
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The available options when **Type*	- (the type of the query criterion) is set to **select*	- (selection type).
	//
	// > When **Type*	- (the type of the query criterion) is set to **input*	- (input type), this parameter returns an empty value.
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
