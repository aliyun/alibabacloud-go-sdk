// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccountFactoryBaselineItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineItems(v []*ListAccountFactoryBaselineItemsResponseBodyBaselineItems) *ListAccountFactoryBaselineItemsResponseBody
	GetBaselineItems() []*ListAccountFactoryBaselineItemsResponseBodyBaselineItems
	SetNextToken(v string) *ListAccountFactoryBaselineItemsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAccountFactoryBaselineItemsResponseBody
	GetRequestId() *string
}

type ListAccountFactoryBaselineItemsResponseBody struct {
	// The baseline items.
	BaselineItems []*ListAccountFactoryBaselineItemsResponseBodyBaselineItems `json:"BaselineItems,omitempty" xml:"BaselineItems,omitempty" type:"Repeated"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAACDGQdAEX3m42z3sQ+f3VTK2Xr2DzYbz/SAfc/zJRqod
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B40D73D8-76AC-5D3C-AC63-4FC8AFCE6671
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAccountFactoryBaselineItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAccountFactoryBaselineItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccountFactoryBaselineItemsResponseBody) GetBaselineItems() []*ListAccountFactoryBaselineItemsResponseBodyBaselineItems {
	return s.BaselineItems
}

func (s *ListAccountFactoryBaselineItemsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAccountFactoryBaselineItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAccountFactoryBaselineItemsResponseBody) SetBaselineItems(v []*ListAccountFactoryBaselineItemsResponseBodyBaselineItems) *ListAccountFactoryBaselineItemsResponseBody {
	s.BaselineItems = v
	return s
}

func (s *ListAccountFactoryBaselineItemsResponseBody) SetNextToken(v string) *ListAccountFactoryBaselineItemsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAccountFactoryBaselineItemsResponseBody) SetRequestId(v string) *ListAccountFactoryBaselineItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccountFactoryBaselineItemsResponseBody) Validate() error {
	if s.BaselineItems != nil {
		for _, item := range s.BaselineItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAccountFactoryBaselineItemsResponseBodyBaselineItems struct {
	// The dependency of the baseline item.
	DependsOn []*ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn `json:"DependsOn,omitempty" xml:"DependsOn,omitempty" type:"Repeated"`
	// The description of the baseline item.
	//
	// example:
	//
	// Notification.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the baseline item.
	//
	// example:
	//
	// ACS-BP_ACCOUNT_FACTORY_ACCOUNT_NOTIFICATION
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the baseline item.
	//
	// example:
	//
	// AccountFactory
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the baseline item.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAccountFactoryBaselineItemsResponseBodyBaselineItems) String() string {
	return dara.Prettify(s)
}

func (s ListAccountFactoryBaselineItemsResponseBodyBaselineItems) GoString() string {
	return s.String()
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItems) GetDependsOn() []*ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn {
	return s.DependsOn
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItems) GetDescription() *string {
	return s.Description
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItems) GetName() *string {
	return s.Name
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItems) GetType() *string {
	return s.Type
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItems) GetVersion() *string {
	return s.Version
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItems) SetDependsOn(v []*ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn) *ListAccountFactoryBaselineItemsResponseBodyBaselineItems {
	s.DependsOn = v
	return s
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItems) SetDescription(v string) *ListAccountFactoryBaselineItemsResponseBodyBaselineItems {
	s.Description = &v
	return s
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItems) SetName(v string) *ListAccountFactoryBaselineItemsResponseBodyBaselineItems {
	s.Name = &v
	return s
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItems) SetType(v string) *ListAccountFactoryBaselineItemsResponseBodyBaselineItems {
	s.Type = &v
	return s
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItems) SetVersion(v string) *ListAccountFactoryBaselineItemsResponseBodyBaselineItems {
	s.Version = &v
	return s
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItems) Validate() error {
	if s.DependsOn != nil {
		for _, item := range s.DependsOn {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn struct {
	// The name of the baseline item.
	//
	// example:
	//
	// ACS-BP_ACCOUNT_FACTORY_VPC
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the baseline item.
	//
	// example:
	//
	// AccountFactory
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the baseline item.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn) String() string {
	return dara.Prettify(s)
}

func (s ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn) GoString() string {
	return s.String()
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn) GetName() *string {
	return s.Name
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn) GetType() *string {
	return s.Type
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn) GetVersion() *string {
	return s.Version
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn) SetName(v string) *ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn {
	s.Name = &v
	return s
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn) SetType(v string) *ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn {
	s.Type = &v
	return s
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn) SetVersion(v string) *ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn {
	s.Version = &v
	return s
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn) Validate() error {
	return dara.Validate(s)
}
